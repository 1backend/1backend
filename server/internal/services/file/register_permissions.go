/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package fileservice

import (
	"context"

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	file "github.com/1backend/1backend/server/internal/services/file/types"
	usertypes "github.com/1backend/1backend/server/internal/services/user/types"
	"github.com/pkg/errors"
)

func (fs *FileService) registerPermits() error {
	ctx := context.Background()
	userSvc := fs.options.ClientFactory.Client(client.WithToken(fs.token)).UserSvcAPI

	req := openapi.UserSvcSavePermitsRequest{
		Permits: []openapi.UserSvcPermitInput{
			{
				App:        openapi.PtrString("*"),
				Slugs:      []string{"docker-svc", "model-svc"},
				Permission: file.PermissionDownloadView,
			},
			{
				App:        openapi.PtrString("*"),
				Slugs:      []string{"model-svc"},
				Permission: file.PermissionDownloadCreate,
			},
			{
				App:        openapi.PtrString("*"),
				Slugs:      []string{"prompt-svc"},
				Permission: file.PermissionUploadCreate,
			},
		},
	}

	for _, role := range []string{
		usertypes.RoleAdmin,
	} {
		for _, permission := range file.AdminPermissions {
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
		return errors.Wrap(err, "cannot save permits")
	}

	return nil
}
