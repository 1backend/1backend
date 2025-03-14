package dockerbackend

import (
	"context"
	"log"
	"strconv"
	"strings"
	"sync"

	"github.com/docker/docker/api/types"
	dockerapitypes "github.com/docker/docker/api/types"
	dockercontainer "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/events"
	"github.com/docker/docker/client"
	container "github.com/openorch/openorch/server/internal/services/container/types"
	"github.com/samber/lo"
)

type ContainerTracker struct {
	mu         sync.Mutex
	containers map[string]container.Container
}

func NewContainerTracker() *ContainerTracker {
	return &ContainerTracker{
		containers: map[string]container.Container{},
	}
}

func StartDockerContainerTracker(cli *client.Client, tracker *ContainerTracker) {
	ctx := context.Background()
	eventChan, errChan := cli.Events(ctx, types.EventsOptions{})

	// Initialize active container list
	containers, err := cli.ContainerList(ctx, dockercontainer.ListOptions{All: true})
	if err != nil {
		log.Fatal("Error listing containers:", err)
	}
	for _, c := range containers {
		tracker.AddContainerFromList(c)
	}

	for {
		select {
		case event := <-eventChan:
			if event.Type == events.ContainerEventType {
				switch event.Action {
				case "start":
					go func(containerID string) {
						info, err := cli.ContainerInspect(ctx, containerID)
						if err != nil {
							log.Println("Error inspecting container:", err)
							return
						}
						tracker.AddContainerFromInspect(info)
					}(event.Actor.ID)
				case "die":
					tracker.RemoveContainer(event.Actor.ID)
				}
			}
		case err := <-errChan:
			if err != nil {
				log.Println("Docker event listener error:", err)
			}
		}
	}
}

func (t *ContainerTracker) AddContainerFromList(c types.Container) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.containers[c.ID] = container.Container{
		Id:    c.ID,
		Names: cleanNames(c.Names),
		Image: c.Image,
		Ports: extractPortsFromList(c.Ports),
		// Labels: c.Labels,
		Status: c.State,
		//Envs:   convertEnvToEnvVar(c.Env), // Convert Env into your custom EnvVar struct
		Capabilities: &container.Capabilities{
			//GPUEnabled: checkGPUEnabled(info),
		},
	}
}

func (t *ContainerTracker) AddContainerFromInspect(info types.ContainerJSON) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.containers[info.ID] = container.Container{
		Id: info.ID,
		Names: []string{
			cleanName(info.Name, 0),
		},
		Image: info.Config.Image,
		Ports: extractPortsFromInspect(info),
		Hash:  info.Config.Labels["hash"],
		// Labels: info.Config.Labels,
		Envs: convertEnvToEnvVar(info.Config.Env), // Convert Env into your custom EnvVar struct
		Capabilities: &container.Capabilities{
			GPUEnabled: checkGPUEnabled(info),
		},
		Status: info.State.Status,
	}
}

// Convert the Docker environment variables (e.g., "KEY=VALUE") to your custom EnvVar format
func convertEnvToEnvVar(env []string) []container.EnvVar {
	var envVars []container.EnvVar
	for _, e := range env {
		parts := strings.SplitN(e, "=", 2)
		if len(parts) == 2 {
			envVars = append(envVars, container.EnvVar{
				Key:   parts[0],
				Value: parts[1],
			})
		}
	}
	return envVars
}

func (t *ContainerTracker) RemoveContainer(containerID string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	delete(t.containers, containerID)
}

func (t *ContainerTracker) GetContainers() []container.Container {
	t.mu.Lock()
	defer t.mu.Unlock()

	var result []container.Container
	for _, c := range t.containers {
		result = append(result, c)
	}
	return result
}

func cleanNames(names []string) []string {
	return lo.Map(names, cleanName)
}

func cleanName(s string, index int) string {
	return strings.Trim(s, "/")
}

func extractPortsFromList(ports []dockerapitypes.Port) []container.PortMapping {
	ret := []container.PortMapping{}

	for _, port := range ports {
		ret = append(ret, container.PortMapping{
			Internal: port.PrivatePort,
			Host:     port.PublicPort,
		})
	}

	return ret
}

func extractPortsFromInspect(info types.ContainerJSON) []container.PortMapping {
	ret := []container.PortMapping{}

	if len(info.NetworkSettings.Ports) > 0 {
		for port, bindings := range info.NetworkSettings.Ports {
			if len(bindings) > 0 {
				hostPort, err := strconv.Atoi(bindings[0].HostPort)
				if err != nil {
					// Handle error if HostPort conversion fails
					continue
				}
				privatePort, err := strconv.Atoi(port.Port())
				if err != nil {
					// Handle error if private port conversion fails
					continue
				}

				ret = append(ret, container.PortMapping{
					Host:     uint16(hostPort),
					Internal: uint16(privatePort),
				})
			}
		}
	}

	return ret
}

func extractKeeps(info types.ContainerJSON) []string {
	var keeps []string
	for _, mount := range info.Mounts {
		keeps = append(keeps, mount.Destination)
	}
	return keeps
}

func checkGPUEnabled(info types.ContainerJSON) bool {
	return len(info.HostConfig.DeviceRequests) > 0
}
