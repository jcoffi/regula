// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fugue/regula/v2/pkg/loader (interfaces: InputDirectory)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	loader "github.com/fugue/regula/v2/pkg/loader"
	gomock "github.com/golang/mock/gomock"
)

// MockInputDirectory is a mock of InputDirectory interface.
type MockInputDirectory struct {
	ctrl     *gomock.Controller
	recorder *MockInputDirectoryMockRecorder
}

// MockInputDirectoryMockRecorder is the mock recorder for MockInputDirectory.
type MockInputDirectoryMockRecorder struct {
	mock *MockInputDirectory
}

// NewMockInputDirectory creates a new mock instance.
func NewMockInputDirectory(ctrl *gomock.Controller) *MockInputDirectory {
	mock := &MockInputDirectory{ctrl: ctrl}
	mock.recorder = &MockInputDirectoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInputDirectory) EXPECT() *MockInputDirectoryMockRecorder {
	return m.recorder
}

// Children mocks base method.
func (m *MockInputDirectory) Children() []loader.InputPath {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Children")
	ret0, _ := ret[0].([]loader.InputPath)
	return ret0
}

// Children indicates an expected call of Children.
func (mr *MockInputDirectoryMockRecorder) Children() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Children", reflect.TypeOf((*MockInputDirectory)(nil).Children))
}

// DetectType mocks base method.
func (m *MockInputDirectory) DetectType(arg0 loader.ConfigurationDetector, arg1 loader.DetectOptions) (loader.IACConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetectType", arg0, arg1)
	ret0, _ := ret[0].(loader.IACConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DetectType indicates an expected call of DetectType.
func (mr *MockInputDirectoryMockRecorder) DetectType(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetectType", reflect.TypeOf((*MockInputDirectory)(nil).DetectType), arg0, arg1)
}

// IsDir mocks base method.
func (m *MockInputDirectory) IsDir() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsDir")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsDir indicates an expected call of IsDir.
func (mr *MockInputDirectoryMockRecorder) IsDir() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsDir", reflect.TypeOf((*MockInputDirectory)(nil).IsDir))
}

// Name mocks base method.
func (m *MockInputDirectory) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockInputDirectoryMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockInputDirectory)(nil).Name))
}

// Path mocks base method.
func (m *MockInputDirectory) Path() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Path")
	ret0, _ := ret[0].(string)
	return ret0
}

// Path indicates an expected call of Path.
func (mr *MockInputDirectoryMockRecorder) Path() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Path", reflect.TypeOf((*MockInputDirectory)(nil).Path))
}

// Walk mocks base method.
func (m *MockInputDirectory) Walk(arg0 loader.WalkFunc) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Walk", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Walk indicates an expected call of Walk.
func (mr *MockInputDirectoryMockRecorder) Walk(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Walk", reflect.TypeOf((*MockInputDirectory)(nil).Walk), arg0)
}
