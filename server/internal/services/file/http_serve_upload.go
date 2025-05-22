package fileservice

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	file "github.com/1backend/1backend/server/internal/services/file/types"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
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
	isLocal, err := fs.isLocal(r.Context(), uploadReplicas)
	if err != nil {
		logger.Error("Error checking if upload is local", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}

	if isLocal {
		fs.serveLocal(uploadReplicas, w, r)
	} else {
		fs.serveRemote(uploadReplicas, w, r)
	}
}

func (fs *FileService) serveLocal(
	uploadReplicas []*file.Upload,
	w http.ResponseWriter,
	r *http.Request,
) {
	upload, err := fs.pickLocal(r.Context(), uploadReplicas)
	if err != nil {
		logger.Error("Failed to pick local upload", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}

	fileInfo, err := os.Stat(upload.FilePath)
	if err != nil || fileInfo.IsDir() {
		endpoint.WriteString(w, http.StatusNotFound, "file not found")
		return
	}

	contentType := mime.TypeByExtension(filepath.Ext(upload.FileName))
	if contentType == "" {
		contentType = "application/octet-stream"
	}
	w.Header().Set("Content-Type", contentType)
	w.Header().
		Set("Content-Disposition", "attachment; filename=\""+sanitizeFilename(upload.FileName)+"\"")
	w.Header().Set("Content-Length", strconv.FormatInt(fileInfo.Size(), 10))

	srcFile, err := os.Open(upload.FilePath)
	if err != nil {
		logger.Error("Failed to open file", slog.Any("error", err))
		endpoint.InternalServerError(w)
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

func (fs *FileService) serveRemote(
	uploadReplicas []*file.Upload,
	w http.ResponseWriter,
	r *http.Request,
) {
	uploads, err := fs.pickRemotes(r.Context(), uploadReplicas)
	if err != nil {
		logger.Error("Failed to pick remote upload", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}

	nodeIds := []string{}
	for _, upload := range uploads {
		nodeIds = append(nodeIds, upload.NodeId)
	}

	token, err := fs.getToken()
	if err != nil {
		logger.Error("Failed to get token", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}

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
		endpoint.InternalServerError(w)
		return
	}
	nodes := nodesRsp.Nodes

	if len(nodes) == 0 {
		endpoint.WriteString(w, http.StatusNotFound, "Not Found")
		return
	}

	node := nodes[0]

	// todo it would be probably better to stream this ourselves here but for now it will do
	file, fileHttpRsp, err := fs.options.ClientFactory.
		Client(client.WithAddress(node.Url), client.WithToken(token)).
		FileSvcAPI.
		ServeUpload(r.Context(), uploads[0].FileId).
		Execute()
	if err != nil {
		logger.Error("Failed to serve upload", slog.Any("error", err))
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

func toUploads(uploadIs []datastore.Row) []*file.Upload {
	ret := []*file.Upload{}
	for _, uploadI := range uploadIs {
		ret = append(ret, uploadI.(*file.Upload))
	}

	return ret
}

func (fs *FileService) isLocal(ctx context.Context, uploads []*file.Upload) (bool, error) {
	if fs.nodeId == "" {
		err := fs.getNodeId(ctx)
		if err != nil {
			return false, errors.Wrap(err, "cannot get node id")
		}
	}

	for _, upload := range uploads {
		if upload.NodeId == fs.nodeId {
			return true, nil
		}
	}

	return false, nil
}

func (fs *FileService) pickLocal(
	ctx context.Context,
	uploads []*file.Upload,
) (*file.Upload, error) {
	if fs.nodeId == "" {
		err := fs.getNodeId(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "cannot get node id")
		}
	}

	for _, upload := range uploads {
		if upload.NodeId == fs.nodeId {
			return upload, nil
		}
	}

	return nil, fmt.Errorf("upload not found")
}

func (fs *FileService) pickRemotes(
	ctx context.Context,
	uploads []*file.Upload,
) ([]*file.Upload, error) {
	if fs.nodeId == "" {
		err := fs.getNodeId(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "cannot get node id")
		}
	}

	ret := []*file.Upload{}
	for _, upload := range uploads {
		if upload.NodeId != fs.nodeId {
			ret = append(ret, upload)
		}
	}

	return ret, nil
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
