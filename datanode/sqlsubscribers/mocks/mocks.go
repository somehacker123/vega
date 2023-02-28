// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/datanode/sqlsubscribers (interfaces: RiskFactorStore,TransferStore,WithdrawalStore,LiquidityProvisionStore,KeyRotationStore,OracleSpecStore,DepositStore,StakeLinkingStore,MarketDataStore,PositionStore,OracleDataStore,MarginLevelsStore,NotaryStore,NodeStore,MarketsStore,MarketSvc)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	entities "code.vegaprotocol.io/vega/datanode/entities"
	gomock "github.com/golang/mock/gomock"
	decimal "github.com/shopspring/decimal"
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

// MockTransferStore is a mock of TransferStore interface.
type MockTransferStore struct {
	ctrl     *gomock.Controller
	recorder *MockTransferStoreMockRecorder
}

// MockTransferStoreMockRecorder is the mock recorder for MockTransferStore.
type MockTransferStoreMockRecorder struct {
	mock *MockTransferStore
}

// NewMockTransferStore creates a new mock instance.
func NewMockTransferStore(ctrl *gomock.Controller) *MockTransferStore {
	mock := &MockTransferStore{ctrl: ctrl}
	mock.recorder = &MockTransferStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransferStore) EXPECT() *MockTransferStoreMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockTransferStore) Upsert(arg0 context.Context, arg1 *entities.Transfer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockTransferStoreMockRecorder) Upsert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockTransferStore)(nil).Upsert), arg0, arg1)
}

// MockWithdrawalStore is a mock of WithdrawalStore interface.
type MockWithdrawalStore struct {
	ctrl     *gomock.Controller
	recorder *MockWithdrawalStoreMockRecorder
}

// MockWithdrawalStoreMockRecorder is the mock recorder for MockWithdrawalStore.
type MockWithdrawalStoreMockRecorder struct {
	mock *MockWithdrawalStore
}

// NewMockWithdrawalStore creates a new mock instance.
func NewMockWithdrawalStore(ctrl *gomock.Controller) *MockWithdrawalStore {
	mock := &MockWithdrawalStore{ctrl: ctrl}
	mock.recorder = &MockWithdrawalStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWithdrawalStore) EXPECT() *MockWithdrawalStoreMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockWithdrawalStore) Upsert(arg0 context.Context, arg1 *entities.Withdrawal) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockWithdrawalStoreMockRecorder) Upsert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockWithdrawalStore)(nil).Upsert), arg0, arg1)
}

// MockLiquidityProvisionStore is a mock of LiquidityProvisionStore interface.
type MockLiquidityProvisionStore struct {
	ctrl     *gomock.Controller
	recorder *MockLiquidityProvisionStoreMockRecorder
}

// MockLiquidityProvisionStoreMockRecorder is the mock recorder for MockLiquidityProvisionStore.
type MockLiquidityProvisionStoreMockRecorder struct {
	mock *MockLiquidityProvisionStore
}

// NewMockLiquidityProvisionStore creates a new mock instance.
func NewMockLiquidityProvisionStore(ctrl *gomock.Controller) *MockLiquidityProvisionStore {
	mock := &MockLiquidityProvisionStore{ctrl: ctrl}
	mock.recorder = &MockLiquidityProvisionStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLiquidityProvisionStore) EXPECT() *MockLiquidityProvisionStoreMockRecorder {
	return m.recorder
}

// Flush mocks base method.
func (m *MockLiquidityProvisionStore) Flush(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Flush", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Flush indicates an expected call of Flush.
func (mr *MockLiquidityProvisionStoreMockRecorder) Flush(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flush", reflect.TypeOf((*MockLiquidityProvisionStore)(nil).Flush), arg0)
}

// Upsert mocks base method.
func (m *MockLiquidityProvisionStore) Upsert(arg0 context.Context, arg1 entities.LiquidityProvision) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockLiquidityProvisionStoreMockRecorder) Upsert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockLiquidityProvisionStore)(nil).Upsert), arg0, arg1)
}

// MockKeyRotationStore is a mock of KeyRotationStore interface.
type MockKeyRotationStore struct {
	ctrl     *gomock.Controller
	recorder *MockKeyRotationStoreMockRecorder
}

