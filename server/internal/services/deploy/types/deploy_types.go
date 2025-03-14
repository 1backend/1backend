/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package deploy_svc

type ErrorResponse struct {
	Error string `json:"error"`
}

type DeploymentStatus string

const (
	DeploymentStatusOK        DeploymentStatus = "OK"
	DeploymentStatusError     DeploymentStatus = "Error"
	DeploymentStatusPending   DeploymentStatus = "Pending"
	DeploymentStatusFailed    DeploymentStatus = "Failed"
	DeploymentStatusDeploying DeploymentStatus = "Deploying"
)

type Deployment struct {
	// ID of the deployment (e.g., "depl_dbOdi5eLQK")
	Id string `json:"id,omitempty" example:"depl_dbOdi5eLQK" binding:"required"`

	// DefinitionId is the id of the definition
	DefinitionId string `json:"definitionId,omitempty" example:"def_deBXZMpxrQ" binding:"required"`

	// Short name for easy reference (e.g., "user-service-v2")
	Name string `json:"name,omitempty" example:"user-service-v2"`

	// Description of what this deployment does
	Description string `json:"description,omitempty" example:"Handles user service requests"`

	// Number of container instances to run
	Replicas int `json:"replicas,omitempty"`

	// Deployment strategy (e.g., rolling update)
	Strategy *DeploymentStrategy `json:"strategy,omitempty"`

	// Resource requirements for each replica
	Resources *ResourceLimits `json:"resources,omitempty"`

	// Optional: Auto-scaling rules
	AutoScaling *AutoScalingConfig `json:"autoScaling,omitempty"`

	// Target deployment regions or clusters
	TargetRegions []TargetRegion `json:"targetRegions,omitempty"`

	// Current status of the deployment (e.g., "OK", "Error", "Pending")
	Status DeploymentStatus `json:"status,omitempty" example:"OK"`

	// Details provides additional information about the deployment's current state,
	// including both success and failure conditions (e.g., "Deployment in progress", "Error pulling image").
	Details string `json:"details,omitempty" example:"Deployment is in progress"`

	// Envars is a map of environment variables that will be passed down to service instances (see Registry Svc Instance)
	// Also see the Registry Svc Definition for required envars.
	Envars map[string]string `json:"envars,omitempty"`
}

func (d Deployment) GetId() string {
	return d.Id
}

type DeploymentStrategy struct {
	Type           StrategyType `json:"type,omitempty"`           // Deployment strategy type (RollingUpdate, Recreate, etc.)
	MaxUnavailable int          `json:"maxUnavailable,omitempty"` // Max unavailable replicas during update
	MaxSurge       int          `json:"maxSurge,omitempty"`       // Max extra replicas during update
}

type StrategyType string

const (
	StrategyRollingUpdate StrategyType = "RollingUpdate"
	StrategyRecreate      StrategyType = "Recreate"
)

type ResourceLimits struct {
	CPU    string `json:"cpu,omitempty"`    // CPU limit, e.g., "500m" for 0.5 cores
	Memory string `json:"memory,omitempty"` // Memory limit, e.g., "128Mi"
	VRAM   string `json:"vram,omitempty"`   // Optional: GPU VRAM requirement, e.g., "48GB"
}

type AutoScalingConfig struct {
	MinReplicas  int `json:"minReplicas,omitempty"`  // Minimum number of replicas to run
	MaxReplicas  int `json:"maxReplicas,omitempty"`  // Maximum number of replicas to run
	CPUThreshold int `json:"cpuThreshold,omitempty"` // CPU usage threshold for scaling (as a percentage)

	// @todo need to measure and describe the utilization of GPUs
}

type TargetRegion struct {
	Cluster string `json:"cluster,omitempty"` // Cluster or node where service should be deployed (e.g., "us-west1", "local-docker")
	Zone    string `json:"zone,omitempty"`    // Optional: Specific zone for the deployment
}

type ListDeploymentsRequest struct {
}

type ListDeploymentsResponse struct {
	Deployments []*Deployment `json:"deployments,omitempty"`
}

type SaveDeploymentRequest struct {
	Deployment *Deployment `json:"deployment,omitempty"`
}

type SaveDeploymentResponse struct {
}

type DeleteDeploymentRequest struct {
	DeploymentId string `json:"deploymentId,omitempty" binding:"required"`
}

type DeleteDeploymentResponse struct {
}
