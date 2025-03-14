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
	"net/http"

	"github.com/gorilla/mux"
	sdk "github.com/openorch/openorch/sdk/go"
	"github.com/openorch/openorch/sdk/go/datastore"
	chat "github.com/openorch/openorch/server/internal/services/chat/types"
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

	isAuthRsp, _, err := a.clientFactory.Client(sdk.WithTokenFromRequest(r)).
		UserSvcAPI.IsAuthorized(r.Context(), *chat.PermissionMessageDelete.Id).
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

	vars := mux.Vars(r)
	threadId := vars["threadId"]

	err = a.deleteMessage(threadId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	jsonData, _ := json.Marshal(map[string]any{})
	w.Write(jsonData)
}

func (a *ChatService) deleteMessage(id string) error {
	return a.messagesStore.Query(
		datastore.Equals(datastore.Field("id"), id),
	).Delete()
}
