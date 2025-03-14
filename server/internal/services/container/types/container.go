/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package container_svc

type ErrorResponse struct {
	Error string `json:"error"`
}

// Resources defines resource constraints for the container.
type Resources struct {
	// CPU cores allocated to the container (e.g., 0.5 = 500m, 2 = 2 cores).
	CPU float64 `json:"cpu,omitempty"`

	// Memory allocated to the container in megabytes.
	MemoryMB int `json:"memoryMB,omitempty"`

	// Disk space allocated to the container in megabytes.
	DiskMB int `json:"diskMB,omitempty"`
}

// Capabilities defines additional runtime features of the container.
type Capabilities struct {
	// GPUEnabled specifies whether GPU support is enabled for the container.
	GPUEnabled bool `json:"gpuEnabled,omitempty"`
}

// Network defines networking details for the container.
type Network struct {
	// Mode specifies the container's network mode (e.g., bridge, host, none, custom).
	Mode string `json:"mode,omitempty"`

	// IPAddress is the assigned IP address of the container.
	IPAddress string `json:"ipAddress,omitempty"`

	// MacAddress is the container's MAC address if applicable.
	MacAddress string `json:"macAddress,omitempty"`
}

// EnvVar represents an environment variable inside the container.
type EnvVar struct {
	// Key is the environment variable name.
	Key string `json:"key" binding:"required"`

	// Value is the environment variable value.
	Value string `json:"value" binding:"required"`
}

// Volume represents a persistent mount or volume inside the container.
type Volume struct {
	// Source is the host path or volume name.
	Source string `json:"source"`

	// Destination is the path inside the container.
	Destination string `json:"destination"`

	// ReadOnly indicates whether the mount is read-only.
	ReadOnly bool `json:"readOnly,omitempty"`
}

// Keep represents a simplified volume where only the persistence of specific internal files matters,
// without concern for their exact location on the host system.
type Keep struct {
	// Path is the absolute path inside the container for the folder that should persist across restarts.
	Path string `json:"path" binding:"required"`

	// ReadOnly indicates whether the keep is read-only.
	ReadOnly bool `json:"readOnly,omitempty"`
}

type PortMapping struct {
	Internal uint16 `json:"internal" binding:"required"`
	Host     uint16 `json:"host" binding:"required"`
}

type Label struct {
	Key   string `json:"key" binding:"required"`
	Value string `json:"value" binding:"required"`
}

type Asset struct {
	EnvVarKey string `json:"envVarKey" binding:"required"`
	Url       string `json:"url" binding:"required"`
}

// Container represents a running container instance.
type Container struct {
	// Id is the unique identifier for the container instance.
	Id string `json:"id"`

	// Node Id
	// Please see the documentation for the envar OPENORCH_NODE_ID
	NodeId string `json:"nodeId"`

	// Names are the human-readable aliases assigned to the container.
	Names []string `json:"names,omitempty"`

	// Image is the Docker image used to create the container.
	Image string `json:"image"`

	// Ports maps host ports (keys) to container ports (values).
	Ports []PortMapping `json:"ports,omitempty"`

	// Hash is a unique identifier associated with the container.
	Hash string `json:"hash,omitempty"`

	// Labels are metadata tags assigned to the container.
	Labels []Label `json:"labels,omitempty"`

	// Assets maps environment variable names to file URLs.
	// Example: {"MODEL": "https://huggingface.co/TheBloke/Mistral-7B-Instruct-v0.2-GGUF/resolve/main/mistral-7b-instruct-v0.2.Q2_K.gguf"}
	// These files are downloaded by the File Svc and mounted in the container.
	// The environment variable `MODEL` will point to the local file path in the container.
	Assets []Asset `json:"assets,omitempty"`

	// Envs are environment variables set within the container.
	Envs []EnvVar `json:"envs,omitempty"`

	// Keeps are paths that persist across container restarts.
	// They function like mounts or volumes, but their external storage location is irrelevant.
	Keeps []Keep `json:"keeps,omitempty"`

	// Volumes mounted by the container.
	Volumes []Volume `json:"volumes,omitempty"`

	// Status indicates the current state of the container (e.g., running, stopped).
	Status string `json:"status"`

	// Runtime specifies the container runtime (e.g., Docker, containerd, etc.).
	Runtime string `json:"runtime,omitempty"`

	// Resources defines CPU, memory, and disk constraints for the container.
	Resources *Resources `json:"resources,omitempty"`

	// Capabilities define additional runtime features, such as GPU support.
	Capabilities *Capabilities `json:"capabilities,omitempty"`

	// Network contains networking-related information for the container.
	Network *Network `json:"network,omitempty"`
}

func (l *Container) GetId() string {
	return l.Id
}

type RunContainerRequest struct {
	// Names are the human-readable aliases assigned to the container.
	Names []string `json:"names,omitempty"`

	// Image is the Docker image to use for the container
	Image string `json:"image" example:"nginx:latest" binding:"required"`

	// Ports maps host ports (keys) to container ports (values).
	Ports []PortMapping `json:"ports,omitempty"`

	// Hash is a unique identifier for the container
	Hash string `json:"hash,omitempty"`

	// Labels are metadata tags assigned to the container.
	Labels []Label `json:"labels,omitempty"`

	// Assets maps environment variable names to file URLs.
	// Example: {"MODEL": "https://huggingface.co/TheBloke/Mistral-7B-Instruct-v0.2-GGUF/resolve/main/mistral-7b-instruct-v0.2.Q2_K.gguf"}
	// These files are downloaded by the File Svc and mounted in the container.
	// The environment variable `MODEL` will point to the local file path in the container.
	Assets []Asset `json:"assets,omitempty"`

	// Envs are environment variables set within the container.
	Envs []EnvVar `json:"envs,omitempty"`

	// Keeps are paths that persist across container restarts.
	// They function like mounts or volumes, but their external storage location is irrelevant.
	Keeps []Keep `json:"keeps,omitempty"`

	// Capabilities define additional runtime features, such as GPU support.
	Capabilities *Capabilities `json:"capabilities,omitempty"`
}

type RunContainerResponse struct {
	Started bool

	// Ports is returned here as host ports might get mapped dynamically.
	Ports []PortMapping
}

type StopContainerRequest struct {
	Id   string `json:"id"   example:"4378b76e05ba"`
	Name string `json:"name" example:"sup-container-x"`
}

type StopContainerResponse struct{}

type GetContainerSummaryRequest struct {
	Hash  string `json:"hash"`
	Lines int    `json:"lines"`
}

type GetContainerSummaryResponse struct {
	Status string `json:"status"  binding:"required"`
	Logs   string `json:"logs"    binding:"required"`

	// DEPRECATED. Summary contains both Status and Logs.
	Summary string `json:"summary" binding:"required"`
}

// This is only used by the backends, not by the HTTP methods
type ContainerIsRunningRequest struct {
	Hash string
	Name string
}

type ContainerIsRunningResponse struct {
	IsRunning bool `json:"isRunning" binding:"required"`
}

type ListContainersRequest struct {
	NodeId      string `json:"nodeId,omitempty"`
	ContainerId string `json:"containerId,omitempty"`
	Limit       int64  `json:"limit,omitempty"`
}

type ListContainersResponse struct {
	Containers []*Container `json:"containers,omitempty"`
}
