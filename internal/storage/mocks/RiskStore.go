// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import mock "github.com/stretchr/testify/mock"
import proto "vega/proto"

// RiskStore is an autogenerated mock type for the RiskStore type
type RiskStore struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *RiskStore) Close() error {
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
func (_m *RiskStore) Commit() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByMarket provides a mock function with given fields: market
func (_m *RiskStore) GetByMarket(market string) (*proto.RiskFactor, error) {
	ret := _m.Called(market)

	var r0 *proto.RiskFactor
	if rf, ok := ret.Get(0).(func(string) *proto.RiskFactor); ok {
		r0 = rf(market)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.RiskFactor)
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

// Post provides a mock function with given fields: risk
func (_m *RiskStore) Post(risk *proto.RiskFactor) error {
	ret := _m.Called(risk)

	var r0 error
	if rf, ok := ret.Get(0).(func(*proto.RiskFactor) error); ok {
		r0 = rf(risk)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Put provides a mock function with given fields: risk
func (_m *RiskStore) Put(risk *proto.RiskFactor) error {
	ret := _m.Called(risk)

	var r0 error
	if rf, ok := ret.Get(0).(func(*proto.RiskFactor) error); ok {
		r0 = rf(risk)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}