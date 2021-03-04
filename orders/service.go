package orders

import (
	"context"
	"sync/atomic"
	"time"

	types "code.vegaprotocol.io/vega/proto"

	"code.vegaprotocol.io/vega/contextutil"
	"code.vegaprotocol.io/vega/logging"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

var (
	// ErrInvalidExpirationDT signals that the time format was invalid
	ErrInvalidExpirationDT = errors.New("invalid expiration datetime (cannot be in the past)")
	// ErrInvalidTimeInForceForMarketOrder signals that the time in force used with a market order is not accepted
	ErrInvalidTimeInForceForMarketOrder = errors.New("invalid time in force for market order (expected: FOK/IOC)")
	// ErrInvalidPriceForLimitOrder signal that a a price was missing for a limit order
	ErrInvalidPriceForLimitOrder = errors.New("invalid limit order (missing required price)")
	// ErrInvalidPriceForMarketOrder signals that a price was specified for a market order but not price is required
	ErrInvalidPriceForMarketOrder = errors.New("invalid market order (no price required)")
	// ErrNonGTTOrderWithExpiry signals that a non GTT order what set with an expiracy
	ErrNonGTTOrderWithExpiry = errors.New("non GTT order with expiry")
	// ErrGTTOrderWithNoExpiry signals that a GTT order was set without an expiracy
	ErrGTTOrderWithNoExpiry = errors.New("GTT order without expiry")
	// ErrEmptyPrepareRequest empty prepare request
	ErrEmptyPrepareRequest = errors.New("empty prepare request")
	// ErrNoParamsInAmendRequest no amended fields have been provided
	ErrNoParamsInAmendRequest = errors.New("no amended fields have been provided")
	// ErrNoTimeInForce no value has been set for the time in force
	ErrNoTimeInForce = errors.New("no value has been set for the time in force")
	// ErrNoSide no value has been set for the side
	ErrNoSide = errors.New("no value has been set for the side")
	// ErrNoType no value has been set for the type
	ErrNoType = errors.New("no value has been set for the type")
	// ErrUnAuthorizedOrderType order type is not allowed (most likely NETWORK)
	ErrUnAuthorizedOrderType = errors.New("unauthorized order type")
	// ErrCancelOrderWithOrderID RequireMarketID a cancel order request with an orderID specified requires the marketID in which the order exists
	ErrCancelOrderWithOrderIDRequireMarketID = errors.New("cancel order with orderID require marketID")
	// ErrCannotAmendToGFA it is not allowed to amend an order to GFA time in force
	ErrCannotAmendToGFA = errors.New("cannot amend to time in force GFA")
	// ErrCannotAmendToGFN it is not allowed to amend an order to GFN time in force
	ErrCannotAmendToGFN = errors.New("cannot amend to time in force GFN")
	// ErrPeggedOrderMustBeLimitOrder pegged order must be limit orders
	ErrPeggedOrderMustBeLimitOrder = errors.New("pegged orders must be limit orders")
	// ErrPeggedOrderMustBeGTTOrGTC pegged orders must be GTT or GTC orders
	ErrPeggedOrderMustBeGTTOrGTC = errors.New("pegged orders must be GTT or GTC orders")
	// ErrPeggedOrderWithoutReferencePrice pegged order message with no reference price
	ErrPeggedOrderWithoutReferencePrice = errors.New("pegged order missing a reference price")
	// ErrPeggedOrderBuyCannotReferenceBestAskPrice pegged buy order cannot reference best ask
	ErrPeggedOrderBuyCannotReferenceBestAskPrice = errors.New("pegged buy order cannot reference best ask")
	// ErrPeggedOrderOffsetMustBeLessOrEqualToZero pegged order offset must be <= 0
	ErrPeggedOrderOffsetMustBeLessOrEqualToZero = errors.New("pegged order offset must be <= 0")
	// ErrPeggedOrderOffsetMustBeLessThanZero pegged order offset must be < 0
	ErrPeggedOrderOffsetMustBeLessThanZero = errors.New("pegged order offset must be < 0")
	// ErrPeggedOrderOffsetMustBeGreaterOrEqualToZero pegged order offset must be >= 0
	ErrPeggedOrderOffsetMustBeGreaterOrEqualToZero = errors.New("pegged order offset must be >= 0")
	// ErrPeggedOrderSellCannotReferenceBestBidPrice pegged sell order cannot reference best bid
	ErrPeggedOrderSellCannotReferenceBestBidPrice = errors.New("pegged sell order cannot reference best bid")
	// ErrPeggedOrderOffsetMustBeGreaterThanZero pegged order offset must be > 0
	ErrPeggedOrderOffsetMustBeGreaterThanZero = errors.New("pegged order offset must be > 0")
)

// TimeService ...
//go:generate go run github.com/golang/mock/mockgen -destination mocks/time_service_mock.go -package mocks code.vegaprotocol.io/vega/orders TimeService
type TimeService interface {
	GetTimeNow() (time.Time, error)
}

// OrderStore ...
//go:generate go run github.com/golang/mock/mockgen -destination mocks/order_store_mock.go -package mocks code.vegaprotocol.io/vega/orders  OrderStore
type OrderStore interface {
	GetByMarketAndID(ctx context.Context, market string, id string) (*types.Order, error)
	GetByPartyAndID(ctx context.Context, party, id string) (*types.Order, error)
	GetByMarket(ctx context.Context, market string, skip, limit uint64, descending bool) ([]*types.Order, error)
	GetByParty(ctx context.Context, party string, skip, limit uint64, descending bool) ([]*types.Order, error)
	GetByReference(ctx context.Context, ref string) (*types.Order, error)
	GetByOrderID(ctx context.Context, id string, version *uint64) (*types.Order, error)
	GetAllVersionsByOrderID(ctx context.Context, id string, skip, limit uint64, descending bool) ([]*types.Order, error)
	Subscribe(orders chan<- []types.Order) uint64
	Unsubscribe(id uint64) error
}

// Svc represents the order service
type Svc struct {
	Config
	log *logging.Logger

	orderStore    OrderStore
	timeService   TimeService
	subscriberCnt int32
}

// NewService creates an Orders service with the necessary dependencies
func NewService(log *logging.Logger, config Config, store OrderStore, time TimeService) (*Svc, error) {
	// setup logger
	log = log.Named(namedLogger)
	log.SetLevel(config.Level.Get())

	return &Svc{
		log:         log,
		Config:      config,
		orderStore:  store,
		timeService: time,
	}, nil
}

// ReloadConf update the internal configuration of the order service
func (s *Svc) ReloadConf(cfg Config) {
	s.log.Info("reloading configuration")
	if s.log.GetLevel() != cfg.Level.Get() {
		s.log.Info("updating log level",
			logging.String("old", s.log.GetLevel().String()),
			logging.String("new", cfg.Level.String()),
		)
		s.log.SetLevel(cfg.Level.Get())
	}

	s.Config = cfg
}

func (s *Svc) PrepareSubmitOrder(ctx context.Context, submission *types.OrderSubmission) error {
	if submission == nil {
		return ErrEmptyPrepareRequest
	}
	if err := s.validateOrderSubmission(submission); err != nil {
		return err
	}
	if submission.Reference == "" {
		submission.Reference = uuid.NewV4().String()
	}
	return nil
}

func (s *Svc) validateOrderSubmission(sub *types.OrderSubmission) error {
	if err := sub.Validate(); err != nil {
		return errors.Wrap(err, "order validation failed")
	}

	if sub.Side == types.Side_SIDE_UNSPECIFIED {
		return ErrNoSide
	}

	if sub.Type == types.Order_TYPE_UNSPECIFIED {
		return ErrNoType
	}

	if sub.TimeInForce == types.Order_TIME_IN_FORCE_UNSPECIFIED {
		return ErrNoTimeInForce
	}

	if sub.TimeInForce == types.Order_TIME_IN_FORCE_GTT {
		if sub.ExpiresAt <= 0 {
			s.log.Error("invalid expiration time")
			return ErrInvalidExpirationDT
		}
	}

	if sub.TimeInForce != types.Order_TIME_IN_FORCE_GTT && sub.ExpiresAt != 0 {
		return ErrNonGTTOrderWithExpiry
	}

	if sub.Type == types.Order_TYPE_MARKET && sub.Price != 0 {
		return ErrInvalidPriceForMarketOrder
	}
	if sub.Type == types.Order_TYPE_MARKET &&
		(sub.TimeInForce != types.Order_TIME_IN_FORCE_FOK && sub.TimeInForce != types.Order_TIME_IN_FORCE_IOC) {
		return ErrInvalidTimeInForceForMarketOrder
	}
	if sub.Type == types.Order_TYPE_LIMIT && sub.Price == 0 &&
		sub.PeggedOrder == nil {
		return ErrInvalidPriceForLimitOrder
	}
	if sub.Type == types.Order_TYPE_NETWORK {
		return ErrUnAuthorizedOrderType
	}

	// Validation for pegged orders
	if sub.PeggedOrder != nil {
		if sub.Type != types.Order_TYPE_LIMIT {
			// All pegged orders must be LIMIT orders
			return ErrPeggedOrderMustBeLimitOrder
		}

		if sub.TimeInForce != types.Order_TIME_IN_FORCE_GTT && sub.TimeInForce != types.Order_TIME_IN_FORCE_GTC {
			// Pegged orders can only be GTC or GTT
			return ErrPeggedOrderMustBeGTTOrGTC
		}

		if sub.PeggedOrder.Reference == types.PeggedReference_PEGGED_REFERENCE_UNSPECIFIED {
			// We must specify a valid reference
			return ErrPeggedOrderWithoutReferencePrice
		}

		if sub.Side == types.Side_SIDE_BUY {
			switch sub.PeggedOrder.Reference {
			case types.PeggedReference_PEGGED_REFERENCE_BEST_ASK:
				return ErrPeggedOrderBuyCannotReferenceBestAskPrice
			case types.PeggedReference_PEGGED_REFERENCE_BEST_BID:
				if sub.PeggedOrder.Offset > 0 {
					return ErrPeggedOrderOffsetMustBeLessOrEqualToZero
				}
			case types.PeggedReference_PEGGED_REFERENCE_MID:
				if sub.PeggedOrder.Offset >= 0 {
					return ErrPeggedOrderOffsetMustBeLessThanZero
				}
			}
		} else {
			switch sub.PeggedOrder.Reference {
			case types.PeggedReference_PEGGED_REFERENCE_BEST_ASK:
				if sub.PeggedOrder.Offset < 0 {
					return ErrPeggedOrderOffsetMustBeGreaterOrEqualToZero
				}
			case types.PeggedReference_PEGGED_REFERENCE_BEST_BID:
				return ErrPeggedOrderSellCannotReferenceBestBidPrice
			case types.PeggedReference_PEGGED_REFERENCE_MID:
				if sub.PeggedOrder.Offset <= 0 {
					return ErrPeggedOrderOffsetMustBeGreaterThanZero
				}
			}
		}
	}
	return nil
}

func (s *Svc) PrepareCancelOrder(ctx context.Context, order *types.OrderCancellation) error {
	if order == nil {
		return ErrEmptyPrepareRequest
	}
	if err := order.Validate(); err != nil {
		return errors.Wrap(err, "order cancellation invalid")
	}

	// ensure that if orderID is specified marketId is as well
	if len(order.OrderId) > 0 && len(order.MarketId) <= 0 {
		return ErrCancelOrderWithOrderIDRequireMarketID
	}

	return nil
}

func (s *Svc) PrepareAmendOrder(ctx context.Context, amendment *types.OrderAmendment) error {
	if amendment == nil {
		return ErrEmptyPrepareRequest
	}
	if err := amendment.Validate(); err != nil {
		return errors.Wrap(err, "order amendment validation failed")
	}

	// Check we are not trying to amend to a GFA
	if amendment.TimeInForce == types.Order_TIME_IN_FORCE_GFA {
		return ErrCannotAmendToGFA
	}

	// Check we are not trying to amend to a GFN
	if amendment.TimeInForce == types.Order_TIME_IN_FORCE_GFN {
		return ErrCannotAmendToGFN
	}

	// Check we have at least one field to update
	if amendment.Price == nil &&
		amendment.SizeDelta == 0 &&
		(amendment.ExpiresAt == nil || amendment.ExpiresAt.Value == 0) &&
		amendment.TimeInForce == types.Order_TIME_IN_FORCE_UNSPECIFIED &&
		amendment.PeggedOffset == nil &&
		amendment.PeggedReference == types.PeggedReference_PEGGED_REFERENCE_UNSPECIFIED {
		return ErrNoParamsInAmendRequest
	}

	// Only update ExpiresAt when TIME_IN_FORCE is related
	if amendment.ExpiresAt != nil && amendment.ExpiresAt.Value > 0 {
		if amendment.TimeInForce != types.Order_TIME_IN_FORCE_GTT &&
			amendment.TimeInForce != types.Order_TIME_IN_FORCE_UNSPECIFIED {
			// We cannot change the expire time for this order type
			return ErrNonGTTOrderWithExpiry
		}
	}

	// if order is GTT convert datetime to blockchain ts
	if amendment.TimeInForce == types.Order_TIME_IN_FORCE_GTT {
		if amendment.ExpiresAt == nil {
			s.log.Error("unable to set trade type to GTT when no expiry given")
			return ErrGTTOrderWithNoExpiry
		}
	}
	return nil
}

// GetByOrderID find an order using its orderID
func (s *Svc) GetByOrderID(ctx context.Context, id string, version uint64) (order *types.Order, err error) {
	if version == 0 {
		return s.orderStore.GetByOrderID(ctx, id, nil)
	}
	return s.orderStore.GetByOrderID(ctx, id, &version)
}

// GetByReference find an order from the store using its reference
func (s *Svc) GetByReference(ctx context.Context, ref string) (*types.Order, error) {
	return s.orderStore.GetByReference(ctx, ref)
}

// GetByMarket returns a list of order for a given market
func (s *Svc) GetByMarket(ctx context.Context, market string, skip, limit uint64, descending bool) (orders []*types.Order, err error) {
	return s.orderStore.GetByMarket(ctx, market, skip, limit, descending)
}

// GetByParty returns a list of order for a given party
func (s *Svc) GetByParty(ctx context.Context, party string, skip, limit uint64, descending bool) (orders []*types.Order, err error) {
	return s.orderStore.GetByParty(ctx, party, skip, limit, descending)
}

// GetByMarketAndID find a order using a marketID and an order id
func (s *Svc) GetByMarketAndID(ctx context.Context, market string, id string) (order *types.Order, err error) {
	o, err := s.orderStore.GetByMarketAndID(ctx, market, id)
	if err != nil {
		return &types.Order{}, err
	}
	return o, err
}

// GetByPartyAndID find an order using a party id and an order id
func (s *Svc) GetByPartyAndID(ctx context.Context, party string, id string) (order *types.Order, err error) {
	o, err := s.orderStore.GetByPartyAndID(ctx, party, id)
	if err != nil {
		return &types.Order{}, err
	}
	return o, err
}

// GetAllVersionsByOrderID returns all available versions for the order specified by id
func (s *Svc) GetAllVersionsByOrderID(ctx context.Context, id string, skip, limit uint64, descending bool) (orders []*types.Order, err error) {
	return s.orderStore.GetAllVersionsByOrderID(ctx, id, skip, limit, descending)
}

// GetOrderSubscribersCount return the number of subscribers to the
// orders updates stream
func (s *Svc) GetOrderSubscribersCount() int32 {
	return atomic.LoadInt32(&s.subscriberCnt)
}

// ObserveOrders add a new subscriber to the stream of order updates
func (s *Svc) ObserveOrders(ctx context.Context, retries int, market *string, party *string) (<-chan []types.Order, uint64) {
	orders := make(chan []types.Order)
	internal := make(chan []types.Order)
	ref := s.orderStore.Subscribe(internal)

	var cancel func()
	ctx, cancel = context.WithCancel(ctx)
	go func() {
		atomic.AddInt32(&s.subscriberCnt, 1)
		defer atomic.AddInt32(&s.subscriberCnt, -1)
		ip, _ := contextutil.RemoteIPAddrFromContext(ctx)
		defer cancel()
		for {
			select {
			case <-ctx.Done():
				s.log.Debug(
					"Orders subscriber closed connection",
					logging.Uint64("id", ref),
					logging.String("ip-address", ip),
				)
				// this error only happens when the subscriber reference doesn't exist
				// so we can still safely close the channels
				if err := s.orderStore.Unsubscribe(ref); err != nil {
					s.log.Error(
						"Failure un-subscribing orders subscriber when context.Done()",
						logging.Uint64("id", ref),
						logging.String("ip-address", ip),
						logging.Error(err),
					)
				}
				close(internal)
				close(orders)
				return
			case v := <-internal:
				// max cap for this slice is the length of the slice we read from the channel
				validatedOrders := make([]types.Order, 0, len(v))
				for _, item := range v {
					// if market is not set, or equals item market and party is not set or equals item party
					if (market == nil || item.MarketId == *market) && (party == nil || item.PartyId == *party) {
						validatedOrders = append(validatedOrders, item)
					}
				}
				if len(validatedOrders) == 0 {
					continue
				}
				retryCount := retries
				success := false
				for !success && retryCount >= 0 {
					select {
					case orders <- validatedOrders:
						s.log.Debug(
							"Orders for subscriber sent successfully",
							logging.Uint64("ref", ref),
							logging.String("ip-address", ip),
						)
						success = true
					default:
						retryCount--
						if retryCount >= 0 {
							s.log.Debug(
								"Orders for subscriber not sent",
								logging.Uint64("ref", ref),
								logging.String("ip-address", ip),
							)
						}
						time.Sleep(time.Duration(10) * time.Millisecond)
					}
				}
				if !success && retryCount <= 0 {
					s.log.Warn(
						"Order subscriber has hit the retry limit",
						logging.Uint64("ref", ref),
						logging.String("ip-address", ip),
						logging.Int("retries", retries),
					)
					cancel()
				}
			}
		}
	}()

	return orders, ref
}
