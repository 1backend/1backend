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
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	chat "github.com/1backend/1backend/server/internal/services/chat/types"
)

// @ID saveThread
// @Summary Save Thread
// @Description Create or update a chat thread.
// @Decription
// @Description Requires the `chat-svc:thread:edit` permission.
// @Tags Chat Svc
// @Accept json
// @Produce json
// @Param body body chat.SaveThreadRequest true "Save Thread Request"
// @Success 200 {object} chat.SaveThreadResponse "Thread successfully created"
// @Failure 400 {string} string "Invalid JSON"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /chat-svc/thread [post]
func (a *ChatService) SaveThread(
	w http.ResponseWriter,
	r *http.Request,
) {

	isAuthRsp, statusCode, err := a.permissionChecker.HasPermission(
		r,
		chat.PermissionThreadEdit,
	)
	if err != nil {
		endpoint.WriteErr(w, statusCode, err)
		return
	}
	if !isAuthRsp.GetAuthorized() {
		endpoint.Unauthorized(w)
		return
	}

	req := chat.SaveThreadRequest{}
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

	if req.Id != "" {
		threads, err := a.listThreads(
			isAuthRsp.User.Id,
			&chat.ListThreadsRequest{
				Ids: []string{req.Id},
			},
		)
		if err != nil {
			endpoint.WriteErr(w, http.StatusInternalServerError, err)
			return
		}
		if len(threads) > 0 {
			thread, err := a.updateThread(isAuthRsp.User.Id, threads[0], &req)
			if err != nil {
				endpoint.WriteErr(w, http.StatusInternalServerError, err)
				return
			}
			jsonData, _ := json.Marshal(chat.SaveThreadResponse{
				Thread: thread,
			})

			w.Write(jsonData)
			return
		}

	}

	req.UserIds = append(req.UserIds, isAuthRsp.User.Id)

	thread, err := a.addThread(r.Context(), &req)
	if err != nil {
		logger.Error("Error saving thread", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}

	jsonData, _ := json.Marshal(chat.SaveThreadResponse{
		Thread: thread,
	})
	w.Write(jsonData)
}

func (a *ChatService) addThread(
	ctx context.Context,
	chatThread *chat.SaveThreadRequest,
) (*chat.Thread, error) {
	if chatThread.Id == "" {
		chatThread.Id = sdk.Id("thr")
	}
	if chatThread.Title == "" {
		chatThread.Title = "New chat"
	}

	if len(chatThread.UserIds) == 0 {
		return nil, errors.New("no user ids")
	}

	now := time.Now()

	thread := &chat.Thread{
		Id:        chatThread.Id,
		Title:     chatThread.Title,
		CreatedAt: now,
		UpdatedAt: now,
		UserIds:   chatThread.UserIds,
	}

	err := a.threadsStore.Create(thread)
	if err != nil {
		return nil, err
	}

	ev := chat.EventThreadAdded{
		ThreadId: chatThread.Id,
	}

	var m map[string]interface{}
	js, _ := json.Marshal(ev)
	json.Unmarshal(js, &m)

	_, err = a.clientFactory.Client(client.WithToken(a.token)).
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

	return thread, nil
}
