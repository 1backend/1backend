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
	"strconv"

	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	container "github.com/1backend/1backend/server/internal/services/container/types"
)

// @ID containerSummary
// @Summary      Get Container Summary
// @Description  Get a summary of the Docker container identified by hash or name, limited to a specified number of lines.
// @Tags         Container Svc
// @Accept       json
// @Produce      json
// @Param        hash           query    string  false  "Container Hash"
// @Param        name           query    string  false  "Container Name"
// @Param        lines          query    int     false  "Number of Lines"
// @Success      200            {object} container.GetContainerSummaryResponse
// @Failure      400            {object} container.ErrorResponse  "missing parameters"
// @Failure      401            {object} container.ErrorResponse  "unauthorized"
// @Failure      500            {object} container.ErrorResponse  "internal server error"
// @Security     BearerAuth
// @Router       /container-svc/container/summary [get]
func (dm *ContainerService) Summary(
	w http.ResponseWriter,
	r *http.Request,
) {
	isAuthRsp, statusCode, err := dm.permissionChecker.HasPermission(
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
	numberOfLines := q.Get("lines")

	name := q.Get("name")
	if name == "" {
		endpoint.WriteString(w, http.StatusNotImplemented, "not implemented")
		return
	}

	lines, err := strconv.ParseInt(numberOfLines, 10, 64)
	if err != nil {
		logger.Error("Failed to parse lines",
			slog.String("numberOfLines", numberOfLines),
			slog.String("error", err.Error()),
		)
		endpoint.InternalServerError(w)
		return
	}

	summary, err := dm.backend.GetContainerSummary(container.GetContainerSummaryRequest{
		Hash:  hash,
		Lines: int(lines),
	})
	if err != nil {
		logger.Error("Failed to get container summary",
			slog.String("hash", hash),
			slog.String("name", name),
			slog.Int("lines", int(lines)),
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	jsonData, _ := json.Marshal(summary)
	_, err = w.Write([]byte(jsonData))
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}
