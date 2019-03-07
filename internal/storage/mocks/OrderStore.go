// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import filtering "code.vegaprotocol.io/vega/internal/filtering"
import mock "github.com/stretchr/testify/mock"
import proto "code.vegaprotocol.io/vega/proto"

// OrderStore is an autogenerated mock type for the OrderStore type
type OrderStore struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *OrderStore) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Commit provides a mock function with given fields:
func (_m *OrderStore) Commit() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByMarket provides a mock function with given fields: market, filters
func (_m *OrderStore) GetByMarket(market string, filters *filtering.OrderQueryFilters) ([]*proto.Order, error) {
	ret := _m.Called(market, filters)

	var r0 []*proto.Order
	if rf, ok := ret.Get(0).(func(string, *filtering.OrderQueryFilters) []*proto.Order); ok {
		r0 = rf(market, filters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*proto.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *filtering.OrderQueryFilters) error); ok {
		r1 = rf(market, filters)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByMarketAndId provides a mock function with given fields: market, id
func (_m *OrderStore) GetByMarketAndId(market string, id string) (*proto.Order, error) {
	ret := _m.Called(market, id)

	var r0 *proto.Order
	if rf, ok := ret.Get(0).(func(string, string) *proto.Order); ok {
		r0 = rf(market, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(market, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByParty provides a mock function with given fields: party, filters
func (_m *OrderStore) GetByParty(party string, filters *filtering.OrderQueryFilters) ([]*proto.Order, error) {
	ret := _m.Called(party, filters)

	var r0 []*proto.Order
	if rf, ok := ret.Get(0).(func(string, *filtering.OrderQueryFilters) []*proto.Order); ok {
		r0 = rf(party, filters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*proto.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *filtering.OrderQueryFilters) error); ok {
		r1 = rf(party, filters)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByPartyAndId provides a mock function with given fields: party, id
func (_m *OrderStore) GetByPartyAndId(party string, id string) (*proto.Order, error) {
	ret := _m.Called(party, id)

	var r0 *proto.Order
	if rf, ok := ret.Get(0).(func(string, string) *proto.Order); ok {
		r0 = rf(party, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(party, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMarketDepth provides a mock function with given fields: market
func (_m *OrderStore) GetMarketDepth(market string) (*proto.MarketDepth, error) {
	ret := _m.Called(market)

	var r0 *proto.MarketDepth
	if rf, ok := ret.Get(0).(func(string) *proto.MarketDepth); ok {
		r0 = rf(market)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.MarketDepth)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(market)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Post provides a mock function with given fields: order
func (_m *OrderStore) Post(order proto.Order) error {
	ret := _m.Called(order)

	var r0 error
	if rf, ok := ret.Get(0).(func(proto.Order) error); ok {
		r0 = rf(order)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Put provides a mock function with given fields: order
func (_m *OrderStore) Put(order proto.Order) error {
	ret := _m.Called(order)

	var r0 error
	if rf, ok := ret.Get(0).(func(proto.Order) error); ok {
		r0 = rf(order)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Subscribe provides a mock function with given fields: orders
func (_m *OrderStore) Subscribe(orders chan<- []proto.Order) uint64 {
	ret := _m.Called(orders)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(chan<- []proto.Order) uint64); ok {
		r0 = rf(orders)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// Unsubscribe provides a mock function with given fields: id
func (_m *OrderStore) Unsubscribe(id uint64) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
