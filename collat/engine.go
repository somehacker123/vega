package collat

import (
	"fmt"
	"sync"
	"time"

	"code.vegaprotocol.io/vega/events"
	"code.vegaprotocol.io/vega/logging"
	types "code.vegaprotocol.io/vega/proto"

	"github.com/pkg/errors"
)

const (
	initialAccountSize = 4096
	// use weird character here, maybe non-displayable ones in the future
	// if needed
	systemOwner = "*"
	noMarket    = "!"
)

var (
	// ErrSystemAccountsMissing signals that a system account is missing, which may means that the
	// collateral engine have not been initialized properly
	ErrSystemAccountsMissing = errors.New("system accounts missing for collateral engine to work")
	// ErrTraderAccountsMissing signals that the accounts for this trader do not exists
	ErrTraderAccountsMissing = errors.New("trader accounts missing, cannot collect")
	// ErrAccountDoesNotExist signals that an account par of a transfer do not exists
	ErrAccountDoesNotExist = errors.New("account do not exists")
)

// AccountBuffer ...
//go:generate go run github.com/golang/mock/mockgen -destination mocks/account_buffer_mock.go -package mocks code.vegaprotocol.io/vega/collat AccountBuffer
type AccountBuffer interface {
	Add(types.Account)
}

// Engine is handling the power of the collateral
type Engine struct {
	Config
	log   *logging.Logger
	cfgMu sync.Mutex

	// map of trader ID's to map of account types + account ID's
	// traderAccounts map[string]map[types.AccountType]map[string]string // by trader, type, and asset
	// marketAccounts map[types.AccountType]map[string]string            // by type and asset

	accs map[string]*types.Account
	buf  AccountBuffer
	// cool be a unix.Time but storing it like this allow us to now time.UnixNano() all the time
	currentTime int64

	idbuf []byte
}

// New instantiates a new collateral engine
func New(log *logging.Logger, conf Config, buf AccountBuffer, now time.Time) (*Engine, error) {
	// setup logger
	log = log.Named(namedLogger)
	log.SetLevel(conf.Level.Get())
	return &Engine{
		log:         log,
		Config:      conf,
		accs:        make(map[string]*types.Account, initialAccountSize),
		buf:         buf,
		currentTime: now.UnixNano(),
		idbuf:       make([]byte, 256),
	}, nil
}

// OnChainTimeUpdate is used to be specified as a callback in over services
// in order to be called when the chain time is updated (basically EndBlock)
func (e *Engine) OnChainTimeUpdate(t time.Time) {
	e.currentTime = t.UnixNano()
}

// ReloadConf updates the internal configuration of the collateral engine
func (e *Engine) ReloadConf(cfg Config) {
	e.log.Info("reloading configuration")
	if e.log.GetLevel() != cfg.Level.Get() {
		e.log.Info("updating log level",
			logging.String("old", e.log.GetLevel().String()),
			logging.String("new", cfg.Level.String()),
		)
		e.log.SetLevel(cfg.Level.Get())
	}

	e.cfgMu.Lock()
	e.Config = cfg
	e.cfgMu.Unlock()
}

// this func uses named returns because it makes body of the func look clearer
func (e *Engine) getSystemAccounts(marketID, asset string) (settle, insurance *types.Account, err error) {

	insID := e.accountID(marketID, systemOwner, asset, types.AccountType_INSURANCE)
	setID := e.accountID(marketID, systemOwner, asset, types.AccountType_SETTLEMENT)

	var ok bool
	if insurance, ok = e.accs[insID]; !ok {
		if e.log.GetLevel() == logging.DebugLevel {
			e.log.Debug("missing system account",
				logging.String("asset", asset),
				logging.String("id", insID),
				logging.String("market", marketID),
			)
		}
		err = ErrSystemAccountsMissing
		return
	}

	if settle, ok = e.accs[setID]; !ok {
		if e.log.GetLevel() == logging.DebugLevel {
			e.log.Debug("missing system account",
				logging.String("asset", asset),
				logging.String("id", setID),
				logging.String("market", marketID),
			)
		}
		err = ErrSystemAccountsMissing
	}

	return
}

