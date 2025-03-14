package fileservice

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	openapi "github.com/openorch/openorch/clients/go"
	sdk "github.com/openorch/openorch/sdk/go"
	file "github.com/openorch/openorch/server/internal/services/file/types"
)

// @ID uploadFile
// @Summary Upload a File
// @Description Uploads a file to the server.
// @Description Currently if using the clients only one file can be uploaded at a time due to this bug https://github.com/OpenAPITools/openapi-generator/issues/11341
// @Description Once that is fixed we should have an `PUT /file-svc/uploads`/uploadFiles (note the plural) endpoints.
// @Description In reality the endpoint "unofficially" supports multiple files. YMMV.
// @Description
// @Description Requires the `file-svc:upload:create` permission.
// @Tags File Svc
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "File to upload"
// @Success 200 {object} file.UploadFileResponse "File uploaded successfully"
// @Failure 400 {object} file.ErrorResponse "Invalid request"
// @Failure 401 {object} file.ErrorResponse "Unauthorized"
// @Failure 500 {object} file.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /file-svc/upload [put]
func (fs *FileService) UploadFile(
	w http.ResponseWriter,
	r *http.Request,
) {
	isAuthRsp, isAuthHttpRsp, err := fs.clientFactory.Client(sdk.WithTokenFromRequest(r)).
		UserSvcAPI.IsAuthorized(r.Context(), *file.PermissionUploadCreate.Id).
		Body(openapi.UserSvcIsAuthorizedRequest{
			GrantedSlugs: []string{"prompt-svc"},
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

	handleError := func(err error, statusCode int, message string) {
		w.WriteHeader(statusCode)
		w.Write([]byte(message + ": " + err.Error()))
	}

	reader, err := r.MultipartReader()
	if err != nil {
		handleError(err, http.StatusBadRequest, "Invalid request")
		return
	}

	var uploadRecord file.Upload
	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			handleError(err, http.StatusInternalServerError, "Failed to read multipart data")
			return
		}

		if part.FileName() == "" {
			continue
		}

		cleanFilename := sanitizeFilename(part.FileName())
		destinationFilePath := filepath.Join(fs.uploadFolder, cleanFilename)
		dstFile, err := os.Create(destinationFilePath)
		if err != nil {
			handleError(err, http.StatusInternalServerError, "Failed to create destination file")
			return
		}
		defer dstFile.Close()

		written, err := io.Copy(dstFile, part)
		if err != nil {
			handleError(err, http.StatusInternalServerError, "Failed to save file")
			return
		}

		if fs.nodeId == "" {
			err := fs.getNodeId(r.Context())
			if err != nil {
				handleError(err, http.StatusInternalServerError, "Failed to get node ID")
				return
			}
		}

		// @todo this is fairly weird that we process multiple files but only a single one is returned
		uploadRecord = file.Upload{
			Id:        sdk.Id("upl"),
			FileId:    sdk.Id("file"),
			NodeId:    fs.nodeId,
			FileName:  part.FileName(),
			FilePath:  destinationFilePath,
			UserId:    *isAuthRsp.GetUser().Id,
			FileSize:  written,
			CreatedAt: time.Now(),
		}
		err = fs.uploadStore.Upsert(uploadRecord)
		if err != nil {
			handleError(err, http.StatusInternalServerError, "Failed to save upload record")
			return
		}
	}

	jsonData, _ := json.Marshal(file.UploadFileResponse{
		Upload: uploadRecord,
	})
	w.Write(jsonData)
}
