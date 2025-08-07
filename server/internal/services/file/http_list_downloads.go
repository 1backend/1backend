/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package fileservice

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	file "github.com/1backend/1backend/server/internal/services/file/types"
)

// @ID listDownloads
// @Summary List Downloads
// @Description List download details.
// @Description
// @Description Requires the `file-svc:download:view` permission.
// @Tags File Svc
// @Accept json
// @Produce json
// @Success 200 {object} file.DownloadsResponse "List of downloads"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /file-svc/downloads [post]
func (ds *FileService) ListDownloads(
	w http.ResponseWriter,
	r *http.Request,
) {

	isAuthRsp, statusCode, err := ds.options.PermissionChecker.HasPermission(
		r,
		file.PermissionDownloadView,
	)
	if err != nil {
		endpoint.WriteErr(w, statusCode, err)
		return
	}
	if !isAuthRsp.Authorized {
		endpoint.Unauthorized(w)
		return
	}

	details, err := ds.list()
	if err != nil {
		logger.Error(
			"Error listing downloads",
			slog.String("error", err.Error()),
		)
		endpoint.InternalServerError(w)
		return
	}

	jsonData, _ := json.Marshal(file.DownloadsResponse{
		Downloads: details,
	})
	_, err = w.Write([]byte(jsonData))
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}
