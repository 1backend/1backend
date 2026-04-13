package userservice_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/1backend/1backend/sdk/go"
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

	clients, tokens, err := test.MakeClients(clientFactory, sdk.DefaultTestAppHost, 1)
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
			NewAppHost: openapi.PtrString(newApp),
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
			NewAppHost: openapi.PtrString(newApp),
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
			NewAppHost: openapi.PtrString(newApp),
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

func TestExchangeToken_SameAppPreservesDeviceScopedActiveOrganization(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		Test: true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)
	ctx := context.Background()

	clients, _, err := test.MakeClients(clientFactory, sdk.DefaultTestAppHost, 1)
	require.NoError(t, err)
	baseClient := clients[0]

	publicKeyRsp, _, err := clientFactory.Client().
		UserSvcAPI.GetPublicKey(ctx).
		Execute()
	require.NoError(t, err)

	createOrg := func(slug string) string {
		rsp, _, err := baseClient.UserSvcAPI.SaveOrganization(ctx).
			Body(openapi.UserSvcSaveOrganizationRequest{
				Activate: openapi.PtrBool(false),
				Slug:     slug,
				Name:     openapi.PtrString(slug),
			}).
			Execute()
		require.NoError(t, err)
		return rsp.Organization.Id
	}

	orgId1 := createOrg("org-1")
	orgId2 := createOrg("org-2")

	loginWithDevice := func(device string) *openapi.APIClient {
		rsp, _, err := clientFactory.Client().UserSvcAPI.Login(ctx).
			Body(openapi.UserSvcLoginRequest{
				AppHost:  sdk.DefaultTestAppHost,
				Slug:     openapi.PtrString("test-user-slug-0"),
				Password: openapi.PtrString("testUserPassword0"),
				Device:   openapi.PtrString(device),
			}).
			Execute()
		require.NoError(t, err)
		return clientFactory.Client(client.WithToken(rsp.Token.Token))
	}

	deviceA := loginWithDevice("browser-a")
	deviceB := loginWithDevice("browser-b")

	activateA, _, err := deviceA.UserSvcAPI.ActivateOrganization(ctx).
		Body(openapi.UserSvcActivateOrganizationRequest{
			OrganizationId: orgId1,
		}).
		Execute()
	require.NoError(t, err)
	deviceA = clientFactory.Client(client.WithToken(activateA.Token.Token))

	activateB, _, err := deviceB.UserSvcAPI.ActivateOrganization(ctx).
		Body(openapi.UserSvcActivateOrganizationRequest{
			OrganizationId: orgId2,
		}).
		Execute()
	require.NoError(t, err)
	deviceB = clientFactory.Client(client.WithToken(activateB.Token.Token))

	defaultApp := sdk.DefaultTestAppHost

	exchangedA, _, err := deviceA.UserSvcAPI.ExchangeToken(ctx).
		Body(openapi.UserSvcExchangeTokenRequest{
			NewAppHost: openapi.PtrString(defaultApp),
		}).
		Execute()
	require.NoError(t, err)

	exchangedAClaims, err := (auth.AuthorizerImpl{}).ParseJWT(
		publicKeyRsp.PublicKey,
		exchangedA.Token.Token,
	)
	require.NoError(t, err)
	require.Equal(t, "browser-a", exchangedAClaims.Device)
	require.Equal(t, orgId1, exchangedAClaims.ActiveOrganizationId)
	require.Contains(t, exchangedAClaims.Roles, fmt.Sprintf("user-svc:org:{%s}:admin", orgId1))
	require.NotContains(t, exchangedAClaims.Roles, fmt.Sprintf("user-svc:org:{%s}:admin", orgId2))

	exchangedB, _, err := deviceB.UserSvcAPI.ExchangeToken(ctx).
		Body(openapi.UserSvcExchangeTokenRequest{
			NewAppHost: openapi.PtrString(defaultApp),
		}).
		Execute()
	require.NoError(t, err)

	exchangedBClaims, err := (auth.AuthorizerImpl{}).ParseJWT(
		publicKeyRsp.PublicKey,
		exchangedB.Token.Token,
	)
	require.NoError(t, err)
	require.Equal(t, "browser-b", exchangedBClaims.Device)
	require.Equal(t, orgId2, exchangedBClaims.ActiveOrganizationId)
	require.Contains(t, exchangedBClaims.Roles, fmt.Sprintf("user-svc:org:{%s}:admin", orgId2))
	require.NotContains(t, exchangedBClaims.Roles, fmt.Sprintf("user-svc:org:{%s}:admin", orgId1))

	exchangedNewDevice, _, err := deviceA.UserSvcAPI.ExchangeToken(ctx).
		Body(openapi.UserSvcExchangeTokenRequest{
			NewAppHost: openapi.PtrString(defaultApp),
			NewDevice:  openapi.PtrString("browser-c"),
		}).
		Execute()
	require.NoError(t, err)

	exchangedNewDeviceClaims, err := (auth.AuthorizerImpl{}).ParseJWT(
		publicKeyRsp.PublicKey,
		exchangedNewDevice.Token.Token,
	)
	require.NoError(t, err)
	require.Equal(t, "browser-c", exchangedNewDeviceClaims.Device)
	require.Empty(t, exchangedNewDeviceClaims.ActiveOrganizationId)
	require.NotContains(t, exchangedNewDeviceClaims.Roles, fmt.Sprintf("user-svc:org:{%s}:admin", orgId1))
	require.NotContains(t, exchangedNewDeviceClaims.Roles, fmt.Sprintf("user-svc:org:{%s}:admin", orgId2))
}
