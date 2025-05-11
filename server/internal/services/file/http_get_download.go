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
	"net/http"
	"net/url"

	"github.com/1backend/1backend/sdk/go/endpoint"
	file "github.com/1backend/1backend/server/internal/services/file/types"
	"github.com/gorilla/mux"
)

// @ID getDownload
// @Summary Get a Download
// @Description Get a download by URL.
// @Description
// @Description Requires the `file-svc:download:view` permission.
// @Tags File Svc
// @Accept json
// @Produce json
// @Param url path string true "url"
// @Success 200 {object} file.GetDownloadResponse
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /file-svc/download/{url} [get]
func (fs *FileService) GetDownload(
	w http.ResponseWriter,
	r *http.Request,
) {

	isAuthRsp, statusCode, err := fs.permissionChecker.HasPermission(
		r,
		file.PermissionDownloadView,
	)
	if err != nil {
		endpoint.WriteErr(w, statusCode, err)
		return
	}
	if !isAuthRsp.GetAuthorized() {
		endpoint.Unauthorized(w)
		return
	}

	vars := mux.Vars(r)
	ur, err := url.PathUnescape(vars["url"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	dl, exists := fs.getDownload(ur)

	jsonData, _ := json.Marshal(file.GetDownloadResponse{
		Exists:   exists,
		Download: downloadToDownloadDetails(dl),
	})
	w.Write(jsonData)
}
