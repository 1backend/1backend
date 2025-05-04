package userservice_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/test"
)

func TestCreateUser(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		Test: true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)

	ctx := context.Background()

	t.Run("non-admins cannot create users", func(t *testing.T) {
		_, _, err = clientFactory.Client().UserSvcAPI.CreateUser(ctx).Body(
			openapi.UserSvcCreateUserRequest{
				User: &openapi.UserSvcUserInput{
					Slug: "test-slug-1",
					Name: openapi.PtrString("Test Name"),
				},
			},
		).Execute()
		require.Error(t, err)
	})

	adminClient, _, err := test.AdminClient(clientFactory)
	require.NoError(t, err)

	t.Run("admins can create users", func(t *testing.T) {
		_, httpRsp, err := adminClient.UserSvcAPI.CreateUser(ctx).Body(
			openapi.UserSvcCreateUserRequest{
				User: &openapi.UserSvcUserInput{
					Slug: "test-slug-1",
					Name: openapi.PtrString("Test Name"),
				},
			},
		).Execute()
		require.NoError(t, err, httpRsp)
	})

	t.Run("slug is taken", func(t *testing.T) {
		_, _, err = adminClient.UserSvcAPI.CreateUser(ctx).Body(
			openapi.UserSvcCreateUserRequest{
				User: &openapi.UserSvcUserInput{
					Slug: "test-slug-1",
					Name: openapi.PtrString("Test Name"),
				},
			},
		).Execute()
		require.Error(t, err)
	})
}
