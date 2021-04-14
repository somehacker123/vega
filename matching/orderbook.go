package matching

import (
	"sort"
	"sync"
	"time"

	"code.vegaprotocol.io/vega/crypto"
	"code.vegaprotocol.io/vega/events"
	"code.vegaprotocol.io/vega/logging"
	"code.vegaprotocol.io/vega/metrics"
	types "code.vegaprotocol.io/vega/proto"

	"github.com/pkg/errors"
)

var (
	// ErrNotEnoughOrders signals that not enough orders were
	// in the book to achieve a given operation
	ErrNotEnoughOrders   = errors.New("insufficient orders")
	ErrOrderDoesNotExist = errors.New("order does not exist")
	ErrInvalidVolume     = errors.New("invalid volume")
)

// OrderBook represents the book holding all orders in the system.
type OrderBook struct {
	log *logging.Logger
	Config

	cfgMu           *sync.Mutex
	marketID        string
	buy             *OrderBookSide
	sell            *OrderBookSide
	lastTradedPrice uint64
	latestTimestamp int64
	ordersByID      map[string]*types.Order
	ordersPerParty  map[string]map[string]struct{}
	auction         bool
	batchID         uint64
}

// CumulativeVolumeLevel represents the cumulative volume at a price level for both bid and ask
type CumulativeVolumeLevel struct {
	price               uint64
	bidVolume           uint64
	askVolume           uint64
	cumulativeBidVolume uint64
	cumulativeAskVolume uint64
	maxTradableAmount   uint64
}

func (b *OrderBook) Hash() []byte {
	return crypto.Hash(append(b.buy.Hash(), b.sell.Hash()...))
}

// NewOrderBook create an order book with a given name.
func NewOrderBook(log *logging.Logger, config Config, marketID string, auction bool) *OrderBook {
	// setup logger
	log = log.Named(namedLogger)
	log.SetLevel(config.Level.Get())

	return &OrderBook{
		log:            log,
		marketID:       marketID,
		cfgMu:          &sync.Mutex{},
		buy:            &OrderBookSide{log: log, side: types.Side_SIDE_BUY},
		sell:           &OrderBookSide{log: log, side: types.Side_SIDE_SELL},
		Config:         config,
		ordersByID:     map[string]*types.Order{},
		auction:        auction,
		batchID:        0,
		ordersPerParty: map[string]map[string]struct{}{},
	}
}

// ReloadConf is used in order to reload the internal configuration of
// the OrderBook
func (b *OrderBook) ReloadConf(cfg Config) {
	b.log.Info("reloading configuration")
	if b.log.GetLevel() != cfg.Level.Get() {
		b.log.Info("updating log level",
			logging.String("old", b.log.GetLevel().String()),
			logging.String("new", cfg.Level.String()),
		)
		b.log.SetLevel(cfg.Level.Get())
	}

	b.cfgMu.Lock()
	b.Config = cfg
	b.cfgMu.Unlock()
}

// GetCloseoutPrice returns the exit price which would be achieved for a given
// volume and give side of the book
func (b *OrderBook) GetCloseoutPrice(volume uint64, side types.Side) (uint64, error) {
	var price uint64
	if b.auction {
		p := b.GetIndicativePrice()
		return p, nil
	}

	if volume == 0 {
		return 0, ErrInvalidVolume
	}
	vol := volume
	if side == types.Side_SIDE_SELL {
		levels := b.sell.getLevels()
		for i := len(levels) - 1; i >= 0; i-- {
			lvl := levels[i]
			if lvl.volume >= vol {
				price += lvl.price * vol
				return price / volume, nil
			}
			price += lvl.price * lvl.volume
			vol -= lvl.volume
		}
		// at this point, we should check vol, make sure it's 0, if not return an error to indicate something is wrong
		// still return the price for the volume we could close out, so the caller can make a decision on what to do
		if vol == volume {
			return b.lastTradedPrice, ErrNotEnoughOrders
		}
		price = price / (volume - vol)
		if vol != 0 {
			return price, ErrNotEnoughOrders
		}
		return price, nil
	}
	// side == buy
	levels := b.buy.getLevels()
	for i := len(levels) - 1; i >= 0; i-- {
		lvl := levels[i]
		if lvl.volume >= vol {
			price += lvl.price * vol
			return price / volume, nil
		}
		price += lvl.price * lvl.volume
		vol -= lvl.volume
	}
	// if we reach this point, chances are vol != 0, in which case we should return an error along with the price
	if vol == volume {
		return b.lastTradedPrice, ErrNotEnoughOrders
	}
	price = price / (volume - vol)
	if vol != 0 {
		return price, ErrNotEnoughOrders
	}
	return price, nil
}

