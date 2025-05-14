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

// @ID prompt
// @Summary Prompt an AI
// @Description Sends a prompt and waits for a response if sync is true. If sync is false, adds the prompt to the queue and returns immediately.
// @Description
// @Description Prompts can be used for `text-to-text`, `text-to-image`, `image-to-image`, and other types of generation.
// @Description If no model ID is specified, the default model will be used (see `Model Svc` for details). The default model may or may not support the requested generation type.
// @Description
// @Description **Prompting Modes**
// @Description - **High-Level Parameters**: Uses predefined parameters relevant to `text-to-image`, `image-to-image`, etc. This mode abstracts away the underlying engine (e.g., LLaMA, Stable Diffusion) and focuses on functionality.
// @Description - **Engine-Specific Parameters**: Uses `engineParameters` to directly specify an AI engine, exposing all available parameters for fine-tuned control.
// @Description
// @Description **Permissions Required:** `prompt-svc:prompt:create`
// @Tags Prompt Svc
// @Accept json
// @Produce json
// @Param body body prompt.PromptRequest true "Add Prompt Request"
// @Success 200 {object} prompt.PromptResponse
// @Failure 400 {object} prompt.ErrorResponse "Invalid JSON"
// @Failure 401 {object} prompt.ErrorResponse "Unauthorized"
// @Failure 500 {object} prompt.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /prompt-svc/prompt [post]
func (p *PromptService) Prompt(
	w http.ResponseWriter,
	r *http.Request,
) {

	isAuthRsp, statusCode, err := p.permissionChecker.HasPermission(
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

	req := &prompt.PromptRequest{}
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

	prsp, err := p.prompt(r.Context(), req, isAuthRsp.User.Id)
	if err != nil {
		logger.Error(
			"Failed to process prompt",
			slog.String("error", err.Error()),
			slog.String("prompt", req.Prompt),
		)
		endpoint.InternalServerError(w)
		return
	}

	bs, _ := json.Marshal(prsp)
	_, err = w.Write(bs)
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}
