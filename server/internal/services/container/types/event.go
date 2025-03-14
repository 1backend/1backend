package container_svc

// @todo nothing to trigger this yet
const EventDockerInfoUpdatedName = "dockerInfoUpdated"

type EventDockerInfoUpdated struct {
	ThreadId string `json:"threadId"`
}

func (e EventDockerInfoUpdated) Name() string {
	return EventDockerInfoUpdatedName
}
