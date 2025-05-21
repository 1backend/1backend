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
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/service"
	"github.com/1backend/1backend/server/internal/universe"
	"github.com/gorilla/mux"
)

type SourceService struct {
	options *universe.Options

	started    bool
	startupErr error

	token string

	credentialStore datastore.DataStore
}

func NewSourceService(
	options *universe.Options,
) (*SourceService, error) {
	credentialStore, err := options.DataStoreFactory.Create(
		"sourceSvcCredentials",
		&auth.Credential{},
	)
	if err != nil {
		return nil, err
	}

	service := &SourceService{
		options:         options,
		credentialStore: credentialStore,
	}

	return service, nil
}

func (ss *SourceService) RegisterRoutes(router *mux.Router) {
	appl := ss.options.Middlewares

	router.HandleFunc("/source-svc/repo/checkout", appl(service.Lazy(ss, func(w http.ResponseWriter, r *http.Request) {
		ss.CheckoutRepo(w, r)
	}))).
		Methods("OPTIONS", "POST")
}

func (cs *SourceService) LazyStart() error {
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
	ss.options.Lock.Acquire(ctx, "source-svc-start")
	defer ss.options.Lock.Release(ctx, "source-svc-start")

	token, err := boot.RegisterServiceAccount(
		ss.options.ClientFactory.Client().UserSvcAPI,
		"source-svc",
		"Source Svc",
		ss.credentialStore,
	)
	if err != nil {
		return err
	}

	ss.token = token.Token

	return ss.registerPermits()
}
