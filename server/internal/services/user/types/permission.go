/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package user_svc

import "time"

type Permission struct {
	// eg. "user.viewer"
	Id        string    `json:"id,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	// eg. "User Viewer"
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

func (c *Permission) GetId() string {
	return c.Id
}

func (c *Permission) GetUpdatedAt() string {
	return c.Id
}

type IsAuthorizedRequest struct {
	GrantedSlugs    []string `json:"grantedSlugs,omitempty"`
	ContactsGranted []string `json:"contactsGranted,omitempty"`
}

type IsAuthorizedResponse struct {
	Authorized bool  `json:"authorized,omitempty"`
	User       *User `json:"user,omitempty"`
}

type CreatePermissionRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreatePermissionResponse struct {
}

type GetPermissionsRequest struct{}

type GetPermissionsResponse struct {
	Permissions []*Permission `json:"permissions"`
}

type SavePermissionsRequest struct {
	Permissions []*Permission `json:"permissions"`
}

type SavePermissionsResponse struct {
	Permissions []*Permission `json:"permissions"`
}

type PermissionLink struct {
	RoleId       string `json:"roleId" binding:"required"`
	PermissionId string `json:"permissionId" binding:"required"`
}

type AssignPermissionsRequest struct {
	PermissionLinks []*PermissionLink `json:"permissionLinks"`
}

type AssignPermissionsResponse struct{}

var PermissionUserCreate = Permission{
	Id:   "user-svc:user:create",
	Name: "User Svc - User Create",
}

var PermissionUserView = Permission{
	Id:   "user-svc:user:view",
	Name: "User Svc - User View",
}

var PermissionUserEdit = Permission{
	Id:   "user-svc:user:edit",
	Name: "User Svc - User Edit",
}

var PermissionUserDelete = Permission{
	Id:   "user-svc:user:delete",
	Name: "User Svc - User Delete",
}

var PermissionUserPasswordChange = Permission{
	Id:   "user-svc:user:passwordChange",
	Name: "User Svc - User Password Change",
}

var PermissionRoleCreate = Permission{
	Id:   "user-svc:role:create",
	Name: "User Svc - Role Create",
}

var PermissionRoleView = Permission{
	Id:   "user-svc:role:view",
	Name: "User Svc - Role View",
}

var PermissionRoleEdit = Permission{
	Id:   "user-svc:role:edit",
	Name: "User Svc - Role Edit",
}

var PermissionRoleDelete = Permission{
	Id:   "user-svc:role:delete",
	Name: "User Svc - Role Delete",
}

var PermissionPermissionCreate = Permission{
	Id:   "user-svc:permission:create",
	Name: "User Svc - Permission Create",
}

var PermissionPermissionEdit = Permission{
	Id:   "user-svc:permission:edit",
	Name: "User Svc - Permission Edit",
}

var PermissionOrganizationCreate = Permission{
	Id:   "user-svc:organization:create",
	Name: "User Svc - Organization Create",
}

var PermissionOrganizationAddUser = Permission{
	Id:   "user-svc:organization:add-user",
	Name: "User Svc - Organization Add User",
}

var PermissionOrganizationRemoveUser = Permission{
	Id:   "user-svc:organization:remove-user",
	Name: "User Svc - Organization Remove User",
}

var PermissionGrantCreate = Permission{
	Id:   "user-svc:grant:create",
	Name: "User Svc - Create Grant",
}

var PermissionGrantView = Permission{
	Id:   "user-svc:grant:view",
	Name: "User Svc - View Grant",
}

var PermissionInviteEdit = Permission{
	Id:   "user-svc:invite:edit",
	Name: "User Svc - Edit Invite",
}

var PermissionInviteView = Permission{
	Id:   "user-svc:invite:view",
	Name: "User Svc - View Invite",
}

var UserPermissions = []*Permission{
	&PermissionUserPasswordChange,

	// Anyone can create and edit their own roles,
	// provided the role ID is prefixed by the caller's slug.
	&PermissionRoleCreate,

	// Anyone can create and edit their own permissions
	// given the permission starts with their slug
	&PermissionPermissionCreate,
	&PermissionPermissionEdit,

	// Anyone can create their own organizations and manage users there.
	// Organization
	&PermissionOrganizationCreate,
	&PermissionOrganizationAddUser,
	&PermissionOrganizationRemoveUser,

	&PermissionInviteEdit,
	&PermissionInviteView,
}

var AdminPermissions = []*Permission{
	&PermissionUserCreate,
	&PermissionUserView,
	&PermissionUserEdit,
	&PermissionUserDelete,
	&PermissionRoleCreate,
	&PermissionRoleEdit,
	&PermissionRoleView,
	&PermissionRoleDelete,
	&PermissionPermissionCreate,
	&PermissionPermissionEdit,
	&PermissionOrganizationCreate,
	&PermissionOrganizationAddUser,
	&PermissionOrganizationRemoveUser,
	&PermissionGrantView,
	&PermissionGrantCreate,
}
