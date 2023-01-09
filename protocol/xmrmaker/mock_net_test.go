// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/athanorlabs/atomic-swap/net (interfaces: NetHost)

// Package xmrmaker is a generated GoMock package.
package xmrmaker

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	network "github.com/libp2p/go-libp2p/core/network"
	peer "github.com/libp2p/go-libp2p/core/peer"
	protocol "github.com/libp2p/go-libp2p/core/protocol"
)

// MockNetHost is a mock of NetHost interface.
type MockNetHost struct {
	ctrl     *gomock.Controller
	recorder *MockNetHostMockRecorder
}

// MockNetHostMockRecorder is the mock recorder for MockNetHost.
type MockNetHostMockRecorder struct {
	mock *MockNetHost
}

// NewMockNetHost creates a new mock instance.
func NewMockNetHost(ctrl *gomock.Controller) *MockNetHost {
	mock := &MockNetHost{ctrl: ctrl}
	mock.recorder = &MockNetHostMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNetHost) EXPECT() *MockNetHostMockRecorder {
	return m.recorder
}

// AddrInfo mocks base method.
func (m *MockNetHost) AddrInfo() peer.AddrInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddrInfo")
	ret0, _ := ret[0].(peer.AddrInfo)
	return ret0
}

// AddrInfo indicates an expected call of AddrInfo.
func (mr *MockNetHostMockRecorder) AddrInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddrInfo", reflect.TypeOf((*MockNetHost)(nil).AddrInfo))
}

// Addresses mocks base method.
func (m *MockNetHost) Addresses() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Addresses")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Addresses indicates an expected call of Addresses.
func (mr *MockNetHostMockRecorder) Addresses() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Addresses", reflect.TypeOf((*MockNetHost)(nil).Addresses))
}

// Advertise mocks base method.
func (m *MockNetHost) Advertise(arg0 []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Advertise", arg0)
}

// Advertise indicates an expected call of Advertise.
func (mr *MockNetHostMockRecorder) Advertise(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Advertise", reflect.TypeOf((*MockNetHost)(nil).Advertise), arg0)
}

// Connect mocks base method.
func (m *MockNetHost) Connect(arg0 context.Context, arg1 peer.AddrInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Connect indicates an expected call of Connect.
func (mr *MockNetHostMockRecorder) Connect(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockNetHost)(nil).Connect), arg0, arg1)
}

// ConnectedPeers mocks base method.
func (m *MockNetHost) ConnectedPeers() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectedPeers")
	ret0, _ := ret[0].([]string)
	return ret0
}

// ConnectedPeers indicates an expected call of ConnectedPeers.
func (mr *MockNetHostMockRecorder) ConnectedPeers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectedPeers", reflect.TypeOf((*MockNetHost)(nil).ConnectedPeers))
}

// Connectedness mocks base method.
func (m *MockNetHost) Connectedness(arg0 peer.ID) network.Connectedness {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connectedness", arg0)
	ret0, _ := ret[0].(network.Connectedness)
	return ret0
}

// Connectedness indicates an expected call of Connectedness.
func (mr *MockNetHostMockRecorder) Connectedness(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connectedness", reflect.TypeOf((*MockNetHost)(nil).Connectedness), arg0)
}

// Discover mocks base method.
func (m *MockNetHost) Discover(arg0 string, arg1 time.Duration) ([]peer.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Discover", arg0, arg1)
	ret0, _ := ret[0].([]peer.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Discover indicates an expected call of Discover.
func (mr *MockNetHostMockRecorder) Discover(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Discover", reflect.TypeOf((*MockNetHost)(nil).Discover), arg0, arg1)
}

// NewStream mocks base method.
func (m *MockNetHost) NewStream(arg0 context.Context, arg1 peer.ID, arg2 protocol.ID) (network.Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewStream", arg0, arg1, arg2)
	ret0, _ := ret[0].(network.Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewStream indicates an expected call of NewStream.
func (mr *MockNetHostMockRecorder) NewStream(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewStream", reflect.TypeOf((*MockNetHost)(nil).NewStream), arg0, arg1, arg2)
}

// PeerID mocks base method.
func (m *MockNetHost) PeerID() peer.ID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PeerID")
	ret0, _ := ret[0].(peer.ID)
	return ret0
}

// PeerID indicates an expected call of PeerID.
func (mr *MockNetHostMockRecorder) PeerID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeerID", reflect.TypeOf((*MockNetHost)(nil).PeerID))
}

// SetShouldAdvertiseFunc mocks base method.
func (m *MockNetHost) SetShouldAdvertiseFunc(arg0 func() bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetShouldAdvertiseFunc", arg0)
}

// SetShouldAdvertiseFunc indicates an expected call of SetShouldAdvertiseFunc.
func (mr *MockNetHostMockRecorder) SetShouldAdvertiseFunc(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetShouldAdvertiseFunc", reflect.TypeOf((*MockNetHost)(nil).SetShouldAdvertiseFunc), arg0)
}

// SetStreamHandler mocks base method.
func (m *MockNetHost) SetStreamHandler(arg0 string, arg1 func(network.Stream)) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetStreamHandler", arg0, arg1)
}

// SetStreamHandler indicates an expected call of SetStreamHandler.
func (mr *MockNetHostMockRecorder) SetStreamHandler(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStreamHandler", reflect.TypeOf((*MockNetHost)(nil).SetStreamHandler), arg0, arg1)
}

// Start mocks base method.
func (m *MockNetHost) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockNetHostMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockNetHost)(nil).Start))
}

// Stop mocks base method.
func (m *MockNetHost) Stop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockNetHostMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockNetHost)(nil).Stop))
}
