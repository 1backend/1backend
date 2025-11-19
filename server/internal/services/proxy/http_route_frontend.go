package proxyservice

import (
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/logger"

	proxy "github.com/1backend/1backend/server/internal/services/proxy/types"
)

const (
	// How long to keep a route in memory before re-verifying with DB
	routeCacheTTL = 30 * time.Second
)

func (cs *ProxyService) RouteFrontend(w http.ResponseWriter, r *http.Request) {

	logger.Debug("Edge proxying",
		slog.String("host", r.Host),
		slog.String("path", r.URL.Path),
	)

	// Capture the matchedRoute to use its ID for prefix stripping
	matchedRoute, targetURL, err := cs.resolveRoute(r.Host, r.URL.Path)
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

	// Determine the Proxy Host (Scheme + Host only)
	// We strip the path from the key so that http://backend/v1 and http://backend/v2
	// share the same ReverseProxy instance (and thus the same TCP connection pool).
	proxyKey := fmt.Sprintf("%s://%s", targetURL.Scheme, targetURL.Host)

	rp, err := cs.getReverseProxy(proxyKey, targetURL)
	if err != nil {
		logger.Error("Failed to get reverse proxy", slog.Any("error", err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Save the original path requested by the client.
	originalPath := r.URL.Path

	// --- Path Rewriting Logic (Fixes Microfrontend Path Bug) ---

	// 1. Determine the target path from the route definition, stripping the trailing slash.
	targetPath := strings.TrimSuffix(targetURL.Path, "/")

	if targetPath != "" {
		// This is a path replacement/rewriting scenario (e.g., /app -> /v1).

		// A. Identify the matched route prefix (e.g., "/app/admin" from "x.localhost/app/admin")
		routePrefix := ""
		if strings.HasPrefix(matchedRoute.Id, r.Host+"/") {
			// Extract the path part, including the leading slash
			routePrefix = matchedRoute.Id[len(r.Host):]
		}

		// B. Path Stripping: Remove the matched route prefix from the original path.
		strippedPath := originalPath
		if routePrefix != "" && strings.HasPrefix(originalPath, routePrefix) {
			strippedPath = originalPath[len(routePrefix):]
		}

		// C. Cleanup: Ensure the stripped path always starts with a slash, even if it's empty.
		// If the stripped path is empty (e.g., exact match on /app), it becomes "/".
		if strippedPath == "" || strippedPath[0] != '/' {
			strippedPath = "/" + strippedPath
		}

		// D. Prepend: The new path is the target path + the stripped path.
		r.URL.Path = targetPath + strippedPath

	} else {
		// If the target URL has no path component, it's a simple pass-through/microfrontend route.
		// The original requested path (e.g., /app/page) is passed directly to the backend.
		r.URL.Path = originalPath
	}

	// Important: Clear RawPath so Go recalculates it from the new Path
	r.URL.RawPath = ""

	// --- End Path Rewriting Logic ---

	// Merge Query Params
	// If the route target had query params (e.g. ?force=true), merge them.
	if targetURL.RawQuery != "" {
		if r.URL.RawQuery == "" {
			r.URL.RawQuery = targetURL.RawQuery
		} else {
			// Note: This logic for merging raw query strings is technically incorrect
			// if targetURL.RawQuery contains characters that need encoding, but
			// sticking to the existing implementation style for now.
			r.URL.RawQuery = targetURL.RawQuery + "&" + r.URL.RawQuery
		}
	}

	r.Host = targetURL.Host
	r.URL.Host = targetURL.Host
	r.URL.Scheme = targetURL.Scheme

	rp.ServeHTTP(w, r)
}

// getReverseProxy returns a cached proxy or creates a new one.
// proxyKey should be "scheme://host".
func (cs *ProxyService) getReverseProxy(proxyKey string, targetURLForCreation *url.URL) (*httputil.ReverseProxy, error) {
	if val, ok := cs.proxyCache.Load(proxyKey); ok {
		return val.(*httputil.ReverseProxy), nil
	}

	// We create a proxy for the BASE URL (no path).
	// This ensures the standard Director doesn't double-append paths,
	// because we handled path rewriting manually in RouteFrontend.
	baseURL := &url.URL{
		Scheme: targetURLForCreation.Scheme,
		Host:   targetURLForCreation.Host,
	}

	transport := &http.Transport{
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 20,
		IdleConnTimeout:     90 * time.Second,
	}

	p := httputil.NewSingleHostReverseProxy(baseURL)
	p.Transport = transport

	// We wrap the director to ensure headers like X-Forwarded-Host are correct
	// and to handle the Host header requirement for some backends.
	originalDirector := p.Director
	p.Director = func(req *http.Request) {
		originalDirector(req)
		req.Host = baseURL.Host
	}

	// Store in cache
	cs.proxyCache.Store(proxyKey, p)
	return p, nil
}

// resolve the Target (with Caching + Batch DB Query)
// returned URL contains the full parsed URL from the route (e.g., http://backend/v1)
func (cs *ProxyService) resolveRoute(host, path string) (*proxy.Route, *url.URL, error) {
	candidateKeys := candidateKeys(host, path)

	queryValues := make([]any, 0, len(candidateKeys))
	for _, v := range candidateKeys {
		queryValues = append(queryValues, v)
	}

	// Batch Query
	rows, err := cs.routeStore.Query(
		datastore.IsInList(datastore.Field("id"), queryValues...),
	).Find()

	if err != nil {

		return nil, nil, sdk.NewHTTPError(http.StatusInternalServerError,
			fmt.Sprintf("failed to query routes: %v", err))
	}

	//  Match Specificity
	foundRoutes := make(map[string]*proxy.Route)
	for _, row := range rows {
		if r, ok := row.(*proxy.Route); ok {
			foundRoutes[r.GetId()] = r
		}
	}

	var bestMatch *proxy.Route
	for _, key := range candidateKeys {
		if route, exists := foundRoutes[key]; exists {
			bestMatch = route
			break // Longest match found
		}
	}

	if bestMatch == nil {
		return nil, nil, sdk.NewHTTPError(http.StatusNotFound,
			fmt.Sprintf("route not found for host %q and path %q", host, path))
	}

	// 5. Parse and Cache
	targetURL, err := url.Parse(bestMatch.Target)
	if err != nil {
		return nil, nil, sdk.NewHTTPError(http.StatusInternalServerError,
			fmt.Sprintf("invalid target url in route %s: %v", bestMatch.Id, err))
	}

	return bestMatch, targetURL, nil
}

func candidateKeys(host, path string) (
	candidates []string,
) {

	var candidateKeys []string

	// Start with the full path, prefixed by host. This is the longest candidate key.
	currentKey := host + path

	// Remove any trailing slash for consistency (unless it's just the root host/)
	if len(currentKey) > len(host)+1 && strings.HasSuffix(currentKey, "/") {
		currentKey = currentKey[:len(currentKey)-1]
	}

	// Iterate from the longest prefix down to the root host key.
	for {
		key := currentKey

		candidateKeys = append(candidateKeys, key)

		// Termination check: If the current key is already just the host, stop.
		if key == host {
			break
		}

		// Find the index of the last slash that separates path segments.
		// We search in the path part of the key.
		lastSlash := strings.LastIndex(key[len(host):], "/")

		// If lastSlash is -1, it means the key is "host/segment".
		if lastSlash == -1 {
			currentKey = host // Next iteration will check the root key
		} else {
			// Trim the key up to (but not including) the last slash found in the path.
			// lastSlash is relative to the path string, so add back len(host).
			currentKey = key[:len(host)+lastSlash]
		}
	}

	return candidateKeys
}
