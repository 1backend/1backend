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
	"sync"
	"syscall"
	"time"

	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"

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
			slog.String("host", r.Host),
			slog.String("path", r.URL.Path),
			slog.String("rawPath", r.URL.RawPath),
			slog.String("escapedPath", r.URL.EscapedPath()),
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

func (cs *ProxyService) routeBackend(w http.ResponseWriter, r *http.Request) (statusCode int, err error) {
	// logger.Debug("Service proxying",
	// 	slog.String("path", r.URL.Path),
	// 	slog.String("method", r.Method),
	// )

	if isHealthProbePath(r.URL.Path) || isHealthProbePath(r.URL.EscapedPath()) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"ok","checks":{"http":"ok"}}`))
		return http.StatusOK, nil
	}

	serviceSlug := getServiceSlug(r.URL.EscapedPath())
	ctx, span := otel.Tracer("github.com/1backend/1backend/server/internal/services/proxy").Start(
		r.Context(),
		"proxy.route_backend",
	)
	defer func() {
		if statusCode == 0 {
			statusCode = http.StatusOK
		}
		span.SetAttributes(
			attribute.String("onebackend.proxy.service", "1backend"),
			attribute.String("onebackend.target_service", serviceSlug),
			attribute.String("http.request.method", r.Method),
			attribute.String("url.path", r.URL.EscapedPath()),
			attribute.Int("http.response.status_code", statusCode),
		)
		if err != nil || statusCode >= http.StatusInternalServerError {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			} else {
				span.SetStatus(codes.Error, http.StatusText(statusCode))
			}
		}
		span.End()
	}()

	// The proxy service's LazyStart() must be called here because OPTIONS requests
	// are not handled the standard Lazy logic. Unlike other services, the Proxy does handle
	// OPTIONS requests and requires initialization (including token acquisition) to do so.
	err = cs.LazyStart()
	if err != nil && r.Method != http.MethodOptions {
		return http.StatusInternalServerError, errors.Wrap(err, "error starting proxy service")
	}

	var instances []openapi.RegistrySvcInstance
	val, ok := cs.instanceCache.Load(serviceSlug)
	entry, _ := val.(cacheEntry)

	if ok && time.Now().Before(entry.expiry) {
		span.SetAttributes(attribute.String("onebackend.proxy.instance_cache", "hit"))
		instances = entry.instances
	} else {
		span.SetAttributes(attribute.String("onebackend.proxy.instance_cache", "miss"))
		// 2. Cache expired or missing -> Use Singleflight
		// This ensures only ONE call to RegistrySvcAPI happens per slug
		res, err, _ := cs.backendSf.Do(serviceSlug, func() (any, error) {
			rsp, _, err := cs.options.ClientFactory.Client(client.WithToken(cs.token)).
				RegistrySvcAPI.ListInstances(ctx).
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

			if len(healthy) > 0 {
				cs.instanceCache.Store(serviceSlug, cacheEntry{
					instances: healthy,
					expiry:    time.Now().Add(10 * time.Second),
				})
			}

			return healthy, nil
		})

		if err != nil {
			return http.StatusInternalServerError, errors.Wrap(err, "registry unavailable")
		}
		instances = res.([]openapi.RegistrySvcInstance)
	}

	if len(instances) == 0 {
		logger.Warn("No instances found",
			slog.String("serviceSlug", serviceSlug),
			slog.String("path", r.URL.EscapedPath()),
		)
		return http.StatusNotFound, errors.New("no instances found")
	}

	instance := instances[rand.IntN(len(instances))]

	var sb strings.Builder
	sb.WriteString(strings.TrimSuffix(instance.Url, "/"))
	if !strings.HasPrefix(r.URL.EscapedPath(), "/") {
		sb.WriteByte('/')
	}
	sb.WriteString(r.URL.EscapedPath())
	if r.URL.RawQuery != "" {
		sb.WriteByte('?')
		sb.WriteString(r.URL.RawQuery)
	}
	uri := sb.String()

	req, err := http.NewRequestWithContext(ctx, r.Method, uri, r.Body)
	if err != nil {
		return http.StatusInternalServerError, errors.Wrap(err, "error creating proxy request")
	}
	req.Header = r.Header.Clone()
	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))

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

	return resp.StatusCode, nil
}

// gets service slug from http request path
// eg. /my-svc/my-endpoint -> my-svc
func getServiceSlug(path string) string {
	// 1. Remove ALL leading slashes (e.g., ///image-svc -> image-svc)
	path = strings.TrimLeft(path, "/")
	if path == "" {
		return ""
	}

	// 2. Get the first segment only
	segments := strings.SplitN(path, "/", 2)
	return segments[0]
}

func isHealthProbePath(path string) bool {
	switch path {
	case "/healthz", "/livez", "/readyz":
		return true
	default:
		return false
	}
}
