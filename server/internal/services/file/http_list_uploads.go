package fileservice

import (
	"encoding/json"
	"net/http"
	"path/filepath"
	"strings"

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/datastore"
	file "github.com/1backend/1backend/server/internal/services/file/types"
)

// @ID listUploads
// @Summary List Uploads
// @Description Lists uploaded files, returning only metadata about each upload.
// @Description To retrieve file content, use the `Serve an Uploaded File` endpoint, which serves a single file per request.
// @Description Note: Retrieving the contents of multiple files in a single request is not supported currently.
// @Description
// @Description Requires the `file-svc:upload:view` permission.
// @Tags File Svc
// @Accept json
// @Produce json
// @Param body body file.ListUploadsRequest false "List Uploads Request"
// @Success 200 {object} file.ListUploadsResponse "List of uploads"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /file-svc/uploads [post]
func (fs *FileService) ListUploads(
	w http.ResponseWriter,
	r *http.Request,
) {
	isAuthRsp, _, err := fs.clientFactory.Client(client.WithTokenFromRequest(r)).
		UserSvcAPI.IsAuthorized(r.Context(), *file.PermissionUploadCreate.Id).
		Body(openapi.UserSvcIsAuthorizedRequest{
			GrantedSlugs: []string{},
		}).
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

	req := &file.ListUploadsRequest{}
	err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	filters := []datastore.Filter{}
	if req.UserId != "" {
		filters = append(filters, datastore.Equals([]string{"userId"}, req.UserId))
	}
	uploadIs, err := fs.uploadStore.Query(filters...).Find()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`Cannot query uploads`))
		return
	}

	rsp := file.ListUploadsResponse{}
	for _, uploadI := range uploadIs {
		upload := uploadI.(*file.Upload)
		rsp.Uploads = append(rsp.Uploads, *upload)
	}

	jsonData, _ := json.Marshal(rsp)
	w.Write(jsonData)
}

func sanitizeFilename(name string) string {
	name = filepath.Clean(name)
	name = filepath.Base(name)
	return strings.ReplaceAll(name, "..", "_")
}
