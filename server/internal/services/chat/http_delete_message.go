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
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	chat "github.com/1backend/1backend/server/internal/services/chat/types"
	"github.com/gorilla/mux"
)

// @ID deleteMessage
// @Summary Delete a Message
// @Description Delete a specific message from a chat thread by its ID
// @Tags Chat Svc
// @Accept json
// @Produce json
// @Param messageId path string true "Message ID"
// @Success 200 {object} map[string]any "Message successfully deleted"
// @Failure 400 {string} string "Invalid JSON"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /chat-svc/message/{messageId} [delete]
func (a *ChatService) DeleteMessage(
	w http.ResponseWriter,
	r *http.Request,
) {
	isAuthResp, statusCode, err := a.options.PermissionChecker.HasPermission(
		r,
		chat.PermissionMessageDelete,
	)
	if err != nil {
		endpoint.WriteErr(w, statusCode, err)
		return
	}
	if !isAuthResp.GetAuthorized() {
		endpoint.Unauthorized(w)
		return
	}

	vars := mux.Vars(r)
	threadId := vars["threadId"]

	err = a.deleteMessage(threadId)
	if err != nil {
		logger.Error("Failed to delete message",
			slog.String("threadId", threadId),
			slog.Any("error", err),
		)
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

func (a *ChatService) deleteMessage(id string) error {
	return a.messagesStore.Query(
		datastore.Equals(datastore.Field("id"), id),
	).Delete()
}
