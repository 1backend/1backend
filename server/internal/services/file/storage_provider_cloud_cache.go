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
	local *LocalProvider
}

func NewCloudCacheProvider(gcs *GCSProvider, local *LocalProvider) *CloudCacheProvider {
	return &CloudCacheProvider{
		gcs:   gcs,
		local: local,
	}
}

func (p *CloudCacheProvider) Open(ctx context.Context, filePath string) (io.ReadCloser, int64, string, error) {
	// 1. Try Local Disk first
	if f, size, fileName, err := p.local.Open(ctx, filePath); err == nil {
		return f, size, fileName, nil
	}

	// 2. Cache Miss: Pull from GCS
	gcsReader, size, fileName, err := p.gcs.Open(ctx, filePath)
	if err != nil {
		return nil, 0, "", err
	}

	// 3. Prepare local path for caching
	absPath := filepath.Join(p.local.uploadFolder, filePath)
	os.MkdirAll(filepath.Dir(absPath), 0755)

	cacheFile, err := os.Create(absPath)
	if err != nil {
		// If disk is full/locked, stream GCS directly so service stays up
		return gcsReader, size, "", nil
	}

	// 4. Stream to User and Disk simultaneously
	return &teeReadCloser{
		reader: io.TeeReader(gcsReader, cacheFile),
		closer: gcsReader,
		file:   cacheFile,
	}, size, fileName, nil
}

func (p *CloudCacheProvider) Save(ctx context.Context, u *file.Upload, content io.Reader) (int64, error) {
	// 1. Write to local disk first
	written, err := p.local.Save(ctx, u, content)
	if err != nil {
		return 0, err
	}

	// 2. Re-open the local file to stream up to GCS
	absPath := filepath.Join(p.local.uploadFolder, u.FilePath)
	f, err := os.Open(absPath)
	if err != nil {
		return written, nil
	}
	defer f.Close()

	// 3. Sync to GCS
	_, err = p.gcs.Save(ctx, u, f)
	return written, err
}

// teeReadCloser is a helper that wraps an io.TeeReader to ensure
// that both the underlying network stream (GCS) and the local cache file
// are closed correctly when the operation finishes.
type teeReadCloser struct {
	reader io.Reader // This will be the io.TeeReader(gcsReader, localFile)
	closer io.Closer // This is the original GCS body (to close the connection)
	file   *os.File  // This is the local file being written to (the cache)
}

// Read satisfies the io.Reader interface by reading from the TeeReader.
// As data is read from GCS, it is automatically written to the local file.
func (t *teeReadCloser) Read(p []byte) (n int, err error) {
	return t.reader.Read(p)
}

// Close satisfies the io.Closer interface.
// This is critical: it ensures the local file is flushed/closed and the
// GCS network connection is released back to the pool.
func (t *teeReadCloser) Close() error {
	// Close the local file first to ensure all data is flushed to disk
	err1 := t.file.Close()

	// Close the GCS stream
	err2 := t.closer.Close()

	if err1 != nil {
		return err1
	}
	return err2
}
