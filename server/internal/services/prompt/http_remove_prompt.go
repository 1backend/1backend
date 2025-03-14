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

	req := &prompt.RemovePromptRequest{}
	err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	err = p.removePrompt(req.PromptId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(`{}`))
}
