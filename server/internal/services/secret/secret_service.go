/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package secretservice

import (
	"context"

	"github.com/pkg/errors"

	secret "github.com/openorch/openorch/server/internal/services/secret/types"

	sdk "github.com/openorch/openorch/sdk/go"
	"github.com/openorch/openorch/sdk/go/datastore"
	"github.com/openorch/openorch/sdk/go/lock"
)

type SecretService struct {
	clientFactory sdk.ClientFactory
	token         string

	authorizer sdk.Authorizer

	lock      lock.DistributedLock
	publicKey string

	credentialStore  datastore.DataStore
	secretStore      datastore.DataStore
	datastoreFactory func(tableName string, instance any) (datastore.DataStore, error)

	encryptionKey string
}

func NewSecretService(
	clientFactory sdk.ClientFactory,
	authorizer sdk.Authorizer,
	lock lock.DistributedLock,
	datastoreFactory func(tableName string, instance any) (datastore.DataStore, error),
	secretEncryptionKey string,
) (*SecretService, error) {

	cs := &SecretService{
		lock:             lock,
		datastoreFactory: datastoreFactory,
		clientFactory:    clientFactory,
		authorizer:       authorizer,
		encryptionKey:    secretEncryptionKey,
	}

	credentialStore, err := cs.datastoreFactory(
		"secretSvcCredentials",
		&sdk.Credential{},
	)
	if err != nil {
		return nil, err
	}
	cs.credentialStore = credentialStore

	secretStore, err := cs.datastoreFactory(
		"secretSvcSecrets",
		&secret.Secret{},
	)
	if err != nil {
		return nil, err
	}
	cs.secretStore = secretStore

	return cs, nil
}

func (cs *SecretService) Start() error {
	if cs.datastoreFactory == nil {
		return errors.New("no datastore factory")
	}

	ctx := context.Background()
	cs.lock.Acquire(ctx, "secret-svc-start")
	defer cs.lock.Release(ctx, "secret-svc-start")

	pk, _, err := cs.clientFactory.Client(sdk.WithToken(cs.token)).
		UserSvcAPI.GetPublicKey(context.Background()).
		Execute()
	if err != nil {
		return err
	}
	cs.publicKey = *pk.PublicKey

	client := cs.clientFactory.Client()

	token, err := sdk.RegisterServiceAccount(
		client.UserSvcAPI,
		"secret-svc",
		"Secret Svc",
		cs.credentialStore,
	)
	if err != nil {
		return err
	}
	cs.token = token

	err = cs.registerPermissions()
	if err != nil {
		return err
	}

	return nil
}