// AddTraderToMarket - when a new trader enters a market, ensure general + margin accounts both exist
func (e *Engine) AddTraderToMarket(marketID, traderID, asset string) error {
	// accountID(marketID, traderID, asset string, ty types.AccountType) accountIDT
	genID := e.accountID("", traderID, asset, types.AccountType_GENERAL)
	marginID := e.accountID(marketID, traderID, asset, types.AccountType_MARGIN)
	_, err := e.GetAccountByID(genID)
	if err != nil {
		e.log.Error(
			"Trader doesn't have a general account somehow?",
			logging.String("trader-id", traderID))
		return ErrTraderAccountsMissing
	}
	_, err = e.GetAccountByID(marginID)
	if err != nil {
		e.log.Error(
			"Trader doesn't have a margin account somehow?",
			logging.String("trader-id", traderID),
			logging.String("Market", marketID))
		return ErrTraderAccountsMissing
	}

	return nil
}

// Transfer will process the list of transfer instructed by other engines
func (e *Engine) Transfer(marketID string, transfers []*types.Transfer) ([]*types.TransferResponse, error) {
	if len(transfers) == 0 {
		return nil, nil
	}
	return nil, nil
	// if isSettle(transfers[0]) {
	// 	return e.collect(marketID, transfers)
	// }
	// // this is a balance top-up or some other thing we haven't implemented yet
	// return nil, nil
}

// MarkToMarket will run the mark to market settlement over a given set of positions
// return ledger move stuff here, too (seperate return value, because we need to stream those)
func (e *Engine) MarkToMarket(marketID string, transfers []events.Transfer) ([]events.Margin, []*types.TransferResponse, error) {
	// stop immediately if there aren't any transfers, channels are closed
	if len(transfers) == 0 {
		return nil, nil, nil
	}
	marginEvts := make([]events.Margin, 0, len(transfers))
	responses := make([]*types.TransferResponse, 0, len(transfers))
	asset := transfers[0].Transfer().Amount.Asset
	// This is where we'll implement everything
	settle, insurance, err := e.getSystemAccounts(marketID, asset)
	if err != nil {
		e.log.Error(
			"Failed to get system accounts required for MTM settlement",
			logging.Error(err),
		)
		return nil, nil, err
	}
	// get the component that calculates the loss socialisation etc... if needed
	distr := &distributor{}
	for _, evt := range transfers {
		transfer := evt.Transfer()
		loss := isLoss(transfer)
		marginEvt := &marginUpdate{
			MarketPosition: evt,
			asset:          transfer.Amount.Asset,
			marketID:       settle.MarketID,
		}
		req, err := e.getTransferRequest(transfer, settle, insurance, marginEvt)
		if err != nil {
			e.log.Error(
				"Failed to build transfer request for event",
				logging.Error(err),
			)
			return nil, nil, err
		}
		// set the amount (this can change the req.Amount value if we entered loss socialisation
		distr.amountCB(req, loss)
		res, err := e.getLedgerEntries(req)
		if err != nil {
			e.log.Error(
				"Failed to transfer funds",
				logging.Error(err),
			)
			return nil, nil, err
		}
		// if this is a loss, we want to update the delta, too
		if loss {
			distr.registerTransfer(res)
		} else {
			// getLedgerEntries updates the from accounts, so losses are handled fine there
			// but the to account isn't updated (losses are deposited in temporary settlement account)
			// but wins are paid out to trader accounts, so we need to update the balance there
			for _, bal := range res.Balances {
				if err := e.UpdateBalance(bal.Account.Id, bal.Balance); err != nil {
					e.log.Error(
						"Could not update the target account in transfer",
						logging.String("account-id", bal.Account.Id),
						logging.Error(err),
					)
					return nil, nil, err
				}
			}
		}
		responses = append(responses, res)
		marginEvts = append(marginEvts, marginEvt)
	}
	return marginEvts, responses, nil
}

