package endpoint

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

type fakeJWTParser struct {
	claimsByToken map[string]*auth.Claims
}

func (p *fakeJWTParser) ParseJWT(_ string, token string) (*auth.Claims, error) {
	claims := p.claimsByToken[token]
	if claims == nil {
		return nil, nil
	}
	return claims, nil
}

type expectedRefresh struct {
	token     string
	expiresAt time.Time
}

func newTestTokenRefresher(
	t *testing.T,
	now time.Time,
	oldToken string,
	issuedAt time.Time,
	expiresAt time.Time,
	refresh *expectedRefresh,
) TokenRefresher {
	t.Helper()
	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	mockClientFactory := client.NewMockClientFactory(ctrl)
	mockUserSvc := openapi.NewMockUserSvcAPI(ctrl)
	mockAPIClient := &openapi.APIClient{UserSvcAPI: mockUserSvc}

	mockClientFactory.EXPECT().Client().Return(mockAPIClient)
	mockUserSvc.EXPECT().GetPublicKey(gomock.Any()).Return(openapi.ApiGetPublicKeyRequest{
		ApiService: mockUserSvc,
	})
	mockUserSvc.EXPECT().GetPublicKeyExecute(gomock.Any()).Return(
		&openapi.UserSvcGetPublicKeyResponse{PublicKey: "public-key"},
		&http.Response{StatusCode: http.StatusOK},
		nil,
	)

	if refresh != nil {
		mockClientFactory.EXPECT().Client(gomock.Any()).Return(mockAPIClient)
		mockUserSvc.EXPECT().RefreshToken(gomock.Any()).Return(openapi.ApiRefreshTokenRequest{
			ApiService: mockUserSvc,
		})
		mockUserSvc.EXPECT().RefreshTokenExecute(gomock.Any()).Return(
			&openapi.UserSvcRefreshTokenResponse{
				Token: openapi.UserSvcToken{
					Token:     refresh.token,
					ExpiresAt: refresh.expiresAt.Format(time.RFC3339),
				},
			},
			&http.Response{StatusCode: http.StatusOK},
			nil,
		)
	}

	refresher, err := NewTokenRefresher(mockClientFactory, &fakeJWTParser{
		claimsByToken: map[string]*auth.Claims{
			oldToken: {
				RegisteredClaims: jwt.RegisteredClaims{
					IssuedAt:  jwt.NewNumericDate(issuedAt),
					ExpiresAt: jwt.NewNumericDate(expiresAt),
				},
			},
		},
	})
	require.NoError(t, err)
	refresher.(*tokenRefresher).currentTime = now

	return refresher
}

func TestEnsureValidTokenRefreshesNearExpiryForNormalLifetime(t *testing.T) {
	now := time.Date(2026, 4, 25, 12, 0, 0, 0, time.UTC)
	oldToken := "old-token"
	newToken := "new-token"
	newExpiry := now.Add(time.Hour)

	refresher := newTestTokenRefresher(
		t,
		now,
		oldToken,
		now.Add(-time.Minute),
		now.Add(2*time.Second),
		&expectedRefresh{token: newToken, expiresAt: newExpiry},
	)

	req := httptest.NewRequest(http.MethodGet, "/secret-svc/secrets", nil)
	req.Header.Set("Authorization", "Bearer "+oldToken)

	token, claims, err := refresher.EnsureValidToken(req)
	require.NoError(t, err)
	require.Equal(t, newToken, token)
	require.Equal(t, "Bearer "+newToken, req.Header.Get("Authorization"))
	require.Equal(t, newExpiry, claims.ExpiresAt.Time)
}

