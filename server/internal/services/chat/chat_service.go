/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package chatservice

import (
	"context"
	"net/http"

	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/boot"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/lock"
	"github.com/1backend/1backend/sdk/go/middlewares"
	"github.com/1backend/1backend/sdk/go/service"
	"github.com/gorilla/mux"

	chattypes "github.com/1backend/1backend/server/internal/services/chat/types"
)

type ChatService struct {
	started    bool
	startupErr error

	clientFactory    client.ClientFactory
	datastoreFactory func(tableName string, instance any) (datastore.DataStore, error)
	token            string

	lock lock.DistributedLock

	messagesStore   datastore.DataStore
	threadsStore    datastore.DataStore
	credentialStore datastore.DataStore
}

func NewChatService(
	clientFactory client.ClientFactory,
	lock lock.DistributedLock,
	datastoreFactory func(tableName string, instance any) (datastore.DataStore, error),
) (*ChatService, error) {

	service := &ChatService{
		clientFactory:    clientFactory,
		datastoreFactory: datastoreFactory,
		lock:             lock,
	}

	return service, nil
}

func (cs *ChatService) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/chat-svc/thread/{threadId}/message", service.Lazy(cs, middlewares.DefaultApplicator(func(w http.ResponseWriter, r *http.Request) {
		cs.SaveMessage(w, r)
	}))).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/chat-svc/message/{messageId}", service.Lazy(cs, middlewares.DefaultApplicator(func(w http.ResponseWriter, r *http.Request) {
		cs.ReadMessage(w, r)
	}))).
		Methods("OPTIONS", "GET")

	router.HandleFunc("/chat-svc/message/{messageId}", service.Lazy(cs, middlewares.DefaultApplicator(func(w http.ResponseWriter, r *http.Request) {
		cs.DeleteMessage(w, r)
	}))).
		Methods("OPTIONS", "DELETE")

	router.HandleFunc("/chat-svc/thread/{threadId}/messages", service.Lazy(cs, middlewares.DefaultApplicator(func(w http.ResponseWriter, r *http.Request) {
		cs.ListMessages(w, r)
	}))).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/chat-svc/thread", service.Lazy(cs, middlewares.DefaultApplicator(func(w http.ResponseWriter, r *http.Request) {
		cs.SaveThread(w, r)
	}))).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/chat-svc/thread/{threadId}", service.Lazy(cs, middlewares.DefaultApplicator(func(w http.ResponseWriter, r *http.Request) {
		cs.DeleteThread(w, r)
	}))).
		Methods("OPTIONS", "DELETE")

	router.HandleFunc("/chat-svc/threads", service.Lazy(cs, middlewares.DefaultApplicator(func(w http.ResponseWriter, r *http.Request) {
		cs.ListThreads(w, r)
	}))).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/chat-svc/thread/{threadId}", service.Lazy(cs, middlewares.DefaultApplicator(func(w http.ResponseWriter, r *http.Request) {
		cs.ReadThread(w, r)
	}))).
		Methods("OPTIONS", "GET")

	router.HandleFunc("/chat-svc/evens", service.Lazy(cs, middlewares.DefaultApplicator(func(w http.ResponseWriter, r *http.Request) {
		cs.Events(w, r)
	}))).
		Methods("OPTIONS", "GET")
}

func (cs *ChatService) LazyStart() error {
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

func (cs *ChatService) start() error {
	threadsStore, err := cs.datastoreFactory("chatSvcThreads", &chattypes.Thread{})
	if err != nil {
		return err
	}
	cs.threadsStore = threadsStore

	messagesStore, err := cs.datastoreFactory(
		"chatSvcMessages",
		&chattypes.Message{},
	)
	if err != nil {
		return err
	}
	cs.messagesStore = messagesStore

	credentialStore, err := cs.datastoreFactory(
		"chatSvcCredentials",
		&auth.Credential{},
	)
	if err != nil {
		return err
	}
	cs.credentialStore = credentialStore

	ctx := context.Background()
	cs.lock.Acquire(ctx, "chat-svc-start")
	defer cs.lock.Release(ctx, "chat-svc-start")

	token, err := boot.RegisterServiceAccount(
		cs.clientFactory.Client().UserSvcAPI,
		"chat-svc",
		"Chat Svc",
		cs.credentialStore,
	)
	if err != nil {
		return err
	}
	cs.token = token.Token

	return cs.registerPermissions()
}
