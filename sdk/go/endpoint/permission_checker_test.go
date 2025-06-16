package endpoint_test

import (
	"context"
	"net/http/httptest"
	"testing"
	"time"

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/endpoint"
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
	}), test.WithUntilFactory(func() time.Time {
		return time.Now().Add(1 * time.Second)
	}), test.WithHasPermissionTimesFactory(func() int {
		return 2
	}))

	mockClientFactory.EXPECT().
		Client(gomock.Any()).
		Return(&openapi.APIClient{
			UserSvcAPI: mockUserSvc,
		}).Times(2)

	// Use a very short cache duration to test expiry
	pc := endpoint.NewPermissionChecker(mockClientFactory)
	pc.(*endpoint.PermissionCheckerImpl).Testing = true

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
	time.Sleep(time.Second)

	// Third call - should hit service again
	resp3, code3, err3 := pc.HasPermission(req, "perm")
	assert.NoError(t, err3)
	assert.Equal(t, 200, code3)
	assert.True(t, resp3.Authorized)
}
