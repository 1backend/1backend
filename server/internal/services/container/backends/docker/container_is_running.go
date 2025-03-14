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

	dockercontainer "github.com/docker/docker/api/types/container"
	container "github.com/openorch/openorch/server/internal/services/container/types"

	"github.com/pkg/errors"
)

func (d *DockerBackend) ContainerIsRunning(req container.ContainerIsRunningRequest) (*container.ContainerIsRunningResponse, error) {
	if req.Hash != "" {
		isRunning, err := d.hashIsRunning(req.Hash)
		if err != nil {
			return nil, err
		}
		return &container.ContainerIsRunningResponse{
			IsRunning: isRunning,
		}, nil
	}

	isRunning, err := d.nameIsRunning(req.Hash)
	if err != nil {
		return nil, err
	}
	return &container.ContainerIsRunningResponse{
		IsRunning: isRunning,
	}, nil
}

func (d *DockerBackend) hashIsRunning(hash string) (bool, error) {
	ctx := context.Background()
	containers, err := d.client.ContainerList(
		ctx,
		dockercontainer.ListOptions{
			All: true,
		},
	)
	if err != nil {
		return false, errors.Wrap(
			err,
			"error listing docker containers when checking for runnign",
		)
	}

	for _, cont := range containers {
		if cont.State != "running" {
			continue
		}
		if cont.Labels["openorch-hash"] == hash {
			return true, nil
		}
	}

	return false, nil
}

func (d *DockerBackend) nameIsRunning(name string) (bool, error) {
	ctx := context.Background()
	containers, err := d.client.ContainerList(
		ctx,
		dockercontainer.ListOptions{
			All: true,
		},
	)
	if err != nil {
		return false, errors.Wrap(
			err,
			"error listing docker containers when checking for runnign",
		)
	}

	for _, cont := range containers {
		if cont.State != "running" {
			continue
		}
		for _, n := range cont.Names {
			if n == name {
				return true, nil
			}
		}
	}

	return false, nil
}
