/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package proxyservice

import (
	"io"
	"log/slog"
	"math/rand/v2"
	"net/http"
	"strings"
	"syscall"
	"time"

	"github.com/pkg/errors"

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/logger"
)

type cacheEntry struct {
	instances []openapi.RegistrySvcInstance
	expiry    time.Time
}

// RouteBackend routes requests that look just like the builtin 1Backend service paths:
// - /my-svc/my-endpoint -> my-svc
// - /my-svc/my-other-endpoint?query=param -> my-svc
// It uses the service slug (the first part of the path) to find the service instance.
// It then proxies the request to an instance of that service.
func (cs *ProxyService) RouteBackend(w http.ResponseWriter, r *http.Request) {
	statusCode, err := cs.routeBackend(w, r)

	if err != nil {
		logger.Error("Error service proxying",
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
			slog.String("error", err.Error()),
		)
		if r.Method == http.MethodOptions {
			// We don't want to write errors to response on OPTIONS requests
			// because the response won't be visible in the chrome dev tools.
			// We log it instead.

			return
		}

		w.WriteHeader(statusCode)
		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			logger.Error("Error writing service proxy response",
				slog.String("error", err.Error()),
			)
		}
		return
	}
}

func (cs *ProxyService) routeBackend(w http.ResponseWriter, r *http.Request) (int, error) {
	// logger.Debug("Service proxying",
	// 	slog.String("path", r.URL.Path),
	// 	slog.String("method", r.Method),
	// )

	// The proxy service's LazyStart() must be called here because OPTIONS requests
	// are not handled the standard Lazy logic. Unlike other services, the Proxy does handle
	// OPTIONS requests and requires initialization (including token acquisition) to do so.
	err := cs.LazyStart()
	if err != nil && r.Method != http.MethodOptions {
		return http.StatusInternalServerError, errors.Wrap(err, "error starting proxy service")
	}

	serviceSlug := getServiceSlug(r)

	var cachedEntry *cacheEntry
	if val, ok := cs.instanceCache.Load(serviceSlug); ok {
		e := val.(cacheEntry)
		cachedEntry = &e
	}

	var instances []openapi.RegistrySvcInstance
	shouldFetch := cachedEntry == nil || time.Now().After(cachedEntry.expiry)

	if shouldFetch {
		// Attempt refresh
		rsp, _, err := cs.options.ClientFactory.Client(client.WithToken(cs.token)).
			RegistrySvcAPI.ListInstances(r.Context()).
			Slug(serviceSlug).
			Execute()

		if err != nil {
			// --- CRITICAL FALLBACK ---
			if cachedEntry != nil {
				logger.Warn("Registry lookup failed, using stale cache",
					slog.String("service", serviceSlug),
					slog.String("error", err.Error()),
				)
				instances = cachedEntry.instances
			} else {
				// No cache and API failed: we must error out
				return http.StatusInternalServerError, errors.Wrap(err, "registry unavailable and no cache exists")
			}
		} else {
			// API Success: Update cache
			instances = rsp.Instances
			cs.instanceCache.Store(serviceSlug, cacheEntry{
				instances: instances,
				expiry:    time.Now().Add(10 * time.Second), // 10s TTL
			})
		}
	} else {
		// Cache is still fresh
		instances = cachedEntry.instances
	}

	if len(instances) == 0 {
		return http.StatusNotFound, errors.New("no instances found")
	}

	// Prioritize healthy instances
	selectedInstances := make([]openapi.RegistrySvcInstance, 0, len(instances))
	for _, instance := range instances {
		if instance.Status == openapi.InstanceStatusHealthy {
			selectedInstances = append(selectedInstances, instance)
		}
	}

	// But fall back to any instances if none found
	if len(selectedInstances) == 0 {
		selectedInstances = instances
	}

	randomIndex := rand.IntN(len(selectedInstances))
	instance := selectedInstances[randomIndex]

	uri := strings.TrimSuffix(instance.Url, "/") +
		"/" +
		strings.TrimLeft(r.URL.Path, "/")
	if r.URL.RawQuery != "" {
		uri += "?" + r.URL.RawQuery
	}

	req, err := http.NewRequestWithContext(r.Context(), r.Method, uri, r.Body)
	if err != nil {
		return http.StatusInternalServerError, errors.Wrap(err, "error creating proxy request")
	}
	req.Header = r.Header

	resp, err := cs.httpClient.Do(req)
	if err != nil {
		return http.StatusInternalServerError, errors.Wrap(err, "error proxying request")
	}
	defer resp.Body.Close()

	for k, v := range resp.Header {
		// Skip "Content-Length" and "Transfer-Encoding" to prevent HTTP response errors.
		// - Go automatically sets "Content-Length" based on the actual body size.
		// - "Transfer-Encoding: chunked" conflicts with "Content-Length", so we avoid copying it.
		if k == "Content-Length" || k == "Transfer-Encoding" {
			continue
		}
		for _, vv := range v {
			w.Header().Add(k, vv)
		}
	}

	// logger.Debug("Service proxy request returned",
	// 	slog.String("path", r.URL.Path),
	// 	slog.String("method", r.Method),
	// 	slog.Int("statusCode", resp.StatusCode),
	// )

	w.WriteHeader(resp.StatusCode)

	_, err = io.Copy(w, resp.Body)
	if err != nil && !errors.Is(err, syscall.EPIPE) && !strings.Contains(err.Error(), "stream closed") {
		return http.StatusInternalServerError, errors.Wrap(err, "error copying response body")
	}

	if flusher, ok := w.(http.Flusher); ok {
		flusher.Flush()
	}

	return http.StatusOK, nil
}

// gets service slug from http request path
// eg. /my-svc/my-endpoint -> my-svc
func getServiceSlug(r *http.Request) string {
	cleanedPath := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(cleanedPath, "/")
	if len(parts) > 0 && parts[0] != "" {
		return parts[0]
	}
	return ""
}
