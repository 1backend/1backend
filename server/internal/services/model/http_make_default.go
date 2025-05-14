/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package modelservice

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	model "github.com/1backend/1backend/server/internal/services/model/types"
	"github.com/gorilla/mux"
)

// @ID makeDefault
// @Summary Make a Model Default
// @Description Sets a model as the default model — when prompts are sent without a Model ID, the default model is used.
// @Tags Model Svc
// @Accept json
// @Produce json
// @Param modelId path string true "Model ID"
// @Success 200 {object} model.MakeDefaultResponse
// @Failure 400 {object} model.ErrorResponse "Invalid Model ID"
// @Failure 401 {object} model.ErrorResponse "Unauthorized"
// @Failure 500 {object} model.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /model-svc/model/{modelId}/make-default [put]
func (ms *ModelService) MakeDefault(
	w http.ResponseWriter,
	r *http.Request,
) {

	isAuthRsp, statusCode, err := ms.permissionChecker.HasPermission(
		r,
		model.PermissionModelEdit,
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
	modelId, err := url.PathUnescape(vars["modelId"])
	if err != nil {
		logger.Error("Model ID in path is not URL encoded", slog.String("error", err.Error()))
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid Model ID")
		return
	}

	err = ms.makeDefault(r.Context(), modelId)
	if err != nil {
		logger.Error("Failed to make model default",
			slog.String("modelId", modelId),
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	jsonData, _ := json.Marshal(model.MakeDefaultResponse{})
	_, err = w.Write([]byte(jsonData))
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}
