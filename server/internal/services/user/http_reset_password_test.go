package userservice_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/test"
)

func TestResetPassword(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		Test: true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)

	manyClients, tokens, err := test.MakeClients(clientFactory, 2)
	require.NoError(t, err)

	client1 := manyClients[0]

	adminClient, _, err := test.AdminClient(clientFactory)
	require.NoError(t, err)

	ctx := context.Background()

	t.Run("user 1 cannot reset password of user 2", func(t *testing.T) {
		_, _, err := client1.UserSvcAPI.ResetPassword(ctx, tokens[1].UserId).Body(
			openapi.UserSvcResetPasswordRequest{
				NewPassword: "testPass123",
			}).Execute()

		require.Error(t, err)
	})

	t.Run("admin can reset password of user 1", func(t *testing.T) {
		_, _, err := adminClient.UserSvcAPI.ResetPassword(ctx, tokens[0].UserId).Body(
			openapi.UserSvcResetPasswordRequest{
				NewPassword: "resetPass123",
			}).Execute()

		require.NoError(t, err)
	})

	t.Run("user 1 can login with new password", func(t *testing.T) {
		rsp, _, err := client1.UserSvcAPI.Login(ctx).Body(
			openapi.UserSvcLoginRequest{
				Slug:     openapi.PtrString("test-user-slug-0"),
				Password: openapi.PtrString("resetPass123"),
			}).Execute()

		require.NoError(t, err)
		require.NotEmpty(t, rsp.Token.Token)
	})
}
