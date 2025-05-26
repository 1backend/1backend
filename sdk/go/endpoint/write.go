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
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/1backend/1backend/sdk/go/logger"
)

// WriteErr writes an error message to the response writer with the specified status code.
// It should only be used when "proxying" errors from other endpoints as
// internal errors should not be indiscriminately returned to the client.
// Usually you should use WriteString instead, unless proxying.
// It logs any errors that occur during the write operation.
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

// WriteString writes a string to the response writer with the specified status code.
// It logs any errors that occur during the write operation.
func WriteString(w http.ResponseWriter, statusCode int, str string) {
	w.WriteHeader(statusCode)
	_, err := w.Write([]byte(str))
	if err != nil {
		logger.Error("Error writing response",
			slog.Any("error", err),
		)
	}
}

// WriteJSON writes a JSON response to the response writer with the specified status code.
func WriteJSON(w http.ResponseWriter, statusCode int, v interface{}) {
	if v == nil {
		return
	}

	jsonData, err := json.Marshal(v)
	if err != nil {
		logger.Error("Error marshalling JSON",
			slog.Any("error", err),
		)
		WriteErr(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(statusCode)
	_, err = w.Write(jsonData)
	if err != nil {
		logger.Error("Error writing JSON response",
			slog.Any("error", err),
		)

		return
	}
}

// InternalServerError is used frequently, so we define it here for convenience.
func InternalServerError(w http.ResponseWriter) {
	WriteString(w, http.StatusInternalServerError, "Internal Server Error")
}

// Unauthorized is used frequently, so we define it here for convenience.
func Unauthorized(w http.ResponseWriter) {
	WriteString(w, http.StatusUnauthorized, "Unauthorized")
}
