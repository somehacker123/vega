// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/governance (interfaces: Buffer)

// Package mocks is a generated GoMock package.
package mocks

import (
	proto "code.vegaprotocol.io/vega/proto"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockBuffer is a mock of Buffer interface
type MockBuffer struct {
	ctrl     *gomock.Controller
	recorder *MockBufferMockRecorder
}

// MockBufferMockRecorder is the mock recorder for MockBuffer
type MockBufferMockRecorder struct {
	mock *MockBuffer
}

// NewMockBuffer creates a new mock instance
func NewMockBuffer(ctrl *gomock.Controller) *MockBuffer {
	mock := &MockBuffer{ctrl: ctrl}
	mock.recorder = &MockBufferMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBuffer) EXPECT() *MockBufferMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockBuffer) Add(arg0 proto.Proposal) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Add", arg0)
}

// Add indicates an expected call of Add
func (mr *MockBufferMockRecorder) Add(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockBuffer)(nil).Add), arg0)
}
