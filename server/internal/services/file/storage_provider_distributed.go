package fileservice

import (
	"context"
	"fmt"
	"io"
	"os"
	"strconv"

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	file "github.com/1backend/1backend/server/internal/services/file/types"
)

type DistributedProvider struct {
	fs *FileService
}

func NewDistributedProvider(fs *FileService) *DistributedProvider {
	return &DistributedProvider{
		fs: fs,
	}
}

func (p *DistributedProvider) Save(ctx context.Context, u *file.Upload, content io.Reader) (int64, error) {
	f, err := os.Create(u.FilePath)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return io.Copy(f, content)
}

func (p *DistributedProvider) Open(ctx context.Context, u *file.Upload) (io.ReadCloser, int64, string, error) {
	// Check if the file is physically on this machine
	if u.NodeId == p.fs.nodeId {
		info, err := os.Stat(u.FilePath)
		if err != nil {
			return nil, 0, "", err
		}
		f, err := os.Open(u.FilePath)
		return f, info.Size(), u.FileName, err
	}

	// File is on another node, open a remote stream
	return p.openRemoteStream(ctx, u)
}

func (p *DistributedProvider) openRemoteStream(ctx context.Context, u *file.Upload) (io.ReadCloser, int64, string, error) {
	token, err := p.fs.getToken()
	if err != nil {
		return nil, 0, "", err
	}

	// Find the node where the file lives
	nodesRsp, _, err := p.fs.options.ClientFactory.
		Client(client.WithToken(token)).
		RegistrySvcAPI.ListNodes(ctx).
		Body(openapi.RegistrySvcListNodesRequest{Ids: []string{u.NodeId}}).
		Execute()

	if err != nil || len(nodesRsp.Nodes) == 0 {
		return nil, 0, "", fmt.Errorf("node %s not found", u.NodeId)
	}
	node := nodesRsp.Nodes[0]

	// Request the file from the remote node
	// IMPORTANT: We use the Client's ability to return the raw stream
	fileStream, fileHttpRsp, err := p.fs.options.ClientFactory.
		Client(client.WithAddress(node.Url), client.WithToken(token)).
		FileSvcAPI.
		ServeUpload(ctx, u.FileId).
		Execute()

	if err != nil {
		return nil, 0, "", err
	}

	size, _ := strconv.ParseInt(fileHttpRsp.Header.Get("Content-Length"), 10, 64)

	// We return the fileStream (io.ReadCloser) directly.
	// The HTTP handler will stream this to the user.
	return fileStream, size, u.FileName, nil
}