func TestEnsureValidTokenDoesNotRefreshOutsidePercentageLeeway(t *testing.T) {
	now := time.Date(2026, 4, 25, 12, 0, 0, 0, time.UTC)
	oldToken := "old-token"

	refresher := newTestTokenRefresher(
		t,
		now,
		oldToken,
		now.Add(-8*time.Second),
		now.Add(2*time.Second),
		nil,
	)

	req := httptest.NewRequest(http.MethodGet, "/secret-svc/secrets", nil)
	req.Header.Set("Authorization", "Bearer "+oldToken)

	token, _, err := refresher.EnsureValidToken(req)
	require.NoError(t, err)
	require.Equal(t, oldToken, token)
	require.Equal(t, "Bearer "+oldToken, req.Header.Get("Authorization"))
}

func TestEnsureValidTokenRefreshesWithinPercentageLeewayForShortLifetime(t *testing.T) {
	now := time.Date(2026, 4, 25, 12, 0, 0, 0, time.UTC)
	oldToken := "old-token"
	newToken := "new-token"
	newExpiry := now.Add(time.Minute)

	refresher := newTestTokenRefresher(
		t,
		now,
		oldToken,
		now.Add(-9*time.Second),
		now.Add(time.Second),
		&expectedRefresh{token: newToken, expiresAt: newExpiry},
	)

	req := httptest.NewRequest(http.MethodGet, "/secret-svc/secrets", nil)
	req.Header.Set("Authorization", "Bearer "+oldToken)

	token, claims, err := refresher.EnsureValidToken(req)
	require.NoError(t, err)
	require.Equal(t, newToken, token)
	require.Equal(t, "Bearer "+newToken, req.Header.Get("Authorization"))
	require.Equal(t, newExpiry, claims.ExpiresAt.Time)
}

func TestEnsureValidTokenExchangesOpaqueAPIKeyBearer(t *testing.T) {
	apiKey := "obk_apik-test_secret"
	exchangedToken := "jwt-from-api-key"
	expiresAt := time.Now().Add(time.Hour).UTC().Truncate(time.Second)
	var calls atomic.Int32

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/user-svc/api-key/exchange", r.URL.Path)
		require.Equal(t, "Bearer "+apiKey, r.Header.Get("Authorization"))
		calls.Add(1)

		w.Header().Set("Content-Type", "application/json")
		require.NoError(t, json.NewEncoder(w).Encode(apiKeyExchangeResponse{
			Token: openapi.UserSvcToken{
				Token:     exchangedToken,
				ExpiresAt: expiresAt.Format(time.RFC3339),
			},
		}))
	}))
	defer server.Close()

	refresher, err := NewTokenRefresher(
		client.NewApiClientFactory(server.URL),
		&fakeJWTParser{},
	)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodGet, "/secret-svc/secrets", nil)
	req.Header.Set("Authorization", "Bearer "+apiKey)

	token, claims, err := refresher.EnsureValidToken(req)
	require.NoError(t, err)
	require.Equal(t, exchangedToken, token)
	require.Equal(t, "Bearer "+exchangedToken, req.Header.Get("Authorization"))
	require.Equal(t, expiresAt, claims.ExpiresAt.Time)
	refresher.(*tokenRefresher).tokenReplacementCache.Wait()

	req2 := httptest.NewRequest(http.MethodGet, "/secret-svc/secrets", nil)
	req2.Header.Set("Authorization", "Bearer "+apiKey)

	token, claims, err = refresher.EnsureValidToken(req2)
	require.NoError(t, err)
	require.Equal(t, exchangedToken, token)
	require.Equal(t, "Bearer "+exchangedToken, req2.Header.Get("Authorization"))
	require.Equal(t, expiresAt, claims.ExpiresAt.Time)
	require.Equal(t, int32(1), calls.Load(), "API key exchange should be cached until the exchanged JWT expires")
}

func TestTokenRefreshLeewayIsPercentageCapped(t *testing.T) {
	now := time.Date(2026, 4, 25, 12, 0, 0, 0, time.UTC)

	require.Equal(t, time.Second, tokenRefreshLeeway(&auth.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(10 * time.Second)),
		},
	}))

	require.Equal(t, maxTokenRefreshLeeway, tokenRefreshLeeway(&auth.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(5 * time.Minute)),
		},
	}))
}
