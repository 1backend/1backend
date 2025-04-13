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

// Lazy is a wrapper that delays starting a service until it’s actually needed.
//
// Some services don’t need to start right when the app launches—especially if
// they don’t run background processes. For these, we can wait to start them
// until the first time they’re used (like when the first HTTP request comes in).
//
// This helps make the server start faster, because we skip setting up services
// that might not even be used.
//
// If starting the service fails, we return a 500 Internal Server Error.
func Lazy(s LazyStarter, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
