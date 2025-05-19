package endpoint

import (
	"context"
	"net/http/httptest"
	"testing"
	"time"

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/test"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestHasPermission(t *testing.T) {
	t.Run("cache miss - service call happens", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockClientFactory := client.NewMockClientFactory(ctrl)

		called := false
		mockUserSvc := test.MockUserSvc(context.Background(), ctrl, test.WithHasPermissionFactory(func() bool {
			called = true
			return true
		}))

		mockClientFactory.EXPECT().
			Client(gomock.Any()).
			Return(&openapi.APIClient{
				UserSvcAPI: mockUserSvc,
			}).AnyTimes()

		mockAuthorizer := auth.NewMockAuthorizer(ctrl)
		mockAuthorizer.EXPECT().
			ParseJWT(gomock.Any(), gomock.Any()).
			Return(&auth.Claims{
				RegisteredClaims: jwt.RegisteredClaims{
					ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
				},
			}, nil).Times(1)

		pc := NewPermissionChecker(mockClientFactory, mockAuthorizer)

		req := httptest.NewRequest("GET", "/", nil)
		req.Header.Set("Authorization", "Bearer testtoken")

		resp, code, err := pc.HasPermission(req, "test-permission")

		assert.NoError(t, err)
		assert.Equal(t, 200, code)
		assert.True(t, resp.Authorized)
		assert.True(t, called, "expected service call to happen on cache miss")
	})

	t.Run("cache hit - no service call", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockClientFactory := client.NewMockClientFactory(ctrl)
		mockAuthorizer := auth.NewMockAuthorizer(ctrl)

		mockUserSvc := test.MockUserSvc(context.Background(), ctrl, test.WithHasPermissionFactory(func() bool {
			return true
		}))

		mockClientFactory.EXPECT().
			Client(gomock.Any()).
			Return(&openapi.APIClient{
				UserSvcAPI: mockUserSvc,
			}).AnyTimes()

		mockAuthorizer.EXPECT().
			ParseJWT(gomock.Any(), gomock.Any()).
			Return(&auth.Claims{
				RegisteredClaims: jwt.RegisteredClaims{
					ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
				},
			}, nil).Times(1)

		pc := NewPermissionChecker(mockClientFactory, mockAuthorizer)

		jwtToken := "Bearer testtoken"
		cacheKey := generateCacheKey(jwtToken, "test-permission")

		_ = pc.(*permissionChecker).permissionCache.SetWithTTL(
			cacheKey,
			&HasPermissionResponse{
				Response: &openapi.UserSvcHasPermissionResponse{
					Authorized: true,
				},
				StatusCode: 200,
			}, 1, 5*time.Minute)
		pc.(*permissionChecker).permissionCache.Wait()

		req := httptest.NewRequest("GET", "/", nil)
		req.Header.Set("Authorization", jwtToken)

		resp, code, err := pc.HasPermission(req, "test-permission")

		assert.NoError(t, err)
		assert.Equal(t, 200, code)
		assert.True(t, resp.Authorized)
	})

	t.Run("expired token - first call hits service, second call uses cache", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockClientFactory := client.NewMockClientFactory(ctrl)

		mockUserSvc := test.MockUserSvc(context.Background(), ctrl, test.WithHasPermissionFactory(func() bool {
			return true
		}))
		mockUserSvc.EXPECT().RefreshToken(gomock.Any()).
			Return(openapi.ApiRefreshTokenRequest{
				ApiService: mockUserSvc,
			}).Times(2)
		mockUserSvc.EXPECT().RefreshTokenExecute(gomock.Any()).
			Return(&openapi.UserSvcRefreshTokenResponse{
				Token: openapi.UserSvcAuthToken{
					Token:     "Hello 2",
					ExpiresAt: time.Now().Add(5 * time.Second).Format(time.RFC3339),
				},
			}, nil, nil).Times(2)

		mockClientFactory.EXPECT().
			Client(gomock.Any()).
			Return(&openapi.APIClient{
				UserSvcAPI: mockUserSvc,
			}).AnyTimes()

		mockAuthorizer := auth.NewMockAuthorizer(ctrl)

		callCount := 0
		mockAuthorizer.EXPECT().
			ParseJWT(gomock.Any(), gomock.Any()).
			DoAndReturn(func(pubKey, token string) (*auth.Claims, error) {
				callCount++
				switch callCount {
				case 1:
					return &auth.Claims{
						RegisteredClaims: jwt.RegisteredClaims{
							ExpiresAt: jwt.NewNumericDate(time.Now().Add(-time.Minute)),
						},
					}, nil // expired token
				case 2:
					return &auth.Claims{
						RegisteredClaims: jwt.RegisteredClaims{
							ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
						},
					}, nil // refreshed token
				case 3:
					return &auth.Claims{
						RegisteredClaims: jwt.RegisteredClaims{
							ExpiresAt: jwt.NewNumericDate(time.Now().Add(-time.Minute)),
						},
					}, nil // expired again
				case 4:
					return &auth.Claims{
						RegisteredClaims: jwt.RegisteredClaims{
							ExpiresAt: jwt.NewNumericDate(time.Now().Add(20 * time.Minute)),
						},
					}, nil // fresh
				}

				return nil, nil
			}).
			Times(3)

		pc := NewPermissionChecker(mockClientFactory, mockAuthorizer).(*permissionChecker)
		pc.currentTime = time.Now()

		req := httptest.NewRequest("GET", "/", nil)
		jwtToken := "Bearer expiredtoken"
		req.Header.Set("Authorization", jwtToken)

		// First call - should miss cache and call service
		resp1, code1, err1 := pc.HasPermission(req, "expired-permission")
		assert.NoError(t, err1)
		assert.Equal(t, 200, code1)
		assert.True(t, resp1.Authorized)

		pc.permissionCache.Wait()

		// Second call - same expired JWT, but should now hit cache
		resp2, code2, err2 := pc.HasPermission(req, "expired-permission")
		assert.NoError(t, err2)
		assert.Equal(t, 200, code2)
		assert.True(t, resp2.Authorized)

		pc.currentTime = time.Now().Add(1 * time.Hour)

		// Third call - even the mapped token is expired
		resp3, code3, err3 := pc.HasPermission(req, "expired-permission")
		assert.NoError(t, err3)
		assert.Equal(t, 200, code3)
		assert.True(t, resp3.Authorized)
	})
}
