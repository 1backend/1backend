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

var PermissionInstanceCreate = openapi.UserSvcPermission{
	Id:   openapi.PtrString("registry-svc:instance:create"),
	Name: openapi.PtrString("Registry Svc - Instance Create"),
}

var PermissionInstanceView = openapi.UserSvcPermission{
	Id:   openapi.PtrString("registry-svc:instance:view"),
	Name: openapi.PtrString("Registry Svc - Instance View"),
}

var PermissionInstanceEdit = openapi.UserSvcPermission{
	Id:   openapi.PtrString("registry-svc:instance:edit"),
	Name: openapi.PtrString("Registry Svc - Instance Edit"),
}

var PermissionInstanceDelete = openapi.UserSvcPermission{
	Id:   openapi.PtrString("registry-svc:instance:delete"),
	Name: openapi.PtrString("Registry Svc - Instance Delete"),
}

var InstanceUserPermissions = []openapi.UserSvcPermission{
	PermissionInstanceView,
	PermissionInstanceEdit,
}

var InstanceAdminPermissions = []openapi.UserSvcPermission{
	PermissionInstanceView,
	PermissionInstanceCreate,
	PermissionInstanceDelete,
}
