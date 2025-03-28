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
	os "os"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockFileSvcAPI is a mock of FileSvcAPI interface.
type MockFileSvcAPI struct {
	ctrl     *gomock.Controller
	recorder *MockFileSvcAPIMockRecorder
	isgomock struct{}
}

// MockFileSvcAPIMockRecorder is the mock recorder for MockFileSvcAPI.
type MockFileSvcAPIMockRecorder struct {
	mock *MockFileSvcAPI
}

// NewMockFileSvcAPI creates a new mock instance.
func NewMockFileSvcAPI(ctrl *gomock.Controller) *MockFileSvcAPI {
	mock := &MockFileSvcAPI{ctrl: ctrl}
	mock.recorder = &MockFileSvcAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileSvcAPI) EXPECT() *MockFileSvcAPIMockRecorder {
	return m.recorder
}

// DownloadFile mocks base method.
func (m *MockFileSvcAPI) DownloadFile(ctx context.Context) ApiDownloadFileRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadFile", ctx)
	ret0, _ := ret[0].(ApiDownloadFileRequest)
	return ret0
}

// DownloadFile indicates an expected call of DownloadFile.
func (mr *MockFileSvcAPIMockRecorder) DownloadFile(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadFile", reflect.TypeOf((*MockFileSvcAPI)(nil).DownloadFile), ctx)
}

// DownloadFileExecute mocks base method.
func (m *MockFileSvcAPI) DownloadFileExecute(r ApiDownloadFileRequest) (map[string]any, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadFileExecute", r)
	ret0, _ := ret[0].(map[string]any)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DownloadFileExecute indicates an expected call of DownloadFileExecute.
func (mr *MockFileSvcAPIMockRecorder) DownloadFileExecute(r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadFileExecute", reflect.TypeOf((*MockFileSvcAPI)(nil).DownloadFileExecute), r)
}

// GetDownload mocks base method.
func (m *MockFileSvcAPI) GetDownload(ctx context.Context, url string) ApiGetDownloadRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDownload", ctx, url)
	ret0, _ := ret[0].(ApiGetDownloadRequest)
	return ret0
}

// GetDownload indicates an expected call of GetDownload.
func (mr *MockFileSvcAPIMockRecorder) GetDownload(ctx, url any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDownload", reflect.TypeOf((*MockFileSvcAPI)(nil).GetDownload), ctx, url)
}

// GetDownloadExecute mocks base method.
func (m *MockFileSvcAPI) GetDownloadExecute(r ApiGetDownloadRequest) (*FileSvcGetDownloadResponse, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDownloadExecute", r)
	ret0, _ := ret[0].(*FileSvcGetDownloadResponse)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetDownloadExecute indicates an expected call of GetDownloadExecute.
func (mr *MockFileSvcAPIMockRecorder) GetDownloadExecute(r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDownloadExecute", reflect.TypeOf((*MockFileSvcAPI)(nil).GetDownloadExecute), r)
}

// ListDownloads mocks base method.
func (m *MockFileSvcAPI) ListDownloads(ctx context.Context) ApiListDownloadsRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDownloads", ctx)
	ret0, _ := ret[0].(ApiListDownloadsRequest)
	return ret0
}

// ListDownloads indicates an expected call of ListDownloads.
func (mr *MockFileSvcAPIMockRecorder) ListDownloads(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDownloads", reflect.TypeOf((*MockFileSvcAPI)(nil).ListDownloads), ctx)
}

// ListDownloadsExecute mocks base method.
func (m *MockFileSvcAPI) ListDownloadsExecute(r ApiListDownloadsRequest) (*FileSvcDownloadsResponse, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDownloadsExecute", r)
	ret0, _ := ret[0].(*FileSvcDownloadsResponse)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListDownloadsExecute indicates an expected call of ListDownloadsExecute.
func (mr *MockFileSvcAPIMockRecorder) ListDownloadsExecute(r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDownloadsExecute", reflect.TypeOf((*MockFileSvcAPI)(nil).ListDownloadsExecute), r)
}

// ListUploads mocks base method.
func (m *MockFileSvcAPI) ListUploads(ctx context.Context) ApiListUploadsRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUploads", ctx)
	ret0, _ := ret[0].(ApiListUploadsRequest)
	return ret0
}

// ListUploads indicates an expected call of ListUploads.
func (mr *MockFileSvcAPIMockRecorder) ListUploads(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUploads", reflect.TypeOf((*MockFileSvcAPI)(nil).ListUploads), ctx)
}

