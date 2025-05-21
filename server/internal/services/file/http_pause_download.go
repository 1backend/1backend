/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package fileservice

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	file "github.com/1backend/1backend/server/internal/services/file/types"
	"github.com/gorilla/mux"
)

// @ID pauseDownload
// @Summary Pause a Download
// @Description Pause a download that is currently in progress.
// @Description
// @Description Requires the `file-svc:download:edit` permission.
// @Tags File Svc
// @Accept json
// @Produce json
// @Param url path string true "Download URL"
// @Success 200 {object} map[string]any "Success response"
// @Failure 400 {string} string "Download ID in path is not URL encoded"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /file-svc/download/{url}/pause [put]
func (ds *FileService) PauseDownload(
	w http.ResponseWriter,
	r *http.Request,
) {
	isAuthRsp, statusCode, err := ds.options.PermissionChecker.HasPermission(
		r,
		file.PermissionDownloadEdit,
	)
	if err != nil {
		endpoint.WriteErr(w, statusCode, err)
		return
	}
	if !isAuthRsp.GetAuthorized() {
		endpoint.WriteString(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	ur, err := url.PathUnescape(mux.Vars(r)["url"])
	if err != nil {
		logger.Error("Failed to unescape URL", slog.Any("error", err))
		endpoint.WriteString(w, http.StatusBadRequest, "Download ID in path is not URL encoded")
		return
	}

	err = ds.pause(ur)
	if err != nil {
		logger.Error("Failed to pause download", slog.Any("error", err))
		endpoint.WriteString(w, http.StatusInternalServerError, "Failed to pause download")
		return
	}

	jsonData, _ := json.Marshal(map[string]any{})
	_, err = w.Write([]byte(jsonData))
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}
