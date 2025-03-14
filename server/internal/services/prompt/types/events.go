package prompt_svc

const EventPromptAddedName = "promptAdded"

type EventPromptAdded struct {
	PromptId string `json:"promptId"`
}

func (e EventPromptAdded) Name() string {
	return EventPromptAddedName
}

const EventPromptRemovedName = "promptRemoved"

type EventPromptRemoved struct {
	PromptId string `json:"promptId"`
}

func (e EventPromptRemoved) Name() string {
	return EventPromptRemovedName
}

const EventPromptProcessingStartedName = "promptProcessingStarted"

type EventPromptProcessingStarted struct {
	PromptId string `json:"promptId"`
}

func (e EventPromptProcessingStarted) Name() string {
	return EventPromptProcessingStartedName
}

const EventPromptProcessingFinishedName = "promptProcessingFinished"

type EventPromptProcessingFinished struct {
	PromptId string `json:"promptId"`
	Error    string `json:"error,omitempty"`
}

func (e EventPromptProcessingFinished) Name() string {
	return EventPromptProcessingFinishedName
}