// ListUploadsExecute mocks base method.
func (m *MockFileSvcAPI) ListUploadsExecute(r ApiListUploadsRequest) (*FileSvcListUploadsResponse, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUploadsExecute", r)
	ret0, _ := ret[0].(*FileSvcListUploadsResponse)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListUploadsExecute indicates an expected call of ListUploadsExecute.
func (mr *MockFileSvcAPIMockRecorder) ListUploadsExecute(r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUploadsExecute", reflect.TypeOf((*MockFileSvcAPI)(nil).ListUploadsExecute), r)
}

// PauseDownload mocks base method.
func (m *MockFileSvcAPI) PauseDownload(ctx context.Context, url string) ApiPauseDownloadRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PauseDownload", ctx, url)
	ret0, _ := ret[0].(ApiPauseDownloadRequest)
	return ret0
}

// PauseDownload indicates an expected call of PauseDownload.
func (mr *MockFileSvcAPIMockRecorder) PauseDownload(ctx, url any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PauseDownload", reflect.TypeOf((*MockFileSvcAPI)(nil).PauseDownload), ctx, url)
}

// PauseDownloadExecute mocks base method.
func (m *MockFileSvcAPI) PauseDownloadExecute(r ApiPauseDownloadRequest) (map[string]any, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PauseDownloadExecute", r)
	ret0, _ := ret[0].(map[string]any)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// PauseDownloadExecute indicates an expected call of PauseDownloadExecute.
func (mr *MockFileSvcAPIMockRecorder) PauseDownloadExecute(r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PauseDownloadExecute", reflect.TypeOf((*MockFileSvcAPI)(nil).PauseDownloadExecute), r)
}

// ServeDownload mocks base method.
func (m *MockFileSvcAPI) ServeDownload(ctx context.Context, url string) ApiServeDownloadRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServeDownload", ctx, url)
	ret0, _ := ret[0].(ApiServeDownloadRequest)
	return ret0
}

// ServeDownload indicates an expected call of ServeDownload.
func (mr *MockFileSvcAPIMockRecorder) ServeDownload(ctx, url any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServeDownload", reflect.TypeOf((*MockFileSvcAPI)(nil).ServeDownload), ctx, url)
}

// ServeDownloadExecute mocks base method.
func (m *MockFileSvcAPI) ServeDownloadExecute(r ApiServeDownloadRequest) (*os.File, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServeDownloadExecute", r)
	ret0, _ := ret[0].(*os.File)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ServeDownloadExecute indicates an expected call of ServeDownloadExecute.
func (mr *MockFileSvcAPIMockRecorder) ServeDownloadExecute(r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServeDownloadExecute", reflect.TypeOf((*MockFileSvcAPI)(nil).ServeDownloadExecute), r)
}

// ServeUpload mocks base method.
func (m *MockFileSvcAPI) ServeUpload(ctx context.Context, fileId string) ApiServeUploadRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServeUpload", ctx, fileId)
	ret0, _ := ret[0].(ApiServeUploadRequest)
	return ret0
}

// ServeUpload indicates an expected call of ServeUpload.
func (mr *MockFileSvcAPIMockRecorder) ServeUpload(ctx, fileId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServeUpload", reflect.TypeOf((*MockFileSvcAPI)(nil).ServeUpload), ctx, fileId)
}

// ServeUploadExecute mocks base method.
func (m *MockFileSvcAPI) ServeUploadExecute(r ApiServeUploadRequest) (*os.File, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServeUploadExecute", r)
	ret0, _ := ret[0].(*os.File)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ServeUploadExecute indicates an expected call of ServeUploadExecute.
func (mr *MockFileSvcAPIMockRecorder) ServeUploadExecute(r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServeUploadExecute", reflect.TypeOf((*MockFileSvcAPI)(nil).ServeUploadExecute), r)
}

// UploadFile mocks base method.
func (m *MockFileSvcAPI) UploadFile(ctx context.Context) ApiUploadFileRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadFile", ctx)
	ret0, _ := ret[0].(ApiUploadFileRequest)
	return ret0
}

// UploadFile indicates an expected call of UploadFile.
func (mr *MockFileSvcAPIMockRecorder) UploadFile(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadFile", reflect.TypeOf((*MockFileSvcAPI)(nil).UploadFile), ctx)
}

// UploadFileExecute mocks base method.
func (m *MockFileSvcAPI) UploadFileExecute(r ApiUploadFileRequest) (*FileSvcUploadFileResponse, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadFileExecute", r)
	ret0, _ := ret[0].(*FileSvcUploadFileResponse)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UploadFileExecute indicates an expected call of UploadFileExecute.
func (mr *MockFileSvcAPIMockRecorder) UploadFileExecute(r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadFileExecute", reflect.TypeOf((*MockFileSvcAPI)(nil).UploadFileExecute), r)
}
