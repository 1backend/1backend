/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package deploy_svc

import (
	openapi "github.com/openorch/openorch/clients/go"
)

var PermissionDeploymentCreate = openapi.UserSvcPermission{
	Id:   openapi.PtrString("deploy-svc:deployment:create"),
	Name: openapi.PtrString("Deploy Svc - Deployment Create"),
}

var PermissionDeploymentView = openapi.UserSvcPermission{
	Id:   openapi.PtrString("deploy-svc:deployment:view"),
	Name: openapi.PtrString("Deploy Svc - Deployment View"),
}

var PermissionDeploymentEdit = openapi.UserSvcPermission{
	Id:   openapi.PtrString("deploy-svc:deployment:create"),
	Name: openapi.PtrString("Deploy Svc - Deployment Create"),
}

var PermissionDeploymentDelete = openapi.UserSvcPermission{
	Id:   openapi.PtrString("deploy-svc:deployment:delete"),
	Name: openapi.PtrString("Deploy Svc - Deployment Delete"),
}

var AdminPermissions = []openapi.UserSvcPermission{
	PermissionDeploymentCreate,
	PermissionDeploymentEdit,
	PermissionDeploymentView,
	PermissionDeploymentDelete,
}
