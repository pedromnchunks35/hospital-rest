// Code generated by MockGen. DO NOT EDIT.
// Source: service/testMock.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRadar is a mock of Radar interface.
type MockRadar struct {
	ctrl     *gomock.Controller
	recorder *MockRadarMockRecorder
}

// MockRadarMockRecorder is the mock recorder for MockRadar.
type MockRadarMockRecorder struct {
	mock *MockRadar
}

// NewMockRadar creates a new mock instance.
func NewMockRadar(ctrl *gomock.Controller) *MockRadar {
	mock := &MockRadar{ctrl: ctrl}
	mock.recorder = &MockRadarMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRadar) EXPECT() *MockRadarMockRecorder {
	return m.recorder
}

// Scan mocks base method.
func (m *MockRadar) Scan(id int) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Scan", id)
	ret0, _ := ret[0].(string)
	return ret0
}

// Scan indicates an expected call of Scan.
func (mr *MockRadarMockRecorder) Scan(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scan", reflect.TypeOf((*MockRadar)(nil).Scan), id)
}
