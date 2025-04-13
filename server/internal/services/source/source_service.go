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
	"net/http"

	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/boot"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/lock"
	"github.com/1backend/1backend/sdk/go/middlewares"
	"github.com/gorilla/mux"
)

type SourceService struct {
	started    bool
	startupErr error

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

func (ss *SourceService) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/source-svc/repo/checkout", middlewares.Lazy(ss, middlewares.DefaultApplicator(func(w http.ResponseWriter, r *http.Request) {
		ss.CheckoutRepo(w, r)
	}))).
		Methods("OPTIONS", "POST")
}

func (cs *SourceService) Start() error {
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

func (ss *SourceService) start() error {
	ctx := context.Background()
	ss.lock.Acquire(ctx, "source-svc-start")
	defer ss.lock.Release(ctx, "source-svc-start")

	token, err := boot.RegisterServiceAccount(
		ss.clientFactory.Client().UserSvcAPI,
		"source-svc",
		"Source Svc",
		ss.credentialStore,
	)
	if err != nil {
		return err
	}

	ss.token = token.Token

	return ss.registerPermissions()
}
