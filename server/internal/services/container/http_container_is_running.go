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

// @ID containerIsRunning
// @Summary      Check If a Container Is Running
// @Description  Check if a Docker container is running, identified by hash or name.
// @Tags         Container Svc
// @Accept       json
// @Produce      json
// @Param        hash  query     string  false  "Container Hash"
// @Param        name  query     string  false  "Container Name"
// @Success      200   {object}  container.ContainerIsRunningResponse
// @Failure      400   {object}  container.ErrorResponse  "Invalid JSON or Missing Parameters"
// @Failure      401   {object}  container.ErrorResponse  "Unauthorized"
// @Failure      500   {object}  container.ErrorResponse  "Internal Server Error"
// @SecurityDefinitions.bearerAuth BearerAuth
// @Security     BearerAuth
// @Router       /container-svc/container/is-running [get]
func (dm *ContainerService) ContainerIsRunning(
	w http.ResponseWriter,
	r *http.Request,
) {

	isAuthRsp, _, err := dm.clientFactory.Client(sdk.WithTokenFromRequest(r)).
		UserSvcAPI.IsAuthorized(r.Context(), *container.PermissionContainerView.Id).
		Body(openapi.UserSvcIsAuthorizedRequest{
			GrantedSlugs: []string{"model-svc"},
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

	q := r.URL.Query()

	hash := q.Get("hash")
	name := q.Get("name")

	if hash == "" && name == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Missing Parameters`))
		return
	}

	if name != "" {
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(`Not Implemented`))
		return
	}

	isRunningRsp, err := dm.backend.ContainerIsRunning(container.ContainerIsRunningRequest{
		Hash: hash,
		Name: name,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	jsonData, _ := json.Marshal(isRunningRsp)
	w.Write(jsonData)
}
