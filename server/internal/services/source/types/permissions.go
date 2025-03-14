/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package source_svc

import (
	openapi "github.com/openorch/openorch/clients/go"
)

var PermissionSourceRepoCheckout = openapi.UserSvcPermission{
	Id:   openapi.PtrString("source-svc:repo:checkout"),
	Name: openapi.PtrString("Source Svc - Repo Checkout"),
}

var AdminPermissions = []openapi.UserSvcPermission{
	PermissionSourceRepoCheckout,
}
