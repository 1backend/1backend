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

func TestHasPermission_CacheHit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClientFactory := client.NewMockClientFactory(ctrl)
	mockUserSvc := test.MockUserSvc(context.Background(), ctrl, test.WithHasPermissionFactory(func() bool {
		return true
	}))
	mockClientFactory.EXPECT().
		Client(gomock.Any()).
		Return(&openapi.APIClient{
			UserSvcAPI: mockUserSvc,
		}).
		AnyTimes()

	mockAuthorizer := auth.NewMockAuthorizer(ctrl)
	mockAuthorizer.EXPECT().
		ParseJWT(gomock.Any(), gomock.Any()).
		Return(&auth.Claims{
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(
					time.Now().Add(1 * time.Hour),
				),
			},
		}, nil)

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

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", jwtToken)

	resp, code, err := pc.HasPermission(req, "test-permission")

	assert.NoError(t, err)
	assert.Equal(t, 200, code)
	assert.True(t, resp.Authorized)
}
