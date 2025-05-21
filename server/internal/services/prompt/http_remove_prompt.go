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

	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	prompt "github.com/1backend/1backend/server/internal/services/prompt/types"
)

// @ID removePrompt
// @Summary Remove Prompt
// @Description Remove a prompt by ID.
// @Tags Prompt Svc
// @Accept json
// @Produce json
// @Param body body prompt.RemovePromptRequest true "Remove Prompt Request"
// @Success 200 {object} prompt.RemovePromptResponse "{}"
// @Failure 400 {object} prompt.ErrorResponse "Invalid JSON"
// @Failure 401 {object} prompt.ErrorResponse "Unauthorized"
// @Failure 500 {object} prompt.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /prompt-svc/remove [post]
func (p *PromptService) RemovePrompt(
	w http.ResponseWriter,
	r *http.Request,
) {

	isAuthRsp, statusCode, err := p.options.PermissionChecker.HasPermission(
		r,
		prompt.PermissionPromptCreate,
	)
	if err != nil {
		endpoint.WriteErr(w, statusCode, err)
		return
	}
	if !isAuthRsp.GetAuthorized() {
		endpoint.Unauthorized(w)
		return
	}

	req := &prompt.RemovePromptRequest{}
	err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		logger.Error(
			"Failed to decode request body",
			slog.String("error", err.Error()),
		)
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	defer r.Body.Close()

	err = p.removePrompt(req.PromptId)
	if err != nil {
		logger.Error(
			"Failed to remove prompt",
			slog.String("error", err.Error()),
			slog.String("promptId", req.PromptId),
		)
		endpoint.InternalServerError(w)
		return
	}

	_, err = w.Write([]byte(`{}`))
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}
