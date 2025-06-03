/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package containerservice

import (
	"context"
	"log/slog"
	"net/http"
	"sync"
	"time"
	"unicode/utf8"

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/boot"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/logger"
	"github.com/1backend/1backend/sdk/go/service"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	"github.com/1backend/1backend/server/internal/services/container/backends"
	container "github.com/1backend/1backend/server/internal/services/container/types"
	"github.com/1backend/1backend/server/internal/universe"
	dockerclient "github.com/docker/docker/client"

	dockerbackend "github.com/1backend/1backend/server/internal/services/container/backends/docker"
	"github.com/1backend/1backend/server/internal/services/container/logaccumulator"
)

type ContainerService struct {
	options *universe.Options

	token string

	backend backends.ContainerBackend

	credentialStore datastore.DataStore
	containerStore  datastore.DataStore
	logStore        datastore.DataStore

	selfNode      *openapi.RegistrySvcNode
	selfNodeMutex sync.Mutex

	containerLoopTrigger chan bool
}

func NewContainerService(
	options *universe.Options,
) (*ContainerService, error) {

	service := &ContainerService{
		options: options,
	}

	return service, nil
}

func (cs *ContainerService) RegisterRoutes(router *mux.Router) {
	appl := cs.options.Middlewares

	router.HandleFunc("/container-svc/daemon/info", appl(service.Lazy(cs, func(w http.ResponseWriter, r *http.Request) {
		cs.DaemonInfo(w, r)
	}))).
		Methods("OPTIONS", "GET")

	router.HandleFunc("/container-svc/image/{imageName}/pullable", appl(service.Lazy(cs, func(w http.ResponseWriter, r *http.Request) {
		cs.ImagePullable(w, r)
	}))).
		Methods("OPTIONS", "GET")

	router.HandleFunc("/container-svc/host", appl(service.Lazy(cs, func(w http.ResponseWriter, r *http.Request) {
		cs.Host(w, r)
	}))).
		Methods("OPTIONS", "GET")

	router.HandleFunc("/container-svc/logs", appl(service.Lazy(cs, func(w http.ResponseWriter, r *http.Request) {
		cs.ListLogs(w, r)
	}))).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/container-svc/containers", appl(service.Lazy(cs, func(w http.ResponseWriter, r *http.Request) {
		cs.ListContainers(w, r)
	}))).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/container-svc/container", appl(service.Lazy(cs, func(w http.ResponseWriter, r *http.Request) {
		cs.RunContainer(w, r)
	}))).
		Methods("OPTIONS", "PUT")

	router.HandleFunc("/container-svc/image", appl(service.Lazy(cs, func(w http.ResponseWriter, r *http.Request) {
		cs.BuildImage(w, r)
	}))).
		Methods("OPTIONS", "PUT")

	router.HandleFunc("/container-svc/container/stop", appl(service.Lazy(cs, func(w http.ResponseWriter, r *http.Request) {
		cs.StopContainer(w, r)
	}))).
		Methods("OPTIONS", "PUT")

	router.HandleFunc("/container-svc/container/is-running", appl(service.Lazy(cs, func(w http.ResponseWriter, r *http.Request) {
		cs.ContainerIsRunning(w, r)
	}))).
		Methods("OPTIONS", "GET")

	router.HandleFunc("/container-svc/container/summary", appl(service.Lazy(cs, func(w http.ResponseWriter, r *http.Request) {
		cs.Summary(w, r)
	}))).
		Methods("OPTIONS", "GET")
}

func (cs *ContainerService) Start() error {
	credentialStore, err := cs.options.DataStoreFactory.Create(
		"containerSvcCredentials",
		&auth.Credential{},
	)
	if err != nil {
		return err
	}
	cs.credentialStore = credentialStore

	containerStore, err := cs.options.DataStoreFactory.Create(
		"containerSvcContainers",
		&container.Container{},
	)
	if err != nil {
		return err
	}
	cs.containerStore = containerStore

	logStore, err := cs.options.DataStoreFactory.Create(
		"containerSvcLogs",
		&container.Log{},
	)
	if err != nil {
		return err
	}
	cs.logStore = logStore

	backend, err := dockerbackend.NewDockerBackend(
		cs.options.VolumeName,
		cs.options.ClientFactory,
		cs.token,
	)
	if err != nil {
		return err
	}
	cs.backend = backend

	go cs.containerLoop()
	go cs.logLoop()
	go cs.containerLoop()

	return nil
}

