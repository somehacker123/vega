package settlement

import (
	"sync"
	"time"

	"code.vegaprotocol.io/vega/internal/events"
	"code.vegaprotocol.io/vega/internal/logging"
	"code.vegaprotocol.io/vega/internal/metrics"

	types "code.vegaprotocol.io/vega/proto"
)

// We should really use a type from the proto package for this, although, these mocks are kind of easy to set up :)

// MarketPosition ...
//go:generate go run github.com/golang/mock/mockgen -destination mocks/market_position_mock.go -package mocks code.vegaprotocol.io/vega/internal/settlement MarketPosition
type MarketPosition interface {
	Party() string
	Size() int64
	Buy() int64
	Sell() int64
	Price() uint64
	ClearPotentials()
}

// Product ...
//go:generate go run github.com/golang/mock/mockgen -destination mocks/settlement_product_mock.go -package mocks code.vegaprotocol.io/vega/internal/settlement Product
type Product interface {
	Settle(entryPrice uint64, netPosition int64) (*types.FinancialAmount, error)
	GetAsset() string
}

// Engine - the main type (of course)
type Engine struct {
	log *logging.Logger

	Config
	product  Product
	pos      map[string]*pos
	posMu    sync.Mutex
	closed   map[string][]*pos
	closedMu sync.Mutex
	market   string
}

// New instanciate a new instance of the settlement engine
func New(log *logging.Logger, conf Config, product Product, market string) *Engine {
	// setup logger
	log = log.Named(namedLogger)
	log.SetLevel(conf.Level.Get())

	return &Engine{
		log:     log,
		Config:  conf,
		product: product,
		pos:     map[string]*pos{},
		market:  market,
		// no need to initialised `closed` map
	}
}

// ReloadConf update the internal configuration of the settlement engined
func (e *Engine) ReloadConf(cfg Config) {
	e.log.Info("reloading configuration")
	if e.log.GetLevel() != cfg.Level.Get() {
		e.log.Info("updating log level",
			logging.String("old", e.log.GetLevel().String()),
			logging.String("new", cfg.Level.String()),
		)
		e.log.SetLevel(cfg.Level.Get())
	}

	e.Config = cfg
}

// Update - takes market positions, keeps track of things
func (e *Engine) Update(positions []MarketPosition) {
	e.posMu.Lock()
	for _, p := range positions {
		party := p.Party()
		ps, ok := e.pos[party]
		if !ok {
			ps = newPos(party)
			e.pos[party] = ps
		}
		// price can come from either position, or trade (Update vs SettleMTM)
		ps.price = p.Price()
		ps.size = p.Size()
	}
	e.posMu.Unlock()
}

func (e *Engine) getCurrentPosition(party string) *pos {
	e.posMu.Lock()
	p, ok := e.pos[party]
	if !ok {
		p = newPos(party)
		e.pos[party] = p
	}
	e.posMu.Unlock()
	return p
}

// RemoveDistressed - remove whatever settlement data we have for distressed traders
// they are being closed out, and shouldn't be part of any MTM settlement or closing settlement
func (e *Engine) RemoveDistressed(traders []events.MarketPosition) {
	e.posMu.Lock()
	e.closedMu.Lock()
	for _, trader := range traders {
		key := trader.Party()
		delete(e.pos, key)
		delete(e.closed, key)
	}
	e.closedMu.Unlock()
	e.posMu.Unlock()
}

// Settle run settlement over all the positions
func (e *Engine) Settle(t time.Time) ([]*types.Transfer, error) {
	e.log.Debugf("Settling market, closed at %s", t.Format(time.RFC3339))
	positions, err := e.settleAll()
	if err != nil {
		e.log.Error(
			"Something went wrong trying to settle positions",
			logging.Error(err),
		)
		return nil, err
	}
	return positions, nil
}

// ListenClosed listen to all positions for distressed trader
// needing to be closed
func (e *Engine) ListenClosed(ch <-chan events.MarketPosition) {
	// lock before we can start
	e.closedMu.Lock()
	go func() {
		// wipe closed map
		e.closed = map[string][]*pos{}
		for ps := range ch {
			trader := ps.Party()
			size := ps.Size()
			price := ps.Price()
			updatePrice := price
			// check current position to see if trade closed out some position
			current := e.getCurrentPosition(trader)
			// if trader is long, and trade closed out (part of) long position, or trader was short, and is now "less short"
			if (current.size > 0 && size < current.size) || (current.size < 0 && size > current.size) {
				closed := current.size
				// trader was long, and still is || trader was short && still is
				if (current.size > 0 && size > 0) || (current.size < 0 && size < 0) {
					// trader was +10, now +5 -> +10 - +5 == MTM on +5 closed positions --> good
					// trader was -10, now -5 -> -10 - -5 == MTM on -5 closed positions --> good
					closed -= size
					updatePrice = current.price
				}
				// let's add this change to the traders' closed positions to be added to the MTM settlement later on
				trades, ok := e.closed[trader]
				if !ok {
					trades = []*pos{}
				}
				pos := newPos(trader)
				pos.size = closed
				pos.price = current.price // we closed out at the old price vs mark price
				e.closed[trader] = append(trades, pos)
			}
			// we've taken the closed out stuff into account, so we can freely update the size here
			current.size = size
			// the position price is possibly updated (e.g. if there was no open position prior to this, or trader went from long to short or vice-versa)
			current.price = updatePrice
		}
		e.closedMu.Unlock()
	}()
}

