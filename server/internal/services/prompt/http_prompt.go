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
	"net/http"

	openapi "github.com/openorch/openorch/clients/go"
	sdk "github.com/openorch/openorch/sdk/go"
	prompt "github.com/openorch/openorch/server/internal/services/prompt/types"
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

	isAuthRsp, _, err := p.clientFactory.Client(sdk.WithTokenFromRequest(r)).
		UserSvcAPI.IsAuthorized(r.Context(), *prompt.PermissionPromptCreate.Id).
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

	req := &prompt.PromptRequest{}
	err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	prsp, err := p.prompt(r.Context(), req, *isAuthRsp.User.Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bs, _ := json.Marshal(prsp)
	w.Write(bs)
}