func isLoss(t *types.Transfer) bool {
	return (t.Type == types.TransferType_LOSS || t.Type == types.TransferType_MTM_LOSS)
}

// getTransferRequest builds the request, and sets the required accounts based on the type of the Transfer argument
func (e *Engine) getTransferRequest(p *types.Transfer, settle, insurance *types.Account, mEvt *marginUpdate) (*types.TransferRequest, error) {
	asset := p.Amount.Asset

	// the accounts for the trader we need
	traderAccs := map[types.AccountType]*types.Account{
		types.AccountType_MARGIN:  nil,
		types.AccountType_GENERAL: nil,
	}
	for t := range traderAccs {
		acc, err := e.GetAccountByID(e.accountID(settle.MarketID, p.Owner, asset, t))
		if err != nil {
			e.log.Error(
				"Failed to get the required trader accounts",
				logging.String("owner-id", p.Owner),
				logging.String("market-id", settle.MarketID),
				logging.String("account-type", t.String()),
				logging.Error(err),
			)
			return nil, err
		}
		traderAccs[t] = acc
	}
	// the return event should contain these accounts
	mEvt.margin = traderAccs[types.AccountType_MARGIN]
	mEvt.general = traderAccs[types.AccountType_GENERAL]
	// final settle, or MTM settle, makes no difference, it's win/loss still
	if p.Type == types.TransferType_LOSS || p.Type == types.TransferType_MTM_LOSS {
		// losses are collected first from the margin account, then the general account, and finally
		// taken out of the insurance pool
		req := types.TransferRequest{
			FromAccount: []*types.Account{
				traderAccs[types.AccountType_MARGIN],
				traderAccs[types.AccountType_GENERAL],
				insurance,
			},
			ToAccount: []*types.Account{
				settle,
			},
			Amount:    uint64(-p.Amount.Amount) * p.Size,
			MinAmount: 0,     // default value, but keep it here explicitly
			Asset:     asset, // TBC
			Reference: p.Type.String(),
		}
		return &req, nil
	}
	if p.Type == types.TransferType_WIN || p.Type == types.TransferType_MTM_WIN {
		// the insurance pool in the FromAccount is not used ATM (losses should fully cover wins
		// or the insurance pool has already been drained).
		return &types.TransferRequest{
			FromAccount: []*types.Account{
				settle,
				insurance,
			},
			ToAccount: []*types.Account{
				traderAccs[types.AccountType_MARGIN],
			},
			Amount:    uint64(p.Amount.Amount) * p.Size,
			MinAmount: 0,     // default value, but keep it here explicitly
			Asset:     asset, // TBC
			Reference: p.Type.String(),
		}, nil
	}

	// just in case...
	if p.Size == 0 {
		p.Size = 1
	}
	if p.Type == types.TransferType_MARGIN_LOW {
		return &types.TransferRequest{
			FromAccount: []*types.Account{
				traderAccs[types.AccountType_GENERAL],
			},
			ToAccount: []*types.Account{
				traderAccs[types.AccountType_MARGIN],
			},
			Amount:    uint64(p.Amount.Amount) * p.Size,
			MinAmount: uint64(p.Amount.MinAmount),
			Asset:     asset,
			Reference: p.Type.String(),
		}, nil
	}
	return &types.TransferRequest{
		FromAccount: []*types.Account{
			traderAccs[types.AccountType_MARGIN],
		},
		ToAccount: []*types.Account{
			traderAccs[types.AccountType_GENERAL],
		},
		Amount:    uint64(p.Amount.Amount) * p.Size,
		MinAmount: uint64(p.Amount.MinAmount),
		Asset:     asset,
		Reference: p.Type.String(),
	}, nil
}

