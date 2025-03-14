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

// @ID getMessages
// @Summary List Messages
// @Description Fetch messages (and associated assets) for a specific chat thread.
// @Tags Chat Svc
// @Accept json
// @Produce json
// @Param threadId path string true "Thread ID"
// @Success 200 {object} chat.GetMessagesResponse "Messages and assets successfully retrieved"
// @Failure 400 {string} string "Invalid JSON"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /chat-svc/thread/{threadId}/messages [post]
func (a *ChatService) GetMessages(
	w http.ResponseWriter,
	r *http.Request,
) {

	isAuthRsp, _, err := a.clientFactory.Client(sdk.WithTokenFromRequest(r)).
		UserSvcAPI.IsAuthorized(r.Context(), *chat.PermissionMessageView.Id).
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

	threadId := mux.Vars(r)["threadId"]

	messages, err := a.getMessages(threadId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	jsonData, _ := json.Marshal(chat.GetMessagesResponse{
		Messages: messages,
	})
	w.Write(jsonData)
}

func (a *ChatService) getMessages(
	threadId string,
) ([]*chat.Message, error) {
	messageIs, err := a.messagesStore.Query(
		datastore.Equals(datastore.Field("threadId"), threadId),
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
