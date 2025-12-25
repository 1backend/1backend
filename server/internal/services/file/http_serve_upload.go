package fileservice

import (
	"fmt"
	"io"
	"log/slog"
	"mime"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	file "github.com/1backend/1backend/server/internal/services/file/types"
	"github.com/gorilla/mux"
)

// ServeUpload now handles the request by delegating to the abstract storage provider.
func (fs *FileService) ServeUpload(
	w http.ResponseWriter,
	r *http.Request,
) {
	vars := mux.Vars(r)
	fileId := vars["fileId"]
	if fileId == "" {
		endpoint.WriteString(w, http.StatusBadRequest, "missing upload ID")
		return
	}

	uploadReplicaIs, err := fs.uploadStore.Query(datastore.Equals(
		[]string{"fileId"},
		fileId,
	)).Find()
	if err != nil {
		logger.Error("Error querying upload", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}
	if len(uploadReplicaIs) == 0 {
		endpoint.WriteString(w, http.StatusNotFound, "File not found")
		return
	}

	uploadReplicas := toUploads(uploadReplicaIs)

	src, size, err := fs.storage.Open(r.Context(), fileId)
	if err != nil {
		logger.Error("Failed to open file stream", slog.Any("error", err))
		endpoint.WriteString(w, http.StatusNotFound, "File not found")
		return
	}
	defer src.Close()

	filename := uploadReplicas[0].FileName

	// 4. Set Headers and stream the response
	contentType := mime.TypeByExtension(filepath.Ext(filename))
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", sanitizeFilename(filename)))
	if size > 0 {
		w.Header().Set("Content-Length", strconv.FormatInt(size, 10))
	}

	_, err = io.Copy(w, src)
	if err != nil {
		logger.Error("Failed to write file to response", slog.Any("error", err))
		return
	}
}

// Internal helpers for the handler
func toUploads(uploadIs []datastore.Row) []*file.Upload {
	ret := []*file.Upload{}
	for _, uploadI := range uploadIs {
		ret = append(ret, uploadI.(*file.Upload))
	}
	return ret
}
