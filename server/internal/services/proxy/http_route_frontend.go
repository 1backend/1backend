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
	"log/slog"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

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

	proxy := httputil.NewSingleHostReverseProxy(targetUrl)

	// Set the Director to fix paths (if findRouteTarget only returns the base)
	// BUT since `findRouteTarget` *already* builds the full path,
	// we need to tell the proxy to use the URL *as-is*.
	//
	// We must rewrite the proxy's Director.
	// By default, NewSingleHostReverseProxy just appends the request path,
	// but `findRouteTarget` already *includes* the path.
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		// Run the original director (which sets host, scheme, etc.)
		originalDirector(req)

		// Override the Path and RawQuery with the values from our target URL.
		// This is the key part!
		req.URL.Path = targetUrl.Path
		req.URL.RawQuery = targetUrl.RawQuery

		// Preserve original header.
		// Maybe make this optional later
		req.Host = r.Host
	}

	proxy.ServeHTTP(w, r)
}

func (cs *ProxyService) findRouteTarget(host, path, rawQuery string) (string, error) {
	if path == "" {
		path = "/"
	}

	candidates := make([]string, 0, strings.Count(path, "/")+1)

	p := path
	for {
		candidates = append(candidates, host+p)

		if p == "/" {
			break
		}

		idx := strings.LastIndexByte(p, '/')
		if idx <= 0 {
			p = "/"
		} else {
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

		ri, err := cs.routeStore.Query(
			datastore.IsInList(datastore.Field("id"), missing...),
		).Find()
		if err != nil {
			return "", sdk.NewHTTPError(
				http.StatusInternalServerError,
				fmt.Sprintf("failed to query route: %v", err),
			)
		}

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
