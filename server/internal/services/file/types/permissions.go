/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package file_svc

import (
	openapi "github.com/openorch/openorch/clients/go"
)

var PermissionDownloadCreate = openapi.UserSvcPermission{
	Id:   openapi.PtrString("file-svc:download:create"),
	Name: openapi.PtrString("File Svc - Download Create"),
}

var PermissionDownloadView = openapi.UserSvcPermission{
	Id:   openapi.PtrString("file-svc:download:view"),
	Name: openapi.PtrString("File Svc - Download View"),
}

var PermissionDownloadEdit = openapi.UserSvcPermission{
	Id:   openapi.PtrString("file-svc:download:edit"),
	Name: openapi.PtrString("File Svc - Download Edit"),
}

var PermissionDownloadDelete = openapi.UserSvcPermission{
	Id:   openapi.PtrString("file-svc:download:delete"),
	Name: openapi.PtrString("File Svc - Download Delete"),
}

var PermissionUploadCreate = openapi.UserSvcPermission{
	Id:   openapi.PtrString("file-svc:upload:create"),
	Name: openapi.PtrString("File Svc - Upload Create"),
}

var PermissionUploadView = openapi.UserSvcPermission{
	Id:   openapi.PtrString("file-svc:upload:view"),
	Name: openapi.PtrString("File Svc - Upload View"),
}

var AdminPermissions = []openapi.UserSvcPermission{
	PermissionDownloadCreate,
	PermissionDownloadView,
	PermissionDownloadEdit,
	PermissionDownloadDelete,
	PermissionUploadCreate,
	PermissionUploadView,
}
