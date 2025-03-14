/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package model_svc

import (
	openapi "github.com/openorch/openorch/clients/go"
)

var PermissionModelCreate = openapi.UserSvcPermission{
	Id:   openapi.PtrString("model-svc:model:create"),
	Name: openapi.PtrString("Model Svc - Model Create"),
}

var PermissionModelView = openapi.UserSvcPermission{
	Id:   openapi.PtrString("model-svc:model:view"),
	Name: openapi.PtrString("Model Svc - Model View"),
}

var PermissionModelEdit = openapi.UserSvcPermission{
	Id:   openapi.PtrString("model-svc:model:edit"),
	Name: openapi.PtrString("Model Svc - Model Edit"),
}

var PermissionModelDelete = openapi.UserSvcPermission{
	Id:   openapi.PtrString("model-svc:model:delete"),
	Name: openapi.PtrString("Model Svc - Model Delete"),
}

var PermissionModelStream = openapi.UserSvcPermission{
	Id:   openapi.PtrString("model-svc:model:stream"),
	Name: openapi.PtrString("Model Svc - Model Stream"),
}

var PermissionPlatformView = openapi.UserSvcPermission{
	Id:   openapi.PtrString("model-svc:platform:view"),
	Name: openapi.PtrString("Model Svc - Platform View"),
}

var PermissionPlatformEdit = openapi.UserSvcPermission{
	Id:   openapi.PtrString("model-svc:platform:edit"),
	Name: openapi.PtrString("Model Svc - Platform Edit"),
}

var AdminPermissions = []openapi.UserSvcPermission{
	PermissionModelCreate,
	PermissionModelView,
	PermissionModelEdit,
	PermissionModelDelete,
	PermissionModelStream,
	PermissionPlatformEdit,
	PermissionPlatformView,
}
