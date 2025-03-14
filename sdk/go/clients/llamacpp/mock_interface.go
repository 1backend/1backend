// Code generated by MockGen. DO NOT EDIT.
//
// Generated by this command:
//
//

// Package llamacpp is a generated GoMock package.
package llamacpp

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockClientI is a mock of ClientI interface.
type MockClientI struct {
	ctrl     *gomock.Controller
	recorder *MockClientIMockRecorder
	isgomock struct{}
}

// MockClientIMockRecorder is the mock recorder for MockClientI.
type MockClientIMockRecorder struct {
	mock *MockClientI
}

// NewMockClientI creates a new mock instance.
func NewMockClientI(ctrl *gomock.Controller) *MockClientI {
	mock := &MockClientI{ctrl: ctrl}
	mock.recorder = &MockClientIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientI) EXPECT() *MockClientIMockRecorder {
	return m.recorder
}

// PostCompletions mocks base method.
func (m *MockClientI) PostCompletions(prompt PostCompletionsRequest) (*CompletionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostCompletions", prompt)
	ret0, _ := ret[0].(*CompletionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostCompletions indicates an expected call of PostCompletions.
func (mr *MockClientIMockRecorder) PostCompletions(prompt any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostCompletions", reflect.TypeOf((*MockClientI)(nil).PostCompletions), prompt)
}

// PostCompletionsStreamed mocks base method.
func (m *MockClientI) PostCompletionsStreamed(prompt PostCompletionsRequest, callback StreamCallback) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostCompletionsStreamed", prompt, callback)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostCompletionsStreamed indicates an expected call of PostCompletionsStreamed.
func (mr *MockClientIMockRecorder) PostCompletionsStreamed(prompt, callback any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostCompletionsStreamed", reflect.TypeOf((*MockClientI)(nil).PostCompletionsStreamed), prompt, callback)
}

// SetAddress mocks base method.
func (m *MockClientI) SetAddress(address string) ClientI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetAddress", address)
	ret0, _ := ret[0].(ClientI)
	return ret0
}

// SetAddress indicates an expected call of SetAddress.
func (mr *MockClientIMockRecorder) SetAddress(address any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAddress", reflect.TypeOf((*MockClientI)(nil).SetAddress), address)
}
