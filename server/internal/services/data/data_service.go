/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package dynamicservice

import (
	"context"
	"net/http"
	"strings"

	"github.com/pkg/errors"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/boot"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/service"
	"github.com/gorilla/mux"

	dynamictypes "github.com/1backend/1backend/server/internal/services/data/types"
	"github.com/1backend/1backend/server/internal/universe"
)

type DataService struct {
	options    *universe.Options
	started    bool
	startupErr error

	token string

	store           datastore.DataStore
	credentialStore datastore.DataStore
	publicKey       string
}

func NewDataService(
	options *universe.Options,
) (*DataService, error) {

	service := &DataService{
		options: options,
	}

	return service, nil
}

func (ds *DataService) RegisterRoutes(router *mux.Router) {
	appl := ds.options.Middlewares

	router.HandleFunc("/data-svc/object", appl(service.Lazy(ds, func(w http.ResponseWriter, r *http.Request) {
		ds.Create(w, r)
	}))).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/data-svc/objects/update", appl(service.Lazy(ds, func(w http.ResponseWriter, r *http.Request) {
		ds.UpdateObjects(w, r)
	}))).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/data-svc/objects/delete", appl(service.Lazy(ds, func(w http.ResponseWriter, r *http.Request) {
		ds.DeleteObjects(w, r)
	}))).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/data-svc/objects", appl(service.Lazy(ds, func(w http.ResponseWriter, r *http.Request) {
		ds.Query(w, r)
	}))).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/data-svc/object/{objectId}", appl(service.Lazy(ds, func(w http.ResponseWriter, r *http.Request) {
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
	store, err := g.options.DataStoreFactory.Create("dataSvcObjects", &dynamictypes.Object{})
	if err != nil {
		return err
	}
	g.store = store

	credentialStore, err := g.options.DataStoreFactory.Create(
		"dataSvcCredentials",
		&auth.Credential{},
	)
	if err != nil {
		return err
	}
	g.credentialStore = credentialStore

	ctx := context.Background()
	g.options.Lock.Acquire(ctx, "data-svc-start")
	defer g.options.Lock.Release(ctx, "data-svc-start")

	pk, _, err := g.options.ClientFactory.Client(client.WithToken(g.token)).
		UserSvcAPI.GetPublicKey(context.Background()).
		Execute()
	if err != nil {
		return err
	}
	g.publicKey = pk.PublicKey

	token, err := boot.RegisterServiceAccount(
		g.options.ClientFactory.Client().UserSvcAPI,
		"data-svc",
		"Data Svc",
		g.credentialStore,
	)
	if err != nil {
		return err
	}
	g.token = token.Token

	return g.registerPermits()
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
