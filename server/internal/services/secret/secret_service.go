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
	"github.com/1backend/1backend/server/internal/universe"

	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/boot"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/service"
)

type SecretService struct {
	options *universe.Options

	started    bool
	startupErr error

	token string

	publicKey string

	credentialStore datastore.DataStore
	secretStore     datastore.DataStore
}

func NewSecretService(
	options *universe.Options,
) (*SecretService, error) {

	cs := &SecretService{
		options: options,
	}

	credentialStore, err := cs.options.DataStoreFactory.Create(
		"secretSvcCredentials",
		&auth.Credential{},
	)
	if err != nil {
		return nil, err
	}
	cs.credentialStore = credentialStore

	secretStore, err := cs.options.DataStoreFactory.Create(
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
	appl := ss.options.Middlewares

	router.HandleFunc("/secret-svc/secrets", appl(service.Lazy(ss, func(w http.ResponseWriter, r *http.Request) {
		ss.ListSecrets(w, r)
	}))).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/secret-svc/secrets", appl(service.Lazy(ss, func(w http.ResponseWriter, r *http.Request) {
		ss.SaveSecrets(w, r)
	}))).
		Methods("OPTIONS", "PUT")

	router.HandleFunc("/secret-svc/encrypt", appl(service.Lazy(ss, func(w http.ResponseWriter, r *http.Request) {
		ss.Encrypt(w, r)
	}))).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/secret-svc/decrypt", appl(service.Lazy(ss, func(w http.ResponseWriter, r *http.Request) {
		ss.Decrypt(w, r)
	}))).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/secret-svc/secrets", appl(service.Lazy(ss, func(w http.ResponseWriter, r *http.Request) {
		ss.RemoveSecrets(w, r)
	}))).
		Methods("OPTIONS", "DELETE")

	router.HandleFunc("/secret-svc/is-secure", appl(service.Lazy(ss, func(w http.ResponseWriter, r *http.Request) {
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
	if cs.options.DataStoreFactory == nil {
		return errors.New("no datastore factory")
	}

	ctx := context.Background()
	cs.options.Lock.Acquire(ctx, "secret-svc-start")
	defer cs.options.Lock.Release(ctx, "secret-svc-start")

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