// EnterAuction Moves the order book into an auction state
func (b *OrderBook) EnterAuction() ([]*types.Order, error) {
	// Scan existing orders to see which ones can be kept or cancelled
	buyCancelledOrders, err := b.buy.getOrdersToCancel(true)
	if err != nil {
		return nil, err
	}

	sellCancelledOrders, err := b.sell.getOrdersToCancel(true)
	if err != nil {
		return nil, err
	}

	// Set the market state
	b.auction = true

	// Return all the orders that have been removed from the book and need to be cancelled
	ordersToCancel := buyCancelledOrders
	ordersToCancel = append(ordersToCancel, sellCancelledOrders...)
	return ordersToCancel, nil
}

// LeaveAuction Moves the order book back into continuous trading state
func (b *OrderBook) LeaveAuction(at time.Time) ([]*types.OrderConfirmation, []*types.Order, error) {
	// Update batchID
	b.batchID++

	ts := at.UnixNano()

	// Uncross the book
	uncrossedOrders, err := b.uncrossBook()
	if err != nil {
		return nil, nil, err
	}

	for _, uo := range uncrossedOrders {
		if uo.Order.Remaining == 0 {
			uo.Order.Status = types.Order_STATUS_FILLED
			// delete from lookup table
			delete(b.ordersByID, uo.Order.Id)
			delete(b.ordersPerParty[uo.Order.PartyId], uo.Order.Id)
		}

		uo.Order.UpdatedAt = ts
		for idx, po := range uo.PassiveOrdersAffected {
			po.UpdatedAt = ts
			// also remove the orders from lookup tables
			if uo.PassiveOrdersAffected[idx].Remaining == 0 {
				uo.PassiveOrdersAffected[idx].Status = types.Order_STATUS_FILLED

				// delete from lookup table
				delete(b.ordersByID, po.Id)
				delete(b.ordersPerParty[po.PartyId], po.Id)
			}
		}
		for _, tr := range uo.Trades {
			tr.Timestamp = ts
		}
	}

	// Remove any orders that will not be valid in continuous trading
	buyOrdersToCancel, err := b.buy.getOrdersToCancel(false)
	if err != nil {
		return nil, nil, err
	}

	sellOrdersToCancel, err := b.sell.getOrdersToCancel(false)
	if err != nil {
		return nil, nil, err
	}
	// Return all the orders that have been cancelled from the book
	ordersToCancel := append(buyOrdersToCancel, sellOrdersToCancel...)

	for _, oc := range ordersToCancel {
		oc.UpdatedAt = ts
	}

	// Flip back to continuous
	b.auction = false

	return uncrossedOrders, ordersToCancel, nil
}

func (b OrderBook) InAuction() bool {
	return b.auction
}

// CanUncross - a clunky name for a somewhat clunky function: this checks if there will be LIMIT orders
// on the book after we uncross the book (at the end of an auction). If this returns false, the opening auction should be extended
func (b *OrderBook) CanUncross() bool {
	return b.canUncross(true)
}

func (b *OrderBook) BidAndAskPresentAfterAuction() bool {
	return b.canUncross(false)
}

