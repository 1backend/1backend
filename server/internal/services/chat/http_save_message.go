/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package chatservice

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"time"

	openapi "github.com/1backend/1backend/clients/go"
	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	chat "github.com/1backend/1backend/server/internal/services/chat/types"
	"github.com/gorilla/mux"
)

// @ID saveMessage
// @Summary Save Message
// @Description Save a new message to a specific thread.
// @Tags Chat Svc
// @Accept json
// @Produce json
// @Param threadId path string true "Thread ID"
// @Param body body chat.SaveMessageRequest true "Save Message Request"
// @Success 200 {object} map[string]any "Message successfully added"
// @Failure 400 {string} string "Invalid JSON"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /chat-svc/thread/{threadId}/message [post]
func (a *ChatService) SaveMessage(
	w http.ResponseWriter,
	r *http.Request,
) {

	isAuthRsp, statusCode, err := a.options.PermissionChecker.HasPermission(
		r,
		chat.PermissionMessageCreate,
	)
	if err != nil {
		endpoint.WriteErr(w, statusCode, err)
		return
	}
	if !isAuthRsp.GetAuthorized() {
		endpoint.Unauthorized(w)
		return
	}

	req := chat.SaveMessageRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logger.Error(
			"Failed to decode request",
			slog.Any("error", err),
		)
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	defer r.Body.Close()

	vars := mux.Vars(r)
	threadId := vars["threadId"]
	req.ThreadId = threadId

	err = a.addMessage(r.Context(), &req)
	if err != nil {
		logger.Error("Error saving message", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}

	jsonData, _ := json.Marshal(map[string]any{})
	_, err = w.Write([]byte(jsonData))
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}

func (a *ChatService) addMessage(
	ctx context.Context,
	chatMessage *chat.SaveMessageRequest,
) error {
	if chatMessage.ThreadId == "" {
		return errors.New("empty chat message thread id")
	}
	if chatMessage.Id == "" {
		chatMessage.Id = sdk.Id("msg")
	}

	threads, err := a.threadsStore.Query(
		datastore.Equals(datastore.Field("id"), chatMessage.ThreadId),
	).Find()
	if err != nil {
		return err
	}

	if len(threads) == 0 {
		return errors.New("thread does not exist")
	}

	logger.Info("Saving chat message",
		slog.String("messageId", chatMessage.Id),
	)

	ev := chat.EventMessageAdded{
		ThreadId: chatMessage.ThreadId,
	}

	var m map[string]interface{}
	js, _ := json.Marshal(ev)
	json.Unmarshal(js, &m)

	_, err = a.options.ClientFactory.Client(client.WithToken(a.token)).
		FirehoseSvcAPI.PublishEvent(context.Background()).
		Event(openapi.FirehoseSvcEventPublishRequest{
			Event: &openapi.FirehoseSvcEvent{
				Name: openapi.PtrString(ev.Name()),
				Data: m,
			},
		}).
		Execute()
	if err != nil {
		logger.Error("Failed to publish firehose event", slog.Any("error", err))
	}

	now := time.Now()

	msg := &chat.Message{
		Id:        chatMessage.Id,
		ThreadId:  chatMessage.ThreadId,
		Text:      chatMessage.Text,
		UserId:    chatMessage.UserId,
		CreatedAt: now,
		UpdatedAt: now,
		FileIds:   chatMessage.FileIds,
		Meta:      chatMessage.Meta,
	}

	return a.messagesStore.Query(
		datastore.Equals(datastore.Field("id"), chatMessage.Id),
	).Upsert(msg)
}
