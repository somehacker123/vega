// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/liquidity (interfaces: RiskModel)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRiskModel is a mock of RiskModel interface
type MockRiskModel struct {
	ctrl     *gomock.Controller
	recorder *MockRiskModelMockRecorder
}

// MockRiskModelMockRecorder is the mock recorder for MockRiskModel
type MockRiskModelMockRecorder struct {
	mock *MockRiskModel
}

// NewMockRiskModel creates a new mock instance
func NewMockRiskModel(ctrl *gomock.Controller) *MockRiskModel {
	mock := &MockRiskModel{ctrl: ctrl}
	mock.recorder = &MockRiskModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRiskModel) EXPECT() *MockRiskModelMockRecorder {
	return m.recorder
}

// GetProjectionHorizon mocks base method
func (m *MockRiskModel) GetProjectionHorizon() float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProjectionHorizon")
	ret0, _ := ret[0].(float64)
	return ret0
}

// GetProjectionHorizon indicates an expected call of GetProjectionHorizon
func (mr *MockRiskModelMockRecorder) GetProjectionHorizon() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjectionHorizon", reflect.TypeOf((*MockRiskModel)(nil).GetProjectionHorizon))
}

// ProbabilityOfTrading mocks base method
func (m *MockRiskModel) ProbabilityOfTrading(arg0, arg1, arg2 float64, arg3, arg4 bool, arg5, arg6 float64) float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProbabilityOfTrading", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(float64)
	return ret0
}

// ProbabilityOfTrading indicates an expected call of ProbabilityOfTrading
func (mr *MockRiskModelMockRecorder) ProbabilityOfTrading(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProbabilityOfTrading", reflect.TypeOf((*MockRiskModel)(nil).ProbabilityOfTrading), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}
