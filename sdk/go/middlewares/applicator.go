/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package middlewares

import (
	"net/http"

	"github.com/1backend/1backend/sdk/go/endpoint"
)

func applicator(
	mws []Middleware,
) func(http.HandlerFunc) http.HandlerFunc {
	return func(h http.HandlerFunc) http.HandlerFunc {
		for _, middleware := range mws {
			h = middleware(h)
		}

		return h
	}
}

func Applicator(
	tr endpoint.TokenRefresher,
	autorefreshOff bool,
) func(http.HandlerFunc) http.HandlerFunc {
	middlewares := []Middleware{
		TokenRefreshMiddleware(tr, autorefreshOff),
		ThrottledLogger,
		Recover,
		CORS,
		GzipDecodeMiddleware,
	}

	return applicator(middlewares)
}
