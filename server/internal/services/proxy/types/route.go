/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package proxy_svc

type ErrorResponse struct {
	Error string `json:"error"`
}

type Route struct {
	// Id is the host itself, e.g. "test.localhost"
	Id        string `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`

	Target string `json:"target"`
}

type RouteInput struct {
	// Id is the host itself, e.g. "test.localhost"
	Id string `json:"id"`

	Target string `json:"target"`
}

func (r *Route) GetId() string {
	return r.Id
}

type SaveRoutesRequest struct {
	Routes []RouteInput `json:"routes"`
}

type SaveRoutesResponse struct {
	Routes []Route `json:"routes"`
}
