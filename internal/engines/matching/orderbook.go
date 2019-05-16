package matching

import (
	"fmt"
	"sync"

	"code.vegaprotocol.io/vega/internal/logging"
	types "code.vegaprotocol.io/vega/proto"
)

const (
	expiringOrdersInitialCap = 4096
)

type OrderBook struct {
	log *logging.Logger
	Config

	cfgMu           *sync.Mutex
	marketID        string
	buy             *OrderBookSide
	sell            *OrderBookSide
	lastTradedPrice uint64
	latestTimestamp int64
	expiringOrders  sortedorders // keep a list of all expiring trades, these will be in timestamp ascending order.
}

// Create an order book with a given name
func NewOrderBook(log *logging.Logger, config Config, marketID string, proRataMode bool) *OrderBook {
	// setup logger
	log = log.Named(namedLogger)
	log.SetLevel(config.Level.Get())

	return &OrderBook{
		log:            log,
		marketID:       marketID,
		cfgMu:          &sync.Mutex{},
		buy:            &OrderBookSide{log: log, proRataMode: proRataMode},
		sell:           &OrderBookSide{log: log, proRataMode: proRataMode},
		Config:         config,
		expiringOrders: newSortedOrders(expiringOrdersInitialCap),
	}
}

func (s *OrderBook) ReloadConf(cfg Config) {
	s.log.Info("reloading configuration")
	if s.log.GetLevel() != cfg.Level.Get() {
		s.log.Info("updating log level",
			logging.String("old", s.log.GetLevel().String()),
			logging.String("new", cfg.Level.String()),
		)
		s.log.SetLevel(cfg.Level.Get())
	}

	s.cfgMu.Lock()
	s.Config = cfg
	s.cfgMu.Unlock()
}

// Cancel an order that is active on an order book. Market and Order ID are validated, however the order must match
// the order on the book with respect to side etc. The caller will typically validate this by using a store, we should
// not trust that the external world can provide these values reliably.
func (b *OrderBook) CancelOrder(order *types.Order) (*types.OrderCancellationConfirmation, error) {
	// Validate Market
	if order.MarketID != b.marketID {
		b.log.Error("Market ID mismatch",
			logging.Order(*order),
			logging.String("order-book", b.marketID))
	}

	// Validate Order ID must be present
	if order.Id == "" || len(order.Id) < 4 {
		b.log.Error("Order ID missing or invalid",
			logging.Order(*order),
			logging.String("order-book", b.marketID))

		return nil, types.ErrInvalidOrderID
	}

	if order.Side == types.Side_Buy {
		if err := b.buy.RemoveOrder(order); err != nil {
			b.log.Error("Failed to remove order (buy side)",
				logging.Order(*order),
				logging.Error(err),
				logging.String("order-book", b.marketID))

			return nil, types.ErrOrderRemovalFailure
		}
	} else {
		if err := b.sell.RemoveOrder(order); err != nil {
			b.log.Error("Failed to remove order (sell side)",
				logging.Order(*order),
				logging.Error(err),
				logging.String("order-book", b.marketID))

			return nil, types.ErrOrderRemovalFailure
		}
	}

	// Important to mark the order as cancelled (and no longer active)
	order.Status = types.Order_Cancelled

	result := &types.OrderCancellationConfirmation{
		Order: order,
	}
	return result, nil
}

func (b *OrderBook) AmendOrder(order *types.Order) error {
	if err := b.validateOrder(order); err != nil {
		b.log.Error("Order validation failure",
			logging.Order(*order),
			logging.Error(err),
			logging.String("order-book", b.marketID))

		return err
	}

	if order.Side == types.Side_Buy {
		if err := b.buy.amendOrder(order); err != nil {
			b.log.Error("Failed to amend (buy side)",
				logging.Order(*order),
				logging.Error(err),
				logging.String("order-book", b.marketID))

			return err
		}
	} else {
		if err := b.sell.amendOrder(order); err != nil {
			b.log.Error("Failed to amend (sell side)",
				logging.Order(*order),
				logging.Error(err),
				logging.String("order-book", b.marketID))

			return err
		}
	}

	return nil
}

