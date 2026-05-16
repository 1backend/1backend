package api_key

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	types "github.com/1backend/1backend/cli/oo/types"
	"github.com/1backend/1backend/cli/oo/util"
	openapi "github.com/1backend/1backend/clients/go"
	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/client"
	sdktest "github.com/1backend/1backend/sdk/go/test"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)

func TestCreateSendsAPIKeyRequest(t *testing.T) {
	t.Setenv("HOME", t.TempDir())

	expiresAt := time.Now().Add(time.Hour).UTC().Truncate(time.Second)
	activeOrgId := "org-test"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		require.Equal(t, http.MethodPost, r.Method)
		require.Equal(t, "/user-svc/api-key", r.URL.Path)
		require.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))

		req := openapi.UserSvcCreateApiKeyRequest{}
		require.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		require.Equal(t, "usr-service", req.GetUserId())
		require.Equal(t, "service-key", req.GetName())
		require.Equal(t, "service.app", req.GetAppHost())
		require.NotNil(t, req.ActiveOrganizationId)
		require.Equal(t, activeOrgId, *req.ActiveOrganizationId)
		require.Equal(t, expiresAt.Format(time.RFC3339), req.GetExpiresAt())

		require.NoError(t, json.NewEncoder(w).Encode(openapi.UserSvcCreateApiKeyResponse{
			ApiKey: openapi.UserSvcApiKeyView{
				Id:                   "apik-test",
				AppId:                "app-test",
				UserId:               "usr-service",
				CreatedAt:            time.Now().Format(time.RFC3339),
				UpdatedAt:            time.Now().Format(time.RFC3339),
				Name:                 openapi.PtrString("service-key"),
				Prefix:               openapi.PtrString("obk_apik-test"),
				ActiveOrganizationId: openapi.PtrString(activeOrgId),
				ExpiresAt:            openapi.PtrString(expiresAt.Format(time.RFC3339)),
			},
			Key: "obk_apik-test_secret",
		}))
	}))
	defer server.Close()

	cmd := testCommand(t, server.URL)
	err := Create(cmd, []string{"service-key"}, createOptions{
		userId:               "usr-service",
		appHost:              "service.app",
		activeOrganizationId: activeOrgId,
		expiresAt:            expiresAt.Format(time.RFC3339),
		keyOnly:              true,
	}, true)
	require.NoError(t, err)
}

func TestListSendsAPIKeyRequest(t *testing.T) {
	t.Setenv("HOME", t.TempDir())

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		require.Equal(t, http.MethodPost, r.Method)
		require.Equal(t, "/user-svc/api-keys", r.URL.Path)
		require.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))

		req := openapi.UserSvcListApiKeysRequest{}
		require.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		require.Equal(t, "usr-service", req.GetUserId())
		require.Equal(t, "app-test", req.GetAppId())
		require.Equal(t, []string{"apik-a", "apik-b"}, req.Ids)
		require.True(t, req.GetIncludeRevoked())

		require.NoError(t, json.NewEncoder(w).Encode(openapi.UserSvcListApiKeysResponse{
			ApiKeys: []openapi.UserSvcApiKeyView{
				{
					Id:        "apik-a",
					AppId:     "app-test",
					UserId:    "usr-service",
					Name:      openapi.PtrString("service-key"),
					Prefix:    openapi.PtrString("obk_apik-a"),
					CreatedAt: time.Now().Format(time.RFC3339),
					UpdatedAt: time.Now().Format(time.RFC3339),
				},
			},
		}))
	}))
	defer server.Close()

	cmd := testCommand(t, server.URL)
	err := List(cmd, listOptions{
		userId:         "usr-service",
		appId:          "app-test",
		ids:            []string{"apik-a", "apik-b"},
		includeRevoked: true,
	})
	require.NoError(t, err)
}

func TestRevokeSendsAPIKeyRequest(t *testing.T) {
	t.Setenv("HOME", t.TempDir())

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodDelete, r.Method)
		require.Equal(t, "/user-svc/api-keys", r.URL.Path)
		require.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))

		req := openapi.UserSvcRevokeApiKeysRequest{}
		require.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		require.Equal(t, "usr-service", req.GetUserId())
		require.Equal(t, []string{"apik-flag", "apik-arg"}, req.Ids)

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	cmd := testCommand(t, server.URL)
	err := Revoke(cmd, []string{"apik-arg"}, revokeOptions{
		userId: "usr-service",
		ids:    []string{"apik-flag"},
	})
	require.NoError(t, err)
}

