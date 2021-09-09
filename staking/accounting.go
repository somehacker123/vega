package staking

import (
	"context"
	"errors"
	"fmt"
	"time"

	vgproto "code.vegaprotocol.io/protos/vega"
	"code.vegaprotocol.io/vega/assets/erc20"
	"code.vegaprotocol.io/vega/events"
	"code.vegaprotocol.io/vega/logging"
	"code.vegaprotocol.io/vega/types"
	"code.vegaprotocol.io/vega/types/num"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcmn "github.com/ethereum/go-ethereum/common"
)

// Broker - the event bus
type Broker interface {
	Send(events.Event)
}

type EthereumClientCaller interface {
	bind.ContractCaller
}

var (
	ErrNoBalanceForParty = errors.New("no balance for party")
)

type Accounting struct {
	log       *logging.Logger
	ethClient EthereumClientCaller
	cfg       Config
	broker    Broker
	accounts  map[string]*StakingAccount

	stakingAssetTotalSupply *num.Uint
	ethCfg                  vgproto.EthereumConfig
}

func NewAccounting(
	log *logging.Logger,
	cfg Config,
	broker Broker,
	ethClient EthereumClientCaller,
) *Accounting {
	log = log.Named(namedLogger)
	log.SetLevel(cfg.Level.Get())

	return &Accounting{
		log:                     log,
		cfg:                     cfg,
		broker:                  broker,
		ethClient:               ethClient,
		accounts:                map[string]*StakingAccount{},
		stakingAssetTotalSupply: num.Zero(),
	}
}

func (a *Accounting) HasBalance(party string) bool {
	_, ok := a.accounts[party]
	return ok
}

func (a *Accounting) AddEvent(ctx context.Context, evt *types.StakeLinking) {
	acc, ok := a.accounts[evt.Party]
	if !ok {
		a.broker.Send(events.NewPartyEvent(ctx, types.Party{Id: evt.Party}))
		acc = NewStakingAccount(evt.Party)
		a.accounts[evt.Party] = acc
	}

	// errors here do not really matter I'd say
	// they are either validation issue, that we can just log
	// but should never happen as things should be created properly
	// or errors from event being received in the wrong order
	// but that we cannot really prevent and that the account
	// would recover from by itself later on.
	// Negative balance is possible when processing orders in disorder,
	// not a big deal
	if err := acc.AddEvent(evt); err != nil && err != ErrNegativeBalance {
		a.log.Error("could not add event to staking account",
			logging.Error(err))
		return
	}
}

//GetAllAvailableBalances returns the staking balance for all parties
func (a *Accounting) GetAllAvailableBalances() map[string]*num.Uint {
	balances := map[string]*num.Uint{}
	for party, acc := range a.accounts {
		balances[party] = acc.GetAvailableBalance()
	}
	return balances
}

func (a *Accounting) OnEthereumConfigUpdate(_ context.Context, rawcfg interface{}) error {
	cfg, ok := rawcfg.(*vgproto.EthereumConfig)
	if !ok {
		return ErrNotAnEthereumConfig
	}

	a.ethCfg = *cfg

	if err := a.updateStakingAssetTotalSupply(); err != nil {
		return fmt.Errorf("failed to update staking asset total supply: %w", err)
	}

	return nil
}

func (a *Accounting) updateStakingAssetTotalSupply() error {
	if len(a.ethCfg.StakingBridgeAddresses) <= 0 {
		a.log.Error("no staking bridge address setup",
			logging.String("eth-cfg", a.ethCfg.String()),
		)
		return nil
	}

	addr := ethcmn.HexToAddress(a.ethCfg.StakingBridgeAddresses[0])

	sc, err := NewStakingCaller(addr, a.ethClient)
	if err != nil {
		return err
	}

	st, err := sc.StakingToken(&bind.CallOpts{})
	if err != nil {
		return err
	}

	tc, err := erc20.NewErc20Caller(st, a.ethClient)
	if err != nil {
		return err
	}

	ts, err := tc.TotalSupply(&bind.CallOpts{})
	if err != nil {
		return err
	}

	totalSupply, overflowed := num.UintFromBig(ts)
	if overflowed {
		return fmt.Errorf("failed to convert big.Int to num.Uint: %s", ts.String())
	}

	a.stakingAssetTotalSupply = totalSupply

	symbol, _ := tc.Symbol(&bind.CallOpts{})
	decimals, _ := tc.Decimals(&bind.CallOpts{})

	a.log.Debug("staking asset loaded",
		logging.String("symbol", symbol),
		logging.Uint8("decimals", decimals),
		logging.String("total-supply", ts.String()),
	)

	return nil
}

func (a *Accounting) GetAvailableBalance(party string) (*num.Uint, error) {
	acc, ok := a.accounts[party]
	if !ok {
		return num.Zero(), ErrNoBalanceForParty
	}

	return acc.GetAvailableBalance(), nil
}

func (a *Accounting) GetAvailableBalanceAt(
	party string, at time.Time) (*num.Uint, error) {
	acc, ok := a.accounts[party]
	if !ok {
		return num.Zero(), ErrNoBalanceForParty
	}

	return acc.GetAvailableBalanceAt(at)
}

func (a *Accounting) GetAvailableBalanceInRange(
	party string, from, to time.Time) (*num.Uint, error) {
	acc, ok := a.accounts[party]
	if !ok {
		return num.Zero(), ErrNoBalanceForParty
	}

	return acc.GetAvailableBalanceInRange(from, to)
}

func (a *Accounting) GetStakingAssetTotalSupply() *num.Uint {
	return a.stakingAssetTotalSupply.Clone()
}
