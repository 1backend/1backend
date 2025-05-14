/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/

package endpoint

import (
	"log/slog"
	"net/http"

	"github.com/1backend/1backend/sdk/go/logger"
)

// WriteErr writes an error message to the response writer with the specified status code.
// It should only be used when "proxying" errors from other endpoints as
// internal errors should not be indiscriminately returned to the client.
// Usually you should use WriteString instead, unless proxying.
func WriteErr(w http.ResponseWriter, statusCode int, err error) {
	errMsg := "error is missing"
	if err != nil {
		errMsg = err.Error()
	}

	w.WriteHeader(statusCode)
	_, err = w.Write([]byte(errMsg))
	if err != nil {
		logger.Error("Error writing response",
			slog.Any("error", err),
		)
	}
}

func WriteString(w http.ResponseWriter, statusCode int, str string) {
	w.WriteHeader(statusCode)
	_, err := w.Write([]byte(str))
	if err != nil {
		logger.Error("Error writing response",
			slog.Any("error", err),
		)
	}
}

// InternalServerError is used frequently, so we define it here for convenience.
func InternalServerError(w http.ResponseWriter) {
	WriteString(w, http.StatusInternalServerError, "Internal Server Error")
}

func Unauthorized(w http.ResponseWriter) {
	WriteString(w, http.StatusUnauthorized, "Unauthorized")
}
