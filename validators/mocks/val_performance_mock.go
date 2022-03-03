// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/validators (interfaces: ValidatorPerformance)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	v1 "code.vegaprotocol.io/protos/vega/snapshot/v1"
	gomock "github.com/golang/mock/gomock"
	decimal "github.com/shopspring/decimal"
)

// MockValidatorPerformance is a mock of ValidatorPerformance interface.
type MockValidatorPerformance struct {
	ctrl     *gomock.Controller
	recorder *MockValidatorPerformanceMockRecorder
}

// MockValidatorPerformanceMockRecorder is the mock recorder for MockValidatorPerformance.
type MockValidatorPerformanceMockRecorder struct {
	mock *MockValidatorPerformance
}

// NewMockValidatorPerformance creates a new mock instance.
func NewMockValidatorPerformance(ctrl *gomock.Controller) *MockValidatorPerformance {
	mock := &MockValidatorPerformance{ctrl: ctrl}
	mock.recorder = &MockValidatorPerformanceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockValidatorPerformance) EXPECT() *MockValidatorPerformanceMockRecorder {
	return m.recorder
}

// BeginBlock mocks base method.
func (m *MockValidatorPerformance) BeginBlock(arg0 context.Context, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BeginBlock", arg0, arg1)
}

// BeginBlock indicates an expected call of BeginBlock.
func (mr *MockValidatorPerformanceMockRecorder) BeginBlock(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginBlock", reflect.TypeOf((*MockValidatorPerformance)(nil).BeginBlock), arg0, arg1)
}

// Deserialize mocks base method.
func (m *MockValidatorPerformance) Deserialize(arg0 *v1.ValidatorPerformance) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Deserialize", arg0)
}

// Deserialize indicates an expected call of Deserialize.
func (mr *MockValidatorPerformanceMockRecorder) Deserialize(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deserialize", reflect.TypeOf((*MockValidatorPerformance)(nil).Deserialize), arg0)
}

// Reset mocks base method.
func (m *MockValidatorPerformance) Reset() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Reset")
}

// Reset indicates an expected call of Reset.
func (mr *MockValidatorPerformanceMockRecorder) Reset() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockValidatorPerformance)(nil).Reset))
}

// Serialize mocks base method.
func (m *MockValidatorPerformance) Serialize() *v1.ValidatorPerformance {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Serialize")
	ret0, _ := ret[0].(*v1.ValidatorPerformance)
	return ret0
}

// Serialize indicates an expected call of Serialize.
func (mr *MockValidatorPerformanceMockRecorder) Serialize() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Serialize", reflect.TypeOf((*MockValidatorPerformance)(nil).Serialize))
}

// ValidatorPerformanceScore mocks base method.
func (m *MockValidatorPerformance) ValidatorPerformanceScore(arg0 string, arg1, arg2 int64) decimal.Decimal {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidatorPerformanceScore", arg0, arg1, arg2)
	ret0, _ := ret[0].(decimal.Decimal)
	return ret0
}

// ValidatorPerformanceScore indicates an expected call of ValidatorPerformanceScore.
func (mr *MockValidatorPerformanceMockRecorder) ValidatorPerformanceScore(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidatorPerformanceScore", reflect.TypeOf((*MockValidatorPerformance)(nil).ValidatorPerformanceScore), arg0, arg1, arg2)
}