func (b *OrderBook) canUncross(requireTrades bool) bool {
	bb, err := b.GetBestBidPrice() // sell
	if err != nil {
		return false
	}
	ba, err := b.GetBestAskPrice() // buy
	if err != nil || bb == 0 || ba == 0 || (requireTrades && bb < ba) {
		return false
	}

	// check all buy price levels below ba, find limit orders
	buyMatch := false
	// iterate from the end, where best is
	for i := len(b.buy.levels) - 1; i >= 0; i-- {
		l := b.buy.levels[i]
		if l.price < ba {
			for _, o := range l.orders {
				// limit order && not just GFA found
				if o.Type == types.Order_TYPE_LIMIT && o.TimeInForce != types.Order_TIME_IN_FORCE_GFA {
					buyMatch = true
					break
				}
			}
		}
	}
	sellMatch := false
	for i := len(b.sell.levels) - 1; i >= 0; i-- {
		l := b.sell.levels[i]
		if l.price > bb {
			for _, o := range l.orders {
				if o.Type == types.Order_TYPE_LIMIT && o.TimeInForce != types.Order_TIME_IN_FORCE_GFA {
					sellMatch = true
					break
				}
			}
		}
	}
	// non-GFA orders outside the price range found on the book, we can uncross
	if buyMatch && sellMatch {
		return true
	}
	_, v, _ := b.GetIndicativePriceAndVolume()
	// no buy orders remaining on the book after uncrossing, it buyMatches exactly
	vol := uint64(0)
	if !buyMatch {
		for i := len(b.buy.levels) - 1; i >= 0; i-- {
			l := b.buy.levels[i]
			// buy orders are ordered ascending
			if l.price < ba {
				break
			}
			for _, o := range l.orders {
				vol += o.Remaining
				// we've filled the uncrossing volume, and found an order that is not GFA
				if vol > v && o.TimeInForce != types.Order_TIME_IN_FORCE_GFA {
					buyMatch = true
					break
				}
			}
		}
		if !buyMatch {
			return false
		}
	}
	// we've had to check buy side - sell side is fine
	if sellMatch {
		return true
	}

	vol = 0
	// for _, l := range b.sell.levels {
	// sell side is ordered descending
	for i := len(b.sell.levels) - 1; i >= 0; i-- {
		l := b.sell.levels[i]
		if l.price > bb {
			break
		}
		for _, o := range l.orders {
			vol += o.Remaining
			if vol > v && o.TimeInForce != types.Order_TIME_IN_FORCE_GFA {
				sellMatch = true
				break
			}
		}
	}

	return sellMatch
}

// GetIndicativePriceAndVolume Calculates the indicative price and volume of the order book without modifying the order book state
func (b *OrderBook) GetIndicativePriceAndVolume() (retprice uint64, retvol uint64, retside types.Side) {

	bestBid, err := b.GetBestBidPrice()
	if err != nil {
		return 0, 0, types.Side_SIDE_UNSPECIFIED
	}
	bestAsk, err := b.GetBestAskPrice()
	if err != nil {
		return 0, 0, types.Side_SIDE_UNSPECIFIED
	}

	// Short circuit if the book is not crossed
	if bestBid < bestAsk || bestBid == 0 || bestAsk == 0 {
		return 0, 0, types.Side_SIDE_UNSPECIFIED
	}

	// Generate a set of price level pairs with their maximum tradable volumes
	cumulativeVolumes, maxTradableAmount := b.buildCumulativePriceLevels(bestBid, bestAsk)

	// Pull out all prices that match that volume
	prices := make([]uint64, 0, len(cumulativeVolumes))
	for _, value := range cumulativeVolumes {
		if value.maxTradableAmount == maxTradableAmount {
			prices = append(prices, value.price)
		}
	}

	// get the maximum volume price from the average of the maximum and minimum tradable price levels
	var uncrossPrice uint64
	var uncrossSide types.Side
	if len(prices) > 0 {
		uncrossPrice = (prices[len(prices)-1] + prices[0]) / 2
	}

	// See which side we should fully process when we uncross
	ordersToFill := int64(maxTradableAmount)
	for _, value := range cumulativeVolumes {
		ordersToFill -= int64(value.bidVolume)
		if ordersToFill == 0 {
			// Buys fill exactly, uncross from the buy side
			uncrossSide = types.Side_SIDE_BUY
			break
		} else if ordersToFill < 0 {
			// Buys are not exact, uncross from the sell side
			uncrossSide = types.Side_SIDE_SELL
			break
		}
	}
	return uncrossPrice, maxTradableAmount, uncrossSide
}

