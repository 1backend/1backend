package userservice_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	openapi "github.com/1backend/1backend/clients/go"
	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/test"
)

func TestRevokeTokens(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		TokenExpiration: 1 * time.Second,
		Test:            true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)

	manyClients, _, err := test.MakeClients(clientFactory, sdk.DefaultTestAppHost, 3)
	require.NoError(t, err)

	client1 := manyClients[0]
	client2 := manyClients[1]
	client3 := manyClients[2]

	t.Run("revoke nonexistent device tokens", func(t *testing.T) {
		t.Run("revoking non-existent device throws no error", func(t *testing.T) {
			_, hrsp, err := client1.UserSvcAPI.RevokeTokens(
				context.Background(),
			).Body(
				openapi.UserSvcRevokeTokensRequest{
					Device: openapi.PtrString("nonexistent-device"),
				},
			).Execute()

			require.NoError(t, err, hrsp)
		})

		t.Run("revoking non-existent device should not affect other tokens", func(t *testing.T) {
			rsp, hrsp, err := client1.UserSvcAPI.
				ReadSelf(context.Background()).
				Execute()

			require.NoError(t, err, hrsp)
			require.NotEmpty(t, rsp.User)
		})

		t.Run("revoking token succeeds", func(t *testing.T) {
			_, hrsp, err := client1.UserSvcAPI.RevokeTokens(
				context.Background(),
			).Body(
				openapi.UserSvcRevokeTokensRequest{
					Device: openapi.PtrString("unknown"),
				},
			).Execute()

			require.NoError(t, err, hrsp)
		})

		t.Run("after JWT expiration read self should fail", func(t *testing.T) {
			time.Sleep(1 * time.Second)

			_, hrsp, err := client1.UserSvcAPI.
				ReadSelf(context.Background()).
				Execute()

			require.Error(t, err, hrsp)
		})
	})

	t.Run("client 2 should still have token", func(t *testing.T) {
		_, hrsp, err := client2.UserSvcAPI.
			ReadSelf(context.Background()).
			Execute()

		require.NoError(t, err, hrsp)
	})

	adminClient, _, err := test.AdminClient(clientFactory, sdk.DefaultTestAppHost)
	require.NoError(t, err)

	t.Run("non-admin cannot revoke all tokens", func(t *testing.T) {
		_, hrsp, err := client3.UserSvcAPI.RevokeTokens(
			context.Background(),
		).Body(
			openapi.UserSvcRevokeTokensRequest{
				AllTokens: openapi.PtrBool(true),
			},
		).Execute()

		require.Error(t, err, hrsp)
	})

	t.Run("admin can revoke all tokens", func(t *testing.T) {
		_, hrsp, err := adminClient.UserSvcAPI.RevokeTokens(
			context.Background(),
		).Body(
			openapi.UserSvcRevokeTokensRequest{
				AllTokens: openapi.PtrBool(true),
			},
		).Execute()

		require.NoError(t, err, hrsp)
	})

	// If this test fails, that's a good indication that token refreshing is not being cached.
	// Ie. Since we already had sleep in this test, the token is already expired.
	// But the old expired token should be mapped to a new one on the backend without a token DB lookup.
	// If that caching is not working, this test will fail.
	t.Run("client 2 should be able to read before expiration", func(t *testing.T) {
		rsp, hrsp, err := client2.UserSvcAPI.
			ReadSelf(context.Background()).
			Execute()

		require.NoError(t, err, hrsp)
		require.NotEmpty(t, rsp.User)
	})

	t.Run("client 3 should be able to read before expiration", func(t *testing.T) {
		_, hrsp, err := client3.UserSvcAPI.
			ReadSelf(context.Background()).
			Execute()

		require.NoError(t, err, hrsp)
	})

	time.Sleep(1 * time.Second)

	t.Run("client 2 should not be able to read after expiration", func(t *testing.T) {
		_, hrsp, err := client2.UserSvcAPI.
			ReadSelf(context.Background()).
			Execute()

		require.Error(t, err, hrsp)
	})

	t.Run("client 3 should not be able to read after expiration", func(t *testing.T) {
		_, hrsp, err := client3.UserSvcAPI.
			ReadSelf(context.Background()).
			Execute()

		require.Error(t, err, hrsp)
	})
}
