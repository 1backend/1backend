/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package secret_svc

var (
	// Secret Service - Secret Permissions

	PermissionSecretList           = "secret-svc:secret:list"
	PermissionSecretSave           = "secret-svc:secret:save"
	PermissionSecretSaveUnprefixed = "secret-svc:secret:save-unprefixed"
	PermissionSecretRemove         = "secret-svc:secret:remove"

	// Secret Permissions Group
	UserPermissions = []string{
		PermissionSecretList,
		PermissionSecretSave,
		PermissionSecretRemove,
	}

	AdminPermissions = []string{
		PermissionSecretSaveUnprefixed,
	}
)
