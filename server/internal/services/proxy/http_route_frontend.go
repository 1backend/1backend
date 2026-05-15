/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package proxyservice

import (
	"bufio"
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/logger"

	proxy "github.com/1backend/1backend/server/internal/services/proxy/types"
)

type cachedResponse struct {
	status    int
	header    http.Header
	body      []byte
	createdAt time.Time
}

func (cs *ProxyService) RouteFrontend(w http.ResponseWriter, r *http.Request) {
	// logger.Debug("Edge proxy request",
	// 	slog.String("host", r.Host),
	// 	slog.String("requestURI", r.RequestURI),
	// 	slog.String("path", r.URL.Path),
	// 	slog.String("rawPath", r.URL.RawPath),
	// 	slog.String("escapedPath", r.URL.EscapedPath()),
	// )

	if isMetricsEndpoint(r.URL.Path) || isMetricsEndpoint(r.URL.EscapedPath()) {
		http.NotFound(w, r)
		return
	}

	if redirect, matchedPathPrefix, found, err := cs.findRedirect(r.Host, r.URL.EscapedPath()); err != nil {
		logger.Error("Error finding redirect",
			slog.String("host", r.Host),
			slog.String("path", r.URL.EscapedPath()),
			slog.Any("error", err),
		)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	} else if found {
		statusCode, err := normalizeRedirectStatusCode(redirect.StatusCode)
		if err != nil {
			logger.Error("Invalid redirect status code",
				slog.String("host", r.Host),
				slog.String("path", r.URL.EscapedPath()),
				slog.Int("statusCode", redirect.StatusCode),
				slog.Any("error", err),
			)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		location := buildRedirectLocation(redirect.Target, matchedPathPrefix, r.URL.EscapedPath(), r.URL.RawQuery)
		if isSelfRedirect(r, location) {
			logger.Warn("Skipping self redirect",
				slog.String("host", r.Host),
				slog.String("path", r.URL.EscapedPath()),
				slog.String("location", location),
			)
		} else {
			w.Header().Set("Location", location)
			w.WriteHeader(statusCode)
			return
		}
	}

	targetString, err := cs.findRouteTarget(r.Host, r.URL.EscapedPath(), r.URL.RawQuery)
	if err != nil {
		if herr, ok := err.(*sdk.HTTPError); ok {
			http.Error(w, herr.Msg, herr.Code)
			return
		}
		logger.Error("Error finding route target",
			slog.String("host", r.Host),
			slog.String("path", r.URL.EscapedPath()),
			slog.Any("error", err),
		)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	r.Header.Del("Accept-Encoding")

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

			w.Header().Set("Vary", "Accept-Encoding, Accept-Language")
			w.Header().Set("X-Content-Type-Options", "nosniff")
			age := time.Since(cr.createdAt).Seconds()
			w.Header().Set("Age", fmt.Sprintf("%.0f", age)) // Age is in seconds
			w.Header().Set("X-Cache", "HIT")

			w.WriteHeader(cr.status)
			w.Write(cr.body)
			return
		}
	}

	ctx := context.WithValue(r.Context(), targetURLKey, targetUrl)

	rr := &responseRecorder{
		w:      w,
		header: w.Header(),
		limit:  cs.maxCachedFileSize,
	}
	cs.reverseProxy.ServeHTTP(rr, r.WithContext(ctx))

	if rr.status == 0 {
		rr.status = http.StatusOK
	}

	if isCacheableRequest {
		if isCacheableResponse, ttl := cs.isCacheableResponse(rr); isCacheableResponse {
			cachedHeader := rr.header.Clone()
			cachedHeader.Del("Content-Encoding")
			cachedHeader.Del("Content-Length")

			cs.edgeCache.SetWithTTL(
				cacheK,
				&cachedResponse{
					status:    rr.status,
					header:    cachedHeader,
					body:      append([]byte(nil), rr.body...),
					createdAt: time.Now(),
				},
				int64(len(rr.body)),
				ttl,
			)
		}
	}
}

