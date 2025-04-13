/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package firehoseservice

import (
	"context"
	"log"
	"net/http"
	"sync"

	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/boot"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/lock"
	"github.com/1backend/1backend/sdk/go/middlewares"
	"github.com/gorilla/mux"

	firehosetypes "github.com/1backend/1backend/server/internal/services/firehose/types"
)

type FirehoseService struct {
	started    bool
	startupErr error

	clientFactory    client.ClientFactory
	datastoreFactory func(tableName string, instance any) (datastore.DataStore, error)
	lock             lock.DistributedLock

	token string

	subscribers map[int]func(events []*firehosetypes.Event)
	mu          sync.Mutex
	nextID      int

	credentialStore datastore.DataStore
}

func NewFirehoseService(
	clientFactory client.ClientFactory,
	lock lock.DistributedLock,
	datastoreFactory func(tableName string, instance any) (datastore.DataStore, error),
) (*FirehoseService, error) {

	service := &FirehoseService{
		clientFactory:    clientFactory,
		datastoreFactory: datastoreFactory,
		lock:             lock,
		subscribers:      make(map[int]func(events []*firehosetypes.Event)),
	}

	return service, nil
}

func (fs *FirehoseService) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/firehose-svc/events/subscribe", middlewares.DefaultApplicator(func(w http.ResponseWriter, r *http.Request) {
		if fs.Start(w, r) {
			return
		}
		fs.Subscribe(w, r)
	})).
		Methods("OPTIONS", "GET")

	router.HandleFunc("/firehose-svc/event", middlewares.DefaultApplicator(func(w http.ResponseWriter, r *http.Request) {
		if fs.Start(w, r) {
			return
		}
		fs.Publish(w, r)
	})).
		Methods("OPTIONS", "POST")
}

func (fs *FirehoseService) Start(
	w http.ResponseWriter,
	r *http.Request,
) bool {
	if fs.started {
		if fs.startupErr != nil {
			endpoint.WriteErr(w, http.StatusInternalServerError, fs.startupErr)
			return true
		}

		return false
	}

	fs.startupErr = fs.start()
	if fs.startupErr != nil {
		endpoint.WriteErr(w, http.StatusInternalServerError, fs.startupErr)
		return true
	}

	fs.started = true
	return false
}

func (fs *FirehoseService) start() error {
	credentialStore, err := fs.datastoreFactory(
		"firehoseSvcCredentials",
		&auth.Credential{},
	)
	if err != nil {
		return err
	}
	fs.credentialStore = credentialStore

	ctx := context.Background()
	fs.lock.Acquire(ctx, "firehose-svc-start")
	defer fs.lock.Release(ctx, "firehose-svc-start")

	token, err := boot.RegisterServiceAccount(
		fs.clientFactory.Client().UserSvcAPI,
		"firehose-svc",
		"Firehose Svc",
		fs.credentialStore,
	)
	if err != nil {
		return err
	}
	fs.token = token.Token

	return fs.registerPermissions()
}

func (fs *FirehoseService) publishMany(events ...*firehosetypes.Event) {
	// for _, event := range events {
	// 	logger.Debug("Event published",
	// 		slog.String("eventName", event.Name),
	// 	)
	// }

	fs.mu.Lock()
	defer fs.mu.Unlock()
	for _, subscriber := range fs.subscribers {
		go func(subscriber func(events []*firehosetypes.Event)) {
			defer func() {
				if r := recover(); r != nil {
					log.Printf("Recovered in subscriber: %v", r)
				}
			}()
			subscriber(events)
		}(subscriber)
	}
}

func (fs *FirehoseService) publish(event *firehosetypes.Event) {
	fs.publishMany(event)
}

func (fs *FirehoseService) subscribe(
	callback func(events []*firehosetypes.Event),
) int {
	fs.mu.Lock()
	defer fs.mu.Unlock()
	id := fs.nextID
	fs.subscribers[id] = callback
	fs.nextID++
	return id
}

func (fs *FirehoseService) unsubscribe(id int) {
	fs.mu.Lock()
	defer fs.mu.Unlock()
	delete(fs.subscribers, id)
}
