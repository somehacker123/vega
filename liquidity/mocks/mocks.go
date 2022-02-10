// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/liquidity (interfaces: RiskModel,PriceMonitor,IDGen)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	num "code.vegaprotocol.io/vega/types/num"
	gomock "github.com/golang/mock/gomock"
	decimal "github.com/shopspring/decimal"
)

// MockRiskModel is a mock of RiskModel interface.
type MockRiskModel struct {
	ctrl     *gomock.Controller
	recorder *MockRiskModelMockRecorder
}

// MockRiskModelMockRecorder is the mock recorder for MockRiskModel.
type MockRiskModelMockRecorder struct {
	mock *MockRiskModel
}

// NewMockRiskModel creates a new mock instance.
func NewMockRiskModel(ctrl *gomock.Controller) *MockRiskModel {
	mock := &MockRiskModel{ctrl: ctrl}
	mock.recorder = &MockRiskModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRiskModel) EXPECT() *MockRiskModelMockRecorder {
	return m.recorder
}

// GetProjectionHorizon mocks base method.
func (m *MockRiskModel) GetProjectionHorizon() decimal.Decimal {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProjectionHorizon")
	ret0, _ := ret[0].(decimal.Decimal)
	return ret0
}

// GetProjectionHorizon indicates an expected call of GetProjectionHorizon.
func (mr *MockRiskModelMockRecorder) GetProjectionHorizon() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjectionHorizon", reflect.TypeOf((*MockRiskModel)(nil).GetProjectionHorizon))
}

// ProbabilityOfTrading mocks base method.
func (m *MockRiskModel) ProbabilityOfTrading(arg0, arg1 *num.Uint, arg2, arg3, arg4 decimal.Decimal, arg5, arg6 bool) decimal.Decimal {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProbabilityOfTrading", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(decimal.Decimal)
	return ret0
}

// ProbabilityOfTrading indicates an expected call of ProbabilityOfTrading.
func (mr *MockRiskModelMockRecorder) ProbabilityOfTrading(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProbabilityOfTrading", reflect.TypeOf((*MockRiskModel)(nil).ProbabilityOfTrading), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// MockPriceMonitor is a mock of PriceMonitor interface.
type MockPriceMonitor struct {
	ctrl     *gomock.Controller
	recorder *MockPriceMonitorMockRecorder
}

// MockPriceMonitorMockRecorder is the mock recorder for MockPriceMonitor.
type MockPriceMonitorMockRecorder struct {
	mock *MockPriceMonitor
}

// NewMockPriceMonitor creates a new mock instance.
func NewMockPriceMonitor(ctrl *gomock.Controller) *MockPriceMonitor {
	mock := &MockPriceMonitor{ctrl: ctrl}
	mock.recorder = &MockPriceMonitorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPriceMonitor) EXPECT() *MockPriceMonitorMockRecorder {
	return m.recorder
}

// GetValidPriceRange mocks base method.
func (m *MockPriceMonitor) GetValidPriceRange() (num.WrappedDecimal, num.WrappedDecimal) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidPriceRange")
	ret0, _ := ret[0].(num.WrappedDecimal)
	ret1, _ := ret[1].(num.WrappedDecimal)
	return ret0, ret1
}

// GetValidPriceRange indicates an expected call of GetValidPriceRange.
func (mr *MockPriceMonitorMockRecorder) GetValidPriceRange() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidPriceRange", reflect.TypeOf((*MockPriceMonitor)(nil).GetValidPriceRange))
}

// MockIDGen is a mock of IDGen interface.
type MockIDGen struct {
	ctrl     *gomock.Controller
	recorder *MockIDGenMockRecorder
}

// MockIDGenMockRecorder is the mock recorder for MockIDGen.
type MockIDGenMockRecorder struct {
	mock *MockIDGen
}

// NewMockIDGen creates a new mock instance.
func NewMockIDGen(ctrl *gomock.Controller) *MockIDGen {
	mock := &MockIDGen{ctrl: ctrl}
	mock.recorder = &MockIDGenMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIDGen) EXPECT() *MockIDGenMockRecorder {
	return m.recorder
}

// NextID mocks base method.
func (m *MockIDGen) NextID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NextID")
	ret0, _ := ret[0].(string)
	return ret0
}

// NextID indicates an expected call of NextID.
func (mr *MockIDGenMockRecorder) NextID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NextID", reflect.TypeOf((*MockIDGen)(nil).NextID))
}
