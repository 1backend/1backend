package userservice_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	openapi "github.com/1backend/1backend/clients/go"
	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/test"
	user "github.com/1backend/1backend/server/internal/services/user/types"
)

func TestApiKeyBearerReadSelfAndListRedactsSecret(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		Test: true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)
	clients, tokens, err := test.MakeClients(clientFactory, sdk.DefaultTestAppHost, 1)
	require.NoError(t, err)

	createRsp := createAPIKey(t, server.Url, tokens[0].Token, user.CreateApiKeyRequest{
		Name: "test key",
	})
	require.True(t, strings.HasPrefix(createRsp.Key, "obk_"))
	require.NotEmpty(t, createRsp.ApiKey.Id)
	require.NotEmpty(t, createRsp.ApiKey.Prefix)
	require.True(t, strings.HasPrefix(createRsp.Key, createRsp.ApiKey.Prefix))

	apiKeyClient := clientFactory.Client(client.WithToken(createRsp.Key))
	selfRsp, _, err := apiKeyClient.UserSvcAPI.ReadSelf(context.Background()).
		Body(openapi.UserSvcReadSelfRequest{CountTokens: openapi.PtrBool(true)}).
		Execute()
	require.NoError(t, err)
	require.Equal(t, "test-user-slug-0", selfRsp.User.Slug)
	require.GreaterOrEqual(t, selfRsp.TokenCount, int32(2), "API key exchange should persist a normal JWT token for its device")

	_, _, err = apiKeyClient.SecretSvcAPI.ListSecrets(context.Background()).
		Body(openapi.SecretSvcListSecretsRequest{}).
		Execute()
	require.NoError(t, err, "API keys should work as ordinary bearers across service permission checks")

	listRsp := listAPIKeys(t, server.Url, tokens[0].Token, user.ListApiKeysRequest{})
	require.Len(t, listRsp.ApiKeys, 1)
	require.Equal(t, createRsp.ApiKey.Id, listRsp.ApiKeys[0].Id)

	listJSON, err := json.Marshal(listRsp)
	require.NoError(t, err)
	require.NotContains(t, string(listJSON), createRsp.Key)
	require.NotContains(t, string(listJSON), "secretHash")

	// The original JWT client should still be normal and unaffected.
	jwtSelfRsp, _, err := clients[0].UserSvcAPI.ReadSelf(context.Background()).Execute()
	require.NoError(t, err)
	require.Equal(t, "test-user-slug-0", jwtSelfRsp.User.Slug)
}

func TestApiKeyExchangeUsesKeyDeviceAndExplicitActiveOrganization(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		Test: true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)
	clients, tokens, err := test.MakeClients(clientFactory, sdk.DefaultTestAppHost, 1)
	require.NoError(t, err)
	userClient := clients[0]

	publicKeyRsp, _, err := clientFactory.Client().UserSvcAPI.GetPublicKey(context.Background()).Execute()
	require.NoError(t, err)

	org1 := createOrg(t, clientFactory, tokens[0].Token, "Org 1", "api-key-org-1")
	org2 := createOrg(t, clientFactory, tokens[0].Token, "Org 2", "api-key-org-2")

	activateRsp, _, err := userClient.UserSvcAPI.ActivateOrganization(context.Background()).
		Body(openapi.UserSvcActivateOrganizationRequest{OrganizationId: org2}).
		Execute()
	require.NoError(t, err)
	activeOrg2Token := activateRsp.Token.Token

	createRsp := createAPIKey(t, server.Url, activeOrg2Token, user.CreateApiKeyRequest{
		Name:                 "org 1 key",
		ActiveOrganizationId: &org1,
	})

	exchangeRsp, exchangeHTTPRsp := exchangeAPIKey(t, server.Url, createRsp.Key)
	require.Equal(t, http.StatusOK, exchangeHTTPRsp.StatusCode)
	require.NotNil(t, exchangeRsp.Token)
	require.NotEmpty(t, exchangeRsp.Token.Token)

	claims, err := (auth.AuthorizerImpl{}).ParseJWT(
		publicKeyRsp.PublicKey,
		exchangeRsp.Token.Token,
	)
	require.NoError(t, err)
	require.Equal(t, "api-key:"+createRsp.ApiKey.Id, claims.Device)
	require.Equal(t, org1, claims.ActiveOrganizationId)
	require.Contains(t, claims.Roles, fmt.Sprintf("user-svc:org:{%s}:admin", org1))
	require.NotContains(t, claims.Roles, fmt.Sprintf("user-svc:org:{%s}:admin", org2))

	selfWithAPIKey, _, err := clientFactory.Client(client.WithToken(createRsp.Key)).
		UserSvcAPI.ReadSelf(context.Background()).Execute()
	require.NoError(t, err)
	require.Equal(t, org1, selfWithAPIKey.GetActiveOrganizationId())

	selfWithOriginalToken, _, err := clientFactory.Client(client.WithToken(activeOrg2Token)).
		UserSvcAPI.ReadSelf(context.Background()).Execute()
	require.NoError(t, err)
	require.Equal(t, org2, selfWithOriginalToken.GetActiveOrganizationId())
}

