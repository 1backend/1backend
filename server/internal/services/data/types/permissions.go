/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package data_svc

var (
	// Data Service - Object Permissions
	PermissionObjectCreate = "data-svc:object:create"
	PermissionObjectView   = "data-svc:object:view"
	PermissionObjectEdit   = "data-svc:object:edit"
	PermissionObjectDelete = "data-svc:object:delete"
	PermissionObjectStream = "data-svc:object:stream"

	// Object Permission Group
	Permissions = []string{
		PermissionObjectCreate,
		PermissionObjectView,
		PermissionObjectEdit,
		PermissionObjectDelete,
		PermissionObjectStream,
	}
)
