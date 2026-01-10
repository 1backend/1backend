/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package proxyservice

import (
	"context"
	"io"
	"log/slog"
	"math/rand/v2"
	"net/http"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/pkg/errors"

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/logger"
)

var copyBufPool = sync.Pool{
	New: func() any {
		return make([]byte, 32*1024) // 32KB standard buffer
	},
}

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

	serviceSlug := getServiceSlug(r.URL.Path)

	var instances []openapi.RegistrySvcInstance
	val, ok := cs.instanceCache.Load(serviceSlug)
	entry, _ := val.(cacheEntry)

	if ok && time.Now().Before(entry.expiry) {
		instances = entry.instances
	} else {
		// 2. Cache expired or missing -> Use Singleflight
		// This ensures only ONE call to RegistrySvcAPI happens per slug
		res, err, _ := cs.backendSf.Do(serviceSlug, func() (any, error) {
			rsp, _, err := cs.options.ClientFactory.Client(client.WithToken(cs.token)).
				RegistrySvcAPI.ListInstances(context.Background()).
				Slug(serviceSlug).
				Execute()

			if err != nil {
				// If API fails but we have stale data, return stale data as fallback
				if ok {
					return entry.instances, nil
				}
				return nil, err
			}

			// Pre-filter healthy instances here to save CPU on every request
			healthy := make([]openapi.RegistrySvcInstance, 0, len(rsp.Instances))
			for _, ins := range rsp.Instances {
				if ins.Status == openapi.InstanceStatusHealthy {
					healthy = append(healthy, ins)
				}
			}

			// If no healthy ones, use all as fallback
			if len(healthy) == 0 {
				healthy = rsp.Instances
			}

			cs.instanceCache.Store(serviceSlug, cacheEntry{
				instances: healthy,
				expiry:    time.Now().Add(10 * time.Second),
			})
			return healthy, nil
		})

		if err != nil {
			return http.StatusInternalServerError, errors.Wrap(err, "registry unavailable")
		}
		instances = res.([]openapi.RegistrySvcInstance)
	}

	if len(instances) == 0 {
		return http.StatusNotFound, errors.New("no instances found")
	}

	instance := instances[rand.IntN(len(instances))]

	var sb strings.Builder
	sb.WriteString(strings.TrimSuffix(instance.Url, "/"))
	if !strings.HasPrefix(r.URL.Path, "/") {
		sb.WriteByte('/')
	}
	sb.WriteString(r.URL.Path)
	if r.URL.RawQuery != "" {
		sb.WriteByte('?')
		sb.WriteString(r.URL.RawQuery)
	}
	uri := sb.String()

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

	wh := w.Header()
	for k, v := range resp.Header {
		// Skip hop-by-hop and content-length headers
		if k == "Content-Length" || k == "Transfer-Encoding" || k == "Connection" {
			continue
		}
		wh[k] = v // Direct map assignment is much faster than Add()
	}

	// logger.Debug("Service proxy request returned",
	// 	slog.String("path", r.URL.Path),
	// 	slog.String("method", r.Method),
	// 	slog.Int("statusCode", resp.StatusCode),
	// )

	w.WriteHeader(resp.StatusCode)

	buf := copyBufPool.Get().([]byte)

	defer copyBufPool.Put(buf)

	// io.CopyBuffer is exactly like io.Copy, but it uses the memory we provide
	// instead of allocating its own 32KB slice internally.
	_, err = io.CopyBuffer(w, resp.Body, buf)

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
func getServiceSlug(path string) string {
	if path == "" || path == "/" {
		return ""
	}
	path = strings.TrimPrefix(path, "/")
	if i := strings.Index(path, "/"); i != -1 {
		return path[:i]
	}
	return path
}
