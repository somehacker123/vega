// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/validators (interfaces: Commander)

// Package mocks is a generated GoMock package.
package mocks

import (
	txn "code.vegaprotocol.io/vega/txn"
	gomock "github.com/golang/mock/gomock"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	reflect "reflect"
)

// MockCommander is a mock of Commander interface
type MockCommander struct {
	ctrl     *gomock.Controller
	recorder *MockCommanderMockRecorder
}

// MockCommanderMockRecorder is the mock recorder for MockCommander
type MockCommanderMockRecorder struct {
	mock *MockCommander
}

// NewMockCommander creates a new mock instance
func NewMockCommander(ctrl *gomock.Controller) *MockCommander {
	mock := &MockCommander{ctrl: ctrl}
	mock.recorder = &MockCommanderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCommander) EXPECT() *MockCommanderMockRecorder {
	return m.recorder
}

// Command mocks base method
func (m *MockCommander) Command(arg0 txn.Command, arg1 protoiface.MessageV1) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Command", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Command indicates an expected call of Command
func (mr *MockCommanderMockRecorder) Command(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Command", reflect.TypeOf((*MockCommander)(nil).Command), arg0, arg1)
}
