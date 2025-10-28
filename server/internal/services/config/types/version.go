/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package config_svc

import "time"

// Version is a historical snapshot of a Config at a point in time.
type Version struct {
	InternalId string `json:"internalId,omitempty" swagger:"ignore"`

	AppId string `json:"appId" binding:"required"`

	// CamelCased slugs of the config owners.
	// Same as Config.Id.
	Id string `json:"id" binding:"required"`

	Branch    string `json:"branch" binding:"required"`
	VersionId string `json:"versionId" binding:"required"`

	CreatedAt time.Time `json:"createdAt" binding:"required"`
	UpdatedAt time.Time `json:"updatedAt" binding:"required"`

	DataJSON string                 `json:"dataJson" binding:"required"`
	Data     map[string]interface{} `json:"data" binding:"required"`
}

func (v Version) GetId() string {
	return v.InternalId
}

type ListVersionsRequest struct {
	AppHost string `json:"appHost" binding:"required" example:"shoes.com"`

	// Ids are camelCased slugs of the config owners.
	// Specifying only the ids will mean all of the config will be returned
	// for that key.
	//
	// If the configs are large, consider using the `Selector` request field.
	Ids []string `json:"ids,omitempty" swagger:"default=[]"`

	// Branch specifies the branch to get versions from.
	Branch string `json:"branch,omitempty" swagger:"default=main"`

	// Selector allows dotPath-based filtering per config owner.
	// Example:
	// {
	//   "user1": ["settings.theme", "featureFlags.enableNewUI"],
	//   "user2": ["settings.language"]
	// }
	Selector map[string][]string `json:"selector,omitempty" swagger:"default={}"`

	VersionIds []string `json:"versionIds,omitempty"`

	AfterJson string `json:"afterJson,omitempty"`
	Limit     int    `json:"limit,omitempty" swagger:"default=20"`
}

type ListVersionsResponse struct {
	// Versions across all IDs, sorted by CreatedAt descending.
	Versions []*Version `json:"versions"`

	// AfterJson is the cursor to fetch the next page, encoded as a JSON array string.
	// e.g. '["2023-08-01T12:00:00Z","versid-123"]'
	AfterJson string `json:"afterJson,omitempty"`
}