// MockKeyRotationStoreMockRecorder is the mock recorder for MockKeyRotationStore.
type MockKeyRotationStoreMockRecorder struct {
	mock *MockKeyRotationStore
}

// NewMockKeyRotationStore creates a new mock instance.
func NewMockKeyRotationStore(ctrl *gomock.Controller) *MockKeyRotationStore {
	mock := &MockKeyRotationStore{ctrl: ctrl}
	mock.recorder = &MockKeyRotationStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKeyRotationStore) EXPECT() *MockKeyRotationStoreMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockKeyRotationStore) Upsert(arg0 context.Context, arg1 *entities.KeyRotation) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockKeyRotationStoreMockRecorder) Upsert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockKeyRotationStore)(nil).Upsert), arg0, arg1)
}

// MockOracleSpecStore is a mock of OracleSpecStore interface.
type MockOracleSpecStore struct {
	ctrl     *gomock.Controller
	recorder *MockOracleSpecStoreMockRecorder
}

// MockOracleSpecStoreMockRecorder is the mock recorder for MockOracleSpecStore.
type MockOracleSpecStoreMockRecorder struct {
	mock *MockOracleSpecStore
}

// NewMockOracleSpecStore creates a new mock instance.
func NewMockOracleSpecStore(ctrl *gomock.Controller) *MockOracleSpecStore {
	mock := &MockOracleSpecStore{ctrl: ctrl}
	mock.recorder = &MockOracleSpecStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOracleSpecStore) EXPECT() *MockOracleSpecStoreMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockOracleSpecStore) Upsert(arg0 context.Context, arg1 *entities.OracleSpec) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockOracleSpecStoreMockRecorder) Upsert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockOracleSpecStore)(nil).Upsert), arg0, arg1)
}

// MockDepositStore is a mock of DepositStore interface.
type MockDepositStore struct {
	ctrl     *gomock.Controller
	recorder *MockDepositStoreMockRecorder
}

// MockDepositStoreMockRecorder is the mock recorder for MockDepositStore.
type MockDepositStoreMockRecorder struct {
	mock *MockDepositStore
}

// NewMockDepositStore creates a new mock instance.
func NewMockDepositStore(ctrl *gomock.Controller) *MockDepositStore {
	mock := &MockDepositStore{ctrl: ctrl}
	mock.recorder = &MockDepositStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDepositStore) EXPECT() *MockDepositStoreMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockDepositStore) Upsert(arg0 context.Context, arg1 *entities.Deposit) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockDepositStoreMockRecorder) Upsert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockDepositStore)(nil).Upsert), arg0, arg1)
}

// MockStakeLinkingStore is a mock of StakeLinkingStore interface.
type MockStakeLinkingStore struct {
	ctrl     *gomock.Controller
	recorder *MockStakeLinkingStoreMockRecorder
}

// MockStakeLinkingStoreMockRecorder is the mock recorder for MockStakeLinkingStore.
type MockStakeLinkingStoreMockRecorder struct {
	mock *MockStakeLinkingStore
}

// NewMockStakeLinkingStore creates a new mock instance.
func NewMockStakeLinkingStore(ctrl *gomock.Controller) *MockStakeLinkingStore {
	mock := &MockStakeLinkingStore{ctrl: ctrl}
	mock.recorder = &MockStakeLinkingStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStakeLinkingStore) EXPECT() *MockStakeLinkingStoreMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockStakeLinkingStore) Upsert(arg0 context.Context, arg1 *entities.StakeLinking) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockStakeLinkingStoreMockRecorder) Upsert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockStakeLinkingStore)(nil).Upsert), arg0, arg1)
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
func (m *MockMarketDataStore) Flush(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Flush", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Flush indicates an expected call of Flush.
func (mr *MockMarketDataStoreMockRecorder) Flush(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flush", reflect.TypeOf((*MockMarketDataStore)(nil).Flush), arg0)
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
func (m *MockPositionStore) Flush(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Flush", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Flush indicates an expected call of Flush.
func (mr *MockPositionStoreMockRecorder) Flush(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flush", reflect.TypeOf((*MockPositionStore)(nil).Flush), arg0)
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

// MockOracleDataStore is a mock of OracleDataStore interface.
type MockOracleDataStore struct {
	ctrl     *gomock.Controller
	recorder *MockOracleDataStoreMockRecorder
}

// MockOracleDataStoreMockRecorder is the mock recorder for MockOracleDataStore.
type MockOracleDataStoreMockRecorder struct {
	mock *MockOracleDataStore
}

// NewMockOracleDataStore creates a new mock instance.
func NewMockOracleDataStore(ctrl *gomock.Controller) *MockOracleDataStore {
	mock := &MockOracleDataStore{ctrl: ctrl}
	mock.recorder = &MockOracleDataStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOracleDataStore) EXPECT() *MockOracleDataStoreMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockOracleDataStore) Add(arg0 context.Context, arg1 *entities.OracleData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockOracleDataStoreMockRecorder) Add(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockOracleDataStore)(nil).Add), arg0, arg1)
}