func TestAPIKeyCLIIntegrationCreateAndListSelfAndOnBehalf(t *testing.T) {
	server, err := sdktest.StartService(sdktest.Options{
		Test: true,
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		server.Cleanup(t)
	})

	t.Setenv("HOME", t.TempDir())

	clientFactory := client.NewApiClientFactory(server.Url)
	_, userTokens, err := sdktest.MakeClients(clientFactory, sdk.DefaultTestAppHost, 2)
	require.NoError(t, err)

	_, adminToken, err := sdktest.AdminClient(clientFactory, sdk.DefaultTestAppHost)
	require.NoError(t, err)

	selfCmd := testCommandWithToken(t, server.Url, "self", userTokens[0].Token)
	selfCreateOut := captureStdout(t, func() {
		err = Create(selfCmd, []string{"self-cli-key"}, createOptions{}, false)
	})
	require.NoError(t, err)
	require.Contains(t, selfCreateOut, "obk_")

	selfListOut := captureStdout(t, func() {
		err = List(selfCmd, listOptions{})
	})
	require.NoError(t, err)
	require.Contains(t, selfListOut, "self-cli-key")
	require.Contains(t, selfListOut, userTokens[0].UserId)

	adminCmd := testCommandWithToken(t, server.Url, "admin", adminToken)
	onBehalfCreateOut := captureStdout(t, func() {
		err = Create(adminCmd, []string{"on-behalf-cli-key"}, createOptions{
			userId: userTokens[1].UserId,
		}, false)
	})
	require.NoError(t, err)
	require.Contains(t, onBehalfCreateOut, "obk_")
	require.Contains(t, onBehalfCreateOut, userTokens[1].UserId)

	onBehalfListOut := captureStdout(t, func() {
		err = List(adminCmd, listOptions{
			userId: userTokens[1].UserId,
		})
	})
	require.NoError(t, err)
	require.Contains(t, onBehalfListOut, "on-behalf-cli-key")
	require.Contains(t, onBehalfListOut, userTokens[1].UserId)
	require.NotContains(t, onBehalfListOut, "self-cli-key")
}

func TestParseExpiration(t *testing.T) {
	_, err := parseExpiration(time.Now().Format(time.RFC3339), "1h")
	require.Error(t, err)

	at := time.Now().Add(time.Hour).UTC().Truncate(time.Second)
	parsedAt, err := parseExpiration(at.Format(time.RFC3339), "")
	require.NoError(t, err)
	require.Equal(t, at.Format(time.RFC3339), parsedAt)

	parsedIn, err := parseExpiration("", "1h")
	require.NoError(t, err)
	parsedInTime, err := time.Parse(time.RFC3339, parsedIn)
	require.NoError(t, err)
	require.True(t, parsedInTime.After(time.Now().Add(55*time.Minute)))
}

func testCommand(t *testing.T, url string) *cobra.Command {
	t.Helper()

	return testCommandWithToken(t, url, "tester", "test-token")
}

func testCommandWithToken(t *testing.T, url string, slug string, token string) *cobra.Command {
	t.Helper()

	require.NoError(t, util.SaveConfig(types.Config{
		SelectedEnvironment: "test",
		Environments: map[string]*types.Environment{
			"test": {
				ShortName:    "test",
				URL:          url,
				SelectedUser: slug,
				Users: map[string]*types.User{
					slug: {
						Slug:            slug,
						SelectedAppHost: "test.app",
						TokensByAppHost: map[string]string{
							"test.app": token,
						},
					},
				},
			},
		},
	}))

	cmd := &cobra.Command{}
	cmd.SetContext(context.Background())
	return cmd
}

func captureStdout(t *testing.T, fn func()) string {
	t.Helper()

	oldStdout := os.Stdout
	readFile, writeFile, err := os.Pipe()
	require.NoError(t, err)
	os.Stdout = writeFile

	fn()

	require.NoError(t, writeFile.Close())
	os.Stdout = oldStdout

	out, err := io.ReadAll(readFile)
	require.NoError(t, err)
	require.NoError(t, readFile.Close())

	return strings.TrimSpace(string(out))
}
