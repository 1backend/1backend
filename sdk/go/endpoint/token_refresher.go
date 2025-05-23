package endpoint

import (
	"context"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/dgraph-io/ristretto"
	jwtlib "github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"

	openapi "github.com/1backend/1backend/clients/go"

	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/client"
)

// Subset of teh auth.Authorizer interface.
type JWTParser interface {
	ParseJWT(userSvcPublicKey, token string) (*auth.Claims, error)
}

type TokenRefresher interface {
	EnsureValidToken(request *http.Request) (string, *auth.Claims, error)
}

type tokenRefresher struct {
	clientFactory         client.ClientFactory
	parser                JWTParser
	tokenReplacementCache *ristretto.Cache
	userServicePublicKey  string
	mutex                 sync.Mutex
	currentTime           time.Time
}

func NewTokenRefresher(
	clientFactory client.ClientFactory,
	parser JWTParser,
) (TokenRefresher, error) {
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e5,
		MaxCost:     1 << 20,
		BufferItems: 64,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create token replacement cache")
	}

	return &tokenRefresher{
		clientFactory:         clientFactory,
		parser:                parser,
		tokenReplacementCache: cache,
	}, nil
}

func (tr *tokenRefresher) getUserServicePublicKey() (string, error) {
	tr.mutex.Lock()
	defer tr.mutex.Unlock()

	if tr.userServicePublicKey == "" {
		keyResp, _, err := tr.clientFactory.Client().
			UserSvcAPI.GetPublicKey(context.Background()).
			Execute()
		if err != nil {
			return "", errors.Wrap(err, "failed to get public key from user service")
		}
		tr.userServicePublicKey = keyResp.PublicKey
	}
	return tr.userServicePublicKey, nil
}

func (tr *tokenRefresher) EnsureValidToken(request *http.Request) (string, *auth.Claims, error) {
	jwt := strings.TrimSpace(strings.TrimPrefix(request.Header.Get("Authorization"), "Bearer"))
	if jwt == "" {
		return "", nil, nil
	}

	now := tr.currentTime
	if now.IsZero() {
		now = time.Now()
	}

	publicKey, err := tr.getUserServicePublicKey()
	if err != nil {
		return "", nil, err
	}

	claims, err := tr.parser.ParseJWT(publicKey, jwt)
	isExpired := false

	if err != nil {
		if strings.Contains(err.Error(), "token is expired") {
			isExpired = true
		} else {
			return "", nil, errors.Wrap(err, "failed to parse JWT")
		}
	}

	if claims == nil ||
		claims.ExpiresAt == nil ||
		// Here we could add a leeway of 5 seconds but in tests token duration is often 1 second
		// and I'm afraid this would mess with that.
		claims.ExpiresAt.Time.Before(now) {
		isExpired = true
	}

	key := generateCacheKey(jwt, "") // no permission needed here

	if isExpired {
		if val, found := tr.tokenReplacementCache.Get(key); found {
			if replacementToken, ok := val.(string); ok {
				newClaims, err := tr.parser.ParseJWT(publicKey, replacementToken)
				if err == nil && newClaims.ExpiresAt != nil &&
					newClaims.ExpiresAt.Time.After(now.Add(5*time.Second)) {
					jwt = replacementToken
					claims = newClaims
					request.Header.Set("Authorization", "Bearer "+jwt)
					return jwt, claims, nil
				}
			}
		}

		tokenResp, _, err := tr.clientFactory.Client(client.WithTokenFromRequest(request)).
			UserSvcAPI.RefreshToken(request.Context()).Execute()
		if err != nil {
			return "", nil, errors.Wrap(err, "token refresh failed")
		}

		jwt = tokenResp.Token.Token
		request.Header.Set("Authorization", "Bearer "+jwt)

		expiresAt, err := time.Parse(time.RFC3339, tokenResp.Token.ExpiresAt)
		if err != nil {
			return "", nil, errors.Wrap(err, "failed to parse token expiry")
		}

		claims = &auth.Claims{
			RegisteredClaims: jwtlib.RegisteredClaims{
				ExpiresAt: jwtlib.NewNumericDate(expiresAt),
			},
		}

		ttl, err := calculateTokenTtl(tokenResp.Token)
		if err != nil {
			return "", nil, errors.Wrap(err, "failed to calculate token TTL")
		}

		tr.tokenReplacementCache.SetWithTTL(key, jwt, 1, ttl)
	}

	return jwt, claims, nil
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
