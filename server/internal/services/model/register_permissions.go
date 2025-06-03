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

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/datastore"
	model "github.com/1backend/1backend/server/internal/services/model/types"
	usertypes "github.com/1backend/1backend/server/internal/services/user/types"
)

func (p *ModelService) registerPermits() error {
	ctx := context.Background()
	userSvc := p.options.ClientFactory.Client(client.WithToken(p.token)).UserSvcAPI

	req := openapi.UserSvcSavePermitsRequest{
		Permits: []openapi.UserSvcPermitInput{
			{
				App:        openapi.PtrString("*"),
				Slugs:      []string{"prompt-svc"},
				Permission: model.PermissionModelView,
			},
			{
				App:        openapi.PtrString("*"),
				Slugs:      []string{"prompt-svc"},
				Permission: model.PermissionPlatformView,
			},
		},
	}

	for _, role := range []string{
		usertypes.RoleAdmin,
	} {
		for _, permission := range model.AdminPermissions {
			req.Permits = append(req.Permits, openapi.UserSvcPermitInput{
				App:        openapi.PtrString("*"),
				Roles:      []string{role},
				Permission: permission,
			})
		}
	}

	_, _, err := userSvc.SavePermits(ctx).
		Body(req).
		Execute()
	if err != nil {
		return err
	}

	return nil
}

func (p *ModelService) bootstrapModels() error {
	platformRows := []datastore.Row{}
	for _, v := range model.Platforms {
		platformRows = append(platformRows, v)
	}
	err := p.platformsStore.UpsertMany(platformRows)
	if err != nil {
		return err
	}

	modelRows := []datastore.Row{}
	for _, v := range model.Models {
		modelRows = append(modelRows, v)
	}

	return p.modelsStore.UpsertMany(modelRows)

}
