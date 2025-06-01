/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package chatservice

import (
	"context"
	"net/http"

	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/boot"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/service"
	"github.com/gorilla/mux"

	chattypes "github.com/1backend/1backend/server/internal/services/chat/types"
	"github.com/1backend/1backend/server/internal/universe"
)

type ChatService struct {
	options *universe.Options

	started    bool
	startupErr error

	token string

	messagesStore   datastore.DataStore
	threadsStore    datastore.DataStore
	credentialStore datastore.DataStore
}

func NewChatService(
	options *universe.Options,
) (*ChatService, error) {

	service := &ChatService{
		options: options,
	}

	return service, nil
}

func (cs *ChatService) RegisterRoutes(router *mux.Router) {
	appl := cs.options.Middlewares

	router.HandleFunc("/chat-svc/thread/{threadId}/message", appl(service.Lazy(cs, func(w http.ResponseWriter, r *http.Request) {
		cs.SaveMessage(w, r)
	}))).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/chat-svc/message/{messageId}", appl(service.Lazy(cs, func(w http.ResponseWriter, r *http.Request) {
		cs.DeleteMessage(w, r)
	}))).
		Methods("OPTIONS", "DELETE")

	router.HandleFunc("/chat-svc/messages", appl(service.Lazy(cs, func(w http.ResponseWriter, r *http.Request) {
		cs.ListMessages(w, r)
	}))).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/chat-svc/thread", appl(service.Lazy(cs, func(w http.ResponseWriter, r *http.Request) {
		cs.SaveThread(w, r)
	}))).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/chat-svc/thread/{threadId}", appl(service.Lazy(cs, func(w http.ResponseWriter, r *http.Request) {
		cs.DeleteThread(w, r)
	}))).
		Methods("OPTIONS", "DELETE")

	router.HandleFunc("/chat-svc/threads", appl(service.Lazy(cs, func(w http.ResponseWriter, r *http.Request) {
		cs.ListThreads(w, r)
	}))).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/chat-svc/evens", appl(service.Lazy(cs, func(w http.ResponseWriter, r *http.Request) {
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
	threadsStore, err := cs.options.DataStoreFactory.Create("chatSvcThreads", &chattypes.Thread{})
	if err != nil {
		return err
	}
	cs.threadsStore = threadsStore

	messagesStore, err := cs.options.DataStoreFactory.Create(
		"chatSvcMessages",
		&chattypes.Message{},
	)
	if err != nil {
		return err
	}
	cs.messagesStore = messagesStore

	credentialStore, err := cs.options.DataStoreFactory.Create(
		"chatSvcCredentials",
		&auth.Credential{},
	)
	if err != nil {
		return err
	}
	cs.credentialStore = credentialStore

	ctx := context.Background()
	cs.options.Lock.Acquire(ctx, "chat-svc-start")
	defer cs.options.Lock.Release(ctx, "chat-svc-start")

	token, err := boot.RegisterServiceAccount(
		cs.options.ClientFactory.Client().UserSvcAPI,
		"chat-svc",
		"Chat Svc",
		cs.credentialStore,
	)
	if err != nil {
		return err
	}
	cs.token = token.Token

	return cs.registerPermits()
}
