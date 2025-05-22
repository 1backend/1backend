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
	"log/slog"
	"net/http"

	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	container "github.com/1backend/1backend/server/internal/services/container/types"
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

	isAuthRsp, statusCode, err := dm.options.PermissionChecker.HasPermission(
		r,
		container.PermissionContainerStop,
	)
	if err != nil {
		endpoint.WriteErr(w, statusCode, err)
		return
	}
	if !isAuthRsp.GetAuthorized() {
		endpoint.Unauthorized(w)
		return
	}

	req := &container.StopContainerRequest{}
	err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		logger.Error("Failed to decode request",
			"error", err,
		)
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	defer r.Body.Close()

	rsp, err := dm.backend.StopContainer(*req)
	if err != nil {
		logger.Error("Failed to stop container",
			"error", err,
		)
		endpoint.InternalServerError(w)
		return
	}

	jsonData, _ := json.Marshal(rsp)
	_, err = w.Write([]byte(jsonData))
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}
