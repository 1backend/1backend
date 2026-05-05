/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package user_svc

import (
	"time"

	"github.com/1backend/1backend/sdk/go/datastore"
)

type TokenRefreshActivity struct {
	InternalId string `json:"internalId,omitempty" swagger:"ignore"`

	Id string `json:"id" binding:"required"`

	CreatedAt time.Time `json:"createdAt" binding:"required"`
	UpdatedAt time.Time `json:"updatedAt" binding:"required"`

	AppId   string `json:"appId" binding:"required"`
	AppHost string `json:"appHost" binding:"required"`
	UserId  string `json:"userId" binding:"required"`
	Device  string `json:"device" binding:"required"`

	BucketStart time.Time `json:"bucketStart" binding:"required"`
	BucketEnd   time.Time `json:"bucketEnd" binding:"required"`

	FirstRefreshedAt time.Time `json:"firstRefreshedAt" binding:"required"`
	LastRefreshedAt  time.Time `json:"lastRefreshedAt" binding:"required"`
	RefreshCount     int64     `json:"refreshCount" binding:"required"`

	Published   bool       `json:"published"`
	PublishedAt *time.Time `json:"publishedAt,omitempty"`
}

func (a *TokenRefreshActivity) GetId() string {
	return a.Id
}

func (a *TokenRefreshActivity) Indexes() []datastore.Index {
	return []datastore.Index{
		{
			Fields:    []string{"published", "bucketEnd"},
			Ascending: true,
		},
	}
}
