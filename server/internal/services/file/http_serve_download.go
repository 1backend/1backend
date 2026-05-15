package fileservice

import (
	"context"
	"fmt"
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
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"

	openapi "github.com/1backend/1backend/clients/go"
	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/client"
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

	downloadReplicaIs, err := fs.downloadStore.Query(datastore.Equals(
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

	if len(downloadReplicaIs) == 0 {
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
			downloadReplicaIs, err = fs.downloadStore.Query(datastore.Equals(
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
			if len(downloadReplicaIs) > 0 {
				break
			}
			time.Sleep(200 * time.Millisecond)
		}

		if len(downloadReplicaIs) == 0 {
			span.SetStatus(codes.Error, "not_found_after_on_demand")
			recordFileDownloadServe(r.Context(), "on_demand", "not_found", -1)
			endpoint.WriteString(w, http.StatusNotFound, "File Not Found")
			return
		}
	}

	downloadReplicas := toDownloads(downloadReplicaIs)
	isLocal, err := fs.isLocalDownload(r.Context(), downloadReplicas)
	if err != nil {
		logger.Error("Error checking if download is local",
			slog.Any("error", err),
		)
		span.RecordError(err)
		span.SetStatus(codes.Error, "locality_check")
		recordFileDownloadServe(r.Context(), "lookup", "locality_error", -1)
		endpoint.InternalServerError(w)
		return
	}
	if isLocal {
		span.SetAttributes(attribute.String("file.download.route", "local"))
		fs.serveLocalDownload(downloadReplicas, w, r)
	} else {
		span.SetAttributes(attribute.String("file.download.route", "remote"))
		fs.serveRemoteDownload(downloadReplicas, w, r)
	}
}

func (fs *FileService) serveLocalDownload(
	downloadReplicas []*file.InternalDownload,
	w http.ResponseWriter,
	r *http.Request,
) {
	ctx, span := startFileSpan(r.Context(), "file.serve_download.local")
	defer span.End()
	r = r.WithContext(ctx)

	storageSource := fileStorageSourceLocal
	download, err := fs.pickLocalDownload(r.Context(), downloadReplicas)
	if err != nil {
		logger.Error("Failed to pick local download", slog.Any("error", err))
		span.RecordError(err)
		span.SetStatus(codes.Error, "pick_local")
		recordFileDownloadServe(r.Context(), "local", "pick_error", -1)
		endpoint.WriteString(w, http.StatusInternalServerError, "Failed to pick local download")
		return
	}

	fileInfo, err := os.Stat(download.FilePath)
	if err != nil || fileInfo.IsDir() {
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
			if err := fs.downloadStore.Query(datastore.Id(download.Id)).Delete(); err != nil {
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

		download = downloadIs[0].(*file.InternalDownload)

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

func (fs *FileService) serveRemoteDownload(
	downloadReplicas []*file.InternalDownload,
	w http.ResponseWriter,
	r *http.Request,
) {
	ctx, span := startFileSpan(r.Context(), "file.serve_download.remote")
	defer span.End()
	r = r.WithContext(ctx)

	if download, ok, err := fs.promoteDownloadFromStorage(r.Context(), downloadReplicas); err != nil {
		logger.Warn("Failed to promote remote download from storage",
			slog.Any("error", err),
		)
		recordFileDownloadRecovery(r.Context(), "remote_storage_promote", "error")
	} else if ok {
		recordFileDownloadRecovery(r.Context(), "remote_storage_promote", "success")
		setFileStorageSourceHeader(w, fileStorageSourceGCS)
		fs.serveLocalDownload([]*file.InternalDownload{download}, w, r)
		return
	} else {
		recordFileDownloadRecovery(r.Context(), "remote_storage_promote", "miss")
	}

	downloads, err := fs.pickRemoteDownloads(r.Context(), downloadReplicas)
	if err != nil {
		logger.Error("Failed to pick remote download", slog.Any("error", err))
		span.RecordError(err)
		recordFileDownloadRecovery(r.Context(), "remote_pick", "error")
		fs.recoverAndServeRemoteDownloadLocally(downloadReplicas, w, r)
		return
	}
	if len(downloads) == 0 {
		recordFileDownloadRecovery(r.Context(), "remote_pick", "empty")
		fs.recoverAndServeRemoteDownloadLocally(downloadReplicas, w, r)
		return
	}
	recordFileDownloadRecovery(r.Context(), "remote_pick", "success")

	nodeIds := []string{}
	for _, download := range downloads {
		nodeIds = append(nodeIds, download.NodeId)
	}

	token, err := fs.getToken()
	if err != nil {
		logger.Error("Failed to get token", slog.Any("error", err))
		span.RecordError(err)
		recordFileDownloadRecovery(r.Context(), "remote_token", "error")
		fs.recoverAndServeRemoteDownloadLocally(downloads, w, r)
		return
	}
	recordFileDownloadRecovery(r.Context(), "remote_token", "success")

	nodesRsp, _, err := fs.options.ClientFactory.
		Client(client.WithToken(token)).
		RegistrySvcAPI.ListNodes(r.Context()).
		Body(
			openapi.RegistrySvcListNodesRequest{
				Ids: nodeIds,
			},
		).Execute()
	if err != nil {
		logger.Error("Failed to list nodes", slog.Any("error", err))
		span.RecordError(err)
		recordFileDownloadRecovery(r.Context(), "remote_list_nodes", "error")
		fs.recoverAndServeRemoteDownloadLocally(downloads, w, r)
		return
	}
	nodes := nodesRsp.Nodes

	if len(nodes) == 0 {
		recordFileDownloadRecovery(r.Context(), "remote_list_nodes", "empty")
		fs.recoverAndServeRemoteDownloadLocally(downloads, w, r)
		return
	}
	recordFileDownloadRecovery(r.Context(), "remote_list_nodes", "success")

	node := nodes[0]

	// todo it would be probably better to stream this ourselves here but for now it will do
	file, fileHttpRsp, err := fs.options.ClientFactory.
		Client(client.WithAddress(node.Url), client.WithToken(token)).
		FileSvcAPI.
		ServeDownload(r.Context(), downloads[0].URL).
		Execute()
	if err != nil {
		logger.Error("Failed to serve download", slog.Any("error", err))
		span.RecordError(err)
		recordFileDownloadRecovery(r.Context(), "remote_proxy", "error")
		fs.recoverAndServeRemoteDownloadLocally(downloads, w, r)
		return
	}
	recordFileDownloadRecovery(r.Context(), "remote_proxy", "success")

	defer file.Close()

	w.Header().Set("Content-Type", fileHttpRsp.Header.Get("Content-Type"))
	w.Header().Set("Content-Disposition", fileHttpRsp.Header.Get("Content-Disposition"))
	setFileStorageSourceHeader(w, fileHttpRsp.Header.Get(fileStorageSourceHeader))
	if w.Header().Get(fileStorageSourceHeader) == fileStorageSourceLocal {
		setFileStorageSourceHeader(w, fileStorageSourceRemote)
	}
	w.Header().Set("Content-Length", fileHttpRsp.Header.Get("Content-Length"))

	written, err := io.Copy(w, file)
	if err != nil {
		logger.Error("Failed to write file to response", slog.Any("error", err))
		span.RecordError(err)
		span.SetStatus(codes.Error, "write_remote_response")
		recordFileDownloadServe(r.Context(), "remote", "write_error", written)
		endpoint.InternalServerError(w)
		return
	}
	recordFileDownloadServe(r.Context(), "remote", "success", written)
}

func (fs *FileService) recoverAndServeRemoteDownloadLocally(
	downloadReplicas []*file.InternalDownload,
	w http.ResponseWriter,
	r *http.Request,
) {
	ctx, span := startFileSpan(r.Context(), "file.serve_download.remote_recover_local")
	defer span.End()
	r = r.WithContext(ctx)

	if len(downloadReplicas) == 0 {
		span.SetStatus(codes.Error, "no_replicas")
		recordFileDownloadRecovery(r.Context(), "remote_local_redownload", "no_replicas")
		recordFileDownloadServe(r.Context(), "remote_recover_local", "not_found", -1)
		endpoint.WriteString(w, http.StatusNotFound, "Not Found")
		return
	}

	downloadURL := downloadReplicas[0].URL
	if fs.nodeId == "" {
		if err := fs.getNodeId(r.Context()); err != nil {
			logger.Error("Failed to get node ID for local download recovery",
				slog.String("url", downloadURL),
				slog.Any("error", err),
			)
			span.RecordError(err)
			span.SetStatus(codes.Error, "node_id")
			recordFileDownloadRecovery(r.Context(), "remote_local_redownload", "node_id_error")
			recordFileDownloadServe(r.Context(), "remote_recover_local", "node_id_error", -1)
			endpoint.WriteString(w, http.StatusInternalServerError, "Failed to recover file")
			return
		}
	}

	localDownload := &file.InternalDownload{
		Id:          sdk.Id("dl"),
		URL:         downloadURL,
		NodeId:      fs.nodeId,
		FilePath:    filepath.Join(fs.downloadFolder, EncodeURLtoFileName(downloadURL)),
		Status:      file.DownloadStatusInProgress,
		RetryCount:  nil,
		NextRetryAt: nil,
		LastError:   nil,
	}

	if err := fs.downloadStore.Upsert(localDownload); err != nil {
		logger.Error("Failed to create local download recovery record",
			slog.String("url", downloadURL),
			slog.Any("error", err),
		)
		span.RecordError(err)
		span.SetStatus(codes.Error, "create_record")
		recordFileDownloadRecovery(r.Context(), "remote_local_redownload", "create_record_error")
		recordFileDownloadServe(r.Context(), "remote_recover_local", "create_record_error", -1)
		endpoint.WriteString(w, http.StatusInternalServerError, "Failed to recover file")
		return
	}

	if err := fs.downloadFile(localDownload); err != nil {
		logger.Error("Failed to redownload remote download locally",
			slog.String("url", downloadURL),
			slog.Any("error", err),
		)
		span.RecordError(err)
		span.SetStatus(codes.Error, "download")
		recordFileDownloadRecovery(r.Context(), "remote_local_redownload", "download_error")
		recordFileDownloadServe(r.Context(), "remote_recover_local", "download_error", -1)
		endpoint.WriteString(w, http.StatusInternalServerError, "Failed to recover file")
		return
	}
	recordFileDownloadRecovery(r.Context(), "remote_local_redownload", "download_success")

	downloadIs, err := fs.downloadStore.Query(datastore.Equals(
		[]string{"url"},
		downloadURL,
	)).Find()
	if err != nil {
		logger.Error("Failed to query locally recovered download",
			slog.String("url", downloadURL),
			slog.Any("error", err),
		)
		span.RecordError(err)
		span.SetStatus(codes.Error, "query_recovered")
		recordFileDownloadServe(r.Context(), "remote_recover_local", "query_error", -1)
		endpoint.WriteString(w, http.StatusInternalServerError, "Failed to recover file")
		return
	}
	if len(downloadIs) == 0 {
		logger.Error("Locally recovered download missing after re-query",
			slog.String("url", downloadURL),
		)
		span.SetStatus(codes.Error, "recovered_missing")
		recordFileDownloadServe(r.Context(), "remote_recover_local", "missing", -1)
		endpoint.WriteString(w, http.StatusNotFound, "file not found")
		return
	}

	fs.serveLocalDownload(toDownloads(downloadIs), w, r)
}

func (fs *FileService) promoteDownloadFromStorage(
	ctx context.Context,
	downloadReplicas []*file.InternalDownload,
) (*file.InternalDownload, bool, error) {
	if fs.downloadStorage == nil {
		return nil, false, nil
	}

	if fs.nodeId == "" {
		err := fs.getNodeId(ctx)
		if err != nil {
			return nil, false, errors.Wrap(err, "cannot get node id")
		}
	}

	for _, download := range downloadReplicas {
		if download.Status != file.DownloadStatusCompleted {
			continue
		}

		localPath := filepath.Join(fs.downloadFolder, EncodeURLtoFileName(download.URL))
		storageFilePath := DownloadStorageFilePath(download.URL)
		restored, size, err := fs.restoreDownloadFromStorage(
			ctx,
			storageFilePath,
			localPath,
		)
		if err != nil {
			return nil, false, err
		}
		if !restored {
			continue
		}

		download.NodeId = fs.nodeId
		download.FilePath = localPath
		download.Status = file.DownloadStatusCompleted
		download.TotalSize = size
		download.DownloadedSize = size
		download.RetryCount = nil
		download.NextRetryAt = nil
		download.LastError = nil

		if err := fs.downloadStore.Upsert(download); err != nil {
			return nil, false, errors.Wrap(err, "failed to update promoted download record")
		}

		logger.Info("Promoted remote download from storage",
			slog.String("url", download.URL),
			slog.String("nodeId", fs.nodeId),
		)

		return download, true, nil
	}

	return nil, false, nil
}

func toDownloads(downloadIs []datastore.Row) []*file.InternalDownload {
	ret := []*file.InternalDownload{}
	for _, downloadI := range downloadIs {
		ret = append(ret, downloadI.(*file.InternalDownload))
	}

	return ret
}

func (fs *FileService) isLocalDownload(
	ctx context.Context,
	downloads []*file.InternalDownload,
) (bool, error) {
	if fs.nodeId == "" {
		err := fs.getNodeId(ctx)
		if err != nil {
			return false, errors.Wrap(err, "cannot get node id")
		}
	}

	for _, download := range downloads {
		if download.NodeId == fs.nodeId {
			return true, nil
		}
	}

	return false, nil
}

func (fs *FileService) pickLocalDownload(
	ctx context.Context,
	downloads []*file.InternalDownload,
) (*file.InternalDownload, error) {
	if fs.nodeId == "" {
		err := fs.getNodeId(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "cannot get node id")
		}
	}

	for _, download := range downloads {
		if download.NodeId == fs.nodeId {
			return download, nil
		}
	}

	return nil, fmt.Errorf("download not found")
}

func (fs *FileService) pickRemoteDownloads(
	ctx context.Context,
	downloads []*file.InternalDownload,
) ([]*file.InternalDownload, error) {
	if fs.nodeId == "" {
		err := fs.getNodeId(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "cannot get node id")
		}
	}

	ret := []*file.InternalDownload{}
	for _, download := range downloads {
		if download.NodeId != fs.nodeId {
			ret = append(ret, download)
		}
	}

	return ret, nil
}
