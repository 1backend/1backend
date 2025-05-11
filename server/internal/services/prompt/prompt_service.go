/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package promptservice

import (
	"context"
	"net/http"
	"sync"

	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/boot"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/lock"
	"github.com/1backend/1backend/sdk/go/middlewares"
	"github.com/1backend/1backend/sdk/go/service"
	"github.com/gorilla/mux"

	"github.com/1backend/1backend/server/internal/clients/llamacpp"
	streammanager "github.com/1backend/1backend/server/internal/services/prompt/stream"
	prompttypes "github.com/1backend/1backend/server/internal/services/prompt/types"
)

type PromptService struct {
	clientFactory    client.ClientFactory
	datastoreFactory func(tableName string, instance any) (datastore.DataStore, error)
	token            string

	llamaCppCLient llamacpp.ClientI
	lock           lock.DistributedLock

	streamManager *streammanager.StreamManager

	promptsStore    datastore.DataStore
	credentialStore datastore.DataStore

	runMutex sync.Mutex
	trigger  chan bool

	permissionChecker endpoint.PermissionChecker
}

func NewPromptService(
	clientFactory client.ClientFactory,
	llamaCppClient llamacpp.ClientI,
	lock lock.DistributedLock,
	datastoreFactory func(tableName string, instance any) (datastore.DataStore, error),
) (*PromptService, error) {

	service := &PromptService{
		clientFactory:    clientFactory,
		datastoreFactory: datastoreFactory,

		llamaCppCLient: llamaCppClient,
		lock:           lock,

		streamManager: streammanager.NewStreamManager(),

		trigger: make(chan bool, 1),
		permissionChecker: endpoint.NewPermissionChecker(
			clientFactory,
		),
	}

	return service, nil
}

func (ps *PromptService) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/prompt-svc/prompt", middlewares.DefaultApplicator(service.Lazy(ps, func(w http.ResponseWriter, r *http.Request) {
		ps.Prompt(w, r)
	}))).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/prompt-svc/prompt/{promptId}", middlewares.DefaultApplicator(service.Lazy(ps, func(w http.ResponseWriter, r *http.Request) {
		ps.RemovePrompt(w, r)
	}))).
		Methods("OPTIONS", "DELETE")

	router.HandleFunc("/prompt-svc/prompts/{threadId}/responses/subscribe", middlewares.DefaultApplicator(service.Lazy(ps, func(w http.ResponseWriter, r *http.Request) {
		ps.SubscribeToPromptResponses(w, r)
	}))).
		Methods("OPTIONS", "GET")

	router.HandleFunc("/prompt-svc/prompts", middlewares.DefaultApplicator(service.Lazy(ps, func(w http.ResponseWriter, r *http.Request) {
		ps.ListPrompts(w, r)
	}))).
		Methods("OPTIONS", "POST")
}

func (cs *PromptService) Start() error {
	promptsStore, err := cs.datastoreFactory(
		"promptSvcPrompts",
		&prompttypes.Prompt{},
	)
	if err != nil {
		return err
	}
	cs.promptsStore = promptsStore

	credentialStore, err := cs.datastoreFactory(
		"promptSvcCredentials",
		&auth.Credential{},
	)
	if err != nil {
		return err
	}
	cs.credentialStore = credentialStore

	promptIs, err := cs.promptsStore.Query(
		datastore.Equals(
			datastore.Field("status"),
			prompttypes.PromptStatusRunning,
		),
	).Find()
	if err != nil {
		return err
	}
	promptIds := []string{}
	for _, promptI := range promptIs {
		promptIds = append(promptIds, promptI.(*prompttypes.Prompt).Id)
	}

	err = cs.promptsStore.Query(
		datastore.Equals(datastore.Field("id"), promptIds),
	).UpdateFields(map[string]any{
		"status": prompttypes.PromptStatusScheduled,
	})
	if err != nil {
		return err
	}

	go cs.processPrompts()

	return nil
}

func (cs *PromptService) LazyStart() error {
	_, err := cs.getToken()
	if err != nil {
		return err
	}

	return nil
}

func (cs *PromptService) getToken() (string, error) {
	if cs.token != "" {
		return cs.token, nil
	}

	ctx := context.Background()
	cs.lock.Acquire(ctx, "prompt-svc-start")
	defer cs.lock.Release(ctx, "prompt-svc-start")

	token, err := boot.RegisterServiceAccount(
		cs.clientFactory.Client().UserSvcAPI,
		"prompt-svc",
		"Prompt Svc",
		cs.credentialStore,
	)
	if err != nil {
		return "", err
	}
	cs.token = token.Token

	err = cs.registerPermits()
	if err != nil {
		return "", err
	}

	return cs.token, nil
}
