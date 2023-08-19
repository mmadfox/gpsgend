// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/mmadfox/projects/gpsio/gpsgend/internal/generator/processes.go

// Package mock_generator is a generated GoMock package.
package mock_generator

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	go_gpsgen "github.com/mmadfox/go-gpsgen"
)

// MockProcesses is a mock of Processes interface.
type MockProcesses struct {
	ctrl     *gomock.Controller
	recorder *MockProcessesMockRecorder
}

// MockProcessesMockRecorder is the mock recorder for MockProcesses.
type MockProcessesMockRecorder struct {
	mock *MockProcesses
}

// NewMockProcesses creates a new mock instance.
func NewMockProcesses(ctrl *gomock.Controller) *MockProcesses {
	mock := &MockProcesses{ctrl: ctrl}
	mock.recorder = &MockProcessesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProcesses) EXPECT() *MockProcessesMockRecorder {
	return m.recorder
}

// Attach mocks base method.
func (m *MockProcesses) Attach(d *go_gpsgen.Device) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Attach", d)
	ret0, _ := ret[0].(error)
	return ret0
}

// Attach indicates an expected call of Attach.
func (mr *MockProcessesMockRecorder) Attach(d interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Attach", reflect.TypeOf((*MockProcesses)(nil).Attach), d)
}

// Detach mocks base method.
func (m *MockProcesses) Detach(deviceID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Detach", deviceID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Detach indicates an expected call of Detach.
func (mr *MockProcessesMockRecorder) Detach(deviceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Detach", reflect.TypeOf((*MockProcesses)(nil).Detach), deviceID)
}

// HasTracker mocks base method.
func (m *MockProcesses) HasTracker(deviceID string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasTracker", deviceID)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasTracker indicates an expected call of HasTracker.
func (mr *MockProcessesMockRecorder) HasTracker(deviceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasTracker", reflect.TypeOf((*MockProcesses)(nil).HasTracker), deviceID)
}

// Lookup mocks base method.
func (m *MockProcesses) Lookup(deviceID string) (*go_gpsgen.Device, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Lookup", deviceID)
	ret0, _ := ret[0].(*go_gpsgen.Device)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Lookup indicates an expected call of Lookup.
func (mr *MockProcessesMockRecorder) Lookup(deviceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lookup", reflect.TypeOf((*MockProcesses)(nil).Lookup), deviceID)
}