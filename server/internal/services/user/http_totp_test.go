package userservice_test

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"github.com/stretchr/testify/require"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/server/internal/di"
	user "github.com/1backend/1backend/server/internal/services/user/types"
	"github.com/1backend/1backend/server/internal/universe"
)

func TestTOTPSetupQRCodeAndLoginEnforcement(t *testing.T) {
	t.Parallel()

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
	email := "totp-login@test.com"
	var register user.RegisterResponse
	httpRsp := doJSON(
		t,
		http.MethodPost,
		server.URL+"/user-svc/register",
		"",
		user.RegisterRequest{
			AppHost:  sdk.DefaultTestAppHost,
			Slug:     "totp-login",
			Password: "testUserPassword0",
			Name:     "TOTP Login",
			Contact: user.ContactInput{
				Id:       email,
				Platform: "email",
			},
		},
		&register,
	)
	require.Equal(t, http.StatusOK, httpRsp.StatusCode)
	require.NotNil(t, register.Token)
	token := register.Token.Token

	require.False(t, readTOTPStatus(t, server.URL, email))
	require.False(t, readTOTPStatus(t, server.URL, "nobody-totp@test.com"))

	var setup user.BeginTOTPSetupResponse
	httpRsp = doJSON(
		t,
		http.MethodPost,
		server.URL+"/user-svc/totp/setup",
		token,
		user.BeginTOTPSetupRequest{},
		&setup,
	)
	require.Equal(t, http.StatusOK, httpRsp.StatusCode)
	require.NotEmpty(t, setup.TOTPId)
	require.NotEmpty(t, setup.Secret)
	require.NotEmpty(t, setup.ProvisioningURI)
	require.Equal(t, "/user-svc/totp/setup/"+setup.TOTPId+"/qr.png", setup.QRImagePath)

	key, err := otp.NewKeyFromURL(setup.ProvisioningURI)
	require.NoError(t, err)
	require.Equal(t, "1Backend", key.Issuer())
	require.Equal(t, "totp-login", key.AccountName())

	selfBefore := readSelfWithServerTypes(t, server.URL, token)
	require.False(t, selfBefore.TOTPEnabled)

	qrReq, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		server.URL+setup.QRImagePath,
		nil,
	)
	require.NoError(t, err)
	qrReq.Header.Set("Authorization", "Bearer "+token)

	qrRsp, err := http.DefaultClient.Do(qrReq)
	require.NoError(t, err)
	defer qrRsp.Body.Close()
	require.Equal(t, http.StatusOK, qrRsp.StatusCode)
	require.Equal(t, "image/png", qrRsp.Header.Get("Content-Type"))

	pngSignature := make([]byte, 8)
	_, err = io.ReadFull(qrRsp.Body, pngSignature)
	require.NoError(t, err)
	require.Equal(t, []byte{0x89, 'P', 'N', 'G', '\r', '\n', 0x1a, '\n'}, pngSignature)

	code := currentTOTPCode(t, setup.Secret)
	var enable user.EnableTOTPResponse
	httpRsp = doJSON(
		t,
		http.MethodPost,
		server.URL+"/user-svc/totp/enable",
		token,
		user.EnableTOTPRequest{
			TOTPId: setup.TOTPId,
			Code:   code,
		},
		&enable,
	)
	require.Equal(t, http.StatusOK, httpRsp.StatusCode)
	require.True(t, enable.Enabled)

	selfAfterEnable := readSelfWithServerTypes(t, server.URL, token)
	require.True(t, selfAfterEnable.TOTPEnabled)
	require.True(t, readTOTPStatus(t, server.URL, email))

	httpRsp = doJSON(
		t,
		http.MethodPost,
		server.URL+"/user-svc/login",
		"",
		user.LoginRequest{
			AppHost:  sdk.DefaultTestAppHost,
			Slug:     "totp-login",
			Password: "testUserPassword0",
		},
		nil,
	)
	require.Equal(t, http.StatusUnauthorized, httpRsp.StatusCode)

	httpRsp = doJSON(
		t,
		http.MethodPost,
		server.URL+"/user-svc/login",
		"",
		user.LoginRequest{
			AppHost:  sdk.DefaultTestAppHost,
			Slug:     "totp-login",
			Password: "testUserPassword0",
			TOTPCode: "000000",
		},
		nil,
	)
	require.Equal(t, http.StatusUnauthorized, httpRsp.StatusCode)

	var login user.LoginResponse
	httpRsp = doJSON(
		t,
		http.MethodPost,
		server.URL+"/user-svc/login",
		"",
		user.LoginRequest{
			AppHost:  sdk.DefaultTestAppHost,
			Slug:     "totp-login",
			Password: "testUserPassword0",
			TOTPCode: currentTOTPCode(t, setup.Secret),
		},
		&login,
	)
	require.Equal(t, http.StatusOK, httpRsp.StatusCode)
	require.NotNil(t, login.Token)
	require.NotEmpty(t, login.Token.Token)

	httpRsp = doJSON(
		t,
		http.MethodDelete,
		server.URL+"/user-svc/totp",
		token,
		user.DisableTOTPRequest{
			Code: currentTOTPCode(t, setup.Secret),
		},
		&user.DisableTOTPResponse{},
	)
	require.Equal(t, http.StatusOK, httpRsp.StatusCode)

	selfAfterDisable := readSelfWithServerTypes(t, server.URL, token)
	require.False(t, selfAfterDisable.TOTPEnabled)
	require.False(t, readTOTPStatus(t, server.URL, email))

	httpRsp = doJSON(
		t,
		http.MethodPost,
		server.URL+"/user-svc/login",
		"",
		user.LoginRequest{
			AppHost:  sdk.DefaultTestAppHost,
			Slug:     "totp-login",
			Password: "testUserPassword0",
		},
		&login,
	)
	require.Equal(t, http.StatusOK, httpRsp.StatusCode)
}

