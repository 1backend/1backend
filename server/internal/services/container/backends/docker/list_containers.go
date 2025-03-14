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

	dockerapicontainer "github.com/docker/docker/api/types/container"
	container "github.com/openorch/openorch/server/internal/services/container/types"
	"github.com/pkg/errors"
)

func (d *DockerBackend) ListContainers(
	request container.ListContainersRequest,
) (*container.ListContainersResponse, error) {
	ctx := context.Background()
	containers, err := d.client.ContainerList(
		ctx,
		dockerapicontainer.ListOptions{
			All: true,
		},
	)
	if err != nil {
		return nil, errors.Wrap(
			err,
			"error listing docker containers when getting logs",
		)
	}

	ret := []*container.Container{}
	for _, cont := range containers {
		c := mapDockerContainerToContainer(cont)
		ret = append(ret, &c)
	}

	return &container.ListContainersResponse{
		Containers: ret,
	}, nil
}
