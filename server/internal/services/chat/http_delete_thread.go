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
	chattypes "github.com/1backend/1backend/server/internal/services/chat/types"
	"github.com/gorilla/mux"
)

// @ID deleteThread
// @Summary Delete a Thread
// @Description Delete a specific chat thread by its ID
// @Tags Chat Svc
// @Accept json
// @Produce json
// @Param threadId path string true "Thread ID"
// @Success 200 {object} map[string]any "Thread successfully deleted"
// @Failure 400 {string} string "Invalid JSON"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /chat-svc/thread/{threadId} [delete]
func (a *ChatService) DeleteThread(
	w http.ResponseWriter,
	r *http.Request,
) {
	isAuthRsp, statusCode, err := a.permissionChecker.HasPermission(
		r,
		chattypes.PermissionThreadDelete,
	)
	if err != nil {
		endpoint.WriteErr(w, statusCode, err)
		return
	}
	if !isAuthRsp.GetAuthorized() {
		endpoint.Unauthorized(w)
		return
	}

	vars := mux.Vars(r)

	err = a.deleteThread(vars["threadId"])
	if err != nil {
		logger.Error(
			"Failed to delete thread",
			slog.String("threadId", vars["threadId"]),
			slog.String("error", err.Error()),
		)
		endpoint.InternalServerError(w)
		return
	}

	jsonData, _ := json.Marshal(map[string]any{})
	w.Write(jsonData)
}

func (a *ChatService) deleteThread(id string) error {
	return a.threadsStore.Query(
		datastore.Equals(datastore.Field("id"), id),
	).Delete()
}
