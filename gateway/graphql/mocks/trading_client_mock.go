// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/gateway/graphql (interfaces: TradingClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	proto "code.vegaprotocol.io/vega/proto"
	api "code.vegaprotocol.io/vega/proto/api"
	context "context"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockTradingClient is a mock of TradingClient interface
type MockTradingClient struct {
	ctrl     *gomock.Controller
	recorder *MockTradingClientMockRecorder
}

// MockTradingClientMockRecorder is the mock recorder for MockTradingClient
type MockTradingClientMockRecorder struct {
	mock *MockTradingClient
}

// NewMockTradingClient creates a new mock instance
func NewMockTradingClient(ctrl *gomock.Controller) *MockTradingClient {
	mock := &MockTradingClient{ctrl: ctrl}
	mock.recorder = &MockTradingClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTradingClient) EXPECT() *MockTradingClientMockRecorder {
	return m.recorder
}

// CancelOrder mocks base method
func (m *MockTradingClient) CancelOrder(arg0 context.Context, arg1 *api.CancelOrderRequest, arg2 ...grpc.CallOption) (*proto.PendingOrder, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CancelOrder", varargs...)
	ret0, _ := ret[0].(*proto.PendingOrder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CancelOrder indicates an expected call of CancelOrder
func (mr *MockTradingClientMockRecorder) CancelOrder(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelOrder", reflect.TypeOf((*MockTradingClient)(nil).CancelOrder), varargs...)
}

// CheckToken mocks base method
func (m *MockTradingClient) CheckToken(arg0 context.Context, arg1 *api.CheckTokenRequest, arg2 ...grpc.CallOption) (*api.CheckTokenResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CheckToken", varargs...)
	ret0, _ := ret[0].(*api.CheckTokenResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckToken indicates an expected call of CheckToken
func (mr *MockTradingClientMockRecorder) CheckToken(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckToken", reflect.TypeOf((*MockTradingClient)(nil).CheckToken), varargs...)
}

// PrepareAmendOrder mocks base method
func (m *MockTradingClient) PrepareAmendOrder(arg0 context.Context, arg1 *api.AmendOrderRequest, arg2 ...grpc.CallOption) (*api.PrepareAmendOrderResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PrepareAmendOrder", varargs...)
	ret0, _ := ret[0].(*api.PrepareAmendOrderResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareAmendOrder indicates an expected call of PrepareAmendOrder
func (mr *MockTradingClientMockRecorder) PrepareAmendOrder(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareAmendOrder", reflect.TypeOf((*MockTradingClient)(nil).PrepareAmendOrder), varargs...)
}

// PrepareCancelOrder mocks base method
func (m *MockTradingClient) PrepareCancelOrder(arg0 context.Context, arg1 *api.CancelOrderRequest, arg2 ...grpc.CallOption) (*api.PrepareCancelOrderResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PrepareCancelOrder", varargs...)
	ret0, _ := ret[0].(*api.PrepareCancelOrderResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareCancelOrder indicates an expected call of PrepareCancelOrder
func (mr *MockTradingClientMockRecorder) PrepareCancelOrder(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareCancelOrder", reflect.TypeOf((*MockTradingClient)(nil).PrepareCancelOrder), varargs...)
}

// PrepareProposal mocks base method
func (m *MockTradingClient) PrepareProposal(arg0 context.Context, arg1 *api.PrepareProposalRequest, arg2 ...grpc.CallOption) (*api.PrepareProposalResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PrepareProposal", varargs...)
	ret0, _ := ret[0].(*api.PrepareProposalResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareProposal indicates an expected call of PrepareProposal
func (mr *MockTradingClientMockRecorder) PrepareProposal(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareProposal", reflect.TypeOf((*MockTradingClient)(nil).PrepareProposal), varargs...)
}

// PrepareSubmitOrder mocks base method
func (m *MockTradingClient) PrepareSubmitOrder(arg0 context.Context, arg1 *api.SubmitOrderRequest, arg2 ...grpc.CallOption) (*api.PrepareSubmitOrderResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PrepareSubmitOrder", varargs...)
	ret0, _ := ret[0].(*api.PrepareSubmitOrderResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareSubmitOrder indicates an expected call of PrepareSubmitOrder
func (mr *MockTradingClientMockRecorder) PrepareSubmitOrder(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareSubmitOrder", reflect.TypeOf((*MockTradingClient)(nil).PrepareSubmitOrder), varargs...)
}

// PrepareVote mocks base method
func (m *MockTradingClient) PrepareVote(arg0 context.Context, arg1 *api.PrepareVoteRequest, arg2 ...grpc.CallOption) (*api.PrepareVoteResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PrepareVote", varargs...)
	ret0, _ := ret[0].(*api.PrepareVoteResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareVote indicates an expected call of PrepareVote
func (mr *MockTradingClientMockRecorder) PrepareVote(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareVote", reflect.TypeOf((*MockTradingClient)(nil).PrepareVote), varargs...)
}

// SignIn mocks base method
func (m *MockTradingClient) SignIn(arg0 context.Context, arg1 *api.SignInRequest, arg2 ...grpc.CallOption) (*api.SignInResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SignIn", varargs...)
	ret0, _ := ret[0].(*api.SignInResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignIn indicates an expected call of SignIn
func (mr *MockTradingClientMockRecorder) SignIn(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignIn", reflect.TypeOf((*MockTradingClient)(nil).SignIn), varargs...)
}

// SubmitOrder mocks base method
func (m *MockTradingClient) SubmitOrder(arg0 context.Context, arg1 *api.SubmitOrderRequest, arg2 ...grpc.CallOption) (*proto.PendingOrder, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SubmitOrder", varargs...)
	ret0, _ := ret[0].(*proto.PendingOrder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitOrder indicates an expected call of SubmitOrder
func (mr *MockTradingClientMockRecorder) SubmitOrder(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitOrder", reflect.TypeOf((*MockTradingClient)(nil).SubmitOrder), varargs...)
}

// SubmitTransaction mocks base method
func (m *MockTradingClient) SubmitTransaction(arg0 context.Context, arg1 *api.SubmitTransactionRequest, arg2 ...grpc.CallOption) (*api.SubmitTransactionResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SubmitTransaction", varargs...)
	ret0, _ := ret[0].(*api.SubmitTransactionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitTransaction indicates an expected call of SubmitTransaction
func (mr *MockTradingClientMockRecorder) SubmitTransaction(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitTransaction", reflect.TypeOf((*MockTradingClient)(nil).SubmitTransaction), varargs...)
}
