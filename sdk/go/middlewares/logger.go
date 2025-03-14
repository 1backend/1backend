/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

// Logger Middleware
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		customWriter := newResponseWriter(w, r.URL.Path) // Pass the endpoint
		next(customWriter, r)

		duration := time.Since(start)
		log.Printf("%s %s %d %v", r.Method, r.URL.Path, customWriter.statusCode, duration)
	}
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
	endpoint   string // Add endpoint field
}

func newResponseWriter(w http.ResponseWriter, endpoint string) *responseWriter {
	// Default status code is 200 OK unless explicitly set otherwise
	return &responseWriter{ResponseWriter: w, statusCode: http.StatusOK, endpoint: endpoint}
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	if rw.statusCode >= 400 {
		// Convert bytes to string and truncate at the first newline to avoid multiline.
		errorMsgStr := strings.Split(string(b), "\n")[0]
		// Escape double quotes to ensure JSON validity.
		errorMsgStr = strings.ReplaceAll(errorMsgStr, `"`, `\"`)

		// Wrap the error message in a JSON object
		errorMsg := fmt.Sprintf(`{"error": "%s"}`, errorMsgStr)

		// Log the error message with the endpoint name, using the original message
		log.Printf("Error at %s: %s", rw.endpoint, strings.TrimSuffix(string(b), "\n"))

		return rw.ResponseWriter.Write([]byte(errorMsg))
	}
	return rw.ResponseWriter.Write(b)
}
