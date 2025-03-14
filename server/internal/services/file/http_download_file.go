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

	openapi "github.com/openorch/openorch/clients/go"
	sdk "github.com/openorch/openorch/sdk/go"
	file "github.com/openorch/openorch/server/internal/services/file/types"
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
// @Failure 500 {object} file.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /file-svc/download [put]
func (ds *FileService) Download(
	w http.ResponseWriter,
	r *http.Request,
) {

	isAuthRsp, isAuthHttpRsp, err := ds.clientFactory.Client(sdk.WithTokenFromRequest(r)).
		UserSvcAPI.IsAuthorized(r.Context(), *file.PermissionDownloadCreate.Id).
		Body(openapi.UserSvcIsAuthorizedRequest{
			GrantedSlugs: []string{"model-svc"},
		}).
		Execute()
	if err != nil {
		w.WriteHeader(isAuthHttpRsp.StatusCode)
		w.Write([]byte(err.Error()))
		return
	}
	if !isAuthRsp.GetAuthorized() {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`Unauthorized`))
		return
	}

	req := file.DownloadFileRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	err = ds.download(r.Context(), req.URL, req.FolderPath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	jsonData, _ := json.Marshal(map[string]any{})
	w.Write(jsonData)
}
