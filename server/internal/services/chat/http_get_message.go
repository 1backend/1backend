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

// @ID getMessage
// @Summary Get Message
// @Description Fetch information about a specific chat message by its ID
// @Tags Chat Svc
// @Accept json
// @Produce json
// @Param messageId path string true "Message ID"
// @Success 200 {object} chat.GetMessageResponse "Message details successfully retrieved"
// @Failure 400 {string} string "Invalid JSON"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /chat-svc/message/{messageId} [get]
func (a *ChatService) GetMessage(
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

	vars := mux.Vars(r)
	messageId := vars["messageId"]

	message, found, err := a.getMessage(messageId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	jsonData, _ := json.Marshal(chat.GetMessageResponse{
		Exists:  found,
		Message: message,
	})
	w.Write(jsonData)
}

func (a *ChatService) getMessage(
	messageId string,
) (*chat.Message, bool, error) {
	messageI, found, err := a.messagesStore.Query(
		datastore.Equals(datastore.Field("id"), messageId),
	).FindOne()
	if err != nil {
		return nil, false, err
	}
	if !found {
		return nil, false, nil
	}

	return messageI.(*chat.Message), true, nil
}
