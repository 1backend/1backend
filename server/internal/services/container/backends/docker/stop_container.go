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
)

func (dm *DockerBackend) StopContainer(
	req container.StopContainerRequest,
) (*container.StopContainerResponse, error) {
	stopID := req.Id
	if stopID == "" {
		stopID = req.Name
	}

	return &container.StopContainerResponse{}, dm.client.ContainerStop(context.Background(), stopID, dockercontainer.StopOptions{})
}
