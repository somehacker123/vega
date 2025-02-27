// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/datanode/service (interfaces: OrderStore,ChainStore,MarketStore,MarketDataStore,PositionStore)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	entities "code.vegaprotocol.io/vega/datanode/entities"
	gomock "github.com/golang/mock/gomock"
)

// MockOrderStore is a mock of OrderStore interface.
type MockOrderStore struct {
	ctrl     *gomock.Controller
	recorder *MockOrderStoreMockRecorder
}

// MockOrderStoreMockRecorder is the mock recorder for MockOrderStore.
type MockOrderStoreMockRecorder struct {
	mock *MockOrderStore
}

// NewMockOrderStore creates a new mock instance.
func NewMockOrderStore(ctrl *gomock.Controller) *MockOrderStore {
	mock := &MockOrderStore{ctrl: ctrl}
	mock.recorder = &MockOrderStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderStore) EXPECT() *MockOrderStoreMockRecorder {
	return m.recorder
}

// GetLiveOrders mocks base method.
func (m *MockOrderStore) GetLiveOrders(arg0 context.Context) ([]entities.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLiveOrders", arg0)
	ret0, _ := ret[0].([]entities.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLiveOrders indicates an expected call of GetLiveOrders.
func (mr *MockOrderStoreMockRecorder) GetLiveOrders(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLiveOrders", reflect.TypeOf((*MockOrderStore)(nil).GetLiveOrders), arg0)
}

// MockChainStore is a mock of ChainStore interface.
type MockChainStore struct {
	ctrl     *gomock.Controller
	recorder *MockChainStoreMockRecorder
}

// MockChainStoreMockRecorder is the mock recorder for MockChainStore.
type MockChainStoreMockRecorder struct {
	mock *MockChainStore
}

// NewMockChainStore creates a new mock instance.
func NewMockChainStore(ctrl *gomock.Controller) *MockChainStore {
	mock := &MockChainStore{ctrl: ctrl}
	mock.recorder = &MockChainStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChainStore) EXPECT() *MockChainStoreMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockChainStore) Get(arg0 context.Context) (entities.Chain, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(entities.Chain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockChainStoreMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockChainStore)(nil).Get), arg0)
}

// Set mocks base method.
func (m *MockChainStore) Set(arg0 context.Context, arg1 entities.Chain) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockChainStoreMockRecorder) Set(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockChainStore)(nil).Set), arg0, arg1)
}

// MockMarketStore is a mock of MarketStore interface.
type MockMarketStore struct {
	ctrl     *gomock.Controller
	recorder *MockMarketStoreMockRecorder
}

// MockMarketStoreMockRecorder is the mock recorder for MockMarketStore.
type MockMarketStoreMockRecorder struct {
	mock *MockMarketStore
}

// NewMockMarketStore creates a new mock instance.
func NewMockMarketStore(ctrl *gomock.Controller) *MockMarketStore {
	mock := &MockMarketStore{ctrl: ctrl}
	mock.recorder = &MockMarketStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMarketStore) EXPECT() *MockMarketStoreMockRecorder {
	return m.recorder
}

// GetAllPaged mocks base method.
func (m *MockMarketStore) GetAllPaged(arg0 context.Context, arg1 string, arg2 entities.CursorPagination, arg3 bool) ([]entities.Market, entities.PageInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllPaged", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]entities.Market)
	ret1, _ := ret[1].(entities.PageInfo)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAllPaged indicates an expected call of GetAllPaged.
func (mr *MockMarketStoreMockRecorder) GetAllPaged(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllPaged", reflect.TypeOf((*MockMarketStore)(nil).GetAllPaged), arg0, arg1, arg2, arg3)
}

// GetByID mocks base method.
func (m *MockMarketStore) GetByID(arg0 context.Context, arg1 string) (entities.Market, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0, arg1)
	ret0, _ := ret[0].(entities.Market)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockMarketStoreMockRecorder) GetByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockMarketStore)(nil).GetByID), arg0, arg1)
}

// GetByTxHash mocks base method.
func (m *MockMarketStore) GetByTxHash(arg0 context.Context, arg1 entities.TxHash) ([]entities.Market, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByTxHash", arg0, arg1)
	ret0, _ := ret[0].([]entities.Market)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByTxHash indicates an expected call of GetByTxHash.
func (mr *MockMarketStoreMockRecorder) GetByTxHash(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByTxHash", reflect.TypeOf((*MockMarketStore)(nil).GetByTxHash), arg0, arg1)
}

// ListSuccessorMarkets mocks base method.
func (m *MockMarketStore) ListSuccessorMarkets(arg0 context.Context, arg1 string, arg2 bool, arg3 entities.CursorPagination) ([]entities.SuccessorMarket, entities.PageInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSuccessorMarkets", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]entities.SuccessorMarket)
	ret1, _ := ret[1].(entities.PageInfo)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListSuccessorMarkets indicates an expected call of ListSuccessorMarkets.
func (mr *MockMarketStoreMockRecorder) ListSuccessorMarkets(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSuccessorMarkets", reflect.TypeOf((*MockMarketStore)(nil).ListSuccessorMarkets), arg0, arg1, arg2, arg3)
}

// Upsert mocks base method.
func (m *MockMarketStore) Upsert(arg0 context.Context, arg1 *entities.Market) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockMarketStoreMockRecorder) Upsert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockMarketStore)(nil).Upsert), arg0, arg1)
}

