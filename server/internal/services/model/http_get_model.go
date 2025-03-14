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
	"net/url"

	"github.com/gorilla/mux"
	openapi "github.com/openorch/openorch/clients/go"
	sdk "github.com/openorch/openorch/sdk/go"
	"github.com/openorch/openorch/sdk/go/datastore"
	model "github.com/openorch/openorch/server/internal/services/model/types"
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
// @Failure 401 {object} model.ErrorResponse "Unauthorized"
// @Failure 500 {object} model.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /model-svc/model/{modelId} [get]
func (ms *ModelService) Get(
	w http.ResponseWriter,
	r *http.Request,
) {

	isAuthRsp, _, err := ms.clientFactory.Client(sdk.WithTokenFromRequest(r)).
		UserSvcAPI.IsAuthorized(r.Context(), *model.PermissionModelView.Id).
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

	vars := mux.Vars(r)
	if vars["modelId"] == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("no model id"))
		return
	}

	mid, err := url.PathUnescape(vars["modelId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	mod, found, err := ms.getModel(mid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if !found {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("model not found"))
		return
	}

	platform, _, err := ms.getPlatform(mod.PlatformId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	jsonData, _ := json.Marshal(model.GetModelResponse{
		Exists:   found,
		Model:    mod,
		Platform: platform,
	})
	w.Write(jsonData)
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
