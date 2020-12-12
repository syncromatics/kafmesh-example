// Code generated by MockGen. DO NOT EDIT.
// Source: /home/jeff/source/syncromatics/kafmesh-example/internal/definitions/heartbeats/enricher_processor.km.go

// Package heartbeats is a generated GoMock package.
package heartbeats

import (
	gomock "github.com/golang/mock/gomock"
	customerId "kafmesh-example/internal/definitions/models/kafmesh/customerId"
	deviceId "kafmesh-example/internal/definitions/models/kafmesh/deviceId"
	reflect "reflect"
	time "time"
)

// MockEnricher_ProcessorContext is a mock of Enricher_ProcessorContext interface
type MockEnricher_ProcessorContext struct {
	ctrl     *gomock.Controller
	recorder *MockEnricher_ProcessorContextMockRecorder
}

// MockEnricher_ProcessorContextMockRecorder is the mock recorder for MockEnricher_ProcessorContext
type MockEnricher_ProcessorContextMockRecorder struct {
	mock *MockEnricher_ProcessorContext
}

// NewMockEnricher_ProcessorContext creates a new mock instance
func NewMockEnricher_ProcessorContext(ctrl *gomock.Controller) *MockEnricher_ProcessorContext {
	mock := &MockEnricher_ProcessorContext{ctrl: ctrl}
	mock.recorder = &MockEnricher_ProcessorContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEnricher_ProcessorContext) EXPECT() *MockEnricher_ProcessorContextMockRecorder {
	return m.recorder
}

// Key mocks base method
func (m *MockEnricher_ProcessorContext) Key() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Key")
	ret0, _ := ret[0].(string)
	return ret0
}

// Key indicates an expected call of Key
func (mr *MockEnricher_ProcessorContextMockRecorder) Key() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Key", reflect.TypeOf((*MockEnricher_ProcessorContext)(nil).Key))
}

// Timestamp mocks base method
func (m *MockEnricher_ProcessorContext) Timestamp() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Timestamp")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// Timestamp indicates an expected call of Timestamp
func (mr *MockEnricher_ProcessorContextMockRecorder) Timestamp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Timestamp", reflect.TypeOf((*MockEnricher_ProcessorContext)(nil).Timestamp))
}

// Lookup_CustomerIDDetails mocks base method
func (m *MockEnricher_ProcessorContext) Lookup_CustomerIDDetails(key string) *customerId.Details {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Lookup_CustomerIDDetails", key)
	ret0, _ := ret[0].(*customerId.Details)
	return ret0
}

// Lookup_CustomerIDDetails indicates an expected call of Lookup_CustomerIDDetails
func (mr *MockEnricher_ProcessorContextMockRecorder) Lookup_CustomerIDDetails(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lookup_CustomerIDDetails", reflect.TypeOf((*MockEnricher_ProcessorContext)(nil).Lookup_CustomerIDDetails), key)
}

// Join_DeviceIDCustomer mocks base method
func (m *MockEnricher_ProcessorContext) Join_DeviceIDCustomer() *deviceId.Customer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Join_DeviceIDCustomer")
	ret0, _ := ret[0].(*deviceId.Customer)
	return ret0
}

// Join_DeviceIDCustomer indicates an expected call of Join_DeviceIDCustomer
func (mr *MockEnricher_ProcessorContextMockRecorder) Join_DeviceIDCustomer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Join_DeviceIDCustomer", reflect.TypeOf((*MockEnricher_ProcessorContext)(nil).Join_DeviceIDCustomer))
}

// Output_DeviceIDEnrichedHeartbeat mocks base method
func (m *MockEnricher_ProcessorContext) Output_DeviceIDEnrichedHeartbeat(key string, message *deviceId.EnrichedHeartbeat) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Output_DeviceIDEnrichedHeartbeat", key, message)
}

// Output_DeviceIDEnrichedHeartbeat indicates an expected call of Output_DeviceIDEnrichedHeartbeat
func (mr *MockEnricher_ProcessorContextMockRecorder) Output_DeviceIDEnrichedHeartbeat(key, message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Output_DeviceIDEnrichedHeartbeat", reflect.TypeOf((*MockEnricher_ProcessorContext)(nil).Output_DeviceIDEnrichedHeartbeat), key, message)
}

// MockEnricher_Processor is a mock of Enricher_Processor interface
type MockEnricher_Processor struct {
	ctrl     *gomock.Controller
	recorder *MockEnricher_ProcessorMockRecorder
}

// MockEnricher_ProcessorMockRecorder is the mock recorder for MockEnricher_Processor
type MockEnricher_ProcessorMockRecorder struct {
	mock *MockEnricher_Processor
}

// NewMockEnricher_Processor creates a new mock instance
func NewMockEnricher_Processor(ctrl *gomock.Controller) *MockEnricher_Processor {
	mock := &MockEnricher_Processor{ctrl: ctrl}
	mock.recorder = &MockEnricher_ProcessorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEnricher_Processor) EXPECT() *MockEnricher_ProcessorMockRecorder {
	return m.recorder
}

// HandleDeviceIDHeartbeat mocks base method
func (m *MockEnricher_Processor) HandleDeviceIDHeartbeat(ctx Enricher_ProcessorContext, message *deviceId.Heartbeat) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleDeviceIDHeartbeat", ctx, message)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleDeviceIDHeartbeat indicates an expected call of HandleDeviceIDHeartbeat
func (mr *MockEnricher_ProcessorMockRecorder) HandleDeviceIDHeartbeat(ctx, message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleDeviceIDHeartbeat", reflect.TypeOf((*MockEnricher_Processor)(nil).HandleDeviceIDHeartbeat), ctx, message)
}
