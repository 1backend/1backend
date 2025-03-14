/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package firehose_svc

import (
	openapi "github.com/openorch/openorch/clients/go"
)

var PermissionEventPublish = openapi.UserSvcPermission{
	Id:   openapi.PtrString("firehose-svc:event:publish"),
	Name: openapi.PtrString("Firehose Svc - Event Publish"),
}

var PermissionFirehoseStream = openapi.UserSvcPermission{
	Id:   openapi.PtrString("firehose-svc:event:stream"),
	Name: openapi.PtrString("Firehose Svc - Event Stream"),
}

var AdminPermissions = []openapi.UserSvcPermission{
	PermissionEventPublish,
	PermissionFirehoseStream,
}
