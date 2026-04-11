package fileservice

import (
	"context"
	"io"
	"os"
	"path/filepath"

	types "github.com/1backend/1backend/server/internal/services/file/types"
)

func (dm *FileService) restoreDownloadFromStorage(
	ctx context.Context,
	storageFilePath string,
	localFilePath string,
) (bool, int64, error) {
	if dm.downloadStorage == nil {
		return false, 0, nil
	}

	src, size, err := dm.downloadStorage.Open(ctx, storageFilePath)
	if err != nil {
		return false, 0, nil
	}
	defer src.Close()

	if err := os.MkdirAll(filepath.Dir(localFilePath), 0755); err != nil {
		return false, 0, err
	}

	dst, err := os.Create(localFilePath)
	if err != nil {
		return false, 0, err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return false, 0, err
	}

	_ = os.Remove(localFilePath + ".part")

	return true, size, nil
}

func (dm *FileService) persistDownloadToStorage(
	ctx context.Context,
	storageFilePath string,
	localFilePath string,
) error {
	if dm.downloadStorage == nil {
		return nil
	}

	f, err := os.Open(localFilePath)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = dm.downloadStorage.Save(ctx, &types.Upload{
		FilePath: storageFilePath,
	}, f)
	return err
}
