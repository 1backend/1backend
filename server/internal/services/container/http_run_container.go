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

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	container "github.com/1backend/1backend/server/internal/services/container/types"
)

// @ID runContainer
// @Summary Run a Container
// @Description Runs a Docker container with the specified parameters.
// @Description
// @Description Requires the `container-svc:container:run` permission.
// @Tags Container Svc
// @Accept json
// @Produce json
// @Param body body container.RunContainerRequest true "Run Container Request"
// @Success 200 {object} container.RunContainerResponse
// @Failure 400 {object} container.ErrorResponse "Invalid JSON"
// @Failure 401 {object} container.ErrorResponse "Unauthorized"
// @Failure 500 {object} container.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /container-svc/container [put]
func (dm *ContainerService) RunContainer(
	w http.ResponseWriter,
	r *http.Request,
) {

	isAuthRsp, _, err := dm.clientFactory.Client(client.WithTokenFromRequest(r)).
		UserSvcAPI.HasPermission(r.Context(), container.PermissionContainerCreate).
		Body(openapi.UserSvcHasPermissionRequest{
			PermittedSlugs: []string{
				"model-svc",
				"deploy-svc",
			},
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

	req := &container.RunContainerRequest{}
	err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	rsp, err := dm.backend.RunContainer(*req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	jsonData, _ := json.Marshal(rsp)
	w.Write(jsonData)
}
