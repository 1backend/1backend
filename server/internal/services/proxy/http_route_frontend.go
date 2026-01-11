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
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"strings"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/logger"

	proxy "github.com/1backend/1backend/server/internal/services/proxy/types"
)

type cachedResponse struct {
	status int
	header http.Header
	body   []byte
}

func cacheKey(r *http.Request) string {
	return r.Host + r.URL.RequestURI()
}

func (cs *ProxyService) RouteFrontend(w http.ResponseWriter, r *http.Request) {
	// logger.Debug("Edge proxying",
	// 	slog.String("host", r.Host),
	// 	slog.String("path", r.URL.Path),
	// )

	targetString, err := cs.findRouteTarget(r.Host, r.URL.Path, r.URL.RawQuery)
	if err != nil {
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

	targetUrl, err := url.Parse(targetString)
	if err != nil {
		logger.Error("Failed to parse target URL", slog.String("target", targetString), slog.Any("error", err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodGet {
		if v, ok := cs.fileCache.Get(cacheKey(r)); ok {
			cr := v.(*cachedResponse)
			for k, vv := range cr.header {
				w.Header()[k] = vv
			}
			w.WriteHeader(cr.status)
			w.Write(cr.body)
			return
		}
	}

	ctx := context.WithValue(r.Context(), targetURLKey, targetUrl)

	rr := &responseRecorder{}
	cs.reverseProxy.ServeHTTP(rr, r.WithContext(ctx))

	for k, vv := range rr.header {
		w.Header()[k] = vv
	}
	if rr.status == 0 {
		rr.status = http.StatusOK
	}
	w.WriteHeader(rr.status)
	w.Write(rr.body)

	if cs.isCacheable(r, rr) {
		cs.fileCache.Set(
			cacheKey(r),
			&cachedResponse{
				status: rr.status,
				header: rr.header.Clone(),
				body:   append([]byte(nil), rr.body...),
			},
			int64(len(rr.body)),
		)
	}
}

func (cs *ProxyService) findRouteTarget(host, path, rawQuery string) (string, error) {

	candidates := make([]string, 0, strings.Count(path, "/")+1)

	p := path
	for {
		candidates = append(candidates, host+p)

		if len(p) == 0 {
			break
		}

		idx := strings.LastIndexByte(p, '/')

		if idx < 0 {
			// Only happens if path is "segment1" (no slash).
			// Next step is root (empty string).
			p = ""
		} else if idx == 0 {
			// We are at "/segment1". The slash is at 0.
			// The substring up to 0 is empty. Next step is root.
			p = ""
		} else {
			// We are at "/segment1/segment2". Slice before the slash.
			p = p[:idx]
		}
	}

	routes := map[string]*proxy.Route{}
	var missing []any

	// 1. Check cache first (including negative cache)
	for _, key := range candidates {
		if v, ok := cs.routeCache.Load(key); ok {
			if v == nil {
				// cached miss
				continue
			}
			routes[key] = v.(*proxy.Route)
		} else {
			missing = append(missing, key)
		}
	}

	// 2. Fetch all missing from DB in one query
	if len(missing) > 0 {
		logger.Debug("Cache miss for routes", slog.Any("missing", missing))

		sfKey := fmt.Sprintf("%s|%s", host, path)
		v, err, _ := cs.sf.Do(sfKey, func() (interface{}, error) {

			ri, err := cs.routeStore.Query(
				datastore.IsInList(datastore.Field("id"), missing...),
			).Find()
			if err != nil {
				return nil, sdk.NewHTTPError(
					http.StatusInternalServerError,
					fmt.Sprintf("failed to query route: %v", err),
				)
			}

			return ri, nil
		})

		if err != nil {
			return "", err
		}

		ri := v.([]datastore.Row)

		foundMap := map[string]*proxy.Route{}

		for _, r := range ri {
			route := r.(*proxy.Route)
			foundMap[route.Id] = route
			cs.routeCache.Store(route.Id, route)
			routes[route.Id] = route
		}

		// 3. Negative cache the rest
		for _, k := range missing {
			key := k.(string)
			if _, ok := foundMap[key]; !ok {
				cs.routeCache.Store(key, nil)
			}
		}
	}

	// 4. Pick longest match (candidates is already longest → shortest)
	var route *proxy.Route
	for _, key := range candidates {
		if r, ok := routes[key]; ok {
			route = r
			break
		}
	}

	if route == nil {
		return "", sdk.NewHTTPError(
			http.StatusNotFound,
			fmt.Sprintf("route not found for host %q and path %q", host, path),
		)
	}

	target := strings.TrimSuffix(route.Target, "/") + path
	if rawQuery != "" {
		target += "?" + rawQuery
	}

	return target, nil
}

type responseRecorder struct {
	status int
	header http.Header
	body   []byte
}

func (r *responseRecorder) Header() http.Header {
	if r.header == nil {
		r.header = make(http.Header)
	}
	return r.header
}

func (r *responseRecorder) WriteHeader(code int) {
	r.status = code
}

func (r *responseRecorder) Write(b []byte) (int, error) {
	r.body = append(r.body, b...)
	return len(b), nil
}

func (cs *ProxyService) isCacheable(req *http.Request, rr *responseRecorder) bool {
	if req.Method != http.MethodGet {
		return false
	}
	if rr.status != http.StatusOK {
		return false
	}
	if len(rr.body) > cs.maxCachedFileSize {
		return false
	}
	if rr.header.Get("Set-Cookie") != "" {
		return false
	}

	ct := rr.header.Get("Content-Type")
	switch {
	case strings.HasPrefix(ct, "text/"):
	case strings.Contains(ct, "javascript"):
	case strings.Contains(ct, "json"):
	case strings.Contains(ct, "css"):
	case strings.Contains(ct, "image/"):
	case strings.Contains(ct, "font/"):
	default:
		return false
	}

	return true
}
