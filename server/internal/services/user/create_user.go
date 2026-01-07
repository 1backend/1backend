/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package userservice

import (
	"time"

	"github.com/pkg/errors"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/datastore"
	usertypes "github.com/1backend/1backend/server/internal/services/user/types"
)

// Creates the user along with contacts, password and roles.
// Accepts the contact as is, verification should be done beforehand if needed.
func (s *UserService) createUserWithoutVerification(
	appId string,
	userInput *usertypes.UserInput,
	contacts []usertypes.Contact,
	password string,
	roles []string,
) error {
	if len(contacts) > 0 {
		_, contactExists, err := s.contactStore.Query(
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

	_, slugExists, err := s.userStore.Query(
		datastore.Equals(datastore.Field("slug"), userInput.Slug),
	).FindOne()
	if err != nil {
		return err
	}

	if slugExists {
		return errors.New("slug already exists")
	}

	now := time.Now()

	if password != "" {
		passwordHash, err := s.hashPassword(password)
		if err != nil {
			return err
		}

		if userInput.Id == "" {
			userInput.Id = sdk.Id("usr")
		}

		err = s.passwordStore.Upsert(&usertypes.Password{
			Id:           sdk.Id("pw"),
			PasswordHash: passwordHash,
			UserId:       userInput.Id,
			CreatedAt:    now,
		})
		if err != nil {
			return errors.Wrap(err, "failed to save password")
		}
	}

	user := &usertypes.User{
		Id:              userInput.Id,
		CreatedAt:       now,
		UpdatedAt:       now,
		Name:            userInput.Name,
		Slug:            userInput.Slug,
		Labels:          userInput.Labels,
		ThumbnailFileId: userInput.ThumbnailFileId,
	}

	err = s.userStore.Create(user)
	if err != nil {
		return nil
	}

	for _, role := range roles {
		err = s.assignRole(appId, user.Id, role)
		if err != nil {
			return err
		}
	}

	contactIs := []datastore.Row{}
	for _, contact := range contacts {
		contact.UserId = user.Id
		contactIs = append(contactIs, contact)
	}
	err = s.contactStore.UpsertMany(contactIs)
	if err != nil {
		return err
	}

	return nil
}
