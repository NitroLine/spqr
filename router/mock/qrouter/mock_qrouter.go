// Code generated by MockGen. DO NOT EDIT.
// Source: ./router/qrouter/qrouter.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	lyx "github.com/pg-sharding/lyx/lyx"
	meta "github.com/pg-sharding/spqr/pkg/meta"
	kr "github.com/pg-sharding/spqr/pkg/models/kr"
	session "github.com/pg-sharding/spqr/pkg/session"
	routingstate "github.com/pg-sharding/spqr/router/routingstate"
)

// MockQueryRouter is a mock of QueryRouter interface.
type MockQueryRouter struct {
	ctrl     *gomock.Controller
	recorder *MockQueryRouterMockRecorder
}

// MockQueryRouterMockRecorder is the mock recorder for MockQueryRouter.
type MockQueryRouterMockRecorder struct {
	mock *MockQueryRouter
}

// NewMockQueryRouter creates a new mock instance.
func NewMockQueryRouter(ctrl *gomock.Controller) *MockQueryRouter {
	mock := &MockQueryRouter{ctrl: ctrl}
	mock.recorder = &MockQueryRouterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQueryRouter) EXPECT() *MockQueryRouterMockRecorder {
	return m.recorder
}

// DataShardsRoutes mocks base method.
func (m *MockQueryRouter) DataShardsRoutes() []*routingstate.DataShardRoute {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DataShardsRoutes")
	ret0, _ := ret[0].([]*routingstate.DataShardRoute)
	return ret0
}

// DataShardsRoutes indicates an expected call of DataShardsRoutes.
func (mr *MockQueryRouterMockRecorder) DataShardsRoutes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataShardsRoutes", reflect.TypeOf((*MockQueryRouter)(nil).DataShardsRoutes))
}

// DeparseKeyWithRangesInternal mocks base method.
func (m *MockQueryRouter) DeparseKeyWithRangesInternal(ctx context.Context, key string, krs []*kr.KeyRange) (*routingstate.DataShardRoute, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeparseKeyWithRangesInternal", ctx, key, krs)
	ret0, _ := ret[0].(*routingstate.DataShardRoute)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeparseKeyWithRangesInternal indicates an expected call of DeparseKeyWithRangesInternal.
func (mr *MockQueryRouterMockRecorder) DeparseKeyWithRangesInternal(ctx, key, krs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeparseKeyWithRangesInternal", reflect.TypeOf((*MockQueryRouter)(nil).DeparseKeyWithRangesInternal), ctx, key, krs)
}

// Initialize mocks base method.
func (m *MockQueryRouter) Initialize() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Initialize")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Initialize indicates an expected call of Initialize.
func (mr *MockQueryRouterMockRecorder) Initialize() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialize", reflect.TypeOf((*MockQueryRouter)(nil).Initialize))
}

// Initialized mocks base method.
func (m *MockQueryRouter) Initialized() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Initialized")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Initialized indicates an expected call of Initialized.
func (mr *MockQueryRouterMockRecorder) Initialized() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialized", reflect.TypeOf((*MockQueryRouter)(nil).Initialized))
}

// Mgr mocks base method.
func (m *MockQueryRouter) Mgr() meta.EntityMgr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Mgr")
	ret0, _ := ret[0].(meta.EntityMgr)
	return ret0
}

// Mgr indicates an expected call of Mgr.
func (mr *MockQueryRouterMockRecorder) Mgr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mgr", reflect.TypeOf((*MockQueryRouter)(nil).Mgr))
}

// Route mocks base method.
func (m *MockQueryRouter) Route(ctx context.Context, stmt lyx.Node, sph session.SessionParamsHolder) (routingstate.RoutingState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Route", ctx, stmt, sph)
	ret0, _ := ret[0].(routingstate.RoutingState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Route indicates an expected call of Route.
func (mr *MockQueryRouterMockRecorder) Route(ctx, stmt, sph interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Route", reflect.TypeOf((*MockQueryRouter)(nil).Route), ctx, stmt, sph)
}

// WorldShardsRoutes mocks base method.
func (m *MockQueryRouter) WorldShardsRoutes() []*routingstate.DataShardRoute {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WorldShardsRoutes")
	ret0, _ := ret[0].([]*routingstate.DataShardRoute)
	return ret0
}

// WorldShardsRoutes indicates an expected call of WorldShardsRoutes.
func (mr *MockQueryRouterMockRecorder) WorldShardsRoutes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WorldShardsRoutes", reflect.TypeOf((*MockQueryRouter)(nil).WorldShardsRoutes))
}
