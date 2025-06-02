/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This email code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package emailservice

import (
	"context"

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	usertypes "github.com/1backend/1backend/server/internal/services/user/types"

	email "github.com/1backend/1backend/server/internal/services/email/types"
)

func (ns *EmailService) registerPermits() error {
	ctx := context.Background()
	userSvc := ns.options.ClientFactory.Client(client.WithToken(ns.token)).UserSvcAPI

	req := openapi.UserSvcSavePermitsRequest{
		Permits: []openapi.UserSvcPermitInput{
			{
				App:        openapi.PtrString("*"),
				Slugs:      []string{"user-svc"},
				Permission: email.PermissionSendEmail,
			},
		},
	}

	for _, role := range []string{
		usertypes.RoleAdmin,
	} {
		for _, permission := range email.AdminPermissions {
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
