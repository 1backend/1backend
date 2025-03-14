/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package containerservice

import (
	"encoding/json"
	"net/http"

	openapi "github.com/openorch/openorch/clients/go"
	sdk "github.com/openorch/openorch/sdk/go"
	container "github.com/openorch/openorch/server/internal/services/container/types"
)

// @ID stopContainer
// @Summary Stop a Container
// @Description Stops a Docker container with the specified parameters.
// @Description
// @Description Requires the `container-svc:container:stop` permission.
// @Tags Container Svc
// @Accept json
// @Produce json
// @Param body body container.StopContainerRequest true "Stop Container Request"
// @Success 200 {object} container.StopContainerResponse
// @Failure 400 {object} container.ErrorResponse "Invalid JSON"
// @Failure 401 {object} container.ErrorResponse "Unauthorized"
// @Failure 500 {object} container.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /container-svc/container/stop [put]
func (dm *ContainerService) StopContainer(
	w http.ResponseWriter,
	r *http.Request,
) {

	isAuthRsp, _, err := dm.clientFactory.Client(sdk.WithTokenFromRequest(r)).
		UserSvcAPI.IsAuthorized(r.Context(), *container.PermissionContainerStop.Id).
		Body(openapi.UserSvcIsAuthorizedRequest{
			GrantedSlugs: []string{"model-svc", "deploy-svc"},
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

	req := &container.StopContainerRequest{}
	err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	rsp, err := dm.backend.StopContainer(*req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	jsonData, _ := json.Marshal(rsp)
	w.Write(jsonData)
}
