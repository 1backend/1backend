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
	"net/http"

	openapi "github.com/openorch/openorch/clients/go"
	sdk "github.com/openorch/openorch/sdk/go"
	model "github.com/openorch/openorch/server/internal/services/model/types"
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

	isAuthRsp, _, err := ms.clientFactory.Client(sdk.WithTokenFromRequest(r)).
		UserSvcAPI.IsAuthorized(r.Context(), *model.PermissionPlatformView.Id).
		Body(openapi.UserSvcIsAuthorizedRequest{
			GrantedSlugs: []string{"prompt-svc"},
		}).
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

	platforms, err := ms.getPlatforms()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	jsonData, _ := json.Marshal(model.ListPlatformsResponse{
		Platforms: platforms,
	})
	w.Write(jsonData)
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
