package endpoint

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/test"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestHasPermissionCaching(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClientFactory := client.NewMockClientFactory(ctrl)

	mockUserSvc := test.MockUserSvc(context.Background(), ctrl, test.WithHasPermissionFactory(func() bool {
		return true
	}))

	mockUserSvc.EXPECT().HasPermission(gomock.Any(), gomock.Any()).
		Return(openapi.ApiHasPermissionRequest{
			ApiService: mockUserSvc,
		}).Times(2)

	mockUserSvc.EXPECT().HasPermissionExecute(gomock.Any()).
		Return(&openapi.UserSvcHasPermissionResponse{
			Authorized: true,
			Until:      time.Now().Add(50 * time.Millisecond).Format(time.RFC3339),
		}, &http.Response{
			StatusCode: http.StatusOK,
		}, nil).Times(2)

	mockClientFactory.EXPECT().
		Client(gomock.Any()).
		Return(&openapi.APIClient{
			UserSvcAPI: mockUserSvc,
		}).Times(2)

	// Use a very short cache duration to test expiry
	pc := NewPermissionChecker(mockClientFactory)
	pc.(*permissionChecker).testing = true

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Bearer token")

	// First call - should hit service
	resp1, code1, err1 := pc.HasPermission(req, "perm")
	assert.NoError(t, err1)
	assert.Equal(t, 200, code1)
	assert.True(t, resp1.Authorized)

	// Second call - should hit cache, not service
	resp2, code2, err2 := pc.HasPermission(req, "perm")
	assert.NoError(t, err2)
	assert.Equal(t, 200, code2)
	assert.True(t, resp2.Authorized)

	// Wait for cache to expire
	time.Sleep(60 * time.Millisecond)

	// Third call - should hit service again
	resp3, code3, err3 := pc.HasPermission(req, "perm")
	assert.NoError(t, err3)
	assert.Equal(t, 200, code3)
	assert.True(t, resp3.Authorized)
}
