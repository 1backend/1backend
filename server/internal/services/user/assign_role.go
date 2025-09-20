/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package userservice

import (
	"github.com/pkg/errors"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/datastore"
	usertypes "github.com/1backend/1backend/server/internal/services/user/types"
)

func (s *UserService) assignRole(
	appId string,
	userId string,
	role string,
) error {
	q := s.userStore.Query(
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

	enrolls, err := s.enrollStore.Query(
		datastore.Equals(datastore.Field("appId"), appId),
		datastore.Equals(datastore.Field("userId"), userId),
	).Find()
	if err != nil {
		return errors.Wrap(err, "failed to query enrolls")
	}

	alreadyHasRole := false
	for _, v := range enrolls {
		enroll := v.(*usertypes.Enroll)
		if enroll.Role == role && enroll.AppId == appId {
			alreadyHasRole = true
		}
	}
	if alreadyHasRole {
		return nil
	}

	id := sdk.Id("enr")

	internalId, err := sdk.InternalId(appId, id)
	if err != nil {
		return errors.Wrap(err, "failed to create internal id")
	}

	inv := &usertypes.Enroll{
		InternalId: internalId,
		Id:         id,
		AppId:      appId,
		Role:       role,
		UserId:     user.Id,
	}
	err = s.enrollStore.Upsert(inv)
	if err != nil {
		return errors.Wrap(err, "failed to add role to user")
	}

	return s.inactivateTokens(appId, userId)
}

func (s *UserService) removeRoleFromUser(userId string, roleId string) error {
	q := s.userStore.Query(
		datastore.Id(userId),
	)
	_, found, err := q.FindOne()
	if err != nil {
		return nil
	}
	if !found {
		return errors.New("user not found")
	}

	enrolls, err := s.enrollStore.Query(
		datastore.Equals(datastore.Field("userId"), userId),
	).Find()
	if err != nil {
		return err
	}

	alreadyHasRole := false
	enrollId := ""
	for _, v := range enrolls {
		if v.(*usertypes.Enroll).Role == roleId {
			alreadyHasRole = true
			enrollId = v.GetId()
		}
	}
	if !alreadyHasRole {
		return nil
	}

	return s.enrollStore.Query(
		datastore.Id(enrollId),
	).Delete()
}
