package userservice_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/test"
)

func TestRefreshToken(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		Test: true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)

	manyClients, tokens, err := test.MakeClients(clientFactory, 1)
	require.NoError(t, err)

	client1 := manyClients[0]

	publicKeyRsp, _, err := clientFactory.Client().
		UserSvcAPI.GetPublicKey(context.Background()).
		Execute()
	require.NoError(t, err)

	originalClaim, err := (auth.AuthorizerImpl{}).ParseJWT(
		publicKeyRsp.PublicKey,
		tokens[0].Token,
	)
	require.NoError(t, err)

	t.Run("refresh token", func(t *testing.T) {
		// Seems like jwt.NumericDate is only accurate to the second
		// so we need to sleep here.
		time.Sleep(1 * time.Second)

		rsp, hrsp, err := client1.UserSvcAPI.RefreshToken(
			context.Background(),
		).Execute()

		require.NoError(t, err, hrsp)
		require.NotEmpty(t, rsp.Token.Token)

		newClaim, err := (auth.AuthorizerImpl{}).ParseJWT(
			publicKeyRsp.PublicKey,
			rsp.Token.Token,
		)
		require.NoError(t, err)
		require.Equal(t, true, originalClaim.ExpiresAt.Before(newClaim.ExpiresAt.Time))
	})
}
