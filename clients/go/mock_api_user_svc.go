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

// MockUserSvcAPI is a mock of UserSvcAPI interface.
type MockUserSvcAPI struct {
	ctrl     *gomock.Controller
	recorder *MockUserSvcAPIMockRecorder
	isgomock struct{}
}

// MockUserSvcAPIMockRecorder is the mock recorder for MockUserSvcAPI.
type MockUserSvcAPIMockRecorder struct {
	mock *MockUserSvcAPI
}

// NewMockUserSvcAPI creates a new mock instance.
func NewMockUserSvcAPI(ctrl *gomock.Controller) *MockUserSvcAPI {
	mock := &MockUserSvcAPI{ctrl: ctrl}
	mock.recorder = &MockUserSvcAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserSvcAPI) EXPECT() *MockUserSvcAPIMockRecorder {
	return m.recorder
}

// AddUserToOrganization mocks base method.
func (m *MockUserSvcAPI) AddUserToOrganization(ctx context.Context, organizationId, userId string) ApiAddUserToOrganizationRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUserToOrganization", ctx, organizationId, userId)
	ret0, _ := ret[0].(ApiAddUserToOrganizationRequest)
	return ret0
}

// AddUserToOrganization indicates an expected call of AddUserToOrganization.
func (mr *MockUserSvcAPIMockRecorder) AddUserToOrganization(ctx, organizationId, userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUserToOrganization", reflect.TypeOf((*MockUserSvcAPI)(nil).AddUserToOrganization), ctx, organizationId, userId)
}

// AddUserToOrganizationExecute mocks base method.
func (m *MockUserSvcAPI) AddUserToOrganizationExecute(r ApiAddUserToOrganizationRequest) (map[string]any, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUserToOrganizationExecute", r)
	ret0, _ := ret[0].(map[string]any)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AddUserToOrganizationExecute indicates an expected call of AddUserToOrganizationExecute.
func (mr *MockUserSvcAPIMockRecorder) AddUserToOrganizationExecute(r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUserToOrganizationExecute", reflect.TypeOf((*MockUserSvcAPI)(nil).AddUserToOrganizationExecute), r)
}

// AssignPermissions mocks base method.
func (m *MockUserSvcAPI) AssignPermissions(ctx context.Context) ApiAssignPermissionsRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignPermissions", ctx)
	ret0, _ := ret[0].(ApiAssignPermissionsRequest)
	return ret0
}

// AssignPermissions indicates an expected call of AssignPermissions.
func (mr *MockUserSvcAPIMockRecorder) AssignPermissions(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignPermissions", reflect.TypeOf((*MockUserSvcAPI)(nil).AssignPermissions), ctx)
}

// AssignPermissionsExecute mocks base method.
func (m *MockUserSvcAPI) AssignPermissionsExecute(r ApiAssignPermissionsRequest) (map[string]any, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignPermissionsExecute", r)
	ret0, _ := ret[0].(map[string]any)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AssignPermissionsExecute indicates an expected call of AssignPermissionsExecute.
func (mr *MockUserSvcAPIMockRecorder) AssignPermissionsExecute(r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignPermissionsExecute", reflect.TypeOf((*MockUserSvcAPI)(nil).AssignPermissionsExecute), r)
}

// AssignRole mocks base method.
func (m *MockUserSvcAPI) AssignRole(ctx context.Context, userId, roleId string) ApiAssignRoleRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignRole", ctx, userId, roleId)
	ret0, _ := ret[0].(ApiAssignRoleRequest)
	return ret0
}

// AssignRole indicates an expected call of AssignRole.
func (mr *MockUserSvcAPIMockRecorder) AssignRole(ctx, userId, roleId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignRole", reflect.TypeOf((*MockUserSvcAPI)(nil).AssignRole), ctx, userId, roleId)
}

// AssignRoleExecute mocks base method.
func (m *MockUserSvcAPI) AssignRoleExecute(r ApiAssignRoleRequest) (map[string]any, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignRoleExecute", r)
	ret0, _ := ret[0].(map[string]any)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AssignRoleExecute indicates an expected call of AssignRoleExecute.
func (mr *MockUserSvcAPIMockRecorder) AssignRoleExecute(r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignRoleExecute", reflect.TypeOf((*MockUserSvcAPI)(nil).AssignRoleExecute), r)
}

// ChangePassword mocks base method.
func (m *MockUserSvcAPI) ChangePassword(ctx context.Context) ApiChangePasswordRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangePassword", ctx)
	ret0, _ := ret[0].(ApiChangePasswordRequest)
	return ret0
}

