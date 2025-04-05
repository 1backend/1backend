package userservice_test

import (
	"context"
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/test"
	"github.com/1backend/1backend/server/internal/di"

	openapi "github.com/1backend/1backend/clients/go"
)

func TestPasswordChange(t *testing.T) {
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

	userSvc := options.ClientFactory.Client().UserSvcAPI

	manyClients, _, err := test.MakeClients(options.ClientFactory, 1)

	require.NoError(t, err)

	userClient := manyClients[0]
	userToken := auth.TokenFromClient(userClient)

	require.Equal(t, true, len(userToken) > 0)

	publicKeyRsp, _, err := userSvc.GetPublicKey(context.Background()).Execute()
	require.NoError(t, err)

	t.Run("user password change", func(t *testing.T) {
		claim, err := options.Authorizer.ParseJWT(
			publicKeyRsp.PublicKey,
			userToken,
		)
		require.NoError(t, err)

		byTokenRsp, _, err := userClient.UserSvcAPI.ReadUserByToken(context.Background()).
			Execute()
		require.NoError(t, err)

		require.Equal(t, "test-user-slug-0", byTokenRsp.User.Slug)
		require.True(t, nil == byTokenRsp.User.PasswordHash)

		require.Equal(t, claim.UserId, byTokenRsp.User.Id)

		changePassReq := openapi.UserSvcChangePasswordRequest{
			Slug:            openapi.PtrString("test-user-slug-0"),
			CurrentPassword: openapi.PtrString("testUserPassword0"),
			NewPassword:     openapi.PtrString("yo"),
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
		changePassReq.CurrentPassword = openapi.PtrString("yoWRONG")
		changePassReq.NewPassword = openapi.PtrString("yo1")

		_, _, err = userClient.UserSvcAPI.ChangePassword(context.Background()).
			Body(changePassReq).
			Execute()
		require.Error(t, err)
	})
}

func TestRegistration(t *testing.T) {
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

	userSvc := options.ClientFactory.Client().UserSvcAPI
	_, _, err = userSvc.Register(context.Background()).Body(
		openapi.UserSvcRegisterRequest{
			Slug: "test-1",
			Contact: &openapi.UserSvcContact{
				Id:       "test1@test.comm",
				Platform: "email",
			},
			Password: openapi.PtrString("test"),
		},
	).Execute()
	require.NoError(t, err)

	t.Run("slug login works", func(t *testing.T) {
		loginReq := openapi.UserSvcLoginRequest{
			Slug:     openapi.PtrString("test-1"),
			Password: openapi.PtrString("test"),
		}
		_, _, err := options.ClientFactory.Client().UserSvcAPI.Login(context.Background()).
			Body(loginReq).
			Execute()
		require.NoError(t, err)
	})

	t.Run("contact login works", func(t *testing.T) {
		loginReq := openapi.UserSvcLoginRequest{
			Contact:  openapi.PtrString("test1@test.comm"),
			Password: openapi.PtrString("test"),
		}
		_, _, err := options.ClientFactory.Client().UserSvcAPI.Login(context.Background()).
			Body(loginReq).
			Execute()
		require.NoError(t, err)
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

			tokenRsp, _, err := userClient.UserSvcAPI.ReadUserByToken(context.Background()).
				Execute()
			require.NoError(t, err)
			require.Equal(t, 1, len(tokenRsp.Organizations))
			require.Equal(
				t,
				openapi.PtrString(orgId1),
				tokenRsp.ActiveOrganizationId,
			)
		},
	)

	t.Run("assign org to user", func(t *testing.T) {
		byTokenRsp, _, err := otherClient.UserSvcAPI.ReadUserByToken(context.Background()).
			Execute()
		require.NoError(t, err)

		_, _, err = userClient.UserSvcAPI.AddUserToOrganization(
			context.Background(),
			orgId1,
			byTokenRsp.User.Id,
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
		require.Contains(
			t,
			claim.Roles,
			fmt.Sprintf("user-svc:org:{%v}:user", orgId1),
		)

		_, _, err = thirdClient.UserSvcAPI.RemoveUserFromOrganization(
			context.Background(),
			orgId1,
			byTokenRsp.User.Id,
		).
			Execute()
		// third user cannot remove the second from the org of the first
		require.Error(t, err)

		_, _, err = userClient.UserSvcAPI.RemoveUserFromOrganization(
			context.Background(),
			orgId1,
			byTokenRsp.User.Id,
		).
			Execute()
		require.NoError(t, err)
	})
}
