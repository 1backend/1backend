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
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/service"
	"github.com/gorilla/mux"

	streammanager "github.com/1backend/1backend/server/internal/services/prompt/stream"
	prompttypes "github.com/1backend/1backend/server/internal/services/prompt/types"
	"github.com/1backend/1backend/server/internal/universe"
)

type PromptService struct {
	options *universe.Options

	token string

	streamManager *streammanager.StreamManager

	promptsStore    datastore.DataStore
	credentialStore datastore.DataStore

	runMutex sync.Mutex
	trigger  chan bool
}

func NewPromptService(
	options *universe.Options,
) (*PromptService, error) {
	service := &PromptService{
		options:       options,
		streamManager: streammanager.NewStreamManager(),
		trigger:       make(chan bool, 1),
	}

	return service, nil
}

func (ps *PromptService) RegisterRoutes(router *mux.Router) {
	appl := ps.options.Middlewares

	router.HandleFunc("/prompt-svc/prompt", appl(service.Lazy(ps, func(w http.ResponseWriter, r *http.Request) {
		ps.Prompt(w, r)
	}))).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/prompt-svc/prompt/{promptId}", appl(service.Lazy(ps, func(w http.ResponseWriter, r *http.Request) {
		ps.RemovePrompt(w, r)
	}))).
		Methods("OPTIONS", "DELETE")

	router.HandleFunc("/prompt-svc/prompts/{threadId}/responses/subscribe", appl(service.Lazy(ps, func(w http.ResponseWriter, r *http.Request) {
		ps.SubscribeToPromptResponses(w, r)
	}))).
		Methods("OPTIONS", "GET")

	router.HandleFunc("/prompt-svc/prompts", appl(service.Lazy(ps, func(w http.ResponseWriter, r *http.Request) {
		ps.ListPrompts(w, r)
	}))).
		Methods("OPTIONS", "POST")
}

func (cs *PromptService) Start() error {
	promptsStore, err := cs.options.DataStoreFactory.Create(
		"promptSvcPrompts",
		&prompttypes.Prompt{},
	)
	if err != nil {
		return err
	}
	cs.promptsStore = promptsStore

	credentialStore, err := cs.options.DataStoreFactory.Create(
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
	cs.options.Lock.Acquire(ctx, "prompt-svc-start")
	defer cs.options.Lock.Release(ctx, "prompt-svc-start")

	token, err := boot.RegisterServiceAccount(
		cs.options.ClientFactory.Client().UserSvcAPI,
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
