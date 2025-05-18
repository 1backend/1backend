/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package user_svc

type HasPermissionResponse struct {
	Authorized bool `json:"authorized" binding:"required"`
	User       User `json:"user" binding:"required"`
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
	PermissionUserPasswordChange = "user-svc:user:password-change"
	PermissionUserPasswordReset  = "user-svc:user:password-reset"

	// Organization Permissions
	PermissionOrganizationCreate     = "user-svc:organization:create"
	PermissionOrganizationAddUser    = "user-svc:organization:add-user"
	PermissionOrganizationRemoveUser = "user-svc:organization:remove-user"
	PermissionOrganizationView       = "user-svc:organization:view"

	// Permit Permissions
	PermissionPermitCreate = "user-svc:permit:create"
	PermissionPermitView   = "user-svc:permit:view"

	// Enroll Permissions
	PermissionEnrollEdit = "user-svc:enroll:edit"
	PermissionEnrollView = "user-svc:enroll:view"

	PermissionTokenRevoke = "user-svc:token:revoke"
)

var UserPermissions = []string{
	// Anyone can change their own password
	PermissionUserPasswordChange,

	// Anyone can create their own organizations
	// and manage users there.
	PermissionOrganizationCreate,
	PermissionOrganizationAddUser,
	PermissionOrganizationRemoveUser,

	// A User can edit their own enrolls
	PermissionEnrollEdit,
	// A User can view their own enrolls
	PermissionEnrollView,
}

var AdminPermissions = []string{
	PermissionUserCreate,
	PermissionUserView,
	PermissionUserEdit,
	PermissionUserDelete,
	PermissionOrganizationCreate,
	PermissionOrganizationAddUser,
	PermissionOrganizationRemoveUser,
	PermissionOrganizationView,
	PermissionPermitCreate,
	PermissionPermitView,
	PermissionEnrollEdit,
	PermissionEnrollView,
	PermissionUserPasswordChange,
	PermissionUserPasswordReset,
	PermissionTokenRevoke,
}