// MockMarginLevelsStore is a mock of MarginLevelsStore interface.
type MockMarginLevelsStore struct {
	ctrl     *gomock.Controller
	recorder *MockMarginLevelsStoreMockRecorder
}

// MockMarginLevelsStoreMockRecorder is the mock recorder for MockMarginLevelsStore.
type MockMarginLevelsStoreMockRecorder struct {
	mock *MockMarginLevelsStore
}

// NewMockMarginLevelsStore creates a new mock instance.
func NewMockMarginLevelsStore(ctrl *gomock.Controller) *MockMarginLevelsStore {
	mock := &MockMarginLevelsStore{ctrl: ctrl}
	mock.recorder = &MockMarginLevelsStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMarginLevelsStore) EXPECT() *MockMarginLevelsStoreMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockMarginLevelsStore) Add(arg0 entities.MarginLevels) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockMarginLevelsStoreMockRecorder) Add(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockMarginLevelsStore)(nil).Add), arg0)
}

// Flush mocks base method.
func (m *MockMarginLevelsStore) Flush(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Flush", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Flush indicates an expected call of Flush.
func (mr *MockMarginLevelsStoreMockRecorder) Flush(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flush", reflect.TypeOf((*MockMarginLevelsStore)(nil).Flush), arg0)
}

// MockNotaryStore is a mock of NotaryStore interface.
type MockNotaryStore struct {
	ctrl     *gomock.Controller
	recorder *MockNotaryStoreMockRecorder
}

// MockNotaryStoreMockRecorder is the mock recorder for MockNotaryStore.
type MockNotaryStoreMockRecorder struct {
	mock *MockNotaryStore
}

// NewMockNotaryStore creates a new mock instance.
func NewMockNotaryStore(ctrl *gomock.Controller) *MockNotaryStore {
	mock := &MockNotaryStore{ctrl: ctrl}
	mock.recorder = &MockNotaryStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNotaryStore) EXPECT() *MockNotaryStoreMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockNotaryStore) Add(arg0 context.Context, arg1 *entities.NodeSignature) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockNotaryStoreMockRecorder) Add(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockNotaryStore)(nil).Add), arg0, arg1)
}

// MockNodeStore is a mock of NodeStore interface.
type MockNodeStore struct {
	ctrl     *gomock.Controller
	recorder *MockNodeStoreMockRecorder
}

// MockNodeStoreMockRecorder is the mock recorder for MockNodeStore.
type MockNodeStoreMockRecorder struct {
	mock *MockNodeStore
}

// NewMockNodeStore creates a new mock instance.
func NewMockNodeStore(ctrl *gomock.Controller) *MockNodeStore {
	mock := &MockNodeStore{ctrl: ctrl}
	mock.recorder = &MockNodeStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNodeStore) EXPECT() *MockNodeStoreMockRecorder {
	return m.recorder
}

// AddNodeAnnouncedEvent mocks base method.
func (m *MockNodeStore) AddNodeAnnouncedEvent(arg0 context.Context, arg1 string, arg2 time.Time, arg3 *entities.ValidatorUpdateAux) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddNodeAnnouncedEvent", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddNodeAnnouncedEvent indicates an expected call of AddNodeAnnouncedEvent.
func (mr *MockNodeStoreMockRecorder) AddNodeAnnouncedEvent(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNodeAnnouncedEvent", reflect.TypeOf((*MockNodeStore)(nil).AddNodeAnnouncedEvent), arg0, arg1, arg2, arg3)
}

