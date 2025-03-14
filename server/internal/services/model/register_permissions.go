/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package modelservice

import (
	"context"

	openapi "github.com/openorch/openorch/clients/go"
	sdk "github.com/openorch/openorch/sdk/go"
	"github.com/openorch/openorch/sdk/go/datastore"
	modeltypes "github.com/openorch/openorch/server/internal/services/model/types"
	usertypes "github.com/openorch/openorch/server/internal/services/user/types"
)

func (p *ModelService) registerPermissions() error {
	ctx := context.Background()
	userSvc := p.clientFactory.Client(sdk.WithToken(p.token)).UserSvcAPI

	_, _, err := userSvc.SavePermissions(ctx).
		Body(openapi.UserSvcSavePermissionsRequest{
			Permissions: modeltypes.AdminPermissions,
		}).
		Execute()
	if err != nil {
		return err
	}

	req := openapi.UserSvcAssignPermissionsRequest{}

	for _, role := range []*usertypes.Role{
		usertypes.RoleAdmin,
	} {
		for _, permission := range modeltypes.AdminPermissions {
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

func (p *ModelService) bootstrapModels() error {
	platformRows := []datastore.Row{}
	for _, v := range modeltypes.Platforms {
		platformRows = append(platformRows, v)
	}
	err := p.platformsStore.UpsertMany(platformRows)
	if err != nil {
		return err
	}

	modelRows := []datastore.Row{}
	for _, v := range modeltypes.Models {
		modelRows = append(modelRows, v)
	}

	return p.modelsStore.UpsertMany(modelRows)

}
