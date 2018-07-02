package trades

import (
	"context"
	"testing"
	"vega/datastore"
	"vega/datastore/mocks"
	"vega/proto"

	"github.com/stretchr/testify/assert"
)

func TestNewTradeService(t *testing.T) {
	var newTradeService = NewTradeService()
	assert.NotNil(t, newTradeService)
}

func TestGetTradesOnAllMarkets(t *testing.T) {
	var market = "MKT/A"

	var ctx = context.Background()
	var tradeStore = mocks.TradeStore{}
	var tradeService = NewTradeService()
	tradeService.Init(&tradeStore)
	tradeStore.On("GetAll", market, datastore.NewLimitMax()).Return([]*datastore.Trade{
		{Trade: msg.Trade{Id: "A", Market: market, Price: 1}},
		{Trade: msg.Trade{Id: "B", Market: market, Price: 2}},
		{Trade: msg.Trade{Id: "C", Market: market, Price: 3}},
	}, nil).Once()

	var tradeSet, err = tradeService.GetTrades(ctx, market, 12345)

	assert.Nil(t, err)
	assert.NotNil(t, tradeSet)
	assert.Equal(t, 3, len(tradeSet))
	tradeStore.AssertExpectations(t)
}

func TestGetTradesForOrderOnMarket(t *testing.T) {
	var market = "MKT/A"
	var orderId = "Z"

	var ctx = context.Background()
	var tradeStore = mocks.TradeStore{}
	var tradeService = NewTradeService()
	tradeService.Init(&tradeStore)
	tradeStore.On("GetByOrderId", market, orderId, datastore.NewLimitMax()).Return([]*datastore.Trade{
		{Trade: msg.Trade{Id: "A", Market: market, Price: 1}, OrderId: orderId},
		{Trade: msg.Trade{Id: "B", Market: market, Price: 2}, OrderId: orderId},
		{Trade: msg.Trade{Id: "C", Market: market, Price: 3}, OrderId: orderId},
		{Trade: msg.Trade{Id: "D", Market: market, Price: 4}, OrderId: orderId},
		{Trade: msg.Trade{Id: "E", Market: market, Price: 5}, OrderId: orderId},
		{Trade: msg.Trade{Id: "F", Market: market, Price: 6}, OrderId: orderId},
	}, nil).Once()

	var tradeSet, err = tradeService.GetTradesForOrder(ctx, market, orderId, 12345)

	assert.Nil(t, err)
	assert.NotNil(t, tradeSet)
	assert.Equal(t, 6, len(tradeSet))
	tradeStore.AssertExpectations(t)
}
