/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package chatservice

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	chat "github.com/1backend/1backend/server/internal/services/chat/types"
)

// @ID listMessages
// @Summary List Messages
// @Description Fetch messages (and associated assets) for a specific chat thread.
// @Tags Chat Svc
// @Accept json
// @Produce json
// @Param body body chat.ListMessagesRequest true "List Messages Request"
// @Success 200 {object} chat.ListMessagesResponse "Messages and assets successfully retrieved"
// @Failure 400 {string} string "Invalid JSON"
// @Failure 401 {string} string "unauthorized"
// @Failure 500 {string} string "internal server error"
// @Security BearerAuth
// @Router /chat-svc/messages [post]
func (a *ChatService) ListMessages(
	w http.ResponseWriter,
	r *http.Request,
) {
	isAuthRsp, statusCode, err := a.options.PermissionChecker.HasPermission(
		r,
		chat.PermissionMessageView,
	)
	if err != nil {
		endpoint.WriteErr(w, statusCode, err)
		return
	}
	if !isAuthRsp.Authorized {
		endpoint.Unauthorized(w)
		return
	}

	req := chat.ListMessagesRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	defer r.Body.Close()

	messages, err := a.listMessages(&req)
	if err != nil {
		logger.Error(
			"Failed to list messages",
			slog.String("threadId", req.ThreadId),
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	jsonData, _ := json.Marshal(chat.ListMessagesResponse{
		Messages: messages,
	})
	_, err = w.Write([]byte(jsonData))
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}

func (a *ChatService) listMessages(
	req *chat.ListMessagesRequest,
) ([]*chat.Message, error) {
	filters := []datastore.Filter{}

	if req.ThreadId != "" {
		filters = append(filters,
			datastore.Equals(
				datastore.Field("threadId"),
				req.ThreadId),
		)
	}

	if req.Ids != nil {
		ids := []any{}
		for _, id := range req.Ids {
			ids = append(ids, id)
		}

		filters = append(filters,
			datastore.IsInList(
				datastore.Field("id"),
				ids...),
		)
	}

	messageIs, err := a.messagesStore.Query(
		filters...,
	).OrderBy(datastore.OrderByField("createdAt", false)).Find()
	if err != nil {
		return nil, err
	}

	messages := []*chat.Message{}
	for _, messageI := range messageIs {
		messages = append(messages, messageI.(*chat.Message))
	}

	return messages, nil
}
