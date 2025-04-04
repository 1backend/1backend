/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package chat_svc

var (
	// Chat Service - Thread Permissions
	PermissionThreadCreate = "chat-svc:thread:create"
	PermissionThreadView   = "chat-svc:thread:view"
	PermissionThreadEdit   = "chat-svc:thread:edit"
	PermissionThreadDelete = "chat-svc:thread:delete"
	PermissionThreadStream = "chat-svc:thread:stream"

	// Thread Permission Group
	ThreadPermissions = []string{
		PermissionThreadCreate,
		PermissionThreadView,
		PermissionThreadEdit,
		PermissionThreadDelete,
		PermissionThreadStream,
	}
)
