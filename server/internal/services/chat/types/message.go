/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package chat_svc

import (
	"time"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type Message struct {
	Id string `json:"id" example:"msg_emSOPlW58o" binding:"required"`

	// ThreadId of the message.
	ThreadId string `json:"threadId" example:"thr_emSOeEUWAg" binding:"required"`

	// Text content of the message eg. "Hi, what's up?"
	Text string `json:"text,omitempty"`

	// UserId is the id of the user who wrote the message.
	// For AI messages this field is empty.
	UserId string `json:"userId,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	// FileIds defines the file attachments the message has.
	FileIds []string `json:"fileIds,omitempty"`

	Meta map[string]interface{} `json:"meta,omitempty"`
}

func (c *Message) GetId() string {
	return c.Id
}

func (c *Message) GetUpdatedAt() string {
	return c.Id
}

type ByTime []*Message

func (a ByTime) Len() int      { return len(a) }
func (a ByTime) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a ByTime) Less(i, j int) bool {
	ti := a[i].CreatedAt
	tj := a[j].CreatedAt

	return ti.Before(tj)
}

type AddMessageRequest struct {
	Message *Message `json:"message"`
}

type AddMessageResponse struct{}

type GetMessagesRequest struct {
	ThreadId string `json:"threadId"`
}

type GetMessagesResponse struct {
	Messages []*Message `json:"messages"`
}

type DeleteMessageRequest struct {
	MessageId string `json:"messageId"`
}

type GetMessageResponse struct {
	Exists  bool     `json:"exists"`
	Message *Message `json:"message"`
}
