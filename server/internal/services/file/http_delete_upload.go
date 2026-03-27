package fileservice

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	file "github.com/1backend/1backend/server/internal/services/file/types"
	"github.com/gorilla/mux"
)

// @ID deleteUpload
// @Summary Delete an Uploaded File
// @Description Deletes an uploaded file and its metadata by `fileId`.
// @Description
// @Description Requires the `file-svc:upload:delete` permission.
// @Tags File Svc
// @Accept json
// @Produce json
// @Param fileId path string true "File ID"
// @Success 200 {object} map[string]any "File deleted successfully"
// @Failure 400 {object} file.ErrorResponse "invalid request"
// @Failure 401 {object} file.ErrorResponse "Unauthorized"
// @Failure 404 {object} file.ErrorResponse "File not found"
// @Failure 500 {object} file.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /file-svc/upload/{fileId} [delete]
func (fs *FileService) DeleteUpload(
	w http.ResponseWriter,
	r *http.Request,
) {
	isAuthRsp, statusCode, err := fs.options.PermissionChecker.HasPermission(
		r,
		file.PermissionUploadDelete,
	)
	if err != nil {
		endpoint.WriteErr(w, statusCode, err)
		return
	}
	if !isAuthRsp.Authorized {
		endpoint.WriteString(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	vars := mux.Vars(r)
	fileId := vars["fileId"]
	if fileId == "" {
		endpoint.WriteString(w, http.StatusBadRequest, "missing file ID")
		return
	}

	uploadIs, err := fs.uploadStore.Query(
		datastore.Equals([]string{"fileId"}, fileId),
	).Find()
	if err != nil {
		logger.Error("Failed to query uploads for deletion", slog.String("fileId", fileId), slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}
	if len(uploadIs) == 0 {
		endpoint.WriteString(w, http.StatusNotFound, "file not found")
		return
	}

	uploads := toUploads(uploadIs)
	deletedPaths := map[string]bool{}
	for _, upload := range uploads {
		if deletedPaths[upload.FilePath] {
			continue
		}
		if err := fs.storage.Delete(r.Context(), upload.FilePath); err != nil {
			logger.Error("Failed to delete uploaded file", slog.String("fileId", fileId), slog.String("filePath", upload.FilePath), slog.Any("error", err))
			endpoint.InternalServerError(w)
			return
		}
		deletedPaths[upload.FilePath] = true
	}

	err = fs.uploadStore.Query(
		datastore.Equals([]string{"fileId"}, fileId),
	).Delete()
	if err != nil {
		logger.Error("Failed to delete upload metadata", slog.String("fileId", fileId), slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}

	fs.cache.Remove(fileId)

	jsonData, _ := json.Marshal(map[string]any{})
	_, err = w.Write([]byte(jsonData))
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}