// GetIndicativePrice Calculates the indicative price of the order book without modifying the order book state
func (b *OrderBook) GetIndicativePrice() (retprice uint64) {
	bestBid, err := b.GetBestBidPrice()
	if err != nil {
		return 0
	}
	bestAsk, err := b.GetBestAskPrice()
	if err != nil {
		return 0
	}

	// Short circuit if the book is not crossed
	if bestBid < bestAsk || bestBid == 0 || bestAsk == 0 {
		return 0
	}

	// Generate a set of price level pairs with their maximum tradable volumes
	cumulativeVolumes, maxTradableAmount := b.buildCumulativePriceLevels(bestBid, bestAsk)

	// Pull out all prices that match that volume
	prices := make([]uint64, 0, len(cumulativeVolumes))
	for _, value := range cumulativeVolumes {
		if value.maxTradableAmount == maxTradableAmount {
			prices = append(prices, value.price)
		}
	}

	// get the maximum volume price from the average of the minimum and maximum tradable price levels
	if len(prices) > 0 {
		return (prices[len(prices)-1] + prices[0]) / 2
	}
	return 0
}

// buildCumulativePriceLevels this returns a slice of all the price levels with the
// cumulative volume for each level. Also returns the max tradable size
func (b *OrderBook) buildCumulativePriceLevels(maxPrice, minPrice uint64) ([]CumulativeVolumeLevel, uint64) {
	type maybePriceLevel struct {
		price  uint64
		buypl  *PriceLevel
		sellpl *PriceLevel
	}

	// we'll keep track of all the pl we encounter
	mplm := map[uint64]maybePriceLevel{}

	for i := len(b.buy.levels) - 1; i >= 0; i-- {
		if b.buy.levels[i].price < minPrice {
			break
		}

		mplm[b.buy.levels[i].price] = maybePriceLevel{price: b.buy.levels[i].price, buypl: b.buy.levels[i]}
	}

	// now we add all the sells
	// to our list of pricelevel
	// making sure we have no duplicates
	for i := len(b.sell.levels) - 1; i >= 0; i-- {
		var price = b.sell.levels[i].price
		if price > maxPrice {
			break
		}

		if mpl, ok := mplm[price]; ok {
			mpl.sellpl = b.sell.levels[i]
			mplm[price] = mpl
		} else {
			mplm[price] = maybePriceLevel{price: price, sellpl: b.sell.levels[i]}
		}
	}

	// now we insert them all in the slice.
	// so we can sort them
	mpls := make([]maybePriceLevel, 0, len(mplm))
	for _, v := range mplm {
		mpls = append(mpls, v)
	}

	// sort the slice so we can go through each levels nicely
	sort.Slice(mpls, func(i, j int) bool { return mpls[i].price > mpls[j].price })

	// now we iterate other all the OK price levels
	var (
		cumulativeVolumeSell, cumulativeVolumeBuy, maxTradable uint64
		cumulativeVolumes                                      = make([]CumulativeVolumeLevel, len(mpls))
		ln                                                     = len(mpls) - 1
	)
	for i := ln; i >= 0; i-- {
		j := ln - i
		cumulativeVolumes[i].price = mpls[i].price
		if mpls[j].buypl != nil {
			cumulativeVolumeBuy += mpls[j].buypl.volume
			cumulativeVolumes[j].bidVolume = mpls[j].buypl.volume
		}

		if mpls[i].sellpl != nil {
			cumulativeVolumeSell += mpls[i].sellpl.volume
			cumulativeVolumes[i].askVolume = mpls[i].sellpl.volume

		}
		cumulativeVolumes[j].cumulativeBidVolume = cumulativeVolumeBuy
		cumulativeVolumes[i].cumulativeAskVolume = cumulativeVolumeSell

		cumulativeVolumes[i].maxTradableAmount = min(cumulativeVolumes[i].cumulativeAskVolume, cumulativeVolumes[i].cumulativeBidVolume)
		cumulativeVolumes[j].maxTradableAmount = min(cumulativeVolumes[j].cumulativeAskVolume, cumulativeVolumes[j].cumulativeBidVolume)
		maxTradable = max(maxTradable, max(cumulativeVolumes[i].maxTradableAmount, cumulativeVolumes[j].maxTradableAmount))
	}

	return cumulativeVolumes, maxTradable
}

