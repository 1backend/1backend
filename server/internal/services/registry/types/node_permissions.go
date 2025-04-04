/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package registry_svc

var (
	// Registry Node Permissions
	PermissionNodeCreate = "registry-svc:node:create"
	PermissionNodeView   = "registry-svc:node:view"
	PermissionNodeEdit   = "registry-svc:node:edit"
	PermissionNodeDelete = "registry-svc:node:delete"
	PermissionNodeStream = "registry-svc:node:stream"

	// Node Permission Groups
	NodeAdminPermissions = []string{
		PermissionNodeCreate,
		PermissionNodeView,
		PermissionNodeEdit,
		PermissionNodeDelete,
		PermissionNodeStream,
	}
)
