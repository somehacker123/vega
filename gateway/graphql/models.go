// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gql

import (
	"fmt"
	"io"
	"strconv"

	"code.vegaprotocol.io/vega/proto"
)

type Oracle interface {
	IsOracle()
}

type Product interface {
	IsProduct()
}

type ProposalChange interface {
	IsProposalChange()
}

type RiskModel interface {
	IsRiskModel()
}

type TradingMode interface {
	IsTradingMode()
}

// A mode where Vega try to execute order as soon as they are received
type ContinuousTrading struct {
	// Size of an increment in price in terms of the quote currency (uint64)
	TickSize *int `json:"tickSize"`
}

func (ContinuousTrading) IsTradingMode() {}

type ContinuousTradingInput struct {
	// Size of an increment in price in terms of the quote currency
	TickSize int `json:"tickSize"`
}

// Some non continuous trading mode
type DiscreteTrading struct {
	// Duration of the trading (uint64)
	Duration *int `json:"duration"`
}

func (DiscreteTrading) IsTradingMode() {}

type DiscreteTradingInput struct {
	// Duration of the trading
	Duration int `json:"duration"`
}

// An Ethereum oracle
type EthereumEvent struct {
	// The ID of the ethereum contract to use (string)
	ContractID string `json:"contractId"`
	// Name of the Ethereum event to listen to. (string)
	Event string `json:"event"`
}

func (EthereumEvent) IsOracle() {}

type EthereumEventInput struct {
	// The ID of the ethereum contract to use
	ContractID string `json:"contractId"`
	// Name of the Ethereum event to listen to
	Event string `json:"event"`
}

// A Future product
type Future struct {
	// The maturity date of the product (string)
	Maturity string `json:"maturity"`
	// The name of the asset (string)
	Asset string `json:"asset"`
	// The oracle used for this product (Oracle union)
	Oracle Oracle `json:"oracle"`
}

func (Future) IsProduct() {}

type FutureInput struct {
	// The maturity date of the product
	Maturity string `json:"maturity"`
	// The name of the asset
	Asset string `json:"asset"`
	// The oracle used for this product
	EthereumOracle *EthereumEventInput `json:"ethereumOracle"`
}

// Describe something that can be traded on Vega
type Instrument struct {
	// Uniquely identify an instrument accrods all instruments available on Vega (string)
	ID string `json:"id"`
	// A short non necessarily unique code used to easily describe the instrument (e.g: FX:BTCUSD/DEC18) (string)
	Code string `json:"code"`
	// Full and fairly descriptive name for the instrument
	Name string `json:"name"`
	// String representing the base (e.g. BTCUSD -> BTC is base)
	BaseName string `json:"baseName"`
	// String representing the quote (e.g. BTCUSD -> USD is quote)
	QuoteName string `json:"quoteName"`
	// Metadata for this instrument
	Metadata *InstrumentMetadata `json:"metadata"`
	// A reference to or instance of a fully specified product, including all required product parameters for that product (Product union)
	Product Product `json:"product"`
}

type InstrumentInput struct {
	// Uniquely identify an instrument accrods all instruments available on Vega
	ID string `json:"id"`
	// A short non necessarily unique code used to easily describe the instrument (e.g: FX:BTCUSD/DEC18)
	Code string `json:"code"`
	// Full and fairly descriptive name for the instrument
	Name string `json:"name"`
	// String representing the base (e.g. BTCUSD -> BTC is base)
	BaseName string `json:"baseName"`
	// String representing the quote (e.g. BTCUSD -> USD is quote)
	QuoteName        string `json:"quoteName"`
	InitialMarkPrice string `json:"initialMarkPrice"`
	// Metadata for this instrument
	Metadata *InstrumentMetadatInput `json:"metadata"`
	// A reference to or instance of a fully specified product, including all required product parameters for that product
	FutureProduct *FutureInput `json:"futureProduct"`
}

type InstrumentMetadatInput struct {
	// An arbitrary list of tags to associated to associate to the Instrument
	Tags []*string `json:"tags"`
}

