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

// @ID serveUpload
// @Summary Serve an Uploaded File
// @Description Retrieves and serves a previously uploaded file using its File ID.
// @Description Note: The `ID` and `FileID` fields of an upload are different.
// @Description - `FileID` is a unique identifier for the file itself.
// @Description - `ID` is a unique identifier for a specific replica of the file.
// @Description Since 1Backend is a distributed system, files can be replicated across multiple nodes.
// @Description This means each uploaded file may have multiple records with the same `FileID` but different `ID`s.
// @Tags File Svc
// @Accept json
// @Produce application/octet-stream
// @Param fileId path string true "FileID uniquely identifies the file itself (not an ID, which represents a specific replica)"
// @Success 200 {file} binary "File served successfully"
// @Failure 400 {object} file.ErrorResponse "Missing Upload ID"
// @Failure 404 {object} file.ErrorResponse "File Not Found"
// @Failure 500 {object} file.ErrorResponse "Internal Server Error"
// @Router /file-svc/serve/upload/{fileId} [get]
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

	var upload *file.Upload

	if cachedUpload, ok := fs.cache.Get(fileId); ok {
		upload = cachedUpload
	} else {
		uploadReplicaIs, err := fs.uploadStore.Query(datastore.Equals(
			[]string{"fileId"},
			fileId,
		)).Find()

		if err != nil || len(uploadReplicaIs) == 0 {
			endpoint.WriteString(w, http.StatusNotFound, "File not found")
			return
		}

		uploadReplicas := toUploads(uploadReplicaIs)
		upload = uploadReplicas[0]

		fs.cache.Add(fileId, upload)
	}

	src, size, err := fs.storage.Open(r.Context(), upload.FilePath)
	if err != nil {
		logger.Error("Failed to open file stream", slog.Any("error", err))
		endpoint.WriteString(w, http.StatusNotFound, "File not found")
		return
	}
	defer src.Close()

	filename := upload.FileName

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
