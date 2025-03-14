/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package chat_svc

import (
	openapi "github.com/openorch/openorch/clients/go"
)

var PermissionMessageCreate = openapi.UserSvcPermission{
	Id:   openapi.PtrString("chat-svc:message:create"),
	Name: openapi.PtrString("Message Create"),
}

var PermissionMessageView = openapi.UserSvcPermission{
	Id:   openapi.PtrString("chat-svc:message:view"),
	Name: openapi.PtrString("Message View"),
}

var PermissionMessageEdit = openapi.UserSvcPermission{
	Id:   openapi.PtrString("chat-svc:message:edit"),
	Name: openapi.PtrString("Message Edit"),
}

var PermissionMessageDelete = openapi.UserSvcPermission{
	Id:   openapi.PtrString("chat-svc:message:delete"),
	Name: openapi.PtrString("Message Delete"),
}

var PermissionMessageStream = openapi.UserSvcPermission{
	Id:   openapi.PtrString("chat-svc:message:stream"),
	Name: openapi.PtrString("Message Stream"),
}

var MessagePermissions = []openapi.UserSvcPermission{
	PermissionMessageCreate,
	PermissionMessageView,
	PermissionMessageEdit,
	PermissionMessageDelete,
	PermissionMessageStream,
}
