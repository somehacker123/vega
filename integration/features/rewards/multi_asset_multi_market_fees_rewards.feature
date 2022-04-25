Feature: Fees rewards with multiple markets and assets

Background:
    Given the following network parameters are set:
      | name                                                |  value   |
      | reward.asset                                        |  VEGA    |
      | validators.epoch.length                             |  10s     |
      | validators.delegation.minAmount                     |  10      |
      | reward.staking.delegation.delegatorShare            |  0.883   |
      | reward.staking.delegation.minimumValidatorStake     |  100     |
      | reward.staking.delegation.maxPayoutPerParticipant   | 100000   |
      | reward.staking.delegation.competitionLevel          |  1.1     |
      | reward.staking.delegation.minValidators             |  5       |
      | reward.staking.delegation.optimalStakeMultiplier    |  5.0     |
      | market.value.windowLength                           | 1h       |
      | market.stake.target.timeWindow                      | 24h      |
      | market.stake.target.scalingFactor                   | 1        |
      | market.liquidity.targetstake.triggering.ratio       | 0        |
      | market.liquidity.providers.fee.distributionTimeStep | 0s       |

    Given time is updated to "2021-08-26T00:00:00Z"
    Given the average block duration is "2"

    And the fees configuration named "fees-config-1":
      | maker fee | infrastructure fee |
      | 0.004     | 0.001             |
    And the price monitoring updated every "1000" seconds named "price-monitoring":
      | horizon | probability | auction extension |
      | 1       | 0.99        | 3                 |

    Given the fees configuration named "fees-config-2":
      | maker fee | infrastructure fee |
      | 0.02      | 0.002              |

    When the simple risk model named "simple-risk-model-1":
      | long | short | max move up | min move down | probability of trading |
      | 0.2  | 0.1   | 100          | -100         | 0.1                    |

    And the markets:
      | id        | quote name | asset | risk model          | margin calculator         | auction duration | fees          | price monitoring | oracle config          |
      | ETH/DEC21 | ETH        | ETH   | simple-risk-model-1 | default-margin-calculator | 1                | fees-config-1 | price-monitoring | default-eth-for-future |
      | ETH/DEC22 | ETH        | ETH   | simple-risk-model-1 | default-margin-calculator | 1                | fees-config-1 | price-monitoring | default-eth-for-future |
      | BTC/DEC21 | BTC        | BTC   | simple-risk-model-1 | default-margin-calculator | 1                | fees-config-2 | price-monitoring | default-eth-for-future |
      | BTC/DEC22 | BTC        | BTC   | simple-risk-model-1 | default-margin-calculator | 1                | fees-config-2 | price-monitoring | default-eth-for-future |

    Given the parties deposit on asset's general account the following amount:
    | party           | asset | amount   |
    | a3c024b4e23230c89884a54a813b1ecb4cb0f827a38641c66eeca466da6b2ddf   | VEGA   | 20000000 |
    | a3c024b4e23230c89884a54a813b1ecb4cb0f827a38641c66eeca466da6b2ddf   | USDT   | 20000000 |
    | a3c024b4e23230c89884a54a813b1ecb4cb0f827a38641c66eeca466da6b2ddf   | USDC   | 20000000 |

    Given the parties deposit on asset's general account the following amount:
      | party  | asset | amount     |
      | lp1    | BTC   | 3000000000 |
      | lp2    | BTC   | 3000000000 |
      | lp3    | BTC   | 3000000000 |
      | party1 | BTC   | 300000000  |
      | party2 | BTC   | 300000000  |
      | lp1    | ETH   | 6000000000 |
      | lp2    | ETH   | 6000000000 |
      | lp3    | ETH   | 6000000000 |
      | party1 | ETH   | 600000000  |
      | party2 | ETH   | 600000000  |

