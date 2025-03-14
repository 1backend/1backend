/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package file_svc

import "time"

type ErrorResponse struct {
	Error string `json:"error"`
}

type DownloadStatus string

var (
	DownloadStatusInProgress DownloadStatus = "inProgress"
	DownloadStatusPaused     DownloadStatus = "paused"
	DownloadStatusCompleted  DownloadStatus = "completed"
	DownloadStatusErrored    DownloadStatus = "errored"
)

// Download is the internal type for downloads.
type InternalDownload struct {
	Id             string         `json:"id"`
	URL            string         `json:"url"`
	NodeId         string         `json:"nodeId"`
	FilePath       string         `json:"filePath"`
	DownloadedSize int64          `json:"downloadedSize"`
	TotalSize      int64          `json:"totalSize"`
	Status         DownloadStatus `json:"status"`
}

func (d InternalDownload) GetId() string {
	return d.Id
}

// Download record
type Download struct {
	Id        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	URL      string  `json:"url"`
	FileName string  `json:"fileName"`
	Progress float64 `json:"progress,omitempty"`

	// DownloadedBytes exists to show the download progress in terms of the number of bytes already downloaded.
	DownloadedBytes int64 `json:"downloadedBytes" format:"int64"`

	// FileSize is the full final downloaded file size.
	FileSize int64          `json:"fileSize,omitempty" format:"int64"`
	Status   DownloadStatus `json:"status"`

	FilePath string `json:"filePath,omitempty"`
	Error    string `json:"error,omitempty"`
}

type DownloadFileRequest struct {
	URL        string `json:"url"                  binding:"required"`
	FolderPath string `json:"folderPath,omitempty"`
	// FileName   *string `json:"fileName,omitempty"`
}

type GetDownloadRequest struct {
	Url string `json:"url"`
}

type GetDownloadResponse struct {
	Exists   bool      `json:"exists"   binding:"required"`
	Download *Download `json:"download"`
}

type DownloadFileResponse struct{}

type DownloadsRequest struct{}

type DownloadsResponse struct {
	Downloads []Download `json:"downloads"`
}

//
// Events
//

const EventDownloadStatusChangeName = "downloadStatusChange"

type EventDownloadStatusChange struct {
}

func (e EventDownloadStatusChange) Name() string {
	return EventDownloadStatusChangeName
}
