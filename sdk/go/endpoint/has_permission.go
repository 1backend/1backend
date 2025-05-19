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
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"

	jwtlib "github.com/golang-jwt/jwt/v5"

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/auth"
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
	clientFactory         client.ClientFactory
	parser                JWTParser
	permissionCache       *ristretto.Cache
	tokenReplacementCache *ristretto.Cache
	userServicePublicKey  string
	mutex                 sync.Mutex

	// currentTime is used to mock the current time in tests.
	currentTime time.Time
}

// Subset of teh auth.Authorizer interface.
type JWTParser interface {
	ParseJWT(userSvcPublicKey, token string) (*auth.Claims, error)
}

func NewPermissionChecker(
	clientFactory client.ClientFactory,
	parser JWTParser,
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

	tokenReplacementCache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e5,
		MaxCost:     1 << 20,
		BufferItems: 64,
	})
	if err != nil {
		panic(errors.Wrap(err, "failed to create token replacement cache").Error())
	}

	return &permissionChecker{
		clientFactory:         clientFactory,
		permissionCache:       permissionCache,
		tokenReplacementCache: tokenReplacementCache,
		parser:                parser,
	}
}

func (pc *permissionChecker) getUserServicePublicKey() (string, error) {
	pc.mutex.Lock()
	defer pc.mutex.Unlock()

	if pc.userServicePublicKey == "" {
		userServicePublicKey, _, err := pc.clientFactory.Client().
			UserSvcAPI.GetPublicKey(context.Background()).Execute()

		if err != nil {
			return "", errors.Wrap(err, "user service get public key endpoint failed")
		}
		pc.userServicePublicKey = userServicePublicKey.PublicKey
	}

	return pc.userServicePublicKey, nil
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
	publicKey, err := pc.getUserServicePublicKey()
	if err != nil {
		return nil, http.StatusInternalServerError, errors.Wrap(err, "failed to get user service public key")
	}

	isExpired := false

	jwt := strings.TrimPrefix(request.Header.Get("Authorization"), "Bearer ")

	now := pc.currentTime
	if now.IsZero() {
		now = time.Now()
	}

	var (
		claims *auth.Claims
	)

	if jwt != "" {
		var err error

		// First we check if the JWT is expired or not as that aspect cannot be cached.
		// If the JWT is expired, we consider it a cache miss
		// and let the request go through to the user service
		claims, err = pc.parser.ParseJWT(
			publicKey,
			jwt,
		)
		if err != nil {
			if strings.Contains(err.Error(), "token is expired") {
				isExpired = true
			} else {
				return nil, http.StatusUnauthorized, errors.Wrap(err, "failed to parse JWT")
			}
		}

		// Handle missing expiresAt for backwards compatibility.
		// Can be removed later.
		if claims == nil ||
			claims.ExpiresAt == nil ||
			claims.ExpiresAt.Time.Before(now.Add(5*time.Second)) {
			isExpired = true
		}
	}

	key := generateCacheKey(
		jwt,
		permission,
	)

	if isExpired {
		if val, found := pc.tokenReplacementCache.Get(key); found {
			if replacementToken, ok := val.(string); ok {
				newClaims, err := pc.parser.ParseJWT(publicKey, replacementToken)
				if err == nil &&
					newClaims.ExpiresAt != nil &&
					newClaims.ExpiresAt.Time.After(now.Add(5*time.Second)) {
					jwt = replacementToken
					request.Header.Set("Authorization", "Bearer "+jwt)
					isExpired = false

					// Claims will be used to calculate the TTL for the cache.
					claims = newClaims
				}
			}
		}

		if isExpired {
			newTokenResp, _, err := pc.clientFactory.Client(client.WithTokenFromRequest(request)).
				UserSvcAPI.RefreshToken(request.Context()).Execute()
			if err != nil {
				return nil, http.StatusUnauthorized, errors.Wrap(err, "token refresh failed")
			}

			jwt = newTokenResp.Token.Token

			expiresAt, err := time.Parse(time.RFC3339, newTokenResp.Token.ExpiresAt)
			if err != nil {
				return nil, http.StatusInternalServerError, errors.Wrap(err, "failed to parse token expiresAt")
			}

			// Claims will be used to calculate the TTL for the cache.
			claims = &auth.Claims{
				RegisteredClaims: jwtlib.RegisteredClaims{
					ExpiresAt: jwtlib.NewNumericDate(expiresAt),
				},
			}

			request.Header.Set("Authorization", "Bearer "+jwt)

			ttl, err := calculateTokenTtl(newTokenResp.Token)
			if err != nil {
				return nil, http.StatusInternalServerError, errors.Wrap(err, "failed to calculate token ttl")
			}

			pc.tokenReplacementCache.SetWithTTL(key, jwt, 1, ttl)
		}
	}

	if jwt != "" && !isExpired {
		if value, found := pc.permissionCache.Get(key); found {
			if cachedResp, ok := value.(*HasPermissionResponse); ok {
				return cachedResp.Response, cachedResp.StatusCode, nil
			}
		}
	}

	// At this point we either don't have a JWT or we have a refreshed one.
	// @todo is the whole refreshing logic above unnecessary because HasPermission
	// can do its own refreshing? Probably not because we don't return the Token itself
	// in HasPermission.
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

	ttl := 5 * time.Minute

	// JWT is not present when the user is not logged in.
	if jwt != "" {
		ttl, err = calculateClaimsTtl(claims)
		if err != nil {
			return nil, http.StatusInternalServerError, errors.Wrap(err, "failed to calculate claims ttl")
		}
	}

	if key != "" {
		pc.permissionCache.SetWithTTL(key, &HasPermissionResponse{
			Response:   isAuthRsp,
			StatusCode: httpResponse.StatusCode,
		}, 1, ttl)
	}

	code := 0
	if httpResponse != nil {
		code = httpResponse.StatusCode
	}

	return isAuthRsp, code, nil
}

func generateCacheKey(token, permission string) string {
	hash := sha256.Sum256([]byte(token))
	return fmt.Sprintf("%s:%s", hex.EncodeToString(hash[:]), permission)
}

func calculateTokenTtl(token openapi.UserSvcAuthToken) (time.Duration, error) {
	if token.ExpiresAt == "" {
		return 0, errors.New("token expiresAt is empty")
	}

	expiresAt, err := time.Parse(time.RFC3339, token.ExpiresAt)
	if err != nil {
		return 0, errors.Wrap(err, "failed to parse token expiresAt")
	}

	ttl := time.Until(expiresAt)
	if ttl < time.Second {
		ttl = time.Second
	}

	return ttl, nil
}

func calculateClaimsTtl(claims *auth.Claims) (time.Duration, error) {
	if claims == nil {
		return 0, errors.New("claims is nil")
	}
	if claims.ExpiresAt == nil {
		return 0, errors.New("expiresAt is nil")
	}

	ttl := time.Until(claims.ExpiresAt.Time)
	if ttl < time.Second {
		ttl = time.Second
	}

	return ttl, nil
}
