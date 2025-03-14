/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package container_svc

import (
	"time"
)

type Log struct {
	Id string `json:"id"`

	// Node Id
	// Please see the documentation for the envar OPENORCH_NODE_ID
	NodeId string `json:"nodeId"`

	// ContainerId is the raw underlying container ID.
	// Eg. Docker container id. Node local.
	ContainerId string `json:"containerId"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	Content string `json:"content"`
}

func (l *Log) GetId() string {
	return l.Id
}

type ListLogsRequest struct {
	NodeId      string `json:"nodeId,omitempty"`
	ContainerId string `json:"containerId,omitempty"`
	Limit       int64  `json:"limit,omitempty"`
}

type ListLogsResponse struct {
	Logs []*Log `json:"logs,omitempty"`
}