// this builds a TransferResponse for a specific request, we collect all of them and aggregate
func (e *Engine) getLedgerEntries(req *types.TransferRequest) (*types.TransferResponse, error) {
	ret := types.TransferResponse{
		Transfers: []*types.LedgerEntry{},
		Balances:  make([]*types.TransferBalance, 0, len(req.ToAccount)),
	}
	for _, t := range req.ToAccount {
		ret.Balances = append(ret.Balances, &types.TransferBalance{
			Account: t,
		})
	}
	amount := int64(req.Amount)
	for _, acc := range req.FromAccount {
		// give each to account an equal share
		parts := amount / int64(len(req.ToAccount))
		// add remaining pennies to last ledger movement
		remainder := amount % int64(len(req.ToAccount))
		var (
			to *types.TransferBalance
			lm *types.LedgerEntry
		)
		// either the account contains enough, or we're having to access insurance pool money
		if acc.Balance >= amount {
			acc.Balance -= amount
			if err := e.IncrementBalance(acc.Id, -amount); err != nil {
				e.log.Error(
					"Failed to update balance for account",
					logging.String("account-id", acc.Id),
					logging.Int64("balance", acc.Balance),
					logging.Error(err),
				)
				return nil, err
			}
			for _, to = range ret.Balances {
				lm = &types.LedgerEntry{
					FromAccount: acc.Id,
					ToAccount:   to.Account.Id,
					Amount:      parts,
					Reference:   req.Reference,
					Type:        "settlement",
					Timestamp:   e.currentTime,
				}
				ret.Transfers = append(ret.Transfers, lm)
				to.Balance += parts
				to.Account.Balance += parts
			}
			// add remainder
			if remainder > 0 {
				lm.Amount += remainder
				to.Balance += remainder
				to.Account.Balance += remainder
			}
			return &ret, nil
		}
		if acc.Balance > 0 {
			amount -= acc.Balance
			// partial amount resolves differently
			parts = acc.Balance / int64(len(req.ToAccount))
			if err := e.UpdateBalance(acc.Id, 0); err != nil {
				e.log.Error(
					"Failed to set balance of account to 0",
					logging.String("account-id", acc.Id),
					logging.Error(err),
				)
				return nil, err
			}
			acc.Balance = 0
			for _, to = range ret.Balances {
				lm = &types.LedgerEntry{
					FromAccount: acc.Id,
					ToAccount:   to.Account.Id,
					Amount:      parts,
					Reference:   req.Reference,
					Type:        "settlement",
					Timestamp:   e.currentTime,
				}
				ret.Transfers = append(ret.Transfers, lm)
				to.Account.Balance += parts
				to.Balance += parts
			}
		}
		if amount == 0 {
			break
		}
	}
	return &ret, nil
}