// ChangePassword indicates an expected call of ChangePassword.
func (mr *MockUserSvcAPIMockRecorder) ChangePassword(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePassword", reflect.TypeOf((*MockUserSvcAPI)(nil).ChangePassword), ctx)
}

// ChangePasswordExecute mocks base method.
func (m *MockUserSvcAPI) ChangePasswordExecute(r ApiChangePasswordRequest) (map[string]any, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangePasswordExecute", r)
	ret0, _ := ret[0].(map[string]any)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ChangePasswordExecute indicates an expected call of ChangePasswordExecute.
func (mr *MockUserSvcAPIMockRecorder) ChangePasswordExecute(r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePasswordExecute", reflect.TypeOf((*MockUserSvcAPI)(nil).ChangePasswordExecute), r)
}

// CreateUser mocks base method.
func (m *MockUserSvcAPI) CreateUser(ctx context.Context) ApiCreateUserRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx)
	ret0, _ := ret[0].(ApiCreateUserRequest)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserSvcAPIMockRecorder) CreateUser(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserSvcAPI)(nil).CreateUser), ctx)
}

// CreateUserExecute mocks base method.
func (m *MockUserSvcAPI) CreateUserExecute(r ApiCreateUserRequest) (map[string]any, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserExecute", r)
	ret0, _ := ret[0].(map[string]any)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateUserExecute indicates an expected call of CreateUserExecute.
func (mr *MockUserSvcAPIMockRecorder) CreateUserExecute(r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserExecute", reflect.TypeOf((*MockUserSvcAPI)(nil).CreateUserExecute), r)
}

// DeleteUser mocks base method.
func (m *MockUserSvcAPI) DeleteUser(ctx context.Context, userId string) ApiDeleteUserRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ctx, userId)
	ret0, _ := ret[0].(ApiDeleteUserRequest)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserSvcAPIMockRecorder) DeleteUser(ctx, userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserSvcAPI)(nil).DeleteUser), ctx, userId)
}

// DeleteUserExecute mocks base method.
func (m *MockUserSvcAPI) DeleteUserExecute(r ApiDeleteUserRequest) (map[string]any, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserExecute", r)
	ret0, _ := ret[0].(map[string]any)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DeleteUserExecute indicates an expected call of DeleteUserExecute.
func (mr *MockUserSvcAPIMockRecorder) DeleteUserExecute(r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserExecute", reflect.TypeOf((*MockUserSvcAPI)(nil).DeleteUserExecute), r)
}

// GetPermissionsByRole mocks base method.
func (m *MockUserSvcAPI) GetPermissionsByRole(ctx context.Context, roleId string) ApiGetPermissionsByRoleRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPermissionsByRole", ctx, roleId)
	ret0, _ := ret[0].(ApiGetPermissionsByRoleRequest)
	return ret0
}

// GetPermissionsByRole indicates an expected call of GetPermissionsByRole.
func (mr *MockUserSvcAPIMockRecorder) GetPermissionsByRole(ctx, roleId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPermissionsByRole", reflect.TypeOf((*MockUserSvcAPI)(nil).GetPermissionsByRole), ctx, roleId)
}

// GetPermissionsByRoleExecute mocks base method.
func (m *MockUserSvcAPI) GetPermissionsByRoleExecute(r ApiGetPermissionsByRoleRequest) (*UserSvcGetPermissionsResponse, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPermissionsByRoleExecute", r)
	ret0, _ := ret[0].(*UserSvcGetPermissionsResponse)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetPermissionsByRoleExecute indicates an expected call of GetPermissionsByRoleExecute.
func (mr *MockUserSvcAPIMockRecorder) GetPermissionsByRoleExecute(r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPermissionsByRoleExecute", reflect.TypeOf((*MockUserSvcAPI)(nil).GetPermissionsByRoleExecute), r)
}

// GetPublicKey mocks base method.
func (m *MockUserSvcAPI) GetPublicKey(ctx context.Context) ApiGetPublicKeyRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPublicKey", ctx)
	ret0, _ := ret[0].(ApiGetPublicKeyRequest)
	return ret0
}

