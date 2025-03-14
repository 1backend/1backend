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
	model "github.com/openorch/openorch/server/internal/services/model/types"
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

	modelId := mux.Vars(r)["id"]
	unesc, err := url.PathUnescape(modelId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	status, err := ms.status(unesc)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	jsonData, _ := json.Marshal(model.StatusResponse{
		Status: status,
	})
	w.Write(jsonData)
}