// ClearMarket will remove all monies or accounts for parties allocated for a market (margin accounts)
// when the market reach end of life (maturity)
func (e *Engine) ClearMarket(mktID, asset string, parties []string) ([]*types.TransferResponse, error) {
	// create a transfer request that we will reuse all the time in order to make allocations smaller
	req := &types.TransferRequest{
		FromAccount: make([]*types.Account, 1),
		ToAccount:   make([]*types.Account, 1),
		Asset:       asset,
	}

	// assume we have as much transfer response than parties
	resps := make([]*types.TransferResponse, 0, len(parties))

	for _, v := range parties {
		marginAcc, err := e.GetAccountByID(e.accountID(mktID, v, asset, types.AccountType_MARGIN))
		if err != nil {
			e.log.Error(
				"Failed to get the margin account",
				logging.String("trader-id", v),
				logging.String("market-id", mktID),
				logging.String("asset", asset),
				logging.Error(err))
			// just try to do other traders
			continue
		}

		generalAcc, err := e.GetAccountByID(e.accountID("", v, asset, types.AccountType_GENERAL))
		if err != nil {
			e.log.Error(
				"Failed to get the general account",
				logging.String("trader-id", v),
				logging.String("market-id", mktID),
				logging.String("asset", asset),
				logging.Error(err))
			// just try to do other traders
			continue
		}

		req.FromAccount[0] = marginAcc
		req.ToAccount[0] = generalAcc
		req.Amount = uint64(marginAcc.Balance)

		if e.log.GetLevel() == logging.DebugLevel {
			e.log.Debug("Clearing trader margin account",
				logging.String("market-id", mktID),
				logging.String("asset", asset),
				logging.String("trader-id", v),
				logging.Int64("margin-before", marginAcc.Balance),
				logging.Int64("general-before", generalAcc.Balance),
				logging.Int64("general-after", generalAcc.Balance+marginAcc.Balance))
		}

		ledgerEntries, err := e.getLedgerEntries(req)
		if err != nil {
			e.log.Error(
				"Failed to move monies from margin to genral account",
				logging.String("trader-id", v),
				logging.String("market-id", mktID),
				logging.String("asset", asset),
				logging.Error(err))
			// just try to do other traders
			continue
		}

		for _, v := range ledgerEntries.Transfers {
			// increment the to account
			if err := e.IncrementBalance(v.ToAccount, v.Amount); err != nil {
				e.log.Error(
					"Failed to increment balance for account",
					logging.String("account-id", v.ToAccount),
					logging.Int64("amount", v.Amount),
					logging.Error(err),
				)
				return nil, err
			}
		}

		resps = append(resps, ledgerEntries)
	}

	return resps, nil
}

// CreateTraderAccount will create trader accounts for a given market
// basically one account per market, per asset for each trader
func (e *Engine) CreateTraderAccount(traderID, marketID, asset string) (marginID, generalID string) {
	// first margin account
	marginID = e.accountID(marketID, traderID, asset, types.AccountType_MARGIN)
	_, ok := e.accs[marginID]
	if !ok {
		acc := &types.Account{
			Id:       marginID,
			Asset:    asset,
			MarketID: marketID,
			Balance:  0,
			Owner:    traderID,
			Type:     types.AccountType_MARGIN,
		}
		e.accs[marginID] = acc
		e.buf.Add(*acc)
	}

	generalID = e.accountID(noMarket, traderID, asset, types.AccountType_GENERAL)
	_, ok = e.accs[generalID]
	if !ok {
		acc := &types.Account{
			Id:       generalID,
			Asset:    asset,
			MarketID: noMarket,
			Balance:  0,
			Owner:    traderID,
			Type:     types.AccountType_GENERAL,
		}
		e.accs[generalID] = acc
		e.buf.Add(*acc)
	}

	return
}

// RemoveDistressed will remove all distressed trader in the event positions
// for a given market and asset
func (e *Engine) RemoveDistressed(traders []events.MarketPosition, marketID, asset string) (*types.TransferResponse, error) {
	tl := len(traders)
	if tl == 0 {
		return nil, nil
	}
	// insurance account is the one we're after
	_, ins, err := e.getSystemAccounts(marketID, asset)
	if err != nil {
		return nil, err
	}
	resp := types.TransferResponse{
		Transfers: make([]*types.LedgerEntry, 0, tl),
	}
	for _, trader := range traders {
		acc, err := e.GetAccountByID(e.accountID(marketID, trader.Party(), asset, types.AccountType_MARGIN))
		if err != nil {
			return nil, err
		}
		// only create a ledger move if the balance is greater than zero
		if acc.Balance > 0 {
			resp.Transfers = append(resp.Transfers, &types.LedgerEntry{
				FromAccount: acc.Id,
				ToAccount:   ins.Id,
				Amount:      acc.Balance,
				Reference:   "close-out distressed",
				Type:        "", // @TODO determine this value
				Timestamp:   e.currentTime,
			})
			if err := e.IncrementBalance(ins.Id, acc.Balance); err != nil {
				return nil, err
			}
			if err := e.UpdateBalance(acc.Id, 0); err != nil {
				return nil, err
			}
		}
		if err := e.removeAccount(acc.Id); err != nil {
			return nil, err
		}
	}
	return &resp, nil
}

