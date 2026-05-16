package endpoint

import (
	"context"
	"encoding/json"
	"io"
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
	"github.com/1backend/1backend/sdk/go/telemetry"
)

// Subset of teh auth.Authorizer interface.
type JWTParser interface {
	ParseJWT(userSvcPublicKey, token string) (*auth.Claims, error)
}

type TokenRefresher interface {
	EnsureValidToken(request *http.Request) (string, *auth.Claims, error)
}

const (
	tokenRefreshLeewayFraction = 0.10
	maxTokenRefreshLeeway      = 5 * time.Second
)

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

type apiKeyExchangeResponse struct {
	Token openapi.UserSvcToken `json:"token"`
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

	tr.clock.Store(time.Now().UnixNano())

	go tr.backgroundClock()

	return tr, nil
}

func (tr *tokenRefresher) getUserServicePublicKey() (string, error) {
	tr.once.Do(func() {
		started := time.Now()
		keyResp, _, err := tr.clientFactory.Client().
			UserSvcAPI.GetPublicKey(context.Background()).
			Execute()
		telemetry.RecordAuthOperation(context.Background(), telemetry.AuthOperationPublicKeyFetch, authResult(err), "network", started, err)

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

	return time.Unix(0, tr.clock.Load())
}

func (tr *tokenRefresher) backgroundClock() {
	ticker := time.NewTicker(500 * time.Millisecond)
	for t := range ticker.C {
		tr.clock.Store(t.UnixNano())
	}
}

func (tr *tokenRefresher) EnsureValidToken(request *http.Request) (token string, claims *auth.Claims, err error) {
	started := time.Now()
	source := "valid"
	defer func() {
		telemetry.RecordAuthOperation(request.Context(), telemetry.AuthOperationTokenRefresh, authResult(err), source, started, err)
	}()

	authHeader := request.Header.Get("Authorization")
	if len(authHeader) < 7 || !strings.HasPrefix(authHeader, "Bearer ") {
		source = "no_token"
		return "", nil, nil
	}
	// Zero-allocation slicing
	jwt := authHeader[7:]

	if strings.HasPrefix(jwt, "obk_") {
		source = "api_key"
		exchangedToken, exchangedClaims, err := tr.exchangeAPIKeyBearer(request.Context(), jwt)
		if err != nil {
			return "", nil, err
		}
		request.Header.Set("Authorization", "Bearer "+exchangedToken)
		return exchangedToken, exchangedClaims, nil
	}

	now := tr.getNow()

	publicKey, err := tr.getUserServicePublicKey()
	if err != nil {
		return "", nil, err
	}

	claims, err = tr.parser.ParseJWT(publicKey, jwt)
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
		shouldRefreshByExpiry(claims, now) {
		isExpired = true
	}

	key := generateCacheKey(jwt, "") // no permission needed here

	if isExpired {
		if val, found := tr.tokenReplacementCache.Get(key); found {
			if replacementToken, ok := val.(string); ok {
				newClaims, err := tr.parser.ParseJWT(publicKey, replacementToken)
				if err == nil && newClaims.ExpiresAt != nil &&
					newClaims.ExpiresAt.Time.After(now.Add(5*time.Second)) {
					source = "cache"
					telemetry.RecordAuthCache(request.Context(), telemetry.AuthCacheTokenRefresh, "hit")
					jwt = replacementToken
					claims = newClaims
					request.Header.Set("Authorization", "Bearer "+jwt)
					return jwt, claims, nil
				}
				telemetry.RecordAuthCache(request.Context(), telemetry.AuthCacheTokenRefresh, "invalid")
			}
		}
		telemetry.RecordAuthCache(request.Context(), telemetry.AuthCacheTokenRefresh, "miss")
		source = "network"

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

func (tr *tokenRefresher) exchangeAPIKeyBearer(
	ctx context.Context,
	apiKey string,
) (string, *auth.Claims, error) {
	now := tr.getNow()
	cacheKey := generateCacheKey(apiKey, "api-key")

	if val, found := tr.tokenReplacementCache.Get(cacheKey); found {
		if cached, ok := val.(*apiKeyExchangeResponse); ok {
			expiresAt, err := time.Parse(time.RFC3339, cached.Token.ExpiresAt)
			if err == nil && expiresAt.After(now.Add(5*time.Second)) {
				telemetry.RecordAuthCache(ctx, telemetry.AuthCacheTokenRefresh, "api_key_hit")
				return cached.Token.Token, claimsWithExpiry(expiresAt), nil
			}
			tr.tokenReplacementCache.Del(cacheKey)
			telemetry.RecordAuthCache(ctx, telemetry.AuthCacheTokenRefresh, "api_key_expired")
		}
	}

	telemetry.RecordAuthCache(ctx, telemetry.AuthCacheTokenRefresh, "api_key_miss")

	apiClient := tr.clientFactory.Client()
	if apiClient == nil || apiClient.GetConfig() == nil {
		return "", nil, errors.New("api client is nil")
	}

	cfg := apiClient.GetConfig()
	serverURL, err := cfg.ServerURL(0, nil)
	if err != nil {
		return "", nil, errors.Wrap(err, "failed to resolve server URL")
	}

	httpClient := cfg.HTTPClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	url := strings.TrimRight(serverURL, "/") + "/user-svc/api-key/exchange"
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, nil)
	if err != nil {
		return "", nil, err
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Accept", "application/json")

	httpResp, err := httpClient.Do(req)
	if err != nil {
		return "", nil, errors.Wrap(err, "API key exchange failed")
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode < 200 || httpResp.StatusCode >= 300 {
		body, _ := io.ReadAll(io.LimitReader(httpResp.Body, 512))
		return "", nil, errors.Errorf("API key exchange failed: status %d: %s", httpResp.StatusCode, strings.TrimSpace(string(body)))
	}

	var exchangeResp apiKeyExchangeResponse
	if err := json.NewDecoder(httpResp.Body).Decode(&exchangeResp); err != nil {
		return "", nil, errors.Wrap(err, "failed to decode API key exchange response")
	}
	if exchangeResp.Token.Token == "" {
		return "", nil, errors.New("API key exchange response token is empty")
	}

	expiresAt, err := time.Parse(time.RFC3339, exchangeResp.Token.ExpiresAt)
	if err != nil {
		return "", nil, errors.Wrap(err, "failed to parse API key exchange token expiry")
	}

	ttl := time.Until(expiresAt)
	if ttl > 0 {
		tr.tokenReplacementCache.SetWithTTL(cacheKey, &exchangeResp, 1, ttl)
	}

	return exchangeResp.Token.Token, claimsWithExpiry(expiresAt), nil
}

func claimsWithExpiry(expiresAt time.Time) *auth.Claims {
	return &auth.Claims{
		RegisteredClaims: jwtlib.RegisteredClaims{
			ExpiresAt: jwtlib.NewNumericDate(expiresAt),
		},
	}
}

func authResult(err error) string {
	if err != nil {
		return "error"
	}
	return "success"
}

func shouldRefreshByExpiry(claims *auth.Claims, now time.Time) bool {
	if claims.ExpiresAt == nil {
		return true
	}

	if !claims.ExpiresAt.Time.After(now) {
		return true
	}

	if claims.IssuedAt == nil {
		return false
	}

	leeway := tokenRefreshLeeway(claims)
	if leeway <= 0 {
		return false
	}

	return !claims.ExpiresAt.Time.After(now.Add(leeway))
}

func tokenRefreshLeeway(claims *auth.Claims) time.Duration {
	if claims == nil || claims.IssuedAt == nil || claims.ExpiresAt == nil {
		return 0
	}

	lifetime := claims.ExpiresAt.Time.Sub(claims.IssuedAt.Time)
	if lifetime <= 0 {
		return 0
	}

	leeway := time.Duration(float64(lifetime) * tokenRefreshLeewayFraction)
	if leeway > maxTokenRefreshLeeway {
		return maxTokenRefreshLeeway
	}

	return leeway
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