// GetPublicKey indicates an expected call of GetPublicKey.
func (mr *MockUserSvcAPIMockRecorder) GetPublicKey(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublicKey", reflect.TypeOf((*MockUserSvcAPI)(nil).GetPublicKey), ctx)
}

// GetPublicKeyExecute mocks base method.
func (m *MockUserSvcAPI) GetPublicKeyExecute(r ApiGetPublicKeyRequest) (*UserSvcGetPublicKeyResponse, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPublicKeyExecute", r)
	ret0, _ := ret[0].(*UserSvcGetPublicKeyResponse)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetPublicKeyExecute indicates an expected call of GetPublicKeyExecute.
func (mr *MockUserSvcAPIMockRecorder) GetPublicKeyExecute(r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublicKeyExecute", reflect.TypeOf((*MockUserSvcAPI)(nil).GetPublicKeyExecute), r)
}

// IsAuthorized mocks base method.
func (m *MockUserSvcAPI) IsAuthorized(ctx context.Context, permission string) ApiIsAuthorizedRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsAuthorized", ctx, permission)
	ret0, _ := ret[0].(ApiIsAuthorizedRequest)
	return ret0
}

// IsAuthorized indicates an expected call of IsAuthorized.
func (mr *MockUserSvcAPIMockRecorder) IsAuthorized(ctx, permission any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAuthorized", reflect.TypeOf((*MockUserSvcAPI)(nil).IsAuthorized), ctx, permission)
}

// IsAuthorizedExecute mocks base method.
func (m *MockUserSvcAPI) IsAuthorizedExecute(r ApiIsAuthorizedRequest) (*UserSvcIsAuthorizedResponse, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsAuthorizedExecute", r)
	ret0, _ := ret[0].(*UserSvcIsAuthorizedResponse)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// IsAuthorizedExecute indicates an expected call of IsAuthorizedExecute.
func (mr *MockUserSvcAPIMockRecorder) IsAuthorizedExecute(r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAuthorizedExecute", reflect.TypeOf((*MockUserSvcAPI)(nil).IsAuthorizedExecute), r)
}

// ListGrants mocks base method.
func (m *MockUserSvcAPI) ListGrants(ctx context.Context) ApiListGrantsRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListGrants", ctx)
	ret0, _ := ret[0].(ApiListGrantsRequest)
	return ret0
}

// ListGrants indicates an expected call of ListGrants.
func (mr *MockUserSvcAPIMockRecorder) ListGrants(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGrants", reflect.TypeOf((*MockUserSvcAPI)(nil).ListGrants), ctx)
}

// ListGrantsExecute mocks base method.
func (m *MockUserSvcAPI) ListGrantsExecute(r ApiListGrantsRequest) (*UserSvcListGrantsResponse, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListGrantsExecute", r)
	ret0, _ := ret[0].(*UserSvcListGrantsResponse)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListGrantsExecute indicates an expected call of ListGrantsExecute.
func (mr *MockUserSvcAPIMockRecorder) ListGrantsExecute(r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGrantsExecute", reflect.TypeOf((*MockUserSvcAPI)(nil).ListGrantsExecute), r)
}

// ListInvites mocks base method.
func (m *MockUserSvcAPI) ListInvites(ctx context.Context) ApiListInvitesRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListInvites", ctx)
	ret0, _ := ret[0].(ApiListInvitesRequest)
	return ret0
}

// ListInvites indicates an expected call of ListInvites.
func (mr *MockUserSvcAPIMockRecorder) ListInvites(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListInvites", reflect.TypeOf((*MockUserSvcAPI)(nil).ListInvites), ctx)
}

// ListInvitesExecute mocks base method.
func (m *MockUserSvcAPI) ListInvitesExecute(r ApiListInvitesRequest) (*UserSvcListInvitesResponse, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListInvitesExecute", r)
	ret0, _ := ret[0].(*UserSvcListInvitesResponse)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListInvitesExecute indicates an expected call of ListInvitesExecute.
func (mr *MockUserSvcAPIMockRecorder) ListInvitesExecute(r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListInvitesExecute", reflect.TypeOf((*MockUserSvcAPI)(nil).ListInvitesExecute), r)
}

// ListRoles mocks base method.
func (m *MockUserSvcAPI) ListRoles(ctx context.Context) ApiListRolesRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRoles", ctx)
	ret0, _ := ret[0].(ApiListRolesRequest)
	return ret0
}

