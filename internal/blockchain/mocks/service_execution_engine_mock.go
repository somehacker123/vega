// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/internal/blockchain (interfaces: ServiceExecutionEngine)

// Package mocks is a generated GoMock package.
package mocks

import (
	proto "code.vegaprotocol.io/vega/proto"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockServiceExecutionEngine is a mock of ServiceExecutionEngine interface
type MockServiceExecutionEngine struct {
	ctrl     *gomock.Controller
	recorder *MockServiceExecutionEngineMockRecorder
}

// MockServiceExecutionEngineMockRecorder is the mock recorder for MockServiceExecutionEngine
type MockServiceExecutionEngineMockRecorder struct {
	mock *MockServiceExecutionEngine
}

// NewMockServiceExecutionEngine creates a new mock instance
func NewMockServiceExecutionEngine(ctrl *gomock.Controller) *MockServiceExecutionEngine {
	mock := &MockServiceExecutionEngine{ctrl: ctrl}
	mock.recorder = &MockServiceExecutionEngineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServiceExecutionEngine) EXPECT() *MockServiceExecutionEngineMockRecorder {
	return m.recorder
}

// AmendOrder mocks base method
func (m *MockServiceExecutionEngine) AmendOrder(arg0 *proto.OrderAmendment) (*proto.OrderConfirmation, error) {
	ret := m.ctrl.Call(m, "AmendOrder", arg0)
	ret0, _ := ret[0].(*proto.OrderConfirmation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AmendOrder indicates an expected call of AmendOrder
func (mr *MockServiceExecutionEngineMockRecorder) AmendOrder(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AmendOrder", reflect.TypeOf((*MockServiceExecutionEngine)(nil).AmendOrder), arg0)
}

// CancelOrder mocks base method
func (m *MockServiceExecutionEngine) CancelOrder(arg0 *proto.Order) (*proto.OrderCancellationConfirmation, error) {
	ret := m.ctrl.Call(m, "CancelOrder", arg0)
	ret0, _ := ret[0].(*proto.OrderCancellationConfirmation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CancelOrder indicates an expected call of CancelOrder
func (mr *MockServiceExecutionEngineMockRecorder) CancelOrder(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelOrder", reflect.TypeOf((*MockServiceExecutionEngine)(nil).CancelOrder), arg0)
}

// Generate mocks base method
func (m *MockServiceExecutionEngine) Generate() error {
	ret := m.ctrl.Call(m, "Generate")
	ret0, _ := ret[0].(error)
	return ret0
}

// Generate indicates an expected call of Generate
func (mr *MockServiceExecutionEngineMockRecorder) Generate() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generate", reflect.TypeOf((*MockServiceExecutionEngine)(nil).Generate))
}

// Process mocks base method
func (m *MockServiceExecutionEngine) Process() error {
	ret := m.ctrl.Call(m, "Process")
	ret0, _ := ret[0].(error)
	return ret0
}

// Process indicates an expected call of Process
func (mr *MockServiceExecutionEngineMockRecorder) Process() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Process", reflect.TypeOf((*MockServiceExecutionEngine)(nil).Process))
}

// SubmitOrder mocks base method
func (m *MockServiceExecutionEngine) SubmitOrder(arg0 *proto.Order) (*proto.OrderConfirmation, error) {
	ret := m.ctrl.Call(m, "SubmitOrder", arg0)
	ret0, _ := ret[0].(*proto.OrderConfirmation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitOrder indicates an expected call of SubmitOrder
func (mr *MockServiceExecutionEngineMockRecorder) SubmitOrder(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitOrder", reflect.TypeOf((*MockServiceExecutionEngine)(nil).SubmitOrder), arg0)
}