// A set of metadata to associate to an instruments
type InstrumentMetadata struct {
	// An arbitrary list of tags to associated to associate to the Instrument (string list)
	Tags []*string `json:"tags"`
}

// Parameters for the log normal risk model
type LogNormalModelParams struct {
	// mu parameter
	Mu float64 `json:"mu"`
	// r parameter
	R float64 `json:"r"`
	// sigma parameter
	Sigma float64 `json:"sigma"`
}

type LogNormalModelParamsInput struct {
	// mu parameter
	Mu float64 `json:"mu"`
	// r parameter
	R float64 `json:"r"`
	// sigma parameter
	Sigma float64 `json:"sigma"`
}

// A type of risk model for futures trading
type LogNormalRiskModel struct {
	// Lambda parameter of the risk model
	RiskAversionParameter float64 `json:"riskAversionParameter"`
	// Tau parameter of the risk model
	Tau float64 `json:"tau"`
	// Params for the log normal risk model
	Params *LogNormalModelParams `json:"params"`
}

func (LogNormalRiskModel) IsRiskModel() {}

type LogNormalRiskModelInput struct {
	// Lambda parameter of the risk model
	RiskAversionParameter float64 `json:"riskAversionParameter"`
	// Tau parameter of the risk model
	Tau float64 `json:"tau"`
	// Params for the log normal risk model
	Params *LogNormalModelParamsInput `json:"params"`
}

type MarginCalculator struct {
	// The scaling factors that will be used for margin calculation
	ScalingFactors *ScalingFactors `json:"scalingFactors"`
}

type MarginCalculatorInput struct {
	// The scaling factors that will be used for margin calculation
	ScalingFactors *ScalingFactorsInput `json:"scalingFactors"`
}

// Represents a product & associated parameters that can be traded on Vega, has an associated OrderBook and Trade history
type Market struct {
	// Market ID
	ID   string `json:"id"`
	Name string `json:"name"`
	// An instance of or reference to a tradable instrument.
	TradableInstrument *TradableInstrument `json:"tradableInstrument"`
	// Definitions and required configuration for the trading mode
	TradingMode TradingMode `json:"tradingMode"`
	// decimalPlaces indicates the number of decimal places that an integer must be shifted by in order to get a correct
	// number denominated in the currency of the Market. (uint64)
	//
	// Examples:
	//   Currency     Balance  decimalPlaces  Real Balance
	//   GBP              100              0       GBP 100
	//   GBP              100              2       GBP   1.00
	//   GBP              100              4       GBP   0.01
	//   GBP                1              4       GBP   0.0001   (  0.01p  )
	//
	//   GBX (pence)      100              0       GBP   1.00     (100p     )
	//   GBX (pence)      100              2       GBP   0.01     (  1p     )
	//   GBX (pence)      100              4       GBP   0.0001   (  0.01p  )
	//   GBX (pence)        1              4       GBP   0.000001 (  0.0001p)
	DecimalPlaces int `json:"decimalPlaces"`
	// Orders on a market
	Orders []*proto.Order `json:"orders"`
	// Get account for a party or market
	Accounts []*proto.Account `json:"accounts"`
	// Trades on a market
	Trades []*proto.Trade `json:"trades"`
	// Current depth on the orderbook for this market
	Depth *proto.MarketDepth `json:"depth"`
	// Candles on a market, for the 'last' n candles, at 'interval' seconds as specified by params
	Candles []*proto.Candle `json:"candles"`
	// Query an order by reference for the given market
	OrderByReference *proto.Order `json:"orderByReference"`
	// marketData for the given market
	Data *proto.MarketData `json:"data"`
}

// Input variation of market details same to those defined in Market type
type MarketInput struct {
	ID                    string                   `json:"id"`
	Name                  string                   `json:"name"`
	TradableInstrument    *TradableInstrumentInput `json:"tradableInstrument"`
	ContinuousTradingMode *ContinuousTradingInput  `json:"continuousTradingMode"`
	DiscreteTradingMode   *DiscreteTradingInput    `json:"discreteTradingMode"`
	DecimalPlaces         int                      `json:"decimalPlaces"`
}

