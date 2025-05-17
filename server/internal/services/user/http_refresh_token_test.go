package userservice_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/test"

	openapi "github.com/1backend/1backend/clients/go"
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

func TestRefreshTokenCountIsBounded(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		TokenExpiration: time.Second,
		Test:            true,
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

	clientFactory.Client().
		UserSvcAPI.GetPublicKey(context.Background()).
		Execute()
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

		self, _, err := client1.UserSvcAPI.ReadSelf(
			context.Background(),
		).Body(
			openapi.UserSvcReadSelfRequest{
				CountTokens: openapi.PtrBool(true),
			},
		).Execute()
		require.NoError(t, err)
		require.LessOrEqual(t, int32(3), self.TokenCount)
	})

	t.Run("refresh token again", func(t *testing.T) {
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

		self, _, err := client1.UserSvcAPI.ReadSelf(
			context.Background(),
		).Body(
			openapi.UserSvcReadSelfRequest{
				CountTokens: openapi.PtrBool(true),
			},
		).Execute()
		require.NoError(t, err)
		require.LessOrEqual(t, self.TokenCount, int32(3))
	})
}

func TestRefreshTokenCountIsBoundedPerDevice(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		TokenExpiration: time.Second,
		Test:            true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)

	manyClients, _, err := test.MakeClients(clientFactory, 1)
	require.NoError(t, err)

	// default device
	deviceA := manyClients[0]

	lrsp, httrsp, err := deviceA.UserSvcAPI.Login(
		context.Background(),
	).Body(openapi.UserSvcLoginRequest{
		Slug:     openapi.PtrString("test-user-slug-0"),
		Password: openapi.PtrString("testUserPassword0"),
		Device:   openapi.PtrString("device-a"),
	}).Execute()
	require.NoError(t, err, httrsp)

	deviceB := clientFactory.Client(
		client.WithToken(lrsp.Token.Token),
	)

	doRefreshAndCount := func(c openapi.UserSvcAPI) int32 {
		time.Sleep(1100 * time.Millisecond)

		rsp, _, err := c.RefreshToken(context.Background()).Execute()
		require.NoError(t, err)
		require.NotEmpty(t, rsp.Token.Token)

		self, _, err := c.ReadSelf(context.Background()).
			Body(openapi.UserSvcReadSelfRequest{CountTokens: openapi.PtrBool(true)}).
			Execute()
		require.NoError(t, err)
		return self.TokenCount
	}

	// make sure 2x3 (3 for each device) tokens are the max

	// Device A: do 4 refreshes, ensure count never exceeds 6
	for i := 0; i < 4; i++ {
		cnt := doRefreshAndCount(deviceA.UserSvcAPI)
		require.LessOrEqual(t, cnt, int32(6), "device A exceeded token bound on iteration %d", i)
	}

	// Device B: do 4 refreshes, ensure count never exceeds 6
	for i := 0; i < 4; i++ {
		cnt := doRefreshAndCount(deviceB.UserSvcAPI)
		require.LessOrEqual(t, cnt, int32(6), "device B exceeded token bound on iteration %d", i)
	}

	finalA, _, err := deviceA.UserSvcAPI.ReadSelf(context.Background()).
		Body(openapi.UserSvcReadSelfRequest{CountTokens: openapi.PtrBool(true)}).
		Execute()
	require.NoError(t, err)
	require.LessOrEqual(t, finalA.TokenCount, int32(6), "test ended with wrong count")
}
