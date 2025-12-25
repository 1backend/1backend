package fileservice

import (
	"context"
	"io"

	"cloud.google.com/go/storage"
	file "github.com/1backend/1backend/server/internal/services/file/types"
)

type GCSProvider struct {
	client *storage.Client
	bucket string
}

func NewGCSProvider(
	client *storage.Client,
	bucket string,
) *GCSProvider {
	return &GCSProvider{
		client: client,
		bucket: bucket,
	}
}

func (p *GCSProvider) Open(ctx context.Context, u *file.Upload) (io.ReadCloser, int64, string, error) {
	obj := p.client.Bucket(p.bucket).Object(u.FileId)
	attrs, err := obj.Attrs(ctx)
	if err != nil {
		return nil, 0, "", err
	}

	reader, err := obj.NewReader(ctx)
	return reader, attrs.Size, u.FileName, err
}

func (p *GCSProvider) Save(ctx context.Context, u *file.Upload, content io.Reader) (int64, error) {
	obj := p.client.Bucket(p.bucket).Object(u.FileId)
	w := obj.NewWriter(ctx)
	written, err := io.Copy(w, content)
	if err != nil {
		return 0, err
	}
	return written, w.Close()
}