// Allows creating new markets on the network
type NewMarket struct {
	Market *Market `json:"market"`
}

func (NewMarket) IsProposalChange() {}

type NewMarketInput struct {
	Market *MarketInput `json:"market"`
}

type PreparedAmendOrder struct {
	// blob: the raw transaction to sign & submit
	Blob string `json:"blob"`
}

type PreparedCancelOrder struct {
	// blob: the raw transaction to sign & submit
	Blob string `json:"blob"`
}

type PreparedProposal struct {
	// Raw transaction data to sign & submit
	Blob string `json:"blob"`
	// The pending proposal
	PendingProposal *proto.GovernanceData `json:"pendingProposal"`
}

type PreparedSubmitOrder struct {
	// blob: the raw transaction to sign & submit
	Blob string `json:"blob"`
}

type PreparedVote struct {
	// Raw, serialised vote to be signed
	Blob string `json:"blob"`
	// The vote serialised in the blob field
	Vote *ProposalVote `json:"vote"`
}

type ProposalTerms struct {
	// ISO-8601 time and date when voting closes for this proposal.
	ClosingDatetime string `json:"closingDatetime"`
	// ISO-8601 time and date when this proposal is executed (if passed). Note that it has to be after closing date time.
	EnactmentDatetime string `json:"enactmentDatetime"`
	// Minimum participation stake required for this proposal to pass
	MinParticipationStake int `json:"minParticipationStake"`
	// Actual change being introduced by the proposal
	Change ProposalChange `json:"change"`
}

// Proposal terms input. Only one kind of change is expected. Proposals with no changes or more than one will not be accepted.
type ProposalTermsInput struct {
	// ISO-8601 time and date when voting closes for this proposal.
	ClosingDatetime string `json:"closingDatetime"`
	// ISO-8601 time and date when this proposal is executed (if passed). Note that it has to be after closing date time.
	EnactmentDatetime string `json:"enactmentDatetime"`
	// Minimum participation stake required for this proposal to pass
	MinParticipationStake int `json:"minParticipationStake"`
	// Optional field to define update market change. If this is set along with another change, proposal will not be accepted.
	UpdateMarket *UpdateMarketInput `json:"updateMarket"`
	// Optional field to define new market change. If this is set along with another change, proposal will not be accepted.
	NewMarket *NewMarketInput `json:"newMarket"`
	// Optional field to define an update of network parameters. If this is set along with another change, proposal will not be accepted.
	UpdateNetwork *UpdateNetworkInput `json:"updateNetwork"`
}

type ProposalVote struct {
	// Cast vote
	Vote *Vote `json:"vote"`
	// Proposal casting the vote on
	ProposalID string `json:"proposalID"`
}

type ScalingFactors struct {
	// the scaling factor that determines the margin level at which we have to search for more money
	SearchLevel float64 `json:"searchLevel"`
	// the scaling factor that determines the optimal margin level
	InitialMargin float64 `json:"initialMargin"`
	// The scaling factor that determines the overflow margin level
	CollateralRelease float64 `json:"collateralRelease"`
}

type ScalingFactorsInput struct {
	// the scaling factor that determines the margin level at which we have to search for more money
	SearchLevel float64 `json:"searchLevel"`
	// the scaling factor that determines the optimal margin level
	InitialMargin float64 `json:"initialMargin"`
	// The scaling factor that determines the overflow margin level
	CollateralRelease float64 `json:"collateralRelease"`
}

// A type of simple/dummy risk model where we can specify the risk factor long and short in params
type SimpleRiskModel struct {
	// Params for the simple risk model
	Params *SimpleRiskModelParams `json:"params"`
}

func (SimpleRiskModel) IsRiskModel() {}

type SimpleRiskModelInput struct {
	// Params for the simple risk model
	Params *SimpleRiskModelParamsInput `json:"params"`
}

