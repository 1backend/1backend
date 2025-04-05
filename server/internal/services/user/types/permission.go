/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package user_svc

type HasPermissionRequest struct {
	GrantedSlugs    []string `json:"grantedSlugs,omitempty"`
	ContactsGranted []string `json:"contactsGranted,omitempty"`
}

type HasPermissionResponse struct {
	Authorized bool  `json:"authorized,omitempty"`
	User       *User `json:"user,omitempty"`
}

type CreatePermissionRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreatePermissionResponse struct {
}

type ListPermissionsRequest struct {
	Roles []string `json:"roles"`
}

type ListPermissionsResponse struct {
	Permissions []string `json:"permissions"`
}

type SavePermissionsRequest struct {
	Permissions []string `json:"permissions"`
}

type SavePermissionsResponse struct {
	Permissions []string `json:"permissions"`
}

type PermissionLink struct {
	Role       string `json:"role" binding:"required"`
	Permission string `json:"permission" binding:"required"`
}

type AssignPermissionsRequest struct {
	PermissionLinks []*PermissionLink `json:"permissionLinks"`
}

type AssignPermissionsResponse struct{}

var (
	// User Permissions
	PermissionUserCreate         = "user-svc:user:create"
	PermissionUserView           = "user-svc:user:view"
	PermissionUserEdit           = "user-svc:user:edit"
	PermissionUserDelete         = "user-svc:user:delete"
	PermissionUserPasswordChange = "user-svc:user:passwordChange"

	// Role Permissions
	PermissionRoleCreate = "user-svc:role:create"
	PermissionRoleView   = "user-svc:role:view"
	PermissionRoleEdit   = "user-svc:role:edit"
	PermissionRoleDelete = "user-svc:role:delete"

	// Permission Permissions
	PermissionPermissionCreate = "user-svc:permission:create"
	PermissionPermissionEdit   = "user-svc:permission:edit"

	// Organization Permissions
	PermissionOrganizationCreate     = "user-svc:organization:create"
	PermissionOrganizationAddUser    = "user-svc:organization:add-user"
	PermissionOrganizationRemoveUser = "user-svc:organization:remove-user"

	// Grant Permissions
	PermissionGrantCreate = "user-svc:grant:create"
	PermissionGrantView   = "user-svc:grant:view"

	// Invite Permissions
	PermissionInviteEdit = "user-svc:invite:edit"
	PermissionInviteView = "user-svc:invite:view"
)

var UserPermissions = []string{
	PermissionUserPasswordChange,

	// Anyone can create and edit their own roles,
	// provided the role ID is prefixed by the caller's slug.
	PermissionRoleCreate,

	// Anyone can create and edit their own permissions
	// given the permission starts with their slug
	PermissionPermissionCreate,
	PermissionPermissionEdit,

	// Anyone can create their own organizations and manage users there.
	// Organization
	PermissionOrganizationCreate,
	PermissionOrganizationAddUser,
	PermissionOrganizationRemoveUser,

	PermissionInviteEdit,
	PermissionInviteView,
}

var AdminPermissions = []string{
	PermissionUserCreate,
	PermissionUserView,
	PermissionUserEdit,
	PermissionUserDelete,
	PermissionRoleCreate,
	PermissionRoleEdit,
	PermissionRoleView,
	PermissionRoleDelete,
	PermissionPermissionCreate,
	PermissionPermissionEdit,
	PermissionOrganizationCreate,
	PermissionOrganizationAddUser,
	PermissionOrganizationRemoveUser,
	PermissionGrantView,
	PermissionGrantCreate,
}
