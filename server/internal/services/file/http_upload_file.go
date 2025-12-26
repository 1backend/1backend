package fileservice

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"path/filepath"
	"time"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	file "github.com/1backend/1backend/server/internal/services/file/types"
	"github.com/pkg/errors"
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

		if fs.nodeId == "" {
			err := fs.getNodeId(r.Context())
			if err != nil {
				logger.Error("Failed to get node ID", slog.Any("error", err))
				endpoint.InternalServerError(w)
				return
			}
		}

		// Calculate the nested path: e.g., "81/d2/file_81d2..."
		// This prevents directory saturation on your SSD.
		intricatePath := calculateIntricatePath(fileId)

		// @todo this is fairly weird that we process multiple files but only a single one is returned
		uploadRecord = file.Upload{
			Id:        sdk.Id("upl"),
			FileId:    fileId,
			NodeId:    fs.nodeId,
			FileName:  part.FileName(),
			FilePath:  intricatePath,
			UserId:    isAuthRsp.User.Id,
			CreatedAt: time.Now(),
		}
		err = fs.uploadStore.Upsert(uploadRecord)
		if err != nil {
			logger.Error("Failed to save upload record", slog.Any("error", err))
			endpoint.InternalServerError(w)
			return
		}

		written, err := fs.storage.Save(r.Context(), &uploadRecord, part)
		if err != nil {
			logger.Error("Failed to save file to storage", slog.Any("error", err))
			endpoint.InternalServerError(w)
			return
		}

		uploadRecord.FileSize = written
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

func (fs *FileService) getNodeId(ctx context.Context) error {
	token, err := fs.getToken()
	if err != nil {
		return errors.Wrap(err, "cannot get token")
	}

	nodeRsp, _, err := fs.options.ClientFactory.
		Client(client.WithToken(token)).
		RegistrySvcAPI.SelfNode(ctx).
		Execute()
	if err != nil {
		return err
	}

	fs.nodeId = nodeRsp.Node.Id
	return nil
}

func calculateIntricatePath(fileId string) string {
	// Input:  "file_81d259fc..."
	// Output: "81/d2/file_81d259fc..."

	prefix := "file_"
	idPart := fileId
	if len(fileId) > len(prefix) && fileId[:len(prefix)] == prefix {
		idPart = fileId[len(prefix):]
	}

	if len(idPart) < 4 {
		return fileId
	}

	return filepath.Join(idPart[0:2], idPart[2:4], fileId)
}
