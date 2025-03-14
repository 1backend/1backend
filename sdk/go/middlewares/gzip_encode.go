/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package middlewares

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

type gzipResponseWriter struct {
	http.ResponseWriter
	Writer io.Writer
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func GzipEncodeMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			gzWriter := gzip.NewWriter(w)
			defer gzWriter.Close()

			w.Header().Set("Content-Encoding", "gzip")
			gzResponseWriter := gzipResponseWriter{
				ResponseWriter: w,
				Writer:         gzWriter,
			}

			next.ServeHTTP(gzResponseWriter, r)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