// UpdateEthereumAddress mocks base method.
func (m *MockNodeStore) UpdateEthereumAddress(arg0 context.Context, arg1 entities.EthereumKeyRotation) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEthereumAddress", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateEthereumAddress indicates an expected call of UpdateEthereumAddress.
func (mr *MockNodeStoreMockRecorder) UpdateEthereumAddress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEthereumAddress", reflect.TypeOf((*MockNodeStore)(nil).UpdateEthereumAddress), arg0, arg1)
}

// UpdatePublicKey mocks base method.
func (m *MockNodeStore) UpdatePublicKey(arg0 context.Context, arg1 *entities.KeyRotation) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePublicKey", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePublicKey indicates an expected call of UpdatePublicKey.
func (mr *MockNodeStoreMockRecorder) UpdatePublicKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePublicKey", reflect.TypeOf((*MockNodeStore)(nil).UpdatePublicKey), arg0, arg1)
}

// UpsertNode mocks base method.
func (m *MockNodeStore) UpsertNode(arg0 context.Context, arg1 *entities.Node) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertNode", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertNode indicates an expected call of UpsertNode.
func (mr *MockNodeStoreMockRecorder) UpsertNode(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertNode", reflect.TypeOf((*MockNodeStore)(nil).UpsertNode), arg0, arg1)
}

// UpsertRanking mocks base method.
func (m *MockNodeStore) UpsertRanking(arg0 context.Context, arg1 *entities.RankingScore, arg2 *entities.RankingScoreAux) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertRanking", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertRanking indicates an expected call of UpsertRanking.
func (mr *MockNodeStoreMockRecorder) UpsertRanking(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertRanking", reflect.TypeOf((*MockNodeStore)(nil).UpsertRanking), arg0, arg1, arg2)
}

// UpsertScore mocks base method.
func (m *MockNodeStore) UpsertScore(arg0 context.Context, arg1 *entities.RewardScore, arg2 *entities.RewardScoreAux) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertScore", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertScore indicates an expected call of UpsertScore.
func (mr *MockNodeStoreMockRecorder) UpsertScore(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertScore", reflect.TypeOf((*MockNodeStore)(nil).UpsertScore), arg0, arg1, arg2)
}

// MockMarketsStore is a mock of MarketsStore interface.
type MockMarketsStore struct {
	ctrl     *gomock.Controller
	recorder *MockMarketsStoreMockRecorder
}

// MockMarketsStoreMockRecorder is the mock recorder for MockMarketsStore.
type MockMarketsStoreMockRecorder struct {
	mock *MockMarketsStore
}

// NewMockMarketsStore creates a new mock instance.
func NewMockMarketsStore(ctrl *gomock.Controller) *MockMarketsStore {
	mock := &MockMarketsStore{ctrl: ctrl}
	mock.recorder = &MockMarketsStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMarketsStore) EXPECT() *MockMarketsStoreMockRecorder {
	return m.recorder
}

// Upsert mocks base method.
func (m *MockMarketsStore) Upsert(arg0 context.Context, arg1 *entities.Market) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockMarketsStoreMockRecorder) Upsert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockMarketsStore)(nil).Upsert), arg0, arg1)
}

// MockMarketSvc is a mock of MarketSvc interface.
type MockMarketSvc struct {
	ctrl     *gomock.Controller
	recorder *MockMarketSvcMockRecorder
}

// MockMarketSvcMockRecorder is the mock recorder for MockMarketSvc.
type MockMarketSvcMockRecorder struct {
	mock *MockMarketSvc
}

// NewMockMarketSvc creates a new mock instance.
func NewMockMarketSvc(ctrl *gomock.Controller) *MockMarketSvc {
	mock := &MockMarketSvc{ctrl: ctrl}
	mock.recorder = &MockMarketSvcMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMarketSvc) EXPECT() *MockMarketSvcMockRecorder {
	return m.recorder
}

// GetMarketScalingFactor mocks base method.
func (m *MockMarketSvc) GetMarketScalingFactor(arg0 context.Context, arg1 string) (decimal.Decimal, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMarketScalingFactor", arg0, arg1)
	ret0, _ := ret[0].(decimal.Decimal)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetMarketScalingFactor indicates an expected call of GetMarketScalingFactor.
func (mr *MockMarketSvcMockRecorder) GetMarketScalingFactor(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMarketScalingFactor", reflect.TypeOf((*MockMarketSvc)(nil).GetMarketScalingFactor), arg0, arg1)
}
