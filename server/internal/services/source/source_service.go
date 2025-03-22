/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package sourceservice

import (
	"context"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/lock"
)

type SourceService struct {
	clientFactory sdk.ClientFactory
	lock          lock.DistributedLock

	token string

	credentialStore datastore.DataStore
}

func NewSourceService(
	clientFactory sdk.ClientFactory,
	lock lock.DistributedLock,
	datastoreFactory sdk.DatastoreConstructor,
) (*SourceService, error) {
	credentialStore, err := datastoreFactory(
		"sourceSvcCredentials",
		&sdk.Credential{},
	)
	if err != nil {
		return nil, err
	}

	service := &SourceService{
		clientFactory:   clientFactory,
		lock:            lock,
		credentialStore: credentialStore,
	}

	return service, nil
}

func (fs *SourceService) Start() error {
	ctx := context.Background()
	fs.lock.Acquire(ctx, "source-svc-start")
	defer fs.lock.Release(ctx, "source-svc-start")

	token, err := sdk.RegisterServiceAccount(
		fs.clientFactory.Client().UserSvcAPI,
		"source-svc",
		"Source Svc",
		fs.credentialStore,
	)
	if err != nil {
		return err
	}

	fs.token = token

	return fs.registerPermissions()
}
