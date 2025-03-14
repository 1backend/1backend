/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package dockerbackend

import (
	"context"
	"fmt"
	"sync"

	sdk "github.com/openorch/openorch/sdk/go"
	container "github.com/openorch/openorch/server/internal/services/container/types"

	"github.com/docker/docker/client"

	dockerapitypes "github.com/docker/docker/api/types"
)

type DockerBackend struct {
	mutex sync.Mutex

	client     *client.Client
	dockerHost string
	dockerPort int

	imagesCache          map[string]bool
	imagePullMutexes     map[string]*sync.Mutex
	imagePullGlobalMutex sync.Mutex
	runContainerMutex    sync.Mutex

	volumeName string

	clientFactory sdk.ClientFactory
	token         string
}

func NewDockerBackend(
	volumeName string,
	clientFactory sdk.ClientFactory,
	token string,
) (*DockerBackend, error) {
	c, err := client.NewClientWithOpts(
		client.FromEnv,
		client.WithAPIVersionNegotiation(),
	)
	if err != nil {
		return nil, err
	}

	return &DockerBackend{
		client: c,

		imagePullMutexes: make(map[string]*sync.Mutex),
		imagesCache:      make(map[string]bool),

		volumeName:    volumeName,
		clientFactory: clientFactory,
	}, nil
}

func (ds *DockerBackend) Client() any {
	return ds.client
}

func (ds *DockerBackend) getDockerPort() int {
	return ds.dockerPort
}

func (d *DockerBackend) getDockerBridgeIP() (string, error) {
	ctx := context.Background()

	networks, err := d.client.NetworkList(ctx, dockerapitypes.NetworkListOptions{})
	if err != nil {
		return "", fmt.Errorf("failed to list Docker networks: %w", err)
	}

	for _, network := range networks {
		if network.Name == "bridge" {
			for _, config := range network.IPAM.Config {
				if config.Gateway != "" {
					return config.Gateway, nil
				}
			}
		}
	}

	return "", fmt.Errorf("Docker bridge network not found")
}

func (ds *DockerBackend) DaemonAddress() (string, error) {
	// Docker host should only exist for cases like WSL when the
	// Docker host address is not localhost.
	// Here instead of trying to return localhost we will try to find the docker bridge
	// ip so containers can address each other.
	if ds.dockerHost == "" {
		return ds.getDockerBridgeIP()
	}

	return ds.dockerHost, nil
}

func mapDockerContainerToContainer(dockerContainer dockerapitypes.Container) container.Container {
	return container.Container{
		Id:     dockerContainer.ID,
		Names:  dockerContainer.Names,
		Image:  dockerContainer.Image,
		Ports:  getPorts(dockerContainer.Ports),
		Hash:   dockerContainer.ImageID,
		Labels: getLabels(dockerContainer.Labels),
		Envs:   nil, // Not directly available
		Keeps:  nil, // No direct equivalent
		Capabilities: &container.Capabilities{
			GPUEnabled: false,
		},
		Status: dockerContainer.State,
	}
}

func getPorts(ports []dockerapitypes.Port) []container.PortMapping {
	ret := []container.PortMapping{}

	for _, port := range ports {
		ret = append(ret, container.PortMapping{
			Host:     port.PublicPort,
			Internal: port.PrivatePort,
		})
	}

	return ret
}

func getLabels(labels map[string]string) []container.Label {
	ret := []container.Label{}

	for k, v := range labels {
		ret = append(ret, container.Label{
			Key:   k,
			Value: v,
		})
	}

	return ret
}
