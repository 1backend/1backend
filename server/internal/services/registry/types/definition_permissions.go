/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package registry_svc

var (
	// Registry Service - Definition Permissions
	PermissionDefinitionCreate = "registry-svc:definition:create"
	PermissionDefinitionView   = "registry-svc:definition:view"
	PermissionDefinitionEdit   = "registry-svc:definition:edit"
	PermissionDefinitionDelete = "registry-svc:definition:delete"

	// Definition Permission Groups
	DefinitionUserPermissions = []string{
		PermissionDefinitionView,
	}

	DefinitionAdminPermissions = []string{
		PermissionDefinitionCreate,
		PermissionDefinitionView,
		PermissionDefinitionEdit,
		PermissionDefinitionDelete,
	}
)
