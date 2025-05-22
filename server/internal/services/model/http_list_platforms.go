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

// @ID listPlatforms
// @Summary List Platforms
// @Description Retrieves a list of AI platforms. Eg. LlamaCpp, StableDiffusion etc.
// @Description
// @Description Requires `model-svc:platform:view` permission.
// @Tags Model Svc
// @Accept json
// @Produce json
// @Success 200 {object} model.ListPlatformsResponse
// @Failure 401 {object} model.ErrorResponse "Unauthorized"
// @Failure 500 {object} model.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /model-svc/platforms [post]
func (ms *ModelService) ListPlatforms(
	w http.ResponseWriter,
	r *http.Request,
) {

	isAuthRsp, statusCode, err := ms.options.PermissionChecker.HasPermission(
		r,
		model.PermissionPlatformView,
	)
	if err != nil {
		endpoint.WriteErr(w, statusCode, err)
		return
	}
	if !isAuthRsp.GetAuthorized() {
		endpoint.Unauthorized(w)
		return
	}

	platforms, err := ms.getPlatforms()
	if err != nil {
		logger.Error("Error getting platforms", slog.String("error", err.Error()))
		endpoint.InternalServerError(w)
		return
	}

	jsonData, _ := json.Marshal(model.ListPlatformsResponse{
		Platforms: platforms,
	})
	_, err = w.Write([]byte(jsonData))
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}

func (ms *ModelService) getPlatforms() ([]*model.Platform, error) {
	platformIs, err := ms.platformsStore.Query().Find()
	if err != nil {
		return nil, err
	}

	platforms := []*model.Platform{}
	for _, platformI := range platformIs {
		platforms = append(platforms, platformI.(*model.Platform))
	}

	return platforms, nil
}
