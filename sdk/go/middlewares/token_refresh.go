/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package middlewares

import (
	"log/slog"
	"net/http"

	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
)

// TokenRefreshMiddleware ensures the JWT token is valid.
// If it's expired, it refreshes it and updates the request header in-place.
func TokenRefreshMiddleware(tr endpoint.TokenRefresher, autorefreshOff bool) func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if !autorefreshOff &&
				// @todo this is cruft, this list should be provided by options
				r.URL.Path != "/user-svc/refresh-token" &&
				r.URL.Path != "/user-svc/public-key" &&
				r.URL.Path != "/user-svc/login" {
				if tr == nil {
					panic("token refresher is nil" + r.URL.Path)
				}

				_, _, err := tr.EnsureValidToken(r)
				if err != nil {
					logger.Error("Token refresh middleware error",
						slog.String("path", r.URL.Path),
						slog.String("method", r.Method),
						slog.Any("error", err),
					)
					http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
					return
				}
			}

			// Token is refreshed in-place in the request header if needed.
			next(w, r)
		}
	}
}