// MockMarketDataStore is a mock of MarketDataStore interface.
type MockMarketDataStore struct {
	ctrl     *gomock.Controller
	recorder *MockMarketDataStoreMockRecorder
}

// MockMarketDataStoreMockRecorder is the mock recorder for MockMarketDataStore.
type MockMarketDataStoreMockRecorder struct {
	mock *MockMarketDataStore
}

// NewMockMarketDataStore creates a new mock instance.
func NewMockMarketDataStore(ctrl *gomock.Controller) *MockMarketDataStore {
	mock := &MockMarketDataStore{ctrl: ctrl}
	mock.recorder = &MockMarketDataStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMarketDataStore) EXPECT() *MockMarketDataStoreMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockMarketDataStore) Add(arg0 *entities.MarketData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockMarketDataStoreMockRecorder) Add(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockMarketDataStore)(nil).Add), arg0)
}

// Flush mocks base method.
func (m *MockMarketDataStore) Flush(arg0 context.Context) ([]*entities.MarketData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Flush", arg0)
	ret0, _ := ret[0].([]*entities.MarketData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Flush indicates an expected call of Flush.
func (mr *MockMarketDataStoreMockRecorder) Flush(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flush", reflect.TypeOf((*MockMarketDataStore)(nil).Flush), arg0)
}

// GetBetweenDatesByID mocks base method.
func (m *MockMarketDataStore) GetBetweenDatesByID(arg0 context.Context, arg1 string, arg2, arg3 time.Time, arg4 entities.Pagination) ([]entities.MarketData, entities.PageInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBetweenDatesByID", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]entities.MarketData)
	ret1, _ := ret[1].(entities.PageInfo)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetBetweenDatesByID indicates an expected call of GetBetweenDatesByID.
func (mr *MockMarketDataStoreMockRecorder) GetBetweenDatesByID(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBetweenDatesByID", reflect.TypeOf((*MockMarketDataStore)(nil).GetBetweenDatesByID), arg0, arg1, arg2, arg3, arg4)
}

// GetFromDateByID mocks base method.
func (m *MockMarketDataStore) GetFromDateByID(arg0 context.Context, arg1 string, arg2 time.Time, arg3 entities.Pagination) ([]entities.MarketData, entities.PageInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFromDateByID", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]entities.MarketData)
	ret1, _ := ret[1].(entities.PageInfo)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetFromDateByID indicates an expected call of GetFromDateByID.
func (mr *MockMarketDataStoreMockRecorder) GetFromDateByID(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFromDateByID", reflect.TypeOf((*MockMarketDataStore)(nil).GetFromDateByID), arg0, arg1, arg2, arg3)
}

// GetMarketDataByID mocks base method.
func (m *MockMarketDataStore) GetMarketDataByID(arg0 context.Context, arg1 string) (entities.MarketData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMarketDataByID", arg0, arg1)
	ret0, _ := ret[0].(entities.MarketData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMarketDataByID indicates an expected call of GetMarketDataByID.
func (mr *MockMarketDataStoreMockRecorder) GetMarketDataByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMarketDataByID", reflect.TypeOf((*MockMarketDataStore)(nil).GetMarketDataByID), arg0, arg1)
}

// GetMarketsData mocks base method.
func (m *MockMarketDataStore) GetMarketsData(arg0 context.Context) ([]entities.MarketData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMarketsData", arg0)
	ret0, _ := ret[0].([]entities.MarketData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMarketsData indicates an expected call of GetMarketsData.
func (mr *MockMarketDataStoreMockRecorder) GetMarketsData(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMarketsData", reflect.TypeOf((*MockMarketDataStore)(nil).GetMarketsData), arg0)
}

// GetToDateByID mocks base method.
func (m *MockMarketDataStore) GetToDateByID(arg0 context.Context, arg1 string, arg2 time.Time, arg3 entities.Pagination) ([]entities.MarketData, entities.PageInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetToDateByID", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]entities.MarketData)
	ret1, _ := ret[1].(entities.PageInfo)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetToDateByID indicates an expected call of GetToDateByID.
func (mr *MockMarketDataStoreMockRecorder) GetToDateByID(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetToDateByID", reflect.TypeOf((*MockMarketDataStore)(nil).GetToDateByID), arg0, arg1, arg2, arg3)
}

// MockPositionStore is a mock of PositionStore interface.
type MockPositionStore struct {
	ctrl     *gomock.Controller
	recorder *MockPositionStoreMockRecorder
}

// MockPositionStoreMockRecorder is the mock recorder for MockPositionStore.
type MockPositionStoreMockRecorder struct {
	mock *MockPositionStore
}

// NewMockPositionStore creates a new mock instance.
func NewMockPositionStore(ctrl *gomock.Controller) *MockPositionStore {
	mock := &MockPositionStore{ctrl: ctrl}
	mock.recorder = &MockPositionStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPositionStore) EXPECT() *MockPositionStoreMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockPositionStore) Add(arg0 context.Context, arg1 entities.Position) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockPositionStoreMockRecorder) Add(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockPositionStore)(nil).Add), arg0, arg1)
}

// Flush mocks base method.
func (m *MockPositionStore) Flush(arg0 context.Context) ([]entities.Position, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Flush", arg0)
	ret0, _ := ret[0].([]entities.Position)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Flush indicates an expected call of Flush.
func (mr *MockPositionStoreMockRecorder) Flush(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flush", reflect.TypeOf((*MockPositionStore)(nil).Flush), arg0)
}

// GetAll mocks base method.
func (m *MockPositionStore) GetAll(arg0 context.Context) ([]entities.Position, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", arg0)
	ret0, _ := ret[0].([]entities.Position)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockPositionStoreMockRecorder) GetAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockPositionStore)(nil).GetAll), arg0)
}

