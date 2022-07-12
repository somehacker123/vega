// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/processor (interfaces: TimeService)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockTimeService is a mock of TimeService interface.
type MockTimeService struct {
	ctrl     *gomock.Controller
	recorder *MockTimeServiceMockRecorder
}

// MockTimeServiceMockRecorder is the mock recorder for MockTimeService.
type MockTimeServiceMockRecorder struct {
	mock *MockTimeService
}

// NewMockTimeService creates a new mock instance.
func NewMockTimeService(ctrl *gomock.Controller) *MockTimeService {
	mock := &MockTimeService{ctrl: ctrl}
	mock.recorder = &MockTimeServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTimeService) EXPECT() *MockTimeServiceMockRecorder {
	return m.recorder
}

// GetTimeLastBatch mocks base method.
func (m *MockTimeService) GetTimeLastBatch() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTimeLastBatch")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetTimeLastBatch indicates an expected call of GetTimeLastBatch.
func (mr *MockTimeServiceMockRecorder) GetTimeLastBatch() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTimeLastBatch", reflect.TypeOf((*MockTimeService)(nil).GetTimeLastBatch))
}

// GetTimeNow mocks base method.
func (m *MockTimeService) GetTimeNow() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTimeNow")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetTimeNow indicates an expected call of GetTimeNow.
func (mr *MockTimeServiceMockRecorder) GetTimeNow() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTimeNow", reflect.TypeOf((*MockTimeService)(nil).GetTimeNow))
}

// NotifyOnTick mocks base method.
func (m *MockTimeService) NotifyOnTick(arg0 ...func(context.Context, time.Time)) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "NotifyOnTick", varargs...)
}

// NotifyOnTick indicates an expected call of NotifyOnTick.
func (mr *MockTimeServiceMockRecorder) NotifyOnTick(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyOnTick", reflect.TypeOf((*MockTimeService)(nil).NotifyOnTick), arg0...)
}

// SetTimeNow mocks base method.
func (m *MockTimeService) SetTimeNow(arg0 context.Context, arg1 time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTimeNow", arg0, arg1)
}

// SetTimeNow indicates an expected call of SetTimeNow.
func (mr *MockTimeServiceMockRecorder) SetTimeNow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTimeNow", reflect.TypeOf((*MockTimeService)(nil).SetTimeNow), arg0, arg1)
}
