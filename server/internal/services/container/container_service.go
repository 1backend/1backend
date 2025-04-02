/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package containerservice

import (
	"context"
	"log/slog"
	"sync"
	"time"
	"unicode/utf8"

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/boot"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/lock"
	"github.com/1backend/1backend/sdk/go/logger"
	"github.com/pkg/errors"

	container "github.com/1backend/1backend/server/internal/services/container/types"
	dockerclient "github.com/docker/docker/client"

	"github.com/1backend/1backend/server/internal/services/container/backends"
	dockerbackend "github.com/1backend/1backend/server/internal/services/container/backends/docker"
	"github.com/1backend/1backend/server/internal/services/container/logaccumulator"
)

type ContainerService struct {
	clientFactory client.ClientFactory
	token         string

	lock lock.DistributedLock

	backend backends.ContainerBackend

	credentialStore datastore.DataStore
	containerStore  datastore.DataStore
	logStore        datastore.DataStore

	selfNode      *openapi.RegistrySvcNode
	selfNodeMutex sync.Mutex

	volumeName           string
	containerLoopTrigger chan bool
}

func NewContainerService(
	volumeName string,
	clientFactory client.ClientFactory,
	lock lock.DistributedLock,
	datastoreFactory func(
		tableName string,
		instance any,
	) (datastore.DataStore, error),
) (*ContainerService, error) {
	credentialStore, err := datastoreFactory(
		"containerSvcCredentials",
		&auth.Credential{},
	)
	if err != nil {
		return nil, err
	}

	containerStore, err := datastoreFactory(
		"containerSvcContainers",
		&container.Container{},
	)
	if err != nil {
		return nil, err
	}

	logStore, err := datastoreFactory(
		"containerSvcLogs",
		&container.Log{},
	)
	if err != nil {
		return nil, err
	}

	service := &ContainerService{
		clientFactory: clientFactory,
		lock:          lock,

		credentialStore: credentialStore,
		containerStore:  containerStore,
		logStore:        logStore,

		volumeName: volumeName,
	}

	return service, nil
}

func (ds *ContainerService) Start() error {
	ctx := context.Background()
	ds.lock.Acquire(ctx, "container-svc-start")
	defer ds.lock.Release(ctx, "container-svc-start")

	token, err := boot.RegisterServiceAccount(
		ds.clientFactory.Client().UserSvcAPI,
		"container-svc",
		"Container Svc",
		ds.credentialStore,
	)
	if err != nil {
		return err
	}
	ds.token = token.Token

	backend, err := dockerbackend.NewDockerBackend(
		ds.volumeName,
		ds.clientFactory,
		ds.token,
	)
	if err != nil {
		return err
	}
	ds.backend = backend

	go ds.containerLoop()
	go ds.logLoop()
	go ds.containerLoop()

	return ds.registerPermissions()
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

	rsp, _, err := ms.clientFactory.Client(client.WithToken(ms.token)).
		RegistrySvcAPI.SelfNode(context.Background()).
		Execute()

	if err != nil {
		return nil, errors.Wrap(err, "error getting self node from model service")
	}

	ms.selfNode = &rsp.Node

	return ms.selfNode, nil
}
