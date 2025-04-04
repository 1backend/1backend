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
	user "github.com/1backend/1backend/server/internal/services/user/types"
	usertypes "github.com/1backend/1backend/server/internal/services/user/types"
)

func (us *UserService) registerPermissions() error {
	links := []*user.PermissionLink{}

	for _, role := range []string{
		usertypes.RoleAdmin,
	} {
		for _, permission := range usertypes.AdminPermissions {
			links = append(links, &user.PermissionLink{
				Role:       role,
				Permission: permission,
			})
		}
	}

	for _, role := range []string{
		usertypes.RoleUser,
	} {
		for _, permission := range usertypes.UserPermissions {
			links = append(links, &user.PermissionLink{
				Role:       role,
				Permission: permission,
			})
		}
	}

	err := us.assignPermissions(
		links,
	)
	if err != nil {
		return err
	}

	return nil
}
