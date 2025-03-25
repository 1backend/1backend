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

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/datastore"
	usertypes "github.com/1backend/1backend/server/internal/services/user/types"
)

func (s *UserService) createUser(
	user *usertypes.User,
	password string,
	roleIds []string,
) error {
	if len(user.Contacts) > 0 {
		_, contactExists, err := s.contactsStore.Query(
			datastore.Equals(
				datastore.Field("id"),
				user.Contacts[0].Id,
			),
		).FindOne()
		if err != nil {
			return err
		}

		if contactExists {
			return errors.New("contact already exists")
		}

	}

	_, slugExists, err := s.usersStore.Query(
		datastore.Equals(datastore.Field("slug"), user.Slug),
	).FindOne()
	if err != nil {
		return err
	}

	if slugExists {
		return errors.New("slug already exists")
	}

	passwordHash, err := s.hashPassword(password)
	if err != nil {
		return err
	}

	if len(roleIds) > 0 {
		roleIdAnys := []any{}
		for _, roleId := range roleIds {
			roleIdAnys = append(roleIdAnys, roleId)
		}

		roles, err := s.rolesStore.Query(
			datastore.IsInList(datastore.Field("id"), roleIdAnys...),
		).Find()
		if err != nil {
			return err
		}
		if len(roles) == 0 {
			return errors.New("no roles found")
		}
		if len(roles) < len(roleIds) {
			return errors.New("some roles are not found")
		}
	}

	user.PasswordHash = passwordHash
	if user.Id == "" {
		user.Id = sdk.Id("usr")
	}

	now := time.Now()
	user.UpdatedAt = now
	user.CreatedAt = now

	err = s.usersStore.Create(user)
	if err != nil {
		return nil
	}

	for _, roleId := range roleIds {
		err = s.assignRole(user.Id, roleId)
		if err != nil {
			return err
		}
	}

	if len(roleIds) == 0 {
		err = s.assignRole(user.Id, usertypes.RoleUser.Id)
		if err != nil {
			return err
		}
	}

	contactIs := []datastore.Row{}
	for _, contact := range user.Contacts {
		contact.UserId = user.Id
		contactIs = append(contactIs, contact)
	}
	err = s.contactsStore.UpsertMany(contactIs)
	if err != nil {
		return err
	}

	return nil
}
