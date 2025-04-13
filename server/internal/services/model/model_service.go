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
	"net/http"
	"sync"

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/boot"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/lock"
	"github.com/1backend/1backend/sdk/go/middlewares"
	"github.com/1backend/1backend/sdk/go/service"

	modeltypes "github.com/1backend/1backend/server/internal/services/model/types"
)

const DefaultModelId = `huggingface/TheBloke/mistral-7b-instruct-v0.2.Q3_K_S.gguf`

type ModelService struct {
	started    bool
	startupErr error

	clientFactory    client.ClientFactory
	datastoreFactory func(tableName string, instance any) (datastore.DataStore, error)
	token            string

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
	clientFactory client.ClientFactory,
	lock lock.DistributedLock,
	datastoreFactory func(tableName string, insance any) (datastore.DataStore, error),
) (*ModelService, error) {
	srv := &ModelService{
		gpuPlatform:      gpuPlatform,
		clientFactory:    clientFactory,
		datastoreFactory: datastoreFactory,
		lock:             lock,
		modelPortMap:     map[int]*modeltypes.ModelState{},
	}

	return srv, nil
}

func (ms *ModelService) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/model-svc/default-model/status", service.Lazy(ms, middlewares.DefaultApplicator(func(w http.ResponseWriter, r *http.Request) {
		ms.DefaultStatus(w, r)
	}))).
		Methods("OPTIONS", "GET")

	router.HandleFunc("/model-svc/model/{modelId}/status", service.Lazy(ms, middlewares.DefaultApplicator(func(w http.ResponseWriter, r *http.Request) {
		ms.Status(w, r)
	}))).
		Methods("OPTIONS", "GET")

	router.HandleFunc("/model-svc/models", service.Lazy(ms, middlewares.DefaultApplicator(func(w http.ResponseWriter, r *http.Request) {
		ms.ListModels(w, r)
	}))).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/model-svc/platforms", service.Lazy(ms, middlewares.DefaultApplicator(func(w http.ResponseWriter, r *http.Request) {
		ms.ListPlatforms(w, r)
	}))).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/model-svc/model/{modelId}", service.Lazy(ms, middlewares.DefaultApplicator(func(w http.ResponseWriter, r *http.Request) {
		ms.Get(w, r)
	}))).
		Methods("OPTIONS", "GET")

	router.HandleFunc("/model-svc/default-model/start", service.Lazy(ms, middlewares.DefaultApplicator(func(w http.ResponseWriter, r *http.Request) {
		ms.StartDefault(w, r)
	}))).
		Methods("OPTIONS", "PUT")

	router.HandleFunc("/model-svc/model/{modelId}/start", service.Lazy(ms, middlewares.DefaultApplicator(func(w http.ResponseWriter, r *http.Request) {
		ms.StartSpecific(w, r)
	}))).
		Methods("OPTIONS", "PUT")

	router.HandleFunc("/model-svc/model/{modelId}/make-default", service.Lazy(ms, middlewares.DefaultApplicator(func(w http.ResponseWriter, r *http.Request) {
		ms.MakeDefault(w, r)
	}))).
		Methods("OPTIONS", "PUT")
}

func (cs *ModelService) LazyStart() error {
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

func (ms *ModelService) start() error {
	modelStore, err := ms.datastoreFactory("modelSvcModels", &modeltypes.Model{})
	if err != nil {
		return err
	}
	ms.modelsStore = modelStore

	platformsStore, err := ms.datastoreFactory(
		"modelSvcPlatforms",
		&modeltypes.Platform{},
	)
	if err != nil {
		return err
	}
	ms.platformsStore = platformsStore

	credentialStore, err := ms.datastoreFactory(
		"modelSvcCredentials",
		&auth.Credential{},
	)
	if err != nil {
		return err
	}
	ms.credentialStore = credentialStore

	err = ms.bootstrapModels()
	if err != nil {
		return err
	}

	ctx := context.Background()
	ms.lock.Acquire(ctx, "model-svc-start")
	defer ms.lock.Release(ctx, "model-svc-start")

	token, err := boot.RegisterServiceAccount(
		ms.clientFactory.Client().UserSvcAPI,
		"model-svc",
		"Model Svc",
		ms.credentialStore,
	)
	if err != nil {
		return err
	}
	ms.token = token.Token

	return ms.registerPermissions()
}

func (ms *ModelService) getNode() (*openapi.RegistrySvcNode, error) {
	ms.selfNodeMutex.Lock()
	defer ms.selfNodeMutex.Unlock()

	if ms.selfNode != nil {
		return ms.selfNode, nil
	}

	rsp, _, err := ms.clientFactory.Client(client.WithToken(ms.token)).
		RegistrySvcAPI.SelfNode(context.Background()).
		Execute()

	if err != nil {
		return nil, errors.Wrap(err, "error getting self node from model service")
	}

	ms.selfNode = &rsp.Node

	return ms.selfNode, nil
}
