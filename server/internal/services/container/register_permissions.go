/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package containerservice

import (
	"context"

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	dockertypes "github.com/1backend/1backend/server/internal/services/container/types"
	usertypes "github.com/1backend/1backend/server/internal/services/user/types"
)

func (p *ContainerService) registerPermits() error {
	ctx := context.Background()
	userSvc := p.options.ClientFactory.Client(client.WithToken(p.token)).UserSvcAPI

	req := openapi.UserSvcSavePermitsRequest{
		Permits: []openapi.UserSvcPermitInput{
			{
				App:        openapi.PtrString("*"),
				Slugs:      []string{"model-svc"},
				Permission: dockertypes.PermissionContainerView,
			},
			{
				App:        openapi.PtrString("*"),
				Slugs:      []string{"model-svc", "deploy-svc"},
				Permission: dockertypes.PermissionLogView,
			},
			{
				App:        openapi.PtrString("*"),
				Slugs:      []string{"model-svc", "deploy-svc"},
				Permission: dockertypes.PermissionContainerCreate,
			},
			{
				App:        openapi.PtrString("*"),
				Slugs:      []string{"model-svc", "deploy-svc"},
				Permission: dockertypes.PermissionContainerCreate,
			},
			{
				App:        openapi.PtrString("*"),
				Slugs:      []string{"model-svc", "deploy-svc"},
				Permission: dockertypes.PermissionContainerStop,
			},
			{
				App:        openapi.PtrString("*"),
				Slugs:      []string{"deploy-svc"},
				Permission: dockertypes.PermissionImageBuild,
			},
		},
	}

	for _, role := range []string{
		usertypes.RoleAdmin,
	} {
		for _, permission := range dockertypes.AdminPermissions {
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