func (cs *ProxyService) findRouteTarget(host, path, rawQuery string) (string, error) {
	snapshot, err := cs.cachedRouteSnapshot()
	if err != nil {
		return "", sdk.NewHTTPError(
			http.StatusInternalServerError,
			fmt.Sprintf("failed to query routes: %v", err),
		)
	}

	// Pick longest match (candidates is already longest to shortest).
	var route *proxy.Route
	for _, candidate := range routeCandidates(host, path) {
		key := candidate.key
		if r, ok := snapshot.route(key); ok {
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

func (cs *ProxyService) findRedirect(host, path string) (*proxy.Redirect, string, bool, error) {
	snapshot, err := cs.cachedRedirectSnapshot()
	if err != nil {
		return nil, "", false, err
	}

	for _, candidate := range routeCandidates(host, path) {
		redirect, ok := snapshot.redirect(candidate.key)
		if ok {
			return redirect, candidate.pathPrefix, true, nil
		}
	}
	return nil, "", false, nil
}

type routeCandidate struct {
	key        string
	pathPrefix string
}

func routeCandidates(host, path string) []routeCandidate {
	candidates := make([]routeCandidate, 0, strings.Count(path, "/")+1)

	p := path
	for {
		candidates = append(candidates, routeCandidate{
			key:        host + p,
			pathPrefix: p,
		})

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

	return candidates
}

func buildRedirectLocation(target, matchedPathPrefix, requestPath, rawQuery string) string {
	base, targetQuery, fragment := splitLocation(target)
	suffix := strings.TrimPrefix(requestPath, matchedPathPrefix)

	if suffix != "" {
		switch {
		case strings.HasSuffix(base, "/") && strings.HasPrefix(suffix, "/"):
			base = strings.TrimRight(base, "/") + suffix
		case !strings.HasSuffix(base, "/") && !strings.HasPrefix(suffix, "/"):
			base += "/" + suffix
		default:
			base += suffix
		}
	}

	query := targetQuery
	if rawQuery != "" {
		if query != "" {
			query += "&" + rawQuery
		} else {
			query = rawQuery
		}
	}
	if query != "" {
		base += "?" + query
	}
	if fragment != "" {
		base += "#" + fragment
	}
	return base
}

func splitLocation(location string) (base string, query string, fragment string) {
	base = location
	if idx := strings.IndexByte(base, '#'); idx >= 0 {
		fragment = base[idx+1:]
		base = base[:idx]
	}
	if idx := strings.IndexByte(base, '?'); idx >= 0 {
		query = base[idx+1:]
		base = base[:idx]
	}
	return base, query, fragment
}

func isSelfRedirect(r *http.Request, location string) bool {
	requestURL, err := requestURLForRedirectComparison(r)
	if err != nil {
		return false
	}

	redirectURL, err := url.Parse(location)
	if err != nil {
		return false
	}
	redirectURL = requestURL.ResolveReference(redirectURL)

	return strings.EqualFold(redirectURL.Scheme, requestURL.Scheme) &&
		equivalentRedirectHost(redirectURL, requestURL) &&
		redirectURL.EscapedPath() == requestURL.EscapedPath() &&
		redirectURL.RawQuery == requestURL.RawQuery
}

func equivalentRedirectHost(a, b *url.URL) bool {
	return strings.EqualFold(a.Hostname(), b.Hostname()) &&
		normalizedURLPort(a.Scheme, a.Port()) == normalizedURLPort(b.Scheme, b.Port())
}

func normalizedURLPort(scheme, port string) string {
	if port != "" {
		return port
	}
	switch strings.ToLower(scheme) {
	case "http":
		return "80"
	case "https":
		return "443"
	default:
		return ""
	}
}

func requestURLForRedirectComparison(r *http.Request) (*url.URL, error) {
	path := r.URL.EscapedPath()
	if path == "" {
		path = "/"
	}

	rawURL := effectiveRequestScheme(r) + "://" + r.Host + path
	if r.URL.RawQuery != "" {
		rawURL += "?" + r.URL.RawQuery
	}

	return url.Parse(rawURL)
}

func effectiveRequestScheme(r *http.Request) string {
	if proto := r.Header.Get("X-Forwarded-Proto"); proto != "" {
		proto = strings.TrimSpace(strings.Split(proto, ",")[0])
		if proto != "" {
			return proto
		}
	}
	if r.TLS != nil {
		return "https"
	}
	return "http"
}

func isMetricsEndpoint(path string) bool {
	path = strings.TrimRight(path, "/")
	if path == "" {
		return false
	}

	lastSegmentIdx := strings.LastIndexByte(path, '/')
	if lastSegmentIdx >= 0 {
		path = path[lastSegmentIdx+1:]
	}

	return strings.EqualFold(path, "metrics")
}

type responseRecorder struct {
	w          http.ResponseWriter
	status     int
	header     http.Header
	body       []byte
	limit      int // Max size to buffer
	overflowed bool
}

func (r *responseRecorder) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	// Corrected: accessing 'r.w' instead of 'r.ResponseWriter'
	h, ok := r.w.(http.Hijacker)
	if !ok {
		return nil, nil, fmt.Errorf("underlying response writer does not support Hijacker")
	}
	return h.Hijack()
}

// Flush allows the recorder to satisfy the http.Flusher interface,
// ensuring the outer compression middleware can stream data correctly.
func (r *responseRecorder) Flush() {
	if f, ok := r.w.(http.Flusher); ok {
		f.Flush()
	}
}

func (r *responseRecorder) Header() http.Header {
	return r.w.Header()
}

func (r *responseRecorder) WriteHeader(code int) {
	if r.status != 0 {
		return // Prevent double-header writing
	}
	r.status = code

	// Inject your proxy-specific headers HERE,
	// just before the real writer sends them out.
	h := r.w.Header()

	// If we are streaming/teeing, we want the proxy to
	// let the outer server handle compression.
	h.Del("Content-Encoding")

	h.Set("Vary", "Accept-Encoding, Accept-Language")
	h.Set("X-Content-Type-Options", "nosniff")
	h.Set("X-Cache", "MISS")

	r.w.WriteHeader(code)
}

func (r *responseRecorder) Write(b []byte) (int, error) {
	// 1. Always send data to the client immediately
	n, err := r.w.Write(b)

	// 2. Only keep in memory if we are under the limit
	if !r.overflowed {
		if len(r.body)+len(b) > r.limit {
			r.overflowed = true
			r.body = nil // Drop the buffer to save memory
		} else {
			r.body = append(r.body, b...)
		}
	}

	return n, err
}

func (r *responseRecorder) Unwrap() http.ResponseWriter {
	return r.w
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
	if rr.status != http.StatusOK || rr.overflowed {
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

	normalizedPath := r.URL.EscapedPath()
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
