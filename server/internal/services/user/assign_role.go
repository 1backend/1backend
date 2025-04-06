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
	"github.com/pkg/errors"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/datastore"
	usertypes "github.com/1backend/1backend/server/internal/services/user/types"
)

func (s *UserService) assignRole(userId string, role string) error {
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

	invites, err := s.invitesStore.Query(
		datastore.Equals(datastore.Field("userId"), userId),
	).Find()
	if err != nil {
		return err
	}

	alreadyHasRole := false
	for _, v := range invites {
		if v.(*usertypes.Invite).Role == role {
			alreadyHasRole = true
		}
	}
	if alreadyHasRole {
		return nil
	}

	err = s.invitesStore.Upsert(&usertypes.Invite{
		Id:     sdk.Id("inv"),
		Role:   role,
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
	_, found, err := q.FindOne()
	if err != nil {
		return nil
	}
	if !found {
		return errors.New("user not found")
	}

	invites, err := s.invitesStore.Query(
		datastore.Equals(datastore.Field("userId"), userId),
	).Find()
	if err != nil {
		return err
	}

	alreadyHasRole := false
	inviteId := ""
	for _, v := range invites {
		if v.(*usertypes.Invite).Role == roleId {
			alreadyHasRole = true
			inviteId = v.GetId()
		}
	}
	if !alreadyHasRole {
		return nil
	}

	return s.invitesStore.Query(
		datastore.Id(inviteId),
	).Delete()
}
