/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package allocator_test

import (
	"testing"

	openapi "github.com/openorch/openorch/clients/go"
	"github.com/openorch/openorch/server/internal/services/deploy/allocator"
	deploy "github.com/openorch/openorch/server/internal/services/deploy/types"
	"github.com/stretchr/testify/require"
)

func TestGenerateCommands_ScaleUp(t *testing.T) {
	nodes := []openapi.RegistrySvcNode{
		{
			Id: "node1",
			Usage: &openapi.RegistrySvcResourceUsage{
				Cpu: &openapi.RegistrySvcUsage{Percent: Float32Ptr(50)},
			},
		},
		{
			Id: "node2",
			Usage: &openapi.RegistrySvcResourceUsage{
				Cpu: &openapi.RegistrySvcUsage{Percent: Float32Ptr(60)},
			},
		},
	}
	serviceInstances := []openapi.RegistrySvcInstance{
		{
			Id:            "instance1",
			DeploymentId:  openapi.PtrString("service-A"),
			LastHeartbeat: StrPtr("valid"),
		},
	}
	deployments := []*deploy.Deployment{
		{Id: "service-A", Replicas: 3},
	}

	commands := allocator.GenerateCommands(nodes, serviceInstances, deployments)

	require.Equal(t, 2, len(commands))
	require.Equal(t, deploy.CommandType("START"), commands[0].Action)
	require.Equal(t, "service-A", *commands[0].DeploymentId)
	require.Equal(t, "node1", commands[0].NodeId)
	require.Equal(t, deploy.CommandType("START"), commands[1].Action)
	require.Equal(t, "service-A", *commands[1].DeploymentId)
	require.Equal(t, "node2", commands[1].NodeId)
}

func TestGenerateCommands_ScaleDown(t *testing.T) {
	nodes := []openapi.RegistrySvcNode{
		{
			Id: "node1",
			Usage: &openapi.RegistrySvcResourceUsage{
				Cpu: &openapi.RegistrySvcUsage{Percent: Float32Ptr(50)},
			},
		},
	}
	serviceInstances := []openapi.RegistrySvcInstance{
		{
			Id:            "instance1",
			DeploymentId:  openapi.PtrString("service-A"),
			LastHeartbeat: StrPtr("valid"),
		},
		{
			Id:            "instance2",
			DeploymentId:  openapi.PtrString("service-A"),
			LastHeartbeat: StrPtr("valid"),
		},
		{
			Id:            "instance3",
			DeploymentId:  openapi.PtrString("service-A"),
			LastHeartbeat: StrPtr("valid"),
		},
	}
	deployments := []*deploy.Deployment{
		{Id: "service-A", Replicas: 1},
	}

	commands := allocator.GenerateCommands(nodes, serviceInstances, deployments)

	require.Equal(t, 2, len(commands))
	require.Equal(t, deploy.CommandType("KILL"), commands[0].Action)
	require.Equal(t, "service-A", *commands[0].DeploymentId)
	require.Equal(t, deploy.CommandType("KILL"), commands[1].Action)
	require.Equal(t, "service-A", *commands[1].DeploymentId)
}

func TestGenerateCommands_KillUnhealthy(t *testing.T) {
	nodes := []openapi.RegistrySvcNode{
		{
			Id: "node1",
			Usage: &openapi.RegistrySvcResourceUsage{
				Cpu: &openapi.RegistrySvcUsage{Percent: Float32Ptr(50)},
			},
		},
	}
	serviceInstances := []openapi.RegistrySvcInstance{
		{Id: "instance1", DeploymentId: openapi.PtrString("service-A"), LastHeartbeat: nil},
	}
	deployments := []*deploy.Deployment{
		{Id: "service-A", Replicas: 1},
	}

	commands := allocator.GenerateCommands(nodes, serviceInstances, deployments)

	require.Equal(t, 1, len(commands))
	require.Equal(t, deploy.CommandType("KILL"), commands[0].Action)
	require.Equal(t, "service-A", *commands[0].DeploymentId)
	require.Equal(t, "instance1", *commands[0].InstanceId)
}

func TestGenerateCommands_NoAction(t *testing.T) {
	nodes := []openapi.RegistrySvcNode{
		{
			Id: "node1",
			Usage: &openapi.RegistrySvcResourceUsage{
				Cpu: &openapi.RegistrySvcUsage{Percent: Float32Ptr(50)},
			},
		},
	}
	serviceInstances := []openapi.RegistrySvcInstance{
		{
			Id:            "instance1",
			DeploymentId:  openapi.PtrString("service-A"),
			LastHeartbeat: StrPtr("valid"),
		},
	}
	deployments := []*deploy.Deployment{
		{Id: "service-A", Replicas: 1},
	}

	commands := allocator.GenerateCommands(nodes, serviceInstances, deployments)

	require.Equal(t, 0, len(commands))
}

func StrPtr(value string) *string {
	return &value
}

func Int64Ptr(value int64) *int64 {
	return &value
}

func Float32Ptr(value float32) *float32 {
	return &value
}
