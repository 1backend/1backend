package fileservice

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
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
	isAuthRsp, statusCode, err := fs.options.PermissionChecker.HasPermission(
		r,
		file.PermissionUploadCreate,
	)
	if err != nil {
		endpoint.WriteErr(w, statusCode, err)
		return
	}
	if !isAuthRsp.Authorized {
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

		intricatePath := calculateIntricatePath(fileId)
		now := time.Now().UTC()

		// @todo this is fairly weird that we process multiple files but only a single one is returned
		uploadRecord = file.Upload{
			Id:        sdk.Id("upl"),
			FileId:    fileId,
			FileName:  part.FileName(),
			FilePath:  intricatePath,
			UserId:    isAuthRsp.User.Id,
			CreatedAt: now,
			UpdatedAt: now,
		}

		written, err := fs.storage.Save(r.Context(), &uploadRecord, part)
		if err != nil {
			logger.Error("Failed to save file to storage", slog.Any("error", err))
			endpoint.InternalServerError(w)
			return
		}

		uploadRecord.FileSize = written

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

// Calculate the nested path: e.g., "81/d2/file_81d2..."
// This prevents directory saturation on your SSD.
//
// Input:  "file_81d259fc..."
// Output: "81/d2/file_81d259fc..."
func calculateIntricatePath(fileId string) string {
	prefix := "file_"
	idPart := fileId
	if len(fileId) > len(prefix) && fileId[:len(prefix)] == prefix {
		idPart = fileId[len(prefix):]
	}

	return shardStoragePathWithBasis(idPart, fileId)
}
