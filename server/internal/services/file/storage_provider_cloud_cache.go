package fileservice

import (
	"context"
	"io"
	"os"
	"path/filepath"

	file "github.com/1backend/1backend/server/internal/services/file/types"
)

type CloudCacheProvider struct {
	gcs   *GCSProvider
	local *DistributedProvider // We use the local save logic here
}

func NewCloudCacheProvider(
	gcs *GCSProvider,
	local *DistributedProvider,
) *CloudCacheProvider {
	return &CloudCacheProvider{
		gcs:   gcs,
		local: local,
	}
}

func (p *CloudCacheProvider) Open(ctx context.Context, u *file.Upload) (io.ReadCloser, int64, string, error) {
	// Try Local absolute path first
	if f, size, name, err := p.local.Open(ctx, u); err == nil {
		return f, size, name, nil
	}

	// Cache Miss: Pull from GCS
	gcsReader, size, name, err := p.gcs.Open(ctx, u)
	if err != nil {
		return nil, 0, "", err
	}

	// Save to Local absolute path while streaming to user
	os.MkdirAll(filepath.Dir(u.FilePath), 0755)
	cacheFile, _ := os.Create(u.FilePath)

	return &teeReadCloser{
		reader: io.TeeReader(gcsReader, cacheFile),
		closer: gcsReader,
		file:   cacheFile,
	}, size, name, nil
}

func (p *CloudCacheProvider) Save(ctx context.Context, u *file.Upload, content io.Reader) (int64, error) {
	// Save locally first (to the absolute path)
	written, err := p.local.Save(ctx, u, content)
	if err != nil {
		return 0, err
	}

	// Sync the local file to GCS
	f, _ := os.Open(u.FilePath)
	defer f.Close()
	_, err = p.gcs.Save(ctx, u, f)

	return written, err
}

// teeReadCloser is a helper that wraps an io.TeeReader to ensure
// that both the underlying network stream and the local cache file
// are closed correctly when the operation finishes.
type teeReadCloser struct {
	reader io.Reader // This will be the io.TeeReader
	closer io.Closer // This is the original GCS body
	file   *os.File  // This is the local file being written to
}

// Read satisfies the io.Reader interface
func (t *teeReadCloser) Read(p []byte) (n int, err error) {
	return t.reader.Read(p)
}

// Close satisfies the io.Closer interface.
// It ensures the local file is flushed to disk and the network connection is closed.
func (t *teeReadCloser) Close() error {
	// Close the local file first to flush writes
	err1 := t.file.Close()

	// Close the GCS stream
	err2 := t.closer.Close()

	if err1 != nil {
		return err1
	}
	return err2
}
