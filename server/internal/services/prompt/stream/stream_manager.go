/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/

// This package is named "prompt_svc_*" so the generated and exported
// swaggo Chunk type is prefixed with PromptSvcStream
package prompt_svc_stream

import (
	"strings"
	"sync"
)

// ChunkType is an enumeration of possible stream event types.
type ChunkType string

const (
	// ChunkTypeText indicates a normal text chunk event.
	ChunkTypeProgress ChunkType = "progress"

	// ChunkTypeDone indicates the stream has completed.
	ChunkTypeDone ChunkType = "done"
)

// Chunk represents a streaming event.
type Chunk struct {
	// TextChunk contains a part of the text output from the stream.
	Text string `json:"text,omitempty"`

	// MessageId is the ChatSvc Message id that the chunk is part of.
	// Might only be available for "done" chunks.
	MessageId string `json:"messageId,omitempty"`

	// Type indicates the type of the stream event (e.g., text, done).
	Type ChunkType `json:"type"`
}

type SubscriberChan chan *Chunk

// StreamManager aims to be a generic streamer tool for
// different AI engines and architectures (ie `Text-to-Text`, `Text-to-Image` etc.)
type StreamManager struct {
	streams map[string][]SubscriberChan
	history map[string][]*Chunk
	lock    sync.RWMutex
}

func NewStreamManager() *StreamManager {
	return &StreamManager{
		streams: make(map[string][]SubscriberChan),
		history: make(map[string][]*Chunk),
	}
}

// Subscribe to prompt responses for a given thread.
// It's acceptable to subscribe even if the thread doesn't exist yet.
// When you subscribe, any pending prompt responses that haven't been saved into a `ChatSvc.Message` are immediately streamed to the subscriber.
func (sm *StreamManager) Subscribe(
	threadId string,
	subscriber SubscriberChan,
) {
	sm.lock.Lock()
	defer sm.lock.Unlock()

	// Add subscriber to the list
	sm.streams[threadId] = append(sm.streams[threadId], subscriber)

	// Send historical messages to the new subscriber
	go func() {
		sm.lock.RLock()
		defer sm.lock.RUnlock()
		for _, msg := range sm.history[threadId] {
			subscriber <- msg
		}
	}()
}

func (sm *StreamManager) Unsubscribe(
	threadId string,
	subscriber SubscriberChan,
) {
	sm.lock.Lock()
	defer sm.lock.Unlock()
	subs := sm.streams[threadId]
	for i, sub := range subs {
		if sub == subscriber {
			sm.streams[threadId] = append(subs[:i], subs[i+1:]...)
			close(
				subscriber,
			) // Close the channel to signify no more data will be sent
			break
		}
	}
}

func (sm *StreamManager) Broadcast(
	threadId string,
	response *Chunk,
) {
	sm.lock.Lock()
	defer sm.lock.Unlock()
	if subscribers, ok := sm.streams[threadId]; ok {
		for _, subscriber := range subscribers {
			select {
			case subscriber <- response:
			default:
				// Handle full channel or unresponsive subscriber
			}
		}
	}

	// Append the new response to the history
	sm.history[threadId] = append(sm.history[threadId], response)
}

// DeleteHistory resets the history of a thread. Use this method from AI engine processors
// when the engine finished processing, right at the time when you save the outputs into a `ChatSvc.Message`.
func (sm *StreamManager) DeleteHistory(threadId string) {
	delete(sm.history, threadId)
}

func (sm *StreamManager) ConcatHistoryText(threadId string) string {
	responses := sm.history[threadId]

	var result strings.Builder

	first := true
	for _, v := range responses {
		if len(v.Text) == 0 {
			continue
		}

		// @todo doing this escaping here is wrong

		var textToAdd string
		if strings.Contains(result.String(), "```") {
			// Handling for inline code formatting if the resulting string is already within a code block
			count := strings.Count(result.String(), "```")
			if count%2 == 1 { // If the count of ``` is odd, we are inside a code block
				textToAdd = v.Text // No escaping needed inside code block
			} else {
				textToAdd = escapeHtml(v.Text) // Apply HTML escaping when outside code blocks
			}
		} else {
			textToAdd = escapeHtml(v.Text) // Apply HTML escaping if there is no code block
		}

		if first {
			textToAdd = strings.TrimLeft(textToAdd, " ")
			first = false
		}

		result.WriteString(textToAdd)

		if v.Type == ChunkTypeDone {
			break
		}
	}

	return result.String()

}

func escapeHtml(input string) string {
	replacer := strings.NewReplacer(
		"&", "&amp;",
		"<", "&lt;",
		">", "&gt;",
		"\"", "&quot;",
		"'", "&#39;",
	)
	return replacer.Replace(input)
}
