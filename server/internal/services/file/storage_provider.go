package fileservice

import (
	"context"
	"io"

	file "github.com/1backend/1backend/server/internal/services/file/types"
)

type StorageProvider interface {
	Open(ctx context.Context, u *file.Upload) (io.ReadCloser, int64, string, error)
	Save(ctx context.Context, u *file.Upload, content io.Reader) (int64, error)
}
