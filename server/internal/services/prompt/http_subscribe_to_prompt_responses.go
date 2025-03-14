/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package promptservice

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	openapi "github.com/openorch/openorch/clients/go"
	sdk "github.com/openorch/openorch/sdk/go"
	"github.com/openorch/openorch/sdk/go/logger"

	streammanager "github.com/openorch/openorch/server/internal/services/prompt/stream"
	prompt "github.com/openorch/openorch/server/internal/services/prompt/types"
)

// @ID subscribeToPromptResponses
// @Summary Subscribe to Prompt Responses by Thread
// @Description Subscribe to prompt responses by thread via Server-Sent Events (SSE).
// @Description You can subscribe to threads before they are created.
// @Description The streamed strings are of type `StreamChunk`, see the PromptTypes endpoint for more details.
// @Tags Prompt Svc
// @Param threadId path string true "Thread ID"
// @Success 200 {string} string "Streaming response"
// @Failure 400 {object} prompt.ErrorResponse "Missing threadId parameter"
// @Failure 401 {object} prompt.ErrorResponse "Unauthorized"
// @Security BearerAuth
// @Router /prompt-svc/prompts/{threadId}/responses/subscribe [get]
func (p *PromptService) SubscribeToPromptResponses(
	w http.ResponseWriter,
	r *http.Request,
) {

	isAuthRsp, _, err := p.clientFactory.Client(sdk.WithTokenFromRequest(r)).
		UserSvcAPI.IsAuthorized(r.Context(), *prompt.PermissionPromptStream.Id).
		Body(openapi.UserSvcIsAuthorizedRequest{}).
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

	if vars["threadId"] == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Missing threadId path parameter`))
		return
	}
	threadId := vars["threadId"]

	w.Header().Set("Content-Type", "text/event-stream")

	subscriber := make(chan *streammanager.Chunk)
	p.streamManager.Subscribe(threadId, subscriber)
	defer p.streamManager.Unsubscribe(threadId, subscriber)

	// Use context to handle client disconnection
	ctx := r.Context()
	go func() {
		<-ctx.Done()
		p.streamManager.Unsubscribe(threadId, subscriber)
	}()

	for resp := range subscriber {
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Printf("Failed to marshal JSON: %v", err)
			continue
		}

		if _, writeErr := w.Write([]byte("data: " + string(jsonResp) + "\n")); writeErr != nil {
			log.Printf("Failed to write streaming response: %v", writeErr)
			break // Exit the loop on write errors
		}

		if flusher, ok := w.(http.Flusher); ok {
			flusher.Flush()
		} else {
			logger.Warn("Warning: ResponseWriter does not support flushing, streaming might be delayed")
		}
	}
}
