package registryservice

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/infra"
	registry "github.com/1backend/1backend/server/internal/services/registry/types"
	"github.com/1backend/1backend/server/internal/universe"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

type fakeRegistryPermissionChecker struct {
	response *openapi.UserSvcHasPermissionResponse
}

func (c *fakeRegistryPermissionChecker) HasPermission(
	_ *http.Request,
	_ string,
) (*openapi.UserSvcHasPermissionResponse, int, error) {
	return c.response, http.StatusOK, nil
}

func TestListInstances_ExpiredAdminCheckFallsBackToOwnerScope(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	authorizer := auth.NewMockAuthorizer(ctrl)
	authorizer.EXPECT().
		IsAdminFromRequest("public-key", gomock.Any()).
		Return(false, errors.New("failed to parse JWT: token has invalid claims: token is expired"))

	dataStoreFactory, err := infra.NewDataStoreFactory(infra.DataStoreConfig{
		Test:    true,
		HomeDir: t.TempDir(),
	})
	require.NoError(t, err)

	const ownerSlug = "test-user-slug-0"
	rs, err := NewRegistryService(&universe.Options{
		DataStoreFactory: dataStoreFactory,
		Authorizer:       authorizer,
		PermissionChecker: &fakeRegistryPermissionChecker{
			response: &openapi.UserSvcHasPermissionResponse{
				Authorized: true,
				Until:      time.Now().Add(time.Minute).Format(time.RFC3339),
				User: openapi.UserSvcUser{
					Slug: ownerSlug,
				},
			},
		},
	})
	require.NoError(t, err)
	rs.publicKey = "public-key"

	err = rs.instanceStore.Create(&registry.Instance{
		Id:   "inst-test",
		URL:  "http://127.0.0.1:1234",
		Slug: ownerSlug,
	})
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodGet, "/registry-svc/instances", nil)
	req.Header.Set("Authorization", "Bearer expired-token")
	rec := httptest.NewRecorder()

	rs.ListInstances(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)

	var rsp registry.ListInstancesResponse
	err = json.Unmarshal(rec.Body.Bytes(), &rsp)
	require.NoError(t, err)
	require.Len(t, rsp.Instances, 1)
	require.Equal(t, "http://127.0.0.1:1234", rsp.Instances[0].URL)
}
