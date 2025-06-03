/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
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

// @ID containerIsRunning
// @Summary      Check If a Container Is Running
// @Description  Check if a Docker container is running, identified by hash or name.
// @Tags         Container Svc
// @Accept       json
// @Produce      json
// @Param        hash  query     string  false  "Container Hash"
// @Param        name  query     string  false  "Container Name"
// @Success      200   {object}  container.ContainerIsRunningResponse
// @Failure      400   {object}  container.ErrorResponse  "Invalid JSON"
// @Failure      400   {object}  container.ErrorResponse  "missing parameters"
// @Failure      401   {object}  container.ErrorResponse  "Unauthorized"
// @Failure      500   {object}  container.ErrorResponse  "Internal Server Error"
// @SecurityDefinitions.bearerAuth BearerAuth
// @Security     BearerAuth
// @Router       /container-svc/container/is-running [get]
func (dm *ContainerService) ContainerIsRunning(
	w http.ResponseWriter,
	r *http.Request,
) {
	isAuthRsp, statusCode, err := dm.options.PermissionChecker.HasPermission(
		r,
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

	q := r.URL.Query()

	hash := q.Get("hash")
	name := q.Get("name")

	if hash == "" && name == "" {
		endpoint.WriteString(w, http.StatusBadRequest, "missing parameters")
		return
	}

	if name != "" {
		endpoint.WriteString(w, http.StatusNotImplemented, "not implemented")
		return
	}

	isRunningRsp, err := dm.backend.ContainerIsRunning(container.ContainerIsRunningRequest{
		Hash: hash,
		Name: name,
	})

	if err != nil {
		logger.Error("Failed to check if container is running",
			slog.String("hash", hash),
			slog.String("name", name),
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	jsonData, _ := json.Marshal(isRunningRsp)
	_, err = w.Write([]byte(jsonData))
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}
