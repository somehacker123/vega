// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/internal/api (interfaces: OrderService)

// Package mocks is a generated GoMock package.
package mocks

import (
	proto "code.vegaprotocol.io/vega/proto"
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockOrderService is a mock of OrderService interface
type MockOrderService struct {
	ctrl     *gomock.Controller
	recorder *MockOrderServiceMockRecorder
}

// MockOrderServiceMockRecorder is the mock recorder for MockOrderService
type MockOrderServiceMockRecorder struct {
	mock *MockOrderService
}

// NewMockOrderService creates a new mock instance
func NewMockOrderService(ctrl *gomock.Controller) *MockOrderService {
	mock := &MockOrderService{ctrl: ctrl}
	mock.recorder = &MockOrderServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOrderService) EXPECT() *MockOrderServiceMockRecorder {
	return m.recorder
}

// GetByMarket mocks base method
func (m *MockOrderService) GetByMarket(arg0 context.Context, arg1 string, arg2, arg3 uint64, arg4 bool, arg5 *bool) ([]*proto.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByMarket", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].([]*proto.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByMarket indicates an expected call of GetByMarket
func (mr *MockOrderServiceMockRecorder) GetByMarket(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByMarket", reflect.TypeOf((*MockOrderService)(nil).GetByMarket), arg0, arg1, arg2, arg3, arg4, arg5)
}

// GetByMarketAndId mocks base method
func (m *MockOrderService) GetByMarketAndId(arg0 context.Context, arg1, arg2 string) (*proto.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByMarketAndId", arg0, arg1, arg2)
	ret0, _ := ret[0].(*proto.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByMarketAndId indicates an expected call of GetByMarketAndId
func (mr *MockOrderServiceMockRecorder) GetByMarketAndId(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByMarketAndId", reflect.TypeOf((*MockOrderService)(nil).GetByMarketAndId), arg0, arg1, arg2)
}

// GetByParty mocks base method
func (m *MockOrderService) GetByParty(arg0 context.Context, arg1 string, arg2, arg3 uint64, arg4 bool, arg5 *bool) ([]*proto.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByParty", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].([]*proto.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByParty indicates an expected call of GetByParty
func (mr *MockOrderServiceMockRecorder) GetByParty(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByParty", reflect.TypeOf((*MockOrderService)(nil).GetByParty), arg0, arg1, arg2, arg3, arg4, arg5)
}

// GetByReference mocks base method
func (m *MockOrderService) GetByReference(arg0 context.Context, arg1 string) (*proto.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByReference", arg0, arg1)
	ret0, _ := ret[0].(*proto.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByReference indicates an expected call of GetByReference
func (mr *MockOrderServiceMockRecorder) GetByReference(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByReference", reflect.TypeOf((*MockOrderService)(nil).GetByReference), arg0, arg1)
}

// GetOrderSubscribersCount mocks base method
func (m *MockOrderService) GetOrderSubscribersCount() int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderSubscribersCount")
	ret0, _ := ret[0].(int32)
	return ret0
}

// GetOrderSubscribersCount indicates an expected call of GetOrderSubscribersCount
func (mr *MockOrderServiceMockRecorder) GetOrderSubscribersCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderSubscribersCount", reflect.TypeOf((*MockOrderService)(nil).GetOrderSubscribersCount))
}

// ObserveOrders mocks base method
func (m *MockOrderService) ObserveOrders(arg0 context.Context, arg1 int, arg2, arg3 *string) (<-chan []proto.Order, uint64) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ObserveOrders", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(<-chan []proto.Order)
	ret1, _ := ret[1].(uint64)
	return ret0, ret1
}

// ObserveOrders indicates an expected call of ObserveOrders
func (mr *MockOrderServiceMockRecorder) ObserveOrders(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObserveOrders", reflect.TypeOf((*MockOrderService)(nil).ObserveOrders), arg0, arg1, arg2, arg3)
}
