package modelservice_test

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/openorch/openorch/sdk/go"
	"github.com/openorch/openorch/sdk/go/test"
	"github.com/openorch/openorch/server/internal/di"
)

func TestListPlatforms(t *testing.T) {
	hs := &di.HandlerSwitcher{}
	server := httptest.NewServer(hs)
	defer server.Close()

	options := &di.Options{
		Test: true,
		Url:  server.URL,
	}

	universe, starterFunc, err := di.BigBang(options)
	require.NoError(t, err)

	hs.UpdateHandler(universe)
	require.NoError(t, starterFunc())

	token, err := sdk.RegisterUserAccount(
		options.ClientFactory.Client().UserSvcAPI,
		"someuser",
		"pw123",
		"Some name",
	)
	require.NoError(t, err)
	userClient := options.ClientFactory.Client(sdk.WithToken(token))

	adminClient, _, err := test.AdminClient(options.ClientFactory)
	require.NoError(t, err)

	t.Run("get models", func(t *testing.T) {
		_, _, err := userClient.ModelSvcAPI.ListPlatforms(context.Background()).
			Execute()
		require.Error(t, err)

		getModelsRsp, _, err := adminClient.ModelSvcAPI.ListPlatforms(context.Background()).
			Execute()

		require.NoError(t, err)

		require.Equal(t, 2, len(getModelsRsp.Platforms))
		require.Equal(t, true, len(getModelsRsp.Platforms[0].Types) > 0)
	})

}
