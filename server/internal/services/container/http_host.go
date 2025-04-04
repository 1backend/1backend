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

// @ID getHost
// @Summary      Get Container Host
// @Description  Retrieve information about the Container host
// @Tags         Container Svc
// @Accept       json
// @Produce      json
// @Success      200   {object}  container.GetHostResponse
// @Failure      401   {object}  container.ErrorResponse  "Unauthorized"
// @Failure      500   {object}  container.ErrorResponse  "Internal Server Error"
// @Security BearerAuth
// @Router       /container-svc/host [get]
func (dm *ContainerService) Host(
	w http.ResponseWriter,
	req *http.Request,
) {

	isAuthRsp, _, err := dm.clientFactory.Client(client.WithTokenFromRequest(req)).
		UserSvcAPI.IsAuthorized(req.Context(), container.PermissionContainerView).
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

	host, err := dm.backend.DaemonAddress()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	jsonData, _ := json.Marshal(container.GetHostResponse{
		Host: host,
	})
	w.Write(jsonData)
}
