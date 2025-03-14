/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package fileservice

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	openapi "github.com/openorch/openorch/clients/go"
	sdk "github.com/openorch/openorch/sdk/go"
	"github.com/openorch/openorch/sdk/go/logger"
	types "github.com/openorch/openorch/server/internal/services/file/types"
	"github.com/pkg/errors"
)

/*
Starts or resumes a download.
Can resume downloads not found in the JSON statefile.
*/
func (dm *FileService) download(ctx context.Context, url, downloadDir string) error {
	if downloadDir == "" {
		downloadDir = dm.downloadFolder
	}

	err := dm.getNodeId(ctx)
	if err != nil {
		return errors.Wrap(err, "cannot get node id")
	}

	safeFileName := EncodeURLtoFileName(url)
	safeFullFilePath := filepath.Join(downloadDir, safeFileName)
	safePartialFilePath := filepath.Join(downloadDir, safeFileName+".part")

	fullSize, fullFileExists, err := checkFileExistsAndSize(safeFullFilePath)
	if err != nil {
		return err
	}
	partialSize, partialFileExists, err := checkFileExistsAndSize(
		safePartialFilePath,
	)
	if err != nil {
		return err
	}

	var (
		download *types.InternalDownload
		exists   bool
	)

	f := func() error {
		download, exists = dm.getDownload(url)

		if !exists {
			if fullFileExists {
				download = &types.InternalDownload{
					Id:             sdk.Id("upl"),
					URL:            url,
					NodeId:         dm.nodeId,
					FilePath:       safeFullFilePath,
					Status:         types.DownloadStatusCompleted,
					TotalSize:      fullSize,
					DownloadedSize: fullSize,
				}
			} else if partialFileExists {
				download = &types.InternalDownload{
					Id:             sdk.Id("upl"),
					URL:            url,
					NodeId:         dm.nodeId,
					FilePath:       safeFullFilePath,
					Status:         types.DownloadStatusInProgress,
					DownloadedSize: partialSize,
				}
			} else {
				download = &types.InternalDownload{
					Id:       sdk.Id("upl"),
					URL:      url,
					NodeId:   dm.nodeId,
					FilePath: safeFullFilePath,
					Status:   types.DownloadStatusInProgress,
				}
			}
		} else {
			// This corrects a potential mismatch between the file size value
			// in the downloads.json and the actual file size which happens
			// if the daemon exists after writing to the file but before reflecting that
			// change in the downloads.json.
			// Search for @transaction-problem in this file
			if partialFileExists {
				download.DownloadedSize = partialSize
			}
			if download.Status == types.DownloadStatusPaused {
				download.Status = types.DownloadStatusInProgress
			}
		}

		return nil
	}

	err = f()
	if err != nil {
		return nil
	}

	err = dm.downloadStore.Upsert(download)
	if err != nil {
		return errors.Wrap(err, "failed to upsert download")
	}

	if dm.SyncDownloads {
		return dm.downloadFile(download)
	} else {
		go func() {
			err := dm.downloadFile(download)
			if err != nil {
				logger.Error("Error downlading file",
					slog.String("url", download.URL),
					slog.String("error", err.Error()),
				)
			}
		}()
	}

	return nil
}

func (dm *FileService) downloadFile(d *types.InternalDownload) error {
	if d.Status == types.DownloadStatusCompleted {
		return nil
	}
	if d.Status == types.DownloadStatusPaused {
		// this should never happen as Do sets this to inProgress
		return fmt.Errorf("cannot download file with status paused")
	}

	out, err := os.OpenFile(
		d.FilePath+".part",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666,
	)
	if err != nil {
		return errors.Wrap(err, "opening file for download")
	}
	defer out.Close()

	req, err := http.NewRequest("GET", d.URL, nil)
	if err != nil {
		return err
	}

	if d.DownloadedSize > 0 {
		req.Header.Set("Range", fmt.Sprintf("bytes=%d-", d.DownloadedSize))
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	totalSize, _ := getTotalSizeFromHeaders(resp)
	if err != nil {
		fmt.Printf("Error retrieving total size: %v\n", err)
	}

	if resp.StatusCode == http.StatusPartialContent ||
		resp.StatusCode == http.StatusOK {
		buffer := make([]byte, 1024*256) // 256KB buffer
		for {
			if d.Status == types.DownloadStatusPaused {
				return nil
			}
			n, err := resp.Body.Read(buffer)
			if n > 0 {
				_, err = out.Write(buffer[:n])
				// @transaction-problem
				if err != nil {
					return err
				}
				d.DownloadedSize += int64(n)
				if d.TotalSize == 0 && totalSize != 0 {
					d.TotalSize = totalSize
				}
				err = dm.downloadStore.Upsert(d)
				if err != nil {
					return errors.Wrap(err, "failed to upsert download")
				}

				ev := types.EventDownloadStatusChange{}
				_, err = dm.clientFactory.Client(sdk.WithToken(dm.token)).
					FirehoseSvcAPI.PublishEvent(context.Background()).
					Event(openapi.FirehoseSvcEventPublishRequest{
						Event: &openapi.FirehoseSvcEvent{
							Name: openapi.PtrString(ev.Name()),
							Data: nil,
						},
					}).
					Execute()
			}
			if err == io.EOF {
				break
			}
			if err != nil {
				return err
			}
		}
		out.Close()

		err = os.Rename(d.FilePath+".part", d.FilePath)
		if err != nil {
			return err
		}

		d.Status = types.DownloadStatusCompleted
		err = dm.downloadStore.Upsert(d)
		if err != nil {
			return errors.Wrap(err, "failed to upsert download")
		}
	} else {
		fmt.Printf("Failed to download: %s, status code: %d\n", d.URL, resp.StatusCode)
	}

	return nil
}

func getTotalSizeFromHeaders(resp *http.Response) (int64, error) {
	// initial downloads without range request headers
	contentLength := resp.Header.Get("Content-Length")
	contentRange := resp.Header.Get("Content-Range")
	if contentLength == "" && contentRange == "" {
		return 0, fmt.Errorf(
			"Content-Length and Content-Range header is missing",
		)
	}

	if contentLength != "" {
		return strconv.ParseInt(contentLength, 10, 64)
	}

	// Content-Range format is typically "bytes start-end/total"
	parts := strings.Split(contentRange, "/")
	if len(parts) != 2 {
		return 0, fmt.Errorf("invalid Content-Range format")
	}

	totalSize, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return 0, fmt.Errorf(
			"error parsing total size from Content-Range: %v",
			err,
		)
	}

	return totalSize, nil
}