// Uncrosses the book to generate the maximum volume set of trades
func (b *OrderBook) uncrossBook() ([]*types.OrderConfirmation, error) {
	// Get the uncrossing price and which side has the most volume at that price
	price, volume, uncrossSide := b.GetIndicativePriceAndVolume()

	// If we have no uncrossing price, we have nothing to do
	if price == 0 && volume == 0 {
		return nil, nil
	}

	var uncrossedOrder *types.OrderConfirmation
	var allOrders []*types.OrderConfirmation

	// Remove all the orders from that side of the book up to the given volume
	if uncrossSide == types.Side_SIDE_BUY {
		// Pull out the trades we want to process
		uncrossOrders, err := b.buy.ExtractOrders(price, volume)
		if err != nil {
			return nil, err
		}

		// Uncross each one
		for _, order := range uncrossOrders {
			trades, affectedOrders, _, err := b.sell.uncross(order, false)

			if err != nil {
				return nil, err
			}
			// Update all the trades to have the correct uncrossing price
			for index := 0; index < len(trades); index++ {
				trades[index].Price = price
			}
			// If the affected order is fully filled set the status
			for _, affectedOrder := range affectedOrders {
				if affectedOrder.Remaining == 0 {
					affectedOrder.Status = types.Order_STATUS_FILLED
				}
			}
			uncrossedOrder = &types.OrderConfirmation{Order: order, PassiveOrdersAffected: affectedOrders, Trades: trades}
			allOrders = append(allOrders, uncrossedOrder)
		}
	} else {
		// Pull out the trades we want to process
		uncrossOrders, err := b.sell.ExtractOrders(price, volume)
		if err != nil {
			return nil, err
		}

		// Uncross each one
		for _, order := range uncrossOrders {
			trades, affectedOrders, _, err := b.buy.uncross(order, false)

			if err != nil {
				return nil, err
			}
			// Update all the trades to have the correct uncrossing price
			for index := 0; index < len(trades); index++ {
				trades[index].Price = price
			}
			// If the affected order is fully filled set the status
			for _, affectedOrder := range affectedOrders {
				if affectedOrder.Remaining == 0 {
					affectedOrder.Status = types.Order_STATUS_FILLED
				}
			}
			uncrossedOrder = &types.OrderConfirmation{Order: order, PassiveOrdersAffected: affectedOrders, Trades: trades}
			allOrders = append(allOrders, uncrossedOrder)
		}
	}
	return allOrders, nil
}

func (b *OrderBook) GetOrdersPerParty(party string) []*types.Order {
	orderIDs := b.ordersPerParty[party]
	if len(orderIDs) <= 0 {
		return []*types.Order{}
	}

	orders := make([]*types.Order, 0, len(orderIDs))
	for oid := range orderIDs {
		orders = append(orders, b.ordersByID[oid])
	}
	return orders
}

// BestBidPriceAndVolume : Return the best bid and volume for the buy side of the book
func (b *OrderBook) BestBidPriceAndVolume() (uint64, uint64, error) {
	return b.buy.BestPriceAndVolume()
}

// BestOfferPriceAndVolume : Return the best bid and volume for the sell side of the book
func (b *OrderBook) BestOfferPriceAndVolume() (uint64, uint64, error) {
	return b.sell.BestPriceAndVolume()
}

func (b *OrderBook) CancelAllOrders(party string) ([]*types.OrderCancellationConfirmation, error) {
	var (
		orders = b.GetOrdersPerParty(party)
		confs  = []*types.OrderCancellationConfirmation{}
		conf   *types.OrderCancellationConfirmation
		err    error
	)

	for _, o := range orders {
		conf, err = b.CancelOrder(o)
		if err != nil {
			return nil, err
		}
		confs = append(confs, conf)
	}

	return confs, err
}

