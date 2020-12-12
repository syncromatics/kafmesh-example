// Code generated by MockGen. DO NOT EDIT.
// Source: ./customerDetailsSink.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	data "kafmesh-example/internal/data"
	reflect "reflect"
)

// MockcustomerRepo is a mock of customerRepo interface
type MockcustomerRepo struct {
	ctrl     *gomock.Controller
	recorder *MockcustomerRepoMockRecorder
}

// MockcustomerRepoMockRecorder is the mock recorder for MockcustomerRepo
type MockcustomerRepoMockRecorder struct {
	mock *MockcustomerRepo
}

// NewMockcustomerRepo creates a new mock instance
func NewMockcustomerRepo(ctrl *gomock.Controller) *MockcustomerRepo {
	mock := &MockcustomerRepo{ctrl: ctrl}
	mock.recorder = &MockcustomerRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockcustomerRepo) EXPECT() *MockcustomerRepoMockRecorder {
	return m.recorder
}

// GetCustomerDetails mocks base method
func (m *MockcustomerRepo) GetCustomerDetails(arg0 context.Context) ([]data.CustomerDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCustomerDetails", arg0)
	ret0, _ := ret[0].([]data.CustomerDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCustomerDetails indicates an expected call of GetCustomerDetails
func (mr *MockcustomerRepoMockRecorder) GetCustomerDetails(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCustomerDetails", reflect.TypeOf((*MockcustomerRepo)(nil).GetCustomerDetails), arg0)
}
