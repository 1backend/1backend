package proxyservice

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"

	"github.com/andybalholm/brotli"
)

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

type compressionBuffer struct {
	http.ResponseWriter
	writer       io.WriteCloser
	encoding     string // "br" or "gzip"
	compressible bool
	wroteHeader  bool
}

func (c *compressionBuffer) WriteHeader(code int) {
	if c.wroteHeader {
		return
	}

	contentType := c.Header().Get("Content-Type")
	baseType := strings.Split(contentType, ";")[0]

	// Only compress if the type is in our list and not already encoded
	if compressibleTypes[baseType] && c.Header().Get("Content-Encoding") == "" {
		c.compressible = true
		c.Header().Set("Content-Encoding", c.encoding)
		c.Header().Del("Content-Length")
		c.Header().Add("Vary", "Accept-Encoding")

		if c.encoding == "br" {
			// Level 4 is a good balance between speed and ratio for on-the-fly
			c.writer = brotli.NewWriterLevel(c.ResponseWriter, 4)
		} else if c.encoding == "gzip" {
			c.writer, _ = gzip.NewWriterLevel(c.ResponseWriter, gzip.DefaultCompression)
		}
	}

	c.wroteHeader = true
	c.ResponseWriter.WriteHeader(code)
}

func (c *compressionBuffer) Write(b []byte) (int, error) {
	if !c.wroteHeader {
		c.WriteHeader(http.StatusOK)
	}
	if c.compressible && c.writer != nil {
		return c.writer.Write(b)
	}
	return c.ResponseWriter.Write(b)
}

func (c *compressionBuffer) Close() {
	if c.writer != nil {
		c.writer.Close()
	}
}

func WithCompression(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accept := r.Header.Get("Accept-Encoding")
		encoding := ""

		// Brotli has priority
		if strings.Contains(accept, "br") {
			encoding = "br"
		} else if strings.Contains(accept, "gzip") {
			encoding = "gzip"
		}

		if encoding == "" {
			next.ServeHTTP(w, r)
			return
		}

		cb := &compressionBuffer{
			ResponseWriter: w,
			encoding:       encoding,
		}
		defer cb.Close()

		next.ServeHTTP(cb, r)
	})
}
