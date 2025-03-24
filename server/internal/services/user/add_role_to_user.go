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
	"fmt"
	"strings"

	"github.com/1backend/1backend/sdk/go/datastore"
	usertypes "github.com/1backend/1backend/server/internal/services/user/types"
)

func (s *UserService) addRoleToUser(userId string, roleId string) error {
	q := s.usersStore.Query(
		datastore.Id(userId),
	)
	userI, found, err := q.FindOne()
	if err != nil {
		return nil
	}
	if !found {
		return errors.New("user not found")
	}
	user := userI.(*usertypes.User)

	q = s.rolesStore.Query(
		datastore.Id(roleId),
	)

	// Only check static roles
	if !strings.Contains(roleId, ":{") {
		roleIs, err := q.Find()
		if err != nil {
			return nil
		}
		if len(roleIs) == 0 {
			return errors.New("role not found")
		}
	}

	roleLinks, err := s.userRoleLinksStore.Query(
		datastore.Equals(datastore.Field("userId"), userId),
	).Find()
	if err != nil {
		return err
	}

	alreadyHasRole := false
	for _, v := range roleLinks {
		if v.(*usertypes.UserRoleLink).RoleId == roleId {
			alreadyHasRole = true
		}
	}
	if alreadyHasRole {
		return nil
	}

	return s.userRoleLinksStore.Upsert(&usertypes.UserRoleLink{
		Id:     fmt.Sprintf("%v:%v", userId, roleId),
		RoleId: roleId,
		UserId: user.Id,
	})
}

func (s *UserService) removeRoleFromUser(userId string, roleId string) error {
	q := s.usersStore.Query(
		datastore.Id(userId),
	)
	userI, found, err := q.FindOne()
	if err != nil {
		return nil
	}
	if !found {
		return errors.New("user not found")
	}
	user := userI.(*usertypes.User)

	roleLinks, err := s.userRoleLinksStore.Query(
		datastore.Equals(datastore.Field("userId"), userId),
	).Find()
	if err != nil {
		return err
	}

	alreadyHasRole := false
	for _, v := range roleLinks {
		if v.(*usertypes.UserRoleLink).RoleId == roleId {
			alreadyHasRole = true
		}
	}
	if !alreadyHasRole {
		return nil
	}

	return s.userRoleLinksStore.Query(
		datastore.Id(
			fmt.Sprintf("%v:%v", user.Id, roleId))).Delete()

}
