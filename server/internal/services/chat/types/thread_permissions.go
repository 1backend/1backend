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

var PermissionThreadCreate = openapi.UserSvcPermission{
	Id:   openapi.PtrString("chat-svc:thread:create"),
	Name: openapi.PtrString("Chat Svc - Thread Create"),
}

var PermissionThreadView = openapi.UserSvcPermission{
	Id:   openapi.PtrString("chat-svc:thread:view"),
	Name: openapi.PtrString("Chat Svc - Thread View"),
}

var PermissionThreadEdit = openapi.UserSvcPermission{
	Id:   openapi.PtrString("chat-svc:thread:edit"),
	Name: openapi.PtrString("Chat Svc - Thread Edit"),
}

var PermissionThreadDelete = openapi.UserSvcPermission{
	Id:   openapi.PtrString("chat-svc:thread:delete"),
	Name: openapi.PtrString("Chat Svc - Thread Delete"),
}

var PermissionThreadStream = openapi.UserSvcPermission{
	Id:   openapi.PtrString("chat-svc:thread:stream"),
	Name: openapi.PtrString("Chat Svc - Thread Stream"),
}

var ThreadPermissions = []openapi.UserSvcPermission{
	PermissionThreadCreate,
	PermissionThreadView,
	PermissionThreadEdit,
	PermissionThreadDelete,
	PermissionThreadStream,
}
