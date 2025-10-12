package userservice_test

import (
	"context"
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/test"
	"github.com/1backend/1backend/server/internal/di"
	"github.com/1backend/1backend/server/internal/universe"

	openapi "github.com/1backend/1backend/clients/go"
)

func TestLogin__IncorrectCredentials(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		TokenExpiration: time.Second,
		Test:            true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)

	ctx := context.Background()
	_, _, err = test.MakeClients(clientFactory, sdk.DefaultTestAppHost, 1)
	require.NoError(t, err)

	time.Sleep(1 * time.Second)

	t.Run("no password", func(t *testing.T) {
		_, _, err := clientFactory.Client().UserSvcAPI.Login(ctx).Body(
			openapi.UserSvcLoginRequest{
				AppHost:  sdk.DefaultTestAppHost,
				Slug:     openapi.PtrString("test-user-slug-0"),
				Password: openapi.PtrString(""),
			},
		).Execute()
		require.Error(t, err)
	})

	t.Run("wrong password", func(t *testing.T) {
		_, _, err := clientFactory.Client().UserSvcAPI.Login(ctx).Body(
			openapi.UserSvcLoginRequest{
				AppHost:  sdk.DefaultTestAppHost,
				Slug:     openapi.PtrString("test-user-slug-0"),
				Password: openapi.PtrString("wrongPassword"),
			},
		).Execute()
		require.Error(t, err)
	})
}

func TestLogin__OnlyOtp(t *testing.T) {
	hs := &di.HandlerSwitcher{}
	server := httptest.NewServer(hs)
	defer server.Close()

	options := &universe.Options{
		Test:           true,
		Url:            server.URL,
		VerifyContacts: true,
	}
	universe, err := di.BigBang(options)
	require.NoError(t, err)

	hs.UpdateHandler(universe.Router)

	err = universe.StarterFunc()
	require.NoError(t, err)

	userSvc := options.ClientFactory.Client().UserSvcAPI

	otpRsp, _, err := userSvc.SendOtp(context.Background()).Body(
		openapi.UserSvcSendOtpRequest{
			AppHost:         sdk.DefaultTestAppHost,
			ContactId:       "test2@test.com",
			ContactPlatform: "email",
		},
	).Execute()
	require.NoError(t, err, otpRsp)
	require.NotNil(t, otpRsp.OtpId)
	require.NotNil(t, otpRsp.Code)

	t.Run("registration with OTP succeeds", func(t *testing.T) {
		_, hrsp, err := userSvc.Register(context.Background()).Body(
			openapi.UserSvcRegisterRequest{
				AppHost: sdk.DefaultTestAppHost,
				Slug:    "test-2",
				Contact: &openapi.UserSvcContactInput{
					Id:       "test2@test.com",
					Platform: "email",
					OtpId:    &otpRsp.OtpId,
					OtpCode:  otpRsp.Code,
				},
			},
		).Execute()
		require.NoError(t, err, hrsp)
	})

	t.Run("log in with same OTP fails", func(t *testing.T) {
		_, hrsp, err := userSvc.Login(context.Background()).Body(
			openapi.UserSvcLoginRequest{
				AppHost: sdk.DefaultTestAppHost,
				Contact: &openapi.UserSvcContactInput{
					Id:       "test2@test.com",
					Platform: "email",
					OtpId:    &otpRsp.OtpId,
					OtpCode:  otpRsp.Code,
				},
			},
		).Execute()
		require.Error(t, err, hrsp)
	})

	t.Run("log in with new OTP succeeds", func(t *testing.T) {
		otpRsp2, _, err := userSvc.SendOtp(context.Background()).Body(
			openapi.UserSvcSendOtpRequest{
				AppHost:         sdk.DefaultTestAppHost,
				ContactId:       "test2@test.com",
				ContactPlatform: "email",
			},
		).Execute()
		require.NoError(t, err, otpRsp2)

		_, hrsp, err := userSvc.Login(context.Background()).Body(
			openapi.UserSvcLoginRequest{
				AppHost: sdk.DefaultTestAppHost,
				Contact: &openapi.UserSvcContactInput{
					Id:       "test2@test.com",
					Platform: "email",
					OtpId:    &otpRsp2.OtpId,
					OtpCode:  otpRsp2.Code,
				},
			},
		).Execute()
		require.NoError(t, err, hrsp)
	})
}

