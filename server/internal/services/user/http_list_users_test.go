package userservice_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	openapi "github.com/1backend/1backend/clients/go"
	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/test"
)

func TestListUsers(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		Test: true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)

	manyClients, tokens, err := test.MakeClients(clientFactory, sdk.DefaultTestAppHost, 5)
	require.NoError(t, err)

	userClient := manyClients[0]

	adminClient, _, err := test.AdminClient(clientFactory, sdk.DefaultTestAppHost)
	require.NoError(t, err)

	t.Run("users can not list users", func(t *testing.T) {
		_, _, err := userClient.UserSvcAPI.ListUsers(
			context.Background(),
		).Execute()
		require.Error(t, err)
	})

	t.Run("admins can list users", func(t *testing.T) {
		rsp, _, err := adminClient.UserSvcAPI.ListUsers(
			context.Background(),
		).Execute()
		require.NoError(t, err)

		require.NotEmpty(t, rsp.Users)
		require.True(t, len(rsp.Users) > 6, rsp)
	})

	t.Run("limit", func(t *testing.T) {
		rsp, _, err := adminClient.UserSvcAPI.ListUsers(
			context.Background(),
		).Body(openapi.UserSvcListUsersRequest{
			Limit: openapi.PtrInt32(3),
		}).Execute()
		require.NoError(t, err)

		require.NotEmpty(t, rsp.Users)
		require.True(t, len(rsp.Users) == 3, rsp)
	})

	t.Run("by ids", func(t *testing.T) {
		rsp, _, err := adminClient.UserSvcAPI.ListUsers(
			context.Background(),
		).Body(openapi.UserSvcListUsersRequest{
			Ids: []string{
				tokens[0].UserId,
				tokens[1].UserId,
				tokens[2].UserId,
			},
		}).Execute()
		require.NoError(t, err)

		require.NotEmpty(t, rsp.Users)
		require.True(t, len(rsp.Users) == 3, rsp)

		for _, user := range rsp.Users {
			require.Contains(t, []string{
				tokens[0].UserId,
				tokens[1].UserId,
				tokens[2].UserId,
			}, user.Id)
		}
	})
}
