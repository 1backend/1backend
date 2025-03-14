/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package secretservice

import (
	"context"

	openapi "github.com/openorch/openorch/clients/go"
	sdk "github.com/openorch/openorch/sdk/go"
	secrettypes "github.com/openorch/openorch/server/internal/services/secret/types"
	usertypes "github.com/openorch/openorch/server/internal/services/user/types"
)

func (p *SecretService) registerPermissions() error {
	ctx := context.Background()
	userSvc := p.clientFactory.Client(sdk.WithToken(p.token)).UserSvcAPI

	_, _, err := userSvc.SavePermissions(ctx).
		Body(openapi.UserSvcSavePermissionsRequest{
			Permissions: secrettypes.Permissions,
		}).
		Execute()
	if err != nil {
		return err
	}

	req := openapi.UserSvcAssignPermissionsRequest{}

	for _, role := range []*usertypes.Role{
		usertypes.RoleAdmin,
		// THIS IS CORRECT, DO NOT PANIC
		// Users can list, save and delete secrets if they
		// are authorized to do so by the "readers", "writers", "deleters"
		// list of a secret.
		usertypes.RoleUser,
	} {
		for _, permission := range secrettypes.Permissions {
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
