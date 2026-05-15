package fileservice

import (
	"context"
	"io"
	"time"

	file "github.com/1backend/1backend/server/internal/services/file/types"
)

type CloudCacheProvider struct {
	cloud StorageProvider
	local StorageProvider
	name  string
}

func NewCloudCacheProvider(
	cloud StorageProvider,
	local StorageProvider,
) *CloudCacheProvider {
	return &CloudCacheProvider{
		cloud: cloud,
		local: local,
		name:  "cloud_cache",
	}
}

func (p *CloudCacheProvider) Open(ctx context.Context, filePath string) (io.ReadCloser, int64, error) {
	// 1. Check local cache
	localStart := time.Now()
	if f, size, err := p.local.Open(ctx, filePath); err == nil {
		recordFileDownloadStorageOperation(ctx, "open", p.backendName("local"), "hit", size, time.Since(localStart))
		return withStorageSource(f, fileStorageSourceLocal), size, nil
	}
	recordFileDownloadStorageOperation(ctx, "open", p.backendName("local"), "miss", -1, time.Since(localStart))

	// 2. Cache Miss: Open cloud stream
	cloudStart := time.Now()
	cloudReader, size, err := p.cloud.Open(ctx, filePath)
	if err != nil {
		recordFileDownloadStorageOperation(ctx, "open", p.backendName("cloud"), "error", -1, time.Since(cloudStart))
		return nil, 0, err
	}
	recordFileDownloadStorageOperation(ctx, "open", p.backendName("cloud"), "success", size, time.Since(cloudStart))

	// 3. Open local writer for caching
	localWriter, err := p.local.NewWriter(ctx, filePath)
	if err != nil {
		// Disk error? Don't break the download, just stream cloud directly.
		return cloudReader, size, nil
	}

	// 4. Return the Tee stream
	return withStorageSource(&teeReadCloser{
		reader: io.TeeReader(cloudReader, localWriter),
		closer: cloudReader,
		writer: localWriter,
	}, fileStorageSourceGCS), size, nil
}

func (p *CloudCacheProvider) OpenCloud(ctx context.Context, filePath string) (io.ReadCloser, int64, error) {
	start := time.Now()
	reader, size, err := p.cloud.Open(ctx, filePath)
	result := "success"
	if err != nil {
		result = "missing"
		size = -1
	}
	recordFileDownloadStorageOperation(ctx, "open", p.backendName("cloud_probe"), result, size, time.Since(start))
	if err != nil {
		return reader, size, err
	}
	return withStorageSource(reader, fileStorageSourceGCS), size, nil
}

func (p *CloudCacheProvider) Save(ctx context.Context, u *file.Upload, content io.Reader) (int64, error) {
	// Write-through: Local first, then Cloud
	localStart := time.Now()
	written, err := p.local.Save(ctx, u, content)
	if err != nil {
		recordFileDownloadStorageOperation(ctx, "save", p.backendName("local"), "error", -1, time.Since(localStart))
		return 0, err
	}
	recordFileDownloadStorageOperation(ctx, "save", p.backendName("local"), "success", written, time.Since(localStart))

	// Use local file as source for Cloud upload
	f, _, err := p.local.Open(ctx, u.FilePath)
	if err != nil {
		return written, err
	}
	defer f.Close()

	cloudStart := time.Now()
	cloudWritten, err := p.cloud.Save(ctx, u, f)
	result := "success"
	if err != nil {
		result = "error"
	}
	recordFileDownloadStorageOperation(ctx, "save", p.backendName("cloud"), result, cloudWritten, time.Since(cloudStart))
	return written, err
}

// GetPath satisfies the interface but we don't need it here anymore
func (p *CloudCacheProvider) NewWriter(ctx context.Context, f string) (io.WriteCloser, error) {
	return p.local.NewWriter(ctx, f)
}

func (p *CloudCacheProvider) Delete(ctx context.Context, filePath string) error {
	_ = p.local.Delete(ctx, filePath)
	return p.cloud.Delete(ctx, filePath)
}

func (p *CloudCacheProvider) backendName(layer string) string {
	if p.name == "" {
		return "cloud_cache_" + layer
	}
	return p.name + "_" + layer
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
