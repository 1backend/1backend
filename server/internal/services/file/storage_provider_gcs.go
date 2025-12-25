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

func (p *GCSProvider) Open(ctx context.Context, filePath string) (io.ReadCloser, int64, string, error) {
	obj := p.client.Bucket(p.bucket).Object(filePath)
	attrs, err := obj.Attrs(ctx)
	if err != nil {
		return nil, 0, "", err
	}

	// Retrieve the name from metadata.
	// Fallback to the filePath (ID) if the metadata is missing.
	name := attrs.Metadata["original_name"]
	if name == "" {
		name = attrs.Name
	}

	reader, err := obj.NewReader(ctx)
	return reader, attrs.Size, name, err
}

func (p *GCSProvider) Save(ctx context.Context, u *file.Upload, content io.Reader) (int64, error) {
	obj := p.client.Bucket(p.bucket).Object(u.FilePath)
	w := obj.NewWriter(ctx)

	// Store the original filename in custom metadata
	w.Metadata = map[string]string{
		"original_name": u.FileName,
	}

	written, err := io.Copy(w, content)
	if err != nil {
		return 0, err
	}
	return written, w.Close()
}
