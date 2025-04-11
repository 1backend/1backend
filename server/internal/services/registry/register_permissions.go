/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package registryservice

import (
	"context"

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	registrytypes "github.com/1backend/1backend/server/internal/services/registry/types"
	usertypes "github.com/1backend/1backend/server/internal/services/user/types"
)

func app(permSlices ...[]string) []string {
	ret := []string{}
	for _, ps := range permSlices {
		ret = append(ret, ps...)
	}
	return ret
}

func (ns *RegistryService) registerPermissions() error {
	ctx := context.Background()
	userSvc := ns.clientFactory.Client(client.WithToken(ns.token)).UserSvcAPI

	req := openapi.UserSvcSavePermitsRequest{}

	for _, role := range []string{
		usertypes.RoleAdmin,
	} {
		for _, permission := range app(
			registrytypes.NodeAdminPermissions,
			registrytypes.InstanceAdminPermissions,
			registrytypes.DefinitionAdminPermissions,
		) {
			req.Permits = append(req.Permits, openapi.UserSvcPermitInput{
				Roles:      []string{role},
				Permission: permission,
			})
		}
	}

	for _, role := range []string{
		usertypes.RoleUser,
	} {
		for _, permission := range app(
			registrytypes.InstanceUserPermissions,
			registrytypes.DefinitionUserPermissions,
		) {
			req.Permits = append(req.Permits, openapi.UserSvcPermitInput{
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
