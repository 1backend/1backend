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

	sdk "github.com/openorch/openorch/sdk/go"
	model "github.com/openorch/openorch/server/internal/services/model/types"
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
// @Failure 400 {object} model.ErrorResponse "Invalid JSON"
// @Failure 401 {object} model.ErrorResponse "Unauthorized"
// @Failure 500 {object} model.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /model-svc/default-model/start [put]
func (ms *ModelService) StartDefault(
	w http.ResponseWriter,
	r *http.Request,
) {

	isAuthRsp, _, err := ms.clientFactory.Client(sdk.WithTokenFromRequest(r)).
		UserSvcAPI.IsAuthorized(r.Context(), *model.PermissionModelCreate.Id).
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

	err = ms.start("")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	jsonData, _ := json.Marshal(model.StartResponse{})
	w.Write(jsonData)
}
