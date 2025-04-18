package userservice_test

import (
	"context"
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/test"
	"github.com/1backend/1backend/server/internal/di"

	openapi "github.com/1backend/1backend/clients/go"
)

func TestOrganization(t *testing.T) {
	hs := &di.HandlerSwitcher{}
	server := httptest.NewServer(hs)
	defer server.Close()

	options := &di.Options{
		Test: true,
		Url:  server.URL,
	}
	universe, err := di.BigBang(options)
	require.NoError(t, err)

	hs.UpdateHandler(universe.Router)
	err = universe.StarterFunc()
	require.NoError(t, err)

	manyClients, _, err := test.MakeClients(options.ClientFactory, 3)
	require.NoError(t, err)

	userClient := manyClients[0]

	userToken := userClient.GetConfig().DefaultHeader["Authorization"]
	userToken = strings.Replace(userToken, "Bearer ", "", -1)
	require.Equal(t, true, len(userToken) > 0)

	otherClient := manyClients[1]
	thirdClient := manyClients[2]

	publicKeyRsp, _, err := options.ClientFactory.Client().
		UserSvcAPI.GetPublicKey(context.Background()).
		Execute()
	require.NoError(t, err)

	orgId1 := sdk.Id("org")

	t.Run(
		"claim contains new organization admin role after creating organization",
		func(t *testing.T) {
			createOrgReq := openapi.UserSvcSaveOrganizationRequest{
				Id:   openapi.PtrString(orgId1),
				Slug: "test-org",
				Name: openapi.PtrString("Test Org"),
			}
			_, rsp, err := userClient.UserSvcAPI.SaveOrganization(context.Background()).
				Body(createOrgReq).
				Execute()
			require.NoError(t, err, rsp)

			claim, err := options.Authorizer.ParseJWT(
				publicKeyRsp.PublicKey,
				userToken,
			)
			require.NoError(t, err)
			require.NotNil(t, claim)
			require.Equal(t, 1, len(claim.Roles), claim.Roles)

			loginReq := openapi.UserSvcLoginRequest{
				Slug:     openapi.PtrString("test-user-slug-0"),
				Password: openapi.PtrString("testUserPassword0"),
			}
			loginRsp, _, err := userClient.UserSvcAPI.Login(context.Background()).
				Body(loginReq).
				Execute()
			require.NoError(t, err)

			claim, err = options.Authorizer.ParseJWT(
				publicKeyRsp.PublicKey,
				loginRsp.Token.Token,
			)
			require.NoError(t, err)
			require.NotNil(t, claim)
			require.Equal(t, 2, len(claim.Roles), claim.Roles)
			require.Contains(
				t,
				claim.Roles,
				fmt.Sprintf("user-svc:org:{%v}:admin", orgId1),
				claim.Roles,
			)
			require.Equal(t, orgId1, claim.ActiveOrganizationId)

			tokenRsp, _, err := userClient.UserSvcAPI.ReadSelf(context.Background()).
				Execute()
			require.NoError(t, err)
			require.Equal(t, 1, len(tokenRsp.Organizations))
			require.Equal(
				t,
				openapi.PtrString(orgId1),
				tokenRsp.ActiveOrganizationId,
			)
			require.NotEmpty(t, tokenRsp.Organizations[0].Slug)
			require.NotEmpty(t, tokenRsp.Organizations[0].Name)
			require.NotEmpty(t, tokenRsp.Roles)
		},
	)

	t.Run("assign org to user", func(t *testing.T) {
		readSelfRsp, _, err := otherClient.UserSvcAPI.ReadSelf(context.Background()).
			Execute()
		require.NoError(t, err)

		_, _, err = userClient.UserSvcAPI.AddUserToOrganization(
			context.Background(),
			orgId1,
			readSelfRsp.User.Id,
		).
			Execute()
		require.NoError(t, err)

		loginReq := openapi.UserSvcLoginRequest{
			Slug:     openapi.PtrString("test-user-slug-1"),
			Password: openapi.PtrString("testUserPassword1"),
		}
		// log in again and see the claim
		loginRsp, _, err := otherClient.UserSvcAPI.Login(context.Background()).
			Body(loginReq).
			Execute()
		require.NoError(t, err)

		claim, err := options.Authorizer.ParseJWT(
			publicKeyRsp.PublicKey,
			loginRsp.Token.Token,
		)

		require.NoError(t, err)
		require.NotNil(t, claim)
		require.Equal(t, 2, len(claim.Roles), claim.Roles)
		require.Equal(t, orgId1, claim.ActiveOrganizationId)
		require.Contains(
			t,
			claim.Roles,
			fmt.Sprintf("user-svc:org:{%v}:user", orgId1),
		)

		_, _, err = thirdClient.UserSvcAPI.RemoveUserFromOrganization(
			context.Background(),
			orgId1,
			readSelfRsp.User.Id,
		).
			Execute()
		// third user cannot remove the second from the org of the first
		require.Error(t, err)

		_, _, err = userClient.UserSvcAPI.RemoveUserFromOrganization(
			context.Background(),
			orgId1,
			readSelfRsp.User.Id,
		).
			Execute()
		require.NoError(t, err)
	})
}
