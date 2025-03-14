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
	user "github.com/openorch/openorch/server/internal/services/user/types"
	usertypes "github.com/openorch/openorch/server/internal/services/user/types"
)

func (us *UserService) registerPermissions() error {
	ps := append(usertypes.UserPermissions, usertypes.AdminPermissions...)

	_, err := us.savePermissions(
		us.serviceUserId,
		&usertypes.SavePermissionsRequest{
			Permissions: ps,
		},
	)
	if err != nil {
		return err
	}

	links := []*user.PermissionLink{}

	for _, role := range []*usertypes.Role{
		usertypes.RoleAdmin,
	} {
		for _, permission := range usertypes.AdminPermissions {
			links = append(links, &user.PermissionLink{
				RoleId:       role.Id,
				PermissionId: permission.Id,
			})
		}
	}

	for _, role := range []*usertypes.Role{
		usertypes.RoleUser,
	} {
		for _, permission := range usertypes.UserPermissions {
			links = append(links, &user.PermissionLink{
				RoleId:       role.Id,
				PermissionId: permission.Id,
			})
		}
	}

	err = us.assignPermissions(
		us.serviceUserId,
		links,
	)
	if err != nil {
		return err
	}

	return nil
}