// CancelOrder cancel an order that is active on an order book. Market and Order ID are validated, however the order must match
// the order on the book with respect to side etc. The caller will typically validate this by using a store, we should
// not trust that the external world can provide these values reliably.
func (b *OrderBook) CancelOrder(order *types.Order) (*types.OrderCancellationConfirmation, error) {
	// Validate Market
	if order.MarketId != b.marketID {
		if b.log.GetLevel() == logging.DebugLevel {
			b.log.Debug("Market ID mismatch",
				logging.Order(*order),
				logging.String("order-book", b.marketID))
		}
		return nil, types.OrderError_ORDER_ERROR_INVALID_MARKET_ID
	}

	// Validate Order ID must be present
	if err := validateOrderID(order.Id); err != nil {
		if b.log.GetLevel() == logging.DebugLevel {
			b.log.Debug("Order ID missing or invalid",
				logging.Order(*order),
				logging.String("order-book", b.marketID))
		}
		return nil, err
	}

	order, err := b.DeleteOrder(order)
	if err != nil {
		return nil, err
	}

	// Important to mark the order as cancelled (and no longer active)
	order.Status = types.Order_STATUS_CANCELLED

	result := &types.OrderCancellationConfirmation{
		Order: order,
	}
	return result, nil
}

// RemoveOrder takes the order off the order book
func (b *OrderBook) RemoveOrder(order *types.Order) error {
	order, err := b.DeleteOrder(order)
	if err != nil {
		return err
	}

	// Important to mark the order as parked (and no longer active)
	order.Status = types.Order_STATUS_PARKED

	return nil
}

// AmendOrder amends an order which is an active order on the book
func (b *OrderBook) AmendOrder(originalOrder, amendedOrder *types.Order) error {
	if originalOrder == nil {
		return types.ErrOrderNotFound
	}

	// If the creation date for the 2 orders is different, something went wrong
	if originalOrder.CreatedAt != amendedOrder.CreatedAt {
		return types.ErrOrderOutOfSequence
	}

	if err := b.validateOrder(amendedOrder); err != nil {
		if b.log.GetLevel() == logging.DebugLevel {
			b.log.Debug("Order validation failure",
				logging.Order(*amendedOrder),
				logging.Error(err),
				logging.String("order-book", b.marketID))
		}
		return err
	}

	if amendedOrder.Side == types.Side_SIDE_BUY {
		if err := b.buy.amendOrder(amendedOrder); err != nil {
			if b.log.GetLevel() == logging.DebugLevel {
				b.log.Debug("Failed to amend (buy side)",
					logging.Order(*amendedOrder),
					logging.Error(err),
					logging.String("order-book", b.marketID))
			}
			return err
		}
	} else {
		if err := b.sell.amendOrder(amendedOrder); err != nil {
			if b.log.GetLevel() == logging.DebugLevel {
				b.log.Debug("Failed to amend (sell side)",
					logging.Order(*amendedOrder),
					logging.Error(err),
					logging.String("order-book", b.marketID))
			}
			return err
		}
	}

	return nil
}

// GetTrades returns the trades a given order generates if we were to submit it now
// this is used to calculate fees, perform price monitoring, etc...
func (b *OrderBook) GetTrades(order *types.Order) ([]*types.Trade, error) {
	if err := b.validateOrder(order); err != nil {
		return nil, err
	}
	if order.CreatedAt > b.latestTimestamp {
		b.latestTimestamp = order.CreatedAt
	}

	if b.auction {
		return nil, nil
	}

	_, trades, err := b.getOppositeSide(order.Side).fakeUncross(order)

	if err != nil {
		if err == ErrWashTrade {
			// we still want to submit this order, there might be trades coming out of it
			return trades, nil
		}
		// some random error happened, return both trades and error
		// this is a case that isn't covered by the current SubmitOrder call
		return trades, err
	}
	// no error uncrossing, in all other cases, return trades (could be empty) without an error
	return trades, nil
}

