// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/libopenstorage/openstorage/cluster (interfaces: Cluster)

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	api "github.com/libopenstorage/openstorage/api"
	cluster "github.com/libopenstorage/openstorage/cluster"
	reflect "reflect"
	time "time"
)

// MockCluster is a mock of Cluster interface
type MockCluster struct {
	ctrl     *gomock.Controller
	recorder *MockClusterMockRecorder
}

// MockClusterMockRecorder is the mock recorder for MockCluster
type MockClusterMockRecorder struct {
	mock *MockCluster
}

// NewMockCluster creates a new mock instance
func NewMockCluster(ctrl *gomock.Controller) *MockCluster {
	mock := &MockCluster{ctrl: ctrl}
	mock.recorder = &MockClusterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCluster) EXPECT() *MockClusterMockRecorder {
	return m.recorder
}

// AddEventListener mocks base method
func (m *MockCluster) AddEventListener(arg0 cluster.ClusterListener) error {
	ret := m.ctrl.Call(m, "AddEventListener", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEventListener indicates an expected call of AddEventListener
func (mr *MockClusterMockRecorder) AddEventListener(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventListener", reflect.TypeOf((*MockCluster)(nil).AddEventListener), arg0)
}

// ClearAlert mocks base method
func (m *MockCluster) ClearAlert(arg0 api.ResourceType, arg1 int64) error {
	ret := m.ctrl.Call(m, "ClearAlert", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClearAlert indicates an expected call of ClearAlert
func (mr *MockClusterMockRecorder) ClearAlert(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearAlert", reflect.TypeOf((*MockCluster)(nil).ClearAlert), arg0, arg1)
}

// DisableUpdates mocks base method
func (m *MockCluster) DisableUpdates() error {
	ret := m.ctrl.Call(m, "DisableUpdates")
	ret0, _ := ret[0].(error)
	return ret0
}

// DisableUpdates indicates an expected call of DisableUpdates
func (mr *MockClusterMockRecorder) DisableUpdates() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableUpdates", reflect.TypeOf((*MockCluster)(nil).DisableUpdates))
}

// EnableUpdates mocks base method
func (m *MockCluster) EnableUpdates() error {
	ret := m.ctrl.Call(m, "EnableUpdates")
	ret0, _ := ret[0].(error)
	return ret0
}

// EnableUpdates indicates an expected call of EnableUpdates
func (mr *MockClusterMockRecorder) EnableUpdates() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableUpdates", reflect.TypeOf((*MockCluster)(nil).EnableUpdates))
}

// Enumerate mocks base method
func (m *MockCluster) Enumerate() (api.Cluster, error) {
	ret := m.ctrl.Call(m, "Enumerate")
	ret0, _ := ret[0].(api.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Enumerate indicates an expected call of Enumerate
func (mr *MockClusterMockRecorder) Enumerate() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enumerate", reflect.TypeOf((*MockCluster)(nil).Enumerate))
}

// EnumerateAlerts mocks base method
func (m *MockCluster) EnumerateAlerts(arg0, arg1 time.Time, arg2 api.ResourceType) (*api.Alerts, error) {
	ret := m.ctrl.Call(m, "EnumerateAlerts", arg0, arg1, arg2)
	ret0, _ := ret[0].(*api.Alerts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnumerateAlerts indicates an expected call of EnumerateAlerts
func (mr *MockClusterMockRecorder) EnumerateAlerts(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnumerateAlerts", reflect.TypeOf((*MockCluster)(nil).EnumerateAlerts), arg0, arg1, arg2)
}

// EraseAlert mocks base method
func (m *MockCluster) EraseAlert(arg0 api.ResourceType, arg1 int64) error {
	ret := m.ctrl.Call(m, "EraseAlert", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EraseAlert indicates an expected call of EraseAlert
func (mr *MockClusterMockRecorder) EraseAlert(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EraseAlert", reflect.TypeOf((*MockCluster)(nil).EraseAlert), arg0, arg1)
}

// GetData mocks base method
func (m *MockCluster) GetData() (map[string]*api.Node, error) {
	ret := m.ctrl.Call(m, "GetData")
	ret0, _ := ret[0].(map[string]*api.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetData indicates an expected call of GetData
func (mr *MockClusterMockRecorder) GetData() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetData", reflect.TypeOf((*MockCluster)(nil).GetData))
}

// GetGossipState mocks base method
func (m *MockCluster) GetGossipState() *cluster.ClusterState {
	ret := m.ctrl.Call(m, "GetGossipState")
	ret0, _ := ret[0].(*cluster.ClusterState)
	return ret0
}

// GetGossipState indicates an expected call of GetGossipState
func (mr *MockClusterMockRecorder) GetGossipState() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGossipState", reflect.TypeOf((*MockCluster)(nil).GetGossipState))
}

// GetTunnelConfig mocks base method
func (m *MockCluster) GetTunnelConfig() api.TunnelConfig {
	ret := m.ctrl.Call(m, "GetTunnelConfig")
	ret0, _ := ret[0].(api.TunnelConfig)
	return ret0
}

// GetTunnelConfig indicates an expected call of GetTunnelConfig
func (mr *MockClusterMockRecorder) GetTunnelConfig() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTunnelConfig", reflect.TypeOf((*MockCluster)(nil).GetTunnelConfig))
}

// Inspect mocks base method
func (m *MockCluster) Inspect(arg0 string) (api.Node, error) {
	ret := m.ctrl.Call(m, "Inspect", arg0)
	ret0, _ := ret[0].(api.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Inspect indicates an expected call of Inspect
func (mr *MockClusterMockRecorder) Inspect(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Inspect", reflect.TypeOf((*MockCluster)(nil).Inspect), arg0)
}

// NodeRemoveDone mocks base method
func (m *MockCluster) NodeRemoveDone(arg0 string, arg1 error) {
	m.ctrl.Call(m, "NodeRemoveDone", arg0, arg1)
}

// NodeRemoveDone indicates an expected call of NodeRemoveDone
func (mr *MockClusterMockRecorder) NodeRemoveDone(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeRemoveDone", reflect.TypeOf((*MockCluster)(nil).NodeRemoveDone), arg0, arg1)
}

// NodeStatus mocks base method
func (m *MockCluster) NodeStatus() (api.Status, error) {
	ret := m.ctrl.Call(m, "NodeStatus")
	ret0, _ := ret[0].(api.Status)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NodeStatus indicates an expected call of NodeStatus
func (mr *MockClusterMockRecorder) NodeStatus() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeStatus", reflect.TypeOf((*MockCluster)(nil).NodeStatus))
}

// PeerStatus mocks base method
func (m *MockCluster) PeerStatus(arg0 string) (map[string]api.Status, error) {
	ret := m.ctrl.Call(m, "PeerStatus", arg0)
	ret0, _ := ret[0].(map[string]api.Status)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PeerStatus indicates an expected call of PeerStatus
func (mr *MockClusterMockRecorder) PeerStatus(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeerStatus", reflect.TypeOf((*MockCluster)(nil).PeerStatus), arg0)
}

// Remove mocks base method
func (m *MockCluster) Remove(arg0 []api.Node, arg1 bool) error {
	ret := m.ctrl.Call(m, "Remove", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove
func (mr *MockClusterMockRecorder) Remove(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockCluster)(nil).Remove), arg0, arg1)
}

// SetFluentDConfig mocks base method
func (m *MockCluster) SetFluentDConfig(arg0 api.FluentDConfig) error {
	ret := m.ctrl.Call(m, "SetFluentDConfig", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetFluentDConfig indicates an expected call of SetFluentDConfig
func (mr *MockClusterMockRecorder) SetFluentDConfig(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFluentDConfig", reflect.TypeOf((*MockCluster)(nil).SetFluentDConfig), arg0)
}

// SetLoggingURL mocks base method
func (m *MockCluster) SetLoggingURL(arg0 string) error {
	ret := m.ctrl.Call(m, "SetLoggingURL", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetLoggingURL indicates an expected call of SetLoggingURL
func (mr *MockClusterMockRecorder) SetLoggingURL(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLoggingURL", reflect.TypeOf((*MockCluster)(nil).SetLoggingURL), arg0)
}

// SetManagementURL mocks base method
func (m *MockCluster) SetManagementURL(arg0 string) error {
	ret := m.ctrl.Call(m, "SetManagementURL", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetManagementURL indicates an expected call of SetManagementURL
func (mr *MockClusterMockRecorder) SetManagementURL(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetManagementURL", reflect.TypeOf((*MockCluster)(nil).SetManagementURL), arg0)
}

// SetSize mocks base method
func (m *MockCluster) SetSize(arg0 int) error {
	ret := m.ctrl.Call(m, "SetSize", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetSize indicates an expected call of SetSize
func (mr *MockClusterMockRecorder) SetSize(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSize", reflect.TypeOf((*MockCluster)(nil).SetSize), arg0)
}

// SetTunnelConfig mocks base method
func (m *MockCluster) SetTunnelConfig(arg0 api.TunnelConfig) error {
	ret := m.ctrl.Call(m, "SetTunnelConfig", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetTunnelConfig indicates an expected call of SetTunnelConfig
func (mr *MockClusterMockRecorder) SetTunnelConfig(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTunnelConfig", reflect.TypeOf((*MockCluster)(nil).SetTunnelConfig), arg0)
}

// Shutdown mocks base method
func (m *MockCluster) Shutdown() error {
	ret := m.ctrl.Call(m, "Shutdown")
	ret0, _ := ret[0].(error)
	return ret0
}

// Shutdown indicates an expected call of Shutdown
func (mr *MockClusterMockRecorder) Shutdown() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockCluster)(nil).Shutdown))
}

// Start mocks base method
func (m *MockCluster) Start(arg0 int, arg1 bool) error {
	ret := m.ctrl.Call(m, "Start", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockClusterMockRecorder) Start(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockCluster)(nil).Start), arg0, arg1)
}

// UpdateData mocks base method
func (m *MockCluster) UpdateData(arg0 map[string]interface{}) error {
	ret := m.ctrl.Call(m, "UpdateData", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateData indicates an expected call of UpdateData
func (mr *MockClusterMockRecorder) UpdateData(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateData", reflect.TypeOf((*MockCluster)(nil).UpdateData), arg0)
}

// UpdateLabels mocks base method
func (m *MockCluster) UpdateLabels(arg0 map[string]string) error {
	ret := m.ctrl.Call(m, "UpdateLabels", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateLabels indicates an expected call of UpdateLabels
func (mr *MockClusterMockRecorder) UpdateLabels(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLabels", reflect.TypeOf((*MockCluster)(nil).UpdateLabels), arg0)
}