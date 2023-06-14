// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/core/oracles (interfaces: EthCallEngine)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	ethcall "code.vegaprotocol.io/vega/core/evtforward/ethcall"
	gomock "github.com/golang/mock/gomock"
)

// MockEthCallEngine is a mock of EthCallEngine interface.
type MockEthCallEngine struct {
	ctrl     *gomock.Controller
	recorder *MockEthCallEngineMockRecorder
}

// MockEthCallEngineMockRecorder is the mock recorder for MockEthCallEngine.
type MockEthCallEngineMockRecorder struct {
	mock *MockEthCallEngine
}

// NewMockEthCallEngine creates a new mock instance.
func NewMockEthCallEngine(ctrl *gomock.Controller) *MockEthCallEngine {
	mock := &MockEthCallEngine{ctrl: ctrl}
	mock.recorder = &MockEthCallEngineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEthCallEngine) EXPECT() *MockEthCallEngineMockRecorder {
	return m.recorder
}

// CallSpec mocks base method.
func (m *MockEthCallEngine) CallSpec(arg0 context.Context, arg1 string, arg2 uint64) (ethcall.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CallSpec", arg0, arg1, arg2)
	ret0, _ := ret[0].(ethcall.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CallSpec indicates an expected call of CallSpec.
func (mr *MockEthCallEngineMockRecorder) CallSpec(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CallSpec", reflect.TypeOf((*MockEthCallEngine)(nil).CallSpec), arg0, arg1, arg2)
}

// MakeResult mocks base method.
func (m *MockEthCallEngine) MakeResult(arg0 string, arg1 []byte) (ethcall.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeResult", arg0, arg1)
	ret0, _ := ret[0].(ethcall.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MakeResult indicates an expected call of MakeResult.
func (mr *MockEthCallEngineMockRecorder) MakeResult(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeResult", reflect.TypeOf((*MockEthCallEngine)(nil).MakeResult), arg0, arg1)
}
