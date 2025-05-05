/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package service

import (
	"net/http"

	"github.com/1backend/1backend/sdk/go/endpoint"
)

// Lazy is a wrapper for endpoints that delays service setup until it’s actually needed.
//
// Some services dependencies don't need to be loaded right away, only when an endpoint is called.
// Delaying the setup can help reduce startup time and resource usage.
//
// If starting the service fails, we return a 500 Internal Server Error.
func Lazy(s LazyStarter, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			// We don't want to throw errors on OPTIONS requests
			// because the response won't be visible in the chrome dev tools.
			next(w, r)
			return
		}

		if err := s.LazyStart(); err != nil {
			endpoint.WriteErr(w, http.StatusInternalServerError, err)
			return
		}

		next(w, r)
	}
}

// LazyStarter is an interface for things that need to be started before use.
// The LazyStart() function should prepare anything that needs to be ready,
// like registering themselves in the registry and setting up permissions.
//
// It's meant to be safe to call multiple times (idempotent), so calling
// LazyStart() again won't cause issues or repeat the setup unnecessarily.
//
// This is useful for services that don’t need to start background tasks right away,
// but still require some setup before being used.
type LazyStarter interface {
	LazyStart() error
}
