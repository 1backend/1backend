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

// @ID getHost
// @Summary      Get Container Host
// @Description  Retrieve information about the Container host
// @Tags         Container Svc
// @Accept       json
// @Produce      json
// @Success      200   {object}  container.GetHostResponse
// @Failure      401   {object}  container.ErrorResponse  "unauthorized"
// @Failure      500   {object}  container.ErrorResponse  "internal server error"
// @Security BearerAuth
// @Router       /container-svc/host [get]
func (dm *ContainerService) Host(
	w http.ResponseWriter,
	req *http.Request,
) {
	isAuthRsp, statusCode, err := dm.permissionChecker.HasPermission(
		req,
		container.PermissionContainerView,
	)
	if err != nil {
		endpoint.WriteErr(w, statusCode, err)
		return
	}
	if !isAuthRsp.GetAuthorized() {
		endpoint.Unauthorized(w)
		return
	}

	host, err := dm.backend.DaemonAddress()
	if err != nil {
		logger.Error(
			"Failed to get host",
			"error", err,
		)
		endpoint.InternalServerError(w)
		return
	}

	jsonData, _ := json.Marshal(container.GetHostResponse{
		Host: host,
	})
	_, err = w.Write([]byte(jsonData))
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}
