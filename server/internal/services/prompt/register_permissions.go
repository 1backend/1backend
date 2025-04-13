/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package promptservice

import (
	"context"

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	prompttypes "github.com/1backend/1backend/server/internal/services/prompt/types"
	usertypes "github.com/1backend/1backend/server/internal/services/user/types"
	"github.com/pkg/errors"
)

func (p *PromptService) registerPermissions() error {
	token, err := p.getToken()
	if err != nil {
		return errors.Wrap(err, "failed to get token")
	}

	ctx := context.Background()
	userSvc := p.clientFactory.Client(client.WithToken(token)).UserSvcAPI

	req := openapi.UserSvcSavePermitsRequest{}

	for _, role := range []string{
		usertypes.RoleAdmin,
	} {
		for _, permission := range prompttypes.AdminPermissions {
			req.Permits = append(req.Permits, openapi.UserSvcPermitInput{
				Roles:      []string{role},
				Permission: permission,
			})
		}
	}

	_, _, err = userSvc.SavePermits(ctx).
		Body(req).
		Execute()
	if err != nil {
		return err
	}

	return nil
}
