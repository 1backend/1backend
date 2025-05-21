/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package proxyservice

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/boot"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/middlewares"
	"github.com/1backend/1backend/server/internal/universe"
)

type ProxyService struct {
	options *universe.Options

	started    bool
	startupErr error

	token string

	publicKey string

	credentialStore datastore.DataStore
}

func NewProxyService(
	options *universe.Options,
) (*ProxyService, error) {
	cs := &ProxyService{
		options: options,
	}

	return cs, nil
}

func (cs *ProxyService) RegisterRoutes(router *mux.Router) {
	tokenRefresherMiddleware := middlewares.TokenRefreshMiddleware(
		cs.options.TokenRefresher,
		cs.options.TokenAutoRefreshOff,
	)

	router.PathPrefix("/").HandlerFunc(tokenRefresherMiddleware(func(w http.ResponseWriter, r *http.Request) {
		cs.Route(w, r)
	}))
}

func (cs *ProxyService) LazyStart() error {
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

func (cs *ProxyService) start() error {
	if cs.options.DataStoreFactory == nil {
		return errors.New("no datastore factory")
	}

	credentialStore, err := cs.options.DataStoreFactory.Create(
		"proxySvcCredentials",
		&auth.Credential{},
	)
	if err != nil {
		return err
	}
	cs.credentialStore = credentialStore

	ctx := context.Background()
	cs.options.Lock.Acquire(ctx, "proxy-svc-start")
	defer cs.options.Lock.Release(ctx, "proxy-svc-start")

	pk, _, err := cs.options.ClientFactory.Client(client.WithToken(cs.token)).
		UserSvcAPI.GetPublicKey(context.Background()).
		Execute()
	if err != nil {
		return err
	}
	cs.publicKey = pk.PublicKey

	client := cs.options.ClientFactory.Client()

	token, err := boot.RegisterServiceAccount(
		client.UserSvcAPI,
		"proxy-svc",
		"Proxy Svc",
		cs.credentialStore,
	)
	if err != nil {
		return err
	}
	cs.token = token.Token

	return nil
}
