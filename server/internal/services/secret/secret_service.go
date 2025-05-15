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
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	secret "github.com/1backend/1backend/server/internal/services/secret/types"

	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/boot"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/lock"
	"github.com/1backend/1backend/sdk/go/middlewares"
	"github.com/1backend/1backend/sdk/go/service"
)

type SecretService struct {
	started    bool
	startupErr error

	clientFactory client.ClientFactory
	token         string

	authorizer auth.Authorizer

	lock      lock.DistributedLock
	publicKey string

	credentialStore  datastore.DataStore
	secretStore      datastore.DataStore
	datastoreFactory func(tableName string, instance any) (datastore.DataStore, error)

	encryptionKey string

	permissionChecker endpoint.PermissionChecker
}

func NewSecretService(
	clientFactory client.ClientFactory,
	authorizer auth.Authorizer,
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
		permissionChecker: endpoint.NewPermissionChecker(
			clientFactory,
			authorizer,
		),
	}

	credentialStore, err := cs.datastoreFactory(
		"secretSvcCredentials",
		&auth.Credential{},
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

func (ss *SecretService) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/secret-svc/secrets", middlewares.DefaultApplicator(service.Lazy(ss, func(w http.ResponseWriter, r *http.Request) {
		ss.ListSecrets(w, r)
	}))).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/secret-svc/secrets", middlewares.DefaultApplicator(service.Lazy(ss, func(w http.ResponseWriter, r *http.Request) {
		ss.SaveSecrets(w, r)
	}))).
		Methods("OPTIONS", "PUT")

	router.HandleFunc("/secret-svc/encrypt", middlewares.DefaultApplicator(service.Lazy(ss, func(w http.ResponseWriter, r *http.Request) {
		ss.Encrypt(w, r)
	}))).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/secret-svc/decrypt", middlewares.DefaultApplicator(service.Lazy(ss, func(w http.ResponseWriter, r *http.Request) {
		ss.Decrypt(w, r)
	}))).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/secret-svc/secrets", middlewares.DefaultApplicator(service.Lazy(ss, func(w http.ResponseWriter, r *http.Request) {
		ss.RemoveSecrets(w, r)
	}))).
		Methods("OPTIONS", "DELETE")

	router.HandleFunc("/secret-svc/is-secure", middlewares.DefaultApplicator(service.Lazy(ss, func(w http.ResponseWriter, r *http.Request) {
		ss.Secure(w, r)
	}))).
		Methods("OPTIONS", "GET")
}

func (cs *SecretService) LazyStart() error {
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

func (cs *SecretService) start() error {
	if cs.datastoreFactory == nil {
		return errors.New("no datastore factory")
	}

	ctx := context.Background()
	cs.lock.Acquire(ctx, "secret-svc-start")
	defer cs.lock.Release(ctx, "secret-svc-start")

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
		"secret-svc",
		"Secret Svc",
		cs.credentialStore,
	)
	if err != nil {
		return err
	}
	cs.token = token.Token

	err = cs.registerPermits()
	if err != nil {
		return err
	}

	return nil
}
