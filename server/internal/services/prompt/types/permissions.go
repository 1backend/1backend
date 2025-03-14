/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package prompt_svc

import (
	openapi "github.com/openorch/openorch/clients/go"
)

var PermissionPromptCreate = openapi.UserSvcPermission{
	Id:   openapi.PtrString("prompt-svc:prompt:create"),
	Name: openapi.PtrString("Prompt Svc - Prompt Create"),
}

var PermissionPromptView = openapi.UserSvcPermission{
	Id:   openapi.PtrString("prompt-svc:prompt:view"),
	Name: openapi.PtrString("Prompt Svc - Prompt View"),
}

var PermissionPromptEdit = openapi.UserSvcPermission{
	Id:   openapi.PtrString("prompt-svc:prompt:edit"),
	Name: openapi.PtrString("Prompt Svc - Prompt Edit"),
}

var PermissionPromptDelete = openapi.UserSvcPermission{
	Id:   openapi.PtrString("prompt-svc:prompt:delete"),
	Name: openapi.PtrString("Prompt Svc - Prompt Delete"),
}

var PermissionPromptStream = openapi.UserSvcPermission{
	Id:   openapi.PtrString("prompt-svc:prompt:stream"),
	Name: openapi.PtrString("Prompt Svc - Prompt Stream"),
}

var AdminPermissions = []openapi.UserSvcPermission{
	PermissionPromptCreate,
	PermissionPromptView,
	PermissionPromptEdit,
	PermissionPromptDelete,
	PermissionPromptStream,
}
