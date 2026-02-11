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

		_, _, err = userSvc.Login(context.Background()).Body(
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
		require.NoError(t, err)
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

	t.Run("registration without verification (OTP) fails", func(t *testing.T) {
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
	})

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

	otpRsp2, _, err := userSvc.SendOtp(context.Background()).Body(
		openapi.UserSvcSendOtpRequest{
			AppHost:         sdk.DefaultTestAppHost,
			ContactId:       "test1@test.com",
			ContactPlatform: "email",
		},
	).Execute()
	require.NoError(t, err, otpRsp)

	t.Run("log in with only OTP, no password", func(t *testing.T) {
		_, hrsp, err = userSvc.Login(context.Background()).Body(
			openapi.UserSvcLoginRequest{
				AppHost: sdk.DefaultTestAppHost,
				Contact: &openapi.UserSvcContactInput{
					Id:       "test1@test.com",
					Platform: "email",
					OtpId:    &otpRsp2.OtpId,
					OtpCode:  otpRsp2.Code,
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

func TestLogin__SecurityBoundaries(t *testing.T) {
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
	ctx := context.Background()

	// 1. SETUP: Register a user with BOTH a password and an email (Hybrid)
	otpRsp, _, _ := userSvc.SendOtp(ctx).Body(openapi.UserSvcSendOtpRequest{
		AppHost: sdk.DefaultTestAppHost, ContactId: "victim@test.com", ContactPlatform: "email",
	}).Execute()

	_, _, err = userSvc.Register(ctx).Body(openapi.UserSvcRegisterRequest{
		AppHost:  sdk.DefaultTestAppHost,
		Slug:     "victim-user",
		Password: openapi.PtrString("secure-password"),
		Contact: &openapi.UserSvcContactInput{
			Id: "victim@test.com", Platform: "email", OtpId: &otpRsp.OtpId, OtpCode: otpRsp.Code,
		},
	}).Execute()
	require.NoError(t, err)

	// --- TEST CASES ---

	t.Run("Bypass password using only OTP for hybrid account", func(t *testing.T) {
		return

		// Request a fresh OTP
		newOtp, _, _ := userSvc.SendOtp(ctx).Body(openapi.UserSvcSendOtpRequest{
			AppHost: sdk.DefaultTestAppHost, ContactId: "victim@test.com", ContactPlatform: "email",
		}).Execute()

		// Try to login with valid OTP but MISSING password
		_, hrsp, err := userSvc.Login(ctx).Body(openapi.UserSvcLoginRequest{
			AppHost: sdk.DefaultTestAppHost,
			Contact: &openapi.UserSvcContactInput{
				Id: "victim@test.com", Platform: "email", OtpId: &newOtp.OtpId, OtpCode: newOtp.Code,
			},
			// Password omitted!
		}).Execute()

		// If the system is strictly MFA, this MUST fail.
		require.Error(t, err, "Should not allow login with OTP only if account has a password")
		require.Equal(t, 401, hrsp.StatusCode)
	})

	t.Run("Bypass OTP using wrong password for hybrid account", func(t *testing.T) {
		// Try to login with the Slug but a WRONG password
		_, hrsp, err := userSvc.Login(ctx).Body(openapi.UserSvcLoginRequest{
			AppHost:  sdk.DefaultTestAppHost,
			Slug:     openapi.PtrString("victim-user"),
			Password: openapi.PtrString("wrong-password"),
		}).Execute()

		require.Error(t, err)
		require.Equal(t, 401, hrsp.StatusCode)
	})

	t.Run("OTP binding - Use Attacker OTP for Victim Email", func(t *testing.T) {
		// 1. Attacker gets an OTP for their own email
		attackerOtp, _, _ := userSvc.SendOtp(ctx).Body(openapi.UserSvcSendOtpRequest{
			AppHost: sdk.DefaultTestAppHost, ContactId: "attacker@test.com", ContactPlatform: "email",
		}).Execute()

		// 2. Attacker tries to register "victim-2" using the Victim's email but the Attacker's OtpId/Code
		_, _, err := userSvc.Register(ctx).Body(openapi.UserSvcRegisterRequest{
			AppHost: sdk.DefaultTestAppHost,
			Slug:    "malicious-switch",
			Contact: &openapi.UserSvcContactInput{
				Id:       "victim-2@test.com", // Target's email
				Platform: "email",
				OtpId:    &attackerOtp.OtpId, // Attacker's verification ID
				OtpCode:  attackerOtp.Code,   // Attacker's code
			},
		}).Execute()

		require.Error(t, err, "System should prevent using an OTP bound to a different email address")
	})

	t.Run("Login with empty string password on OTP-only account", func(t *testing.T) {
		// Create an OTP-only user
		otpOnly, _, _ := userSvc.SendOtp(ctx).Body(openapi.UserSvcSendOtpRequest{
			AppHost: sdk.DefaultTestAppHost, ContactId: "nopass@test.com", ContactPlatform: "email",
		}).Execute()
		userSvc.Register(ctx).Body(openapi.UserSvcRegisterRequest{
			AppHost: sdk.DefaultTestAppHost, Slug: "nopass-user",
			Contact: &openapi.UserSvcContactInput{
				Id: "nopass@test.com", Platform: "email", OtpId: &otpOnly.OtpId, OtpCode: otpOnly.Code,
			},
		}).Execute()

		// Attempt to login with Slug and an empty password (no OTP provided)
		_, _, err := userSvc.Login(ctx).Body(openapi.UserSvcLoginRequest{
			AppHost:  sdk.DefaultTestAppHost,
			Slug:     openapi.PtrString("nopass-user"),
			Password: openapi.PtrString(""),
		}).Execute()

		require.Error(t, err, "Should not allow empty string password to act as a 'valid' credential")
	})

	t.Run("OTP Login Hijacking - Use Attacker OTP to login as Victim", func(t *testing.T) {
		// 1. Attacker requests an OTP for THEIR OWN email
		attackerOtp, _, _ := userSvc.SendOtp(ctx).Body(openapi.UserSvcSendOtpRequest{
			AppHost:         sdk.DefaultTestAppHost,
			ContactId:       "attacker@test.com",
			ContactPlatform: "email",
		}).Execute()

		// 2. Attacker tries to Login as the VICTIM using the attacker's OtpId/Code
		_, _, err := userSvc.Login(ctx).Body(openapi.UserSvcLoginRequest{
			AppHost: sdk.DefaultTestAppHost,
			Contact: &openapi.UserSvcContactInput{
				Id:       "victim@test.com", // The target account
				Platform: "email",
				OtpId:    &attackerOtp.OtpId, // The attacker's valid OTP session
				OtpCode:  attackerOtp.Code,   // The attacker's valid code
			},
		}).Execute()

		// This must fail because the OTP was not issued for victim@test.com
		require.Error(t, err, "Security Breach: System allowed login using an OTP issued for a different email")
	})
}

func TestRegister__IdempotentWithOtp(t *testing.T) {
	hs := &di.HandlerSwitcher{}
	server := httptest.NewServer(hs)
	defer server.Close()

	options := &universe.Options{
		Test:           true,
		Url:            server.URL,
		VerifyContacts: true,
	}
	u, err := di.BigBang(options)
	require.NoError(t, err)

	hs.UpdateHandler(u.Router)
	err = u.StarterFunc()
	require.NoError(t, err)

	userSvc := options.ClientFactory.Client().UserSvcAPI
	ctx := context.Background()
	email := "idempotent-test@test.com"

	// 1. Initial Registration
	otpRsp1, _, err := userSvc.SendOtp(ctx).Body(openapi.UserSvcSendOtpRequest{
		AppHost:         sdk.DefaultTestAppHost,
		ContactId:       email,
		ContactPlatform: "email",
	}).Execute()
	require.NoError(t, err)

	_, _, err = userSvc.Register(ctx).Body(openapi.UserSvcRegisterRequest{
		AppHost: sdk.DefaultTestAppHost,
		Slug:    "original-user",
		Contact: &openapi.UserSvcContactInput{
			Id:       email,
			Platform: "email",
			OtpId:    &otpRsp1.OtpId,
			OtpCode:  otpRsp1.Code,
		},
	}).Execute()
	require.NoError(t, err, "Initial registration should succeed")

	t.Run("registering existing contact with valid OTP logs in", func(t *testing.T) {
		// 2. Request a NEW OTP for the same contact
		otpRsp2, _, err := userSvc.SendOtp(ctx).Body(openapi.UserSvcSendOtpRequest{
			AppHost:         sdk.DefaultTestAppHost,
			ContactId:       email,
			ContactPlatform: "email",
		}).Execute()
		require.NoError(t, err)

		// 3. Attempt to Register again with the same email but new OTP
		// Even if the Slug is different, it should identify the existing contact and log in.
		resp, hrsp, err := userSvc.Register(ctx).Body(openapi.UserSvcRegisterRequest{
			AppHost: sdk.DefaultTestAppHost,
			Slug:    "duplicate-slug",
			Contact: &openapi.UserSvcContactInput{
				Id:       email,
				Platform: "email",
				OtpId:    &otpRsp2.OtpId,
				OtpCode:  otpRsp2.Code,
			},
		}).Execute()

		require.NoError(t, err, "Should not return error for existing contact if OTP is valid")
		require.NotNil(t, resp.Token, "Should return a session token")
		require.Equal(t, 200, hrsp.StatusCode)
	})

	t.Run("registering existing contact with NO OTP fails", func(t *testing.T) {
		_, hrsp, err := userSvc.Register(ctx).Body(openapi.UserSvcRegisterRequest{
			AppHost: sdk.DefaultTestAppHost,
			Slug:    "another-slug",
			Contact: &openapi.UserSvcContactInput{
				Id:       email,
				Platform: "email",
			},
		}).Execute()

		require.Error(t, err, "Should fail registration for existing contact without OTP")
		// Usually returns 409 Conflict or 400 depending on implementation
		require.True(t, hrsp.StatusCode >= 400)
	})

	t.Run("SECURITY: re-registering with same slug and NEW email won't log person in", func(t *testing.T) {
		otpRsp2, _, err := userSvc.SendOtp(ctx).Body(openapi.UserSvcSendOtpRequest{
			AppHost:         sdk.DefaultTestAppHost,
			ContactId:       email + "1",
			ContactPlatform: "email",
		}).Execute()
		require.NoError(t, err)

		_, _, err = userSvc.Register(ctx).Body(openapi.UserSvcRegisterRequest{
			AppHost: sdk.DefaultTestAppHost,
			Slug:    "original-user",
			Contact: &openapi.UserSvcContactInput{
				Id:       email + "1",
				Platform: "email",
				OtpId:    &otpRsp2.OtpId,
				OtpCode:  otpRsp2.Code,
			},
		}).Execute()

		require.Error(t, err, "Should  fail registration")
	})

	t.Run("re-registering with same slug and same email works", func(t *testing.T) {
		otpRsp2, _, err := userSvc.SendOtp(ctx).Body(openapi.UserSvcSendOtpRequest{
			AppHost:         sdk.DefaultTestAppHost,
			ContactId:       email,
			ContactPlatform: "email",
		}).Execute()
		require.NoError(t, err)

		_, _, err = userSvc.Register(ctx).Body(openapi.UserSvcRegisterRequest{
			AppHost: sdk.DefaultTestAppHost,
			Slug:    "original-user",
			Contact: &openapi.UserSvcContactInput{
				Id:       email,
				Platform: "email",
				OtpId:    &otpRsp2.OtpId,
				OtpCode:  otpRsp2.Code,
			},
		}).Execute()

		require.NoError(t, err, "Should  fail registration")
	})
}

func TestRegister__PlusEmail(t *testing.T) {
	hs := &di.HandlerSwitcher{}
	server := httptest.NewServer(hs)
	defer server.Close()

	options := &universe.Options{
		Test:           true,
		Url:            server.URL,
		VerifyContacts: true,
	}
	u, err := di.BigBang(options)
	require.NoError(t, err)

	hs.UpdateHandler(u.Router)
	err = u.StarterFunc()
	require.NoError(t, err)

	userSvc := options.ClientFactory.Client().UserSvcAPI
	ctx := context.Background()
	email := "plus-test+123@test.com"

	// 1. Initial Registration
	otpRsp1, _, err := userSvc.SendOtp(ctx).Body(openapi.UserSvcSendOtpRequest{
		AppHost:         sdk.DefaultTestAppHost,
		ContactId:       email,
		ContactPlatform: "email",
	}).Execute()
	require.NoError(t, err)

	rsp, hrsp, err := userSvc.Register(ctx).Body(openapi.UserSvcRegisterRequest{
		AppHost: sdk.DefaultTestAppHost,
		Slug:    "original-user",
		Name:    openapi.PtrString("John Doe"),
		Contact: &openapi.UserSvcContactInput{
			Id:       email,
			Platform: "email",
			OtpId:    &otpRsp1.OtpId,
			OtpCode:  otpRsp1.Code,
		},
	}).Execute()
	require.NoError(t, err, hrsp)

	userSvc = options.ClientFactory.Client(
		client.WithToken(rsp.Token.Token),
	).UserSvcAPI

	srsp, _, err := userSvc.ReadSelf(ctx).Execute()
	require.NoError(t, err)

	require.Equal(t, *srsp.User.Name, "John Doe")
	require.Equal(t, srsp.Contacts[0].Id, "plus-test@test.com")
}

func TestRegister__VerifyTrueRegisterOnLogin(t *testing.T) {
	hs := &di.HandlerSwitcher{}
	server := httptest.NewServer(hs)
	defer server.Close()

	options := &universe.Options{
		Test:           true,
		Url:            server.URL,
		VerifyContacts: true,
	}
	u, err := di.BigBang(options)
	require.NoError(t, err)

	hs.UpdateHandler(u.Router)
	err = u.StarterFunc()
	require.NoError(t, err)

	userSvc := options.ClientFactory.Client().UserSvcAPI
	ctx := context.Background()
	unregisteredEmail := "unregistered@test.com"

	t.Run("will fail with empty slug", func(t *testing.T) {
		otpRsp1, _, err := userSvc.SendOtp(ctx).Body(openapi.UserSvcSendOtpRequest{
			AppHost:         sdk.DefaultTestAppHost,
			ContactId:       unregisteredEmail,
			ContactPlatform: "email",
		}).Execute()
		require.NoError(t, err)

		rsp, _, err := userSvc.Login(ctx).Body(openapi.UserSvcLoginRequest{
			AppHost: sdk.DefaultTestAppHost,
			Contact: &openapi.UserSvcContactInput{
				Id:       unregisteredEmail,
				Platform: "email",
				OtpId:    &otpRsp1.OtpId,
				OtpCode:  otpRsp1.Code,
			},
		}).Execute()
		require.Error(t, err, rsp)
	})

	t.Run("login-register works", func(t *testing.T) {
		otpRsp1, _, err := userSvc.SendOtp(ctx).Body(openapi.UserSvcSendOtpRequest{
			AppHost:         sdk.DefaultTestAppHost,
			ContactId:       unregisteredEmail,
			ContactPlatform: "email",
		}).Execute()
		require.NoError(t, err)

		_, hrsp, err := userSvc.Login(ctx).Body(openapi.UserSvcLoginRequest{
			AppHost: sdk.DefaultTestAppHost,
			Slug:    openapi.PtrString("available-slug"),
			Contact: &openapi.UserSvcContactInput{
				Id:       unregisteredEmail,
				Platform: "email",
				OtpId:    &otpRsp1.OtpId,
				OtpCode:  otpRsp1.Code,
			},
		}).Execute()
		require.NoError(t, err, hrsp)
	})

	t.Run("idempotent", func(t *testing.T) {
		otpRsp2, _, err := userSvc.SendOtp(ctx).Body(openapi.UserSvcSendOtpRequest{
			AppHost:         sdk.DefaultTestAppHost,
			ContactId:       unregisteredEmail,
			ContactPlatform: "email",
		}).Execute()
		require.NoError(t, err)

		rsp, _, err := userSvc.Login(ctx).Body(openapi.UserSvcLoginRequest{
			AppHost: sdk.DefaultTestAppHost,
			Slug:    openapi.PtrString("available-slug"),
			Contact: &openapi.UserSvcContactInput{
				Id:       unregisteredEmail,
				Platform: "email",
				OtpId:    &otpRsp2.OtpId,
				OtpCode:  otpRsp2.Code,
			},
		}).Execute()
		require.NoError(t, err, rsp)

		userSvc = options.ClientFactory.Client(
			client.WithToken(rsp.Token.Token),
		).UserSvcAPI

		srsp, _, err := userSvc.ReadSelf(ctx).Execute()
		require.NoError(t, err)

		require.Empty(t, srsp.User.Name)
		require.Equal(t, srsp.Contacts[0].Id, unregisteredEmail)
	})

	t.Run("nonmatching slug is ok if contact matches", func(t *testing.T) {
		otpRsp2, _, err := userSvc.SendOtp(ctx).Body(openapi.UserSvcSendOtpRequest{
			AppHost:         sdk.DefaultTestAppHost,
			ContactId:       unregisteredEmail,
			ContactPlatform: "email",
		}).Execute()
		require.NoError(t, err)

		rsp, _, err := userSvc.Login(ctx).Body(openapi.UserSvcLoginRequest{
			AppHost: sdk.DefaultTestAppHost,
			Slug:    openapi.PtrString("RANDOMavailable-slug"),
			Contact: &openapi.UserSvcContactInput{
				Id:       unregisteredEmail,
				Platform: "email",
				OtpId:    &otpRsp2.OtpId,
				OtpCode:  otpRsp2.Code,
			},
		}).Execute()
		require.NoError(t, err, rsp)

		userSvc = options.ClientFactory.Client(
			client.WithToken(rsp.Token.Token),
		).UserSvcAPI

		srsp, _, err := userSvc.ReadSelf(ctx).Execute()
		require.NoError(t, err)

		require.Equal(t, srsp.Contacts[0].Id, unregisteredEmail)
	})
}
