/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package modelservice

import (
	"context"
	"sync"

	openapi "github.com/openorch/openorch/clients/go"
	"github.com/pkg/errors"

	sdk "github.com/openorch/openorch/sdk/go"
	"github.com/openorch/openorch/sdk/go/datastore"
	"github.com/openorch/openorch/sdk/go/lock"

	modeltypes "github.com/openorch/openorch/server/internal/services/model/types"
)

const DefaultModelId = `huggingface/TheBloke/mistral-7b-instruct-v0.2.Q3_K_S.gguf`

type ModelService struct {
	clientFactory sdk.ClientFactory
	token         string

	modelStateMutex sync.Mutex
	modelPortMap    map[int]*modeltypes.ModelState

	lock lock.DistributedLock

	modelsStore    datastore.DataStore
	platformsStore datastore.DataStore

	credentialStore datastore.DataStore

	gpuPlatform string
	llmHost     string

	selfNode      *openapi.RegistrySvcNode
	selfNodeMutex sync.Mutex
}

func NewModelService(
	// @todo GPU platform maybe this could be autodetected
	gpuPlatform string,
	llmHost string,
	clientFactory sdk.ClientFactory,
	lock lock.DistributedLock,
	datastoreFactory func(tableName string, insance any) (datastore.DataStore, error),
) (*ModelService, error) {
	srv := &ModelService{
		gpuPlatform:   gpuPlatform,
		clientFactory: clientFactory,
		lock:          lock,
		modelPortMap:  map[int]*modeltypes.ModelState{},
	}
	modelStore, err := datastoreFactory("modelSvcModels", &modeltypes.Model{})
	if err != nil {
		return nil, err
	}
	srv.modelsStore = modelStore

	platformsStore, err := datastoreFactory(
		"modelSvcPlatforms",
		&modeltypes.Platform{},
	)
	if err != nil {
		return nil, err
	}
	srv.platformsStore = platformsStore

	credentialStore, err := datastoreFactory(
		"modelSvcCredentials",
		&sdk.Credential{},
	)
	if err != nil {
		return nil, err
	}
	srv.credentialStore = credentialStore

	err = srv.bootstrapModels()
	if err != nil {
		return nil, err
	}

	return srv, nil
}

func (ms *ModelService) Start() error {
	ctx := context.Background()
	ms.lock.Acquire(ctx, "model-svc-start")
	defer ms.lock.Release(ctx, "model-svc-start")

	token, err := sdk.RegisterServiceAccount(
		ms.clientFactory.Client().UserSvcAPI,
		"model-svc",
		"Model Svc",
		ms.credentialStore,
	)
	if err != nil {
		return err
	}
	ms.token = token

	return ms.registerPermissions()
}

func (ms *ModelService) getNode() (*openapi.RegistrySvcNode, error) {
	ms.selfNodeMutex.Lock()
	defer ms.selfNodeMutex.Unlock()

	if ms.selfNode != nil {
		return ms.selfNode, nil
	}

	rsp, _, err := ms.clientFactory.Client(sdk.WithToken(ms.token)).
		RegistrySvcAPI.SelfNode(context.Background()).
		Execute()

	if err != nil {
		return nil, errors.Wrap(err, "error getting self node from model service")
	}

	ms.selfNode = &rsp.Node

	return ms.selfNode, nil
}
