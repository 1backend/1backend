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

	req := openapi.UserSvcSaveGrantsRequest{}

	for _, role := range []string{
		usertypes.RoleAdmin,
		usertypes.RoleUser,
	} {
		for _, permission := range append(chattypes.ThreadPermissions, chattypes.MessagePermissions...) {
			req.Grants = append(req.Grants, openapi.UserSvcGrant{
				Roles:      []string{role},
				Permission: permission,
			})
		}
	}

	_, _, err := userSvc.SaveGrants(ctx).
		Body(req).
		Execute()
	if err != nil {
		return err
	}

	return nil
}
