/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package config_svc

import "time"

type ErrorResponse struct {
	Error string `json:"error"`
}

type Config struct {
	// Id of the config.
	// It is deterministically created from the app and the key.
	Id        string    `json:"id" binding:"required"`
	CreatedAt time.Time `json:"createdAt" binding:"required"`
	UpdatedAt time.Time `json:"updatedAt" binding:"required"`

	App string `json:"app,omitempty"`
	Key string `json:"key,omitempty"`

	DataJSON string                 `json:"dataJson" binding:"required"`
	Data     map[string]interface{} `json:"data" binding:"required"`
}

func (c Config) GetId() string {
	return c.Id
}

type ListConfigsRequest struct {
	App string `json:"app" swagger:"default=default"`

	// Keys are camelCased slugs of the config owners.
	// Specifying only the keys will mean all of the config will be returned
	// for that key.
	//
	// If the configs are large, consider using the `Selector` request field.
	Keys []string `json:"keys,omitempty" swagger:"default=[]"`

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

type SaveConfigRequest struct {
	// App can only be specified by users who have the
	// `config-svc:config:edit-on-behalf` permission, who are typically admins.
	//
	// If not specified, the config will be saved for the current app of the user's token.
	App string `json:"app,omitempty"`

	// Key is the slug of the owner to save the config for.
	// Only user with the `config-svc:config:edit-on-behalf` can specify this.
	// For everyone else, it is automatically set to the slug of the caller user.
	Key string `json:"key,omitempty"`

	DataJSON string                 `json:"dataJson,omitempty"`
	Data     map[string]interface{} `json:"data,omitempty"`
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
