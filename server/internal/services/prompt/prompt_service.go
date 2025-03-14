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
	"sync"

	sdk "github.com/openorch/openorch/sdk/go"
	"github.com/openorch/openorch/sdk/go/clients/llamacpp"
	"github.com/openorch/openorch/sdk/go/datastore"
	"github.com/openorch/openorch/sdk/go/lock"

	streammanager "github.com/openorch/openorch/server/internal/services/prompt/stream"
	prompttypes "github.com/openorch/openorch/server/internal/services/prompt/types"
)

type PromptService struct {
	clientFactory sdk.ClientFactory
	token         string

	llamaCppCLient llamacpp.ClientI
	lock           lock.DistributedLock

	streamManager *streammanager.StreamManager

	promptsStore    datastore.DataStore
	credentialStore datastore.DataStore

	runMutex sync.Mutex
	trigger  chan bool
}

func NewPromptService(
	clientFactory sdk.ClientFactory,
	llamaCppClient llamacpp.ClientI,
	lock lock.DistributedLock,
	datastoreFactory func(tableName string, instance any) (datastore.DataStore, error),
) (*PromptService, error) {
	promptsStore, err := datastoreFactory(
		"promptSvcPrompts",
		&prompttypes.Prompt{},
	)
	if err != nil {
		return nil, err
	}

	credentialStore, err := datastoreFactory(
		"promptSvcCredentials",
		&sdk.Credential{},
	)
	if err != nil {
		return nil, err
	}

	service := &PromptService{
		clientFactory: clientFactory,

		llamaCppCLient: llamaCppClient,
		lock:           lock,

		streamManager: streammanager.NewStreamManager(),

		promptsStore:    promptsStore,
		credentialStore: credentialStore,

		trigger: make(chan bool, 1),
	}

	promptIs, err := service.promptsStore.Query(
		datastore.Equals(
			datastore.Field("status"),
			prompttypes.PromptStatusRunning,
		),
	).Find()
	if err != nil {
		return nil, err
	}
	promptIds := []string{}
	for _, promptI := range promptIs {
		promptIds = append(promptIds, promptI.(*prompttypes.Prompt).Id)
	}

	err = service.promptsStore.Query(
		datastore.Equals(datastore.Field("id"), promptIds),
	).UpdateFields(map[string]any{
		"status": prompttypes.PromptStatusScheduled,
	})
	if err != nil {
		return nil, err
	}

	go service.processPrompts()

	return service, nil
}

func (cs *PromptService) Start() error {
	ctx := context.Background()
	cs.lock.Acquire(ctx, "prompt-svc-start")
	defer cs.lock.Release(ctx, "prompt-svc-start")

	token, err := sdk.RegisterServiceAccount(
		cs.clientFactory.Client().UserSvcAPI,
		"prompt-svc",
		"Prompt Svc",
		cs.credentialStore,
	)
	if err != nil {
		return err
	}
	cs.token = token

	return cs.registerPermissions()
}
