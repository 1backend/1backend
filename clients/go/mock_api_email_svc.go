// Code generated by MockGen. DO NOT EDIT.
//
// Generated by this command:
//
//

// Package openapi is a generated GoMock package.
package openapi

import (
	context "context"
	http "net/http"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockEmailSvcAPI is a mock of EmailSvcAPI interface.
type MockEmailSvcAPI struct {
	ctrl     *gomock.Controller
	recorder *MockEmailSvcAPIMockRecorder
	isgomock struct{}
}

// MockEmailSvcAPIMockRecorder is the mock recorder for MockEmailSvcAPI.
type MockEmailSvcAPIMockRecorder struct {
	mock *MockEmailSvcAPI
}

// NewMockEmailSvcAPI creates a new mock instance.
func NewMockEmailSvcAPI(ctrl *gomock.Controller) *MockEmailSvcAPI {
	mock := &MockEmailSvcAPI{ctrl: ctrl}
	mock.recorder = &MockEmailSvcAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmailSvcAPI) EXPECT() *MockEmailSvcAPIMockRecorder {
	return m.recorder
}

// SendEmail mocks base method.
func (m *MockEmailSvcAPI) SendEmail(ctx context.Context) ApiSendEmailRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendEmail", ctx)
	ret0, _ := ret[0].(ApiSendEmailRequest)
	return ret0
}

// SendEmail indicates an expected call of SendEmail.
func (mr *MockEmailSvcAPIMockRecorder) SendEmail(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendEmail", reflect.TypeOf((*MockEmailSvcAPI)(nil).SendEmail), ctx)
}

// SendEmailExecute mocks base method.
func (m *MockEmailSvcAPI) SendEmailExecute(r ApiSendEmailRequest) (*EmailSvcSendEmailResponse, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendEmailExecute", r)
	ret0, _ := ret[0].(*EmailSvcSendEmailResponse)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SendEmailExecute indicates an expected call of SendEmailExecute.
func (mr *MockEmailSvcAPIMockRecorder) SendEmailExecute(r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendEmailExecute", reflect.TypeOf((*MockEmailSvcAPI)(nil).SendEmailExecute), r)
}
