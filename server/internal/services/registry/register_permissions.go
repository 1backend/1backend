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
	registry "github.com/1backend/1backend/server/internal/services/registry/types"
	registrytypes "github.com/1backend/1backend/server/internal/services/registry/types"
	usertypes "github.com/1backend/1backend/server/internal/services/user/types"
	"github.com/pkg/errors"
)

func app(permSlices ...[]string) []string {
	ret := []string{}
	for _, ps := range permSlices {
		ret = append(ret, ps...)
	}
	return ret
}

func (ns *RegistryService) registerPermits() error {
	ctx := context.Background()
	userSvc := ns.options.ClientFactory.Client(client.WithToken(ns.token)).UserSvcAPI

	req := openapi.UserSvcSavePermitsRequest{
		Permits: []openapi.UserSvcPermitInput{
			{
				Slugs:      []string{"deploy-svc"},
				Permission: registry.PermissionDefinitionDelete,
			},
			{
				Slugs:      []string{"deploy-svc"},
				Permission: registry.PermissionNodeDelete,
			},
			{
				Slugs:      []string{"deploy-svc"},
				Permission: registry.PermissionDefinitionView,
			},
			{
				Slugs:      []string{"deploy-svc", "proxy-svc"},
				Permission: registry.PermissionInstanceView,
			},
			{
				Slugs:      []string{"deploy-svc", "file-svc", "model-svc"},
				Permission: registry.PermissionNodeView,
			},
			{
				Slugs:      []string{"deploy-svc"},
				Permission: registry.PermissionInstanceEdit,
			},
			{
				Slugs:      []string{"deploy-svc"},
				Permission: registry.PermissionInstanceDelete,
			},
		},
	}

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
		return errors.Wrap(err, "failed to save permits")
	}

	return nil
}
