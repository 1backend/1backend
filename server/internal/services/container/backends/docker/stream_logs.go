package dockerbackend

import (
	"bufio"
	"context"
	"log"
	"sync"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	dockercontainer "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/events"
	"github.com/docker/docker/client"
	"github.com/openorch/openorch/server/internal/services/container/logaccumulator"
)

// Start listening to Docker logs
func StartDockerLogListener(cli *client.Client, la *logaccumulator.LogAccumulator) {
	ctx := context.Background()
	eventChan, errChan := cli.Events(ctx, types.EventsOptions{})

	activeContainers := make(map[string]bool)
	var mu sync.Mutex

	containers, err := cli.ContainerList(ctx, dockercontainer.ListOptions{})
	if err != nil {
		log.Fatal("Error listing containers:", err)
	}
	for _, container := range containers {
		go streamLogs(cli, la, container.ID, &mu, activeContainers)
	}

	for {
		select {
		case event := <-eventChan:
			if event.Type == events.ContainerEventType {
				if event.Action == "start" {
					go streamLogs(cli, la, event.Actor.ID, &mu, activeContainers)
				} else if event.Action == "die" {
					mu.Lock()
					delete(activeContainers, event.Actor.ID)
					mu.Unlock()
				}
			}
		case err := <-errChan:
			if err != nil {
				log.Println("Docker event listener error:", err)
			}
		}
	}
}

// Stream logs from containers
func streamLogs(cli *client.Client, la *logaccumulator.LogAccumulator, containerID string, mu *sync.Mutex, activeContainers map[string]bool) {
	mu.Lock()
	if activeContainers[containerID] {
		mu.Unlock()
		return
	}
	activeContainers[containerID] = true
	mu.Unlock()

	ctx := context.Background()
	options := container.LogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow:     true,
		Timestamps: false,
	}

	logs, err := cli.ContainerLogs(ctx, containerID, options)
	if err != nil {
		log.Println("Error getting logs:", err)
		return
	}
	defer logs.Close()

	scanner := bufio.NewScanner(logs)
	for scanner.Scan() {
		entry := logaccumulator.LogEntry{
			ProducerID: containerID,
			Message:    scanner.Text(),
		}
		la.AddLog(entry)
	}

	if err := scanner.Err(); err != nil {
		log.Println("Error reading log stream:", err)
	}

	mu.Lock()
	delete(activeContainers, containerID)
	mu.Unlock()
}
