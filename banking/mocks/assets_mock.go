// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/banking (interfaces: Assets)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	assets "code.vegaprotocol.io/vega/assets"
	gomock "github.com/golang/mock/gomock"
)

// MockAssets is a mock of Assets interface.
type MockAssets struct {
	ctrl     *gomock.Controller
	recorder *MockAssetsMockRecorder
}

// MockAssetsMockRecorder is the mock recorder for MockAssets.
type MockAssetsMockRecorder struct {
	mock *MockAssets
}

// NewMockAssets creates a new mock instance.
func NewMockAssets(ctrl *gomock.Controller) *MockAssets {
	mock := &MockAssets{ctrl: ctrl}
	mock.recorder = &MockAssetsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAssets) EXPECT() *MockAssetsMockRecorder {
	return m.recorder
}

// Enable mocks base method.
func (m *MockAssets) Enable(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Enable", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Enable indicates an expected call of Enable.
func (mr *MockAssetsMockRecorder) Enable(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enable", reflect.TypeOf((*MockAssets)(nil).Enable), arg0)
}

// Get mocks base method.
func (m *MockAssets) Get(arg0 string) (*assets.Asset, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*assets.Asset)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockAssetsMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockAssets)(nil).Get), arg0)
}
