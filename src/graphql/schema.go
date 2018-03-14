package graphql

import (
	"github.com/graph-gophers/graphql-go"
)

var schema *graphql.Schema

func Schema() *graphql.Schema {
	if schema == nil {
		schema = graphql.MustParseSchema(schemaString, &Resolver{})
	}
	return schema
}

//TODO: investigate a way to build this into the binary from graphql/schema.graphql e.g. go-bindata
const schemaString = `
	# Represents a date/time
	scalar DateTime
	
	
	# For now just a symbol, will almost certainly be it's own type later
	enum Asset {
	  ETH
	  BTC
	  XLM
	}
	
	
	# Secures the purchase of 'contract size' contracts of the 'underlying' at 'maturity'
	type Future {
	
	  # The asset which is purchasable at maturity
	  underlying: Asset!
	
	  # The amount of the underlying to be purchased for at maturity
	  contractSize: Int!
	
	  # The time and date when the future matures (e.g. when the asset is delivered)
	  maturity: DateTime!
	}
	
	
	# Valued at the difference between the initial trading price and the current price of an 'underlying' divided by the 'contract size', a perpetual product with no maturity/settlement
	type ContractForDifference {
	
	  # The asset which is purchasable for the strike price if the option is exercised
	  underlying: Asset!
	
	  # The amount of the underlying that counts as a 1 point move
	  contractSize: Int!
	}
	
	
	# Pays 'contrract size' at 'maturity' if the 'condition' (e.g. Arsenal win a football match or the FTSE 100 closes over 10,000) is deemed to have been met, otherwise pays nothing
	type BinaryOption {
	
	  # The time and date when the condition is evaluated and the product settles
	  maturity: DateTime!
	
	  # The proposition or test to be applied at maturity to determine if it settles at a price of 1 or 0
	  condition: ID!  # TBC
	}
	
	
	
	# Determines the type of a standard option
	enum StandardOptionType {
	
	  # This option type gives the right to buy at the strike price at maturity only
	  EUROPEAN,
	
	  # This option type gives the right to buy at the strike price any time between the trade occurring and maturity
	  AMERICAN
	}
	
	
	# Confers the right but not the obligation to purchase the 'underlying' asset at a given 'strike' price
	type StandardOption {
	
	  # The asset which is purchasable for the strike price if the option is exercised
	  underlying: Asset!
	
	  # The price the underlying asset will be purchased for if the option is exercised
	  strike: Int!
	
	  # The time and date when the option expires
	  maturity: DateTime!
	
	  # The type of standard option in question (currently European or American)
	  type: StandardOptionType!
	}
	
	
	# Union type for all possible products
	union Product = StandardOption | Future | ContractForDifference | BinaryOption
	
	
	# Represents a product & associated parameters that can be traded on Vega, has an associated OrderBook and Trade history
	type Market {
	
	  # Hash of the market's "indentifiable parameters" that uniquely identifies a market without duplicates
	  id: ID!
	
	  # Market full name
	  name: String! @fake(type:productName)
	
	  # Product traded on this market
	  product: Product!
	
	  # Ticker symbol for the market
	  symbol: String! @fake(type:tickerSymbol)
	
	  # Whether the market is active for trading
	  active: Boolean!
	
	  # The asset/currency in which orders and trades are priced and cash settlement occurs
	  baseCurrency: Asset!
	
	  # Timestamp of last data update
	  updatedAt: DateTime! @fake(type:pastDate)
	
	  # Current order book
	  orderBook: OrderBook!
	
	  # All trades that have occurred on this market in order
	  trades: [Trade] @fake(type:trade)
	
	  # The most recent trade on this market
	  lastTrade: Trade @fake(type:trade)
	}
	
	
	# Represents the data in the order book for a market
	type OrderBook {
	
	  # The highest price anyone is currently willing to pay to buy at on this market
	  bestBid: Int @fake(type:money)
	
	  # The lowest price anyone is currently willing to sell at on this market
	  bestOffer: Int @fake(type:money)
	
	  # The half way point between the best bid and best offer
	  midPrice: Float @fake(type:money)
	
	  # All orders on the order book by timestamp when they were added
	  orders: [Order]!
	
	  # All price levels for which there is at least 1 buy order (ordered best/highest price to lowest)
	  buySide: [PriceLevel]!
	
	  # All price levels for which there is at least 1 sell order (ordered best/lowest price to highest)
	  sellSide: [PriceLevel]!
	}
	
	
	# Represents a price on either the buy or sell side and all the orders at that price
	type PriceLevel {
	
	  # Does this PriceLevel contain buy or sell orders?
	  side: Side!
	
	  # The price of all the orders at this level
	  price: Int!
	
	  # The total remaiming size of all orders at this level
	  volume: Int!
	
	  # List of orders sorted by time added from oldest to newest
	  order: [Order]!
	}
	
	
	# Valid order types, these determine what happens when an order is added to the book
	enum OrderType {
	
	  # The order either trades completely (remainingSize == 0 after adding) or not at all, does not remain on the book if it doesn't trade
	  FILL_OR_KILL,
	
	  # The order trades any amount and as much as possible but does not remain on the book (whether it trades or not)
	  EXECUTE_AND_ELIMINATE,
	
	  # This order trades any amount and as much as possible and remains on the book until it either trades completely or is cancelled
	  GOOD_TILL_CANCELLED,
	
	  # This order type trades any amount and as much as possible and remains on the book until they either trade completely, are cancelled, or expires at a set time
	  # NOTE: this may in future be multiple types or have sub types for orders that provide different ways of specifying expiry
	  GOOD_TILL_TIME,
	}
	
	
	# Whether the placer of an order is aiming to buy or sell on the market
	enum Side {
	  BUY
	  SELL
	}
	
	
	# An order in Vega, if active it will be on the OrderBoook for the market
	type Order {
	
	  # Hash of the order data
	  id: ID!
	
	  # The worst price the order will trade at (e.g. buy for price or less, sell for price or more)
	  price: Int! @fake(type:money)
	
	  # The type of order (determines how and if it executes, and whether it persists on the book)
	  type: OrderType!
	
	  # Whether the order is to buy or sell
	  side: Side!
	
	  # The market the order is trading on (probably stored internally as a hash of the market details)
	  market: Market!
	
	  # If the order can expire, specifies the expiry date/time
	  expiry: DateTime
	
	  # Total number of contracts that may be bought or sold (immutable)
	  size: Int!
	
	  # Number of contracts remaining of the total that have not yet been bought or sold
	  remainingSize: Int!
	
	  # Flag that is true if the order has been cancelled by the trader
	  cancelled: Boolean!
	
	  # Flag that is true if the order has been filled (remainingSize == 0)
	  filled: Boolean!
	
	  # Flag that is true while the order remains on the order book
	  active: Boolean!
	
	  # The trader who place the order (probably stored internally as the trader's public key)
	  counterparty: Counterparty!
	
	  # If the order was added to the book or uncrossed at any point, the timestamp when that was done
	  timestamp: DateTime
	}
	
	
	# A trader or market maker on Vega
	type Counterparty {
	
	  # Public key of the counterparty
	  id: ID!
	}
	
	
	# A trade on Vega, the result of two orders being "matched" in the market
	type Trade {
	
	  # The hash of the trade data
	  id: ID!
	
	  # The market the trade occurred on
	  market: Market!
	
	  # The order that bought
	  buyOrder: Order!
	
	  # The order that sold
	  sellOrder: Order!
	
	  # The order that was newly added to the book and caused the trade
	  aggressiveOrder: Order!
	
	  # The order that was already on the book and was "hit" by the incoming order
	  passiveOrder: Order!
	
	  # The price of the trade (probably initially the passive order price, other determination algorithms are possible though)
	  price: Int!
	
	  # The number of contracts trades, will always be <= the remaining size of both orders immediately before the trade
	  size: Int!
	
	  # When the trade occured, probably the timestamp of the agressive order
	  timestamp: DateTime!
	}
	
	type Query {
	  markets(asset: Asset): [Market!]
	}
	
	type Subscription {
	  markets(market: String): [Market]
	}
	
	type Mutation {
	  order(market: ID, Price: Int, type: OrderType, side: Side, expiry: DateTime): Order
	}
	
	schema {
	  query: Query
subscription: Subscription
mutation: Mutation
}`