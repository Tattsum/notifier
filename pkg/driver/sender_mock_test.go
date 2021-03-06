// Code generated by MockGen. DO NOT EDIT.
// Source: sender.go

// Package driver is a generated GoMock package.
package driver

import (
	parser "github.com/bannzai/notifier/pkg/parser"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockSender is a mock of Sender interface
type MockSender struct {
	ctrl     *gomock.Controller
	recorder *MockSenderMockRecorder
}

// MockSenderMockRecorder is the mock recorder for MockSender
type MockSenderMockRecorder struct {
	mock *MockSender
}

// NewMockSender creates a new mock instance
func NewMockSender(ctrl *gomock.Controller) *MockSender {
	mock := &MockSender{ctrl: ctrl}
	mock.recorder = &MockSenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSender) EXPECT() *MockSenderMockRecorder {
	return m.recorder
}

// Send mocks base method
func (m *MockSender) Send(content parser.Content) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", content)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockSenderMockRecorder) Send(content interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockSender)(nil).Send), content)
}
