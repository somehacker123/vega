// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/data-node/gateway/graphql (interfaces: TradingDataServiceClientV2)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	v2 "code.vegaprotocol.io/protos/data-node/api/v2"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockTradingDataServiceClientV2 is a mock of TradingDataServiceClientV2 interface.
type MockTradingDataServiceClientV2 struct {
	ctrl     *gomock.Controller
	recorder *MockTradingDataServiceClientV2MockRecorder
}

// MockTradingDataServiceClientV2MockRecorder is the mock recorder for MockTradingDataServiceClientV2.
type MockTradingDataServiceClientV2MockRecorder struct {
	mock *MockTradingDataServiceClientV2
}

// NewMockTradingDataServiceClientV2 creates a new mock instance.
func NewMockTradingDataServiceClientV2(ctrl *gomock.Controller) *MockTradingDataServiceClientV2 {
	mock := &MockTradingDataServiceClientV2{ctrl: ctrl}
	mock.recorder = &MockTradingDataServiceClientV2MockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTradingDataServiceClientV2) EXPECT() *MockTradingDataServiceClientV2MockRecorder {
	return m.recorder
}

// GetCandleData mocks base method.
func (m *MockTradingDataServiceClientV2) GetCandleData(arg0 context.Context, arg1 *v2.GetCandleDataRequest, arg2 ...grpc.CallOption) (*v2.GetCandleDataResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCandleData", varargs...)
	ret0, _ := ret[0].(*v2.GetCandleDataResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCandleData indicates an expected call of GetCandleData.
func (mr *MockTradingDataServiceClientV2MockRecorder) GetCandleData(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCandleData", reflect.TypeOf((*MockTradingDataServiceClientV2)(nil).GetCandleData), varargs...)
}

// GetCandlesForMarket mocks base method.
func (m *MockTradingDataServiceClientV2) GetCandlesForMarket(arg0 context.Context, arg1 *v2.GetCandlesForMarketRequest, arg2 ...grpc.CallOption) (*v2.GetCandlesForMarketResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCandlesForMarket", varargs...)
	ret0, _ := ret[0].(*v2.GetCandlesForMarketResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCandlesForMarket indicates an expected call of GetCandlesForMarket.
func (mr *MockTradingDataServiceClientV2MockRecorder) GetCandlesForMarket(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCandlesForMarket", reflect.TypeOf((*MockTradingDataServiceClientV2)(nil).GetCandlesForMarket), varargs...)
}

// GetERC20MultiSigSignerAddedBundles mocks base method.
func (m *MockTradingDataServiceClientV2) GetERC20MultiSigSignerAddedBundles(arg0 context.Context, arg1 *v2.GetERC20MultiSigSignerAddedBundlesRequest, arg2 ...grpc.CallOption) (*v2.GetERC20MultiSigSignerAddedBundlesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetERC20MultiSigSignerAddedBundles", varargs...)
	ret0, _ := ret[0].(*v2.GetERC20MultiSigSignerAddedBundlesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetERC20MultiSigSignerAddedBundles indicates an expected call of GetERC20MultiSigSignerAddedBundles.
func (mr *MockTradingDataServiceClientV2MockRecorder) GetERC20MultiSigSignerAddedBundles(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetERC20MultiSigSignerAddedBundles", reflect.TypeOf((*MockTradingDataServiceClientV2)(nil).GetERC20MultiSigSignerAddedBundles), varargs...)
}

// GetERC20MultiSigSignerRemovedBundles mocks base method.
func (m *MockTradingDataServiceClientV2) GetERC20MultiSigSignerRemovedBundles(arg0 context.Context, arg1 *v2.GetERC20MultiSigSignerRemovedBundlesRequest, arg2 ...grpc.CallOption) (*v2.GetERC20MultiSigSignerRemovedBundlesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetERC20MultiSigSignerRemovedBundles", varargs...)
	ret0, _ := ret[0].(*v2.GetERC20MultiSigSignerRemovedBundlesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetERC20MultiSigSignerRemovedBundles indicates an expected call of GetERC20MultiSigSignerRemovedBundles.
func (mr *MockTradingDataServiceClientV2MockRecorder) GetERC20MultiSigSignerRemovedBundles(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetERC20MultiSigSignerRemovedBundles", reflect.TypeOf((*MockTradingDataServiceClientV2)(nil).GetERC20MultiSigSignerRemovedBundles), varargs...)
}

// GetMarketDataHistoryByID mocks base method.
func (m *MockTradingDataServiceClientV2) GetMarketDataHistoryByID(arg0 context.Context, arg1 *v2.GetMarketDataHistoryByIDRequest, arg2 ...grpc.CallOption) (*v2.GetMarketDataHistoryByIDResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMarketDataHistoryByID", varargs...)
	ret0, _ := ret[0].(*v2.GetMarketDataHistoryByIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMarketDataHistoryByID indicates an expected call of GetMarketDataHistoryByID.
func (mr *MockTradingDataServiceClientV2MockRecorder) GetMarketDataHistoryByID(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMarketDataHistoryByID", reflect.TypeOf((*MockTradingDataServiceClientV2)(nil).GetMarketDataHistoryByID), varargs...)
}

// GetNetworkLimits mocks base method.
func (m *MockTradingDataServiceClientV2) GetNetworkLimits(arg0 context.Context, arg1 *v2.GetNetworkLimitsRequest, arg2 ...grpc.CallOption) (*v2.GetNetworkLimitsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetNetworkLimits", varargs...)
	ret0, _ := ret[0].(*v2.GetNetworkLimitsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNetworkLimits indicates an expected call of GetNetworkLimits.
func (mr *MockTradingDataServiceClientV2MockRecorder) GetNetworkLimits(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetworkLimits", reflect.TypeOf((*MockTradingDataServiceClientV2)(nil).GetNetworkLimits), varargs...)
}

// GetOracleDataBySpecID mocks base method.
func (m *MockTradingDataServiceClientV2) GetOracleDataBySpecID(arg0 context.Context, arg1 *v2.GetOracleDataBySpecIDRequest, arg2 ...grpc.CallOption) (*v2.GetOracleDataBySpecIDResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetOracleDataBySpecID", varargs...)
	ret0, _ := ret[0].(*v2.GetOracleDataBySpecIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOracleDataBySpecID indicates an expected call of GetOracleDataBySpecID.
func (mr *MockTradingDataServiceClientV2MockRecorder) GetOracleDataBySpecID(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOracleDataBySpecID", reflect.TypeOf((*MockTradingDataServiceClientV2)(nil).GetOracleDataBySpecID), varargs...)
}

// GetOracleSpecByID mocks base method.
func (m *MockTradingDataServiceClientV2) GetOracleSpecByID(arg0 context.Context, arg1 *v2.GetOracleSpecByIDRequest, arg2 ...grpc.CallOption) (*v2.GetOracleSpecByIDResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetOracleSpecByID", varargs...)
	ret0, _ := ret[0].(*v2.GetOracleSpecByIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOracleSpecByID indicates an expected call of GetOracleSpecByID.
func (mr *MockTradingDataServiceClientV2MockRecorder) GetOracleSpecByID(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOracleSpecByID", reflect.TypeOf((*MockTradingDataServiceClientV2)(nil).GetOracleSpecByID), varargs...)
}

// ListOracleData mocks base method.
func (m *MockTradingDataServiceClientV2) ListOracleData(arg0 context.Context, arg1 *v2.ListOracleDataRequest, arg2 ...grpc.CallOption) (*v2.ListOracleDataResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListOracleData", varargs...)
	ret0, _ := ret[0].(*v2.ListOracleDataResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOracleData indicates an expected call of ListOracleData.
func (mr *MockTradingDataServiceClientV2MockRecorder) ListOracleData(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOracleData", reflect.TypeOf((*MockTradingDataServiceClientV2)(nil).ListOracleData), varargs...)
}

// ListOracleSpecs mocks base method.
func (m *MockTradingDataServiceClientV2) ListOracleSpecs(arg0 context.Context, arg1 *v2.ListOracleSpecsRequest, arg2 ...grpc.CallOption) (*v2.ListOracleSpecsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListOracleSpecs", varargs...)
	ret0, _ := ret[0].(*v2.ListOracleSpecsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOracleSpecs indicates an expected call of ListOracleSpecs.
func (mr *MockTradingDataServiceClientV2MockRecorder) ListOracleSpecs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOracleSpecs", reflect.TypeOf((*MockTradingDataServiceClientV2)(nil).ListOracleSpecs), varargs...)
}

// OrderVersionsByID mocks base method.
func (m *MockTradingDataServiceClientV2) OrderVersionsByID(arg0 context.Context, arg1 *v2.OrderVersionsByIDRequest, arg2 ...grpc.CallOption) (*v2.OrderVersionsByIDResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "OrderVersionsByID", varargs...)
	ret0, _ := ret[0].(*v2.OrderVersionsByIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OrderVersionsByID indicates an expected call of OrderVersionsByID.
func (mr *MockTradingDataServiceClientV2MockRecorder) OrderVersionsByID(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderVersionsByID", reflect.TypeOf((*MockTradingDataServiceClientV2)(nil).OrderVersionsByID), varargs...)
}

// OrdersByMarket mocks base method.
func (m *MockTradingDataServiceClientV2) OrdersByMarket(arg0 context.Context, arg1 *v2.OrdersByMarketRequest, arg2 ...grpc.CallOption) (*v2.OrdersByMarketResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "OrdersByMarket", varargs...)
	ret0, _ := ret[0].(*v2.OrdersByMarketResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OrdersByMarket indicates an expected call of OrdersByMarket.
func (mr *MockTradingDataServiceClientV2MockRecorder) OrdersByMarket(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrdersByMarket", reflect.TypeOf((*MockTradingDataServiceClientV2)(nil).OrdersByMarket), varargs...)
}

// QueryBalanceHistory mocks base method.
func (m *MockTradingDataServiceClientV2) QueryBalanceHistory(arg0 context.Context, arg1 *v2.QueryBalanceHistoryRequest, arg2 ...grpc.CallOption) (*v2.QueryBalanceHistoryResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryBalanceHistory", varargs...)
	ret0, _ := ret[0].(*v2.QueryBalanceHistoryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryBalanceHistory indicates an expected call of QueryBalanceHistory.
func (mr *MockTradingDataServiceClientV2MockRecorder) QueryBalanceHistory(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryBalanceHistory", reflect.TypeOf((*MockTradingDataServiceClientV2)(nil).QueryBalanceHistory), varargs...)
}

// SubscribeToCandleData mocks base method.
func (m *MockTradingDataServiceClientV2) SubscribeToCandleData(arg0 context.Context, arg1 *v2.SubscribeToCandleDataRequest, arg2 ...grpc.CallOption) (v2.TradingDataService_SubscribeToCandleDataClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SubscribeToCandleData", varargs...)
	ret0, _ := ret[0].(v2.TradingDataService_SubscribeToCandleDataClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubscribeToCandleData indicates an expected call of SubscribeToCandleData.
func (mr *MockTradingDataServiceClientV2MockRecorder) SubscribeToCandleData(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeToCandleData", reflect.TypeOf((*MockTradingDataServiceClientV2)(nil).SubscribeToCandleData), varargs...)
}
