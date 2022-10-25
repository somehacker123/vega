// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/wallet/api/node (interfaces: CoreClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	v1 "code.vegaprotocol.io/vega/protos/vega/api/v1"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockCoreClient is a mock of CoreClient interface.
type MockCoreClient struct {
	ctrl     *gomock.Controller
	recorder *MockCoreClientMockRecorder
}

// MockCoreClientMockRecorder is the mock recorder for MockCoreClient.
type MockCoreClientMockRecorder struct {
	mock *MockCoreClient
}

// NewMockCoreClient creates a new mock instance.
func NewMockCoreClient(ctrl *gomock.Controller) *MockCoreClient {
	mock := &MockCoreClient{ctrl: ctrl}
	mock.recorder = &MockCoreClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCoreClient) EXPECT() *MockCoreClientMockRecorder {
	return m.recorder
}

// CheckTransaction mocks base method.
func (m *MockCoreClient) CheckTransaction(arg0 context.Context, arg1 *v1.CheckTransactionRequest, arg2 ...grpc.CallOption) (*v1.CheckTransactionResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CheckTransaction", varargs...)
	ret0, _ := ret[0].(*v1.CheckTransactionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckTransaction indicates an expected call of CheckTransaction.
func (mr *MockCoreClientMockRecorder) CheckTransaction(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckTransaction", reflect.TypeOf((*MockCoreClient)(nil).CheckTransaction), varargs...)
}

// GetVegaTime mocks base method.
func (m *MockCoreClient) GetVegaTime(arg0 context.Context, arg1 *v1.GetVegaTimeRequest, arg2 ...grpc.CallOption) (*v1.GetVegaTimeResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetVegaTime", varargs...)
	ret0, _ := ret[0].(*v1.GetVegaTimeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVegaTime indicates an expected call of GetVegaTime.
func (mr *MockCoreClientMockRecorder) GetVegaTime(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVegaTime", reflect.TypeOf((*MockCoreClient)(nil).GetVegaTime), varargs...)
}

// Host mocks base method.
func (m *MockCoreClient) Host() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Host")
	ret0, _ := ret[0].(string)
	return ret0
}

// Host indicates an expected call of Host.
func (mr *MockCoreClientMockRecorder) Host() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Host", reflect.TypeOf((*MockCoreClient)(nil).Host))
}

// LastBlockHeight mocks base method.
func (m *MockCoreClient) LastBlockHeight(arg0 context.Context, arg1 *v1.LastBlockHeightRequest, arg2 ...grpc.CallOption) (*v1.LastBlockHeightResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "LastBlockHeight", varargs...)
	ret0, _ := ret[0].(*v1.LastBlockHeightResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LastBlockHeight indicates an expected call of LastBlockHeight.
func (mr *MockCoreClientMockRecorder) LastBlockHeight(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastBlockHeight", reflect.TypeOf((*MockCoreClient)(nil).LastBlockHeight), varargs...)
}

// Stop mocks base method.
func (m *MockCoreClient) Stop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockCoreClientMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockCoreClient)(nil).Stop))
}

// SubmitTransaction mocks base method.
func (m *MockCoreClient) SubmitTransaction(arg0 context.Context, arg1 *v1.SubmitTransactionRequest, arg2 ...grpc.CallOption) (*v1.SubmitTransactionResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SubmitTransaction", varargs...)
	ret0, _ := ret[0].(*v1.SubmitTransactionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitTransaction indicates an expected call of SubmitTransaction.
func (mr *MockCoreClientMockRecorder) SubmitTransaction(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitTransaction", reflect.TypeOf((*MockCoreClient)(nil).SubmitTransaction), varargs...)
}