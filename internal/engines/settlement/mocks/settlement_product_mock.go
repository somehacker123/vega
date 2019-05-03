// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/internal/engines/settlement (interfaces: Product)

// Package mocks is a generated GoMock package.
package mocks

import (
	proto "code.vegaprotocol.io/vega/proto"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockProduct is a mock of Product interface
type MockProduct struct {
	ctrl     *gomock.Controller
	recorder *MockProductMockRecorder
}

// MockProductMockRecorder is the mock recorder for MockProduct
type MockProductMockRecorder struct {
	mock *MockProduct
}

// NewMockProduct creates a new mock instance
func NewMockProduct(ctrl *gomock.Controller) *MockProduct {
	mock := &MockProduct{ctrl: ctrl}
	mock.recorder = &MockProductMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProduct) EXPECT() *MockProductMockRecorder {
	return m.recorder
}

// Settle mocks base method
func (m *MockProduct) Settle(arg0 uint64, arg1 int64) (*proto.FinancialAmount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Settle", arg0, arg1)
	ret0, _ := ret[0].(*proto.FinancialAmount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Settle indicates an expected call of Settle
func (mr *MockProductMockRecorder) Settle(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Settle", reflect.TypeOf((*MockProduct)(nil).Settle), arg0, arg1)
}