// SubmitOrder Add an order and attempt to uncross the book, returns a TradeSet protobuf message object
func (b *OrderBook) SubmitOrder(order *types.Order) (*types.OrderConfirmation, error) {
	timer := metrics.NewTimeCounter(b.marketID, "matching", "SubmitOrder")

	if err := b.validateOrder(order); err != nil {
		timer.EngineTimeCounterAdd()
		return nil, err
	}

	if order.CreatedAt > b.latestTimestamp {
		b.latestTimestamp = order.CreatedAt
	}

	if b.LogPriceLevelsDebug {
		b.PrintState("Entry state:")
	}

	var trades []*types.Trade
	var impactedOrders []*types.Order
	var lastTradedPrice uint64
	var err error

	order.BatchId = b.batchID

	if !b.auction {
		// uncross with opposite
		trades, impactedOrders, lastTradedPrice, err = b.getOppositeSide(order.Side).uncross(order, true)
		if lastTradedPrice != 0 {
			b.lastTradedPrice = lastTradedPrice
		}
		// if state of the book changed show state
		if b.LogPriceLevelsDebug && len(trades) != 0 {
			b.PrintState("After uncross state:")
		}
	}

	// if order is persistent type add to order book to the correct side
	// and we did not hit a error / wash trade error
	if order.IsPersistent() && err == nil {
		b.getSide(order.Side).addOrder(order)

		if b.LogPriceLevelsDebug {
			b.PrintState("After addOrder state:")
		}
	}

	// Was the aggressive order fully filled?
	if order.Remaining == 0 {
		order.Status = types.Order_STATUS_FILLED
	}

	// What is an Immediate or Cancel Order?
	// An immediate or cancel order (IOC) is an order to buy or sell that executes all
	// or part immediately and cancels any unfilled portion of the order.
	if order.TimeInForce == types.Order_TIME_IN_FORCE_IOC && order.Remaining > 0 {
		// Stopped as not filled at all
		if order.Remaining == order.Size {
			order.Status = types.Order_STATUS_STOPPED
		} else {
			// IOC so we set status as Cancelled.
			order.Status = types.Order_STATUS_PARTIALLY_FILLED
		}
	}

	// What is Fill Or Kill?
	// Fill or kill (FOK) is a type of time-in-force designation used in trading that instructs
	// the protocol to execute an order immediately and completely or not at all.
	// The order must be filled in its entirety or cancelled (killed).
	if order.TimeInForce == types.Order_TIME_IN_FORCE_FOK && order.Remaining == order.Size {
		// FOK and didnt trade at all we set status as Stopped
		order.Status = types.Order_STATUS_STOPPED
	}

	for idx := range impactedOrders {
		if impactedOrders[idx].Remaining == 0 {
			impactedOrders[idx].Status = types.Order_STATUS_FILLED

			// delete from lookup table
			delete(b.ordersByID, impactedOrders[idx].Id)
			delete(b.ordersPerParty[impactedOrders[idx].PartyId], impactedOrders[idx].Id)
		}
	}

	// if we did hit a wash trade, set the status to STOPPED
	if err == ErrWashTrade {
		if order.Size > order.Remaining {
			order.Status = types.Order_STATUS_PARTIALLY_FILLED
		} else {
			order.Status = types.Order_STATUS_STOPPED
		}
		order.Reason = types.OrderError_ORDER_ERROR_SELF_TRADING
	}

	if order.Status == types.Order_STATUS_ACTIVE {
		b.ordersByID[order.Id] = order
		if orders, ok := b.ordersPerParty[order.PartyId]; !ok {
			b.ordersPerParty[order.PartyId] = map[string]struct{}{
				order.Id: {},
			}
		} else {
			orders[order.Id] = struct{}{}
		}
	}

	orderConfirmation := makeResponse(order, trades, impactedOrders)
	timer.EngineTimeCounterAdd()
	return orderConfirmation, nil
}

// DeleteOrder remove a given order on a given side from the book
func (b *OrderBook) DeleteOrder(
	order *types.Order) (*types.Order, error) {
	dorder, err := b.getSide(order.Side).RemoveOrder(order)
	if err != nil {
		if b.log.GetLevel() == logging.DebugLevel {
			b.log.Debug("Failed to remove order",
				logging.Order(*order),
				logging.Error(err),
				logging.String("order-book", b.marketID))
		}
		return nil, types.ErrOrderRemovalFailure
	}
	delete(b.ordersByID, order.Id)
	delete(b.ordersPerParty[order.PartyId], order.Id)
	return dorder, err
}

