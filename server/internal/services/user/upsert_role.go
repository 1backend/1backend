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
	"errors"
	"time"

	"github.com/openorch/openorch/sdk/go/datastore"
	user "github.com/openorch/openorch/server/internal/services/user/types"
	usertypes "github.com/openorch/openorch/server/internal/services/user/types"
)

func (s *UserService) upsertRole(
	userId, id, name, description string,
	permissionIds []string,
) error {
	permissions, err := s.permissionsStore.Query(
		datastore.Equals(datastore.Field("id"), permissionIds),
	).Find()
	if err != nil {
		return err
	}

	if len(permissions) < len(permissionIds) {
		return errors.New("nonexistent permissions")
	}

	roleI, found, err := s.rolesStore.Query(
		datastore.Equals(datastore.Field("id"), id),
	).FindOne()
	if err != nil {
		return err
	}
	if !found {
		role := &usertypes.Role{
			Id:          id,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Name:        name,
			Description: description,
		}
		err = s.rolesStore.Create(role)
		if err != nil {
			return err
		}
	} else {
		err = s.rolesStore.Query(
			datastore.Equals(datastore.Field("id"), id),
		).Update(roleI)
		if err != nil {
			return err
		}
	}

	links := []*user.PermissionLink{}

	for _, permissionId := range permissionIds {
		links = append(links, &user.PermissionLink{
			RoleId:       roleI.GetId(),
			PermissionId: permissionId,
		})
	}

	err = s.assignPermissions(userId, links)
	if err != nil {
		return err
	}

	return nil
}
