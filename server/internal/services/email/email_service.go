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
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/service"
	email "github.com/1backend/1backend/server/internal/services/email/types"
	"github.com/1backend/1backend/server/internal/universe"
	"github.com/gorilla/mux"
)

type EmailService struct {
	options *universe.Options

	started    bool
	startupErr error

	token string

	credentialStore datastore.DataStore
	emailStore      datastore.DataStore
}

func NewEmailService(
	options *universe.Options,
) (*EmailService, error) {

	service := &EmailService{
		options: options,
	}

	return service, nil
}

func (es *EmailService) RegisterRoutes(router *mux.Router) {
	appl := es.options.Middlewares

	router.HandleFunc("/email-svc/email", appl(service.Lazy(es, func(w http.ResponseWriter, r *http.Request) {
		es.SendEmail(w, r)
	}))).
		Methods("OPTIONS", "PUT")
}

func (es *EmailService) LazyStart() error {
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
	credentialStore, err := fs.options.DataStoreFactory.Create(
		"emailSvcCredentials",
		&auth.Credential{},
	)
	if err != nil {
		return err
	}
	fs.credentialStore = credentialStore

	emailStore, err := fs.options.DataStoreFactory.Create(
		"emailSvcEmails",
		&email.Email{},
	)
	if err != nil {
		return err
	}
	fs.emailStore = emailStore

	ctx := context.Background()

	fs.options.Lock.Acquire(ctx, "email-svc-start")
	defer fs.options.Lock.Release(ctx, "email-svc-start")

	token, err := boot.RegisterServiceAccount(
		fs.options.ClientFactory.Client().UserSvcAPI,
		"email-svc",
		"Email Svc",
		fs.credentialStore,
	)
	if err != nil {
		return err
	}

	fs.token = token.Token

	return fs.registerPermits()
}
