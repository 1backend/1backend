package proxyservice

import (
	"compress/gzip"
	"net/http"
	"strings"
)

// List of types worth compressing
var compressibleTypes = map[string]bool{
	"text/html":                true,
	"text/css":                 true,
	"text/plain":               true,
	"text/javascript":          true,
	"application/javascript":   true,
	"application/json":         true,
	"application/xml":          true,
	"application/x-javascript": true,
	"image/svg+xml":            true,
}

type gzipResponseBuffer struct {
	http.ResponseWriter
	gw           *gzip.Writer
	compressible bool
	wroteHeader  bool
}

func (g *gzipResponseBuffer) WriteHeader(code int) {
	if g.wroteHeader {
		return
	}

	// Check if the backend response is a type we want to compress
	contentType := g.Header().Get("Content-Type")
	// Handle types like "text/html; charset=utf-8"
	baseType := strings.Split(contentType, ";")[0]

	if compressibleTypes[baseType] && g.Header().Get("Content-Encoding") == "" {
		g.compressible = true
		g.Header().Set("Content-Encoding", "gzip")
		g.Header().Del("Content-Length")
		g.Header().Add("Vary", "Accept-Encoding")

		// Initialize the writer only if we are actually compressing
		g.gw, _ = gzip.NewWriterLevel(g.ResponseWriter, 5)
	}

	g.wroteHeader = true
	g.ResponseWriter.WriteHeader(code)
}

func (g *gzipResponseBuffer) Write(b []byte) (int, error) {
	if !g.wroteHeader {
		g.WriteHeader(http.StatusOK)
	}
	if g.compressible && g.gw != nil {
		return g.gw.Write(b)
	}
	return g.ResponseWriter.Write(b)
}

func (g *gzipResponseBuffer) Close() {
	if g.gw != nil {
		g.gw.Close()
	}
}

func withSmartGzip(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. Skip if client doesn't support gzip
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}

		// 2. Wrap the response writer
		gb := &gzipResponseBuffer{ResponseWriter: w}
		defer gb.Close()

		next.ServeHTTP(gb, r)
	})
}
