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

	"github.com/gorilla/mux"
	openapi "github.com/openorch/openorch/clients/go"
	sdk "github.com/openorch/openorch/sdk/go"
	"github.com/openorch/openorch/sdk/go/datastore"
	"github.com/openorch/openorch/sdk/go/logger"
	chat "github.com/openorch/openorch/server/internal/services/chat/types"
)

// @ID addMessage
// @Summary Add Message
// @Description Add a new message to a specific thread.
// @Tags Chat Svc
// @Accept json
// @Produce json
// @Param threadId path string true "Thread ID"
// @Param body body chat.AddMessageRequest true "Add Message Request"
// @Success 200 {object} map[string]any "Message successfully added"
// @Failure 400 {string} string "Invalid JSON"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /chat-svc/thread/{threadId}/message [post]
func (a *ChatService) AddMessage(
	w http.ResponseWriter,
	r *http.Request,
) {

	isAuthRsp, _, err := a.clientFactory.Client(sdk.WithTokenFromRequest(r)).
		UserSvcAPI.IsAuthorized(r.Context(), *chat.PermissionMessageCreate.Id).
		Execute()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if !isAuthRsp.GetAuthorized() {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`Unauthorized`))
		return
	}

	req := chat.AddMessageRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	vars := mux.Vars(r)
	threadId := vars["threadId"]
	req.Message.ThreadId = threadId

	err = a.addMessage(r.Context(), req.Message)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	jsonData, _ := json.Marshal(map[string]any{})
	w.Write(jsonData)
}

func (a *ChatService) addMessage(
	ctx context.Context,
	chatMessage *chat.Message,
) error {
	if chatMessage.ThreadId == "" {
		return errors.New("empty chat message thread id")
	}
	if chatMessage.Id == "" {
		chatMessage.Id = sdk.Id("msg")
	}
	if chatMessage.CreatedAt.IsZero() {
		chatMessage.CreatedAt = time.Now()
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

	_, err = a.clientFactory.Client(sdk.WithToken(a.token)).
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

	return a.messagesStore.Query(
		datastore.Equals(datastore.Field("id"), chatMessage.Id),
	).Upsert(chatMessage)
}
