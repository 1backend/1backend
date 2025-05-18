package userservice_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	openapi "github.com/1backend/1backend/clients/go"
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

	manyClients, _, err := test.MakeClients(clientFactory, 2)
	require.NoError(t, err)

	client1 := manyClients[0]
	client2 := manyClients[1]

	t.Run("revoke nonexisten device tokens", func(t *testing.T) {
		_, hrsp, err := client1.UserSvcAPI.RevokeTokens(
			context.Background(),
		).Body(
			openapi.UserSvcRevokeTokensRequest{
				Device: openapi.PtrString("nonexistent-device"),
			},
		).Execute()
		require.NoError(t, err, hrsp)

		rsp, hrsp, err := client1.UserSvcAPI.
			ReadSelf(context.Background()).
			Execute()
		require.NoError(t, err, hrsp)
		require.NotEmpty(t, rsp.User)

		_, hrsp, err = client1.UserSvcAPI.RevokeTokens(
			context.Background(),
		).Body(
			openapi.UserSvcRevokeTokensRequest{
				Device: openapi.PtrString("default"),
			},
		).Execute()
		require.NoError(t, err, hrsp)

		time.Sleep(1 * time.Second)

		_, _, err = client1.UserSvcAPI.
			ReadSelf(context.Background()).
			Execute()
		require.Error(t, err, hrsp)
	})

	t.Run("client 2 should still have token", func(t *testing.T) {
		_, hrsp, err := client2.UserSvcAPI.
			ReadSelf(context.Background()).
			Execute()
		require.NoError(t, err, hrsp)
	})
}
