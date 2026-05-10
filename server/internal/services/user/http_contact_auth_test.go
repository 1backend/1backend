package userservice_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"

	openapi "github.com/1backend/1backend/clients/go"
	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/server/internal/di"
	usertypes "github.com/1backend/1backend/server/internal/services/user/types"
	"github.com/1backend/1backend/server/internal/universe"
)

type fakeContactAuthVerifier struct {
	claims *universe.ContactAuthClaims
	err    error
}

func (f fakeContactAuthVerifier) VerifyContactAuthToken(
	ctx context.Context,
	provider string,
	token string,
) (*universe.ContactAuthClaims, error) {
	if f.err != nil {
		return nil, f.err
	}
	return f.claims, nil
}

func TestContactAuthGoogleLoginExistingContact(t *testing.T) {
	hs := &di.HandlerSwitcher{}
	server := httptest.NewServer(hs)
	defer server.Close()

	options := &universe.Options{
		Test:           true,
		Url:            server.URL,
		VerifyContacts: true,
		ContactAuthVerifier: fakeContactAuthVerifier{
			claims: &universe.ContactAuthClaims{
				Email:         "owner@example.com",
				EmailVerified: true,
				Name:          "Contact Owner",
			},
		},
	}
	u, err := di.BigBang(options)
	require.NoError(t, err)
	hs.UpdateHandler(u.Router)
	require.NoError(t, u.StarterFunc())

	ctx := context.Background()
	userSvc := options.ClientFactory.Client().UserSvcAPI
	otpRsp, _, err := userSvc.SendOtp(ctx).Body(openapi.UserSvcSendOtpRequest{
		AppHost:         sdk.DefaultTestAppHost,
		ContactId:       "owner@example.com",
		ContactPlatform: "email",
	}).Execute()
	require.NoError(t, err)

	_, _, err = userSvc.Register(ctx).Body(openapi.UserSvcRegisterRequest{
		AppHost: sdk.DefaultTestAppHost,
		Slug:    "contact-owner",
		Contact: &openapi.UserSvcContactInput{
			Id:       "owner@example.com",
			Platform: "email",
			OtpId:    &otpRsp.OtpId,
			OtpCode:  otpRsp.Code,
		},
	}).Execute()
	require.NoError(t, err)

	loginRsp := postContactAuthLogin(t, server.URL, "google", map[string]any{
		"appHost": sdk.DefaultTestAppHost,
		"token":   "fake-google-token",
	})
	require.NotNil(t, loginRsp.Token)
	require.NotEmpty(t, loginRsp.Token.Token)

	publicKeyRsp, _, err := userSvc.GetPublicKey(ctx).Execute()
	require.NoError(t, err)
	claims, err := options.Authorizer.ParseJWT(publicKeyRsp.PublicKey, loginRsp.Token.Token)
	require.NoError(t, err)
	require.Equal(t, "contact-owner", claims.Slug)
}

func TestContactAuthGoogleLoginCreatesVerifiedContactUser(t *testing.T) {
	hs := &di.HandlerSwitcher{}
	server := httptest.NewServer(hs)
	defer server.Close()

	options := &universe.Options{
		Test:           true,
		Url:            server.URL,
		VerifyContacts: true,
		ContactAuthVerifier: fakeContactAuthVerifier{
			claims: &universe.ContactAuthClaims{
				Email:         "new.user+tag@example.com",
				EmailVerified: true,
				Name:          "New User",
			},
		},
	}
	u, err := di.BigBang(options)
	require.NoError(t, err)
	hs.UpdateHandler(u.Router)
	require.NoError(t, u.StarterFunc())

	loginRsp := postContactAuthLogin(t, server.URL, "google", map[string]any{
		"appHost": sdk.DefaultTestAppHost,
		"token":   "fake-google-token",
	})
	require.NotNil(t, loginRsp.Token)
	require.NotEmpty(t, loginRsp.Token.Token)

	readSelfRsp, _, err := client.NewApiClientFactory(server.URL).
		Client(client.WithToken(loginRsp.Token.Token)).
		UserSvcAPI.ReadSelf(context.Background()).
		Body(openapi.UserSvcReadSelfRequest{}).
		Execute()
	require.NoError(t, err)
	require.Equal(t, "new-user", readSelfRsp.User.Slug)
	require.Len(t, readSelfRsp.Contacts, 1)
	require.Equal(t, "new.user@example.com", readSelfRsp.Contacts[0].Id)
	require.True(t, readSelfRsp.Contacts[0].GetVerified())
}

