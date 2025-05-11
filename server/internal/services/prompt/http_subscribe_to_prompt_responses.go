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
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"

	streammanager "github.com/1backend/1backend/server/internal/services/prompt/stream"
	prompt "github.com/1backend/1backend/server/internal/services/prompt/types"
)

// @ID subscribeToPromptResponses
// @Summary Subscribe to Prompt Responses by Thread
// @Description Subscribe to prompt responses by thread via Server-Sent Events (SSE).
// @Description You can subscribe to threads before they are created.
// @Description The streamed strings are of type `StreamChunk`, see the PromptTypes endpoint for more details.
// @Tags Prompt Svc
// @Param threadId path string true "Thread ID"
// @Success 200 {string} string "Streaming response"
// @Failure 400 {object} prompt.ErrorResponse "Missing Parameter"
// @Failure 401 {object} prompt.ErrorResponse "Unauthorized"
// @Security BearerAuth
// @Router /prompt-svc/prompts/{threadId}/responses/subscribe [get]
func (p *PromptService) SubscribeToPromptResponses(
	w http.ResponseWriter,
	r *http.Request,
) {

	isAuthRsp, statusCode, err := p.permissionChecker.HasPermission(
		r,
		prompt.PermissionPromptStream,
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

	if vars["threadId"] == "" {
		endpoint.WriteString(w, http.StatusBadRequest, "Missing Parameter")
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
			logger.Error("Failed to marshal JSON",
				slog.Any("error", err),
			)
			continue
		}

		if _, writeErr := w.Write([]byte("data: " + string(jsonResp) + "\n")); writeErr != nil {
			logger.Error("Failed to write streaming response",
				slog.Any("error", writeErr),
			)
			break // Exit the loop on write errors
		}

		if flusher, ok := w.(http.Flusher); ok {
			flusher.Flush()
		} else {
			logger.Warn("Warning: ResponseWriter does not support flushing, streaming might be delayed")
		}
	}
}
