package userservice_test

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	openapi "github.com/openorch/openorch/clients/go"
	"github.com/openorch/openorch/sdk/go/test"
	"github.com/openorch/openorch/server/internal/di"
)

func TestCreateUser(t *testing.T) {
	hs := &di.HandlerSwitcher{}
	server := httptest.NewServer(hs)
	defer server.Close()

	options := &di.Options{
		Test: true,
		Url:  server.URL,
	}
	universe, starterFunc, err := di.BigBang(options)
	require.NoError(t, err)

	hs.UpdateHandler(universe)

	err = starterFunc()
	require.NoError(t, err)

	ctx := context.Background()

	t.Run("non-admins cannot create users", func(t *testing.T) {
		_, _, err = options.ClientFactory.Client().UserSvcAPI.CreateUser(ctx).Body(
			openapi.UserSvcCreateUserRequest{
				User: &openapi.UserSvcUser{
					Slug: openapi.PtrString("test-slug-1"),
					Name: openapi.PtrString("Test Name"),
				},
			},
		).Execute()
		require.Error(t, err)
	})

	adminClient, _, err := test.AdminClient(options.ClientFactory)
	require.NoError(t, err)

	t.Run("admins can create users", func(t *testing.T) {
		_, httpRsp, err := adminClient.UserSvcAPI.CreateUser(ctx).Body(
			openapi.UserSvcCreateUserRequest{
				User: &openapi.UserSvcUser{
					Slug: openapi.PtrString("test-slug-1"),
					Name: openapi.PtrString("Test Name"),
				},
			},
		).Execute()
		require.NoError(t, err, httpRsp)
	})

	t.Run("slug is taken", func(t *testing.T) {
		_, _, err = adminClient.UserSvcAPI.CreateUser(ctx).Body(
			openapi.UserSvcCreateUserRequest{
				User: &openapi.UserSvcUser{
					Slug: openapi.PtrString("test-slug-1"),
					Name: openapi.PtrString("Test Name"),
				},
			},
		).Execute()
		require.Error(t, err)
	})
}
