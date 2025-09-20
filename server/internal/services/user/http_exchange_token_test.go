package userservice_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/test"

	openapi "github.com/1backend/1backend/clients/go"
)

func TestExchangeToken(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		Test: true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)

	clients, tokens, err := test.MakeClients(clientFactory, "test", 1)
	require.NoError(t, err)

	client1 := clients[0]
	originalToken := tokens[0].Token

	publicKeyRsp, _, err := clientFactory.Client().
		UserSvcAPI.GetPublicKey(context.Background()).
		Execute()
	require.NoError(t, err)

	originalClaim, err := (auth.AuthorizerImpl{}).ParseJWT(
		publicKeyRsp.PublicKey,
		originalToken,
	)
	require.NoError(t, err)

	newApp := "some-other-app"

	t.Run("exchange", func(t *testing.T) {
		exchangeReq := openapi.UserSvcExchangeTokenRequest{
			NewAppHost: newApp,
		}

		rsp, httpRsp, err := client1.UserSvcAPI.ExchangeToken(context.Background()).
			Body(exchangeReq).
			Execute()

		require.NoError(t, err, httpRsp)
		require.NotEmpty(t, rsp.Token.Token)

		exchangedClaim, err := (auth.AuthorizerImpl{}).ParseJWT(
			publicKeyRsp.PublicKey,
			rsp.Token.Token,
		)
		require.NoError(t, err)

		require.Equal(t, originalClaim.UserId, exchangedClaim.UserId)
		require.Equal(t, originalClaim.Device, exchangedClaim.Device)
		require.Equal(t, newApp, rsp.Token.App.Host)
	})

	newToken := ""

	t.Run("exchange new device", func(t *testing.T) {
		exchangeReq := openapi.UserSvcExchangeTokenRequest{
			NewAppHost: newApp,
			NewDevice:  openapi.PtrString("new-device"),
		}

		rsp, httpRsp, err := client1.UserSvcAPI.ExchangeToken(context.Background()).
			Body(exchangeReq).
			Execute()

		require.NoError(t, err, httpRsp)
		require.NotEmpty(t, rsp.Token.Token)

		exchangedClaim, err := (auth.AuthorizerImpl{}).ParseJWT(
			publicKeyRsp.PublicKey,
			rsp.Token.Token,
		)
		require.NoError(t, err)

		require.Equal(t, originalClaim.UserId, exchangedClaim.UserId)
		require.Equal(t, "new-device", exchangedClaim.Device)
		require.Equal(t, newApp, rsp.Token.App.Host)

		newToken = rsp.Token.Token
	})

	t.Run("exchange, device not specified", func(t *testing.T) {
		exchangeReq := openapi.UserSvcExchangeTokenRequest{
			NewAppHost: newApp,
		}

		rsp, httpRsp, err := clientFactory.Client(client.WithToken(newToken)).
			UserSvcAPI.ExchangeToken(context.Background()).
			Body(exchangeReq).
			Execute()

		require.NoError(t, err, httpRsp)
		require.NotEmpty(t, rsp.Token.Token)

		exchangedClaim, err := (auth.AuthorizerImpl{}).ParseJWT(
			publicKeyRsp.PublicKey,
			rsp.Token.Token,
		)
		require.NoError(t, err)

		require.Equal(t, "new-device", exchangedClaim.Device)
		require.Equal(t, newApp, rsp.Token.App.Host)
	})
}