func TestApiKeyRevocationExpiryAndTampering(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		Test: true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)
	_, tokens, err := test.MakeClients(clientFactory, sdk.DefaultTestAppHost, 1)
	require.NoError(t, err)

	cached := createAPIKey(t, server.Url, tokens[0].Token, user.CreateApiKeyRequest{
		Name: "cached key",
	})
	_, _, err = clientFactory.Client(client.WithToken(cached.Key)).
		UserSvcAPI.ReadSelf(context.Background()).Execute()
	require.NoError(t, err)
	revokeAPIKeys(t, server.Url, tokens[0].Token, user.RevokeApiKeysRequest{
		Id: cached.ApiKey.Id,
	})
	_, cachedHTTPRsp, err := clientFactory.Client(client.WithToken(cached.Key)).
		UserSvcAPI.ReadSelf(context.Background()).Execute()
	require.Error(t, err)
	require.Equal(t, http.StatusUnauthorized, cachedHTTPRsp.StatusCode)

	revoked := createAPIKey(t, server.Url, tokens[0].Token, user.CreateApiKeyRequest{
		Name: "revoked key",
	})
	revokeAPIKeys(t, server.Url, tokens[0].Token, user.RevokeApiKeysRequest{
		Id: revoked.ApiKey.Id,
	})

	_, revokedHTTPRsp := exchangeAPIKey(t, server.Url, revoked.Key)
	require.Equal(t, http.StatusUnauthorized, revokedHTTPRsp.StatusCode)

	_, httpRsp, err := clientFactory.Client(client.WithToken(revoked.Key)).
		UserSvcAPI.ReadSelf(context.Background()).Execute()
	require.Error(t, err)
	require.Equal(t, http.StatusUnauthorized, httpRsp.StatusCode)

	expiresAt := time.Now().Add(500 * time.Millisecond)
	expiring := createAPIKey(t, server.Url, tokens[0].Token, user.CreateApiKeyRequest{
		Name:      "expiring key",
		ExpiresAt: &expiresAt,
	})
	_, validHTTPRsp := exchangeAPIKey(t, server.Url, expiring.Key)
	require.Equal(t, http.StatusOK, validHTTPRsp.StatusCode)

	time.Sleep(800 * time.Millisecond)

	_, expiredHTTPRsp := exchangeAPIKey(t, server.Url, expiring.Key)
	require.Equal(t, http.StatusUnauthorized, expiredHTTPRsp.StatusCode)

	tampered := expiring.Key + "x"
	_, tamperedHTTPRsp := exchangeAPIKey(t, server.Url, tampered)
	require.Equal(t, http.StatusUnauthorized, tamperedHTTPRsp.StatusCode)
}

