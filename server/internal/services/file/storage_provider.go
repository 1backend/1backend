package fileservice

import (
	"context"
	"io"

	file "github.com/1backend/1backend/server/internal/services/file/types"
)

// StorageProvider defines the abstraction for file operations.
// It allows the service to switch between Local, GCS, or Cached storage
// without changing the core business logic in ServeUpload or UploadFile.
type StorageProvider interface {
	// Open retrieves a file for reading.
	// To avoid unnecessary database lookups, it takes a relative 'filePath'
	// (the FileID) and returns a ReadCloser for streaming, the file size,
	// and the original filename (retrieved from disk or cloud metadata).
	Open(ctx context.Context, filePath string) (rc io.ReadCloser, size int64, err error)

	// Save performs an atomic, blocking write.
	// It is primarily used during the initial Upload phase where the
	// full metadata (*file.Upload) is available to set storage-specific
	// attributes like GCS metadata or specific disk paths.
	Save(ctx context.Context, u *file.Upload, content io.Reader) (written int64, err error)

	// NewWriter is a stream-to-location factory.
	// Unlike Save, it does not consume a reader immediately. Instead, it returns
	// an io.WriteCloser that allows data to be "pushed" to the destination.
	// This is the engine behind the performant CloudCacheProvider: it allows
	// us to pipe data to the local disk while simultaneously serving the user.
	NewWriter(ctx context.Context, filePath string) (io.WriteCloser, error)
}
