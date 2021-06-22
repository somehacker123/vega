// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/processor (interfaces: Assets)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	assets "code.vegaprotocol.io/vega/assets"
	proto "code.vegaprotocol.io/vega/proto"
	gomock "github.com/golang/mock/gomock"
)

// MockAssets is a mock of Assets interface
type MockAssets struct {
	ctrl     *gomock.Controller
	recorder *MockAssetsMockRecorder
}

// MockAssetsMockRecorder is the mock recorder for MockAssets
type MockAssetsMockRecorder struct {
	mock *MockAssets
}

// NewMockAssets creates a new mock instance
func NewMockAssets(ctrl *gomock.Controller) *MockAssets {
	mock := &MockAssets{ctrl: ctrl}
	mock.recorder = &MockAssetsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAssets) EXPECT() *MockAssetsMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockAssets) Get(arg0 string) (*assets.Asset, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*assets.Asset)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockAssetsMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockAssets)(nil).Get), arg0)
}

// IsEnabled mocks base method
func (m *MockAssets) IsEnabled(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsEnabled", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsEnabled indicates an expected call of IsEnabled
func (mr *MockAssetsMockRecorder) IsEnabled(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsEnabled", reflect.TypeOf((*MockAssets)(nil).IsEnabled), arg0)
}

// NewAsset mocks base method
func (m *MockAssets) NewAsset(arg0 string, arg1 *proto.AssetDetails) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewAsset", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewAsset indicates an expected call of NewAsset
func (mr *MockAssetsMockRecorder) NewAsset(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAsset", reflect.TypeOf((*MockAssets)(nil).NewAsset), arg0, arg1)
}
