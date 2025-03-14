/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package secret_svc

import (
	openapi "github.com/openorch/openorch/clients/go"
)

var PermissionSecretList = openapi.UserSvcPermission{
	Id:   openapi.PtrString("secret-svc:secret:list"),
	Name: openapi.PtrString("Secret Svc - Secret List"),
}

var PermissionSecretSave = openapi.UserSvcPermission{
	Id:   openapi.PtrString("secret-svc:secret:save"),
	Name: openapi.PtrString("Secret Svc - Secret Save"),
}

var PermissionSecretRemove = openapi.UserSvcPermission{
	Id:   openapi.PtrString("secret-svc:secret:remove"),
	Name: openapi.PtrString("Secret Svc - Secret Remove"),
}

// These sensitive looking permissions are not
// just for the admins because there is custom authorization
// (eg. 'readers', 'writers', 'deleters' etc) in the endpoints.
var Permissions = []openapi.UserSvcPermission{
	PermissionSecretList,
	PermissionSecretSave,
	PermissionSecretRemove,
}
