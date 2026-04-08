package userservice_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/test"

	openapi "github.com/1backend/1backend/clients/go"
)

func TestActivateOrganization(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		Test: true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)

	clients, _, err := test.MakeClients(clientFactory, sdk.DefaultTestAppHost, 2)
	require.NoError(t, err)

	userClient := clients[0]
	otherClient := clients[1]

	publicKeyRsp, _, err := clientFactory.Client().
		UserSvcAPI.GetPublicKey(context.Background()).
		Execute()
	require.NoError(t, err)

	createOrg := func(name, slug string, token string) (string, string) {
		rsp, _, err := clientFactory.Client(client.WithToken(token)).
			UserSvcAPI.SaveOrganization(context.Background()).
			Body(openapi.UserSvcSaveOrganizationRequest{
				Name: openapi.PtrString(name),
				Slug: slug,
			}).
			Execute()
		require.NoError(t, err)

		return rsp.Organization.Id, rsp.Token.Token
	}

	tokenFromResponse := func(resp *http.Response) string {
		defer resp.Body.Close()

		var payload struct {
			Token struct {
				Token string `json:"token"`
			} `json:"token"`
		}

		err := json.NewDecoder(resp.Body).Decode(&payload)
		require.NoError(t, err)
		require.NotEmpty(t, payload.Token.Token)

		return payload.Token.Token
	}

	activate := func(token string, organizationId string) *http.Response {
		body, err := json.Marshal(map[string]string{
			"organizationId": organizationId,
		})
		require.NoError(t, err)

		req, err := http.NewRequest(
			http.MethodPost,
			server.Url+"/user-svc/organization/activate",
			bytes.NewReader(body),
		)
		require.NoError(t, err)
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		require.NoError(t, err)

		return resp
	}

	loginRsp, _, err := userClient.UserSvcAPI.Login(context.Background()).
		Body(openapi.UserSvcLoginRequest{
			AppHost:  sdk.DefaultTestAppHost,
			Slug:     openapi.PtrString("test-user-slug-0"),
			Password: openapi.PtrString("testUserPassword0"),
		}).
		Execute()
	require.NoError(t, err)

	orgId1, currentToken := createOrg("Org 1", "org-1", loginRsp.Token.Token)
	orgId2, currentToken := createOrg("Org 2", "org-2", currentToken)

	t.Run("switch active organization and mint a fresh token", func(t *testing.T) {
		resp := activate(currentToken, orgId1)
		require.Equal(t, http.StatusOK, resp.StatusCode)

		newToken := tokenFromResponse(resp)

		claims, err := (auth.AuthorizerImpl{}).ParseJWT(
			publicKeyRsp.PublicKey,
			newToken,
		)
		require.NoError(t, err)
		require.Equal(t, orgId1, claims.ActiveOrganizationId)

		readSelfRsp, _, err := clientFactory.Client(client.WithToken(newToken)).
			UserSvcAPI.ReadSelf(context.Background()).
			Execute()
		require.NoError(t, err)
		require.Equal(t, openapi.PtrString(orgId1), readSelfRsp.ActiveOrganizationId)
		require.Len(t, readSelfRsp.Organizations, 2)
	})

	t.Run("keep current active organization unchanged on invalid target", func(t *testing.T) {
		resp := activate(currentToken, "org-does-not-exist")
		defer resp.Body.Close()
		require.Equal(t, http.StatusUnauthorized, resp.StatusCode)

		readSelfRsp, _, err := clientFactory.Client(client.WithToken(currentToken)).
			UserSvcAPI.ReadSelf(context.Background()).
			Execute()
		require.NoError(t, err)
		require.Equal(t, openapi.PtrString(orgId2), readSelfRsp.ActiveOrganizationId)
	})

	t.Run("reject switching to an organization the caller is not a member of", func(t *testing.T) {
		otherLoginRsp, _, err := otherClient.UserSvcAPI.Login(context.Background()).
			Body(openapi.UserSvcLoginRequest{
				AppHost:  sdk.DefaultTestAppHost,
				Slug:     openapi.PtrString("test-user-slug-1"),
				Password: openapi.PtrString("testUserPassword1"),
			}).
			Execute()
		require.NoError(t, err)

		resp := activate(otherLoginRsp.Token.Token, orgId2)
		defer resp.Body.Close()
		require.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	})
}