// Parameters for the simple risk model
type SimpleRiskModelParams struct {
	// Risk factor for long
	FactorLong float64 `json:"factorLong"`
	// Risk factor for short
	FactorShort float64 `json:"factorShort"`
}

type SimpleRiskModelParamsInput struct {
	// Risk factor for long
	FactorLong float64 `json:"factorLong"`
	// Risk factor for short
	FactorShort float64 `json:"factorShort"`
}

// A tradable instrument is a combination of an instrument and a risk model
type TradableInstrument struct {
	// An instance of or reference to a fully specified instrument.
	Instrument *Instrument `json:"instrument"`
	// A reference to a risk model that is valid for the instrument
	RiskModel RiskModel `json:"riskModel"`
	// Margin calculation info, currently only the scaling factors (search, initial, release) for this tradable instrument
	MarginCalculator *MarginCalculator `json:"marginCalculator"`
}

// Input variation of tradable instrument details
type TradableInstrumentInput struct {
	Instrument         *InstrumentInput         `json:"instrument"`
	SimpleRiskModel    *SimpleRiskModelInput    `json:"simpleRiskModel"`
	LogNormalRiskModel *LogNormalRiskModelInput `json:"logNormalRiskModel"`
	MarginCalculator   *MarginCalculatorInput   `json:"marginCalculator"`
}

type TransactionSubmitted struct {
	Success bool `json:"success"`
}

// Incomplete change definition for governance proposal terms
// TODO: complete the type
type UpdateMarket struct {
	MarketID string `json:"marketId"`
}

func (UpdateMarket) IsProposalChange() {}

type UpdateMarketInput struct {
	MarketID string `json:"marketId"`
}

// Allows submitting a proposal for changing governance network parameters
type UpdateNetwork struct {
	// Network parameter that restricts when the earliest a proposal
	// can be set to close voting. Value represents duration in seconds.
	MinCloseInSeconds *int `json:"minCloseInSeconds"`
	// Network parameter that restricts when the latest a proposal
	// can be set to close voting. Value represents duration in seconds.
	MaxCloseInSeconds *int `json:"maxCloseInSeconds"`
	// Network parameter that restricts when the earliest a proposal
	// can be set to be executed (if that proposal passed).
	// Value represents duration in seconds.
	MinEnactInSeconds *int `json:"minEnactInSeconds"`
	// Network parameter that restricts when the latest a proposal
	// can be set to be executed (if that proposal passed).
	// Value represents duration in seconds.
	MaxEnactInSeconds *int `json:"maxEnactInSeconds"`
	// Network parameter that restricts the minimum participation stake
	// required for a proposal to pass.
	MinParticipationStake *int `json:"minParticipationStake"`
}

func (UpdateNetwork) IsProposalChange() {}

// Allows submitting a proposal for changing governance network parameters
type UpdateNetworkInput struct {
	// Network parameter that restricts when the earliest a proposal
	// can be set to close voting. Value represents duration in seconds.
	MinCloseInSeconds *int `json:"minCloseInSeconds"`
	// Network parameter that restricts when the latest a proposal
	// can be set to close voting. Value represents duration in seconds.
	MaxCloseInSeconds *int `json:"maxCloseInSeconds"`
	// Network parameter that restricts when the earliest a proposal
	// can be set to be executed (if that proposal passed).
	// Value represents duration in seconds.
	MinEnactInSeconds *int `json:"minEnactInSeconds"`
	// Network parameter that restricts when the latest a proposal
	// can be set to be executed (if that proposal passed).
	// Value represents duration in seconds.
	MaxEnactInSeconds *int `json:"maxEnactInSeconds"`
	// Network parameter that restricts the minimum participation stake
	// required for a proposal to pass.
	MinParticipationStake *int `json:"minParticipationStake"`
}

type Vote struct {
	// The vote value cast
	Value VoteValue `json:"value"`
	// The party casting the vote
	Party *proto.Party `json:"party"`
	// ISO-8601 time and date when the vote reached Vega network
	Datetime string `json:"datetime"`
}

// The various account types we have (used by collateral)
type AccountType string