// SettleOrder - settlements based on order-level, can take several update positions, and marks all to market
// if party size and price were both updated (ie party was a trader), we're combining both MTM's (old + new position)
// and creating a single transfer from that
func (e *Engine) SettleOrder(markPrice uint64, positions []events.MarketPosition) []events.Transfer {
	timer := metrics.NewTimeCounter("-", "settlement", "SettleOrder")

	transfers := make([]events.Transfer, 0, len(positions))
	winTransfers := make([]events.Transfer, 0, len(positions)/2)
	// see if we've got closed out positions
	e.closedMu.Lock()
	closed := e.closed
	// reset map here in case we're going to call this with just an updated mark price
	e.closed = map[string][]*pos{}
	for _, pos := range positions {
		size := pos.Size()
		// price := pos.Price()
		trader := pos.Party()
		current := e.getCurrentPosition(trader)
		// markPrice was already set by positions engine
		// e.g. position avg -> 90, mark price 100:
		// short -> (100 - 90) * -10 => -100 ==> MTM_LOSS
		// long -> (100-90) * 10 => 100 ==> MTM_WIN
		// short -> (100 - 110) * -10 => 100 ==> MTM_WIN
		// long -> (100 - 110) * 10 => -100 ==> MTM_LOSS
		closedOut, _ := closed[trader]
		// updated price is mark price, mark against that using current known price
		// if price == markPrice {
		// 	price = current.price
		// }
		// first we MTM the old position, exclude closed positions
		mtmShare := calcMTM(int64(markPrice), current.size, int64(current.price), closedOut)
		// now MTM the new position if needed
		// mtmShare += calcMTM(int64(markPrice), size, int64(price), closedOut)
		// update position
		current.size = size
		current.price = markPrice
		// there's nothing to mark to market
		if mtmShare == 0 {
			continue
		}
		settle := &types.Transfer{
			Owner: current.party,
			Size:  1, // this is an absolute delta based on volume, so size is always 1
			Amount: &types.FinancialAmount{
				Amount: mtmShare, // current delta -> mark price minus current position average
				Asset:  e.product.GetAsset(),
			},
		}
		if mtmShare > 0 {
			settle.Type = types.TransferType_MTM_WIN
			winTransfers = append(winTransfers, &mtmTransfer{
				MarketPosition: pos,
				transfer:       settle,
			})
		} else {
			// losses are prepended
			settle.Type = types.TransferType_MTM_LOSS
			transfers = append(transfers, &mtmTransfer{
				MarketPosition: pos,
				transfer:       settle,
			})
		}
	}
	e.closedMu.Unlock()
	transfers = append(transfers, winTransfers...)
	timer.EngineTimeCounterAdd()
	return transfers
}

// calcMTM only handles futures ATM. The formula is simple:
// amount =  prev_vol * (current_price - prev_mark_price) + SUM(new_trade := range trades)( new_trade(i).volume(party)*(current_price - new_trade(i).price )
// given that the new trades price will equal new mark price,  the sum(trades) bit will probably == 0 for nicenet
func calcMTM(markPrice, size, price int64, closed []*pos) (mtmShare int64) {
	mtmShare = (markPrice - price) * size
	for _, c := range closed {
		// add MTM compared to trade price for the positions that were closed out
		mtmShare += (markPrice - int64(c.price)) * c.size
	}
	return
}

// simplified settle call
func (e *Engine) settleAll() ([]*types.Transfer, error) {
	e.posMu.Lock()
	defer e.posMu.Unlock()
	// there should be as many positions as there are traders (obviously)
	aggregated := make([]*types.Transfer, 0, len(e.pos))
	// traders who are in the black should be appended (collect first).
	// The split won't always be 50-50, but it's a reasonable approximation
	owed := make([]*types.Transfer, 0, len(e.pos)/2)
	for party, pos := range e.pos {
		// this is possible now, with the Mark to Market stuff, it's possible we've settled any and all positions for a given trader
		if pos.size == 0 {
			continue
		}
		e.log.Debug("Settling position for trader", logging.String("trader-id", party))
		// @TODO - there was something here... the final amount had to be oracle - market or something
		// check with Tamlyn why that was, because we're only handling open positions here...
		amt, err := e.product.Settle(pos.price, pos.size)
		// for now, product.Settle returns the total value, we need to only settle the delta between a traders current position
		// and the final price coming from the oracle, so oracle_price - mark_price * volume (check with Tamlyn whether this should be absolute or not)
		if err != nil {
			e.log.Error(
				"Failed to settle position for trader",
				logging.String("trader-id", party),
				logging.Error(err),
			)
			return nil, err
		}
		settlePos := &types.Transfer{
			Owner:  party,
			Size:   1,
			Amount: amt,
		}
		e.log.Debug(
			"Settled position for trader",
			logging.String("trader-id", party),
			logging.Int64("amount", amt.Amount),
		)
		if amt.Amount < 0 {
			// trader is winning...
			settlePos.Type = types.TransferType_LOSS
			aggregated = append(aggregated, settlePos)
		} else {
			// bad name again, but SELL means trader is owed money
			settlePos.Type = types.TransferType_WIN
			owed = append(owed, settlePos)
		}
	}
	// append the traders in the black to the end
	aggregated = append(aggregated, owed...)
	return aggregated, nil
}