func TestApiKeySelfServiceAndOnBehalfAuthorization(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		Test: true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)
	_, tokens, err := test.MakeClients(clientFactory, sdk.DefaultTestAppHost, 2)
	require.NoError(t, err)

	selfRsp := createAPIKey(t, server.Url, tokens[0].Token, user.CreateApiKeyRequest{
		Name: "self key",
	})
	require.NotEmpty(t, selfRsp.Key)

	httpRsp := doJSON(t, http.MethodPost, server.Url+"/user-svc/api-key", tokens[0].Token, user.CreateApiKeyRequest{
		UserId: tokens[1].UserId,
		Name:   "other user key",
	}, nil)
	require.Equal(t, http.StatusUnauthorized, httpRsp.StatusCode)
}

func createOrg(
	t *testing.T,
	clientFactory client.ClientFactory,
	token string,
	name string,
	slug string,
) string {
	t.Helper()

	rsp, _, err := clientFactory.Client(client.WithToken(token)).
		UserSvcAPI.SaveOrganization(context.Background()).
		Body(openapi.UserSvcSaveOrganizationRequest{
			Activate: openapi.PtrBool(false),
			Name:     openapi.PtrString(name),
			Slug:     slug,
		}).
		Execute()
	require.NoError(t, err)
	return rsp.Organization.Id
}

func createAPIKey(
	t *testing.T,
	serverURL string,
	token string,
	req user.CreateApiKeyRequest,
) user.CreateApiKeyResponse {
	t.Helper()

	var rsp user.CreateApiKeyResponse
	httpRsp := doJSON(t, http.MethodPost, serverURL+"/user-svc/api-key", token, req, &rsp)
	require.Equal(t, http.StatusOK, httpRsp.StatusCode)
	return rsp
}

func listAPIKeys(
	t *testing.T,
	serverURL string,
	token string,
	req user.ListApiKeysRequest,
) user.ListApiKeysResponse {
	t.Helper()

	var rsp user.ListApiKeysResponse
	httpRsp := doJSON(t, http.MethodPost, serverURL+"/user-svc/api-keys", token, req, &rsp)
	require.Equal(t, http.StatusOK, httpRsp.StatusCode)
	return rsp
}

func revokeAPIKeys(
	t *testing.T,
	serverURL string,
	token string,
	req user.RevokeApiKeysRequest,
) {
	t.Helper()

	httpRsp := doJSON(t, http.MethodDelete, serverURL+"/user-svc/api-keys", token, req, nil)
	require.Equal(t, http.StatusOK, httpRsp.StatusCode)
}

func exchangeAPIKey(
	t *testing.T,
	serverURL string,
	apiKey string,
) (user.ExchangeApiKeyResponse, *http.Response) {
	t.Helper()

	var rsp user.ExchangeApiKeyResponse
	httpRsp := doJSON(t, http.MethodPost, serverURL+"/user-svc/api-key/exchange", apiKey, nil, &rsp)
	return rsp, httpRsp
}

func doJSON(
	t *testing.T,
	method string,
	url string,
	bearer string,
	body any,
	out any,
) *http.Response {
	t.Helper()

	var reader io.Reader
	if body != nil {
		bs, err := json.Marshal(body)
		require.NoError(t, err)
		reader = bytes.NewReader(bs)
	}

	req, err := http.NewRequest(method, url, reader)
	require.NoError(t, err)
	req.Header.Set("Authorization", "Bearer "+bearer)
	req.Header.Set("Accept", "application/json")
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	httpRsp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer httpRsp.Body.Close()

	bs, err := io.ReadAll(httpRsp.Body)
	require.NoError(t, err)
	if out != nil && len(bs) > 0 && httpRsp.StatusCode >= 200 && httpRsp.StatusCode < 300 {
		require.NoError(t, json.Unmarshal(bs, out), string(bs))
	}

	return httpRsp
}
