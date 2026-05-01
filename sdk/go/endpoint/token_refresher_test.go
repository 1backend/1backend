package endpoint

import (
	"net/http"
	"net/http/httptest"
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
