/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package chat_svc

var (
	// Chat Service - Message Permissions
	PermissionMessageCreate = "chat-svc:message:create"
	PermissionMessageView   = "chat-svc:message:view"
	PermissionMessageEdit   = "chat-svc:message:edit"
	PermissionMessageDelete = "chat-svc:message:delete"
	PermissionMessageStream = "chat-svc:message:stream"

	// Message Permission Group
	MessagePermissions = []string{
		PermissionMessageCreate,
		PermissionMessageView,
		PermissionMessageEdit,
		PermissionMessageDelete,
		PermissionMessageStream,
	}
)
