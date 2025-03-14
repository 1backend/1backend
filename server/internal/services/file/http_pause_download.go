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

	"github.com/gorilla/mux"
	sdk "github.com/openorch/openorch/sdk/go"
	file "github.com/openorch/openorch/server/internal/services/file/types"
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
// @Failure 400 {string} string "Invalid JSON"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /file-svc/download/{url}/pause [put]
func (ds *FileService) PauseDownload(
	w http.ResponseWriter,
	r *http.Request,
) {

	isAuthRsp, _, err := ds.clientFactory.Client(sdk.WithTokenFromRequest(r)).
		UserSvcAPI.IsAuthorized(r.Context(), *file.PermissionDownloadEdit.Id).
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

	ur, err := url.PathUnescape(mux.Vars(r)["url"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Download ID in path is not URL encoded"))
		return
	}

	err = ds.pause(ur)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	jsonData, _ := json.Marshal(map[string]any{})
	w.Write(jsonData)
}
