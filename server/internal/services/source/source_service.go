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

	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/boot"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/lock"
)

type SourceService struct {
	clientFactory client.ClientFactory
	lock          lock.DistributedLock

	token string

	credentialStore datastore.DataStore
}

func NewSourceService(
	clientFactory client.ClientFactory,
	lock lock.DistributedLock,
	datastoreFactory func(tableName string, instance any) (datastore.DataStore, error),
) (*SourceService, error) {
	credentialStore, err := datastoreFactory(
		"sourceSvcCredentials",
		&auth.Credential{},
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

	token, err := boot.RegisterServiceAccount(
		fs.clientFactory.Client().UserSvcAPI,
		"source-svc",
		"Source Svc",
		fs.credentialStore,
	)
	if err != nil {
		return err
	}

	fs.token = token.Token

	return fs.registerPermissions()
}
