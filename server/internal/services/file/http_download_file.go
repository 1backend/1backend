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

// @ID downloadFile
// @Summary Download a File
// @Description Start or resume the download for a specified URL.
// @Description
// @Description Requires the `file-svc:download:create` permission.
// @Tags File Svc
// @Accept json
// @Produce json
// @Param body body file.DownloadFileRequest true "Download Request"
// @Success 200 {object} map[string]any "Download initiated successfully"
// @Failure 400 {object} file.ErrorResponse "Invalid JSON"
// @Failure 401 {object} file.ErrorResponse "Unauthorized"
// @Failure 500 {object} file.ErrorResponse "Failed to download file"
// @Security BearerAuth
// @Router /file-svc/download [put]
func (ds *FileService) Download(
	w http.ResponseWriter,
	r *http.Request,
) {
	isAuthRsp, statusCode, err := ds.options.PermissionChecker.HasPermission(
		r,
		file.PermissionDownloadCreate,
	)
	if err != nil {
		endpoint.WriteErr(w, statusCode, err)
		return
	}
	if !isAuthRsp.GetAuthorized() {
		endpoint.WriteString(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	req := file.DownloadFileRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	defer r.Body.Close()

	err = ds.download(r.Context(), req.URL, req.FolderPath)
	if err != nil {
		endpoint.WriteString(w, http.StatusInternalServerError, "Failed to download file")
		return
	}

	jsonData, _ := json.Marshal(map[string]any{})
	_, err = w.Write([]byte(jsonData))
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}
