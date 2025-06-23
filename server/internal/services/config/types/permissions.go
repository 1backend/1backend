/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package config_svc

var (
	// Config Service - Config Permissions
	PermissionConfigCreate       = "config-svc:config:create"
	PermissionConfigEdit         = "config-svc:config:edit"
	PermissionConfigEditOnBehalf = "config-svc:config:edit-on-behalf"
	PermissionConfigDelete       = "config-svc:config:delete"
	PermissionConfigStream       = "config-svc:config:stream"

	UserPermissions = []string{
		PermissionConfigEdit,
	}

	// Admin Permission Group
	AdminPermissions = []string{
		PermissionConfigCreate,
		PermissionConfigEdit,
		PermissionConfigDelete,
		PermissionConfigStream,
		PermissionConfigEditOnBehalf,
	}
)
