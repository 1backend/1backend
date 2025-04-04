/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package container_svc

var (
	// Container Service Permissions
	PermissionContainerCreate = "container-svc:container:create"
	PermissionContainerView   = "container-svc:container:view"
	PermissionContainerEdit   = "container-svc:container:edit"
	PermissionContainerStop   = "container-svc:container:stop"
	PermissionImageBuild      = "container-svc:image:build"
	PermissionLogView         = "container-svc:log:view"

	// Admin Permission Group
	AdminPermissions = []string{
		PermissionContainerCreate,
		PermissionContainerView,
		PermissionContainerEdit,
		PermissionContainerStop,
		PermissionImageBuild,
		PermissionLogView,
	}
)