// ListRoles indicates an expected call of ListRoles.
func (mr *MockUserSvcAPIMockRecorder) ListRoles(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRoles", reflect.TypeOf((*MockUserSvcAPI)(nil).ListRoles), ctx)
}

// ListRolesExecute mocks base method.
func (m *MockUserSvcAPI) ListRolesExecute(r ApiListRolesRequest) (*UserSvcListRolesResponse, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRolesExecute", r)
	ret0, _ := ret[0].(*UserSvcListRolesResponse)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListRolesExecute indicates an expected call of ListRolesExecute.
func (mr *MockUserSvcAPIMockRecorder) ListRolesExecute(r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRolesExecute", reflect.TypeOf((*MockUserSvcAPI)(nil).ListRolesExecute), r)
}

// ListUsers mocks base method.
func (m *MockUserSvcAPI) ListUsers(ctx context.Context) ApiListUsersRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsers", ctx)
	ret0, _ := ret[0].(ApiListUsersRequest)
	return ret0
}

// ListUsers indicates an expected call of ListUsers.
func (mr *MockUserSvcAPIMockRecorder) ListUsers(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsers", reflect.TypeOf((*MockUserSvcAPI)(nil).ListUsers), ctx)
}

// ListUsersExecute mocks base method.
func (m *MockUserSvcAPI) ListUsersExecute(r ApiListUsersRequest) (*UserSvcListUsersResponse, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsersExecute", r)
	ret0, _ := ret[0].(*UserSvcListUsersResponse)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListUsersExecute indicates an expected call of ListUsersExecute.
func (mr *MockUserSvcAPIMockRecorder) ListUsersExecute(r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsersExecute", reflect.TypeOf((*MockUserSvcAPI)(nil).ListUsersExecute), r)
}

// Login mocks base method.
func (m *MockUserSvcAPI) Login(ctx context.Context) ApiLoginRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx)
	ret0, _ := ret[0].(ApiLoginRequest)
	return ret0
}

// Login indicates an expected call of Login.
func (mr *MockUserSvcAPIMockRecorder) Login(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUserSvcAPI)(nil).Login), ctx)
}

// LoginExecute mocks base method.
func (m *MockUserSvcAPI) LoginExecute(r ApiLoginRequest) (*UserSvcLoginResponse, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginExecute", r)
	ret0, _ := ret[0].(*UserSvcLoginResponse)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// LoginExecute indicates an expected call of LoginExecute.
func (mr *MockUserSvcAPIMockRecorder) LoginExecute(r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginExecute", reflect.TypeOf((*MockUserSvcAPI)(nil).LoginExecute), r)
}

// ReadUserByToken mocks base method.
func (m *MockUserSvcAPI) ReadUserByToken(ctx context.Context) ApiReadUserByTokenRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadUserByToken", ctx)
	ret0, _ := ret[0].(ApiReadUserByTokenRequest)
	return ret0
}

// ReadUserByToken indicates an expected call of ReadUserByToken.
func (mr *MockUserSvcAPIMockRecorder) ReadUserByToken(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadUserByToken", reflect.TypeOf((*MockUserSvcAPI)(nil).ReadUserByToken), ctx)
}

// ReadUserByTokenExecute mocks base method.
func (m *MockUserSvcAPI) ReadUserByTokenExecute(r ApiReadUserByTokenRequest) (*UserSvcReadUserByTokenResponse, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadUserByTokenExecute", r)
	ret0, _ := ret[0].(*UserSvcReadUserByTokenResponse)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ReadUserByTokenExecute indicates an expected call of ReadUserByTokenExecute.
func (mr *MockUserSvcAPIMockRecorder) ReadUserByTokenExecute(r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadUserByTokenExecute", reflect.TypeOf((*MockUserSvcAPI)(nil).ReadUserByTokenExecute), r)
}

// Register mocks base method.
func (m *MockUserSvcAPI) Register(ctx context.Context) ApiRegisterRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", ctx)
	ret0, _ := ret[0].(ApiRegisterRequest)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockUserSvcAPIMockRecorder) Register(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockUserSvcAPI)(nil).Register), ctx)
}

