package fileservice

import (
	"context"
	"io"
	"path"

	file "github.com/1backend/1backend/server/internal/services/file/types"
)

// PrefixedStorageProvider scopes all paths under a fixed prefix.
// Useful for separating uploads vs downloads in the same cloud bucket.
type PrefixedStorageProvider struct {
	base   StorageProvider
	prefix string
}

func (p *PrefixedStorageProvider) prefixed(filePath string) string {
	if p.prefix == "" {
		return filePath
	}
	return path.Join(p.prefix, filePath)
}

func (p *PrefixedStorageProvider) Open(ctx context.Context, filePath string) (io.ReadCloser, int64, error) {
	return p.base.Open(ctx, p.prefixed(filePath))
}

func (p *PrefixedStorageProvider) Save(ctx context.Context, u *file.Upload, content io.Reader) (int64, error) {
	cp := *u
	cp.FilePath = p.prefixed(cp.FilePath)
	return p.base.Save(ctx, &cp, content)
}

func (p *PrefixedStorageProvider) NewWriter(ctx context.Context, filePath string) (io.WriteCloser, error) {
	return p.base.NewWriter(ctx, p.prefixed(filePath))
}

func (p *PrefixedStorageProvider) Delete(ctx context.Context, filePath string) error {
	return p.base.Delete(ctx, p.prefixed(filePath))
}