Scenario: all sort of fees with multiple assets and multiple markets pay rewards on epoch end
    When the parties submit the following liquidity provision:
      | id  | party | market id | commitment amount | fee   | side | pegged reference | proportion | offset | lp type |
      | lp1 | lp1   | ETH/DEC21 | 4000              | 0.001 | buy  | BID              | 1          | 2      | submission |
      | lp1 | lp1   | ETH/DEC21 | 4000              | 0.001 | buy  | MID              | 2          | 1      | amendment |
      | lp1 | lp1   | ETH/DEC21 | 4000              | 0.001 | sell | ASK              | 1          | 2      | amendment |
      | lp1 | lp1   | ETH/DEC21 | 4000              | 0.001 | sell | MID              | 2          | 1      | amendment |
      | lp1 | lp1   | ETH/DEC22 | 8000              | 0.001 | buy  | BID              | 1          | 2      | submission |
      | lp1 | lp1   | ETH/DEC22 | 8000              | 0.001 | buy  | MID              | 2          | 1      | amendment |
      | lp1 | lp1   | ETH/DEC22 | 8000              | 0.001 | sell | ASK              | 1          | 2      | amendment |
      | lp1 | lp1   | ETH/DEC22 | 8000              | 0.001 | sell | MID              | 2          | 1      | amendment |
      | lp1 | lp1   | BTC/DEC21 | 2000              | 0.001 | buy  | BID              | 1          | 2      | submission |
      | lp1 | lp1   | BTC/DEC21 | 2000              | 0.001 | buy  | MID              | 2          | 1      | amendment |
      | lp1 | lp1   | BTC/DEC21 | 2000              | 0.001 | sell | ASK              | 1          | 2      | amendment |
      | lp1 | lp1   | BTC/DEC21 | 2000              | 0.001 | sell | MID              | 2          | 1      | amendment |
      | lp1 | lp1   | BTC/DEC22 | 4000              | 0.001 | buy  | BID              | 1          | 2      | submission |
      | lp1 | lp1   | BTC/DEC22 | 4000              | 0.001 | buy  | MID              | 2          | 1      | amendment |
      | lp1 | lp1   | BTC/DEC22 | 4000              | 0.001 | sell | ASK              | 1          | 2      | amendment |
      | lp1 | lp1   | BTC/DEC22 | 4000              | 0.001 | sell | MID              | 2          | 1      | amendment |
      | lp2 | lp2   | ETH/DEC21 | 1000              | 0.002 | buy  | BID              | 1          | 2      | submission |
      | lp2 | lp2   | ETH/DEC21 | 1000              | 0.002 | buy  | MID              | 2          | 1      | amendment |
      | lp2 | lp2   | ETH/DEC21 | 1000              | 0.002 | sell | ASK              | 1          | 2      | amendment |
      | lp2 | lp2   | ETH/DEC21 | 1000              | 0.002 | sell | MID              | 2          | 1      | amendment |
      | lp2 | lp2   | ETH/DEC22 | 2000              | 0.002 | buy  | BID              | 1          | 2      | submission |
      | lp2 | lp2   | ETH/DEC22 | 2000              | 0.002 | buy  | MID              | 2          | 1      | amendment |
      | lp2 | lp2   | ETH/DEC22 | 2000              | 0.002 | sell | ASK              | 1          | 2      | amendment |
      | lp2 | lp2   | ETH/DEC22 | 2000              | 0.002 | sell | MID              | 2          | 1      | amendment |
      | lp2 | lp2   | BTC/DEC21 | 500               | 0.002 | buy  | BID              | 1          | 2      | submission |
      | lp2 | lp2   | BTC/DEC21 | 500               | 0.002 | buy  | MID              | 2          | 1      | amendment |
      | lp2 | lp2   | BTC/DEC21 | 500               | 0.002 | sell | ASK              | 1          | 2      | amendment |
      | lp2 | lp2   | BTC/DEC21 | 500               | 0.002 | sell | MID              | 2          | 1      | amendment |
      | lp2 | lp2   | BTC/DEC22 | 1000              | 0.002 | buy  | BID              | 1          | 2      | submission |
      | lp2 | lp2   | BTC/DEC22 | 1000              | 0.002 | buy  | MID              | 2          | 1      | amendment |
      | lp2 | lp2   | BTC/DEC22 | 1000              | 0.002 | sell | ASK              | 1          | 2      | amendment |
      | lp2 | lp2   | BTC/DEC22 | 1000              | 0.002 | sell | MID              | 2          | 1      | amendment |

    Then the parties place the following orders:
      | party  | market id | side | volume | price | resulting trades | type       | tif     |
      | party1 | ETH/DEC21 | buy  | 1      | 900   | 0                | TYPE_LIMIT | TIF_GTC |
      | party1 | ETH/DEC21 | buy  | 60     | 1000  | 0                | TYPE_LIMIT | TIF_GTC |
      | party2 | ETH/DEC21 | sell | 1      | 1100  | 0                | TYPE_LIMIT | TIF_GTC |
      | party2 | ETH/DEC21 | sell | 60     | 1000  | 0                | TYPE_LIMIT | TIF_GTC |
      | party1 | ETH/DEC22 | buy  | 2      | 950   | 0                | TYPE_LIMIT | TIF_GTC |
      | party1 | ETH/DEC22 | buy  | 30     | 1050  | 0                | TYPE_LIMIT | TIF_GTC |
      | party2 | ETH/DEC22 | sell | 2      | 1150  | 0                | TYPE_LIMIT | TIF_GTC |
      | party2 | ETH/DEC22 | sell | 30     | 1050  | 0                | TYPE_LIMIT | TIF_GTC |
      | party1 | BTC/DEC21 | buy  | 3      | 800   | 0                | TYPE_LIMIT | TIF_GTC |
      | party1 | BTC/DEC21 | buy  | 30     | 850   | 0                | TYPE_LIMIT | TIF_GTC |
      | party2 | BTC/DEC21 | sell | 3      | 1100  | 0                | TYPE_LIMIT | TIF_GTC |
      | party2 | BTC/DEC21 | sell | 30     | 850   | 0                | TYPE_LIMIT | TIF_GTC |
      | party1 | BTC/DEC22 | buy  | 4      | 950   | 0                | TYPE_LIMIT | TIF_GTC |
      | party1 | BTC/DEC22 | buy  | 25     | 1030  | 0                | TYPE_LIMIT | TIF_GTC |
      | party2 | BTC/DEC22 | sell | 4      | 1150  | 0                | TYPE_LIMIT | TIF_GTC |
      | party2 | BTC/DEC22 | sell | 25     | 1030  | 0                | TYPE_LIMIT | TIF_GTC |

    When the opening auction period ends for market "ETH/DEC21"
    Then the following trades should be executed:
      | buyer   | price | size | seller  |
      | party1  | 1000  | 60   |  party2 |

    When the opening auction period ends for market "ETH/DEC22"
    Then the following trades should be executed:
      | buyer   | price | size | seller  |
      | party1  | 1050  | 30   |  party2 |

    When the opening auction period ends for market "BTC/DEC21"
    Then the following trades should be executed:
      | buyer   | price | size | seller  |
      | party1  | 850  | 30   |  party2 |

    When the opening auction period ends for market "BTC/DEC22"
    Then the following trades should be executed:
      | buyer   | price | size | seller  |
      | party1  | 1030  | 25   |  party2 |

    And the trading mode should be "TRADING_MODE_CONTINUOUS" for the market "ETH/DEC21"
    And the trading mode should be "TRADING_MODE_CONTINUOUS" for the market "ETH/DEC22"
    And the trading mode should be "TRADING_MODE_CONTINUOUS" for the market "BTC/DEC21"
    And the trading mode should be "TRADING_MODE_CONTINUOUS" for the market "BTC/DEC22"

    Then the parties place the following orders:
    | party  | market id | side | volume | price | resulting trades | type       | tif     | reference    |
    | party1 | ETH/DEC21 | sell | 20     | 1000  | 0                | TYPE_LIMIT | TIF_GTC | party1-sell1 |
    | party2 | ETH/DEC21 | buy  | 20     | 1000  | 9                | TYPE_LIMIT | TIF_GTC | party2-buy1  |
    | party1 | ETH/DEC22 | sell | 30     | 1050  | 0                | TYPE_LIMIT | TIF_GTC | party1-sell2 |
    | party2 | ETH/DEC22 | buy  | 30     | 1050  | 1                | TYPE_LIMIT | TIF_GTC | party2-buy2  |
    | party2 | BTC/DEC21 | sell | 5      | 850   | 0                | TYPE_LIMIT | TIF_GTC | party2-sell1 |
    | party1 | BTC/DEC21 | buy  | 10     | 850   | 1                | TYPE_LIMIT | TIF_GTC | party1-buy1  |
    | party2 | BTC/DEC22 | buy  | 5      | 1030  | 0                | TYPE_LIMIT | TIF_GTC | party2-buy3  |
    | party1 | BTC/DEC22 | sell | 20     | 1030  | 1                | TYPE_LIMIT | TIF_GTC | party1-sell3 |

    And the following trades should be executed:
      | buyer  | price | size | seller |
      | party2 | 951   | 2    | lp1    |
      | party2 | 951   | 2    | lp1    |
      | party2 | 951   | 2    | lp1    |
      | party2 | 951   | 2    | lp1    |
      | party2 | 951   | 1    | lp2    |
      | party2 | 951   | 1    | lp2    |
      | party2 | 951   | 1    | lp2    |
      | party2 | 951   | 1    | lp2    |
      | party2 | 1000  | 8    | party1 |
      | party2 | 1050  | 30   | party1 |
      | party1 | 850   | 5    | party2 |
      | party2 | 1030  | 5    | party1 |

    Given the parties submit the following one off transfers:
    | id  |                             from                                 |  from_account_type    |                                to                                 |   to_account_type                       | asset  |   market   |amount |       delivery_time   |
    | 1   | a3c024b4e23230c89884a54a813b1ecb4cb0f827a38641c66eeca466da6b2ddf |  ACCOUNT_TYPE_GENERAL |  0000000000000000000000000000000000000000000000000000000000000000 | ACCOUNT_TYPE_REWARD_MAKER_RECEIVED_FEES | VEGA   |  ETH/DEC21 | 10000 | 2021-08-26T00:00:00Z  |
    | 2   | a3c024b4e23230c89884a54a813b1ecb4cb0f827a38641c66eeca466da6b2ddf |  ACCOUNT_TYPE_GENERAL |  0000000000000000000000000000000000000000000000000000000000000000 | ACCOUNT_TYPE_REWARD_TAKER_PAID_FEES     | USDT   |  ETH/DEC21 | 20000 | 2021-08-26T00:00:10Z  |
    | 3   | a3c024b4e23230c89884a54a813b1ecb4cb0f827a38641c66eeca466da6b2ddf |  ACCOUNT_TYPE_GENERAL |  0000000000000000000000000000000000000000000000000000000000000000 | ACCOUNT_TYPE_REWARD_LP_RECEIVED_FEES    | USDC   |  ETH/DEC21 | 5000  | 2021-08-26T00:00:10Z  |
    | 4   | a3c024b4e23230c89884a54a813b1ecb4cb0f827a38641c66eeca466da6b2ddf |  ACCOUNT_TYPE_GENERAL |  0000000000000000000000000000000000000000000000000000000000000000 | ACCOUNT_TYPE_REWARD_MAKER_RECEIVED_FEES | VEGA   |  ETH/DEC22 | 10000 | 2021-08-26T00:00:00Z  |
    | 5   | a3c024b4e23230c89884a54a813b1ecb4cb0f827a38641c66eeca466da6b2ddf |  ACCOUNT_TYPE_GENERAL |  0000000000000000000000000000000000000000000000000000000000000000 | ACCOUNT_TYPE_REWARD_TAKER_PAID_FEES     | USDT   |  ETH/DEC22 | 20000 | 2021-08-26T00:00:10Z  |
    | 6   | a3c024b4e23230c89884a54a813b1ecb4cb0f827a38641c66eeca466da6b2ddf |  ACCOUNT_TYPE_GENERAL |  0000000000000000000000000000000000000000000000000000000000000000 | ACCOUNT_TYPE_REWARD_LP_RECEIVED_FEES    | USDC   |  ETH/DEC22 | 5000  | 2021-08-26T00:00:10Z  |
    | 7   | a3c024b4e23230c89884a54a813b1ecb4cb0f827a38641c66eeca466da6b2ddf |  ACCOUNT_TYPE_GENERAL |  0000000000000000000000000000000000000000000000000000000000000000 | ACCOUNT_TYPE_REWARD_MAKER_RECEIVED_FEES | VEGA   |  BTC/DEC21 | 1000  | 2021-08-26T00:00:00Z  |
    | 8   | a3c024b4e23230c89884a54a813b1ecb4cb0f827a38641c66eeca466da6b2ddf |  ACCOUNT_TYPE_GENERAL |  0000000000000000000000000000000000000000000000000000000000000000 | ACCOUNT_TYPE_REWARD_TAKER_PAID_FEES     | USDT   |  BTC/DEC21 | 2000  | 2021-08-26T00:00:10Z  |
    | 9   | a3c024b4e23230c89884a54a813b1ecb4cb0f827a38641c66eeca466da6b2ddf |  ACCOUNT_TYPE_GENERAL |  0000000000000000000000000000000000000000000000000000000000000000 | ACCOUNT_TYPE_REWARD_LP_RECEIVED_FEES    | USDC   |  BTC/DEC21 | 500   | 2021-08-26T00:00:10Z  |
    | 10  | a3c024b4e23230c89884a54a813b1ecb4cb0f827a38641c66eeca466da6b2ddf |  ACCOUNT_TYPE_GENERAL |  0000000000000000000000000000000000000000000000000000000000000000 | ACCOUNT_TYPE_REWARD_MAKER_RECEIVED_FEES | VEGA   |  BTC/DEC22 | 1000  | 2021-08-26T00:00:00Z  |
    | 11  | a3c024b4e23230c89884a54a813b1ecb4cb0f827a38641c66eeca466da6b2ddf |  ACCOUNT_TYPE_GENERAL |  0000000000000000000000000000000000000000000000000000000000000000 | ACCOUNT_TYPE_REWARD_TAKER_PAID_FEES     | USDT   |  BTC/DEC22 | 2000  | 2021-08-26T00:00:10Z  |
    | 12  | a3c024b4e23230c89884a54a813b1ecb4cb0f827a38641c66eeca466da6b2ddf |  ACCOUNT_TYPE_GENERAL |  0000000000000000000000000000000000000000000000000000000000000000 | ACCOUNT_TYPE_REWARD_LP_RECEIVED_FEES    | USDC   |  BTC/DEC22 | 500   | 2021-08-26T00:00:10Z  |

    Then "party1" should have general account balance of "599982442" for asset "ETH"
    Then "party2" should have general account balance of "599995070" for asset "ETH"
    Then "lp1" should have general account balance of "5999990005" for asset "ETH"
    Then "lp2" should have general account balance of "5999996665" for asset "ETH"

    Then "party1" should have general account balance of "299981368" for asset "BTC"
    Then "party2" should have general account balance of "299982246" for asset "BTC"
    Then "lp1" should have general account balance of "3000000000" for asset "BTC"
    Then "lp2" should have general account balance of "3000000000" for asset "BTC"

    #complete the epoch for rewards to take place
    Then the network moves ahead "7" blocks

    # calculation of maker fees received reward - given in VEGA
    # ETH/DEC21 maker fees received:
    # party1 - 0.4 * 10000 = 4000
    # lp1 - 0.4 * 10000 = 4000
    # lp2 - 0.2 * 10000 = 2000

    # ETH/DEC22 maker fees received:
    # party1 - 1 * 10000 = 10000
    
    # BTC/DEC21 maker fees received:
    # party2 - 1 * 1000 = 1000
    
    # BTC/DEC22 maker fees received:
    # party2 - 1 * 1000 = 1000
    
    Then "party1" should have general account balance of "599982442" for asset "ETH"
    Then "party1" should have general account balance of "14000" for asset "VEGA"
    Then "party2" should have general account balance of "2000" for asset "VEGA"
    Then "lp1" should have general account balance of "4000" for asset "VEGA"
    Then "lp2" should have general account balance of "2000" for asset "VEGA"

    # calculation of taker fees paid reward - given in USDT
    # ETH/DEC21 taker fees paid:
    # party2 - 1 * 20000 = 20000
    
    # ETH/DEC22 taker fees paid:
    # party2 - 1 * 20000 = 20000
    
    # BTC/DEC21 taker fees paid:
    # party1 - 1 * 2000 = 2000
    
    # BTC/DEC22 taker fees paid:
    # party1 - 1 * 2000 = 2000

    Then "party1" should have general account balance of "4000" for asset "USDT"
    Then "party2" should have general account balance of "40000" for asset "USDT"
    
    # calculation of taker lp fees received - paid in USDC
    # ETH/DEC21 taker fees paid:
    # lp1 - 0.8 * 5000 = 4000
    # lp2 - 0.2 * 5000 = 1000
    
    Then "lp1" should have general account balance of "4000" for asset "USDC"
    Then "lp2" should have general account balance of "1000" for asset "USDC"
    
   