// RegisterExecute mocks base method.
func (m *MockUserSvcAPI) RegisterExecute(r ApiRegisterRequest) (*UserSvcRegisterResponse, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterExecute", r)
	ret0, _ := ret[0].(*UserSvcRegisterResponse)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RegisterExecute indicates an expected call of RegisterExecute.
func (mr *MockUserSvcAPIMockRecorder) RegisterExecute(r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterExecute", reflect.TypeOf((*MockUserSvcAPI)(nil).RegisterExecute), r)
}

// RemoveUserFromOrganization mocks base method.
func (m *MockUserSvcAPI) RemoveUserFromOrganization(ctx context.Context, organizationId, userId string) ApiRemoveUserFromOrganizationRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveUserFromOrganization", ctx, organizationId, userId)
	ret0, _ := ret[0].(ApiRemoveUserFromOrganizationRequest)
	return ret0
}

// RemoveUserFromOrganization indicates an expected call of RemoveUserFromOrganization.
func (mr *MockUserSvcAPIMockRecorder) RemoveUserFromOrganization(ctx, organizationId, userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveUserFromOrganization", reflect.TypeOf((*MockUserSvcAPI)(nil).RemoveUserFromOrganization), ctx, organizationId, userId)
}

// RemoveUserFromOrganizationExecute mocks base method.
func (m *MockUserSvcAPI) RemoveUserFromOrganizationExecute(r ApiRemoveUserFromOrganizationRequest) (map[string]any, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveUserFromOrganizationExecute", r)
	ret0, _ := ret[0].(map[string]any)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RemoveUserFromOrganizationExecute indicates an expected call of RemoveUserFromOrganizationExecute.
func (mr *MockUserSvcAPIMockRecorder) RemoveUserFromOrganizationExecute(r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveUserFromOrganizationExecute", reflect.TypeOf((*MockUserSvcAPI)(nil).RemoveUserFromOrganizationExecute), r)
}

// ResetPassword mocks base method.
func (m *MockUserSvcAPI) ResetPassword(ctx context.Context, userId string) ApiResetPasswordRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetPassword", ctx, userId)
	ret0, _ := ret[0].(ApiResetPasswordRequest)
	return ret0
}

// ResetPassword indicates an expected call of ResetPassword.
func (mr *MockUserSvcAPIMockRecorder) ResetPassword(ctx, userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetPassword", reflect.TypeOf((*MockUserSvcAPI)(nil).ResetPassword), ctx, userId)
}

// ResetPasswordExecute mocks base method.
func (m *MockUserSvcAPI) ResetPasswordExecute(r ApiResetPasswordRequest) (map[string]any, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetPasswordExecute", r)
	ret0, _ := ret[0].(map[string]any)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ResetPasswordExecute indicates an expected call of ResetPasswordExecute.
func (mr *MockUserSvcAPIMockRecorder) ResetPasswordExecute(r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetPasswordExecute", reflect.TypeOf((*MockUserSvcAPI)(nil).ResetPasswordExecute), r)
}

// SaveGrants mocks base method.
func (m *MockUserSvcAPI) SaveGrants(ctx context.Context) ApiSaveGrantsRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveGrants", ctx)
	ret0, _ := ret[0].(ApiSaveGrantsRequest)
	return ret0
}

// SaveGrants indicates an expected call of SaveGrants.
func (mr *MockUserSvcAPIMockRecorder) SaveGrants(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveGrants", reflect.TypeOf((*MockUserSvcAPI)(nil).SaveGrants), ctx)
}

// SaveGrantsExecute mocks base method.
func (m *MockUserSvcAPI) SaveGrantsExecute(r ApiSaveGrantsRequest) (map[string]any, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveGrantsExecute", r)
	ret0, _ := ret[0].(map[string]any)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SaveGrantsExecute indicates an expected call of SaveGrantsExecute.
func (mr *MockUserSvcAPIMockRecorder) SaveGrantsExecute(r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveGrantsExecute", reflect.TypeOf((*MockUserSvcAPI)(nil).SaveGrantsExecute), r)
}

// SaveInvites mocks base method.
func (m *MockUserSvcAPI) SaveInvites(ctx context.Context) ApiSaveInvitesRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveInvites", ctx)
	ret0, _ := ret[0].(ApiSaveInvitesRequest)
	return ret0
}

