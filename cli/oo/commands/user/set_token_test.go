package user

import (
	"context"
	"testing"

	types "github.com/1backend/1backend/cli/oo/types"
	"github.com/1backend/1backend/cli/oo/util"
	openapi "github.com/1backend/1backend/clients/go"
	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/test"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)

func TestSetTokenStoresExactTokenForActualUser(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		Test: true,
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		server.Cleanup(t)
	})

	t.Setenv("HOME", t.TempDir())

	conf := types.Config{
		SelectedEnvironment: "test",
		Environments: map[string]*types.Environment{
			"test": {
				ShortName: "test",
				URL:       server.Url,
			},
		},
	}
	require.NoError(t, util.SaveConfig(conf))

	clientFactory := client.NewApiClientFactory(server.Url)
	loginRsp, _, err := clientFactory.Client().
		UserSvcAPI.Login(context.Background()).
		Body(openapi.UserSvcLoginRequest{
			AppHost:  sdk.DefaultTestAppHost,
			Slug:     openapi.PtrString("test-user-slug-0"),
			Password: openapi.PtrString("testUserPassword0"),
		}).
		Execute()
	require.NoError(t, err)

	cmd := &cobra.Command{}
	cmd.SetContext(context.Background())

	require.NoError(t, SetToken(cmd, []string{loginRsp.Token.Token}))

	publicKeyRsp, _, err := clientFactory.Client().
		UserSvcAPI.GetPublicKey(context.Background()).
		Execute()
	require.NoError(t, err)

	loaded, err := util.LoadConfig()
	require.NoError(t, err)

	env := loaded.Environments["test"]
	require.Equal(t, "test-user-slug-0", env.SelectedUser)

	usr := env.Users["test-user-slug-0"]
	require.NotNil(t, usr)
	require.NotEmpty(t, usr.SelectedAppHost)

	storedToken := usr.TokensByAppHost[usr.SelectedAppHost]
	require.NotEmpty(t, storedToken)
	require.Equal(t, loginRsp.Token.Token, storedToken)

	claims, err := auth.AuthorizerImpl{}.ParseJWT(publicKeyRsp.PublicKey, storedToken)
	require.NoError(t, err)
	require.Equal(t, "test-user-slug-0", claims.Slug)
	require.Equal(t, "unknown", claims.Device)
}