func (cs *ContainerService) LazyStart() error {
	_, err := cs.getToken()
	if err != nil {
		return errors.Wrap(err, "failed to get token")
	}

	return nil
}

func (cs *ContainerService) getToken() (string, error) {
	if cs.token != "" {
		return cs.token, nil
	}

	ctx := context.Background()
	cs.options.Lock.Acquire(ctx, "container-svc-start")
	defer cs.options.Lock.Release(ctx, "container-svc-start")

	token, err := boot.RegisterServiceAccount(
		cs.options.ClientFactory.Client().UserSvcAPI,
		"container-svc",
		"Container Svc",
		cs.credentialStore,
	)
	if err != nil {
		return "", err
	}
	cs.token = token.Token

	err = cs.registerPermits()
	if err != nil {
		return "", errors.Wrap(err, "failed to register permissions")
	}

	return cs.token, nil
}

func (ms *ContainerService) logLoop() {
	la := logaccumulator.NewLogAccumulator(0, 0, func(ls []*logaccumulator.LogChunk) {
		// logs := make([]datastore.Row, len(ls))

		for _, l := range ls {
			log := &container.Log{
				Id:          l.ChunkID,
				ContainerId: l.ProducerID,
				// @todo save node id

				// Without trimming we get this:
				// invalid byte sequence for encoding \"UTF8\": 0x00
				Content: string(cleanInvalidUTF8(l.Buffer.Bytes())),
			}

			// logs = append(logs, log)

			// @todo remove single upsert once upsertmany is fixed
			err := ms.logStore.Upsert(log)
			if err != nil {
				logger.Error("Error saving container log",
					slog.String("error", err.Error()),
				)
			}
		}

		// @todo Fix upsertmany and use that as it's more
		// performant.
		//
		// err := ms.logStore.UpsertMany(logs)
		// if err != nil {
		// 	logger.Error("Error saving container logs",
		// 		slog.String("error", err.Error()),
		// 	)
		// }
	})

	go dockerbackend.StartDockerLogListener(ms.backend.Client().(*dockerclient.Client), la)
}

func (ms *ContainerService) containerLoop() {
	ctracker := dockerbackend.NewContainerTracker()

	go dockerbackend.StartDockerContainerTracker(
		ms.backend.Client().(*dockerclient.Client),
		ctracker,
	)

	ticker := time.NewTicker(1000 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			for _, c := range ctracker.GetContainers() {
				err := ms.containerStore.Upsert(&c)
				if err != nil {
					logger.Error("Error saving container",
						slog.String("error", err.Error()),
					)
				}
			}
		}
	}
}

// Remove invalid UTF-8 sequences and unwanted control characters
func cleanInvalidUTF8(data []byte) []byte {
	result := make([]byte, 0, len(data)) // Preallocate for efficiency

	for len(data) > 0 {
		r, size := utf8.DecodeRune(data)
		if r == utf8.RuneError && size == 1 {
			// Skip invalid UTF-8 bytes
			data = data[size:]
			continue
		}

		// Skip null byte (0x00) and other control characters (0x01 - 0x1F, excluding \t, \n, \r)
		if r == '\x00' || (r < 32 && r != '\t' && r != '\n' && r != '\r') {
			data = data[size:]
			continue
		}

		// Append valid characters
		result = append(result, data[:size]...)
		data = data[size:]
	}

	return result
}

func (ms *ContainerService) containerLoopCycle() error {
	//node, err := ms.selfNode()
	//if err != nil {
	//	return err
	//}
	//
	//containers, err := ms.client.ContainerList(context.Background(), container.ListOptions{})
	//if err != nil {
	//	return nil
	//}

	return nil
}

func (ms *ContainerService) getNode() (*openapi.RegistrySvcNode, error) {
	ms.selfNodeMutex.Lock()
	defer ms.selfNodeMutex.Unlock()

	if ms.selfNode != nil {
		return ms.selfNode, nil
	}

	rsp, _, err := ms.options.ClientFactory.Client(client.WithToken(ms.token)).
		RegistrySvcAPI.SelfNode(context.Background()).
		Execute()

	if err != nil {
		return nil, errors.Wrap(err, "error getting self node from model service")
	}

	ms.selfNode = &rsp.Node

	return ms.selfNode, nil
}