// Add an order and attempt to uncross the book, returns a TradeSet protobuf message object
func (b *OrderBook) SubmitOrder(order *types.Order) (*types.OrderConfirmation, error) {
	if err := b.validateOrder(order); err != nil {
		return nil, err
	}

	if order.CreatedAt > b.latestTimestamp {
		b.latestTimestamp = order.CreatedAt
	}

	if b.LogPriceLevelsDebug {
		b.PrintState("Entry state:")
	}

	// uncross with opposite
	trades, impactedOrders, lastTradedPrice := b.getOppositeSide(order.Side).uncross(order)
	if lastTradedPrice != 0 {
		b.lastTradedPrice = lastTradedPrice
	}

	// if state of the book changed show state
	if b.LogPriceLevelsDebug && len(trades) != 0 {
		b.PrintState("After uncross state:")
	}

	// if order is persistent type add to order book to the correct side
	if (order.Type == types.Order_GTC || order.Type == types.Order_GTT) && order.Remaining > 0 {

		// GTT orders need to be added to the expiring orders table, these orders will be removed when expired.
		if order.Type == types.Order_GTT {
			b.expiringOrders = b.expiringOrders.insert(*order)
		}

		b.getSide(order.Side).addOrder(order, order.Side)

		if b.LogPriceLevelsDebug {
			b.PrintState("After addOrder state:")
		}
	}

	// did we fully fill the originating order?
	if order.Remaining == 0 {
		order.Status = types.Order_Filled
	}

	// update order statuses based on the order types if they didn't trade
	if (order.Type == types.Order_FOK || order.Type == types.Order_ENE) && order.Remaining == order.Size {
		order.Status = types.Order_Stopped
	}

	for idx := range impactedOrders {
		if impactedOrders[idx].Remaining == 0 {
			impactedOrders[idx].Status = types.Order_Filled

			// Ensure any fully filled impacted GTT orders are removed
			// from internal matching engine pending orders list
			if impactedOrders[idx].Type == types.Order_GTT {
				b.removePendingGttOrder(*impactedOrders[idx])
			}
		}
	}

	orderConfirmation := makeResponse(order, trades, impactedOrders)
	return orderConfirmation, nil
}

func (b *OrderBook) DeleteOrder(order *types.Order) error {
	err := b.getSide(order.Side).RemoveOrder(order)
	return err
}

// RemoveExpiredOrders removes any GTT orders that will expire on or before the expiration timestamp (epoch+nano).
// expirationTimestamp must be of the format unix epoch seconds with nanoseconds e.g. 1544010789803472469.
// RemoveExpiredOrders returns a slice of Orders that were removed, internally it will remove the orders from the
// matching engine price levels. The returned orders will have an Order_Expired status, ready to update in stores.
func (b *OrderBook) RemoveExpiredOrders(expirationTs int64) []types.Order {
	var expired []types.Order
	expired, b.expiringOrders = b.expiringOrders.removeExpired(expirationTs)
	for _, v := range expired {
		v := v
		b.DeleteOrder(&v)
		v.Status = types.Order_Expired
	}

	if b.LogRemovedOrdersDebug {
		b.log.Debug("Removed expired orders from order book",
			logging.String("order-book", b.marketID),
			logging.Int("expired-orders", len(expired)),
			logging.Int("remaining-orders", len(b.expiringOrders)))
	}

	return expired
}

func (b OrderBook) getSide(orderSide types.Side) *OrderBookSide {
	if orderSide == types.Side_Buy {
		return b.buy
	} else {
		return b.sell
	}
}

func (b *OrderBook) getOppositeSide(orderSide types.Side) *OrderBookSide {
	if orderSide == types.Side_Buy {
		return b.sell
	} else {
		return b.buy
	}
}

func (b OrderBook) removePendingGttOrder(order types.Order) bool {
	found := -1
	for idx, or := range b.expiringOrders {
		if or.Id == order.Id {
			found = idx
		}
	}
	if found > -1 {
		b.expiringOrders = append(b.expiringOrders[:found], b.expiringOrders[found+1:]...)
		return true
	}
	return false
}

func makeResponse(order *types.Order, trades []*types.Trade, impactedOrders []*types.Order) *types.OrderConfirmation {
	return &types.OrderConfirmation{
		Order:                 order,
		PassiveOrdersAffected: impactedOrders,
		Trades:                trades,
	}
}

func (b *OrderBook) PrintState(types string) {
	b.log.Debug(fmt.Sprintf("%s", types))
	b.log.Debug("------------------------------------------------------------")
	b.log.Debug("                        BUY SIDE                            ")
	for _, priceLevel := range b.buy.getLevels() {
		if len(priceLevel.orders) > 0 {
			priceLevel.print(b.log)
		}
	}
	b.log.Debug("------------------------------------------------------------")
	b.log.Debug("                        SELL SIDE                           ")
	for _, priceLevel := range b.sell.getLevels() {
		if len(priceLevel.orders) > 0 {
			priceLevel.print(b.log)
		}
	}
	b.log.Debug("------------------------------------------------------------")
}