const (
	// Insurance pool account - only for 'system' party
	AccountTypeInsurance AccountType = "Insurance"
	// Settlement - only for 'system' party
	AccountTypeSettlement AccountType = "Settlement"
	// Margin - The leverage account for traders
	AccountTypeMargin AccountType = "Margin"
	// General account - the account containing 'unused' collateral for traders
	AccountTypeGeneral AccountType = "General"
)

var AllAccountType = []AccountType{
	AccountTypeInsurance,
	AccountTypeSettlement,
	AccountTypeMargin,
	AccountTypeGeneral,
}

func (e AccountType) IsValid() bool {
	switch e {
	case AccountTypeInsurance, AccountTypeSettlement, AccountTypeMargin, AccountTypeGeneral:
		return true
	}
	return false
}

func (e AccountType) String() string {
	return string(e)
}

func (e *AccountType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AccountType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AccountType", str)
	}
	return nil
}

func (e AccountType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The interval for trade candles when subscribing via VEGA graphql, default is I15M
type Interval string

const (
	// 1 minute interval
	IntervalI1m Interval = "I1M"
	// 5 minute interval
	IntervalI5m Interval = "I5M"
	// 15 minute interval (default)
	IntervalI15m Interval = "I15M"
	// 1 hour interval
	IntervalI1h Interval = "I1H"
	// 6 hour interval
	IntervalI6h Interval = "I6H"
	// 1 day interval
	IntervalI1d Interval = "I1D"
)

var AllInterval = []Interval{
	IntervalI1m,
	IntervalI5m,
	IntervalI15m,
	IntervalI1h,
	IntervalI6h,
	IntervalI1d,
}

func (e Interval) IsValid() bool {
	switch e {
	case IntervalI1m, IntervalI5m, IntervalI15m, IntervalI1h, IntervalI6h, IntervalI1d:
		return true
	}
	return false
}

func (e Interval) String() string {
	return string(e)
}

func (e *Interval) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Interval(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Interval", str)
	}
	return nil
}

func (e Interval) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Valid order statuses, these determine several states for an order that cannot be expressed with other fields in Order.
type OrderStatus string

const (
	// The order is active and not cancelled or expired, it could be unfilled, partially filled or fully filled.
	// Active does not necessarily mean it's still on the order book.
	OrderStatusActive OrderStatus = "Active"
	// The order is cancelled, the order could be partially filled or unfilled before it was cancelled. It is not possible to cancel an order with 0 remaining.
	OrderStatusCancelled OrderStatus = "Cancelled"
	// This order trades any amount and as much as possible and remains on the book until it either trades completely or expires.
	OrderStatusExpired OrderStatus = "Expired"
	// This order was of type IOC or FOK and could not be processed by the matching engine due to lack of liquidity.
	OrderStatusStopped OrderStatus = "Stopped"
	// This order is fully filled with remaining equals zero.
	OrderStatusFilled OrderStatus = "Filled"
	// This order was rejected while beeing processed in the core.
	OrderStatusRejected OrderStatus = "Rejected"
)

var AllOrderStatus = []OrderStatus{
	OrderStatusActive,
	OrderStatusCancelled,
	OrderStatusExpired,
	OrderStatusStopped,
	OrderStatusFilled,
	OrderStatusRejected,
}

func (e OrderStatus) IsValid() bool {
	switch e {
	case OrderStatusActive, OrderStatusCancelled, OrderStatusExpired, OrderStatusStopped, OrderStatusFilled, OrderStatusRejected:
		return true
	}
	return false
}

func (e OrderStatus) String() string {
	return string(e)
}

func (e *OrderStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrderStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrderStatus", str)
	}
	return nil
}

