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

func Applicator(
	mws []Middleware,
) func(http.HandlerFunc) http.HandlerFunc {
	return func(h http.HandlerFunc) http.HandlerFunc {
		for _, middleware := range mws {
			h = middleware(h)
		}

		return h
	}
}

var DefaultMiddlewares = []Middleware{
	ThrottledLogger,
	Recover,
	CORS,
	GzipDecodeMiddleware,
}

var DefaultApplicator = Applicator(DefaultMiddlewares)

type Starter interface {
	// Idempotent start function
	Start() error
}

// Lazy is a wrapper for services that can be lazy loaded.
// Eg., services which don't have background process loops that need to be started immediately
// can be started on demand; when the first request comes in.
//
// This speeds up the startup time of the server.
func Lazy(s Starter, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := s.Start(); err != nil {
			endpoint.WriteErr(w, http.StatusInternalServerError, err)
			return
		}

		next(w, r)
	}
}
