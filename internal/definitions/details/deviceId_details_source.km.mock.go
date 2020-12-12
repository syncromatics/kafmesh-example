// Code generated by MockGen. DO NOT EDIT.
// Source: /home/jeff/source/syncromatics/kafmesh-example/internal/definitions/details/deviceId_details_source.km.go

// Package details is a generated GoMock package.
package details

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockDeviceIDDetails_Source is a mock of DeviceIDDetails_Source interface
type MockDeviceIDDetails_Source struct {
	ctrl     *gomock.Controller
	recorder *MockDeviceIDDetails_SourceMockRecorder
}

// MockDeviceIDDetails_SourceMockRecorder is the mock recorder for MockDeviceIDDetails_Source
type MockDeviceIDDetails_SourceMockRecorder struct {
	mock *MockDeviceIDDetails_Source
}

// NewMockDeviceIDDetails_Source creates a new mock instance
func NewMockDeviceIDDetails_Source(ctrl *gomock.Controller) *MockDeviceIDDetails_Source {
	mock := &MockDeviceIDDetails_Source{ctrl: ctrl}
	mock.recorder = &MockDeviceIDDetails_SourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDeviceIDDetails_Source) EXPECT() *MockDeviceIDDetails_SourceMockRecorder {
	return m.recorder
}

// Emit mocks base method
func (m *MockDeviceIDDetails_Source) Emit(message DeviceIDDetails_Source_Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Emit", message)
	ret0, _ := ret[0].(error)
	return ret0
}

// Emit indicates an expected call of Emit
func (mr *MockDeviceIDDetails_SourceMockRecorder) Emit(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Emit", reflect.TypeOf((*MockDeviceIDDetails_Source)(nil).Emit), message)
}

// EmitBulk mocks base method
func (m *MockDeviceIDDetails_Source) EmitBulk(ctx context.Context, messages []DeviceIDDetails_Source_Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EmitBulk", ctx, messages)
	ret0, _ := ret[0].(error)
	return ret0
}

// EmitBulk indicates an expected call of EmitBulk
func (mr *MockDeviceIDDetails_SourceMockRecorder) EmitBulk(ctx, messages interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EmitBulk", reflect.TypeOf((*MockDeviceIDDetails_Source)(nil).EmitBulk), ctx, messages)
}

// Delete mocks base method
func (m *MockDeviceIDDetails_Source) Delete(key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockDeviceIDDetails_SourceMockRecorder) Delete(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDeviceIDDetails_Source)(nil).Delete), key)
}
