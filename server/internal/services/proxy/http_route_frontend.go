/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package proxyservice

import (
	"log/slog"
	"net/http"

	"github.com/1backend/1backend/sdk/go/logger"
)

func (cs *ProxyService) RouteFrontend(w http.ResponseWriter, r *http.Request) {
	statusCode, err := cs.routeFrontend(w, r)

	if err != nil {
		logger.Error("Error proxying OPTIONS request",
			slog.String("path", r.URL.Path),
			slog.String("error", err.Error()),
		)
		if r.Method == http.MethodOptions {
			// We don't want to write errors to response on OPTIONS requests
			// because the response won't be visible in the chrome dev tools.
			// We log it instead.

			return
		}

		w.WriteHeader(statusCode)
		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			logger.Error("Error writing response",
				slog.String("error", err.Error()),
			)
		}
		return
	}
}

func (cs *ProxyService) routeFrontend(w http.ResponseWriter, r *http.Request) (int, error) {
	return 0, nil
}
