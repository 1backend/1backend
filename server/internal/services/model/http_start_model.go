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

// @ID startModel
// @Summary Start a Model
// @Description Starts a model by ID
// @Tags Model Svc
// @Accept json
// @Produce json
// @Param modelId path string true "Model ID"
// @Success 200 {object} model.StartResponse
// @Failure 401 {object} model.ErrorResponse "Missing Parameter"
// @Failure 401 {object} model.ErrorResponse "Unauthorized"
// @Failure 500 {object} model.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /model-svc/model/{modelId}/start [put]
func (ms *ModelService) StartSpecific(
	w http.ResponseWriter,
	r *http.Request,
) {

	v := mux.Vars(r)
	if v["id"] == "" {
		endpoint.WriteString(w, http.StatusBadRequest, "Missing Parameter")
		return
	}

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

	req := model.StartRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logger.Error("Error decoding JSON", slog.String("error", err.Error()))
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	defer r.Body.Close()

	modelId, err := url.PathUnescape(v["modelId"])
	if err != nil {
		logger.Error("Model ID in path is not URL encoded", slog.String("error", err.Error()))
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid Model ID")
		return
	}

	err = ms.startModel(modelId)
	if err != nil {
		logger.Error("Failed to start model",
			slog.String("modelId", modelId),
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	jsonData, _ := json.Marshal(model.StartResponse{})
	w.Write(jsonData)
}
