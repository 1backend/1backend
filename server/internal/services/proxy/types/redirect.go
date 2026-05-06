/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package proxy_svc

type Redirect struct {
	// Id is the redirect matching key: host plus optional path prefix.
	// It follows the same longest-prefix lookup model as Route.Id.
	Id string `json:"id" binding:"required"`

	// Target is the redirect destination base URL or absolute path.
	// The unmatched request path suffix and query string are appended.
	Target string `json:"target" binding:"required"`

	// StatusCode is the HTTP redirect status code.
	// If omitted, the proxy uses 308 Permanent Redirect.
	StatusCode int `json:"statusCode,omitempty"`
}

type RedirectInput struct {
	Id         string `json:"id" binding:"required"`
	Target     string `json:"target" binding:"required"`
	StatusCode int    `json:"statusCode,omitempty"`
}

func (r *Redirect) GetId() string {
	return r.Id
}

type SaveRedirectsRequest struct {
	Redirects []RedirectInput `json:"redirects"`
}

type SaveRedirectsResponse struct {
	Redirects []Redirect `json:"redirects"`
}

type ListRedirectsRequest struct {
	Ids []string `json:"ids,omitempty"`
}

type ListRedirectsResponse struct {
	Redirects []Redirect `json:"redirects"`
}

type DeleteRedirectsRequest struct {
	Ids []string `json:"ids,omitempty"`
}

type DeleteRedirectsResponse struct{}
