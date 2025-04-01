package modelservice_test

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/1backend/1backend/sdk/go/boot"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/test"
	"github.com/1backend/1backend/server/internal/di"
)

func TestListPlatforms(t *testing.T) {
	hs := &di.HandlerSwitcher{}
	server := httptest.NewServer(hs)
	defer server.Close()

	options := &di.Options{
		Test: true,
		Url:  server.URL,
	}

	universe, err := di.BigBang(options)
	require.NoError(t, err)

	hs.UpdateHandler(universe.Router)
	require.NoError(t, universe.StarterFunc())

	token, err := boot.RegisterUserAccount(
		options.ClientFactory.Client().UserSvcAPI,
		"someuser",
		"pw123",
		"Some name",
	)
	require.NoError(t, err)
	userClient := options.ClientFactory.Client(client.WithToken(token.Token))

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
