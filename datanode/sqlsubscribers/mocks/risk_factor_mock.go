// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/data-node/datanode/sqlsubscribers (interfaces: RiskFactorStore)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	entities "code.vegaprotocol.io/data-node/datanode/entities"
	gomock "github.com/golang/mock/gomock"
)

// MockRiskFactorStore is a mock of RiskFactorStore interface.
type MockRiskFactorStore struct {
	ctrl     *gomock.Controller
	recorder *MockRiskFactorStoreMockRecorder
}

// MockRiskFactorStoreMockRecorder is the mock recorder for MockRiskFactorStore.
type MockRiskFactorStoreMockRecorder struct {
	mock *MockRiskFactorStore
}

// NewMockRiskFactorStore creates a new mock instance.
func NewMockRiskFactorStore(ctrl *gomock.Controller) *MockRiskFactorStore {
	mock := &MockRiskFactorStore{ctrl: ctrl}
	mock.recorder = &MockRiskFactorStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRiskFactorStore) EXPECT() *MockRiskFactorStoreMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockRiskFactorStore) Upsert(arg0 context.Context, arg1 *entities.RiskFactor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockRiskFactorStoreMockRecorder) Upsert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockRiskFactorStore)(nil).Upsert), arg0, arg1)
}
