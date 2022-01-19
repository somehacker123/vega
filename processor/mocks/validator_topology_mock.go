// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/processor (interfaces: ValidatorTopology)

// Package mocks is a generated GoMock package.
package mocks

import (
	v1 "code.vegaprotocol.io/protos/vega/commands/v1"
	context "context"
	gomock "github.com/golang/mock/gomock"
	types "github.com/tendermint/tendermint/abci/types"
	types0 "github.com/tendermint/tendermint/types"
	reflect "reflect"
)

// MockValidatorTopology is a mock of ValidatorTopology interface
type MockValidatorTopology struct {
	ctrl     *gomock.Controller
	recorder *MockValidatorTopologyMockRecorder
}

// MockValidatorTopologyMockRecorder is the mock recorder for MockValidatorTopology
type MockValidatorTopologyMockRecorder struct {
	mock *MockValidatorTopology
}

// NewMockValidatorTopology creates a new mock instance
func NewMockValidatorTopology(ctrl *gomock.Controller) *MockValidatorTopology {
	mock := &MockValidatorTopology{ctrl: ctrl}
	mock.recorder = &MockValidatorTopologyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockValidatorTopology) EXPECT() *MockValidatorTopologyMockRecorder {
	return m.recorder
}

// AddKeyRotate mocks base method
func (m *MockValidatorTopology) AddKeyRotate(arg0 context.Context, arg1 string, arg2 uint64, arg3 *v1.KeyRotateSubmission) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddKeyRotate", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddKeyRotate indicates an expected call of AddKeyRotate
func (mr *MockValidatorTopologyMockRecorder) AddKeyRotate(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddKeyRotate", reflect.TypeOf((*MockValidatorTopology)(nil).AddKeyRotate), arg0, arg1, arg2, arg3)
}

// AddNodeRegistration mocks base method
func (m *MockValidatorTopology) AddNodeRegistration(arg0 context.Context, arg1 *v1.NodeRegistration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddNodeRegistration", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddNodeRegistration indicates an expected call of AddNodeRegistration
func (mr *MockValidatorTopologyMockRecorder) AddNodeRegistration(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNodeRegistration", reflect.TypeOf((*MockValidatorTopology)(nil).AddNodeRegistration), arg0, arg1)
}

// AllVegaPubKeys mocks base method
func (m *MockValidatorTopology) AllVegaPubKeys() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllVegaPubKeys")
	ret0, _ := ret[0].([]string)
	return ret0
}

// AllVegaPubKeys indicates an expected call of AllVegaPubKeys
func (mr *MockValidatorTopologyMockRecorder) AllVegaPubKeys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllVegaPubKeys", reflect.TypeOf((*MockValidatorTopology)(nil).AllVegaPubKeys))
}

// BeginBlock mocks base method
func (m *MockValidatorTopology) BeginBlock(arg0 context.Context, arg1 types.RequestBeginBlock, arg2 []*types0.Validator) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BeginBlock", arg0, arg1, arg2)
}

// BeginBlock indicates an expected call of BeginBlock
func (mr *MockValidatorTopologyMockRecorder) BeginBlock(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginBlock", reflect.TypeOf((*MockValidatorTopology)(nil).BeginBlock), arg0, arg1, arg2)
}

// IsValidator mocks base method
func (m *MockValidatorTopology) IsValidator() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsValidator")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsValidator indicates an expected call of IsValidator
func (mr *MockValidatorTopologyMockRecorder) IsValidator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsValidator", reflect.TypeOf((*MockValidatorTopology)(nil).IsValidator))
}

// IsValidatorNodeID mocks base method
func (m *MockValidatorTopology) IsValidatorNodeID(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsValidatorNodeID", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsValidatorNodeID indicates an expected call of IsValidatorNodeID
func (mr *MockValidatorTopologyMockRecorder) IsValidatorNodeID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsValidatorNodeID", reflect.TypeOf((*MockValidatorTopology)(nil).IsValidatorNodeID), arg0)
}

// IsValidatorVegaPubKey mocks base method
func (m *MockValidatorTopology) IsValidatorVegaPubKey(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsValidatorVegaPubKey", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsValidatorVegaPubKey indicates an expected call of IsValidatorVegaPubKey
func (mr *MockValidatorTopologyMockRecorder) IsValidatorVegaPubKey(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsValidatorVegaPubKey", reflect.TypeOf((*MockValidatorTopology)(nil).IsValidatorVegaPubKey), arg0)
}

// Len mocks base method
func (m *MockValidatorTopology) Len() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Len")
	ret0, _ := ret[0].(int)
	return ret0
}

// Len indicates an expected call of Len
func (mr *MockValidatorTopologyMockRecorder) Len() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Len", reflect.TypeOf((*MockValidatorTopology)(nil).Len))
}

// UpdateValidatorSet mocks base method
func (m *MockValidatorTopology) UpdateValidatorSet(arg0 []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateValidatorSet", arg0)
}

// UpdateValidatorSet indicates an expected call of UpdateValidatorSet
func (mr *MockValidatorTopologyMockRecorder) UpdateValidatorSet(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateValidatorSet", reflect.TypeOf((*MockValidatorTopology)(nil).UpdateValidatorSet), arg0)
}
