package userservice_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	openapi "github.com/1backend/1backend/clients/go"
	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/test"
)

func TestActivateOrganization(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		Test: true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)

	clients, _, err := test.MakeClients(clientFactory, sdk.DefaultTestAppHost, 2)
	require.NoError(t, err)

	userClient := clients[0]
	otherClient := clients[1]

	publicKeyRsp, _, err := clientFactory.Client().
		UserSvcAPI.GetPublicKey(context.Background()).
		Execute()
	require.NoError(t, err)

	createOrg := func(name, slug string, token string) string {
		rsp, _, err := clientFactory.Client(client.WithToken(token)).
			UserSvcAPI.SaveOrganization(context.Background()).
			Body(openapi.UserSvcSaveOrganizationRequest{
				Activate: openapi.PtrBool(false),
				Name: openapi.PtrString(name),
				Slug: slug,
			}).
			Execute()
		require.NoError(t, err)

		return rsp.Organization.Id
	}

	activate := func(token string, organizationId string) (*openapi.UserSvcActivateOrganizationResponse, *http.Response, error) {
		return clientFactory.Client(client.WithToken(token)).
			UserSvcAPI.ActivateOrganization(context.Background()).
			Body(openapi.UserSvcActivateOrganizationRequest{
				OrganizationId: organizationId,
			}).
			Execute()
	}

	loginRsp, _, err := userClient.UserSvcAPI.Login(context.Background()).
		Body(openapi.UserSvcLoginRequest{
			AppHost:  sdk.DefaultTestAppHost,
			Slug:     openapi.PtrString("test-user-slug-0"),
			Password: openapi.PtrString("testUserPassword0"),
		}).
		Execute()
	require.NoError(t, err)

	orgId1 := createOrg("Org 1", "org-1", loginRsp.Token.Token)
	orgId2 := createOrg("Org 2", "org-2", loginRsp.Token.Token)

	initialActivateRsp, _, err := activate(loginRsp.Token.Token, orgId2)
	require.NoError(t, err)
	currentToken := initialActivateRsp.Token.Token

	t.Run("switch active organization and mint a fresh token", func(t *testing.T) {
		activateRsp, httpResp, err := activate(currentToken, orgId1)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, httpResp.StatusCode)
		require.NotNil(t, activateRsp)
		require.NotEmpty(t, activateRsp.Token.Token)

		newToken := activateRsp.Token.Token

		claims, err := (auth.AuthorizerImpl{}).ParseJWT(
			publicKeyRsp.PublicKey,
			newToken,
		)
		require.NoError(t, err)
		require.Equal(t, orgId1, claims.ActiveOrganizationId)
		require.Contains(t, claims.Roles, fmt.Sprintf("user-svc:org:{%s}:admin", orgId1))
		require.NotContains(t, claims.Roles, fmt.Sprintf("user-svc:org:{%s}:admin", orgId2))

		readSelfRsp, _, err := clientFactory.Client(client.WithToken(newToken)).
			UserSvcAPI.ReadSelf(context.Background()).
			Execute()
		require.NoError(t, err)
		require.Equal(t, orgId1, readSelfRsp.GetActiveOrganizationId())
		require.Len(t, readSelfRsp.Organizations, 2)
	})

	t.Run("keep current active organization unchanged on invalid target", func(t *testing.T) {
		_, httpResp, err := activate(currentToken, "org-does-not-exist")
		require.Error(t, err)
		require.Equal(t, http.StatusUnauthorized, httpResp.StatusCode)

		readSelfRsp, _, err := clientFactory.Client(client.WithToken(currentToken)).
			UserSvcAPI.ReadSelf(context.Background()).
			Execute()
		require.NoError(t, err)
		require.Equal(t, orgId1, readSelfRsp.GetActiveOrganizationId())
	})

	t.Run("reject switching to an organization the caller is not a member of", func(t *testing.T) {
		otherLoginRsp, _, err := otherClient.UserSvcAPI.Login(context.Background()).
			Body(openapi.UserSvcLoginRequest{
				AppHost:  sdk.DefaultTestAppHost,
				Slug:     openapi.PtrString("test-user-slug-1"),
				Password: openapi.PtrString("testUserPassword1"),
			}).
			Execute()
		require.NoError(t, err)

		_, httpResp, err := activate(otherLoginRsp.Token.Token, orgId2)
		require.Error(t, err)
		require.Equal(t, http.StatusUnauthorized, httpResp.StatusCode)
	})

	t.Run("keep active organization independent per device", func(t *testing.T) {
		deviceA := "browser-a"
		deviceB := "browser-b"

		loginWithDevice := func(device string) string {
			loginRsp, _, err := userClient.UserSvcAPI.Login(context.Background()).
				Body(openapi.UserSvcLoginRequest{
					AppHost:  sdk.DefaultTestAppHost,
					Slug:     openapi.PtrString("test-user-slug-0"),
					Password: openapi.PtrString("testUserPassword0"),
					Device:   openapi.PtrString(device),
				}).
				Execute()
			require.NoError(t, err)
			return loginRsp.Token.Token
		}

		tokenA := loginWithDevice(deviceA)
		tokenB := loginWithDevice(deviceB)

		activateRspA, httpRespA, err := activate(tokenA, orgId1)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, httpRespA.StatusCode)
		require.NotNil(t, activateRspA)
		require.NotEmpty(t, activateRspA.Token.Token)
		newTokenA := activateRspA.Token.Token

		activateRspB, httpRespB, err := activate(tokenB, orgId2)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, httpRespB.StatusCode)
		require.NotNil(t, activateRspB)
		require.NotEmpty(t, activateRspB.Token.Token)
		newTokenB := activateRspB.Token.Token

		claimsA, err := (auth.AuthorizerImpl{}).ParseJWT(publicKeyRsp.PublicKey, newTokenA)
		require.NoError(t, err)
		require.Equal(t, orgId1, claimsA.ActiveOrganizationId)
		require.Contains(t, claimsA.Roles, fmt.Sprintf("user-svc:org:{%s}:admin", orgId1))
		require.NotContains(t, claimsA.Roles, fmt.Sprintf("user-svc:org:{%s}:admin", orgId2))

		claimsB, err := (auth.AuthorizerImpl{}).ParseJWT(publicKeyRsp.PublicKey, newTokenB)
		require.NoError(t, err)
		require.Equal(t, orgId2, claimsB.ActiveOrganizationId)
		require.Contains(t, claimsB.Roles, fmt.Sprintf("user-svc:org:{%s}:admin", orgId2))
		require.NotContains(t, claimsB.Roles, fmt.Sprintf("user-svc:org:{%s}:admin", orgId1))
	})
}
