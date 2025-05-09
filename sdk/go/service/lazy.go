/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package service

import (
	"net/http"
	"strings"
	"sync"

	"github.com/1backend/1backend/sdk/go/endpoint"
)

var (
	muMap   = map[string]*sync.Mutex{}
	muGuard sync.Mutex // protects muMap
)

type LazyOptions struct {
	// SkipLock turns off mutex locking.
	// In some scenarios, for example when a service endpoint calls itself on the same
	// node, can cause a deadlock.
	SkipLock bool
}

// LazyOption modifies LazyOptions.
type LazyOption func(*LazyOptions)

// WithSkipLock returns a LazyOption that sets SkipLock to true.
func WithSkipLock() LazyOption {
	return func(lo *LazyOptions) {
		lo.SkipLock = true
	}
}

// Lazy is a wrapper for endpoints that delays service setup until it’s actually needed.
//
// Some services dependencies don't need to be loaded right away, only when an endpoint is called.
// Delaying the setup can help reduce startup time and resource usage.
//
// If starting the service fails, we return a 500 Internal Server Error.
func Lazy(
	s LazyStarter,
	next http.HandlerFunc,
	opts ...LazyOption,
) http.HandlerFunc {
	var options LazyOptions
	for _, opt := range opts {
		opt(&options)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			// We don't want to throw errors on OPTIONS requests
			// because the response won't be visible in the chrome dev tools.
			next(w, r)
			return
		}

		firstSegment := getFirstSegment(r.URL.Path)
		mutex := getMutexForPathSegment(firstSegment)

		if !options.SkipLock {
			mutex.Lock()
			defer mutex.Unlock()
		}

		if err := s.LazyStart(); err != nil {
			endpoint.WriteErr(w, http.StatusInternalServerError, err)
			return
		}

		next(w, r)
	}
}

// LazyStarter is an interface for things that need to be started before use.
// The LazyStart() function should prepare anything that needs to be ready,
// like registering themselves in the registry and setting up permissions.
//
// It's meant to be safe to call multiple times (idempotent), so calling
// LazyStart() again won't cause issues or repeat the setup unnecessarily.
//
// This is useful for services that don’t need to start background tasks right away,
// but still require some setup before being used.
type LazyStarter interface {
	LazyStart() error
}

func getMutexForPathSegment(segment string) *sync.Mutex {
	muGuard.Lock()
	defer muGuard.Unlock()

	if m, exists := muMap[segment]; exists {
		return m
	}
	m := &sync.Mutex{}
	muMap[segment] = m
	return m
}

func getFirstSegment(path string) string {
	path = strings.TrimPrefix(path, "/")
	segments := strings.SplitN(path, "/", 2)
	if len(segments) > 0 {
		return segments[0]
	}
	return ""
}
