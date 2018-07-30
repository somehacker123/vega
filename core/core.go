package core

import (
	"fmt"
	"time"

	"vega/datastore"
	"vega/matching"
	"vega/msg"
	"vega/risk"
	"vega/log"
)

const marketName = "BTC/DEC18"

type Config struct{
	GenesisTime time.Time
}

type Vega struct {
	config         *Config
	markets        map[string]*matching.OrderBook
	OrderStore     datastore.OrderStore
	TradeStore     datastore.TradeStore
	matchingEngine matching.MatchingEngine
	State          *State
	riskEngine     risk.RiskEngine
}

func New(config *Config, store *datastore.MemoryStoreProvider) *Vega {

	// Initialise matching engine
	matchingEngine := matching.NewMatchingEngine()

	// Initialise risk engine
	riskEngine := risk.New()

	return &Vega{
		config:         config,
		markets:        make(map[string]*matching.OrderBook),
		OrderStore:     store.OrderStore(),
		TradeStore:     store.TradeStore(),
		matchingEngine: matchingEngine,
		riskEngine:     riskEngine,
		State:          newState(),
	}
}

func (v *Vega) SetGenesisTime(genesisTime time.Time) {
	v.config.GenesisTime = genesisTime
}

func GetConfig() *Config {
	return &Config{}
}

func (v *Vega) GetAbciHeight() int64 {
	return v.State.Height
}

func (v *Vega) GetTime() time.Time {
	//genesisTime, _ := time.Parse(time.RFC3339, genesisTimeStr)
	genesisTime := v.config.GenesisTime
	return genesisTime.Add(time.Duration(v.State.Height) * time.Second)
}

func (v *Vega) InitialiseMarkets() {
	v.matchingEngine.CreateMarket(marketName)
}

func (v *Vega) SubmitOrder(order *msg.Order) (*msg.OrderConfirmation, msg.OrderError) {

	order.Id = fmt.Sprintf("V%d-%d", v.State.Height, v.State.Size)
	order.Timestamp = uint64(v.State.Height)

	// -----------------------------------------------//
	//----------------- MATCHING ENGINE --------------//
	// 1) submit order to matching engine
	confirmation, errorMsg := v.matchingEngine.SubmitOrder(order)
	if confirmation == nil || errorMsg != msg.OrderError_NONE {
		return nil, errorMsg
	}

	// ------------------------------------------------//
	// 2) --------------- RISK ENGINE -----------------//

	log.Infof("Risk BEFORE calling model calculation = ", order.RiskFactor)

	v.riskEngine.Assess(order)
	confirmation.Order = order

	log.Infof("Risk AFTER calling model calculation = ", order.RiskFactor)

	// -----------------------------------------------//
	//-------------------- STORES --------------------//
	// 3) save to stores

	// insert aggressive remaining order
	err := v.OrderStore.Post(*datastore.NewOrderFromProtoMessage(order))
	if err != nil {
		// Note: writing to store should not prevent flow to other engines
		log.Infof("OrderStore.Post error: %+v\n", err)
	}

	if confirmation.PassiveOrdersAffected != nil {
		// insert all passive orders siting on the book
		for _, order := range confirmation.PassiveOrdersAffected {
			// Note: writing to store should not prevent flow to other engines
			err := v.OrderStore.Put(*datastore.NewOrderFromProtoMessage(order))
			if err != nil {
				log.Infof("OrderStore.Put error: %+v\n", err)
			}
		}
	}

	if confirmation.Trades != nil {
		// insert all trades resulted from the executed order
		for idx, trade := range confirmation.Trades {
			trade.Id = fmt.Sprintf("%s-%d", order.Id, idx)

			t := datastore.NewTradeFromProtoMessage(trade, order.Id, confirmation.PassiveOrdersAffected[idx].Id)
			if err := v.TradeStore.Post(*t); err != nil {
				// Note: writing to store should not prevent flow to other engines
				log.Infof("TradeStore.Post error: %+v\n", err)
			}
		}
	}


	// ------------------------------------------------//

	return confirmation, msg.OrderError_NONE
}

func (v *Vega) CancelOrder(order *msg.Order) (*msg.OrderCancellation, msg.OrderError) {

	log.Infof("CancelOrder: %+v", order)

	// -----------------------------------------------//
	//----------------- MATCHING ENGINE --------------//
	// 1) cancel order in matching engine
	cancellation, errorMsg := v.matchingEngine.CancelOrder(order)
	if cancellation == nil || errorMsg != msg.OrderError_NONE {
		return nil, errorMsg
	}

	// -----------------------------------------------//
	//-------------------- STORES --------------------//
	// 2) if OK update stores

	// insert aggressive remaining order
	err := v.OrderStore.Put(*datastore.NewOrderFromProtoMessage(order))
	if err != nil {
		// Note: writing to store should not prevent flow to other
		log.Infof("OrderStore.Put error: %+v\n", err)
	}

	return cancellation, msg.OrderError_NONE
}