// CreateMarketAccounts will create all required accounts for a market once
// a new market is accepted through the network
func (e *Engine) CreateMarketAccounts(marketID, asset string, insurance int64) (insuranceID, settleID string) {
	insuranceID = e.accountID(marketID, "", asset, types.AccountType_INSURANCE)
	_, ok := e.accs[insuranceID]
	if !ok {
		insAcc := &types.Account{
			Id:       insuranceID,
			Asset:    asset,
			Owner:    systemOwner,
			Balance:  insurance,
			MarketID: marketID,
			Type:     types.AccountType_INSURANCE,
		}
		e.accs[insuranceID] = insAcc
		e.buf.Add(*insAcc)

	}
	settleID = e.accountID(marketID, "", asset, types.AccountType_SETTLEMENT)
	_, ok = e.accs[settleID]
	if !ok {
		setAcc := &types.Account{
			Id:       settleID,
			Asset:    asset,
			Owner:    systemOwner,
			Balance:  0,
			MarketID: marketID,
			Type:     types.AccountType_SETTLEMENT,
		}
		e.accs[settleID] = setAcc
		e.buf.Add(*setAcc)
	}

	return
}

// Withdraw will remove the specified amount from the trader
// general account
func (e *Engine) Withdraw(partyID, asset string, amount uint64) error {
	acc, err := e.GetAccountByID(e.accountID("", partyID, asset, types.AccountType_GENERAL))
	if err != nil {
		return ErrAccountDoesNotExist
	}

	// check we have more money than required to withdraw
	if uint64(acc.Balance) < amount {
		// if we have less balance than required to withdraw, just set it to 0
		// and return an error
		if err := e.UpdateBalance(acc.Id, 0); err != nil {
			return err
		}
		return fmt.Errorf("withdraw error, required=%v, available=%v", amount, acc.Balance)
	}

	if err := e.IncrementBalance(acc.Id, -int64(amount)); err != nil {
		return err
	}
	return nil
}

// UpdateBalance will update the balance of a given account
func (e *Engine) UpdateBalance(id string, balance int64) error {
	acc, ok := e.accs[id]
	if !ok {
		return ErrAccountDoesNotExist
	}
	acc.Balance = balance
	e.buf.Add(*acc)
	return nil
}

// IncrementBalance will increment the balance of a given account
// using the given value
func (e *Engine) IncrementBalance(id string, inc int64) error {
	acc, ok := e.accs[id]
	if !ok {
		return ErrAccountDoesNotExist
	}
	acc.Balance += inc
	e.buf.Add(*acc)
	return nil
}

// GetAccountByID will return an account using the given id
func (e *Engine) GetAccountByID(id string) (*types.Account, error) {
	acc, ok := e.accs[id]
	if !ok {
		return nil, ErrAccountDoesNotExist
	}
	acccpy := *acc
	return &acccpy, nil
}

func (e *Engine) removeAccount(id string) error {
	delete(e.accs, id)
	return nil
}

// @TODO this function uses a single slice for each call. This is fine now, as we're processing
// everything sequentially, and so there's no possible data-races here. Once we start doing things
// like cleaning up expired markets asynchronously, then this func is not safe for concurrent use
func (e *Engine) accountID(marketID, partyID, asset string, ty types.AccountType) string {
	if len(marketID) <= 0 {
		marketID = noMarket
	}

	// market account
	if len(partyID) <= 0 {
		partyID = systemOwner
	}

	copy(e.idbuf, marketID)
	ln := len(marketID)
	copy(e.idbuf[ln:], partyID)
	ln += len(partyID)
	copy(e.idbuf[ln:], asset)
	ln += len(asset)
	e.idbuf[ln] = byte(ty + 48)
	return string(e.idbuf[:ln+1])
}
