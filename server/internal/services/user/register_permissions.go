/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package userservice

import (
	"context"

	user "github.com/1backend/1backend/server/internal/services/user/types"
	usertypes "github.com/1backend/1backend/server/internal/services/user/types"
)

func (us *UserService) registerPermissions() error {
	grants := []*user.Grant{}

	for _, role := range []string{
		usertypes.RoleAdmin,
	} {
		for _, permission := range usertypes.AdminPermissions {
			grants = append(grants, &user.Grant{
				Roles:      []string{role},
				Permission: permission,
			})
		}
	}

	for _, role := range []string{
		usertypes.RoleUser,
	} {
		for _, permission := range usertypes.UserPermissions {
			grants = append(grants, &user.Grant{
				Roles:      []string{role},
				Permission: permission,
			})
		}
	}

	err := us.saveGrants(
		context.Background(),
		&usertypes.SaveGrantsRequest{
			Grants: grants,
		},
	)
	if err != nil {
		return err
	}

	return nil
}
