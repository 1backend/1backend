/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package config_svc

type ErrorResponse struct {
	Error string `json:"error"`
}

type Config struct {
	Id        string                 `json:"id,omitempty"`
	Namespace string                 `json:"namespace"          swagger:"default=default"`
	DataJSON  string                 `json:"dataJson,omitempty"`
	Data      map[string]interface{} `json:"data"                                         binding:"required"`
}

func (c Config) GetId() string {
	return c.Id
}

type GetConfigRequest struct {
	Namespace string `json:"namespace" swagger:"default=default"`
}

type GetConfigResponse struct {
	Config *Config `json:"config"`
}

type SaveConfigRequest struct {
	Config *Config `json:"config"`
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
