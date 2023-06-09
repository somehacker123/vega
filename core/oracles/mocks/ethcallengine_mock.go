// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/core/oracles (interfaces: EthCallEngine)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	big "math/big"
	reflect "reflect"

	oracles "code.vegaprotocol.io/vega/core/oracles"
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

// CallContract mocks base method.
func (m *MockEthCallEngine) CallContract(arg0 context.Context, arg1 string, arg2 *big.Int) (oracles.EthCallResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CallContract", arg0, arg1, arg2)
	ret0, _ := ret[0].(oracles.EthCallResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CallContract indicates an expected call of CallContract.
func (mr *MockEthCallEngineMockRecorder) CallContract(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CallContract", reflect.TypeOf((*MockEthCallEngine)(nil).CallContract), arg0, arg1, arg2)
}

// MakeResult mocks base method.
func (m *MockEthCallEngine) MakeResult(arg0 string, arg1 []byte) (oracles.EthCallResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeResult", arg0, arg1)
	ret0, _ := ret[0].(oracles.EthCallResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MakeResult indicates an expected call of MakeResult.
func (mr *MockEthCallEngineMockRecorder) MakeResult(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeResult", reflect.TypeOf((*MockEthCallEngine)(nil).MakeResult), arg0, arg1)
}
