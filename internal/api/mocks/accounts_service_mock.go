// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/internal/api (interfaces: AccountsService)

// Package mocks is a generated GoMock package.
package mocks

import (
	proto "code.vegaprotocol.io/vega/proto"
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAccountsService is a mock of AccountsService interface
type MockAccountsService struct {
	ctrl     *gomock.Controller
	recorder *MockAccountsServiceMockRecorder
}

// MockAccountsServiceMockRecorder is the mock recorder for MockAccountsService
type MockAccountsServiceMockRecorder struct {
	mock *MockAccountsService
}

// NewMockAccountsService creates a new mock instance
func NewMockAccountsService(ctrl *gomock.Controller) *MockAccountsService {
	mock := &MockAccountsService{ctrl: ctrl}
	mock.recorder = &MockAccountsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAccountsService) EXPECT() *MockAccountsServiceMockRecorder {
	return m.recorder
}

// GetAccountSubscribersCount mocks base method
func (m *MockAccountsService) GetAccountSubscribersCount() int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountSubscribersCount")
	ret0, _ := ret[0].(int32)
	return ret0
}

// GetAccountSubscribersCount indicates an expected call of GetAccountSubscribersCount
func (mr *MockAccountsServiceMockRecorder) GetAccountSubscribersCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountSubscribersCount", reflect.TypeOf((*MockAccountsService)(nil).GetAccountSubscribersCount))
}

// GetTraderAccounts mocks base method
func (m *MockAccountsService) GetTraderAccounts(arg0 string) ([]*proto.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTraderAccounts", arg0)
	ret0, _ := ret[0].([]*proto.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTraderAccounts indicates an expected call of GetTraderAccounts
func (mr *MockAccountsServiceMockRecorder) GetTraderAccounts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTraderAccounts", reflect.TypeOf((*MockAccountsService)(nil).GetTraderAccounts), arg0)
}

// GetTraderAccountsForMarket mocks base method
func (m *MockAccountsService) GetTraderAccountsForMarket(arg0, arg1 string) ([]*proto.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTraderAccountsForMarket", arg0, arg1)
	ret0, _ := ret[0].([]*proto.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTraderAccountsForMarket indicates an expected call of GetTraderAccountsForMarket
func (mr *MockAccountsServiceMockRecorder) GetTraderAccountsForMarket(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTraderAccountsForMarket", reflect.TypeOf((*MockAccountsService)(nil).GetTraderAccountsForMarket), arg0, arg1)
}

// GetTraderMarketBalance mocks base method
func (m *MockAccountsService) GetTraderMarketBalance(arg0, arg1 string) ([]*proto.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTraderMarketBalance", arg0, arg1)
	ret0, _ := ret[0].([]*proto.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTraderMarketBalance indicates an expected call of GetTraderMarketBalance
func (mr *MockAccountsServiceMockRecorder) GetTraderMarketBalance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTraderMarketBalance", reflect.TypeOf((*MockAccountsService)(nil).GetTraderMarketBalance), arg0, arg1)
}

// ObserveAccounts mocks base method
func (m *MockAccountsService) ObserveAccounts(arg0 context.Context, arg1 int, arg2, arg3 string, arg4 proto.AccountType) (<-chan *proto.Account, uint64) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ObserveAccounts", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(<-chan *proto.Account)
	ret1, _ := ret[1].(uint64)
	return ret0, ret1
}

// ObserveAccounts indicates an expected call of ObserveAccounts
func (mr *MockAccountsServiceMockRecorder) ObserveAccounts(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObserveAccounts", reflect.TypeOf((*MockAccountsService)(nil).ObserveAccounts), arg0, arg1, arg2, arg3, arg4)
}
