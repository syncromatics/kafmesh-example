// Code generated by MockGen. DO NOT EDIT.
// Source: /home/jeff/source/kafmesh-example/internal/definitions/details/enricher_processor.km.go

// Package details is a generated GoMock package.
package details

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

// Lookup_KafmeshCustomerIDDetails mocks base method
func (m *MockEnricher_ProcessorContext) Lookup_KafmeshCustomerIDDetails(key string) *customerId.Details {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Lookup_KafmeshCustomerIDDetails", key)
	ret0, _ := ret[0].(*customerId.Details)
	return ret0
}

// Lookup_KafmeshCustomerIDDetails indicates an expected call of Lookup_KafmeshCustomerIDDetails
func (mr *MockEnricher_ProcessorContextMockRecorder) Lookup_KafmeshCustomerIDDetails(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lookup_KafmeshCustomerIDDetails", reflect.TypeOf((*MockEnricher_ProcessorContext)(nil).Lookup_KafmeshCustomerIDDetails), key)
}

// Output_KafmeshDeviceIDEnrichedDetails mocks base method
func (m *MockEnricher_ProcessorContext) Output_KafmeshDeviceIDEnrichedDetails(key string, message *deviceId.EnrichedDetails) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Output_KafmeshDeviceIDEnrichedDetails", key, message)
}

// Output_KafmeshDeviceIDEnrichedDetails indicates an expected call of Output_KafmeshDeviceIDEnrichedDetails
func (mr *MockEnricher_ProcessorContextMockRecorder) Output_KafmeshDeviceIDEnrichedDetails(key, message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Output_KafmeshDeviceIDEnrichedDetails", reflect.TypeOf((*MockEnricher_ProcessorContext)(nil).Output_KafmeshDeviceIDEnrichedDetails), key, message)
}

// SaveState mocks base method
func (m *MockEnricher_ProcessorContext) SaveState(state *deviceId.EnrichedDetailsState) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SaveState", state)
}

// SaveState indicates an expected call of SaveState
func (mr *MockEnricher_ProcessorContextMockRecorder) SaveState(state interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveState", reflect.TypeOf((*MockEnricher_ProcessorContext)(nil).SaveState), state)
}

// State mocks base method
func (m *MockEnricher_ProcessorContext) State() *deviceId.EnrichedDetailsState {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "State")
	ret0, _ := ret[0].(*deviceId.EnrichedDetailsState)
	return ret0
}

// State indicates an expected call of State
func (mr *MockEnricher_ProcessorContextMockRecorder) State() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "State", reflect.TypeOf((*MockEnricher_ProcessorContext)(nil).State))
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

// HandleKafmeshDeviceIDDetails mocks base method
func (m *MockEnricher_Processor) HandleKafmeshDeviceIDDetails(ctx Enricher_ProcessorContext, message *deviceId.Details) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleKafmeshDeviceIDDetails", ctx, message)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleKafmeshDeviceIDDetails indicates an expected call of HandleKafmeshDeviceIDDetails
func (mr *MockEnricher_ProcessorMockRecorder) HandleKafmeshDeviceIDDetails(ctx, message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleKafmeshDeviceIDDetails", reflect.TypeOf((*MockEnricher_Processor)(nil).HandleKafmeshDeviceIDDetails), ctx, message)
}

// HandleKafmeshDeviceIDCustomer mocks base method
func (m *MockEnricher_Processor) HandleKafmeshDeviceIDCustomer(ctx Enricher_ProcessorContext, message *deviceId.Customer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleKafmeshDeviceIDCustomer", ctx, message)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleKafmeshDeviceIDCustomer indicates an expected call of HandleKafmeshDeviceIDCustomer
func (mr *MockEnricher_ProcessorMockRecorder) HandleKafmeshDeviceIDCustomer(ctx, message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleKafmeshDeviceIDCustomer", reflect.TypeOf((*MockEnricher_Processor)(nil).HandleKafmeshDeviceIDCustomer), ctx, message)
}
