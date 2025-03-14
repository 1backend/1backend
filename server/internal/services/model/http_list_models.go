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
	"github.com/openorch/openorch/sdk/go/datastore"
	model "github.com/openorch/openorch/server/internal/services/model/types"
)

// @ID listModels
// @Summary List Models
// @Description Retrieves a list of models.
// @Description
// @Description Requires `model-svc:model:view` permission.
// @Tags Model Svc
// @Accept json
// @Produce json
// @Success 200 {object} model.ListModelsResponse
// @Failure 401 {object} model.ErrorResponse "Unauthorized"
// @Failure 500 {object} model.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /model-svc/models [post]
func (ms *ModelService) ListModels(
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

	models, err := ms.listModels()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	jsonData, _ := json.Marshal(model.ListModelsResponse{
		Models: models,
	})
	w.Write(jsonData)
}

func (ms *ModelService) listModels() ([]*model.Model, error) {
	modelIs, err := ms.modelsStore.Query().Find()
	if err != nil {
		return nil, err
	}

	models := []*model.Model{}
	for _, modelI := range modelIs {
		models = append(models, modelI.(*model.Model))
	}

	return models, nil
}

func (ms *ModelService) getModel(
	modelId string,
) (*model.Model, bool, error) {
	modelIs, err := ms.modelsStore.Query(
		datastore.Id(modelId),
	).Find()
	if err != nil {
		return nil, false, err
	}

	models := []*model.Model{}
	for _, modelI := range modelIs {
		models = append(models, modelI.(*model.Model))
	}

	if len(models) == 0 {
		return nil, false, nil
	}

	return models[0], true, nil
}
