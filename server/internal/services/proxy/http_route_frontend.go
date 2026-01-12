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
	"time"

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

	isCacheableRequest := cs.isCacheableRequest(r)

	cacheK := cacheKey(r)

	if isCacheableRequest {
		if v, ok := cs.edgeCache.Get(cacheK); ok {
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

	if isCacheableRequest {
		if isCacheableResponse, ttl := cs.isCacheableResponse(rr); isCacheableResponse {
			cs.edgeCache.SetWithTTL(
				cacheK,
				&cachedResponse{
					status: rr.status,
					header: rr.header.Clone(),
					body:   append([]byte(nil), rr.body...),
				},
				int64(len(rr.body)),
				ttl,
			)
		}
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

func (cs *ProxyService) isCacheableRequest(req *http.Request) bool {
	if req.Method != http.MethodGet {
		return false
	}

	if req.Header.Get("Authorization") != "" || hasAuthCookie(req) {
		return false
	}

	return true
}

func (cs *ProxyService) isCacheableResponse(rr *responseRecorder) (bool, time.Duration) {
	if rr.status != http.StatusOK {
		return false, 0
	}

	// Respect Backend 'Opt-Out'

	// rr.header.Get() handles the KEY case-insensitivity.
	// strings.ToLower() handles the VALUE case-insensitivity.
	cc := strings.ToLower(rr.header.Get("Cache-Control"))

	if strings.Contains(cc, "no-store") ||
		strings.Contains(cc, "no-cache") ||
		strings.Contains(cc, "private") ||
		strings.ToLower(rr.header.Get("Pragma")) == "no-cache" {
		return false, 0
	}

	if len(rr.body) > cs.maxCachedFileSize {
		return false, 0
	}
	if rr.header.Get("Set-Cookie") != "" {
		return false, 0
	}

	ct := rr.header.Get("Content-Type")

	// Handle HTML (Short-lived for Guests).
	// Note: This provides a "Public Cache." Anonymous users will see the same
	// cached version of dynamic pages for up to 5 minutes. This should be
	// disabled via flag if the site requires real-time accuracy for guests.
	if strings.Contains(ct, "text/html") {
		return true, 5 * time.Minute
	}

	// Handle Static Assets (Long-lived)
	switch {
	case strings.HasPrefix(ct, "text/"),
		strings.Contains(ct, "javascript"),
		strings.Contains(ct, "css"),
		strings.Contains(ct, "image/"),
		strings.Contains(ct, "font/"):
		return true, 24 * time.Hour
	}

	// Fallback for everything else (JSON APIs, PDFs, etc.)
	return false, 0
}

// hasAuthCookie checks if the request contains any common authentication or session cookies.
func hasAuthCookie(req *http.Request) bool {
	// Use a map for O(1) lookups
	staticAuthCookies := map[string]struct{}{
		"access_token": {}, "token": {}, "the_token": {},
		"session_id": {}, "session": {}, "id_token": {},
		"refresh_token": {}, "connect.sid": {}, "_session_id": {},
		"sessionid": {}, "PHPSESSID": {}, "JSESSIONID": {},
		".AspNetCore.Cookies": {}, "Cookies": {},
	}

	// Iterate through the cookies actually sent by the browser
	for _, c := range req.Cookies() {
		if _, found := staticAuthCookies[c.Name]; found {
			return true
		}
	}
	return false
}

func cacheKey(r *http.Request) string {
	// Determine the Scheme accurately.
	// If behind a load balancer, r.URL.Scheme is often empty.
	scheme := r.URL.Scheme
	if scheme == "" {
		scheme = r.Header.Get("X-Forwarded-Proto")
	}
	if scheme == "" {
		scheme = "http" // Fallback
	}

	// Canonicalize Query Parameters.
	// URL.Query() parses params into a map; Encode() returns them sorted alphabetically.
	// This ensures ?b=2&a=1 and ?a=1&b=2 generate the same key.
	params := r.URL.Query()
	sortedQuery := params.Encode()

	normalizedPath := r.URL.Path
	if sortedQuery != "" {
		normalizedPath += "?" + sortedQuery
	}

	// Incorporate Content Negotiation Headers.
	// Accept-Language: prevents serving the wrong translation.
	// Accept-Encoding: prevents serving Gzip/Brotli to clients that can't decode it.
	lang := r.Header.Get("Accept-Language")
	encoding := r.Header.Get("Accept-Encoding")

	// Using a separator like '|' helps prevent "key collision" attacks where
	// parts of a path might blend into headers.
	return fmt.Sprintf("%s://%s%s|lang:%s|enc:%s",
		scheme,
		r.Host,
		normalizedPath,
		lang,
		encoding,
	)
}
