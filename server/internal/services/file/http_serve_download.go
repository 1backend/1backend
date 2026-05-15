package fileservice

import (
	"io"
	"log/slog"
	"mime"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	file "github.com/1backend/1backend/server/internal/services/file/types"
)

// @ID serveDownload
// @Summary Serve a Downloaded file
// @Description Serves a previously downloaded file based on its URL.
// @Tags File Svc
// @Produce application/octet-stream
// @Param url path string true "URL of the file. Even after downloading, the file is still referenced by its original internet URL."
// @Success 200 {file} binary "File served successfully"
// @Failure 400 {object} file.ErrorResponse "Invalid Download URL"
// @Failure 400 {object} file.ErrorResponse "Error Parsing Download URL"
// @Failure 404 {object} file.ErrorResponse "File Not Found"
// @Failure 500 {object} file.ErrorResponse "Internal Server Error"
// @Router /file-svc/serve/download/{url} [get]
func (fs *FileService) ServeDownload(
	w http.ResponseWriter,
	r *http.Request,
) {
	vars := mux.Vars(r)
	ur, err := url.PathUnescape(vars["url"])
	if err != nil {
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid Download URL")
		return
	}
	ctx, span := startFileSpan(r.Context(), "file.serve_download")
	defer span.End()
	r = r.WithContext(ctx)

	downloadRows, err := fs.downloadStore.Query(datastore.Equals(
		[]string{"url"},
		ur,
	)).Find()
	if err != nil {
		logger.Error("Error Querying Download", slog.Any("error", err))
		span.RecordError(err)
		span.SetStatus(codes.Error, "query_download")
		recordFileDownloadServe(r.Context(), "lookup", "query_error", -1)
		endpoint.WriteString(w, http.StatusInternalServerError, err.Error())
		return
	}

	if len(downloadRows) == 0 {
		err := fs.download(r.Context(), ur, "", true)
		if err != nil {
			logger.Error("Failed to download on demand",
				slog.String("url", ur),
				slog.Any("error", err),
			)
			span.RecordError(err)
			span.SetStatus(codes.Error, "on_demand_download")
			recordFileDownloadRecovery(r.Context(), "on_demand", "error")
			recordFileDownloadServe(r.Context(), "on_demand", "error", -1)
			endpoint.WriteString(w, http.StatusInternalServerError, "Failed to download file")
			return
		}
		recordFileDownloadRecovery(r.Context(), "on_demand", "success")

		for i := 0; i < 10; i++ {
			downloadRows, err = fs.downloadStore.Query(datastore.Equals(
				[]string{"url"},
				ur,
			)).Find()
			if err != nil {
				logger.Error("Error querying download after on-demand download", slog.Any("error", err))
				span.RecordError(err)
				span.SetStatus(codes.Error, "query_after_on_demand")
				recordFileDownloadServe(r.Context(), "on_demand", "query_error", -1)
				endpoint.WriteString(w, http.StatusInternalServerError, err.Error())
				return
			}
			if len(downloadRows) > 0 {
				break
			}
			time.Sleep(200 * time.Millisecond)
		}

		if len(downloadRows) == 0 {
			span.SetStatus(codes.Error, "not_found_after_on_demand")
			recordFileDownloadServe(r.Context(), "on_demand", "not_found", -1)
			endpoint.WriteString(w, http.StatusNotFound, "File Not Found")
			return
		}
	}

	span.SetAttributes(attribute.String("file.download.route", "local"))
	fs.serveLocalDownload(toDownloads(downloadRows), w, r)
}

