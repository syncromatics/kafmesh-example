// Code generated by MockGen. DO NOT EDIT.
// Source: /home/jeff/source/kafmesh-example/internal/definitions/assignments/deviceId_customer_view.km.go

// Package assignments is a generated GoMock package.
package assignments

import (
	gomock "github.com/golang/mock/gomock"
	deviceId "kafmesh-example/internal/definitions/models/kafmesh/deviceId"
	reflect "reflect"
)

// MockDeviceIDCustomer_View is a mock of DeviceIDCustomer_View interface
type MockDeviceIDCustomer_View struct {
	ctrl     *gomock.Controller
	recorder *MockDeviceIDCustomer_ViewMockRecorder
}

// MockDeviceIDCustomer_ViewMockRecorder is the mock recorder for MockDeviceIDCustomer_View
type MockDeviceIDCustomer_ViewMockRecorder struct {
	mock *MockDeviceIDCustomer_View
}

// NewMockDeviceIDCustomer_View creates a new mock instance
func NewMockDeviceIDCustomer_View(ctrl *gomock.Controller) *MockDeviceIDCustomer_View {
	mock := &MockDeviceIDCustomer_View{ctrl: ctrl}
	mock.recorder = &MockDeviceIDCustomer_ViewMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDeviceIDCustomer_View) EXPECT() *MockDeviceIDCustomer_ViewMockRecorder {
	return m.recorder
}

// Keys mocks base method
func (m *MockDeviceIDCustomer_View) Keys() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Keys indicates an expected call of Keys
func (mr *MockDeviceIDCustomer_ViewMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockDeviceIDCustomer_View)(nil).Keys))
}

// Get mocks base method
func (m *MockDeviceIDCustomer_View) Get(key string) (*deviceId.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].(*deviceId.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockDeviceIDCustomer_ViewMockRecorder) Get(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDeviceIDCustomer_View)(nil).Get), key)
}