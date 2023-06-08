// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/core/oracles (interfaces: EthCall)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	ethcall "code.vegaprotocol.io/vega/core/evtforward/ethcall"
	gomock "github.com/golang/mock/gomock"
)

// MockEthCall is a mock of EthCall interface.
type MockEthCall struct {
	ctrl     *gomock.Controller
	recorder *MockEthCallMockRecorder
}

// MockEthCallMockRecorder is the mock recorder for MockEthCall.
type MockEthCallMockRecorder struct {
	mock *MockEthCall
}

// NewMockEthCall creates a new mock instance.
func NewMockEthCall(ctrl *gomock.Controller) *MockEthCall {
	mock := &MockEthCall{ctrl: ctrl}
	mock.recorder = &MockEthCallMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEthCall) EXPECT() *MockEthCallMockRecorder {
	return m.recorder
}

// GetDataSource mocks base method.
func (m *MockEthCall) GetDataSource(arg0 string) (ethcall.DataSource, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDataSource", arg0)
	ret0, _ := ret[0].(ethcall.DataSource)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetDataSource indicates an expected call of GetDataSource.
func (mr *MockEthCallMockRecorder) GetDataSource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDataSource", reflect.TypeOf((*MockEthCall)(nil).GetDataSource), arg0)
}
