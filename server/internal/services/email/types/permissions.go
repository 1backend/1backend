/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This email code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package email_svc

import (
	openapi "github.com/openorch/openorch/clients/go"
)

var PermissionSendEmail = openapi.UserSvcPermission{
	Id:   openapi.PtrString("email-svc:email:send"),
	Name: openapi.PtrString("Email Svc - Email Send"),
}

var AdminPermissions = []openapi.UserSvcPermission{
	PermissionSendEmail,
}
