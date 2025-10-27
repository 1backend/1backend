/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package config_svc

import (
	"time"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type Config struct {
	InternalId string `json:"internalId,omitempty" swagger:"ignore"`

	AppId string `json:"appId" binding:"required"`

	// CamelCased slugs of the config owners
	Id        string    `json:"id" binding:"required"`
	CreatedAt time.Time `json:"createdAt" binding:"required"`
	UpdatedAt time.Time `json:"updatedAt" binding:"required"`

	DataJSON string                 `json:"dataJson" binding:"required"`
	Data     map[string]interface{} `json:"data" binding:"required"`

	Version string `json:"version"`

	// Tags are used to filter on versions, not active configs.
	Tags []string `json:"tags,omitempty"`
}

func (c Config) GetId() string {
	return c.InternalId
}

type Version struct {
	InternalId string `json:"internalId,omitempty" swagger:"ignore"`

	AppId string `json:"appId" binding:"required"`

	// CamelCased slugs of the config owners
	Id      string `json:"id" binding:"required"`
	Version string `json:"version"`

	CreatedAt time.Time `json:"createdAt" binding:"required"`
	UpdatedAt time.Time `json:"updatedAt" binding:"required"`

	DataJSON string                 `json:"dataJson" binding:"required"`
	Data     map[string]interface{} `json:"data" binding:"required"`

	Tags []string `json:"tags,omitempty"`
}

func (v Version) GetId() string {
	return v.InternalId
}

type ListConfigsRequest struct {
	AppHost string `json:"appHost" binding:"required" example:"shoes.com"`

	// Ids are camelCased slugs of the config owners.
	// Specifying only the ids will mean all of the config will be returned
	// for that key.
	//
	// If the configs are large, consider using the `Selector` request field.
	Ids []string `json:"ids,omitempty" swagger:"default=[]"`

	// Selector allows dotPath-based filtering per config owner.
	// Example:
	// {
	//   "user1": ["settings.theme", "featureFlags.enableNewUI"],
	//   "user2": ["settings.language"]
	// }
	Selector map[string][]string `json:"selector,omitempty" swagger:"default={}"`
}

type ListConfigsResponse struct {
	// Configs is a map of camelcase owner slug to Config.
	// Eg.
	// {
	// 	"testUserSlug0": {
	//   "id": "testUserSlug0",
	//   "data": {
	//     "key1": "value1",
	//     "key2": "value2"
	//   }
	// 	}
	// }
	Configs map[string]*Config `json:"configs" binding:"required"`
}

type ListVersionsRequest struct {
	AppHost string `json:"appHost" binding:"required" example:"shoes.com"`

	// Ids are camelCased slugs of the config owners.
	// Specifying only the ids will mean all of the config will be returned
	// for that key.
	//
	// If the configs are large, consider using the `Selector` request field.
	Ids []string `json:"ids,omitempty" swagger:"default=[]"`

	// Selector allows dotPath-based filtering per config owner.
	// Example:
	// {
	//   "user1": ["settings.theme", "featureFlags.enableNewUI"],
	//   "user2": ["settings.language"]
	// }
	Selector map[string][]string `json:"selector,omitempty" swagger:"default={}"`

	Tags []string `json:"tags,omitempty" swagger:"default=[]"`

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

type SaveConfigRequest struct {
	// AppHost can only be specified by users who have the
	// `config-svc:config:edit-on-behalf` permission, who are typically admins.
	//
	// If not specified, the config will be saved for the current app of the user's token.
	AppHost string `json:"appHost,omitempty"`

	// Id is the slug of the owner to save the config for.
	// Only user with the `config-svc:config:edit-on-behalf` can specify this.
	// For everyone else, it is automatically set to the slug of the caller user.
	Id string `json:"id,omitempty"`

	DataJSON string                 `json:"dataJson,omitempty"`
	Data     map[string]interface{} `json:"data,omitempty"`

	Tags []string `json:"tags,omitempty" swagger:"default=[]"`
}

type SaveConfigResponse struct {
}

//
// Event
//

const EventConfigUpdateName = "configUpdate"

type EventConfigUpdate struct {
}

func (e EventConfigUpdate) Name() string {
	return EventConfigUpdateName
}
