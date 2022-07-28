// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/data-node/datanode/api (interfaces: LiquidityService)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	vega "code.vegaprotocol.io/protos/vega"
	gomock "github.com/golang/mock/gomock"
)

// MockLiquidityService is a mock of LiquidityService interface.
type MockLiquidityService struct {
	ctrl     *gomock.Controller
	recorder *MockLiquidityServiceMockRecorder
}

// MockLiquidityServiceMockRecorder is the mock recorder for MockLiquidityService.
type MockLiquidityServiceMockRecorder struct {
	mock *MockLiquidityService
}

// NewMockLiquidityService creates a new mock instance.
func NewMockLiquidityService(ctrl *gomock.Controller) *MockLiquidityService {
	mock := &MockLiquidityService{ctrl: ctrl}
	mock.recorder = &MockLiquidityServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLiquidityService) EXPECT() *MockLiquidityServiceMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockLiquidityService) Get(arg0, arg1 string) ([]*vega.LiquidityProvision, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].([]*vega.LiquidityProvision)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockLiquidityServiceMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockLiquidityService)(nil).Get), arg0, arg1)
}
