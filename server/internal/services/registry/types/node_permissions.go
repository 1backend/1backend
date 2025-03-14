/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package registry_svc

import (
	openapi "github.com/openorch/openorch/clients/go"
)

var PermissionNodeCreate = openapi.UserSvcPermission{
	Id:   openapi.PtrString("registry-svc:node:create"),
	Name: openapi.PtrString("Registry Svc - Node Create"),
}

var PermissionNodeView = openapi.UserSvcPermission{
	Id:   openapi.PtrString("registry-svc:node:view"),
	Name: openapi.PtrString("Registry Svc - Node View"),
}

var PermissionNodeEdit = openapi.UserSvcPermission{
	Id:   openapi.PtrString("registry-svc:node:edit"),
	Name: openapi.PtrString("Registry Svc - Node Edit"),
}

var PermissionNodeDelete = openapi.UserSvcPermission{
	Id:   openapi.PtrString("registry-svc:node:delete"),
	Name: openapi.PtrString("Registry Svc - Node Delete"),
}

var PermissionNodeStream = openapi.UserSvcPermission{
	Id:   openapi.PtrString("registry-svc:node:stream"),
	Name: openapi.PtrString("Registry Svc - Node Stream"),
}

var NodeAdminPermissions = []openapi.UserSvcPermission{
	PermissionNodeCreate,
	PermissionNodeView,
	PermissionNodeEdit,
	PermissionNodeDelete,
	PermissionNodeStream,
}
