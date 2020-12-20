// Code generated by MockGen. DO NOT EDIT.
// Source: /home/jeff/source/kafmesh-example/internal/definitions/assignments/customerId_details_view.km.go

// Package assignments is a generated GoMock package.
package assignments

import (
	gomock "github.com/golang/mock/gomock"
	customerId "kafmesh-example/internal/definitions/models/kafmesh/customerId"
	reflect "reflect"
)

// MockCustomerIDDetails_View is a mock of CustomerIDDetails_View interface
type MockCustomerIDDetails_View struct {
	ctrl     *gomock.Controller
	recorder *MockCustomerIDDetails_ViewMockRecorder
}

// MockCustomerIDDetails_ViewMockRecorder is the mock recorder for MockCustomerIDDetails_View
type MockCustomerIDDetails_ViewMockRecorder struct {
	mock *MockCustomerIDDetails_View
}

// NewMockCustomerIDDetails_View creates a new mock instance
func NewMockCustomerIDDetails_View(ctrl *gomock.Controller) *MockCustomerIDDetails_View {
	mock := &MockCustomerIDDetails_View{ctrl: ctrl}
	mock.recorder = &MockCustomerIDDetails_ViewMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCustomerIDDetails_View) EXPECT() *MockCustomerIDDetails_ViewMockRecorder {
	return m.recorder
}

// Keys mocks base method
func (m *MockCustomerIDDetails_View) Keys() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Keys indicates an expected call of Keys
func (mr *MockCustomerIDDetails_ViewMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockCustomerIDDetails_View)(nil).Keys))
}

// Get mocks base method
func (m *MockCustomerIDDetails_View) Get(key string) (*customerId.Details, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].(*customerId.Details)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockCustomerIDDetails_ViewMockRecorder) Get(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCustomerIDDetails_View)(nil).Get), key)
}