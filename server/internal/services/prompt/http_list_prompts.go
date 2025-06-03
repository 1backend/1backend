/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package promptservice

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	prompt "github.com/1backend/1backend/server/internal/services/prompt/types"
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

	isAuthRsp, statusCode, err := p.options.PermissionChecker.HasPermission(
		r,
		prompt.PermissionPromptView,
	)
	if err != nil {
		endpoint.WriteErr(w, statusCode, err)
		return
	}
	if !isAuthRsp.GetAuthorized() {
		endpoint.Unauthorized(w)
		return
	}

	req := prompt.ListPromptsRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logger.Error(
			"Failed to decode request body",
			slog.String("error", err.Error()),
		)
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid JSON")
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
		logger.Error(
			"Failed to list prompts",
			slog.String("error", err.Error()),
		)
		endpoint.InternalServerError(w)
		return
	}

	for i := range prompts {
		if prompts[i].UserId != isAuthRsp.User.Id {
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
	_, err = w.Write(bs)
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}
