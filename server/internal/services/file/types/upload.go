/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package file_svc

import "time"

// Upload represents a single instance of a file upload in the system.
// A single file upload can result in multiple Upload records due to replication
// across different nodes or instances. Each record corresponds to a specific
// replica of the file on a particular node.
type Upload struct {
	// Unique ID for this replica
	Id string `json:"id"`

	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	// ID of the node storing this replica
	NodeId string `json:"nodeId"`

	// Logical file ID spanning all replicas
	FileId string `json:"fileId"`

	// FilePath is the full node local path of the file
	FilePath string `json:"filePath"`

	// Filename is the original name of the file
	FileName string `json:"fileName"`

	UserId   string `json:"userId,omitempty"`
	FileSize int64  `json:"fileSize"         binding:"required" format:"int64"`
}

func (u Upload) GetId() string {
	return u.Id
}

type UploadFileResponse struct {
	Upload Upload `json:"upload,omitempty"`
}

type ListUploadsRequest struct {
	UserId string `json:"userId,omitempty"`
	Limit  int    `json:"limit,omitempty"`

	// After time value
	After string `json:"after,omitempty"`
}

type ListUploadsResponse struct {
	Uploads []Upload `json:"uploads,omitempty"`
}
