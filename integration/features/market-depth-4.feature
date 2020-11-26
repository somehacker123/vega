Feature: Test market depth events for pegged orders (cancelling pegged orders)

  Background:
    Given the insurance pool initial balance for the markets is "0":
 #   And the markets starts on "2019-11-30T00:00:00Z" and expires on "2019-12-31T23:59:59Z"
    And the executon engine have these markets:
      | name      | baseName | quoteName | asset | markprice | risk model | lamd/long | tau/short | mu |     r | sigma | release factor | initial factor | search factor | settlementPrice | openAuction | trading mode | makerFee | infrastructureFee | liquidityFee | p. m. update freq. | p. m. horizons | p. m. probs | p. m. durations |
      | ETH/DEC19 | ETH      | BTC       | BTC   |       100 | simple     |          0 |        0 |  0 | 0.016 |   2.0 |            1.4 |            1.2 |           1.1 |              42 | 0           | continuous   |        0 |                 0 |            0 |                 0  |                |             |                 |

  Scenario: Check order events with larger pegged orders, and lower balance
# setup accounts
    Given the following traders:
      | name             |    amount |
      | sellSideProvider | 100000000 |
      | buySideProvider  | 100000000 |
      | pegged1          |     50000 |
      | pegged2          |     50000 |
      | pegged3          |     50000 |
      | pegged4          |     50000 |
    Then I Expect the traders to have new general account:
      | name             | asset |
      | pegged1          | BTC   |
      | pegged2          | BTC   |
      | pegged3          | BTC   |
      | pegged4          | BTC   |
      | sellSideProvider | BTC   |
      | buySideProvider  | BTC   |
# setup pegged orders
    Then traders place pegged orders:
      | trader   | id        | side | volume | reference | offset | price |
      | pegged1  | ETH/DEC19 | sell |    500 | ASK       | 10     | 100   |
      | pegged2  | ETH/DEC19 | sell |    500 | MID       | 15     | 100   |
      | pegged3  | ETH/DEC19 | buy  |    500 | BID       | -10    | 100   |
      | pegged4  | ETH/DEC19 | buy  |    500 | MID       | -10    | 100   |
#   And dump orders
    Then I see the following order events:
      | trader   | id        | side | volume | reference | offset | price | status        |
      | pegged1  | ETH/DEC19 | sell |    500 | ASK       | 10     | 100   | STATUS_PARKED |
      | pegged2  | ETH/DEC19 | sell |    500 | MID       | 15     | 100   | STATUS_PARKED |
      | pegged3  | ETH/DEC19 | buy  |    500 | BID       | -10    | 100   | STATUS_PARKED |
      | pegged4  | ETH/DEC19 | buy  |    500 | MID       | -10    | 100   | STATUS_PARKED |
# keep things simple: remove the events we've just verified
    And clear order events
# setup orderbook
    Then traders place following orders with references:
      | trader           | id        | type | volume | price | resulting trades | type       | tif     | reference       |
      | sellSideProvider | ETH/DEC19 | sell |   1000 |   120 |                0 | TYPE_LIMIT | TIF_GTC | sell-provider-1 |
      | buySideProvider  | ETH/DEC19 | buy  |   1000 |    80 |                0 | TYPE_LIMIT | TIF_GTC | buy-provider-1  |
#   And dump orders
    Then I see the following order events:
      | trader            | id        | side | volume | reference | offset | price | status        |
      | sellSideProvider  | ETH/DEC19 | sell |   1000 |           | 0      | 120   | STATUS_ACTIVE |
      | buySideProvider   | ETH/DEC19 | buy  |   1000 |           | 0      | 80    | STATUS_ACTIVE |
# Checked out, remove the order events we've checked, now let's have a look at the pegged order events
    And clear order events by reference:
      | trader            | reference       |
      | sellSideProvider  | sell-provider-1 |
      | buySideProvider   | buy-provider-1  |
#   And dump orders
# Now check what happened to our pegged orders
    Then I see the following order events:
      | trader   | id        | side | volume | reference | offset | price | status          |
      | pegged1  | ETH/DEC19 | sell |    500 | ASK       | 10     | 130   | STATUS_ACTIVE   |
      | pegged1  | ETH/DEC19 | sell |    500 | ASK       | 10     | 130   | STATUS_ACTIVE   |
      | pegged1  | ETH/DEC19 | sell |    500 | ASK       | 10     | 130   | STATUS_ACTIVE   |
      | pegged2  | ETH/DEC19 | sell |    500 | MID       | 15     | 115   | STATUS_ACTIVE   |
      | pegged3  | ETH/DEC19 | buy  |    500 | BID       | -10    | 70    | STATUS_ACTIVE   |
      | pegged4  | ETH/DEC19 | buy  |    500 | MID       | -10    | 90    | STATUS_ACTIVE   |
##  Cancel some pegged events, and clear order event buffer so we can ignore the events we checked above 
    Then traders cancel pegged orders and clear:
      | trader  | MarketID  |
      | pegged1 | ETH/DEC19 |
      | pegged3 | ETH/DEC19 |
      | pegged2 | ETH/DEC19 |
#   And dump orders
    Then I see the following order events:
      | trader   | id        | side | volume | reference | offset | price | status           |
      | pegged3  | ETH/DEC19 | buy  |    500 | BID       | -10    | 70    | STATUS_CANCELLED |
      | pegged1  | ETH/DEC19 | sell |    500 | ASK       | 10     | 130   | STATUS_CANCELLED |
      | pegged2  | ETH/DEC19 | sell |    500 | MID       | 15     | 115   | STATUS_CANCELLED |
#   And dump orders