func (fs *FileService) serveLocalDownload(
	downloads []*file.InternalDownload,
	w http.ResponseWriter,
	r *http.Request,
) {
	ctx, span := startFileSpan(r.Context(), "file.serve_download.local")
	defer span.End()
	r = r.WithContext(ctx)

	storageSource := fileStorageSourceLocal
	download := pickDownload(downloads)
	if download == nil {
		span.SetStatus(codes.Error, "missing_record")
		recordFileDownloadServe(r.Context(), "local", "missing_record", -1)
		endpoint.WriteString(w, http.StatusNotFound, "file not found")
		return
	}

	fileInfo, err := os.Stat(download.FilePath)
	if err != nil || fileInfo.IsDir() {
		downloadURL := download.URL
		logger.Warn("Local download missing on disk, attempting recovery",
			slog.String("url", download.URL),
			slog.String("filePath", download.FilePath),
			slog.Any("error", err),
		)

		storageFilePath := DownloadStorageFilePath(download.URL)
		restored, _, restoreErr := fs.restoreDownloadFromStorage(
			r.Context(),
			storageFilePath,
			download.FilePath,
		)
		if restoreErr != nil {
			logger.Warn("Failed to restore missing local download from storage",
				slog.String("url", download.URL),
				slog.String("filePath", download.FilePath),
				slog.Any("error", restoreErr),
			)
			recordFileDownloadRecovery(r.Context(), "storage_restore", "error")
		}

		if restored {
			recordFileDownloadRecovery(r.Context(), "storage_restore", "success")
			storageSource = fileStorageSourceGCS
			download.Status = file.DownloadStatusCompleted
			if err := fs.downloadStore.Upsert(download); err != nil {
				logger.Error("Failed to update restored download record",
					slog.String("url", download.URL),
					slog.String("filePath", download.FilePath),
					slog.Any("error", err),
				)
				span.RecordError(err)
				span.SetStatus(codes.Error, "update_restored")
				recordFileDownloadServe(r.Context(), "local", "recovery_update_error", -1)
				endpoint.WriteString(w, http.StatusInternalServerError, "Failed to recover file")
				return
			}
		} else {
			recordFileDownloadRecovery(r.Context(), "storage_restore", "miss")
			if err := fs.downloadStore.Query(datastore.Equals([]string{"url"}, download.URL)).Delete(); err != nil {
				logger.Error("Failed to delete stale download record",
					slog.String("url", download.URL),
					slog.String("filePath", download.FilePath),
					slog.Any("error", err),
				)
				span.RecordError(err)
				span.SetStatus(codes.Error, "delete_stale")
				recordFileDownloadServe(r.Context(), "local", "recovery_delete_error", -1)
				endpoint.WriteString(w, http.StatusInternalServerError, "Failed to recover file")
				return
			}

			// Synchronous redownload also persists the completed file to
			// downloadStorage, which is GCS when cloud storage is enabled.
			if err := fs.download(r.Context(), download.URL, "", true); err != nil {
				logger.Error("Failed to redownload missing local download",
					slog.String("url", download.URL),
					slog.String("filePath", download.FilePath),
					slog.Any("error", err),
				)
				span.RecordError(err)
				span.SetStatus(codes.Error, "redownload_missing_local")
				recordFileDownloadRecovery(r.Context(), "local_redownload", "error")
				recordFileDownloadServe(r.Context(), "local", "recovery_download_error", -1)
				endpoint.WriteString(w, http.StatusInternalServerError, "Failed to recover file")
				return
			}
			recordFileDownloadRecovery(r.Context(), "local_redownload", "success")
		}

		downloadIs, qerr := fs.downloadStore.Query(datastore.Equals(
			[]string{"url"},
			download.URL,
		)).Find()
		if qerr != nil {
			logger.Error("Failed to query recovered download",
				slog.String("url", download.URL),
				slog.Any("error", qerr),
			)
			span.RecordError(qerr)
			span.SetStatus(codes.Error, "query_recovered")
			recordFileDownloadServe(r.Context(), "local", "recovery_query_error", -1)
			endpoint.WriteString(w, http.StatusInternalServerError, "Failed to recover file")
			return
		}
		if len(downloadIs) == 0 {
			logger.Error("Recovered download missing after re-query",
				slog.String("url", download.URL),
			)
			span.SetStatus(codes.Error, "recovered_missing")
			recordFileDownloadServe(r.Context(), "local", "recovery_missing", -1)
			endpoint.WriteString(w, http.StatusNotFound, "file not found")
			return
		}

		download = pickDownload(toDownloads(downloadIs))
		if download == nil {
			logger.Error("Recovered download missing after re-query",
				slog.String("url", downloadURL),
			)
			span.SetStatus(codes.Error, "recovered_missing")
			recordFileDownloadServe(r.Context(), "local", "recovery_missing", -1)
			endpoint.WriteString(w, http.StatusNotFound, "file not found")
			return
		}

		fileInfo, err = os.Stat(download.FilePath)
		if err != nil || fileInfo.IsDir() {
			logger.Error("Recovered download still missing on disk",
				slog.String("url", download.URL),
				slog.String("filePath", download.FilePath),
				slog.Any("error", err),
			)
			if err != nil {
				span.RecordError(err)
			}
			span.SetStatus(codes.Error, "recovered_missing_on_disk")
			recordFileDownloadServe(r.Context(), "local", "recovery_missing_disk", -1)
			endpoint.WriteString(w, http.StatusNotFound, "file not found")
			return
		}
	}

	written, err := fs.writeDownloadFileToResponse(download, fileInfo, w, storageSource)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "write_response")
		recordFileDownloadServe(r.Context(), "local", "write_error", written)
		return
	}
	recordFileDownloadServe(r.Context(), "local", "success", written)
	go fs.maybeBackfillDownloadToStorage(download)
}

func (fs *FileService) writeDownloadFileToResponse(
	download *file.InternalDownload,
	fileInfo os.FileInfo,
	w http.ResponseWriter,
	storageSource string,
) (int64, error) {
	parsedURL, err := url.Parse(download.URL)
	if err != nil {
		endpoint.WriteString(w, http.StatusBadRequest, "error parsing download URL")
		return 0, err
	}

	fileName := path.Base(parsedURL.Path)
	contentType := mime.TypeByExtension(filepath.Ext(fileName))
	if contentType == "" {
		contentType = "application/octet-stream"
	}
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Disposition", "attachment; filename=\""+sanitizeFilename(fileName)+"\"")
	setDefaultFileStorageSourceHeader(w, storageSource)
	w.Header().Set("Content-Length", strconv.FormatInt(fileInfo.Size(), 10))

	srcFile, err := os.Open(download.FilePath)
	if err != nil {
		endpoint.WriteString(w, http.StatusInternalServerError, "Failed to open file")
		return 0, err
	}
	defer srcFile.Close()

	written, err := io.Copy(w, srcFile)
	if err != nil {
		logger.Error("Failed to write file to response", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return written, err
	}
	return written, nil
}

func toDownloads(downloadIs []datastore.Row) []*file.InternalDownload {
	ret := []*file.InternalDownload{}
	for _, downloadI := range downloadIs {
		ret = append(ret, downloadI.(*file.InternalDownload))
	}

	return ret
}

func pickDownload(
	downloads []*file.InternalDownload,
) *file.InternalDownload {
	for _, download := range downloads {
		if download.Status == file.DownloadStatusCompleted {
			return download
		}
	}

	if len(downloads) == 0 {
		return nil
	}
	return downloads[0]
}