// SaveInvites indicates an expected call of SaveInvites.
func (mr *MockUserSvcAPIMockRecorder) SaveInvites(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveInvites", reflect.TypeOf((*MockUserSvcAPI)(nil).SaveInvites), ctx)
}

// SaveInvitesExecute mocks base method.
func (m *MockUserSvcAPI) SaveInvitesExecute(r ApiSaveInvitesRequest) (*UserSvcSaveInvitesResponse, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveInvitesExecute", r)
	ret0, _ := ret[0].(*UserSvcSaveInvitesResponse)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SaveInvitesExecute indicates an expected call of SaveInvitesExecute.
func (mr *MockUserSvcAPIMockRecorder) SaveInvitesExecute(r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveInvitesExecute", reflect.TypeOf((*MockUserSvcAPI)(nil).SaveInvitesExecute), r)
}

// SaveOrganization mocks base method.
func (m *MockUserSvcAPI) SaveOrganization(ctx context.Context) ApiSaveOrganizationRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveOrganization", ctx)
	ret0, _ := ret[0].(ApiSaveOrganizationRequest)
	return ret0
}

// SaveOrganization indicates an expected call of SaveOrganization.
func (mr *MockUserSvcAPIMockRecorder) SaveOrganization(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveOrganization", reflect.TypeOf((*MockUserSvcAPI)(nil).SaveOrganization), ctx)
}

// SaveOrganizationExecute mocks base method.
func (m *MockUserSvcAPI) SaveOrganizationExecute(r ApiSaveOrganizationRequest) (*UserSvcSaveOrganizationResponse, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveOrganizationExecute", r)
	ret0, _ := ret[0].(*UserSvcSaveOrganizationResponse)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SaveOrganizationExecute indicates an expected call of SaveOrganizationExecute.
func (mr *MockUserSvcAPIMockRecorder) SaveOrganizationExecute(r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveOrganizationExecute", reflect.TypeOf((*MockUserSvcAPI)(nil).SaveOrganizationExecute), r)
}

// SaveSelf mocks base method.
func (m *MockUserSvcAPI) SaveSelf(ctx context.Context, userId string) ApiSaveSelfRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveSelf", ctx, userId)
	ret0, _ := ret[0].(ApiSaveSelfRequest)
	return ret0
}

// SaveSelf indicates an expected call of SaveSelf.
func (mr *MockUserSvcAPIMockRecorder) SaveSelf(ctx, userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveSelf", reflect.TypeOf((*MockUserSvcAPI)(nil).SaveSelf), ctx, userId)
}

// SaveSelfExecute mocks base method.
func (m *MockUserSvcAPI) SaveSelfExecute(r ApiSaveSelfRequest) (map[string]any, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveSelfExecute", r)
	ret0, _ := ret[0].(map[string]any)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SaveSelfExecute indicates an expected call of SaveSelfExecute.
func (mr *MockUserSvcAPIMockRecorder) SaveSelfExecute(r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveSelfExecute", reflect.TypeOf((*MockUserSvcAPI)(nil).SaveSelfExecute), r)
}

// SaveUser mocks base method.
func (m *MockUserSvcAPI) SaveUser(ctx context.Context, userId string) ApiSaveUserRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveUser", ctx, userId)
	ret0, _ := ret[0].(ApiSaveUserRequest)
	return ret0
}

// SaveUser indicates an expected call of SaveUser.
func (mr *MockUserSvcAPIMockRecorder) SaveUser(ctx, userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveUser", reflect.TypeOf((*MockUserSvcAPI)(nil).SaveUser), ctx, userId)
}

// SaveUserExecute mocks base method.
func (m *MockUserSvcAPI) SaveUserExecute(r ApiSaveUserRequest) (map[string]any, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveUserExecute", r)
	ret0, _ := ret[0].(map[string]any)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SaveUserExecute indicates an expected call of SaveUserExecute.
func (mr *MockUserSvcAPIMockRecorder) SaveUserExecute(r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveUserExecute", reflect.TypeOf((*MockUserSvcAPI)(nil).SaveUserExecute), r)
}
