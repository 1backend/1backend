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

var PermissionDefinitionCreate = openapi.UserSvcPermission{
	Id:   openapi.PtrString("registry-svc:definition:create"),
	Name: openapi.PtrString("Registry Svc - Definition Create"),
}

var PermissionDefinitionView = openapi.UserSvcPermission{
	Id:   openapi.PtrString("registry-svc:definition:view"),
	Name: openapi.PtrString("Registry Svc - Definition View"),
}

var PermissionDefinitionEdit = openapi.UserSvcPermission{
	Id:   openapi.PtrString("registry-svc:definition:edit"),
	Name: openapi.PtrString("Registry Svc - Definition Edit"),
}

var PermissionDefinitionDelete = openapi.UserSvcPermission{
	Id:   openapi.PtrString("registry-svc:definition:delete"),
	Name: openapi.PtrString("Registry Svc - Definition Delete"),
}

var DefinitionUserPermissions = []openapi.UserSvcPermission{
	PermissionDefinitionView,
}

var DefinitionAdminPermissions = []openapi.UserSvcPermission{
	PermissionDefinitionCreate,
	PermissionDefinitionView,
	PermissionDefinitionEdit,
	PermissionDefinitionDelete,
}