// GetByMarket mocks base method.
func (m *MockPositionStore) GetByMarket(arg0 context.Context, arg1 string) ([]entities.Position, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByMarket", arg0, arg1)
	ret0, _ := ret[0].([]entities.Position)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByMarket indicates an expected call of GetByMarket.
func (mr *MockPositionStoreMockRecorder) GetByMarket(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByMarket", reflect.TypeOf((*MockPositionStore)(nil).GetByMarket), arg0, arg1)
}

// GetByMarketAndParties mocks base method.
func (m *MockPositionStore) GetByMarketAndParties(arg0 context.Context, arg1 string, arg2 []string) ([]entities.Position, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByMarketAndParties", arg0, arg1, arg2)
	ret0, _ := ret[0].([]entities.Position)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByMarketAndParties indicates an expected call of GetByMarketAndParties.
func (mr *MockPositionStoreMockRecorder) GetByMarketAndParties(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByMarketAndParties", reflect.TypeOf((*MockPositionStore)(nil).GetByMarketAndParties), arg0, arg1, arg2)
}

// GetByMarketAndParty mocks base method.
func (m *MockPositionStore) GetByMarketAndParty(arg0 context.Context, arg1, arg2 string) (entities.Position, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByMarketAndParty", arg0, arg1, arg2)
	ret0, _ := ret[0].(entities.Position)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByMarketAndParty indicates an expected call of GetByMarketAndParty.
func (mr *MockPositionStoreMockRecorder) GetByMarketAndParty(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByMarketAndParty", reflect.TypeOf((*MockPositionStore)(nil).GetByMarketAndParty), arg0, arg1, arg2)
}

// GetByParty mocks base method.
func (m *MockPositionStore) GetByParty(arg0 context.Context, arg1 string) ([]entities.Position, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByParty", arg0, arg1)
	ret0, _ := ret[0].([]entities.Position)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByParty indicates an expected call of GetByParty.
func (mr *MockPositionStoreMockRecorder) GetByParty(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByParty", reflect.TypeOf((*MockPositionStore)(nil).GetByParty), arg0, arg1)
}

// GetByPartyConnection mocks base method.
func (m *MockPositionStore) GetByPartyConnection(arg0 context.Context, arg1, arg2 []string, arg3 entities.CursorPagination) ([]entities.Position, entities.PageInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByPartyConnection", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]entities.Position)
	ret1, _ := ret[1].(entities.PageInfo)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetByPartyConnection indicates an expected call of GetByPartyConnection.
func (mr *MockPositionStoreMockRecorder) GetByPartyConnection(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByPartyConnection", reflect.TypeOf((*MockPositionStore)(nil).GetByPartyConnection), arg0, arg1, arg2, arg3)
}

// GetByTxHash mocks base method.
func (m *MockPositionStore) GetByTxHash(arg0 context.Context, arg1 entities.TxHash) ([]entities.Position, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByTxHash", arg0, arg1)
	ret0, _ := ret[0].([]entities.Position)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByTxHash indicates an expected call of GetByTxHash.
func (mr *MockPositionStoreMockRecorder) GetByTxHash(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByTxHash", reflect.TypeOf((*MockPositionStore)(nil).GetByTxHash), arg0, arg1)
}
