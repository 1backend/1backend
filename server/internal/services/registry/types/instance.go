/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package registry_svc

import (
	"errors"
	"fmt"
	"time"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type InstanceStatus string

const (
	InstanceStatusUnknown     InstanceStatus = "Unknown"     // Default state when the instance status is not yet determined
	InstanceStatusHealthy     InstanceStatus = "Healthy"     // Instance is fully operational and responding as expected
	InstanceStatusDegraded    InstanceStatus = "Degraded"    // Instance is responding but with performance issues or partial failures
	InstanceStatusUnreachable InstanceStatus = "Unreachable" // Instance is not reachable, possibly down or with network issues
	InstanceStatusError       InstanceStatus = "Error"       // Instance encountered errors or failed multiple health checks
)

type Instance struct {
	// Required: ID of the instance
	Id string `json:"id,omitempty" example:"inst_di9riJEvH2" binding:"required"`

	// Full address URL of the instance.
	URL string `json:"url,omitempty" example:"https://myserver.com:5981" binding:"required"`

	// The ID of the deployment that this instance is an instance of.
	// Only instances deployed by OpenOrch have a DeploymentId.
	// Services can be deployed through other means (Docker Compose, K8s, anything),
	// in that case they self-register and will not have a DeploymentId.
	DeploymentId string `json:"deploymentId,omitempty" example:"depl_deBUCtJirc"`

	// NodeURL is the URL of the OpenOrch server the instance is running on.
	// To have a NodeURL the instance must either:
	// - Be deployed by OpenOrch
	// - Declare the OpenOrch server URL when registering its instance
	NodeURL string `json:"nodeUrl,omitempty" example:"https://myserver.com:58231"`

	// Last time the instance gave a sign of life
	LastHeartbeat time.Time `json:"lastHeartbeat,omitempty"`

	// Path of the instance address. Optional (e.g., "/api")
	Path string `json:"path,omitempty" example:"/your-svc"`

	// URL alternative fields

	// Scheme of the instance address. Required if URL is not provided.
	Scheme string `json:"scheme,omitempty" example:"https"`

	// Host of the instance address. Required if URL is not provided
	Host string `json:"host,omitempty" example:"myserver.com"`

	// IP of the instance address. Optional: to register by IP instead of host
	IP string `json:"ip,omitempty" example:"8.8.8.8"`

	// Port of the instance address. Required if URL is not provided
	Port int `json:"port,omitempty" example:"8080"`

	// Status
	Status InstanceStatus `json:"status,omitempty" example:"Healthy" binding:"required"`

	// Details
	Details string `json:"details,omitempty" example:"Instance is healthy"`

	// Slug of the account that owns this instance
	// Services that want to be proxied by their slug are advised to self register
	// their instance at startup.
	// Keep in mind, instances might be deployed by OpenOrch yet they still won't be OpenOrch services
	// and they won't have slugs. Think NGINX, MySQL, etc.
	Slug string `json:"slug,omitempty" example:"my-svc"`

	// Tags are used to filter instances
	Tags []string `json:"tags,omitempty" example:"tag1,tag2"`
}

func (s *Instance) GetId() string {
	return s.Id
}

func (s *Instance) BuildURL() string {
	if s.URL != "" {
		return s.URL
	}

	var constructedURL string
	if s.Host != "" {
		constructedURL = fmt.Sprintf("%s://%s:%d", s.Scheme, s.Host, s.Port)
	} else {
		constructedURL = fmt.Sprintf("%s://%s:%d", s.Scheme, s.IP, s.Port)
	}

	if s.Path != "" {
		return fmt.Sprintf("%s/%s", constructedURL, s.Path)
	}

	return constructedURL
}

var ErrNotFound = errors.New("service not found")

// RegisterInstanceRequest represents the request body to register a service.
//
// The user must provide either:
// 1. A full URL (e.g., "https://myserver.com:5981") -OR-
// 2. A combination of the following fields:
//   - scheme: "http" or "https" (required if URL is not provided)
//   - host: the domain of the service (required if URL is not provided)
//   - port: the port number (required if URL is not provided)
//
// Additionally, if both host and port are provided, they cannot both be specified at the same time.
// The IP field is optional and can be used for registration by IP instead of host.
type RegisterInstanceRequest struct {
	Id string `json:"id,omitempty" example:"inst_di9riJEvH2"`

	// The ID of the deployment that this instance is an instance of.
	DeploymentId string `json:"deploymentId,omitempty" example:"depl_deBUCtJirc"`

	// Full address URL of the instance.
	URL string `json:"url,omitempty" example:"https://myserver.com:5981" binding:"required"`

	// Scheme of the instance address. Required if URL is not provided.
	Scheme string `json:"scheme,omitempty" example:"https"`

	// Host of the instance address. Required if URL is not provided
	Host string `json:"host,omitempty" example:"myserver.com"`

	// IP of the instance address. Optional: to register by IP instead of host
	IP string `json:"ip,omitempty" example:"8.8.8.8"`

	// Port of the instance address. Required if URL is not provided
	Port int `json:"port,omitempty" example:"8080"`

	// Path of the instance address. Optional (e.g., "/api")
	Path string `json:"path,omitempty" example:"/your-svc"`
}

type RegisterInstanceResponse struct {
}

type ListInstancesResponse struct {
	Instances []*Instance `json:"instances,omitempty"`
}
