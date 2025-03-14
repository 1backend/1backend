/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package registry_svc

// Definition can be roughly thought of as a container image + some additional details
// like clients, api definition etc.
type Definition struct {
	Id string `json:"id,omitempty" binding:"required"`

	// Container specifications for Docker, K8s, etc.
	// Use this to deploy already built images.
	Image *ImageSpec `json:"image,omitempty"`

	// Repository based definitions is an alternative to Image definitions.
	// Instead of deploying an already built and pushed image, a source code repository
	// url can be provided. The container will be built from the source.
	Repository *RepositorySpec `json:"repository,omitempty"`

	// API Specs such as OpenAPI definitions etc.
	APISpecs []APISpec `json:"apiSpecs,omitempty"`

	// Programming language clients such as on npm or GitHub.
	Clients []Client `json:"clients,omitempty"`

	// Envars is a map of Renvironment variables that a deployment (see Deploy Svc Deployment) of this definition will REQUIRE to run.
	// E.g., {"DB_URL": "mysql://user:password@host:port/db"}
	// These will be injected into the service instances (see Registry Svc Instance) at runtime.
	// The value of a key here is the default value. The actual value can be overridden at deployment time.
	Envars []EnvVar `json:"envars,omitempty"`

	// Ports have host ports and internal ports currently but they
	// really only should have internal ports as host ports should be assigned
	// by the system. Host ports might go away in the future.
	Ports []PortMapping `json:"ports,omitempty"`
}

func (s Definition) GetId() string {
	return s.Id
}

type PortMapping struct {
	Internal uint16 `json:"internal" binding:"required"`
	Host     uint16 `json:"host" binding:"required"`
}

// EnvVar represents an environment variable inside the container.
type EnvVar struct {
	// Key is the environment variable name.
	Key string `json:"key" binding:"required"`

	// Value is the environment variable value.
	Value string `json:"value" binding:"required"`
}

type RepositorySpec struct {
	// URL is the URL to the repository
	URL string `json:"url,omitempty" example:"https://github.com/openorch/openorch.git" binding:"required"`

	// Version of the code to use
	Version string `json:"version,omitempty" example:"v1.0.0"`

	// Context is the path to the image build context
	BuildContext string `json:"buildContext,omitempty" example:"path/to/subfolder"`

	// ContainerFile is the path to the file that contains the container build instructions
	// Relative from the build context. By default, it is assumed to be a Dockerfile.
	ContainerFile string `json:"containerFile,omitempty" example:"docker/Dockerfile"`

	// Ports the container will listen on internally
	InternalPorts []uint16 `json:"ports" binding:"required"`
}

type ImageSpec struct {
	// Name is the container image name/URL to use for the container
	Name string `json:"name" example:"nginx:latest" binding:"required"`

	// InternalPorts are the ports the container will listen on internally
	InternalPorts []uint16 `json:"internalPorts" binding:"required"`
}

type APISpec struct {
	URL          string            `json:"url,omitempty"`          // URL to the OpenAPI file or other API definition
	ProtocolType string            `json:"protocolType,omitempty"` // Protocol type (e.g., OpenAPI, Swagger, etc.)
	Version      string            `json:"version,omitempty"`      // Version of the API specification (e.g., 3.0.0)
	Metadata     map[string]string `json:"metadata,omitempty"`     // Additional metadata about the API (e.g., author, license, etc.)
}

type Client struct {
	Language Language `json:"language,omitempty" example:"JavaScript"                    binding:"required"` // Programming language.
	URL      string   `json:"url,omitempty"      example:"https://example.com/client.js" binding:"required"` // The URL of the client.
}

type RuntimeType string

const (
	RuntimeDocker     RuntimeType = "Docker"
	RuntimeContainerd RuntimeType = "containerd"
	RuntimeCRIO       RuntimeType = "CRI-O"
)

type Language string

const (
	JavaScript Language = "JavaScript"
	Python     Language = "Python"
	Java       Language = "Java"
	CSharp     Language = "C#"
	CPlusPlus  Language = "C++"
	Ruby       Language = "Ruby"
	Go         Language = "Go"
	Swift      Language = "Swift"
	PHP        Language = "PHP"
	TypeScript Language = "TypeScript"
	Kotlin     Language = "Kotlin"
	Scala      Language = "Scala"
	Perl       Language = "Perl"
	Rust       Language = "Rust"
	Haskell    Language = "Haskell"
	Clojure    Language = "Clojure"
	Elixir     Language = "Elixir"
	ObjectiveC Language = "Objective-C"
	FSharp     Language = "F#"
)

type SaveDefinitionRequest struct {
	Definition Definition `json:"definition,omitempty"`
}

type SaveDefinitionResponse struct {
}

type ListDefinitionsRequest struct{}

type ListDefinitionsResponse struct {
	Definitions []*Definition `json:"definitions,omitempty"`
}
