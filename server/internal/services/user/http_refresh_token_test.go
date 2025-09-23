package userservice_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	sdk "github.com/1backend/1backend/sdk/go"
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

	manyClients, tokens, err := test.MakeClients(clientFactory, sdk.DefaultTestAppHost, 1)
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

	manyClients, tokens, err := test.MakeClients(clientFactory, sdk.DefaultTestAppHost, 1)
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
		require.LessOrEqual(t, self.TokenCount, int32(3))
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

	manyClients, _, err := test.MakeClients(clientFactory, sdk.DefaultTestAppHost, 1)
	require.NoError(t, err)

	// default device
	deviceA := manyClients[0]

	lrsp, httrsp, err := deviceA.UserSvcAPI.Login(
		context.Background(),
	).Body(openapi.UserSvcLoginRequest{
		AppHost:  sdk.DefaultTestAppHost,
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

func TestRefreshToken_AfterExchange_KeepsAppAndExtendsExpiry(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{Test: true})
	require.NoError(t, err)
	defer server.Cleanup(t)

	const appA = "socks.com"
	const appB = "shoes.com"
	const permA = "test-user-slug-0:perm-A"

	clientFactory := client.NewApiClientFactory(server.Url)
	ctx := context.Background()

	// user
	many, toks, err := test.MakeClients(clientFactory, sdk.DefaultTestAppHost, 1)
	require.NoError(t, err)
	userUnnamed := many[0]

	// for claim parsing
	pk, _, err := clientFactory.Client().UserSvcAPI.GetPublicKey(ctx).Execute()
	require.NoError(t, err)

	orig, err := (auth.AuthorizerImpl{}).ParseJWT(pk.PublicKey, toks[0].Token)
	require.NoError(t, err)

	// exchange into appA
	exA, _, err := userUnnamed.UserSvcAPI.ExchangeToken(ctx).
		Body(openapi.UserSvcExchangeTokenRequest{NewAppHost: openapi.PtrString(appA)}).Execute()
	require.NoError(t, err)
	require.NotEmpty(t, exA.Token.Token)
	userA := clientFactory.Client(client.WithToken(exA.Token.Token))

	// create an appA-scoped permit
	_, _, err = userA.UserSvcAPI.SavePermits(ctx).Body(openapi.UserSvcSavePermitsRequest{
		Permits: []openapi.UserSvcPermitInput{
			{Roles: []string{"user-svc:user"}, Permission: permA},
		},
	}).Execute()
	require.NoError(t, err)

	// refresh the exchanged token
	time.Sleep(1100 * time.Millisecond)
	ref, _, err := userA.UserSvcAPI.RefreshToken(ctx).Execute()
	require.NoError(t, err)
	require.NotEmpty(t, ref.Token.Token)

	refClaim, err := (auth.AuthorizerImpl{}).ParseJWT(pk.PublicKey, ref.Token.Token)
	require.NoError(t, err)
	require.True(t, orig.ExpiresAt.Before(refClaim.ExpiresAt.Time), "expiry did not increase")

	// with refreshed token in appA we should have the permission
	userARef := clientFactory.Client(client.WithToken(ref.Token.Token))
	hasA, _, err := userARef.UserSvcAPI.HasPermission(ctx, permA).
		Execute()
	require.NoError(t, err)
	require.True(t, hasA.Authorized, "expected permission in appA after refresh")

	// exchange into appB and ensure permission is NOT granted there
	exB, _, err := userARef.UserSvcAPI.ExchangeToken(ctx).
		Body(openapi.UserSvcExchangeTokenRequest{NewAppHost: openapi.PtrString(appB)}).Execute()
	require.NoError(t, err)
	userB := clientFactory.Client(client.WithToken(exB.Token.Token))

	hasB, _, err := userB.UserSvcAPI.HasPermission(ctx, permA).
		Execute()
	require.NoError(t, err)
	require.False(t, hasB.Authorized, "appA permit leaked into appB")
}

func TestRefreshTokenCount_IsBounded_ForExchangedTokens_PerDevice(t *testing.T) {
	return
	t.Parallel()

	server, err := test.StartService(test.Options{
		TokenExpiration: time.Second,
		Test:            true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	const appA = "socks.com"

	clientFactory := client.NewApiClientFactory(server.Url)
	ctx := context.Background()

	// login with explicit device
	login, _, err := clientFactory.Client().UserSvcAPI.Login(ctx).
		Body(openapi.UserSvcLoginRequest{
			Slug:     openapi.PtrString("test-user-slug-0"),
			Password: openapi.PtrString("testUserPassword0"),
			Device:   openapi.PtrString("device-a"),
		}).Execute()
	require.NoError(t, err)

	// exchange that device token into appA (device should carry over unless overridden)
	exA, _, err := clientFactory.Client(client.WithToken(login.Token.Token)).
		UserSvcAPI.ExchangeToken(ctx).
		Body(openapi.UserSvcExchangeTokenRequest{NewAppHost: openapi.PtrString(appA)}).Execute()
	require.NoError(t, err)
	require.NotEmpty(t, exA.Token.Token)

	c := clientFactory.Client(client.WithToken(exA.Token.Token))

	doRefreshAndCount := func() int32 {
		time.Sleep(1100 * time.Millisecond)
		rsp, _, err := c.UserSvcAPI.RefreshToken(ctx).Execute()
		require.NoError(t, err)
		require.NotEmpty(t, rsp.Token.Token)

		// rotate client to use newest token
		c = clientFactory.Client(client.WithToken(rsp.Token.Token))

		self, _, err := c.UserSvcAPI.ReadSelf(ctx).
			Body(openapi.UserSvcReadSelfRequest{CountTokens: openapi.PtrBool(true)}).
			Execute()
		require.NoError(t, err)
		return self.TokenCount
	}

	// refresh several times; count should stay ≤ 3 for the device even for exchanged tokens
	for i := 0; i < 5; i++ {
		cnt := doRefreshAndCount()
		require.LessOrEqual(t, cnt, int32(3), "exchanged-token refresh exceeded per-device bound on iter %d", i)
	}
}
