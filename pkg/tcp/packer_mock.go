// Code generated by MockGen. DO NOT EDIT.
// Source: packer.go

// Package server is a generated GoMock package.
package server

import (
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPacker is a mock of Packer interface.
type MockPacker struct {
	ctrl     *gomock.Controller
	recorder *MockPackerMockRecorder
}

// MockPackerMockRecorder is the mock recorder for MockPacker.
type MockPackerMockRecorder struct {
	mock *MockPacker
}

// NewMockPacker creates a new mock instance.
func NewMockPacker(ctrl *gomock.Controller) *MockPacker {
	mock := &MockPacker{ctrl: ctrl}
	mock.recorder = &MockPackerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPacker) EXPECT() *MockPackerMockRecorder {
	return m.recorder
}

// Pack mocks base method.
func (m *MockPacker) Pack(message *Message) []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pack", message)
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Pack indicates an expected call of Pack.
func (mr *MockPackerMockRecorder) Pack(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pack", reflect.TypeOf((*MockPacker)(nil).Pack), message)
}

// Unpack mocks base method.
func (m *MockPacker) Unpack(reader io.Reader) (*Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unpack", reader)
	ret0, _ := ret[0].(*Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unpack indicates an expected call of Unpack.
func (mr *MockPackerMockRecorder) Unpack(reader interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unpack", reflect.TypeOf((*MockPacker)(nil).Unpack), reader)
}
