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

	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	model "github.com/1backend/1backend/server/internal/services/model/types"
)

// @ID startDefaultModel
// @Summary Start the Default Model
// @Description Starts The Default Model.
// @Description
// @Description Requires the `model-svc:model:create` permission.
// @Tags Model Svc
// @Accept json
// @Produce json
// @Success 200 {object} model.StartResponse
// @Failure 401 {object} model.ErrorResponse "Unauthorized"
// @Failure 500 {object} model.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /model-svc/default-model/start [put]
func (ms *ModelService) StartDefault(
	w http.ResponseWriter,
	r *http.Request,
) {

	isAuthRsp, statusCode, err := ms.permissionChecker.HasPermission(
		r,
		model.PermissionModelCreate,
	)
	if err != nil {
		endpoint.WriteErr(w, statusCode, err)
		return
	}
	if !isAuthRsp.GetAuthorized() {
		endpoint.Unauthorized(w)
		return
	}

	err = ms.startModel("")
	if err != nil {
		logger.Error("Error starting default model", slog.String("error", err.Error()))
		endpoint.InternalServerError(w)
		return
	}

	jsonData, _ := json.Marshal(model.StartResponse{})
	_, err = w.Write([]byte(jsonData))
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}
