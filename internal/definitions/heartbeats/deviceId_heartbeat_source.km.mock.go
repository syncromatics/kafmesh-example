// Code generated by MockGen. DO NOT EDIT.
// Source: /home/jeff/source/syncromatics/kafmesh-example/internal/definitions/heartbeats/deviceId_heartbeat_source.km.go

// Package heartbeats is a generated GoMock package.
package heartbeats

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockDeviceIDHeartbeat_Source is a mock of DeviceIDHeartbeat_Source interface
type MockDeviceIDHeartbeat_Source struct {
	ctrl     *gomock.Controller
	recorder *MockDeviceIDHeartbeat_SourceMockRecorder
}

// MockDeviceIDHeartbeat_SourceMockRecorder is the mock recorder for MockDeviceIDHeartbeat_Source
type MockDeviceIDHeartbeat_SourceMockRecorder struct {
	mock *MockDeviceIDHeartbeat_Source
}

// NewMockDeviceIDHeartbeat_Source creates a new mock instance
func NewMockDeviceIDHeartbeat_Source(ctrl *gomock.Controller) *MockDeviceIDHeartbeat_Source {
	mock := &MockDeviceIDHeartbeat_Source{ctrl: ctrl}
	mock.recorder = &MockDeviceIDHeartbeat_SourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDeviceIDHeartbeat_Source) EXPECT() *MockDeviceIDHeartbeat_SourceMockRecorder {
	return m.recorder
}

// Emit mocks base method
func (m *MockDeviceIDHeartbeat_Source) Emit(message DeviceIDHeartbeat_Source_Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Emit", message)
	ret0, _ := ret[0].(error)
	return ret0
}

// Emit indicates an expected call of Emit
func (mr *MockDeviceIDHeartbeat_SourceMockRecorder) Emit(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Emit", reflect.TypeOf((*MockDeviceIDHeartbeat_Source)(nil).Emit), message)
}

// EmitBulk mocks base method
func (m *MockDeviceIDHeartbeat_Source) EmitBulk(ctx context.Context, messages []DeviceIDHeartbeat_Source_Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EmitBulk", ctx, messages)
	ret0, _ := ret[0].(error)
	return ret0
}

// EmitBulk indicates an expected call of EmitBulk
func (mr *MockDeviceIDHeartbeat_SourceMockRecorder) EmitBulk(ctx, messages interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EmitBulk", reflect.TypeOf((*MockDeviceIDHeartbeat_Source)(nil).EmitBulk), ctx, messages)
}

// Delete mocks base method
func (m *MockDeviceIDHeartbeat_Source) Delete(key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockDeviceIDHeartbeat_SourceMockRecorder) Delete(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDeviceIDHeartbeat_Source)(nil).Delete), key)
}
