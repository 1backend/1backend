/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package deploy_svc

var (
	// Deploy Service - Deployment Permissions
	PermissionDeploymentCreate = "deploy-svc:deployment:create"
	PermissionDeploymentView   = "deploy-svc:deployment:view"
	PermissionDeploymentEdit   = "deploy-svc:deployment:edit"
	PermissionDeploymentDelete = "deploy-svc:deployment:delete"

	// Admin Permission Group
	AdminPermissions = []string{
		PermissionDeploymentCreate,
		PermissionDeploymentEdit,
		PermissionDeploymentView,
		PermissionDeploymentDelete,
	}
)