func TestContactAuthRejectsUnverifiedEmail(t *testing.T) {
	hs := &di.HandlerSwitcher{}
	server := httptest.NewServer(hs)
	defer server.Close()

	options := &universe.Options{
		Test: true,
		Url:  server.URL,
		ContactAuthVerifier: fakeContactAuthVerifier{
			claims: &universe.ContactAuthClaims{
				Email:         "unverified@example.com",
				EmailVerified: false,
			},
		},
	}
	u, err := di.BigBang(options)
	require.NoError(t, err)
	hs.UpdateHandler(u.Router)
	require.NoError(t, u.StarterFunc())

	resp := postContactAuthLoginRaw(t, server.URL, "google", map[string]any{
		"appHost": sdk.DefaultTestAppHost,
		"token":   "fake-google-token",
	})
	defer resp.Body.Close()
	require.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

func TestContactAuthRejectsRelayEmail(t *testing.T) {
	hs := &di.HandlerSwitcher{}
	server := httptest.NewServer(hs)
	defer server.Close()

	options := &universe.Options{
		Test: true,
		Url:  server.URL,
		ContactAuthVerifier: fakeContactAuthVerifier{
			claims: &universe.ContactAuthClaims{
				Email:         "123+person@users.noreply.github.com",
				EmailVerified: true,
			},
		},
	}
	u, err := di.BigBang(options)
	require.NoError(t, err)
	hs.UpdateHandler(u.Router)
	require.NoError(t, u.StarterFunc())

	resp := postContactAuthLoginRaw(t, server.URL, "google", map[string]any{
		"appHost": sdk.DefaultTestAppHost,
		"token":   "fake-google-token",
	})
	defer resp.Body.Close()
	require.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

func TestContactAuthRejectsAppleProvider(t *testing.T) {
	hs := &di.HandlerSwitcher{}
	server := httptest.NewServer(hs)
	defer server.Close()

	options := &universe.Options{
		Test: true,
		Url:  server.URL,
		ContactAuthVerifier: fakeContactAuthVerifier{
			claims: &universe.ContactAuthClaims{
				Email:         "person@example.com",
				EmailVerified: true,
			},
		},
	}
	u, err := di.BigBang(options)
	require.NoError(t, err)
	hs.UpdateHandler(u.Router)
	require.NoError(t, u.StarterFunc())

	resp := postContactAuthLoginRaw(t, server.URL, "apple", map[string]any{
		"appHost": sdk.DefaultTestAppHost,
		"token":   "fake-apple-token",
	})
	defer resp.Body.Close()
	require.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

func TestContactAuthStartRejectsCrossHostReturnTo(t *testing.T) {
	hs := &di.HandlerSwitcher{}
	server := httptest.NewServer(hs)
	defer server.Close()

	options := &universe.Options{
		Test: true,
		Url:  server.URL,
		ContactAuthProviders: map[string]universe.ContactAuthProviderConfig{
			"google": {
				ClientID: "google-client-id",
			},
		},
	}
	u, err := di.BigBang(options)
	require.NoError(t, err)
	hs.UpdateHandler(u.Router)
	require.NoError(t, u.StarterFunc())

	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf(
			"%s/user-svc/auth/google/start?appHost=%s&returnTo=%s",
			server.URL,
			url.QueryEscape(sdk.DefaultTestAppHost),
			url.QueryEscape("https://evil.example/callback"),
		),
		nil,
	)
	require.NoError(t, err)

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()
	require.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestListContactAuthProviders(t *testing.T) {
	hs := &di.HandlerSwitcher{}
	server := httptest.NewServer(hs)
	defer server.Close()

	options := &universe.Options{
		Test: true,
		Url:  server.URL,
		ContactAuthProviders: map[string]universe.ContactAuthProviderConfig{
			"google": {
				ClientID:     "google-client-id",
				ClientSecret: "google-client-secret",
			},
			"facebook": {
				ClientID:     "facebook-app-id",
				ClientSecret: "facebook-app-secret",
			},
			"slack": {
				ClientID:     "slack-client-id",
				ClientSecret: "slack-client-secret",
			},
			"linkedin": {
				ClientID:     "linkedin-client-id",
				ClientSecret: "linkedin-client-secret",
			},
			"gitlab": {
				ClientID:     "gitlab-client-id",
				ClientSecret: "gitlab-client-secret",
			},
			"github": {
				ClientID:     "github-client-id",
				ClientSecret: "github-client-secret",
			},
			"discord": {
				ClientID:     "discord-client-id",
				ClientSecret: "discord-client-secret",
			},
			"okta": {
				IssuerURL:    "https://example.okta.com",
				ClientID:     "okta-client-id",
				ClientSecret: "okta-client-secret",
			},
			"apple": {
				IssuerURL:    "https://appleid.apple.com",
				ClientID:     "apple-client-id",
				ClientSecret: "apple-client-secret",
			},
		},
	}
	u, err := di.BigBang(options)
	require.NoError(t, err)
	hs.UpdateHandler(u.Router)
	require.NoError(t, u.StarterFunc())

	resp, err := http.Get(server.URL + "/user-svc/auth/providers")
	require.NoError(t, err)
	defer resp.Body.Close()
	require.Equal(t, http.StatusOK, resp.StatusCode)

	body := usertypes.ListContactAuthProvidersResponse{}
	require.NoError(t, json.NewDecoder(resp.Body).Decode(&body))

	providers := map[string]usertypes.ContactAuthProviderInfo{}
	for _, provider := range body.Providers {
		providers[provider.Id] = provider
	}

	require.Contains(t, providers, "google")
	require.Equal(t, "oidc", providers["google"].Kind)
	require.Contains(t, providers, "slack")
	require.Contains(t, providers, "linkedin")
	require.Contains(t, providers, "gitlab")
	require.Contains(t, providers, "facebook")
	require.Equal(t, "oauth2", providers["facebook"].Kind)
	require.Contains(t, providers, "github")
	require.Equal(t, "oauth2", providers["github"].Kind)
	require.Contains(t, providers, "discord")
	require.Equal(t, "oauth2", providers["discord"].Kind)
	require.Contains(t, providers, "okta")
	require.NotContains(t, providers, "apple")
	require.Equal(t, "/user-svc/auth/google/start", providers["google"].StartUrl)
	require.Equal(t, "/user-svc/auth/github/login", providers["github"].LoginUrl)
}

func postContactAuthLogin(
	t *testing.T,
	baseURL string,
	provider string,
	body map[string]any,
) usertypes.ContactAuthLoginResponse {
	t.Helper()

	resp := postContactAuthLoginRaw(t, baseURL, provider, body)
	defer resp.Body.Close()

	bs, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode, string(bs))

	loginRsp := usertypes.ContactAuthLoginResponse{}
	require.NoError(t, json.Unmarshal(bs, &loginRsp))
	return loginRsp
}

func postContactAuthLoginRaw(
	t *testing.T,
	baseURL string,
	provider string,
	body map[string]any,
) *http.Response {
	t.Helper()

	bs, err := json.Marshal(body)
	require.NoError(t, err)
	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/user-svc/auth/%s/login", baseURL, provider),
		bytes.NewReader(bs),
	)
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	return resp
}
