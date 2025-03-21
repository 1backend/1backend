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

	clients "github.com/1backend/1backend/clients/go"
)

func TestRegistration(t *testing.T) {
	hs := &di.HandlerSwitcher{}
	server := httptest.NewServer(hs)
	defer server.Close()

	options := &di.Options{
		Test: true,
		Url:  server.URL,
	}
	universe, starterFunc, err := di.BigBang(options)
	require.NoError(t, err)

	hs.UpdateHandler(universe)

	err = starterFunc()
	require.NoError(t, err)

	userSvc := options.ClientFactory.Client().UserSvcAPI

	manyClients, err := test.MakeClients(options.ClientFactory, 1)

	require.NoError(t, err)

	userClient := manyClients[0]
	userToken := userClient.GetConfig().DefaultHeader["Authorization"]
	userToken = strings.Replace(userToken, "Bearer ", "", -1)
	require.Equal(t, true, len(userToken) > 0)

	publicKeyRsp, _, err := userSvc.GetPublicKey(context.Background()).Execute()
	require.NoError(t, err)

	t.Run("user password change", func(t *testing.T) {
		claim, err := options.Authorizer.ParseJWT(
			*publicKeyRsp.PublicKey,
			userToken,
		)
		require.NoError(t, err)

		byTokenRsp, _, err := userClient.UserSvcAPI.ReadUserByToken(context.Background()).
			Execute()
		require.NoError(t, err)

		require.Equal(t, "test-user-slug-0", *byTokenRsp.User.Slug)
		require.True(t, nil == byTokenRsp.User.PasswordHash)

		require.Equal(t, &claim.UserId, byTokenRsp.User.Id)

		changePassReq := clients.UserSvcChangePasswordRequest{
			Slug:            clients.PtrString("test-user-slug-0"),
			CurrentPassword: clients.PtrString("testUserPassword0"),
			NewPassword:     clients.PtrString("yo"),
		}
		_, _, err = userSvc.ChangePassword(context.Background()).
			Body(changePassReq).
			Execute()
		require.Error(t, err)

		_, _, err = userClient.UserSvcAPI.ChangePassword(context.Background()).
			Body(changePassReq).
			Execute()
		require.NoError(t, err)

		// changing with wrong password should error
		changePassReq.CurrentPassword = clients.PtrString("yoWRONG")
		changePassReq.NewPassword = clients.PtrString("yo1")

		_, _, err = userClient.UserSvcAPI.ChangePassword(context.Background()).
			Body(changePassReq).
			Execute()
		require.Error(t, err)
	})
}

func TestOrganization(t *testing.T) {
	hs := &di.HandlerSwitcher{}
	server := httptest.NewServer(hs)
	defer server.Close()

	options := &di.Options{
		Test: true,
		Url:  server.URL,
	}
	universe, starterFunc, err := di.BigBang(options)
	require.NoError(t, err)

	hs.UpdateHandler(universe)
	err = starterFunc()
	require.NoError(t, err)

	manyClients, err := test.MakeClients(options.ClientFactory, 3)
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
			createOrgReq := clients.UserSvcCreateOrganizationRequest{
				Id:   clients.PtrString(orgId1),
				Slug: clients.PtrString("test-org"),
				Name: clients.PtrString("Test Org"),
			}
			_, _, err := userClient.UserSvcAPI.CreateOrganization(context.Background()).
				Body(createOrgReq).
				Execute()
			require.NoError(t, err)

			claim, err := options.Authorizer.ParseJWT(
				*publicKeyRsp.PublicKey,
				userToken,
			)
			require.NoError(t, err)
			require.NotNil(t, claim)
			require.Equal(t, 1, len(claim.RoleIds), claim.RoleIds)

			loginReq := clients.UserSvcLoginRequest{
				Slug:     clients.PtrString("test-user-slug-0"),
				Password: clients.PtrString("testUserPassword0"),
			}
			loginRsp, _, err := userClient.UserSvcAPI.Login(context.Background()).
				Body(loginReq).
				Execute()
			require.NoError(t, err)

			claim, err = options.Authorizer.ParseJWT(
				*publicKeyRsp.PublicKey,
				*loginRsp.Token.Token,
			)
			require.NoError(t, err)
			require.NotNil(t, claim)
			require.Equal(t, 2, len(claim.RoleIds), claim.RoleIds)
			require.Contains(
				t,
				claim.RoleIds,
				fmt.Sprintf("user-svc:org:{%v}:admin", orgId1),
				claim.RoleIds,
			)

			tokenRsp, _, err := userClient.UserSvcAPI.ReadUserByToken(context.Background()).
				Execute()
			require.NoError(t, err)
			require.Equal(t, 1, len(tokenRsp.Organizations))
			require.Equal(
				t,
				clients.PtrString(orgId1),
				tokenRsp.ActiveOrganizationId,
			)
		},
	)

	t.Run("assign org to user", func(t *testing.T) {
		byTokenRsp, _, err := otherClient.UserSvcAPI.ReadUserByToken(context.Background()).
			Execute()
		require.NoError(t, err)

		addUserReq := clients.UserSvcAddUserToOrganizationRequest{
			UserId: byTokenRsp.User.Id,
		}
		_, _, err = userClient.UserSvcAPI.AddUserToOrganization(context.Background(), orgId1).
			Body(addUserReq).
			Execute()
		require.NoError(t, err)

		loginReq := clients.UserSvcLoginRequest{
			Slug:     clients.PtrString("test-user-slug-1"),
			Password: clients.PtrString("testUserPassword1"),
		}
		// log in again and see the claim
		loginRsp, _, err := otherClient.UserSvcAPI.Login(context.Background()).
			Body(loginReq).
			Execute()
		require.NoError(t, err)

		claim, err := options.Authorizer.ParseJWT(
			*publicKeyRsp.PublicKey,
			*loginRsp.Token.Token,
		)

		require.NoError(t, err)
		require.NotNil(t, claim)
		require.Equal(t, 2, len(claim.RoleIds), claim.RoleIds)
		require.Contains(
			t,
			claim.RoleIds,
			fmt.Sprintf("user-svc:org:{%v}:user", orgId1),
		)

		_, _, err = thirdClient.UserSvcAPI.RemoveUserFromOrganization(context.Background(), orgId1, *byTokenRsp.User.Id).
			Execute()
		// third user cannot remove the second from the org of the first
		require.Error(t, err)

		_, _, err = userClient.UserSvcAPI.RemoveUserFromOrganization(context.Background(), orgId1, *byTokenRsp.User.Id).
			Execute()
		require.NoError(t, err)
	})
}
