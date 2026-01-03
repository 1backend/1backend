package userservice_test

import (
	"context"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	openapi "github.com/1backend/1backend/clients/go"
	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/server/internal/di"
	userservice "github.com/1backend/1backend/server/internal/services/user"
	"github.com/1backend/1backend/server/internal/universe"
)

func TestRegister(t *testing.T) {
	hs := &di.HandlerSwitcher{}
	server := httptest.NewServer(hs)
	defer server.Close()

	options := &universe.Options{
		Test: true,
		Url:  server.URL,
	}
	universe, err := di.BigBang(options)
	require.NoError(t, err)

	hs.UpdateHandler(universe.Router)

	err = universe.StarterFunc()
	require.NoError(t, err)

	ctx := context.Background()

	tokenId := ""
	publicKeyRsp, _, err := options.ClientFactory.Client().
		UserSvcAPI.GetPublicKey(context.Background()).
		Execute()
	require.NoError(t, err)

	t.Run("anyone can register", func(t *testing.T) {
		rsp, _, err := options.ClientFactory.Client().UserSvcAPI.Register(ctx).Body(
			openapi.UserSvcRegisterRequest{
				AppHost:  sdk.DefaultTestAppHost,
				Slug:     "test-slug-1",
				Name:     openapi.PtrString("Test Name"),
				Password: openapi.PtrString("testPass123"),
			},
		).Execute()
		require.NoError(t, err)

		tokenId = rsp.Token.Id

		claim, err := options.Authorizer.ParseJWT(
			publicKeyRsp.PublicKey,
			rsp.Token.Token,
		)
		require.NoError(t, err)
		require.Equal(t, 1, len(claim.Roles), claim.Roles)
		require.NotEmpty(t, claim.ExpiresAt)
		// ~  5minutes
		require.Equal(t, claim.ExpiresAt.Time.After(time.Now().Add(298*time.Second)), true)
		require.Equal(t, claim.ExpiresAt.Time.Before(time.Now().Add(302*time.Second)), true)
	})

	t.Run("uppercase fails so ui must be forced to deal with it", func(t *testing.T) {
		_, _, err := options.ClientFactory.Client().UserSvcAPI.Register(ctx).Body(
			openapi.UserSvcRegisterRequest{
				Slug:     "Test-slug-1",
				Name:     openapi.PtrString("Test Name"),
				Password: openapi.PtrString("testPass123"),
			},
		).Execute()
		require.Error(t, err)
	})

	t.Run("cannot re-register", func(t *testing.T) {
		_, _, err := options.ClientFactory.Client().UserSvcAPI.Register(ctx).Body(
			openapi.UserSvcRegisterRequest{
				Slug:     "test-slug-1",
				Name:     openapi.PtrString("Test Name"),
				Password: openapi.PtrString("testPass123"),
			},
		).Execute()
		require.Error(t, err)
	})

	t.Run("login doesn't produce a new token", func(t *testing.T) {
		rsp, _, err := options.ClientFactory.Client().UserSvcAPI.Login(ctx).Body(
			openapi.UserSvcLoginRequest{
				AppHost:  sdk.DefaultTestAppHost,
				Slug:     openapi.PtrString("test-slug-1"),
				Password: openapi.PtrString("testPass123"),
			},
		).Execute()
		require.NoError(t, err)
		require.Equal(t, true, rsp.Token.Id == tokenId)
	})
}

func TestRegistration(t *testing.T) {
	hs := &di.HandlerSwitcher{}
	server := httptest.NewServer(hs)
	defer server.Close()

	options := &universe.Options{
		Test: true,
		Url:  server.URL,
	}
	universe, err := di.BigBang(options)
	require.NoError(t, err)

	hs.UpdateHandler(universe.Router)

	err = universe.StarterFunc()
	require.NoError(t, err)

	userSvc := options.ClientFactory.Client().UserSvcAPI
	_, hrsp, err := userSvc.Register(context.Background()).Body(
		openapi.UserSvcRegisterRequest{
			AppHost: sdk.DefaultTestAppHost,
			Slug:    "test-1",
			Contact: &openapi.UserSvcContactInput{
				Id:       "test1@test.com",
				Platform: "email",
			},
			Password: openapi.PtrString("test"),
		},
	).Execute()
	require.NoError(t, err, hrsp)

	t.Run("slug login works", func(t *testing.T) {
		loginReq := openapi.UserSvcLoginRequest{
			AppHost:  sdk.DefaultTestAppHost,
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
			AppHost: sdk.DefaultTestAppHost,
			Contact: &openapi.UserSvcContactInput{
				Id:       "test1@test.com",
				Platform: "email",
			},
			Password: openapi.PtrString("test"),
		}
		_, _, err := options.ClientFactory.Client().UserSvcAPI.Login(context.Background()).
			Body(loginReq).
			Execute()
		require.NoError(t, err)
	})
}

