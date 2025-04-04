/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package dynamicservice

import (
	"context"

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	data "github.com/1backend/1backend/server/internal/services/data/types"
	usertypes "github.com/1backend/1backend/server/internal/services/user/types"
)

func (p *DataService) registerPermissions() error {
	ctx := context.Background()
	userSvc := p.clientFactory.Client(client.WithToken(p.token)).UserSvcAPI

	req := openapi.UserSvcAssignPermissionsRequest{}

	for _, role := range []string{
		usertypes.RoleAdmin,
		usertypes.RoleUser,
	} {
		for _, permission := range data.Permissions {
			req.PermissionLinks = append(req.PermissionLinks, openapi.UserSvcPermissionLink{
				Role:       role,
				Permission: permission,
			})
		}
	}

	_, _, err := userSvc.AssignPermissions(ctx).
		Body(req).
		Execute()
	if err != nil {
		return err
	}

	return nil
}
