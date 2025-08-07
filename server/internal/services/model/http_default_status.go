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

	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	model "github.com/1backend/1backend/server/internal/services/model/types"
)

// @ID getDefaultModelStatus
// @Summary Get Default Model Status
// @Description Retrieves the status of the default model.
// @Description
// @Description Requires the `model-svc:model:view` permission.
// @Tags Model Svc
// @Accept json
// @Produce json
// @Success 200 {object} model.StatusResponse "Model status retrieved successfully"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /model-svc/default-model/status [get]
func (ms *ModelService) DefaultStatus(
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

	status, err := ms.status("")
	if err != nil {
		logger.Error("Error getting default model status", slog.String("error", err.Error()))
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