func (e OrderStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Valid order types, these determine what happens when an order is added to the book
type OrderTimeInForce string

const (
	// The order either trades completely (remainingSize == 0 after adding) or not at all, does not remain on the book if it doesn't trade
	OrderTimeInForceFok OrderTimeInForce = "FOK"
	// The order trades any amount and as much as possible but does not remain on the book (whether it trades or not)
	OrderTimeInForceIoc OrderTimeInForce = "IOC"
	// This order trades any amount and as much as possible and remains on the book until it either trades completely or is cancelled
	OrderTimeInForceGtc OrderTimeInForce = "GTC"
	// This order type trades any amount and as much as possible and remains on the book until they either trade completely, are cancelled, or expires at a set time
	// NOTE: this may in future be multiple types or have sub types for orders that provide different ways of specifying expiry
	OrderTimeInForceGtt OrderTimeInForce = "GTT"
)

var AllOrderTimeInForce = []OrderTimeInForce{
	OrderTimeInForceFok,
	OrderTimeInForceIoc,
	OrderTimeInForceGtc,
	OrderTimeInForceGtt,
}

func (e OrderTimeInForce) IsValid() bool {
	switch e {
	case OrderTimeInForceFok, OrderTimeInForceIoc, OrderTimeInForceGtc, OrderTimeInForceGtt:
		return true
	}
	return false
}

func (e OrderTimeInForce) String() string {
	return string(e)
}

func (e *OrderTimeInForce) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrderTimeInForce(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrderTimeInForce", str)
	}
	return nil
}

func (e OrderTimeInForce) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type OrderType string

const (
	// the default order type
	OrderTypeMarket OrderType = "MARKET"
	// mentioned in ticket, but as yet unused order type
	OrderTypeLimit OrderType = "LIMIT"
	// Used for distressed traders, an order placed by the network to close out distressed traders
	// similar to MARKET order, only no party is attached to the order.
	OrderTypeNetwork OrderType = "NETWORK"
)

var AllOrderType = []OrderType{
	OrderTypeMarket,
	OrderTypeLimit,
	OrderTypeNetwork,
}

func (e OrderType) IsValid() bool {
	switch e {
	case OrderTypeMarket, OrderTypeLimit, OrderTypeNetwork:
		return true
	}
	return false
}

func (e OrderType) String() string {
	return string(e)
}

func (e *OrderType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrderType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrderType", str)
	}
	return nil
}

func (e OrderType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Varoius states a proposal can transition through:
//   Open ->
//       - Passed -> Enacted.
//       - Rejected.
//   Proposal can enter Failed state from any other state.
type ProposalState string

const (
	// Proposal became invalid and cannot be processed
	ProposalStateFailed ProposalState = "Failed"
	// Proposal is open for voting
	ProposalStateOpen ProposalState = "Open"
	// Proposal has gained enough support to be executed
	ProposalStatePassed ProposalState = "Passed"
	// Proposal didn't get enough votes
	ProposalStateDeclined ProposalState = "Declined"
	// Proposal has could not gain enough support to be executed
	ProposalStateRejected ProposalState = "Rejected"
	// Proposal has been executed and the changes under this proposal have now been applied
	ProposalStateEnacted ProposalState = "Enacted"
)

var AllProposalState = []ProposalState{
	ProposalStateFailed,
	ProposalStateOpen,
	ProposalStatePassed,
	ProposalStateDeclined,
	ProposalStateRejected,
	ProposalStateEnacted,
}

func (e ProposalState) IsValid() bool {
	switch e {
	case ProposalStateFailed, ProposalStateOpen, ProposalStatePassed, ProposalStateDeclined, ProposalStateRejected, ProposalStateEnacted:
		return true
	}
	return false
}

func (e ProposalState) String() string {
	return string(e)
}

func (e *ProposalState) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ProposalState(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ProposalState", str)
	}
	return nil
}

func (e ProposalState) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Reason for the order beeing rejected by the core node
type RejectionReason string

