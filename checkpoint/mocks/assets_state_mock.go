// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/checkpoint (interfaces: AssetsState)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	types "code.vegaprotocol.io/vega/types"
	gomock "github.com/golang/mock/gomock"
)

// MockAssetsState is a mock of AssetsState interface.
type MockAssetsState struct {
	ctrl     *gomock.Controller
	recorder *MockAssetsStateMockRecorder
}

// MockAssetsStateMockRecorder is the mock recorder for MockAssetsState.
type MockAssetsStateMockRecorder struct {
	mock *MockAssetsState
}

// NewMockAssetsState creates a new mock instance.
func NewMockAssetsState(ctrl *gomock.Controller) *MockAssetsState {
	mock := &MockAssetsState{ctrl: ctrl}
	mock.recorder = &MockAssetsStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAssetsState) EXPECT() *MockAssetsStateMockRecorder {
	return m.recorder
}

// Checkpoint mocks base method.
func (m *MockAssetsState) Checkpoint() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Checkpoint")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Checkpoint indicates an expected call of Checkpoint.
func (mr *MockAssetsStateMockRecorder) Checkpoint() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Checkpoint", reflect.TypeOf((*MockAssetsState)(nil).Checkpoint))
}

// GetEnabledAssets mocks base method.
func (m *MockAssetsState) GetEnabledAssets() []*types.Asset {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEnabledAssets")
	ret0, _ := ret[0].([]*types.Asset)
	return ret0
}

// GetEnabledAssets indicates an expected call of GetEnabledAssets.
func (mr *MockAssetsStateMockRecorder) GetEnabledAssets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEnabledAssets", reflect.TypeOf((*MockAssetsState)(nil).GetEnabledAssets))
}

// Load mocks base method.
func (m *MockAssetsState) Load(arg0 context.Context, arg1 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Load", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Load indicates an expected call of Load.
func (mr *MockAssetsStateMockRecorder) Load(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Load", reflect.TypeOf((*MockAssetsState)(nil).Load), arg0, arg1)
}

// Name mocks base method.
func (m *MockAssetsState) Name() types.CheckpointName {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(types.CheckpointName)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockAssetsStateMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockAssetsState)(nil).Name))
}
