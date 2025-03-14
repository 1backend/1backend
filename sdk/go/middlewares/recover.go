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
	"runtime/debug"

	"github.com/openorch/openorch/sdk/go/logger"
)

// Panic Recovery Middleware
func Recover(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				stackTrace := debug.Stack()

				logger.Logger.Error("Recovered from panic",
					slog.Any("error", err),
					slog.String("stackTrace", string(stackTrace)),
				)

				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
			}
		}()
		next(w, r)
	}
}
