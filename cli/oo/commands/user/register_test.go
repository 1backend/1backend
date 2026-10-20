package user

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	types "github.com/1backend/1backend/cli/oo/types"
	"github.com/1backend/1backend/cli/oo/util"
	openapi "github.com/1backend/1backend/clients/go"
	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)

func TestRegisterSelectsAppHostAndStoresToken(t *testing.T) {
	t.Setenv("HOME", t.TempDir())

	const token = "test-register-token"

	var registerReq openapi.UserSvcRegisterRequest
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		require.Equal(t, "/user-svc/register", r.URL.Path)
		require.NoError(t, json.NewDecoder(r.Body).Decode(&registerReq))

		w.Header().Set("Content-Type", "application/json")
		require.NoError(t, json.NewEncoder(w).Encode(openapi.UserSvcRegisterResponse{
			Token: &openapi.UserSvcToken{Token: token},
		}))
	}))
	defer server.Close()

	require.NoError(t, util.SaveConfig(types.Config{
		SelectedEnvironment: "test",
		Environments: map[string]*types.Environment{
			"test": {
				ShortName: "test",
				URL:       server.URL,
			},
		},
	}))

	cmd := &cobra.Command{}
	cmd.SetContext(context.Background())

	require.NoError(t, Register(cmd, []string{"new-user", "testPassword1"}))

	require.Equal(t, sdk.DefaultAppHost, registerReq.AppHost)
	require.Equal(t, "new-user", registerReq.Slug)

	loaded, err := util.LoadConfig()
	require.NoError(t, err)

	env := loaded.Environments["test"]
	require.Equal(t, "new-user", env.SelectedUser)

	usr := env.Users["new-user"]
	require.NotNil(t, usr)
	require.Equal(t, sdk.DefaultAppHost, usr.SelectedAppHost)
	require.Equal(t, token, usr.TokensByAppHost[usr.SelectedAppHost])

	// The freshly stored token must be the one later commands use.
	url, storedToken, err := util.GetSelectedUrlAndToken(cmd)
	require.NoError(t, err)
	require.Equal(t, server.URL, url)
	require.Equal(t, token, storedToken)
}

func TestRegisterHandlesExistingUserWithoutTokens(t *testing.T) {
	t.Setenv("HOME", t.TempDir())

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		require.NoError(t, json.NewEncoder(w).Encode(openapi.UserSvcRegisterResponse{
			Token: &openapi.UserSvcToken{Token: "test-register-token"},
		}))
	}))
	defer server.Close()

	// Config files written by other tools or earlier versions may hold a user
	// entry without a tokensByAppHost mapping.
	configDir := filepath.Join(os.Getenv("HOME"), ".1backend")
	require.NoError(t, os.MkdirAll(configDir, 0755))
	require.NoError(t, os.WriteFile(
		filepath.Join(configDir, "cliConfig.yaml"),
		[]byte(`selectedEnvironment: test
environments:
  test:
    url: `+server.URL+`
    shortName: test
    users:
      existing-user:
        slug: existing-user
        tokensByAppHost: null
`),
		0644,
	))

	cmd := &cobra.Command{}
	cmd.SetContext(context.Background())

	require.NotPanics(t, func() {
		require.NoError(t, Register(cmd, []string{"existing-user", "testPassword1"}))
	})

	loaded, err := util.LoadConfig()
	require.NoError(t, err)

	usr := loaded.Environments["test"].Users["existing-user"]
	require.NotNil(t, usr)
	require.Equal(t, sdk.DefaultAppHost, usr.SelectedAppHost)
	require.Equal(t, "test-register-token", usr.TokensByAppHost[usr.SelectedAppHost])
}
