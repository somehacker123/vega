// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import context "context"
import core "vega/core"
import datastore "vega/datastore"
import mock "github.com/stretchr/testify/mock"
import msg "vega/msg"
import time "time"

// TradeService is an autogenerated mock type for the TradeService type
type TradeService struct {
	mock.Mock
}

// GetByMarket provides a mock function with given fields: ctx, market, limit
func (_m *TradeService) GetByMarket(ctx context.Context, market string, limit uint64) ([]*msg.Trade, error) {
	ret := _m.Called(ctx, market, limit)

	var r0 []*msg.Trade
	if rf, ok := ret.Get(0).(func(context.Context, string, uint64) []*msg.Trade); ok {
		r0 = rf(ctx, market, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*msg.Trade)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, uint64) error); ok {
		r1 = rf(ctx, market, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByMarketAndId provides a mock function with given fields: ctx, market, id
func (_m *TradeService) GetByMarketAndId(ctx context.Context, market string, id string) (*msg.Trade, error) {
	ret := _m.Called(ctx, market, id)

	var r0 *msg.Trade
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *msg.Trade); ok {
		r0 = rf(ctx, market, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*msg.Trade)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, market, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByParty provides a mock function with given fields: ctx, party, limit
func (_m *TradeService) GetByParty(ctx context.Context, party string, limit uint64) ([]*msg.Trade, error) {
	ret := _m.Called(ctx, party, limit)

	var r0 []*msg.Trade
	if rf, ok := ret.Get(0).(func(context.Context, string, uint64) []*msg.Trade); ok {
		r0 = rf(ctx, party, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*msg.Trade)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, uint64) error); ok {
		r1 = rf(ctx, party, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByPartyAndId provides a mock function with given fields: ctx, party, id
func (_m *TradeService) GetByPartyAndId(ctx context.Context, party string, id string) (*msg.Trade, error) {
	ret := _m.Called(ctx, party, id)

	var r0 *msg.Trade
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *msg.Trade); ok {
		r0 = rf(ctx, party, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*msg.Trade)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, party, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCandles provides a mock function with given fields: ctx, market, since, interval
func (_m *TradeService) GetCandles(ctx context.Context, market string, since time.Time, interval uint64) (msg.Candles, error) {
	ret := _m.Called(ctx, market, since, interval)

	var r0 msg.Candles
	if rf, ok := ret.Get(0).(func(context.Context, string, time.Time, uint64) msg.Candles); ok {
		r0 = rf(ctx, market, since, interval)
	} else {
		r0 = ret.Get(0).(msg.Candles)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, time.Time, uint64) error); ok {
		r1 = rf(ctx, market, since, interval)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Init provides a mock function with given fields: app, tradeStore
func (_m *TradeService) Init(app *core.Vega, tradeStore datastore.TradeStore) {
	_m.Called(app, tradeStore)
}
