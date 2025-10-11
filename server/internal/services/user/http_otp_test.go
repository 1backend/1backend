package userservice_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	openapi "github.com/1backend/1backend/clients/go"
	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/server/internal/di"
	usertypes "github.com/1backend/1backend/server/internal/services/user/types"
	"github.com/1backend/1backend/server/internal/universe"
)

func TestOTPLoginFlow(t *testing.T) {
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
	userSvc := options.ClientFactory.Client().UserSvcAPI

	contact := "otp-login@example.com"
	password := "supersecret"

	_, _, err = userSvc.Register(ctx).Body(openapi.UserSvcRegisterRequest{
		AppHost: sdk.DefaultTestAppHost,
		Slug:    "otp-login-user",
		Contact: &openapi.UserSvcContactInput{
			Id:       contact,
			Platform: string(usertypes.ContactPlatformEmail),
		},
		Password: openapi.PtrString(password),
	}).Execute()
	require.NoError(t, err)

	requestBody, err := json.Marshal(usertypes.RequestOTPRequest{
		ContactId: contact,
		AppHost:   sdk.DefaultTestAppHost,
	})
	require.NoError(t, err)

	resp, err := http.Post(server.URL+"/user-svc/otp/request", "application/json", bytes.NewReader(requestBody))
	require.NoError(t, err)
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode, string(respBody))

	otpResp := usertypes.RequestOTPResponse{}
	err = json.Unmarshal(respBody, &otpResp)
	require.NoError(t, err)
	require.NotEmpty(t, otpResp.OtpId)
	require.WithinDuration(t, time.Now().Add(10*time.Minute), otpResp.ExpiresAt, time.Minute)
	require.NotEmpty(t, otpResp.Code)
	t.Logf("OTP code: %s", otpResp.Code)

	verifyReq, err := json.Marshal(usertypes.VerifyOTPRequest{
		OtpId:     otpResp.OtpId,
		ContactId: contact,
		Code:      otpResp.Code,
		Device:    "cli",
	})
	require.NoError(t, err)

	verifyResp, err := http.Post(server.URL+"/user-svc/otp/verify", "application/json", bytes.NewReader(verifyReq))
	require.NoError(t, err)
	defer verifyResp.Body.Close()
	verifyBody, err := io.ReadAll(verifyResp.Body)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, verifyResp.StatusCode, string(verifyBody))

	otpVerifyResp := usertypes.VerifyOTPResponse{}
	err = json.Unmarshal(verifyBody, &otpVerifyResp)
	require.NoError(t, err)
	require.True(t, otpVerifyResp.Verified)
	require.NotNil(t, otpVerifyResp.Token)
	require.NotEmpty(t, otpVerifyResp.Token.Token)

	// OTP cannot be reused
	secondResp, err := http.Post(server.URL+"/user-svc/otp/verify", "application/json", bytes.NewReader(verifyReq))
	require.NoError(t, err)
	defer secondResp.Body.Close()
	require.Equal(t, http.StatusUnauthorized, secondResp.StatusCode)

}

func TestOTPContactVerification(t *testing.T) {
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
	userSvc := options.ClientFactory.Client().UserSvcAPI

	contact := "otp-verify@example.com"

	_, _, err = userSvc.Register(ctx).Body(openapi.UserSvcRegisterRequest{
		AppHost: sdk.DefaultTestAppHost,
		Slug:    "otp-verify-user",
		Contact: &openapi.UserSvcContactInput{
			Id:       contact,
			Platform: string(usertypes.ContactPlatformEmail),
		},
		Password: openapi.PtrString("password"),
	}).Execute()
	require.NoError(t, err)

	requestBody, err := json.Marshal(usertypes.RequestOTPRequest{
		ContactId: contact,
		Purpose:   usertypes.OTPPurposeVerifyContact,
	})
	require.NoError(t, err)

	resp, err := http.Post(server.URL+"/user-svc/otp/request", "application/json", bytes.NewReader(requestBody))
	require.NoError(t, err)
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode, string(respBody))

	otpResp := usertypes.RequestOTPResponse{}
	err = json.Unmarshal(respBody, &otpResp)
	require.NoError(t, err)
	require.NotEmpty(t, otpResp.OtpId)
	require.NotEmpty(t, otpResp.Code)

	verifyPayload, err := json.Marshal(usertypes.VerifyOTPRequest{
		OtpId:     otpResp.OtpId,
		ContactId: contact,
		Purpose:   usertypes.OTPPurposeVerifyContact,
		Code:      otpResp.Code,
	})
	require.NoError(t, err)

	verifyResp, err := http.Post(server.URL+"/user-svc/otp/verify", "application/json", bytes.NewReader(verifyPayload))
	require.NoError(t, err)
	defer verifyResp.Body.Close()
	verifyBody, err := io.ReadAll(verifyResp.Body)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, verifyResp.StatusCode, string(verifyBody))

	otpVerifyResp := usertypes.VerifyOTPResponse{}
	err = json.Unmarshal(verifyBody, &otpVerifyResp)
	require.NoError(t, err)
	require.True(t, otpVerifyResp.Verified)
	require.Nil(t, otpVerifyResp.Token)

}
