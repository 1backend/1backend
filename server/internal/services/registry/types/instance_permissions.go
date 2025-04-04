/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package registry_svc

var (
	// Registry Instance Permissions
	PermissionInstanceCreate = "registry-svc:instance:create"
	PermissionInstanceView   = "registry-svc:instance:view"
	PermissionInstanceEdit   = "registry-svc:instance:edit"
	PermissionInstanceDelete = "registry-svc:instance:delete"

	// Instance Permission Groups
	InstanceUserPermissions = []string{
		PermissionInstanceView,
		PermissionInstanceEdit,
	}

	InstanceAdminPermissions = []string{
		PermissionInstanceView,
		PermissionInstanceCreate,
		PermissionInstanceDelete,
	}
)
