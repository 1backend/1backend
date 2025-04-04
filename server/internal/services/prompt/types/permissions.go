/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package prompt_svc

var (
	// Prompt Service - Prompt Permissions
	PermissionPromptCreate = "prompt-svc:prompt:create"
	PermissionPromptView   = "prompt-svc:prompt:view"
	PermissionPromptEdit   = "prompt-svc:prompt:edit"
	PermissionPromptDelete = "prompt-svc:prompt:delete"
	PermissionPromptStream = "prompt-svc:prompt:stream"

	// Admin Permission Group
	AdminPermissions = []string{
		PermissionPromptCreate,
		PermissionPromptView,
		PermissionPromptEdit,
		PermissionPromptDelete,
		PermissionPromptStream,
	}
)
