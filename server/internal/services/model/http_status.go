/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
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

// @ID getModelStatus
// @Summary Get Model Status
// @Description Retrieves the status of a model by ID.
// @Description
// @Description Requires the `model-svc:model:view` permission.
// @Tags Model Svc
// @Accept json
// @Produce json
// @Param modelId path string true "Model ID" // Path parameter for the model ID
// @Success 200 {object} model.StatusResponse "Model status retrieved successfully"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /model-svc/model/{modelId}/status [get]
func (ms *ModelService) Status(
	w http.ResponseWriter,
	r *http.Request,
) {

	isAuthRsp, statusCode, err := ms.options.PermissionChecker.HasPermission(
		r,
		model.PermissionModelView,
	)
	if err != nil {
		endpoint.WriteErr(w, statusCode, err)
		return
	}
	if !isAuthRsp.Authorized {
		endpoint.Unauthorized(w)
		return
	}

	modelId := mux.Vars(r)["id"]
	unesc, err := url.PathUnescape(modelId)
	if err != nil {
		logger.Error("Error unescaping model ID", slog.String("error", err.Error()))
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid Model ID")
		return
	}

	status, err := ms.status(unesc)
	if err != nil {
		logger.Error("Error getting model status",
			slog.String("modelId", unesc),
			slog.String("error", err.Error()),
		)
		endpoint.InternalServerError(w)
		return
	}

	jsonData, _ := json.Marshal(model.StatusResponse{
		Status: status,
	})
	_, err = w.Write([]byte(jsonData))
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}
