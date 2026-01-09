package endpoint

import (
	"context"
	"log/slog"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/dgraph-io/ristretto"
	jwtlib "github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"

	openapi "github.com/1backend/1backend/clients/go"

	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/logger"
)

// Subset of teh auth.Authorizer interface.
type JWTParser interface {
	ParseJWT(userSvcPublicKey, token string) (*auth.Claims, error)
}

type TokenRefresher interface {
	EnsureValidToken(request *http.Request) (string, *auth.Claims, error)
}

type fastClock struct {
	now atomic.Int64
}

func (c *fastClock) start() {
	for {
		c.now.Store(time.Now().Unix())
		time.Sleep(500 * time.Millisecond) // Update twice a second
	}
}

func (c *fastClock) get() time.Time {
	return time.Unix(c.now.Load(), 0)
}

type tokenRefresher struct {
	clientFactory         client.ClientFactory
	parser                JWTParser
	tokenReplacementCache *ristretto.Cache
	userServicePublicKey  string
	mutex                 sync.Mutex
	currentTime           time.Time
	once                  sync.Once
	clock                 atomic.Int64
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

	tr := &tokenRefresher{
		clientFactory:         clientFactory,
		parser:                parser,
		tokenReplacementCache: cache,
	}

	tr.clock.Store(time.Now().Unix())

	go tr.backgroundClock()

	return tr, nil
}

func (tr *tokenRefresher) getUserServicePublicKey() (string, error) {
	tr.once.Do(func() {
		keyResp, _, err := tr.clientFactory.Client().
			UserSvcAPI.GetPublicKey(context.Background()).
			Execute()

		if err != nil {
			logger.Error("Failed to get public key",
				slog.Any("error", err),
			)
			return
		}
		tr.userServicePublicKey = keyResp.PublicKey
	})

	return tr.userServicePublicKey, nil
}

func (tr *tokenRefresher) getNow() time.Time {
	// Priority 1: Manual override for tests
	if !tr.currentTime.IsZero() {
		return tr.currentTime
	}
	// Priority 2: High-performance atomic clock for production
	return time.Unix(tr.clock.Load(), 0)
}

func (tr *tokenRefresher) backgroundClock() {
	ticker := time.NewTicker(500 * time.Millisecond)
	for t := range ticker.C {
		tr.clock.Store(t.Unix())
	}
}

func (tr *tokenRefresher) EnsureValidToken(request *http.Request) (string, *auth.Claims, error) {
	authHeader := request.Header.Get("Authorization")
	if len(authHeader) < 7 || !strings.HasPrefix(authHeader, "Bearer ") {
		return "", nil, nil
	}
	// Zero-allocation slicing
	jwt := authHeader[7:]

	now := tr.getNow()

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

func calculateTokenTtl(token openapi.UserSvcToken) (time.Duration, error) {
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