func TestTOTPSetupCustomIssuer(t *testing.T) {
	t.Parallel()

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

	var register user.RegisterResponse
	httpRsp := doJSON(
		t,
		http.MethodPost,
		server.URL+"/user-svc/register",
		"",
		user.RegisterRequest{
			AppHost:  sdk.DefaultTestAppHost,
			Slug:     "totp-custom-issuer",
			Password: "testUserPassword0",
			Name:     "TOTP Custom Issuer",
			Contact: user.ContactInput{
				Id:       "totp-custom-issuer@test.com",
				Platform: "email",
			},
		},
		&register,
	)
	require.Equal(t, http.StatusOK, httpRsp.StatusCode)
	require.NotNil(t, register.Token)
	token := register.Token.Token

	httpRsp = doJSON(
		t,
		http.MethodPost,
		server.URL+"/user-svc/totp/setup",
		token,
		user.BeginTOTPSetupRequest{Issuer: "bad:issuer"},
		nil,
	)
	require.Equal(t, http.StatusBadRequest, httpRsp.StatusCode)

	var setup user.BeginTOTPSetupResponse
	httpRsp = doJSON(
		t,
		http.MethodPost,
		server.URL+"/user-svc/totp/setup",
		token,
		user.BeginTOTPSetupRequest{Issuer: " auth.example.com "},
		&setup,
	)
	require.Equal(t, http.StatusOK, httpRsp.StatusCode)
	require.NotEmpty(t, setup.Secret)

	key, err := otp.NewKeyFromURL(setup.ProvisioningURI)
	require.NoError(t, err)
	require.Equal(t, "auth.example.com", key.Issuer())
	require.Equal(t, "totp-custom-issuer", key.AccountName())

	var enable user.EnableTOTPResponse
	httpRsp = doJSON(
		t,
		http.MethodPost,
		server.URL+"/user-svc/totp/enable",
		token,
		user.EnableTOTPRequest{
			TOTPId: setup.TOTPId,
			Code:   currentTOTPCode(t, setup.Secret),
		},
		&enable,
	)
	require.Equal(t, http.StatusOK, httpRsp.StatusCode)
	require.True(t, enable.Enabled)
}

func readTOTPStatus(
	t *testing.T,
	serverURL string,
	email string,
) bool {
	t.Helper()

	var rsp user.ReadTOTPStatusResponse
	httpRsp := doJSON(
		t,
		http.MethodPost,
		serverURL+"/user-svc/totp/status",
		"",
		user.ReadTOTPStatusRequest{Email: email},
		&rsp,
	)
	require.Equal(t, http.StatusOK, httpRsp.StatusCode)

	return rsp.TOTPEnabled
}

func readSelfWithServerTypes(
	t *testing.T,
	serverURL string,
	token string,
) user.ReadSelfResponse {
	t.Helper()

	var rsp user.ReadSelfResponse
	httpRsp := doJSON(
		t,
		http.MethodPost,
		serverURL+"/user-svc/self",
		token,
		user.ReadSelfRequest{},
		&rsp,
	)
	require.Equal(t, http.StatusOK, httpRsp.StatusCode)

	return rsp
}

func currentTOTPCode(t *testing.T, secret string) string {
	t.Helper()

	code, err := totp.GenerateCodeCustom(secret, time.Now(), totp.ValidateOpts{
		Period:    30,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA1,
	})
	require.NoError(t, err)

	return code
}