const (
	// Market id is invalid
	RejectionReasonInvalidMarketID RejectionReason = "InvalidMarketId"
	// Order id is invalid
	RejectionReasonInvalidOrderID RejectionReason = "InvalidOrderId"
	// Order is out of sequence
	RejectionReasonOrderOutOfSequence RejectionReason = "OrderOutOfSequence"
	// Remaining size in the order is invalid
	RejectionReasonInvalidRemainingSize RejectionReason = "InvalidRemainingSize"
	// Time has failed us
	RejectionReasonTimeFailure RejectionReason = "TimeFailure"
	// Unable to remove the order
	RejectionReasonOrderRemovalFailure RejectionReason = "OrderRemovalFailure"
	// Expiration time is invalid
	RejectionReasonInvalidExpirationTime RejectionReason = "InvalidExpirationTime"
	// Order reference is invalid
	RejectionReasonInvalidOrderReference RejectionReason = "InvalidOrderReference"
	// Edit is not allowed
	RejectionReasonEditNotAllowed RejectionReason = "EditNotAllowed"
	// Order amend fail
	RejectionReasonOrderAmendFailure RejectionReason = "OrderAmendFailure"
	// Order does not exist
	RejectionReasonOrderNotFound RejectionReason = "OrderNotFound"
	// Party id is invalid
	RejectionReasonInvalidPartyID RejectionReason = "InvalidPartyId"
	// Market is closed
	RejectionReasonMarketClosed RejectionReason = "MarketClosed"
	// Margin check failed
	RejectionReasonMarginCheckFailed RejectionReason = "MarginCheckFailed"
	// An internal error happend
	RejectionReasonInternalError RejectionReason = "InternalError"
)

var AllRejectionReason = []RejectionReason{
	RejectionReasonInvalidMarketID,
	RejectionReasonInvalidOrderID,
	RejectionReasonOrderOutOfSequence,
	RejectionReasonInvalidRemainingSize,
	RejectionReasonTimeFailure,
	RejectionReasonOrderRemovalFailure,
	RejectionReasonInvalidExpirationTime,
	RejectionReasonInvalidOrderReference,
	RejectionReasonEditNotAllowed,
	RejectionReasonOrderAmendFailure,
	RejectionReasonOrderNotFound,
	RejectionReasonInvalidPartyID,
	RejectionReasonMarketClosed,
	RejectionReasonMarginCheckFailed,
	RejectionReasonInternalError,
}

func (e RejectionReason) IsValid() bool {
	switch e {
	case RejectionReasonInvalidMarketID, RejectionReasonInvalidOrderID, RejectionReasonOrderOutOfSequence, RejectionReasonInvalidRemainingSize, RejectionReasonTimeFailure, RejectionReasonOrderRemovalFailure, RejectionReasonInvalidExpirationTime, RejectionReasonInvalidOrderReference, RejectionReasonEditNotAllowed, RejectionReasonOrderAmendFailure, RejectionReasonOrderNotFound, RejectionReasonInvalidPartyID, RejectionReasonMarketClosed, RejectionReasonMarginCheckFailed, RejectionReasonInternalError:
		return true
	}
	return false
}

func (e RejectionReason) String() string {
	return string(e)
}

func (e *RejectionReason) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RejectionReason(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RejectionReason", str)
	}
	return nil
}

func (e RejectionReason) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Whether the placer of an order is aiming to buy or sell on the market
type Side string

const (
	// The Placer of the order is aiming to buy
	SideBuy Side = "Buy"
	// The placer of the order is aiming to sell
	SideSell Side = "Sell"
)

var AllSide = []Side{
	SideBuy,
	SideSell,
}

func (e Side) IsValid() bool {
	switch e {
	case SideBuy, SideSell:
		return true
	}
	return false
}

func (e Side) String() string {
	return string(e)
}

func (e *Side) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Side(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Side", str)
	}
	return nil
}

func (e Side) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type VoteValue string

const (
	// NO reject a proposal
	VoteValueNo VoteValue = "NO"
	// YES accept a proposal
	VoteValueYes VoteValue = "YES"
)

var AllVoteValue = []VoteValue{
	VoteValueNo,
	VoteValueYes,
}

func (e VoteValue) IsValid() bool {
	switch e {
	case VoteValueNo, VoteValueYes:
		return true
	}
	return false
}

func (e VoteValue) String() string {
	return string(e)
}

func (e *VoteValue) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = VoteValue(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid VoteValue", str)
	}
	return nil
}

func (e VoteValue) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
