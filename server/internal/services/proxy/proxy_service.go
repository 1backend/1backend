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
	"strings"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/boot"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/lock"
	"github.com/1backend/1backend/sdk/go/middlewares"
)

type ProxyService struct {
	started    bool
	startupErr error

	clientFactory client.ClientFactory
	token         string

	authorizer auth.Authorizer

	lock      lock.DistributedLock
	publicKey string

	credentialStore  datastore.DataStore
	datastoreFactory func(tableName string, instance any) (datastore.DataStore, error)
}

func NewProxyService(
	clientFactory client.ClientFactory,
	authorizer auth.Authorizer,
	lock lock.DistributedLock,
	datastoreFactory func(tableName string, instance any) (datastore.DataStore, error),
) (*ProxyService, error) {
	cs := &ProxyService{
		lock:             lock,
		datastoreFactory: datastoreFactory,
		clientFactory:    clientFactory,
		authorizer:       authorizer,
	}

	return cs, nil
}

func (cs *ProxyService) RegisterRoutes(router *mux.Router) {
	router.PathPrefix("/").HandlerFunc(middlewares.Lazy(cs, func(w http.ResponseWriter, r *http.Request) {
		cs.Route(w, r)
	}))
}

func (cs *ProxyService) Start() error {
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

// gets service slug from http request path
// eg. /my-svc/my-endpoint -> my-svc
func getServiceSlug(r *http.Request) string {
	cleanedPath := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(cleanedPath, "/")
	if len(parts) > 0 && parts[0] != "" {
		return parts[0]
	}
	return ""
}

func (cs *ProxyService) start() error {
	if cs.datastoreFactory == nil {
		return errors.New("no datastore factory")
	}

	credentialStore, err := cs.datastoreFactory(
		"proxySvcCredentials",
		&auth.Credential{},
	)
	if err != nil {
		return err
	}
	cs.credentialStore = credentialStore

	ctx := context.Background()
	cs.lock.Acquire(ctx, "proxy-svc-start")
	defer cs.lock.Release(ctx, "proxy-svc-start")

	pk, _, err := cs.clientFactory.Client(client.WithToken(cs.token)).
		UserSvcAPI.GetPublicKey(context.Background()).
		Execute()
	if err != nil {
		return err
	}
	cs.publicKey = pk.PublicKey

	client := cs.clientFactory.Client()

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
