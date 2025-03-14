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
	"github.com/openorch/openorch/sdk/go/datastore"
	prompt "github.com/openorch/openorch/server/internal/services/prompt/types"
)

// @ID listPrompts
// @Summary List Prompts
// @Description List prompts that satisfy a query.
// @Tags Prompt Svc
// @Accept json
// @Produce json
// @Param body body prompt.ListPromptsRequest false "List Prompts Request"
// @Success 200 {object} prompt.ListPromptsResponse
// @Failure 400 {object} prompt.ErrorResponse "Invalid JSON"
// @Failure 401 {object} prompt.ErrorResponse "Unauthorized"
// @Failure 500 {object} prompt.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /prompt-svc/prompts [post]
func (p *PromptService) ListPrompts(
	w http.ResponseWriter,
	r *http.Request,
) {

	isAuthRsp, _, err := p.clientFactory.Client(sdk.WithTokenFromRequest(r)).
		UserSvcAPI.IsAuthorized(r.Context(), *prompt.PermissionPromptView.Id).
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

	req := prompt.ListPromptsRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	options := &prompt.ListPromptOptions{
		Query: req.Query,
	}
	if options.Query == nil {
		options.Query = &datastore.Query{}
	}

	if !options.Query.HasFieldFilter("status") {
		options.Query.Filters = append(
			options.Query.Filters,
			datastore.IsInList(datastore.Field("status"),
				prompt.PromptStatusRunning,
				prompt.PromptStatusScheduled,
			),
		)
	}
	if options.Query.Limit == 0 {
		options.Query.Limit = 20
	}

	prompts, count, err := p.listPrompts(options)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	for i := range prompts {
		if prompts[i].UserId != *isAuthRsp.User.Id {
			// do not let users see other peoples promtps,
			// not even if they are admins
			// eg. imagine a sysadmin looking at the CEO's prompt
			prompts[i].Prompt = ""
		}
	}

	response := prompt.ListPromptsResponse{
		Prompts: prompts,
		Count:   count,
	}
	if len(prompts) >= 20 {
		response.After = prompts[len(prompts)-1].CreatedAt
	}

	bs, _ := json.Marshal(response)
	w.Write(bs)
}
