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

func TestRegistration__PasswordOtpBoth(t *testing.T) {
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

func TestRegistration__OnlyOtp(t *testing.T) {
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
		require.NoError(t, err, hrsp)
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
