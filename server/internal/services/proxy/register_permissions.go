/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package proxyservice

import (
	"context"

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	proxytypes "github.com/1backend/1backend/server/internal/services/proxy/types"
	usertypes "github.com/1backend/1backend/server/internal/services/user/types"
)

func (p *ProxyService) registerPermits() error {
	ctx := context.Background()
	userSvc := p.options.ClientFactory.Client(client.WithToken(p.token)).UserSvcAPI

	req := openapi.UserSvcSavePermitsRequest{}

	for _, role := range []string{
		usertypes.RoleAdmin,
	} {
		for _, permission := range proxytypes.AdminPermissions {
			req.Permits = append(req.Permits, openapi.UserSvcPermitInput{
				AppHost:    openapi.PtrString("*"),
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
