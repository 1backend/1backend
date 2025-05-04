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
)

// @ID listThreads
// @Summary List Threads
// @Description Fetch all chat threads associated with a specific user
// @Tags Chat Svc
// @Accept json
// @Produce json
// @Param body body chat.ListThreadsRequest true "List Threads Request"
// @Success 200 {object} chat.ListThreadsResponse "Threads successfully retrieved"
// @Failure 400 {string} string "Invalid JSON"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /chat-svc/threads [post]
func (a *ChatService) ListThreads(
	w http.ResponseWriter,
	r *http.Request,
) {

	isAuthRsp, _, err := a.clientFactory.Client(client.WithTokenFromRequest(r)).
		UserSvcAPI.HasPermission(r.Context(), chat.PermissionThreadView).
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

	req := chat.ListThreadsRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	threads, err := a.listThreads(isAuthRsp.User.Id, &req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	jsonData, _ := json.Marshal(chat.ListThreadsResponse{
		Threads: threads,
	})
	w.Write(jsonData)
}

func (a *ChatService) listThreads(
	callerUserId string,
	req *chat.ListThreadsRequest,
) ([]*chat.Thread, error) {
	filters := []datastore.Filter{
		datastore.Equals(datastore.Field("userIds"), callerUserId),
	}

	if req.Ids != nil {
		ids := []any{}
		for _, id := range req.Ids {
			ids = append(ids, id)
		}
		filters = append(filters,
			datastore.IsInList(
				datastore.Field("id"),
				ids...,
			),
		)
	}

	threadIs, err := a.threadsStore.Query(
		filters...,
	).
		OrderBy(
			datastore.OrderByField("createdAt", true),
		).Find()
	if err != nil {
		return nil, err
	}

	threads := []*chat.Thread{}
	for _, threadI := range threadIs {
		threads = append(threads, threadI.(*chat.Thread))
	}

	return threads, nil
}
