// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/processor (interfaces: Banking)

// Package mocks is a generated GoMock package.
package mocks

import (
	proto "code.vegaprotocol.io/vega/proto"
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockBanking is a mock of Banking interface
type MockBanking struct {
	ctrl     *gomock.Controller
	recorder *MockBankingMockRecorder
}

// MockBankingMockRecorder is the mock recorder for MockBanking
type MockBankingMockRecorder struct {
	mock *MockBanking
}

// NewMockBanking creates a new mock instance
func NewMockBanking(ctrl *gomock.Controller) *MockBanking {
	mock := &MockBanking{ctrl: ctrl}
	mock.recorder = &MockBankingMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBanking) EXPECT() *MockBankingMockRecorder {
	return m.recorder
}

// DepositBuiltinAsset mocks base method
func (m *MockBanking) DepositBuiltinAsset(arg0 context.Context, arg1 *proto.BuiltinAssetDeposit, arg2 string, arg3 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DepositBuiltinAsset", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// DepositBuiltinAsset indicates an expected call of DepositBuiltinAsset
func (mr *MockBankingMockRecorder) DepositBuiltinAsset(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DepositBuiltinAsset", reflect.TypeOf((*MockBanking)(nil).DepositBuiltinAsset), arg0, arg1, arg2, arg3)
}

// DepositERC20 mocks base method
func (m *MockBanking) DepositERC20(arg0 context.Context, arg1 *proto.ERC20Deposit, arg2 string, arg3, arg4 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DepositERC20", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// DepositERC20 indicates an expected call of DepositERC20
func (mr *MockBankingMockRecorder) DepositERC20(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DepositERC20", reflect.TypeOf((*MockBanking)(nil).DepositERC20), arg0, arg1, arg2, arg3, arg4)
}

// EnableBuiltinAsset mocks base method
func (m *MockBanking) EnableBuiltinAsset(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableBuiltinAsset", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnableBuiltinAsset indicates an expected call of EnableBuiltinAsset
func (mr *MockBankingMockRecorder) EnableBuiltinAsset(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableBuiltinAsset", reflect.TypeOf((*MockBanking)(nil).EnableBuiltinAsset), arg0, arg1)
}

// EnableERC20 mocks base method
func (m *MockBanking) EnableERC20(arg0 context.Context, arg1 *proto.ERC20AssetList, arg2, arg3 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableERC20", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnableERC20 indicates an expected call of EnableERC20
func (mr *MockBankingMockRecorder) EnableERC20(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableERC20", reflect.TypeOf((*MockBanking)(nil).EnableERC20), arg0, arg1, arg2, arg3)
}

// LockWithdrawalERC20 mocks base method
func (m *MockBanking) LockWithdrawalERC20(arg0 context.Context, arg1, arg2, arg3 string, arg4 uint64, arg5 *proto.Erc20WithdrawExt) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LockWithdrawalERC20", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// LockWithdrawalERC20 indicates an expected call of LockWithdrawalERC20
func (mr *MockBankingMockRecorder) LockWithdrawalERC20(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LockWithdrawalERC20", reflect.TypeOf((*MockBanking)(nil).LockWithdrawalERC20), arg0, arg1, arg2, arg3, arg4, arg5)
}

// WithdrawalBuiltinAsset mocks base method
func (m *MockBanking) WithdrawalBuiltinAsset(arg0 context.Context, arg1, arg2, arg3 string, arg4 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithdrawalBuiltinAsset", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// WithdrawalBuiltinAsset indicates an expected call of WithdrawalBuiltinAsset
func (mr *MockBankingMockRecorder) WithdrawalBuiltinAsset(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithdrawalBuiltinAsset", reflect.TypeOf((*MockBanking)(nil).WithdrawalBuiltinAsset), arg0, arg1, arg2, arg3, arg4)
}

// WithdrawalERC20 mocks base method
func (m *MockBanking) WithdrawalERC20(arg0 *proto.ERC20Withdrawal, arg1, arg2 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithdrawalERC20", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// WithdrawalERC20 indicates an expected call of WithdrawalERC20
func (mr *MockBankingMockRecorder) WithdrawalERC20(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithdrawalERC20", reflect.TypeOf((*MockBanking)(nil).WithdrawalERC20), arg0, arg1, arg2)
}
