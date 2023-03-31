// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/core/api (interfaces: EventService)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	events "code.vegaprotocol.io/vega/core/events"
	subscribers "code.vegaprotocol.io/vega/libs/subscribers"
	v1 "code.vegaprotocol.io/vega/protos/vega/events/v1"
	gomock "github.com/golang/mock/gomock"
)

// MockEventService is a mock of EventService interface.
type MockEventService struct {
	ctrl     *gomock.Controller
	recorder *MockEventServiceMockRecorder
}

// MockEventServiceMockRecorder is the mock recorder for MockEventService.
type MockEventServiceMockRecorder struct {
	mock *MockEventService
}

// NewMockEventService creates a new mock instance.
func NewMockEventService(ctrl *gomock.Controller) *MockEventService {
	mock := &MockEventService{ctrl: ctrl}
	mock.recorder = &MockEventServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventService) EXPECT() *MockEventServiceMockRecorder {
	return m.recorder
}

// ObserveEvents mocks base method.
func (m *MockEventService) ObserveEvents(arg0 context.Context, arg1 int, arg2 []events.Type, arg3 int, arg4 ...subscribers.EventFilter) (<-chan []*v1.BusEvent, chan<- int) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3}
	for _, a := range arg4 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ObserveEvents", varargs...)
	ret0, _ := ret[0].(<-chan []*v1.BusEvent)
	ret1, _ := ret[1].(chan<- int)
	return ret0, ret1
}

// ObserveEvents indicates an expected call of ObserveEvents.
func (mr *MockEventServiceMockRecorder) ObserveEvents(arg0, arg1, arg2, arg3 interface{}, arg4 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3}, arg4...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObserveEvents", reflect.TypeOf((*MockEventService)(nil).ObserveEvents), varargs...)
}
