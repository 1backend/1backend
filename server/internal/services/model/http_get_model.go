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

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	model "github.com/1backend/1backend/server/internal/services/model/types"
	"github.com/gorilla/mux"
)

// @ID getModel
// @Summary Get a Model
// @Description Retrieves the details of a model by its ID.
// @Description
// @Description the Requires `model.view` permission.
// @Tags Model Svc
// @Accept json
// @Produce json
// @Param modelId path string true "Model ID"
// @Success 200 {object} model.GetModelResponse
// @Failure 400 {object} model.ErrorResponse "Missing Parameter"
// @Failure 400 {object} model.ErrorResponse "Invalid Model ID"
// @Failure 401 {object} model.ErrorResponse "Unauthorized"
// @Failure 404 {object} model.ErrorResponse "Model Not Found"
// @Failure 500 {object} model.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /model-svc/model/{modelId} [get]
func (ms *ModelService) Get(
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

	vars := mux.Vars(r)
	if vars["modelId"] == "" {
		endpoint.WriteString(w, http.StatusBadRequest, "Missing Parameter")
		return
	}

	mid, err := url.PathUnescape(vars["modelId"])
	if err != nil {
		logger.Error("Failed to unescape model ID",
			slog.String("modelId", vars["modelId"]),
			slog.Any("error", err),
		)
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid model ID")
		return
	}

	mod, found, err := ms.getModel(mid)
	if err != nil {
		logger.Error("Failed to get model",
			slog.String("modelId", mid),
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}
	if !found {
		endpoint.WriteString(w, http.StatusNotFound, "Model Not Found")
		return
	}

	platform, _, err := ms.getPlatform(mod.PlatformId)
	if err != nil {
		logger.Error("Failed to get platform",
			slog.String("platformId", mod.PlatformId),
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	jsonData, _ := json.Marshal(model.GetModelResponse{
		Exists:   found,
		Model:    mod,
		Platform: platform,
	})
	_, err = w.Write([]byte(jsonData))
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}

func (ms *ModelService) getPlatform(platformId string) (*model.Platform, bool, error) {
	platformIs, err := ms.platformsStore.Query(
		datastore.Id(platformId),
	).Find()
	if err != nil {
		return nil, false, err
	}

	platforms := []*model.Platform{}
	for _, platformI := range platformIs {
		platforms = append(platforms, platformI.(*model.Platform))
	}

	return platforms[0], true, nil
}