// GetOrderByID returns order by its ID (IDs are not expected to collide within same market)
func (b *OrderBook) GetOrderByID(orderID string) (*types.Order, error) {
	if err := validateOrderID(orderID); err != nil {
		if b.log.GetLevel() == logging.DebugLevel {
			b.log.Debug("Order ID missing or invalid",
				logging.String("order-id", orderID))
		}
		return nil, err
	}
	// First look for the order in the order book
	order, exists := b.ordersByID[orderID]
	if !exists {
		return nil, ErrOrderDoesNotExist
	}
	return order, nil
}

// RemoveDistressedOrders remove from the book all order holding distressed positions
func (b *OrderBook) RemoveDistressedOrders(
	parties []events.MarketPosition,
) ([]*types.Order, error) {
	rmorders := []*types.Order{}

	for _, party := range parties {
		orders := []*types.Order{}
		for _, l := range b.buy.levels {
			rm := l.getOrdersByParty(party.Party())
			orders = append(orders, rm...)
		}
		for _, l := range b.sell.levels {
			rm := l.getOrdersByParty(party.Party())
			orders = append(orders, rm...)
		}
		for _, o := range orders {
			confirm, err := b.CancelOrder(o)
			if err != nil {
				if b.log.GetLevel() == logging.DebugLevel {
					b.log.Debug(
						"Failed to cancel a given order for party",
						logging.Order(*o),
						logging.String("party", party.Party()),
						logging.Error(err))
				}
				// let's see whether we need to handle this further down
				continue
			}
			// here we set the status of the order as stopped as the system triggered it as well.
			confirm.Order.Status = types.Order_STATUS_STOPPED
			rmorders = append(rmorders, confirm.Order)
		}
	}
	return rmorders, nil
}

func (b OrderBook) getSide(orderSide types.Side) *OrderBookSide {
	if orderSide == types.Side_SIDE_BUY {
		return b.buy
	}
	return b.sell
}

func (b *OrderBook) getOppositeSide(orderSide types.Side) *OrderBookSide {
	if orderSide == types.Side_SIDE_BUY {
		return b.sell
	}
	return b.buy
}

func makeResponse(order *types.Order, trades []*types.Trade, impactedOrders []*types.Order) *types.OrderConfirmation {

	return &types.OrderConfirmation{
		Order:                 order,
		PassiveOrdersAffected: impactedOrders,
		Trades:                trades,
	}
}

func (b *OrderBook) GetBestBidPrice() (uint64, error) {
	price, _, err := b.buy.BestPriceAndVolume()
	return price, err
}

func (b *OrderBook) GetBestStaticBidPrice() (uint64, error) {
	return b.buy.BestStaticPrice()
}

func (b *OrderBook) GetBestStaticBidPriceAndVolume() (uint64, uint64, error) {
	return b.buy.BestStaticPriceAndVolume()
}

func (b *OrderBook) GetBestAskPrice() (uint64, error) {
	price, _, err := b.sell.BestPriceAndVolume()
	return price, err
}

func (b *OrderBook) GetBestStaticAskPrice() (uint64, error) {
	return b.sell.BestStaticPrice()
}

func (b *OrderBook) GetBestStaticAskPriceAndVolume() (uint64, uint64, error) {
	return b.sell.BestStaticPriceAndVolume()
}

// PrintState prints the actual state of the book.
// this should be use only in debug / non production environment as it
// rely a lot on logging
func (b *OrderBook) PrintState(types string) {
	b.log.Debug("PrintState",
		logging.String("types", types))
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

// GetTotalNumberOfOrders is a debug/testing function to return the total number of orders in the book
func (b *OrderBook) GetTotalNumberOfOrders() int64 {
	return b.buy.getOrderCount() + b.sell.getOrderCount()
}

// GetTotalVolume is a debug/testing function to return the total volume in the order book
func (b *OrderBook) GetTotalVolume() int64 {
	return b.buy.getTotalVolume() + b.sell.getTotalVolume()
}
