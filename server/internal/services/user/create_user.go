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
	"time"

	"github.com/pkg/errors"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/datastore"
	usertypes "github.com/1backend/1backend/server/internal/services/user/types"
)

func (s *UserService) createUser(
	user *usertypes.User,
	contacts []usertypes.Contact,
	password string,
	roles []string,
) error {
	if len(contacts) > 0 {
		_, contactExists, err := s.contactsStore.Query(
			datastore.Equals(
				datastore.Field("id"),
				contacts[0].Id,
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

	if user.Id == "" {
		user.Id = sdk.Id("usr")

	}

	now := time.Now()

	err = s.passwordsStore.Upsert(&usertypes.Password{
		Id:           sdk.Id("pw"),
		PasswordHash: passwordHash,
		UserId:       user.Id,
		CreatedAt:    now,
	})
	if err != nil {
		return errors.Wrap(err, "failed to save password")
	}

	user.UpdatedAt = now
	user.CreatedAt = now

	err = s.usersStore.Create(user)
	if err != nil {
		return nil
	}

	for _, role := range roles {
		err = s.assignRole(user.Id, role)
		if err != nil {
			return err
		}
	}

	contactIs := []datastore.Row{}
	for _, contact := range contacts {
		contact.UserId = user.Id
		contactIs = append(contactIs, contact)
	}
	err = s.contactsStore.UpsertMany(contactIs)
	if err != nil {
		return err
	}

	return nil
}
