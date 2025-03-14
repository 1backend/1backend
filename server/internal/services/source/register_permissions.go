/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package sourceservice

import (
	"context"

	openapi "github.com/openorch/openorch/clients/go"
	sdk "github.com/openorch/openorch/sdk/go"
	usertypes "github.com/openorch/openorch/server/internal/services/user/types"

	sourcetypes "github.com/openorch/openorch/server/internal/services/source/types"
)

func (ns *SourceService) registerPermissions() error {
	ctx := context.Background()
	userSvc := ns.clientFactory.Client(sdk.WithToken(ns.token)).UserSvcAPI

	_, _, err := userSvc.SavePermissions(ctx).
		Body(openapi.UserSvcSavePermissionsRequest{
			Permissions: sourcetypes.AdminPermissions,
		}).
		Execute()
	if err != nil {
		return err
	}

	req := openapi.UserSvcAssignPermissionsRequest{}

	for _, role := range []*usertypes.Role{
		usertypes.RoleAdmin,
	} {
		for _, permission := range sourcetypes.AdminPermissions {
			req.PermissionLinks = append(req.PermissionLinks, openapi.UserSvcPermissionLink{
				RoleId:       openapi.PtrString(role.Id),
				PermissionId: permission.Id,
			})
		}
	}

	_, _, err = userSvc.AssignPermissions(ctx).
		Body(req).
		Execute()
	if err != nil {
		return err
	}

	return nil
}
