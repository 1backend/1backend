/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package model_svc

import (
	"sync"

	prompt "github.com/openorch/openorch/server/internal/services/prompt/types"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

/*
Platform (~AI Platform) represents an AI container and its settings.
It defines the AI runtime, supported prompt types, and the underlying hardware architecture.
*/
type Platform struct {
	Id            string        `json:"id"`
	Name          string        `json:"name,omitempty"`
	Version       int           `json:"version,omitempty"`
	Architectures Architectures `json:"architectures"`

	// List of prompt types that the AI engine supports.
	Types []prompt.PromptType `json:"types"`
}

// GetId returns the platform's unique identifier.
func (p Platform) GetId() string {
	return p.Id
}

/*
Architectures defines container configurations for different hardware platforms.
Includes default settings and CUDA-specific configurations for GPU acceleration.
*/
type Architectures struct {
	// Default container configuration for non-GPU environments.
	Default DefaultParameters `json:"default"`

	// CUDA-specific container parameters, if applicable.
	Cuda CudaParameters `json:"cuda,omitempty"`
}

/* DefaultParameters holds the container configuration for CPU-based execution. */
type DefaultParameters struct {
	Container Container `json:"container"`
}

/*
CudaParameters defines CUDA-specific settings for GPU-accelerated containers.
It specifies the CUDA and cuDNN versions along with container image details.
*/
type CudaParameters struct {
	// Container configuration related to CUDA usage.
	Container Container `json:"container"`

	// Default CUDA version to use (e.g., "12.2" or "12.2.0").
	DefaultCudaVersion string `json:"defaultCudaVersion"`

	// Level of precision for selecting the CUDA version when resolving the container image.
	// - 2 -> Use "major.minor" (e.g., "12.2")
	// - 3 -> Use "major.minor.patch" (e.g., "12.2.0")
	CudaVersionPrecision int `json:"cudaVersionPrecision"`

	// Default cuDNN version to use alongside CUDA.
	DefaultCudnnVersion string `json:"defaultCudnnVersion"`
}

// EnvVar represents an environment variable inside the container.
//
// Inspired by `container-svc`'s EnvVar type.
type EnvVar struct {
	// Key is the environment variable name.
	Key string `json:"key"`

	// Value is the environment variable value.
	Value string `json:"value"`
}

// Keep represents a simplified volume where only the persistence of specific internal files matters,
// without concern for their exact location on the host system.
//
// Inspired by `container-svc`'s Keep type.
type Keep struct {
	// Path is the absolute path inside the container for the folder that should persist across restarts.
	Path string `json:"path"`

	// ReadOnly indicates whether the keep is read-only.
	ReadOnly bool `json:"readOnly,omitempty"`
}

/*
Container represents a deployable container configuration, including
image details, environment variables, and persistence settings.
*/
type Container struct {
	// Internal port exposed by the container.
	Port int `json:"port"`

	// Template for constructing the container image name.
	ImageTemplate string `json:"imageTemplate"`

	// Environment variables to be passed to the container (e.g., "DEVICES=all").
	Envars []EnvVar `json:"envars"`

	// List of container paths that should persist across restarts.
	Keeps []Keep `json:"keeps,omitempty"`
}

type Asset struct {
	EnvVarKey string `json:"envVarKey" binding:"required"`
	Url       string `json:"url" binding:"required"`
}

type Model struct {
	Id             string   `json:"id" binding:"required"`
	PlatformId     string   `json:"platformId" binding:"required"`
	Name           string   `json:"name" binding:"required"`
	Parameters     string   `json:"parameters,omitempty"`
	Flavour        string   `json:"flavour,omitempty"`
	Version        string   `json:"version,omitempty"`
	Quality        string   `json:"quality,omitempty"`
	Extension      string   `json:"extension,omitempty"`
	FullName       string   `json:"fullName,omitempty"`
	Tags           []string `json:"tags,omitempty"`
	Mirrors        []string `json:"mirrors,omitempty"`
	Size           float64  `json:"size,omitempty"`
	Uncensored     bool     `json:"uncensored,omitempty"`
	MaxRam         float64  `json:"maxRam,omitempty"`
	Description    string   `json:"description,omitempty"`
	PromptTemplate string   `json:"promptTemplate,omitempty"`
	QuantComment   string   `json:"quantComment,omitempty"`
	MaxBits        int      `json:"maxBits,omitempty"`
	Bits           int      `json:"bits,omitempty"`
	Assets         []Asset  `json:"assets,omitempty"`
}

func (g Model) GetId() string {
	return g.Id
}

/* Internal type for ModelService */
type ModelState struct {
	sync.Mutex
	Answering         bool
	HasCheckerRunning bool
}

// Setter methods for each field
func (m *ModelState) SetAnswering(v bool) {
	m.Lock()
	defer m.Unlock()
	m.Answering = v
}

func (m *ModelState) SetHasCheckerRunning(v bool) {
	m.Lock()
	defer m.Unlock()
	m.HasCheckerRunning = v
}

type ModelStatus struct {
	AssetsReady bool `json:"assetsReady" binding:"required"`
	/* Running triggers onModelLaunch on the frontend.
	Running is true when the model is both running and answering
	- fully loaded. */
	Running bool   `json:"running"     binding:"required"`
	Address string `json:"address"     binding:"required"`
}

type StatusRequest struct {
	ModelId string `json:"modelId"`
}

type StatusResponse struct {
	Status *ModelStatus `json:"status"`
}

type StartRequest struct {
}

type StartResponse struct {
}

type StartDefaultRequest struct {
}

type StartDefaultResponse struct {
}

type MakeDefaultRequest struct {
}

type MakeDefaultResponse struct {
}

type ListModelsRequest struct{}

type ListModelsResponse struct {
	Models []*Model `json:"models,omitempty"`
}

type GetModelRequest struct {
	Id string `json:"id,omitempty"`
}

type GetModelResponse struct {
	Exists   bool      `json:"exists"             binding:"required"`
	Model    *Model    `json:"model,omitempty"    binding:"required"`
	Platform *Platform `json:"platform,omitempty" binding:"required"`
}

type ListPlatformsRequest struct{}
type ListPlatformsResponse struct {
	Platforms []*Platform `json:"platforms,omitempty" binding:"required"`
}

//
// Events
//

const EventModelReadyName = "modelReady"

type EventModelReady struct {
	ThreadId string `json:"threadId"`
}

func (e EventModelReady) Name() string {
	return EventModelReadyName
}