func TestLogin__PasswordOtpBoth(t *testing.T) {
	hs := &di.HandlerSwitcher{}
	server := httptest.NewServer(hs)
	defer server.Close()

	options := &universe.Options{
		Test:           true,
		Url:            server.URL,
		VerifyContacts: true,
	}
	universe, err := di.BigBang(options)
	require.NoError(t, err)

	hs.UpdateHandler(universe.Router)

	err = universe.StarterFunc()
	require.NoError(t, err)

	userSvc := options.ClientFactory.Client().UserSvcAPI
	_, hrsp, err := userSvc.Register(context.Background()).Body(
		openapi.UserSvcRegisterRequest{
			AppHost:  sdk.DefaultTestAppHost,
			Slug:     "test-1",
			Password: openapi.PtrString("test"),
		},
	).Execute()
	require.NoError(t, err, hrsp)

	_, hrsp, err = userSvc.Register(context.Background()).Body(
		openapi.UserSvcRegisterRequest{
			AppHost: sdk.DefaultTestAppHost,
			Slug:    "test-2",
			Contact: &openapi.UserSvcContactInput{
				Id:       "test1@test.com",
				Platform: "email",
			},
			Password: openapi.PtrString("test"),
		},
	).Execute()
	require.Error(t, err, hrsp)

	otpRsp, _, err := userSvc.SendOtp(context.Background()).Body(
		openapi.UserSvcSendOtpRequest{
			AppHost:         sdk.DefaultTestAppHost,
			ContactId:       "test1@test.com",
			ContactPlatform: "email",
		},
	).Execute()
	require.NoError(t, err, otpRsp)

	t.Run("registration with OTP succeeds", func(t *testing.T) {
		_, hrsp, err = userSvc.Register(context.Background()).Body(
			openapi.UserSvcRegisterRequest{
				AppHost: sdk.DefaultTestAppHost,
				Slug:    "test-2",
				Contact: &openapi.UserSvcContactInput{
					Id:       "test1@test.com",
					Platform: "email",
					OtpId:    &otpRsp.OtpId,
					OtpCode:  otpRsp.Code,
				},
				Password: openapi.PtrString("test"),
			},
		).Execute()
		require.NoError(t, err, hrsp)
	})

	t.Run("log in with only OTP, no password", func(t *testing.T) {
		_, hrsp, err = userSvc.Login(context.Background()).Body(
			openapi.UserSvcLoginRequest{
				AppHost: sdk.DefaultTestAppHost,
				Contact: &openapi.UserSvcContactInput{
					Id:       "test1@test.com",
					Platform: "email",
					OtpId:    &otpRsp.OtpId,
					OtpCode:  otpRsp.Code,
				},
				Password: openapi.PtrString("test"),
			},
		).Execute()
		require.NoError(t, err, hrsp)
	})

	t.Run("log in with password", func(t *testing.T) {
		_, hrsp, err = userSvc.Login(context.Background()).Body(
			openapi.UserSvcLoginRequest{
				AppHost:  sdk.DefaultTestAppHost,
				Slug:     openapi.PtrString("test-2"),
				Password: openapi.PtrString("test"),
			},
		).Execute()
		require.NoError(t, err, hrsp)
	})
}

