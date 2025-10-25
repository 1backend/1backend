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
)

/*
CORS middleware for multi-tenant APIs.

Context:
- One shared API backend serves multiple untrusted tenant frontends.
- Each tenant authenticates using explicit header tokens (Authorization), not cookies.

Design:
1. Access-Control-Allow-Origin
   - Echoes the incoming Origin header.
   - Safe because no implicit credentials (cookies, HTTP auth) are used.
   - Each frontend sends its own Authorization header, isolating tenants by token scope.

2. Access-Control-Allow-Credentials
   - Set to "false" (or omitted).
   - Prevents browsers from sending cookies or HTTP auth automatically.
   - If true + blind echo, a hostile tenant site could read another tenant’s responses.

3. Vary: Origin
   - Always present.
   - Ensures shared caches treat each Origin as a distinct variant, avoiding cross-tenant leakage.

4. Access-Control-Allow-Headers
   - Includes Authorization since auth tokens travel in headers.

5. Security model
   - Requests from any origin are allowed, but only valid tokens grant access.
   - Safe for adversarial tenants as long as credentials remain explicit.
   - If cookie-based auth is ever introduced, strict origin validation becomes mandatory.

Summary:
- Token-based, multi-tenant-safe CORS.
- No cookies, no implicit credentials.
- Cache isolation via Vary header.
*/

// CORS Middleware
func CORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Content-Encrypted")
		w.Header().Set("Access-Control-Allow-Credentials", "false")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Vary", "Origin")

		origin := r.Header.Get("Origin")
		if origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		} else {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}
