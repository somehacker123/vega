// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/core/oracles (interfaces: ContractCaller)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	big "math/big"
	reflect "reflect"

	ethereum "github.com/ethereum/go-ethereum"
	gomock "github.com/golang/mock/gomock"
)

// MockContractCaller is a mock of ContractCaller interface.
type MockContractCaller struct {
	ctrl     *gomock.Controller
	recorder *MockContractCallerMockRecorder
}

// MockContractCallerMockRecorder is the mock recorder for MockContractCaller.
type MockContractCallerMockRecorder struct {
	mock *MockContractCaller
}

// NewMockContractCaller creates a new mock instance.
func NewMockContractCaller(ctrl *gomock.Controller) *MockContractCaller {
	mock := &MockContractCaller{ctrl: ctrl}
	mock.recorder = &MockContractCallerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContractCaller) EXPECT() *MockContractCallerMockRecorder {
	return m.recorder
}

// CallContract mocks base method.
func (m *MockContractCaller) CallContract(arg0 context.Context, arg1 ethereum.CallMsg, arg2 *big.Int) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CallContract", arg0, arg1, arg2)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CallContract indicates an expected call of CallContract.
func (mr *MockContractCallerMockRecorder) CallContract(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CallContract", reflect.TypeOf((*MockContractCaller)(nil).CallContract), arg0, arg1, arg2)
}