func TestLoginAfterTokenExpiry(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		TokenExpiration: time.Second,
		Test:            true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)

	ctx := context.Background()
	_, _, err = test.MakeClients(clientFactory, sdk.DefaultTestAppHost, 1)
	require.NoError(t, err)

	time.Sleep(1 * time.Second)

	rsp, _, err := clientFactory.Client().UserSvcAPI.Login(ctx).Body(
		openapi.UserSvcLoginRequest{
			AppHost:  sdk.DefaultTestAppHost,
			Slug:     openapi.PtrString("test-user-slug-0"),
			Password: openapi.PtrString("testUserPassword0"),
		},
	).Execute()
	require.NoError(t, err)
	require.NotEmpty(t, rsp.Token.Token)
	expiresAt, err := time.Parse(time.RFC3339, rsp.Token.ExpiresAt)
	require.NoError(t, err)
	require.True(t, expiresAt.After(time.Now().Add(900*time.Millisecond)))
}

func TestOrganization(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		Test: true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)

	manyClients, _, err := test.MakeClients(clientFactory, sdk.DefaultTestAppHost, 3)
	require.NoError(t, err)

	userClient := manyClients[0]

	userToken := userClient.GetConfig().DefaultHeader["Authorization"]
	userToken = strings.Replace(userToken, "Bearer ", "", -1)
	require.Equal(t, true, len(userToken) > 0)

	otherClient := manyClients[1]
	thirdClient := manyClients[2]

	publicKeyRsp, _, err := clientFactory.Client().
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

			t.Run("token refresh still works", func(t *testing.T) {
				// Creating an org mints a new token.
				// We've had an issue in the past where the token minted during org save could not be refreshed.
				_, hrsp, err := userClient.UserSvcAPI.RefreshToken(context.Background()).
					Execute()
				require.NoError(t, err, hrsp)
			})

			claim, err := auth.AuthorizerImpl{}.ParseJWT(
				publicKeyRsp.PublicKey,
				userToken,
			)
			require.NoError(t, err)
			require.NotNil(t, claim)
			require.Equal(t, 1, len(claim.Roles), claim.Roles)
			require.NotEmpty(t, claim.ExpiresAt)
			// ~  5minutes
			require.Equal(t, claim.ExpiresAt.Time.After(time.Now().Add(298*time.Second)), true)
			require.Equal(t, claim.ExpiresAt.Time.Before(time.Now().Add(302*time.Second)), true)

			loginReq := openapi.UserSvcLoginRequest{
				AppHost:  sdk.DefaultTestAppHost,
				Slug:     openapi.PtrString("test-user-slug-0"),
				Password: openapi.PtrString("testUserPassword0"),
			}
			loginRsp, _, err := userClient.UserSvcAPI.Login(context.Background()).
				Body(loginReq).
				Execute()
			require.NoError(t, err)

			claim, err = auth.AuthorizerImpl{}.ParseJWT(
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

		_, _, err = userClient.UserSvcAPI.SaveMembership(
			context.Background(),
			orgId1,
			readSelfRsp.User.Id,
		).
			Execute()
		require.NoError(t, err)

		loginReq := openapi.UserSvcLoginRequest{
			AppHost:  sdk.DefaultTestAppHost,
			Slug:     openapi.PtrString("test-user-slug-1"),
			Password: openapi.PtrString("testUserPassword1"),
		}
		// log in again and see the claim
		loginRsp, _, err := otherClient.UserSvcAPI.Login(context.Background()).
			Body(loginReq).
			Execute()
		require.NoError(t, err)

		claim, err := auth.AuthorizerImpl{}.ParseJWT(
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

		_, _, err = thirdClient.UserSvcAPI.DeleteMembership(
			context.Background(),
			orgId1,
			readSelfRsp.User.Id,
		).
			Execute()
		// third user cannot remove the second from the org of the first
		require.Error(t, err)

		_, _, err = userClient.UserSvcAPI.DeleteMembership(
			context.Background(),
			orgId1,
			readSelfRsp.User.Id,
		).
			Execute()
		require.NoError(t, err)
	})
}
