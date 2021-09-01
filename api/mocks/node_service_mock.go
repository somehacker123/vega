// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/data-node/api (interfaces: NodeService)

// Package mocks is a generated GoMock package.
package mocks

import (
	vega "code.vegaprotocol.io/protos/vega"
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockNodeService is a mock of NodeService interface
type MockNodeService struct {
	ctrl     *gomock.Controller
	recorder *MockNodeServiceMockRecorder
}

// MockNodeServiceMockRecorder is the mock recorder for MockNodeService
type MockNodeServiceMockRecorder struct {
	mock *MockNodeService
}

// NewMockNodeService creates a new mock instance
func NewMockNodeService(ctrl *gomock.Controller) *MockNodeService {
	mock := &MockNodeService{ctrl: ctrl}
	mock.recorder = &MockNodeServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNodeService) EXPECT() *MockNodeServiceMockRecorder {
	return m.recorder
}

// GetNodeByID mocks base method
func (m *MockNodeService) GetNodeByID(arg0 context.Context, arg1 string) (*vega.Node, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNodeByID", arg0, arg1)
	ret0, _ := ret[0].(*vega.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNodeByID indicates an expected call of GetNodeByID
func (mr *MockNodeServiceMockRecorder) GetNodeByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNodeByID", reflect.TypeOf((*MockNodeService)(nil).GetNodeByID), arg0, arg1)
}

// GetNodeData mocks base method
func (m *MockNodeService) GetNodeData(arg0 context.Context) (*vega.NodeData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNodeData", arg0)
	ret0, _ := ret[0].(*vega.NodeData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNodeData indicates an expected call of GetNodeData
func (mr *MockNodeServiceMockRecorder) GetNodeData(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNodeData", reflect.TypeOf((*MockNodeService)(nil).GetNodeData), arg0)
}

// GetNodes mocks base method
func (m *MockNodeService) GetNodes(arg0 context.Context) ([]*vega.Node, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNodes", arg0)
	ret0, _ := ret[0].([]*vega.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNodes indicates an expected call of GetNodes
func (mr *MockNodeServiceMockRecorder) GetNodes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNodes", reflect.TypeOf((*MockNodeService)(nil).GetNodes), arg0)
}
