// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/internal/trades (interfaces: RiskStore)

// Package mocks is a generated GoMock package.
package mocks

import (
	proto "code.vegaprotocol.io/vega/proto"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRiskStore is a mock of RiskStore interface
type MockRiskStore struct {
	ctrl     *gomock.Controller
	recorder *MockRiskStoreMockRecorder
}

// MockRiskStoreMockRecorder is the mock recorder for MockRiskStore
type MockRiskStoreMockRecorder struct {
	mock *MockRiskStore
}

// NewMockRiskStore creates a new mock instance
func NewMockRiskStore(ctrl *gomock.Controller) *MockRiskStore {
	mock := &MockRiskStore{ctrl: ctrl}
	mock.recorder = &MockRiskStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRiskStore) EXPECT() *MockRiskStoreMockRecorder {
	return m.recorder
}

// GetByMarket mocks base method
func (m *MockRiskStore) GetByMarket(arg0 string) (*proto.RiskFactor, error) {
	ret := m.ctrl.Call(m, "GetByMarket", arg0)
	ret0, _ := ret[0].(*proto.RiskFactor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByMarket indicates an expected call of GetByMarket
func (mr *MockRiskStoreMockRecorder) GetByMarket(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByMarket", reflect.TypeOf((*MockRiskStore)(nil).GetByMarket), arg0)
}