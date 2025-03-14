/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package registry_svc

import "time"

type Node struct {
	// Required: ID of the instance
	Id string `json:"id,omitempty" example:"node_di9riJEvH2" binding:"required"`

	// URL of the daemon running on the node.
	// If not configured defaults to hostname + default OpenOrch server port.
	URL string `json:"url" binding:"required"`

	// The availability zone of the node
	AvailabilityZone string `json:"availabilityZone,omitempty"`

	// Last time the instance gave a sign of life
	LastHeartbeat time.Time `json:"lastHeartbeat,omitempty"`

	// The region of the node
	Region string `json:"region,omitempty"`

	// Resource usage metrics of the node.
	Usage ResourceUsage `json:"usage,omitempty"`

	// List of GPUs available on the node
	GPUs []*GPU `json:"gpus,omitempty"`
}

func (n Node) GetId() string {
	return n.Id
}

// Usage represents the usage metrics for a resource.
type Usage struct {
	Used    uint64  `json:"used"    format:"int64"` // Used amount (in bytes).
	Total   uint64  `json:"total"   format:"int64"` // Total available amount (in bytes).
	Percent float64 `json:"percent"`                // Usage percentage.
}

type ResourceUsage struct {
	CPU    Usage `json:"cpu"`    // CPU usage metrics.
	Memory Usage `json:"memory"` // Memory usage metrics.
	Disk   Usage `json:"disk"`   // Disk usage metrics.
}

type GPU struct {
	// Id Node.URL + IntraNodeId
	Id               string    `json:"id"`
	IntraNodeId      int       `json:"intraNodeId"`
	Name             string    `json:"name"`
	BusId            string    `json:"busId"`
	Temperature      float64   `json:"temperature"`
	PerformanceState string    `json:"performanceState"`
	PowerUsage       float64   `json:"powerUsage"`
	PowerCap         float64   `json:"powerCap"`
	MemoryUsage      int       `json:"memoryUsage"`
	MemoryTotal      int       `json:"memoryTotal"`
	GPUUtilization   float64   `json:"gpuUtilization"`
	ComputeMode      string    `json:"computeMode"`
	ProcessDetails   []Process `json:"processDetails,omitempty"`
	CudaVersion      string    `json:"cudaVersion,omitempty"`
}

type Process struct {
	Pid         int    `json:"pid"`
	ProcessName string `json:"processName"`
	MemoryUsage int    `json:"memoryUsage"`
}

type ListNodesRequest struct {
	// Node IDs to filter on
	Ids []string `json:"ids"`
}

type ListNodesResponse struct {
	Nodes []*Node `json:"nodes"`
}

type NodeSelfRequest struct{}

type NodeSelfResponse struct {
	Node Node `json:"node" binding:"required"`
}
