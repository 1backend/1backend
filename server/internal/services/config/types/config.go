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
	// Id is simply the app of the config.
	Id        string    `json:"id" binding:"required"`
	CreatedAt time.Time `json:"createdAt" binding:"required"`
	UpdatedAt time.Time `json:"updatedAt" binding:"required"`

	DataJSON string                 `json:"dataJson" binding:"required"`
	Data     map[string]interface{} `json:"data" binding:"required"`
}

func (c Config) GetId() string {
	return c.Id
}

type GetConfigRequest struct {
	App string `json:"app" swagger:"default=default"`
}

type GetConfigResponse struct {
	Config *Config `json:"config"`
}

type SaveConfigRequest struct {
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
