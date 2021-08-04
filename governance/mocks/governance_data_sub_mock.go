// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/governance (interfaces: GovernanceDataSub)

// Package mocks is a generated GoMock package.
package mocks

import (
	vega "code.vegaprotocol.io/protos/vega"
	subscribers "code.vegaprotocol.io/vega/subscribers"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockGovernanceDataSub is a mock of GovernanceDataSub interface
type MockGovernanceDataSub struct {
	ctrl     *gomock.Controller
	recorder *MockGovernanceDataSubMockRecorder
}

// MockGovernanceDataSubMockRecorder is the mock recorder for MockGovernanceDataSub
type MockGovernanceDataSubMockRecorder struct {
	mock *MockGovernanceDataSub
}

// NewMockGovernanceDataSub creates a new mock instance
func NewMockGovernanceDataSub(ctrl *gomock.Controller) *MockGovernanceDataSub {
	mock := &MockGovernanceDataSub{ctrl: ctrl}
	mock.recorder = &MockGovernanceDataSubMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGovernanceDataSub) EXPECT() *MockGovernanceDataSubMockRecorder {
	return m.recorder
}

// Filter mocks base method
func (m *MockGovernanceDataSub) Filter(arg0 bool, arg1 ...subscribers.ProposalFilter) []*vega.GovernanceData {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Filter", varargs...)
	ret0, _ := ret[0].([]*vega.GovernanceData)
	return ret0
}

// Filter indicates an expected call of Filter
func (mr *MockGovernanceDataSubMockRecorder) Filter(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Filter", reflect.TypeOf((*MockGovernanceDataSub)(nil).Filter), varargs...)
}
