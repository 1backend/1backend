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
	"path/filepath"

	types "github.com/openorch/openorch/server/internal/services/file/types"
	"github.com/pkg/errors"
)

func (ds *FileService) list() ([]types.Download, error) {
	ds.lock.Lock()
	defer ds.lock.Unlock()

	downloadIs, err := ds.downloadStore.Query().Find()
	if err != nil {
		return nil, errors.Wrap(err, "error listing downloads")
	}

	var downloadDetailsList []types.Download
	for _, downloadI := range downloadIs {
		download := downloadI.(*types.InternalDownload)
		downloadDetail := downloadToDownloadDetails(download)
		downloadDetailsList = append(downloadDetailsList, *downloadDetail)
	}

	return downloadDetailsList, nil
}

func downloadToDownloadDetails(
	download *types.InternalDownload,
) *types.Download {
	if download == nil {
		return nil
	}
	fileName := filepath.Base(download.FilePath)

	var progress float64
	if download.TotalSize > 0 {
		computedProgress := float64(
			(download.DownloadedSize * 100),
		) / float64(
			download.TotalSize,
		)
		progress = computedProgress
	}

	var fullFileSize int64
	if download.TotalSize > 0 {
		fullFileSize = download.TotalSize
	}

	downloadDetail := types.Download{
		Id:              download.Id,
		URL:             download.URL,
		FileName:        fileName,
		Progress:        progress,
		DownloadedBytes: download.DownloadedSize,
		FileSize:        fullFileSize,
		Status:          download.Status,
		FilePath:        download.FilePath,
	}

	return &downloadDetail
}
