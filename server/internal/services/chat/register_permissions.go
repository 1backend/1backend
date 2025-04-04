/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package chatservice

import (
	"context"

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	chattypes "github.com/1backend/1backend/server/internal/services/chat/types"
	usertypes "github.com/1backend/1backend/server/internal/services/user/types"
)

func (p *ChatService) registerPermissions() error {
	ctx := context.Background()
	userSvc := p.clientFactory.Client(client.WithToken(p.token)).UserSvcAPI

	_, _, err := userSvc.SavePermissions(ctx).
		Body(openapi.UserSvcSavePermissionsRequest{
			Permissions: append(chattypes.ThreadPermissions, chattypes.MessagePermissions...),
		}).
		Execute()
	if err != nil {
		return err
	}

	req := openapi.UserSvcAssignPermissionsRequest{}

	for _, role := range []*usertypes.Role{
		usertypes.RoleAdmin,
		usertypes.RoleUser,
	} {
		for _, permission := range append(chattypes.ThreadPermissions, chattypes.MessagePermissions...) {
			req.PermissionLinks = append(req.PermissionLinks, openapi.UserSvcPermissionLink{
				RoleId:       role.Id,
				PermissionId: *permission.Id,
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
