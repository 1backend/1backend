/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package backends

import (
	container "github.com/openorch/openorch/server/internal/services/container/types"
)

type ContainerBackend interface {
	// Get
	GetContainerSummary(container.GetContainerSummaryRequest) (*container.GetContainerSummaryResponse, error)

	// Build and image
	BuildImage(container.BuildImageRequest) (*container.BuildImageResponse, error)

	// List containers
	ListContainers(container.ListContainersRequest) (*container.ListContainersResponse, error)

	// Stop a container
	StopContainer(container.StopContainerRequest) (*container.StopContainerResponse, error)

	// Check if a container is running
	ContainerIsRunning(container.ContainerIsRunningRequest) (*container.ContainerIsRunningResponse, error)

	DaemonInfo(container.DaemonInfoRequest) (*container.DaemonInfoResponse, error)

	// Daemon address (eg. Docker daemon address)
	// @todo overlaps with DaemonInfo
	DaemonAddress() (string, error)

	// RunContainer starts a container or does nothing if a container with the
	// same requests parameters is already running. Idempotent.
	RunContainer(container.RunContainerRequest) (*container.RunContainerResponse, error)

	// The client returns the underlying backend specific client
	// (eg. Docker Client) to support features that require the underlying connection
	// such as acquiring logs etc.
	Client() any
}
