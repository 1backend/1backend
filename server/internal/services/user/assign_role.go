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
	"fmt"

	"github.com/pkg/errors"

	"github.com/1backend/1backend/sdk/go/datastore"
	usertypes "github.com/1backend/1backend/server/internal/services/user/types"
)

func (s *UserService) assignRole(userId string, roleId string) error {
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
	if alreadyHasRole {
		return nil
	}

	err = s.userRoleLinksStore.Upsert(&usertypes.UserRoleLink{
		Id:     fmt.Sprintf("%v:%v", userId, roleId),
		RoleId: roleId,
		UserId: user.Id,
	})
	if err != nil {
		return errors.Wrap(err, "failed to add role to user")
	}

	return s.inactivateTokens(userId)
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
