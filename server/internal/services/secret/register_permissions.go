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

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	secrettypes "github.com/1backend/1backend/server/internal/services/secret/types"
	usertypes "github.com/1backend/1backend/server/internal/services/user/types"
)

func (p *SecretService) registerPermits() error {
	ctx := context.Background()
	userSvc := p.options.ClientFactory.Client(client.WithToken(p.token)).UserSvcAPI

	req := openapi.UserSvcSavePermitsRequest{
		Permits: []openapi.UserSvcPermitInput{
			{
				App:        openapi.PtrString("*"),
				Slugs:      []string{"model-svc"},
				Permission: secrettypes.PermissionSecretSave,
			},
		},
	}

	for _, role := range []string{
		usertypes.RoleAdmin,
		// THIS IS CORRECT, DO NOT PANIC
		// Users can list, save and delete secrets if they
		// are authorized to do so by the "readers", "writers", "deleters"
		// list of a secret.
		usertypes.RoleUser,
	} {
		for _, permission := range secrettypes.UserPermissions {
			req.Permits = append(req.Permits, openapi.UserSvcPermitInput{
				App:        openapi.PtrString("*"),
				Roles:      []string{role},
				Permission: permission,
			})
		}
	}

	for _, role := range []string{
		usertypes.RoleAdmin,
	} {
		for _, permission := range secrettypes.AdminPermissions {
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
