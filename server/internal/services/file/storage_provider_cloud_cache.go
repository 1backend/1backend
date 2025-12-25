package fileservice

import (
	"context"
	"io"

	file "github.com/1backend/1backend/server/internal/services/file/types"
)

type CloudCacheProvider struct {
	cloud StorageProvider
	local StorageProvider
}

func NewCloudCacheProvider(
	cloud StorageProvider,
	local StorageProvider,
) *CloudCacheProvider {
	return &CloudCacheProvider{
		cloud: cloud,
		local: local,
	}
}

func (p *CloudCacheProvider) Open(ctx context.Context, filePath string) (io.ReadCloser, int64, error) {
	// 1. Check local cache
	if f, size, err := p.local.Open(ctx, filePath); err == nil {
		return f, size, nil
	}

	// 2. Cache Miss: Open cloud stream
	cloudReader, size, err := p.cloud.Open(ctx, filePath)
	if err != nil {
		return nil, 0, err
	}

	// 3. Open local writer for caching
	localWriter, err := p.local.NewWriter(ctx, filePath)
	if err != nil {
		// Disk error? Don't break the download, just stream cloud directly.
		return cloudReader, size, nil
	}

	// 4. Return the Tee stream
	return &teeReadCloser{
		reader: io.TeeReader(cloudReader, localWriter),
		closer: cloudReader,
		writer: localWriter,
	}, size, nil
}

func (p *CloudCacheProvider) Save(ctx context.Context, u *file.Upload, content io.Reader) (int64, error) {
	// Write-through: Local first, then Cloud
	written, err := p.local.Save(ctx, u, content)
	if err != nil {
		return 0, err
	}

	// Use local file as source for Cloud upload
	f, _, err := p.local.Open(ctx, u.FilePath)
	if err != nil {
		return written, nil
	}
	defer f.Close()

	_, err = p.cloud.Save(ctx, u, f)
	return written, err
}

// GetPath satisfies the interface but we don't need it here anymore
func (p *CloudCacheProvider) NewWriter(ctx context.Context, f string) (io.WriteCloser, error) {
	return p.local.NewWriter(ctx, f)
}

// teeReadCloser is a helper that wraps an io.TeeReader to ensure
// that both the underlying network stream (Cloud) and the local cache writer
// are closed correctly when the operation finishes.
type teeReadCloser struct {
	reader io.Reader      // The io.TeeReader(cloudReader, localWriter)
	closer io.Closer      // The original cloud body stream
	writer io.WriteCloser // The local cache file/stream
}

// Read satisfies the io.Reader interface.
func (t *teeReadCloser) Read(p []byte) (n int, err error) {
	return t.reader.Read(p)
}

// Close ensures both the local cache is finalized and the cloud connection is released.
func (t *teeReadCloser) Close() error {
	// 1. Close the local writer first to flush the cache to disk
	err1 := t.writer.Close()

	// 2. Close the cloud stream
	err2 := t.closer.Close()

	if err1 != nil {
		return err1
	}
	return err2
}
