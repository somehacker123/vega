package settlement

import (
	"sync"
	"time"

	"code.vegaprotocol.io/vega/internal/logging"
	"code.vegaprotocol.io/vega/internal/products"
	types "code.vegaprotocol.io/vega/proto"
)

//go:generate go run github.com/golang/mock/mockgen -destination mocks/market_position_mock.go -package mocks code.vegaprotocol.io/vega/internal/engines/settlement MarketPosition
type MarketPosition interface {
	Party() string
	Size() int64
	Price() uint64
}

type pos struct {
	size  int64
	price uint64
}

type Engine struct {
	log *logging.Logger

	Config
	mu      *sync.Mutex
	product products.Product
	pos     map[string]*pos
}

func New(log *logging.Logger, conf Config) *Engine {
	// setup logger
	log = log.Named(namedLogger)
	log.SetLevel(conf.Level.Get())

	return &Engine{
		log:    log,
		Config: conf,
		mu:     &sync.Mutex{},
		pos:    map[string]*pos{},
	}
}

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
	e.mu.Lock()
	for _, p := range positions {
		e.updatePosition(p, p.Price())
	}
	e.mu.Unlock()
}

func (e *Engine) updatePosition(p MarketPosition, price uint64) {
	party := p.Party()
	ps, ok := e.pos[party]
	if !ok {
		ps = &pos{}
		e.pos[party] = ps
	}
	// price can come from either position, or trade (Update vs SettleMTM)
	ps.price = price
	ps.size = p.Size()
}

func (e *Engine) Settle(t time.Time) ([]*types.SettlePosition, error) {
	e.mu.Lock()
	e.log.Debugf("Settling market, closed at %s", t.Format(time.RFC3339))
	positions, err := e.settleAll()
	if err != nil {
		e.log.Error(
			"Something went wrong trying to settle positions",
			logging.Error(err),
		)
		return nil, err
	}
	e.mu.Unlock()
	return positions, nil
}

func (e *Engine) SettleMTM(trade *types.Trade, ch <-chan MarketPosition) <-chan []*types.SettlePosition {
	// put the positions on here once we've worked out what all we need to settle
	sch := make(chan []*types.SettlePosition)
	go func() {
		posSlice := make([]*types.SettlePosition, 0, cap(ch))
		winSlice := make([]*types.SettlePosition, 0, cap(ch)/2)
		e.mu.Lock()
		for pos := range ch {
			if pos == nil {
				break
			}
			// update position for trader - always keep track of latest position
			pp := pos.Price()
			ps := pos.Size()
			// all positions need to be updated to the new market price
			e.updatePosition(pos, trade.Price)
			// we should set the new position to market price here... somehow
			// e.Update([]MarketPosition{pos})
			if pp == trade.Price || ps == 0 {
				// nothing has changed or there's no position to settle
				continue
			}
			// e.g. position avg -> 90, market price 100:
			// short -> (100 - 90) * -10 => -100 ==> MTM_LOSS
			// long -> (100-90) * 10 => 100 ==> MTM_WIN
			// short -> (100 - 110) * -10 => 100 ==> MTM_WIN
			// long -> (100 - 110) * 10 => -100 ==> MTM_LOSS
			mtmShare := (int64(trade.Price) - int64(pp)) * ps
			settle := &types.SettlePosition{
				Owner: pos.Party(),
				Size:  1, // this is an absolute delta based on volume, so size is always 1
				Amount: &types.FinancialAmount{
					Amount: mtmShare, // current delta -> mark price minus current position average
				},
			}
			if mtmShare > 0 {
				settle.Type = types.SettleType_MTM_WIN
				winSlice = append(winSlice, settle)
			} else {
				settle.Type = types.SettleType_MTM_LOSS
				posSlice = append(posSlice, settle)
			}
		}
		e.mu.Unlock()
		posSlice = append(posSlice, winSlice...)
		sch <- posSlice
		close(sch)
	}()
	return sch
}

// simplified settle call
func (e *Engine) settleAll() ([]*types.SettlePosition, error) {
	// there should be as many positions as there are traders (obviously)
	aggregated := make([]*types.SettlePosition, 0, len(e.pos))
	// traders who are in the black should be appended (collect first) obvioulsy.
	// The split won't always be 50-50, but it's a reasonable approximation
	owed := make([]*types.SettlePosition, 0, len(e.pos)/2)
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
		// and the final price coming from the oracle, so oracle_price - market_price * volume (check with Tamlyn whether this should be absolute or not)
		if err != nil {
			e.log.Error(
				"Failed to settle position for trader",
				logging.String("trader-id", party),
				logging.Error(err),
			)
			return nil, err
		}
		settlePos := &types.SettlePosition{
			Owner:  party,
			Size:   uint64(pos.size),
			Amount: amt,
			Type:   types.SettleType_LOSS, // this is a poor name, will be changed later
		}
		if pos.size < 0 {
			// ensure absolute value
			settlePos.Size = uint64(-pos.size)
		}
		e.log.Debug(
			"Settled position for trader",
			logging.String("trader-id", party),
			logging.Int64("amount", amt.Amount),
		)
		if amt.Amount < 0 {
			aggregated = append(aggregated, settlePos)
		} else {
			// bad name again, but SELL means trader is owed money
			settlePos.Type = types.SettleType_WIN
			owed = append(owed, settlePos)
		}
	}
	// append the traders in the black to the end
	aggregated = append(aggregated, owed...)
	return aggregated, nil
}
