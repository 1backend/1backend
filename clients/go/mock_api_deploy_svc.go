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

// MockDeploySvcAPI is a mock of DeploySvcAPI interface.
type MockDeploySvcAPI struct {
	ctrl     *gomock.Controller
	recorder *MockDeploySvcAPIMockRecorder
	isgomock struct{}
}

// MockDeploySvcAPIMockRecorder is the mock recorder for MockDeploySvcAPI.
type MockDeploySvcAPIMockRecorder struct {
	mock *MockDeploySvcAPI
}

// NewMockDeploySvcAPI creates a new mock instance.
func NewMockDeploySvcAPI(ctrl *gomock.Controller) *MockDeploySvcAPI {
	mock := &MockDeploySvcAPI{ctrl: ctrl}
	mock.recorder = &MockDeploySvcAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeploySvcAPI) EXPECT() *MockDeploySvcAPIMockRecorder {
	return m.recorder
}

// DeleteDeployment mocks base method.
func (m *MockDeploySvcAPI) DeleteDeployment(ctx context.Context) ApiDeleteDeploymentRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDeployment", ctx)
	ret0, _ := ret[0].(ApiDeleteDeploymentRequest)
	return ret0
}

// DeleteDeployment indicates an expected call of DeleteDeployment.
func (mr *MockDeploySvcAPIMockRecorder) DeleteDeployment(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDeployment", reflect.TypeOf((*MockDeploySvcAPI)(nil).DeleteDeployment), ctx)
}

// DeleteDeploymentExecute mocks base method.
func (m *MockDeploySvcAPI) DeleteDeploymentExecute(r ApiDeleteDeploymentRequest) (map[string]any, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDeploymentExecute", r)
	ret0, _ := ret[0].(map[string]any)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DeleteDeploymentExecute indicates an expected call of DeleteDeploymentExecute.
func (mr *MockDeploySvcAPIMockRecorder) DeleteDeploymentExecute(r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDeploymentExecute", reflect.TypeOf((*MockDeploySvcAPI)(nil).DeleteDeploymentExecute), r)
}

// ListDeployments mocks base method.
func (m *MockDeploySvcAPI) ListDeployments(ctx context.Context) ApiListDeploymentsRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDeployments", ctx)
	ret0, _ := ret[0].(ApiListDeploymentsRequest)
	return ret0
}

// ListDeployments indicates an expected call of ListDeployments.
func (mr *MockDeploySvcAPIMockRecorder) ListDeployments(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDeployments", reflect.TypeOf((*MockDeploySvcAPI)(nil).ListDeployments), ctx)
}

// ListDeploymentsExecute mocks base method.
func (m *MockDeploySvcAPI) ListDeploymentsExecute(r ApiListDeploymentsRequest) (*DeploySvcListDeploymentsResponse, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDeploymentsExecute", r)
	ret0, _ := ret[0].(*DeploySvcListDeploymentsResponse)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListDeploymentsExecute indicates an expected call of ListDeploymentsExecute.
func (mr *MockDeploySvcAPIMockRecorder) ListDeploymentsExecute(r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDeploymentsExecute", reflect.TypeOf((*MockDeploySvcAPI)(nil).ListDeploymentsExecute), r)
}

// SaveDeployment mocks base method.
func (m *MockDeploySvcAPI) SaveDeployment(ctx context.Context) ApiSaveDeploymentRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveDeployment", ctx)
	ret0, _ := ret[0].(ApiSaveDeploymentRequest)
	return ret0
}

// SaveDeployment indicates an expected call of SaveDeployment.
func (mr *MockDeploySvcAPIMockRecorder) SaveDeployment(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveDeployment", reflect.TypeOf((*MockDeploySvcAPI)(nil).SaveDeployment), ctx)
}

// SaveDeploymentExecute mocks base method.
func (m *MockDeploySvcAPI) SaveDeploymentExecute(r ApiSaveDeploymentRequest) (map[string]any, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveDeploymentExecute", r)
	ret0, _ := ret[0].(map[string]any)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SaveDeploymentExecute indicates an expected call of SaveDeploymentExecute.
func (mr *MockDeploySvcAPIMockRecorder) SaveDeploymentExecute(r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveDeploymentExecute", reflect.TypeOf((*MockDeploySvcAPI)(nil).SaveDeploymentExecute), r)
}
