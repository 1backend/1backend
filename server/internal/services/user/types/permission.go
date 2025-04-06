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

type ListPermissionsRequest struct {
	Roles []string `json:"roles"`
}

type ListPermissionsResponse struct {
	Permissions []string `json:"permissions"`
}

var (
	// User Permissions
	PermissionUserCreate         = "user-svc:user:create"
	PermissionUserView           = "user-svc:user:view"
	PermissionUserEdit           = "user-svc:user:edit"
	PermissionUserDelete         = "user-svc:user:delete"
	PermissionUserPasswordChange = "user-svc:user:passwordChange"

	// Organization Permissions
	PermissionOrganizationCreate     = "user-svc:organization:create"
	PermissionOrganizationAddUser    = "user-svc:organization:add-user"
	PermissionOrganizationRemoveUser = "user-svc:organization:remove-user"

	// Grant Permissions
	PermissionGrantCreate = "user-svc:grant:create"

	// Invite Permissions
	PermissionInviteEdit = "user-svc:invite:edit"
	PermissionInviteView = "user-svc:invite:view"
)

var UserPermissions = []string{
	PermissionUserPasswordChange,

	// Anyone can create their own organizations and manage users there.
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
	PermissionOrganizationCreate,
	PermissionOrganizationAddUser,
	PermissionOrganizationRemoveUser,
	PermissionGrantCreate,
	PermissionInviteEdit,
	PermissionInviteView,
}
