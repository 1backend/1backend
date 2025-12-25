package fileservice

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	file "github.com/1backend/1backend/server/internal/services/file/types"
)

type LocalProvider struct {
	uploadFolder string
}

func NewLocalProvider(uploadFolder string) *LocalProvider {
	return &LocalProvider{uploadFolder: uploadFolder}
}

func (p *LocalProvider) Open(ctx context.Context, filePath string) (io.ReadCloser, int64, string, error) {
	// Construct absolute path: /root/.1backend/uploads/ + file_abc
	absPath := filepath.Join(p.uploadFolder, filePath)

	info, err := os.Stat(absPath)
	if err != nil {
		return nil, 0, "", err
	}
	if info.IsDir() {
		return nil, 0, "", fmt.Errorf("path is a directory")
	}

	f, err := os.Open(absPath)
	return f, info.Size(), info.Name(), err
}

func (p *LocalProvider) Save(ctx context.Context, u *file.Upload, content io.Reader) (int64, error) {
	absPath := filepath.Join(p.uploadFolder, u.FilePath)

	// Create directory if it doesn't exist
	if err := os.MkdirAll(filepath.Dir(absPath), 0755); err != nil {
		return 0, err
	}

	dstFile, err := os.Create(absPath)
	if err != nil {
		return 0, err
	}
	defer dstFile.Close()

	return io.Copy(dstFile, content)
}
