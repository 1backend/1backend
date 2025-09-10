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
	// Id is the routing key: host plus optional path prefix.
	// Example:
	//   "x.com"              -> root of the domain
	//   "x.com/path1"        -> microfrontend at /path1
	//   "x.com/path1/path2"  -> deeper microfrontend mounted at /path1/path2
	//
	// Use case: multiple microfrontends served under the same host but
	// separated by URL path segments. For example:
	//   - Marketing site at x.com
	//   - Dashboard at x.com/app
	//   - Admin UI at x.com/app/admin
	//
	// Lookup algorithm:
	//   1. Take the request host and path (e.g. "x.com/app/admin/users").
	//   2. Try to match the longest registered Id by progressively stripping
	//      trailing path segments:
	//         - x.com/app/admin/users   (no match)
	//         - x.com/app/admin         (match -> admin UI)
	//   3. If still no match, strip again:
	//         - x.com/app               (match -> dashboard)
	//   4. If still no match, fallback to host-only route:
	//         - x.com                   (match -> marketing site)
	//   5. If no host-only route exists, return 404.
	//
	// This provides deterministic longest-prefix routing without regex or
	// rule engines, keeping the model simple but enabling path-based
	// microfrontend composition.
	Id     string `json:"id"`
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

type ListRoutesRequest struct {
	Ids []string `json:"ids,omitempty"`
}

type ListRoutesResponse struct {
	Routes []Route `json:"routes"`
}
