/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package proxyservice

import (
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net"
	"net/http"
	"strings"
	"syscall"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/logger"

	proxy "github.com/1backend/1backend/server/internal/services/proxy/types"
)

func (cs *ProxyService) RouteFrontend(w http.ResponseWriter, r *http.Request) {
	logger.Debug("Edge proxying",
		slog.String("host", r.Host),
		slog.String("path", r.URL.Path),
	)

	target, err := cs.findRouteTarget(r.Host, r.URL.Path, r.URL.RawQuery)
	if err != nil {
		// check for HTTPError
		if herr, ok := err.(*sdk.HTTPError); ok {
			http.Error(w, herr.Msg, herr.Code)
			return
		}
		logger.Error("Error finding route target",
			slog.String("host", r.Host),
			slog.String("path", r.URL.Path),
			slog.Any("error", err),
		)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if !strings.HasPrefix(target, "http") {
		target = "http://" + target
	}

	req, err := http.NewRequest(r.Method, target, r.Body)
	if err != nil {
		logger.Error(
			"Failed to create request",
			slog.String("target", target),
			slog.String("host", r.Host),
			slog.String("path", r.URL.Path),
			slog.String("target", target),
			slog.Any("error", err),
		)
		http.Error(w, "Failed to create proxy request", http.StatusInternalServerError)
		return
	}

	req.Header = make(http.Header)
	for k, vv := range r.Header {
		for _, v := range vv {
			req.Header.Add(k, v)
		}
	}

	req.Host = r.Host

	proto := r.Header.Get("X-Forwarded-Proto")
	if proto == "" {
		if r.TLS != nil {
			proto = "https"
		} else {
			proto = "http"
		}
	}
	req.Header.Set("X-Forwarded-Proto", proto)

	clientIP := r.RemoteAddr
	if prior, ok := r.Header["X-Forwarded-For"]; ok && len(prior) > 0 {
		clientIP = prior[0] + ", " + clientIP
	}
	req.Header.Set("X-Forwarded-For", clientIP)
	req.Header.Set("X-Forwarded-Host", r.Host)

	host, port, err := net.SplitHostPort(r.Host)
	if err != nil || port == "" {
		if r.TLS != nil {
			port = "443"
		} else {
			port = "80"
		}
	}
	req.Header.Set("X-Forwarded-Port", port)

	req.Header.Set("X-Real-IP", host)

	forwarded := fmt.Sprintf("for=%q;host=%q;proto=%q", r.RemoteAddr, r.Host, proto)
	req.Header.Set("Forwarded", forwarded)

	if req.Header.Get("User-Agent") == "" {
		req.Header.Set("User-Agent", "1backend-proxy")
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("Error edge proxying",
			slog.String("host", r.Host),
			slog.String("path", r.URL.Path),
			slog.Any("error", err),
		)

		http.Error(w, "Failed to proxy request", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	for k, v := range resp.Header {
		switch http.CanonicalHeaderKey(k) {
		// Hop-by-hop headers are meaningful only for a single transport link.
		// They must not be forwarded by proxies or gateways per RFC 9110 §7.6.1.
		// Keeping them can corrupt HTTP/2 framing or break keep-alive logic.

		case "Content-Length", // handled automatically by Go based on written bytes
			"Transfer-Encoding", // conflicts with automatic chunked encoding
			"Connection",        // hop-by-hop; invalid in HTTP/2
			"Keep-Alive",        // deprecated; invalid in HTTP/2
			"Upgrade",           // used only for protocol switches (e.g., WebSocket)
			"Proxy-Connection",  // nonstandard, used by some clients incorrectly
			"Trailer":           // managed by Go's chunked encoder if needed
			continue
		}
		for _, vv := range v {
			w.Header().Add(k, vv)
		}
	}

	w.WriteHeader(resp.StatusCode)

	// Use a fixed 32 KB buffer for predictable memory use and fewer allocations.
	// Copy the upstream response body directly to the downstream client.
	// This streams the data instead of buffering the entire response in memory.
	buf := make([]byte, 32*1024)
	_, err = io.CopyBuffer(w, resp.Body, buf)

	// Ignore common disconnect errors — they just mean the client closed the connection early.
	// Log only unexpected I/O problems.
	if err != nil && !errors.Is(err, syscall.EPIPE) && !strings.Contains(err.Error(), "stream closed") {
		logger.Error("Proxy stream error", slog.String("error", err.Error()))
	}

	// Optionally flush if ResponseWriter supports it (e.g., for HTTP/1.1 chunked or slow clients).
	// In HTTP/2, Go auto-flushes frames, so it's rarely needed. Harmless but not required.
	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	}
}

func (cs *ProxyService) findRouteTarget(host, path, rawQuery string) (string, error) {
	if path == "" {
		path = "/"
	}
	var candidates []string
	parts := strings.Split(path, "/")
	for i := len(parts); i >= 0; i-- {
		prefix := strings.Join(parts[:i], "/")
		if prefix == "" {
			candidates = append(candidates, host)
		} else {
			candidates = append(candidates, host+prefix)
		}
	}

	var route *proxy.Route
	for _, key := range candidates {
		ri, found, err := cs.routeStore.Query(datastore.Id(key)).FindOne()
		if err != nil {
			// datastore failure = 500
			return "", sdk.NewHTTPError(http.StatusInternalServerError,
				fmt.Sprintf("failed to query route: %v", err))
		}
		if found {
			var ok bool
			route, ok = ri.(*proxy.Route)
			if !ok {
				return "", sdk.NewHTTPError(http.StatusInternalServerError,
					fmt.Sprintf("invalid route type: %T", ri))
			}
			break
		}
	}
	if route == nil {
		// not found = 404
		return "", sdk.NewHTTPError(http.StatusNotFound,
			fmt.Sprintf("route not found for host %q and path %q", host, path))
	}

	target := strings.TrimSuffix(route.Target, "/") + path
	if rawQuery != "" {
		target += "?" + rawQuery
	}

	return target, nil
}
