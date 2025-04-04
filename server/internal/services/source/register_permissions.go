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

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	usertypes "github.com/1backend/1backend/server/internal/services/user/types"

	sourcetypes "github.com/1backend/1backend/server/internal/services/source/types"
)

func (ns *SourceService) registerPermissions() error {
	ctx := context.Background()
	userSvc := ns.clientFactory.Client(client.WithToken(ns.token)).UserSvcAPI

	req := openapi.UserSvcAssignPermissionsRequest{}

	for _, role := range []string{
		usertypes.RoleAdmin,
	} {
		for _, permission := range sourcetypes.AdminPermissions {
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
