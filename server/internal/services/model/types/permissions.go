/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package model_svc

var (
	// Model Service - Model Permissions
	PermissionModelCreate = "model-svc:model:create"
	PermissionModelView   = "model-svc:model:view"
	PermissionModelEdit   = "model-svc:model:edit"
	PermissionModelDelete = "model-svc:model:delete"
	PermissionModelStream = "model-svc:model:stream"

	// Model Service - Platform Permissions
	PermissionPlatformView = "model-svc:platform:view"
	PermissionPlatformEdit = "model-svc:platform:edit"

	// Admin Permission Group
	AdminPermissions = []string{
		PermissionModelCreate,
		PermissionModelView,
		PermissionModelEdit,
		PermissionModelDelete,
		PermissionModelStream,
		PermissionPlatformEdit,
		PermissionPlatformView,
	}
)
