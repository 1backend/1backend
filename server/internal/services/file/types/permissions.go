/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package file_svc

var (
	// File Service - Download Permissions
	PermissionDownloadCreate = "file-svc:download:create"
	PermissionDownloadView   = "file-svc:download:view"
	PermissionDownloadEdit   = "file-svc:download:edit"
	PermissionDownloadDelete = "file-svc:download:delete"

	// File Service - Upload Permissions
	PermissionUploadCreate = "file-svc:upload:create"
	PermissionUploadView   = "file-svc:upload:view"

	// Admin Permission Group
	AdminPermissions = []string{
		PermissionDownloadCreate,
		PermissionDownloadView,
		PermissionDownloadEdit,
		PermissionDownloadDelete,
		PermissionUploadCreate,
		PermissionUploadView,
	}
)
