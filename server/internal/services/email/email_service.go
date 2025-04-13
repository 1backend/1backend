/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This email code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package emailservice

import (
	"context"
	"net/http"

	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/boot"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/lock"
	"github.com/1backend/1backend/sdk/go/middlewares"
	email "github.com/1backend/1backend/server/internal/services/email/types"
	"github.com/gorilla/mux"
)

type EmailService struct {
	started    bool
	startupErr error

	clientFactory    client.ClientFactory
	datastoreFactory func(tableName string, instance any) (datastore.DataStore, error)
	lock             lock.DistributedLock

	token string

	credentialStore datastore.DataStore
	emailStore      datastore.DataStore
}

func NewEmailService(
	clientFactory client.ClientFactory,
	lock lock.DistributedLock,
	datastoreFactory func(tableName string, instance any) (datastore.DataStore, error),
) (*EmailService, error) {

	service := &EmailService{
		clientFactory:    clientFactory,
		datastoreFactory: datastoreFactory,
		lock:             lock,
	}

	return service, nil
}

func (es *EmailService) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/email-svc/email", middlewares.Lazy(es, middlewares.DefaultApplicator(func(w http.ResponseWriter, r *http.Request) {
		es.SendEmail(w, r)
	}))).
		Methods("OPTIONS", "PUT")
}

func (es *EmailService) Start() error {
	if es.started {
		return es.startupErr
	}

	es.startupErr = es.start()
	if es.startupErr != nil {
		return es.startupErr
	}

	es.started = true
	return nil
}

func (fs *EmailService) start() error {
	credentialStore, err := fs.datastoreFactory(
		"emailSvcCredentials",
		&auth.Credential{},
	)
	if err != nil {
		return err
	}
	fs.credentialStore = credentialStore

	emailStore, err := fs.datastoreFactory(
		"emailSvcEmails",
		&email.Email{},
	)
	if err != nil {
		return err
	}
	fs.emailStore = emailStore

	ctx := context.Background()

	fs.lock.Acquire(ctx, "email-svc-start")
	defer fs.lock.Release(ctx, "email-svc-start")

	token, err := boot.RegisterServiceAccount(
		fs.clientFactory.Client().UserSvcAPI,
		"email-svc",
		"Email Svc",
		fs.credentialStore,
	)
	if err != nil {
		return err
	}

	fs.token = token.Token

	return fs.registerPermissions()
}
