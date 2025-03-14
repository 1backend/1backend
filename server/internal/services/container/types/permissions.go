/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package container_svc

import (
	openapi "github.com/openorch/openorch/clients/go"
)

var PermissionContainerCreate = openapi.UserSvcPermission{
	Id:   openapi.PtrString("container-svc:container:create"),
	Name: openapi.PtrString("Container Svc - Container Create"),
}

var PermissionContainerView = openapi.UserSvcPermission{
	Id:   openapi.PtrString("container-svc:container:view"),
	Name: openapi.PtrString("Container Svc - Container View"),
}

var PermissionContainerEdit = openapi.UserSvcPermission{
	Id:   openapi.PtrString("container-svc:container:edit"),
	Name: openapi.PtrString("Container Svc - Container Edit"),
}

var PermissionContainerStop = openapi.UserSvcPermission{
	Id:   openapi.PtrString("container-svc:container:stop"),
	Name: openapi.PtrString("Container Svc - Container Stop"),
}

var PermissionImageBuild = openapi.UserSvcPermission{
	Id:   openapi.PtrString("container-svc:image:build"),
	Name: openapi.PtrString("Container Svc - Image Build"),
}

var PermissionLogView = openapi.UserSvcPermission{
	Id:   openapi.PtrString("container-svc:log:view"),
	Name: openapi.PtrString("Container Svc - Log View"),
}

var AdminPermissions = []openapi.UserSvcPermission{
	PermissionContainerCreate,
	PermissionContainerView,
	PermissionContainerEdit,
	PermissionContainerStop,
	PermissionImageBuild,
	PermissionLogView,
}
