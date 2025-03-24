/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This email code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package emailservice

import (
	"context"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/lock"
	email "github.com/1backend/1backend/server/internal/services/email/types"
)

type EmailService struct {
	clientFactory sdk.ClientFactory
	lock          lock.DistributedLock

	token string

	credentialStore datastore.DataStore
	emailStore      datastore.DataStore
}

func NewEmailService(
	clientFactory sdk.ClientFactory,
	lock lock.DistributedLock,
	datastoreFactory func(tableName string, instance any) (datastore.DataStore, error),
) (*EmailService, error) {
	credentialStore, err := datastoreFactory(
		"emailSvcCredentials",
		&sdk.Credential{},
	)
	if err != nil {
		return nil, err
	}
	emailStore, err := datastoreFactory(
		"emailSvcEmails",
		&email.Email{},
	)
	if err != nil {
		return nil, err
	}

	service := &EmailService{
		clientFactory:   clientFactory,
		lock:            lock,
		credentialStore: credentialStore,
		emailStore:      emailStore,
	}

	return service, nil
}

func (fs *EmailService) Start() error {
	ctx := context.Background()

	fs.lock.Acquire(ctx, "email-svc-start")
	defer fs.lock.Release(ctx, "email-svc-start")

	token, err := sdk.RegisterServiceAccount(
		fs.clientFactory.Client().UserSvcAPI,
		"email-svc",
		"Email Svc",
		fs.credentialStore,
	)
	if err != nil {
		return err
	}

	fs.token = token.Token

	return fs.registerPermissions()
}
