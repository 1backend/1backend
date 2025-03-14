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
	"log/slog"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/openorch/openorch/sdk/go/logger"
)

// Define constants for log interval
const LogInterval = 240 * time.Second

type endpointData struct {
	LastLogTime  time.Time
	RequestCount int
	ErrorCount   int
}

var (
	endpointStats map[string]*endpointData
	mu            sync.Mutex
)

func init() {
	endpointStats = make(map[string]*endpointData)
}

// ThrottledLogger Middleware
func ThrottledLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		customWriter := newThrottledResponseWriter(w, r.URL.Path)
		next(customWriter, r)

		duration := time.Since(start)

		mu.Lock()
		data, exists := endpointStats[r.URL.Path]
		if !exists {
			data = &endpointData{LastLogTime: time.Now().Add(-LogInterval), RequestCount: 0}
			endpointStats[r.URL.Path] = data
		}
		data.RequestCount++
		requestCount := data.RequestCount
		shouldLog := time.Since(data.LastLogTime) > LogInterval
		if shouldLog && false {
			logger.Debug("Endpoint request count",
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.Int("statusCode", customWriter.statusCode),
				slog.String("duration", duration.String()),
				slog.Int("requestCount", requestCount),
				slog.Int64("logInterval", int64(LogInterval)))
			data.LastLogTime = time.Now()
			data.RequestCount = 0 // Reset the count after logging
		}
		mu.Unlock()
	}
}

type throttledResponseWriter struct {
	http.ResponseWriter
	statusCode int
	endpoint   string
}

func newThrottledResponseWriter(w http.ResponseWriter, endpoint string) *throttledResponseWriter {
	return &throttledResponseWriter{ResponseWriter: w, statusCode: http.StatusOK, endpoint: endpoint}
}

func (rw *throttledResponseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *throttledResponseWriter) Write(b []byte) (int, error) {
	if rw.statusCode >= 400 {
		rw.Header().Set("Content-Type", "application/json")

		// Convert bytes to string and truncate at the first newline to avoid multiline.
		errorMsgStr := strings.Split(string(b), "\n")[0]
		// Escape double quotes to ensure JSON validity.
		errorMsgStr = strings.ReplaceAll(errorMsgStr, `"`, `\"`)

		// Wrap the error message in a JSON object
		errorMsg := fmt.Sprintf(`{"error": "%s"}`, errorMsgStr)

		// Log the error message with the endpoint name, using the original message
		// Log for 500+ errors and selected 4xx errors
		if rw.statusCode >= 500 || rw.statusCode == 403 || rw.statusCode == 429 {
			logger.Debug("Endpoint returned error",
				slog.String("endpoint", rw.endpoint),
				slog.Int("statusCode", rw.statusCode),
				slog.String("error", strings.TrimSuffix(string(b), "\n")))
		}

		return rw.ResponseWriter.Write([]byte(errorMsg))
	}

	return rw.ResponseWriter.Write(b)
}

func (tw *throttledResponseWriter) Flush() {
	if flusher, ok := tw.ResponseWriter.(http.Flusher); ok {
		flusher.Flush()
	} else {
		logger.Warn("Flushing is not supported by the underlying responsewriter of throttledResponseWriter")
	}
}
