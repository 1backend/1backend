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

	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	openapi "github.com/1backend/1backend/clients/go"
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
// @Failure 400 {object} file.ErrorResponse "invalid download URL"
// @Failure 400 {object} file.ErrorResponse "error parsing download URL"
// @Failure 404 {object} file.ErrorResponse "file not found"
// @Failure 500 {object} file.ErrorResponse "internal server error"
// @Router /file-svc/serve/download/{url} [get]
func (fs *FileService) ServeDownload(
	w http.ResponseWriter,
	r *http.Request,
) {
	vars := mux.Vars(r)
	ur, err := url.PathUnescape(vars["url"])
	if err != nil {
		endpoint.WriteString(w, http.StatusBadRequest, "invalid download URL")
		return
	}

	downloadReplicaIs, err := fs.downloadStore.Query(datastore.Equals(
		[]string{"url"},
		ur,
	)).Find()
	if err != nil {
		logger.Error("Error Querying Download", slog.Any("error", err))
		endpoint.WriteString(w, http.StatusInternalServerError, err.Error())
		return
	}
	if len(downloadReplicaIs) == 0 {
		endpoint.WriteString(w, http.StatusNotFound, "file not found")
		return
	}

	downloadReplicas := toDownloads(downloadReplicaIs)
	isLocal, err := fs.isLocalDownload(r.Context(), downloadReplicas)
	if err != nil {
		logger.Error("Error checking if download is local", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}
	if isLocal {
		fs.serveLocalDownload(downloadReplicas, w, r)
	} else {
		fs.serveRemoteDownload(downloadReplicas, w, r)
	}
}

func (fs *FileService) serveLocalDownload(
	downloadReplicas []*file.InternalDownload,
	w http.ResponseWriter,
	r *http.Request,
) {
	download, err := fs.pickLocalDownload(r.Context(), downloadReplicas)
	if err != nil {
		logger.Error("Failed to pick local download", slog.Any("error", err))
		endpoint.WriteString(w, http.StatusInternalServerError, "Failed to pick local download")
		return
	}

	fileInfo, err := os.Stat(download.FilePath)
	if err != nil || fileInfo.IsDir() {
		endpoint.WriteString(w, http.StatusNotFound, "file not found")
		return
	}

	parsedURL, err := url.Parse(download.URL)
	if err != nil {
		endpoint.WriteString(w, http.StatusBadRequest, "error parsing download URL")
		return
	}

	fileName := path.Base(parsedURL.Path)
	contentType := mime.TypeByExtension(filepath.Ext(fileName))
	if contentType == "" {
		contentType = "application/octet-stream"
	}
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Disposition", "attachment; filename=\""+sanitizeFilename(fileName)+"\"")
	w.Header().Set("Content-Length", strconv.FormatInt(fileInfo.Size(), 10))

	srcFile, err := os.Open(download.FilePath)
	if err != nil {
		endpoint.WriteString(w, http.StatusInternalServerError, "Failed to open file")
		return
	}
	defer srcFile.Close()

	_, err = io.Copy(w, srcFile)
	if err != nil {
		logger.Error("Failed to write file to response", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}
}

func (fs *FileService) serveRemoteDownload(
	downloadReplicas []*file.InternalDownload,
	w http.ResponseWriter,
	r *http.Request,
) {
	downloads, err := fs.pickRemoteDownloads(r.Context(), downloadReplicas)
	if err != nil {
		logger.Error("Failed to pick remote download", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}

	nodeIds := []string{}
	for _, download := range downloads {
		nodeIds = append(nodeIds, download.NodeId)
	}

	token, err := fs.getToken()
	if err != nil {
		logger.Error("Failed to get token", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}

	nodesRsp, _, err := fs.clientFactory.
		Client(client.WithToken(token)).
		RegistrySvcAPI.ListNodes(r.Context()).
		Body(
			openapi.RegistrySvcListNodesRequest{
				Ids: nodeIds,
			},
		).Execute()
	if err != nil {
		logger.Error("Failed to list nodes", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}
	nodes := nodesRsp.Nodes

	if len(nodes) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("not found"))
		return
	}

	node := nodes[0]

	// todo it would be probably better to stream this ourselves here but for now it will do
	file, fileHttpRsp, err := fs.clientFactory.
		Client(client.WithAddress(node.Url), client.WithToken(token)).
		FileSvcAPI.
		ServeDownload(r.Context(), downloads[0].URL).
		Execute()
	if err != nil {
		logger.Error("Failed to serve download", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}

	defer file.Close()

	w.Header().Set("Content-Type", fileHttpRsp.Header.Get("Content-Type"))
	w.Header().Set("Content-Disposition", fileHttpRsp.Header.Get("Content-Disposition"))
	w.Header().Set("Content-Length", fileHttpRsp.Header.Get("Content-Length"))

	_, err = io.Copy(w, file)
	if err != nil {
		logger.Error("Failed to write file to response", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}
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
