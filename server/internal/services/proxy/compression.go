package proxyservice

import (
	"compress/gzip"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/andybalholm/brotli"
)

const minCompressSize = 1024 // Only compress if > 1KB

var (
	gzipPool = sync.Pool{
		New: func() any {
			w, _ := gzip.NewWriterLevel(io.Discard, gzip.DefaultCompression)
			return w
		},
	}
	brPool = sync.Pool{
		New: func() any {
			return brotli.NewWriterLevel(io.Discard, 4)
		},
	}
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
	"application/x-font-ttf":   true,
	"font/opentype":            true,
}

type compressionBuffer struct {
	http.ResponseWriter
	writer       io.WriteCloser
	encoding     string
	compressible bool
	wroteHeader  bool
}

func (c *compressionBuffer) WriteHeader(code int) {
	if c.wroteHeader {
		return
	}

	// 1. Check Content-Length if present
	contentLengthStr := c.Header().Get("Content-Length")
	if contentLengthStr != "" {
		if size, err := strconv.Atoi(contentLengthStr); err == nil && size < minCompressSize {
			c.compressible = false
			c.wroteHeader = true
			c.ResponseWriter.WriteHeader(code)
			return
		}
	}

	contentType := c.Header().Get("Content-Type")
	baseType := strings.Split(contentType, ";")[0]
	isNoTransform := strings.Contains(c.Header().Get("Cache-Control"), "no-transform")

	// 2. Validate Type and Encoding status
	if compressibleTypes[baseType] && c.Header().Get("Content-Encoding") == "" && !isNoTransform {
		c.compressible = true
		c.Header().Set("Content-Encoding", c.encoding)
		c.Header().Del("Content-Length")
		c.Header().Add("Vary", "Accept-Encoding")

		if c.encoding == "gzip" {
			gw := gzipPool.Get().(*gzip.Writer)
			gw.Reset(c.ResponseWriter)
			c.writer = gw
		} else if c.encoding == "br" {
			bw := brPool.Get().(*brotli.Writer)
			bw.Reset(c.ResponseWriter)
			c.writer = bw
		}
	}

	c.wroteHeader = true
	c.ResponseWriter.WriteHeader(code)
}

func (c *compressionBuffer) Write(b []byte) (int, error) {
	if !c.wroteHeader {
		// If we don't know the size yet and the first chunk is tiny,
		// we could technically buffer here, but for a proxy,
		// it's usually safer to just proceed with WriteHeader.
		c.WriteHeader(http.StatusOK)
	}
	if c.compressible && c.writer != nil {
		return c.writer.Write(b)
	}
	return c.ResponseWriter.Write(b)
}

func (c *compressionBuffer) Close() {
	if c.writer == nil {
		return
	}
	c.writer.Close()

	if c.encoding == "gzip" {
		if gw, ok := c.writer.(*gzip.Writer); ok {
			gzipPool.Put(gw)
		}
	} else if c.encoding == "br" {
		if bw, ok := c.writer.(*brotli.Writer); ok {
			brPool.Put(bw)
		}
	}
	// Important: nil out the writer so double closes don't break pools
	c.writer = nil
}

func WithCompression(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accept := r.Header.Get("Accept-Encoding")
		encoding := ""

		if strings.Contains(accept, "gzip") {
			encoding = "gzip"
		} else if strings.Contains(accept, "br") {
			encoding = "br"
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
