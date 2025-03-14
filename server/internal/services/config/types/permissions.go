/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package config_svc

import (
	openapi "github.com/openorch/openorch/clients/go"
)

var PermissionConfigCreate = openapi.UserSvcPermission{
	Id:   openapi.PtrString("config-svc:config:create"),
	Name: openapi.PtrString("Config Svc - Config Create"),
}

// There is no view permission as configs are public facing,
// can be read even without logging in

var PermissionConfigEdit = openapi.UserSvcPermission{
	Id:   openapi.PtrString("config-svc:config:edit"),
	Name: openapi.PtrString("Config Svc - Config Edit"),
}

var PermissionConfigDelete = openapi.UserSvcPermission{
	Id:   openapi.PtrString("config-svc:config:delete"),
	Name: openapi.PtrString("Config Svc - Delete"),
}

var PermissionConfigStream = openapi.UserSvcPermission{
	Id:   openapi.PtrString("config-svc:config:stream"),
	Name: openapi.PtrString("Config Svc - Config Stream"),
}

var AdminPermissions = []openapi.UserSvcPermission{
	PermissionConfigCreate,
	PermissionConfigEdit,
	PermissionConfigDelete,
	PermissionConfigStream,
}
