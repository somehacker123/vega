// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/core/oracles (interfaces: EthereumConfirmations)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockEthereumConfirmations is a mock of EthereumConfirmations interface.
type MockEthereumConfirmations struct {
	ctrl     *gomock.Controller
	recorder *MockEthereumConfirmationsMockRecorder
}

// MockEthereumConfirmationsMockRecorder is the mock recorder for MockEthereumConfirmations.
type MockEthereumConfirmationsMockRecorder struct {
	mock *MockEthereumConfirmations
}

// NewMockEthereumConfirmations creates a new mock instance.
func NewMockEthereumConfirmations(ctrl *gomock.Controller) *MockEthereumConfirmations {
	mock := &MockEthereumConfirmations{ctrl: ctrl}
	mock.recorder = &MockEthereumConfirmationsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEthereumConfirmations) EXPECT() *MockEthereumConfirmationsMockRecorder {
	return m.recorder
}

// Check mocks base method.
func (m *MockEthereumConfirmations) Check(arg0 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Check", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Check indicates an expected call of Check.
func (mr *MockEthereumConfirmationsMockRecorder) Check(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Check", reflect.TypeOf((*MockEthereumConfirmations)(nil).Check), arg0)
}
