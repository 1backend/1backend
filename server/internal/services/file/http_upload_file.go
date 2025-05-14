package fileservice

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"time"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	file "github.com/1backend/1backend/server/internal/services/file/types"
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
// @Failure 400 {object} file.ErrorResponse "invalid request"
// @Failure 401 {object} file.ErrorResponse "Unauthorized"
// @Failure 500 {object} file.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /file-svc/upload [put]
func (fs *FileService) UploadFile(
	w http.ResponseWriter,
	r *http.Request,
) {
	isAuthRsp, statusCode, err := fs.permissionChecker.HasPermission(
		r,
		file.PermissionUploadCreate,
	)
	if err != nil {
		endpoint.WriteErr(w, statusCode, err)
		return
	}
	if !isAuthRsp.GetAuthorized() {
		endpoint.WriteString(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	reader, err := r.MultipartReader()
	if err != nil {
		logger.Error("Failed to create multipart reader", slog.Any("error", err))
		endpoint.WriteString(w, http.StatusBadRequest, "invalid request")
		return
	}

	var uploadRecord file.Upload
	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			logger.Error("Failed to read multipart data", slog.Any("error", err))
			endpoint.InternalServerError(w)
			return
		}

		if part.FileName() == "" {
			continue
		}

		// File IDs should not use the sdk.Id as they must be more unique to
		// prevent enumeration, as there is no concept of file ownership.
		fileId := sdk.OpaqueId("file")

		destinationFilePath := filepath.Join(fs.uploadFolder, fileId)
		dstFile, err := os.Create(destinationFilePath)
		if err != nil {
			logger.Error("Failed to create destination file", slog.Any("error", err))
			endpoint.InternalServerError(w)
			return
		}
		defer dstFile.Close()

		written, err := io.Copy(dstFile, part)
		if err != nil {
			logger.Error("Failed to copy file data", slog.Any("error", err))
			endpoint.InternalServerError(w)
			return
		}

		if fs.nodeId == "" {
			err := fs.getNodeId(r.Context())
			if err != nil {
				logger.Error("Failed to get node ID", slog.Any("error", err))
				endpoint.InternalServerError(w)
				return
			}
		}

		// @todo this is fairly weird that we process multiple files but only a single one is returned
		uploadRecord = file.Upload{
			Id:        sdk.Id("upl"),
			FileId:    fileId,
			NodeId:    fs.nodeId,
			FileName:  part.FileName(),
			FilePath:  destinationFilePath,
			UserId:    isAuthRsp.GetUser().Id,
			FileSize:  written,
			CreatedAt: time.Now(),
		}
		err = fs.uploadStore.Upsert(uploadRecord)
		if err != nil {
			logger.Error("Failed to save upload record", slog.Any("error", err))
			endpoint.InternalServerError(w)
			return
		}
	}

	jsonData, _ := json.Marshal(file.UploadFileResponse{
		Upload: uploadRecord,
	})
	_, err = w.Write([]byte(jsonData))
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}
