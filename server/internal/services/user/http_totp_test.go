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

	var staleSetup user.BeginTOTPSetupResponse
	httpRsp = doJSON(
		t,
		http.MethodPost,
		server.URL+"/user-svc/totp/setup",
		token,
		user.BeginTOTPSetupRequest{},
		&staleSetup,
	)
	require.Equal(t, http.StatusOK, httpRsp.StatusCode)
	require.NotEmpty(t, staleSetup.TOTPId)
	require.NotEqual(t, setup.TOTPId, staleSetup.TOTPId)

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
		server.URL+"/user-svc/totp/enable",
		token,
		user.EnableTOTPRequest{
			TOTPId: staleSetup.TOTPId,
			Code:   currentTOTPCode(t, staleSetup.Secret),
		},
		nil,
	)
	require.Equal(t, http.StatusConflict, httpRsp.StatusCode)

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
		http.MethodPost,
		server.URL+"/user-svc/totp/setup",
		token,
		user.BeginTOTPSetupRequest{},
		nil,
	)
	require.Equal(t, http.StatusConflict, httpRsp.StatusCode)

	httpRsp = doJSON(
		t,
		http.MethodPost,
		server.URL+"/user-svc/totp/setup",
		token,
		user.BeginTOTPSetupRequest{CurrentCode: "000000"},
		nil,
	)
	require.Equal(t, http.StatusUnauthorized, httpRsp.StatusCode)

	var replacementSetup user.BeginTOTPSetupResponse
	httpRsp = doJSON(
		t,
		http.MethodPost,
		server.URL+"/user-svc/totp/setup",
		token,
		user.BeginTOTPSetupRequest{
			Issuer:      "updated.example.com",
			AccountName: "$email",
			CurrentCode: currentTOTPCode(t, setup.Secret),
		},
		&replacementSetup,
	)
	require.Equal(t, http.StatusOK, httpRsp.StatusCode)
	require.NotEmpty(t, replacementSetup.TOTPId)
	require.NotEmpty(t, replacementSetup.Secret)
	require.NotEqual(t, setup.TOTPId, replacementSetup.TOTPId)
	require.Equal(t, setup.Secret, replacementSetup.Secret)

	replacementKey, err := otp.NewKeyFromURL(replacementSetup.ProvisioningURI)
	require.NoError(t, err)
	require.Equal(t, "updated.example.com", replacementKey.Issuer())
	require.Equal(t, email, replacementKey.AccountName())

	var replacementEnable user.EnableTOTPResponse
	httpRsp = doJSON(
		t,
		http.MethodPost,
		server.URL+"/user-svc/totp/enable",
		token,
		user.EnableTOTPRequest{
			TOTPId: replacementSetup.TOTPId,
			Code:   currentTOTPCode(t, replacementSetup.Secret),
		},
		&replacementEnable,
	)
	require.Equal(t, http.StatusOK, httpRsp.StatusCode)
	require.True(t, replacementEnable.Enabled)
	require.True(t, readTOTPStatus(t, server.URL, email))

	var replacementLogin user.LoginResponse
	httpRsp = doJSON(
		t,
		http.MethodPost,
		server.URL+"/user-svc/login",
		"",
		user.LoginRequest{
			AppHost:  sdk.DefaultTestAppHost,
			Slug:     "totp-login",
			Password: "testUserPassword0",
			TOTPCode: currentTOTPCode(t, replacementSetup.Secret),
		},
		&replacementLogin,
	)
	require.Equal(t, http.StatusOK, httpRsp.StatusCode)
	require.NotNil(t, replacementLogin.Token)
	require.NotEmpty(t, replacementLogin.Token.Token)

	var rotatedSetup user.BeginTOTPSetupResponse
	httpRsp = doJSON(
		t,
		http.MethodPost,
		server.URL+"/user-svc/totp/setup",
		token,
		user.BeginTOTPSetupRequest{
			Issuer:       "rotated.example.com",
			AccountName:  "$slug",
			CurrentCode:  currentTOTPCode(t, replacementSetup.Secret),
			RotateSecret: true,
		},
		&rotatedSetup,
	)
	require.Equal(t, http.StatusOK, httpRsp.StatusCode)
	require.NotEmpty(t, rotatedSetup.TOTPId)
	require.NotEmpty(t, rotatedSetup.Secret)
	require.NotEqual(t, replacementSetup.TOTPId, rotatedSetup.TOTPId)
	require.NotEqual(t, replacementSetup.Secret, rotatedSetup.Secret)

	rotatedKey, err := otp.NewKeyFromURL(rotatedSetup.ProvisioningURI)
	require.NoError(t, err)
	require.Equal(t, "rotated.example.com", rotatedKey.Issuer())
	require.Equal(t, "totp-login", rotatedKey.AccountName())

	var rotatedEnable user.EnableTOTPResponse
	httpRsp = doJSON(
		t,
		http.MethodPost,
		server.URL+"/user-svc/totp/enable",
		token,
		user.EnableTOTPRequest{
			TOTPId: rotatedSetup.TOTPId,
			Code:   currentTOTPCode(t, rotatedSetup.Secret),
		},
		&rotatedEnable,
	)
	require.Equal(t, http.StatusOK, httpRsp.StatusCode)
	require.True(t, rotatedEnable.Enabled)

	httpRsp = doJSON(
		t,
		http.MethodPost,
		server.URL+"/user-svc/login",
		"",
		user.LoginRequest{
			AppHost:  sdk.DefaultTestAppHost,
			Slug:     "totp-login",
			Password: "testUserPassword0",
			TOTPCode: currentTOTPCodeDistinctFrom(t, replacementSetup.Secret, rotatedSetup.Secret),
		},
		nil,
	)
	require.Equal(t, http.StatusUnauthorized, httpRsp.StatusCode)

	var rotatedLogin user.LoginResponse
	httpRsp = doJSON(
		t,
		http.MethodPost,
		server.URL+"/user-svc/login",
		"",
		user.LoginRequest{
			AppHost:  sdk.DefaultTestAppHost,
			Slug:     "totp-login",
			Password: "testUserPassword0",
			TOTPCode: currentTOTPCode(t, rotatedSetup.Secret),
		},
		&rotatedLogin,
	)
	require.Equal(t, http.StatusOK, httpRsp.StatusCode)
	require.NotNil(t, rotatedLogin.Token)
	require.NotEmpty(t, rotatedLogin.Token.Token)

	httpRsp = doJSON(
		t,
		http.MethodDelete,
		server.URL+"/user-svc/totp",
		token,
		user.DisableTOTPRequest{
			Code: currentTOTPCode(t, rotatedSetup.Secret),
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

	httpRsp = doJSON(
		t,
		http.MethodPost,
		server.URL+"/user-svc/totp/setup",
		token,
		user.BeginTOTPSetupRequest{AccountName: "bad\naccount"},
		nil,
	)
	require.Equal(t, http.StatusBadRequest, httpRsp.StatusCode)

	httpRsp = doJSON(
		t,
		http.MethodPost,
		server.URL+"/user-svc/totp/setup",
		token,
		user.BeginTOTPSetupRequest{AccountName: "$unknown"},
		nil,
	)
	require.Equal(t, http.StatusBadRequest, httpRsp.StatusCode)

	var setup user.BeginTOTPSetupResponse
	httpRsp = doJSON(
		t,
		http.MethodPost,
		server.URL+"/user-svc/totp/setup",
		token,
		user.BeginTOTPSetupRequest{
			Issuer:      " auth.example.com ",
			AccountName: " $name <$email> ",
		},
		&setup,
	)
	require.Equal(t, http.StatusOK, httpRsp.StatusCode)
	require.NotEmpty(t, setup.Secret)

	key, err := otp.NewKeyFromURL(setup.ProvisioningURI)
	require.NoError(t, err)
	require.Equal(t, "auth.example.com", key.Issuer())
	require.Equal(t, "TOTP Custom Issuer <totp-custom-issuer@test.com>", key.AccountName())

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

func currentTOTPCodeDistinctFrom(
	t *testing.T,
	secret string,
	otherSecret string,
) string {
	t.Helper()

	deadline := time.Now().Add(65 * time.Second)
	for {
		code := currentTOTPCode(t, secret)
		otherCode := currentTOTPCode(t, otherSecret)
		if code != otherCode {
			return code
		}

		if time.Now().After(deadline) {
			t.Fatalf("TOTP codes for different secrets stayed equal")
		}
		time.Sleep(time.Second)
	}
}
