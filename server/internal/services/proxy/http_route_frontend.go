/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package proxyservice

import (
	"fmt"
	"io"
	"log/slog"
	"net"
	"net/http"
	"strings"

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/logger"

	proxy "github.com/1backend/1backend/server/internal/services/proxy/types"
)

func (cs *ProxyService) RouteFrontend(w http.ResponseWriter, r *http.Request) {
	logger.Debug("Edge proxying",
		slog.String("host", r.Host),
		slog.String("path", r.URL.Path),
	)

	// Try to find a matching route for this host
	routeI, found, err := cs.routeStore.Query(
		datastore.Id(r.Host),
	).FindOne()
	if err != nil {
		logger.Error("Error querying route store", slog.String("error", err.Error()))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if !found {
		http.Error(w, "Route not found", http.StatusNotFound)
		return
	}

	route, ok := routeI.(*proxy.Route)
	if !ok {
		logger.Error("Invalid route type", slog.String("type", fmt.Sprintf("%T", routeI)))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	target := strings.TrimSuffix(route.Target, "/") + r.URL.Path
	if r.URL.RawQuery != "" {
		target += "?" + r.URL.RawQuery
	}

	req, err := http.NewRequest(r.Method, target, r.Body)
	if err != nil {
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
		http.Error(w, "Failed to proxy request", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	for k, v := range resp.Header {
		if k == "Content-Length" || k == "Transfer-Encoding" {
			continue
		}
		for _, vv := range v {
			w.Header().Add(k, vv)
		}
	}

	w.WriteHeader(resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Error reading response body", slog.String("error", err.Error()))
		return
	}
	_, err = w.Write(body)
	if err != nil {
		logger.Error("Error writing response body", slog.String("error", err.Error()))
		return
	}

	if flusher, ok := w.(http.Flusher); ok {
		flusher.Flush()
	}
}