func TestRegistration__VerifyContacts(t *testing.T) {
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
}

func TestRegister__EscalationAttack(t *testing.T) {
	hs := &di.HandlerSwitcher{}
	server := httptest.NewServer(hs)
	defer server.Close()

	options := &universe.Options{
		Test:           true,
		Url:            server.URL,
		VerifyContacts: true,
	}
	u, _ := di.BigBang(options)
	hs.UpdateHandler(u.Router)
	u.StarterFunc()

	userSvc := options.ClientFactory.Client().UserSvcAPI
	ctx := context.Background()

	// 1. SETUP: Create a legitimate victim account
	victimEmail := "victim@target.com"
	victimSlug := "victim-slug"

	// (Assume victim is already registered via some flow)
	// For this test, we'll quickly register them properly first.
	otpV, _, _ := userSvc.SendOtp(ctx).Body(openapi.UserSvcSendOtpRequest{
		AppHost: sdk.DefaultTestAppHost, ContactId: victimEmail, ContactPlatform: "email",
	}).Execute()
	userSvc.Register(ctx).Body(openapi.UserSvcRegisterRequest{
		AppHost: sdk.DefaultTestAppHost,
		Slug:    victimSlug,
		Contact: &openapi.UserSvcContactInput{
			Id: victimEmail, Platform: "email", OtpId: &otpV.OtpId, OtpCode: otpV.Code,
		},
	}).Execute()

	// 2. ATTACK: Attacker tries to "Register" using Victim's Slug but Attacker's Email
	attackerEmail := "attacker@evil.com"

	// Attacker gets a VALID OTP for THEIR OWN email
	otpA, _, _ := userSvc.SendOtp(ctx).Body(openapi.UserSvcSendOtpRequest{
		AppHost: sdk.DefaultTestAppHost, ContactId: attackerEmail, ContactPlatform: "email",
	}).Execute()

	t.Run("Attacker tries to register with Victim Slug and Attacker OTP", func(t *testing.T) {
		// This request provides:
		// - The Victim's Slug (to target their account)
		// - The Attacker's Email + Valid OTP (to pass the 'verified' check)
		_, hrsp, err := userSvc.Register(ctx).Body(openapi.UserSvcRegisterRequest{
			AppHost: sdk.DefaultTestAppHost,
			Slug:    victimSlug, // Targeting the victim
			Contact: &openapi.UserSvcContactInput{
				Id:       attackerEmail,
				Platform: "email",
				OtpId:    &otpA.OtpId,
				OtpCode:  otpA.Code,
			},
		}).Execute()

		// IF THE CODE IS VULNERABLE:
		// 1. verifyContactOTP(attackerEmail) passes.
		// 2. createUserWithoutVerification fails because "victim-slug" exists.
		// 3. The code calls s.login(victimSlug, otpAlreadyVerified=true).
		// 4. The attacker gets a token for the Victim.

		if err == nil {
			// Check if the token issued actually belongs to the attacker or victim
			// By parsing the token or calling ReadSelf
			t.Errorf("SECURITY FLAW: Registration allowed login escalation. Status: %d", hrsp.StatusCode)

			// Further verification: Check who the token belongs to
			// (Assuming you add the token to the client config here)
		}
	})
}

func TestSlugRegexp(t *testing.T) {
	shouldWork := []string{
		"1backend",
		"test-slug",
		"backend1",
		"test-1-slug",
	}

	shouldNotWork := []string{
		"Test-Slug",
		"-test-slug",
		"test-slug-",
		"test slug",
		"test_slug",
		"test@slug",
		"test#slug",
		"test$slug",
		"test%slug",
		"test^slug",
		"test&slug",
		"test*slug",
		"test(slug)",
	}

	for _, slug := range shouldWork {
		if !userservice.SlugRegexp.MatchString(slug) {
			t.Errorf("Expected %s to be valid", slug)
		}
	}

	for _, slug := range shouldNotWork {
		if userservice.SlugRegexp.MatchString(slug) {
			t.Errorf("Expected %s to be invalid", slug)
		}
	}
}
