/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/

package endpoint

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/pkg/errors"

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/dgraph-io/ristretto"
)

type HasPermissionResponse struct {
	Response   *openapi.UserSvcHasPermissionResponse
	StatusCode int

	// Not caching the full http response here to avoid unnecessary memory usage.
}

type PermissionChecker interface {
	// HasPermission caches the result of the has-permission endpoint calls.
	// Uses a SHA-256 hash of the JWT and permission string as the cache key.
	// Cached values expire after 5 minutes. No invalidation on token revocation yet.
	HasPermission(
		request *http.Request,
		permission string,
	) (*openapi.UserSvcHasPermissionResponse, int, error)
}

type permissionChecker struct {
	testing         bool
	clientFactory   client.ClientFactory
	permissionCache *ristretto.Cache
}

func NewPermissionChecker(
	clientFactory client.ClientFactory,
) PermissionChecker {
	permissionCache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e5,     // number of keys to track frequency (10x max items)
		MaxCost:     1 << 20, // max cost in bytes (~1 MiB)
		BufferItems: 64,      // recommended
	})
	if err != nil {
		// ??
		panic(errors.Wrap(err, "failed to create ristretto cache").Error())
	}

	return &permissionChecker{
		clientFactory:   clientFactory,
		permissionCache: permissionCache,
	}
}

// HasPermission checks if the user has the specified permission.
// It first checks the cache for a cached response.
// If a cached response is found, it returns that.
// If not, it makes a request to the user service to check the permission.
//
// It also handles JWT expiration.
// If the JWT is expired, it refreshes the token and maps the old request to the new one.
func (pc *permissionChecker) HasPermission(
	request *http.Request,
	permission string,
) (*openapi.UserSvcHasPermissionResponse, int, error) {

	jwt := strings.TrimPrefix(request.Header.Get("Authorization"), "Bearer ")

	key := generateCacheKey(
		jwt,
		permission,
	)

	if jwt != "" {
		if value, found := pc.permissionCache.Get(key); found {
			if cachedResp, ok := value.(*HasPermissionResponse); ok {
				return cachedResp.Response, cachedResp.StatusCode, nil
			}
		}
	}

	if pc.clientFactory == nil {
		return nil, http.StatusInternalServerError, errors.New("client factory is nil")
	}

	isAuthRsp, httpResponse, err := pc.clientFactory.Client(
		client.WithTokenFromRequest(request),
	).
		UserSvcAPI.HasPermission(
		request.Context(),
		permission,
	).
		Execute()
	if err != nil {
		return nil, httpResponse.StatusCode, err
	}

	code := 0
	if httpResponse != nil {
		code = httpResponse.StatusCode
	}

	if key != "" {
		expiresAt, err := time.Parse(time.RFC3339, isAuthRsp.Until)
		if err != nil {
			return nil, http.StatusInternalServerError, errors.Wrap(err, "failed to parse expiresAt")
		}

		pc.permissionCache.SetWithTTL(key, &HasPermissionResponse{
			Response:   isAuthRsp,
			StatusCode: code,
		}, 1, expiresAt.Sub(time.Now()))
		if pc.testing {
			pc.permissionCache.Wait()
		}
	}

	return isAuthRsp, code, nil
}

func generateCacheKey(token, permission string) string {
	hash := sha256.Sum256([]byte(token))
	return fmt.Sprintf("%s:%s", hex.EncodeToString(hash[:]), permission)
}
