// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import mock "github.com/stretchr/testify/mock"
import msg "vega/msg"

// CandleStore is an autogenerated mock type for the CandleStore type
type CandleStore struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *CandleStore) Close() {
	_m.Called()
}

// GenerateCandles provides a mock function with given fields: trade
func (_m *CandleStore) GenerateCandles(trade *msg.Trade) error {
	ret := _m.Called(trade)

	var r0 error
	if rf, ok := ret.Get(0).(func(*msg.Trade) error); ok {
		r0 = rf(trade)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GenerateEmptyCandles provides a mock function with given fields: market, timestamp
func (_m *CandleStore) GenerateEmptyCandles(market string, timestamp uint64) error {
	ret := _m.Called(market, timestamp)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, uint64) error); ok {
		r0 = rf(market, timestamp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetCandles provides a mock function with given fields: market, sinceTimestamp, interval
func (_m *CandleStore) GetCandles(market string, sinceTimestamp uint64, interval msg.Interval) []*msg.Candle {
	ret := _m.Called(market, sinceTimestamp, interval)

	var r0 []*msg.Candle
	if rf, ok := ret.Get(0).(func(string, uint64, msg.Interval) []*msg.Candle); ok {
		r0 = rf(market, sinceTimestamp, interval)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*msg.Candle)
		}
	}

	return r0
}

// Notify provides a mock function with given fields:
func (_m *CandleStore) Notify() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// QueueEvent provides a mock function with given fields: candle, interval
func (_m *CandleStore) QueueEvent(candle msg.Candle, interval msg.Interval) error {
	ret := _m.Called(candle, interval)

	var r0 error
	if rf, ok := ret.Get(0).(func(msg.Candle, msg.Interval) error); ok {
		r0 = rf(candle, interval)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Subscribe provides a mock function with given fields: internalTransport
func (_m *CandleStore) Subscribe(internalTransport map[msg.Interval]chan msg.Candle) uint64 {
	ret := _m.Called(internalTransport)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(map[msg.Interval]chan msg.Candle) uint64); ok {
		r0 = rf(internalTransport)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// Unsubscribe provides a mock function with given fields: id
func (_m *CandleStore) Unsubscribe(id uint64) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
