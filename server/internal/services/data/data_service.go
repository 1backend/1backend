/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package dynamicservice

import (
	"context"
	"errors"
	"strings"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/boot"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/lock"

	dynamictypes "github.com/1backend/1backend/server/internal/services/data/types"
)

type DataService struct {
	clientFactory client.ClientFactory
	token         string

	lock       lock.DistributedLock
	authorizer auth.Authorizer

	store           datastore.DataStore
	credentialStore datastore.DataStore
	publicKey       string
}

func NewDataService(
	clientFactory client.ClientFactory,
	lock lock.DistributedLock,
	authorizer auth.Authorizer,
	datastoreFactory func(tableName string, instance any) (datastore.DataStore, error),
) (*DataService, error) {
	store, err := datastoreFactory("dataSvcObjects", &dynamictypes.Object{})
	if err != nil {
		return nil, err
	}
	credentialStore, err := datastoreFactory(
		"dataSvcCredentials",
		&auth.Credential{},
	)
	if err != nil {
		return nil, err
	}

	service := &DataService{
		clientFactory: clientFactory,

		credentialStore: credentialStore,
		authorizer:      authorizer,

		lock:  lock,
		store: store,
	}

	return service, nil
}

func (g *DataService) Start() error {
	ctx := context.Background()
	g.lock.Acquire(ctx, "data-svc-start")
	defer g.lock.Release(ctx, "data-svc-start")

	pk, _, err := g.clientFactory.Client(client.WithToken(g.token)).
		UserSvcAPI.GetPublicKey(context.Background()).
		Execute()
	if err != nil {
		return err
	}
	g.publicKey = pk.PublicKey

	token, err := boot.RegisterServiceAccount(
		g.clientFactory.Client().UserSvcAPI,
		"data-svc",
		"Data Svc",
		g.credentialStore,
	)
	if err != nil {
		return err
	}
	g.token = token.Token

	return g.registerPermissions()
}

func (g *DataService) createMany(
	request *dynamictypes.CreateManyRequest,
) error {
	objectIs := []datastore.Row{}
	for _, object := range request.Objects {
		if object.Id == "" {
			object.Id = sdk.Id(object.Table)
		}
		if !strings.HasPrefix(object.Id, object.Table) {
			return errors.New("wrong prefix")
		}
		objectIs = append(objectIs, object)
	}

	return g.store.CreateMany(objectIs)
}

func intersects(slice1, slice2 []string) bool {
	elemMap := make(map[string]struct{})
	for _, elem := range slice1 {
		elemMap[elem] = struct{}{}
	}

	for _, elem := range slice2 {
		if _, found := elemMap[elem]; found {
			return true
		}
	}

	return false
}
