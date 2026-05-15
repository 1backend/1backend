package fileservice

import (
	"context"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"github.com/1backend/1backend/sdk/go/logger"
	types "github.com/1backend/1backend/server/internal/services/file/types"
)

type cloudOpener interface {
	OpenCloud(ctx context.Context, filePath string) (io.ReadCloser, int64, error)
}

func (dm *FileService) restoreDownloadFromStorage(
	ctx context.Context,
	storageFilePath string,
	localFilePath string,
) (bool, int64, error) {
	if dm.downloadStorage == nil {
		return false, 0, nil
	}

	start := time.Now()
	src, size, err := dm.downloadStorage.Open(ctx, storageFilePath)
	if err != nil {
		recordFileDownloadStorageOperation(ctx, "open", "download_storage", "miss", -1, time.Since(start))
		return false, 0, nil
	}
	defer src.Close()
	recordFileDownloadStorageOperation(ctx, "open", "download_storage", "success", size, time.Since(start))

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

func (dm *FileService) maybeBackfillDownloadToStorage(
	download *types.InternalDownload,
) {
	if dm.downloadStorage == nil ||
		dm.downloadStorageBackfillEvery == 0 ||
		download == nil ||
		download.Status != types.DownloadStatusCompleted {
		return
	}

	if dm.downloadStorageBackfillCounter.Add(1)%dm.downloadStorageBackfillEvery != 0 {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	start := time.Now()
	storageFilePath := DownloadStorageFilePath(download.URL)

	opener := func(ctx context.Context, filePath string) (io.ReadCloser, int64, error) {
		if cloud, ok := dm.downloadStorage.(cloudOpener); ok {
			return cloud.OpenCloud(ctx, filePath)
		}
		return dm.downloadStorage.Open(ctx, filePath)
	}

	src, _, err := opener(ctx, storageFilePath)
	if err == nil {
		_ = src.Close()
		recordFileDownloadGCSBackfill(ctx, "present", time.Since(start))
		return
	}

	if _, statErr := os.Stat(download.FilePath); statErr != nil {
		logger.Warn("Cannot backfill download to storage because local file is missing",
			slog.String("url", download.URL),
			slog.String("filePath", download.FilePath),
			slog.Any("error", statErr),
		)
		recordFileDownloadGCSBackfill(ctx, "missing_local", time.Since(start))
		return
	}

	if err := dm.persistDownloadToStorage(ctx, storageFilePath, download.FilePath); err != nil {
		logger.Warn("Failed to backfill download to storage",
			slog.String("url", download.URL),
			slog.String("filePath", download.FilePath),
			slog.Any("error", err),
		)
		recordFileDownloadGCSBackfill(ctx, "save_error", time.Since(start))
		return
	}

	recordFileDownloadGCSBackfill(ctx, "saved", time.Since(start))
}

func (dm *FileService) persistDownloadToStorage(
	ctx context.Context,
	storageFilePath string,
	localFilePath string,
) error {
	if dm.downloadStorage == nil {
		return nil
	}

	stat, statErr := os.Stat(localFilePath)
	if statErr != nil {
		return statErr
	}

	f, err := os.Open(localFilePath)
	if err != nil {
		return err
	}
	defer f.Close()

	start := time.Now()
	written, err := dm.downloadStorage.Save(ctx, &types.Upload{
		FilePath: storageFilePath,
	}, f)
	result := "success"
	bytes := written
	if err != nil {
		result = "error"
		bytes = stat.Size()
	}
	recordFileDownloadStorageOperation(ctx, "save", "download_storage", result, bytes, time.Since(start))
	return err
}
