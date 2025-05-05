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
	"net/http"
	"strings"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/boot"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/lock"
	"github.com/1backend/1backend/sdk/go/middlewares"
	"github.com/1backend/1backend/sdk/go/service"
	"github.com/gorilla/mux"

	dynamictypes "github.com/1backend/1backend/server/internal/services/data/types"
)

type DataService struct {
	started    bool
	startupErr error

	clientFactory    client.ClientFactory
	datastoreFactory func(tableName string, instance any) (datastore.DataStore, error)
	token            string

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

	service := &DataService{
		clientFactory:    clientFactory,
		datastoreFactory: datastoreFactory,
		authorizer:       authorizer,
		lock:             lock,
	}

	return service, nil
}

func (ds *DataService) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/data-svc/object", middlewares.DefaultApplicator(service.Lazy(ds, func(w http.ResponseWriter, r *http.Request) {
		ds.Create(w, r)
	}))).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/data-svc/objects/update", middlewares.DefaultApplicator(service.Lazy(ds, func(w http.ResponseWriter, r *http.Request) {
		ds.UpdateObjects(w, r)
	}))).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/data-svc/objects/delete", middlewares.DefaultApplicator(service.Lazy(ds, func(w http.ResponseWriter, r *http.Request) {
		ds.DeleteObjects(w, r)
	}))).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/data-svc/objects", middlewares.DefaultApplicator(service.Lazy(ds, func(w http.ResponseWriter, r *http.Request) {
		ds.Query(w, r)
	}))).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/data-svc/object/{objectId}", middlewares.DefaultApplicator(service.Lazy(ds, func(w http.ResponseWriter, r *http.Request) {
		ds.Upsert(w, r)
	}))).
		Methods("OPTIONS", "PUT")
}

func (cs *DataService) LazyStart() error {
	if cs.started {
		return cs.startupErr
	}

	cs.startupErr = cs.start()
	if cs.startupErr != nil {
		return cs.startupErr
	}

	cs.started = true
	return nil
}

func (g *DataService) start() error {
	store, err := g.datastoreFactory("dataSvcObjects", &dynamictypes.Object{})
	if err != nil {
		return err
	}
	g.store = store

	credentialStore, err := g.datastoreFactory(
		"dataSvcCredentials",
		&auth.Credential{},
	)
	if err != nil {
		return err
	}
	g.credentialStore = credentialStore

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
