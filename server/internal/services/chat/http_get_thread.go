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

	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/datastore"
	chat "github.com/1backend/1backend/server/internal/services/chat/types"
	"github.com/gorilla/mux"
)

// @ID getThread
// @Summary Get Thread
// @Description Fetch information about a specific chat thread by its ID
// @Tags Chat Svc
// @Accept json
// @Produce json
// @Param threadId path string true "Thread ID"
// @Success 200 {object} chat.GetThreadResponse "Thread details successfully retrieved"
// @Failure 400 {string} string "Invalid JSON"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /chat-svc/thread/{threadId} [get]
func (a *ChatService) GetThread(
	w http.ResponseWriter,
	r *http.Request,
) {

	isAuthRsp, _, err := a.clientFactory.Client(client.WithTokenFromRequest(r)).
		UserSvcAPI.HasPermission(r.Context(), chat.PermissionThreadCreate).
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

	thread, found, err := a.getThread(threadId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	jsonData, _ := json.Marshal(chat.GetThreadResponse{
		Exists: found,
		Thread: thread,
	})
	w.Write(jsonData)
}

func (a *ChatService) getThread(
	threadId string,
) (*chat.Thread, bool, error) {
	threadI, found, err := a.threadsStore.Query(
		datastore.Equals(datastore.Field("id"), threadId),
	).FindOne()
	if err != nil {
		return nil, false, err
	}
	if !found {
		return nil, false, nil
	}

	return threadI.(*chat.Thread), true, nil
}
