package modelservice_test

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/boot"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/test"
	"github.com/1backend/1backend/server/internal/di"
	"github.com/1backend/1backend/server/internal/universe"
)

func TestModel(t *testing.T) {
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
	require.NoError(t, universe.StarterFunc())

	token, err := boot.RegisterUserAccount(
		options.ClientFactory.Client().UserSvcAPI,
		sdk.DefaultTestAppHost,
		"someuser",
		"pw123",
		"Some name",
	)
	require.NoError(t, err)
	userClient := options.ClientFactory.Client(client.WithToken(token.Token))

	adminClient, _, err := test.AdminClient(options.ClientFactory, sdk.DefaultTestAppHost)
	require.NoError(t, err)

	t.Run("get models", func(t *testing.T) {
		_, _, err := userClient.ModelSvcAPI.ListModels(context.Background()).
			Execute()
		require.Error(t, err)

		getModelsRsp, _, err := adminClient.ModelSvcAPI.ListModels(context.Background()).
			Execute()
		require.NoError(t, err)

		require.Equal(t, 1, len(getModelsRsp.Models[0].Assets))
	})

	t.Run("model status is not running, not ready", func(t *testing.T) {
		_, _, err := userClient.ModelSvcAPI.GetModelStatus(context.Background(), "huggingface/TheBloke/mistral-7b-instruct-v0.2.Q2_K.gguf").
			Execute()
		require.Error(t, err)

		statusRsp, _, err := adminClient.ModelSvcAPI.GetModelStatus(context.Background(), "huggingface/TheBloke/mistral-7b-instruct-v0.2.Q2_K.gguf").
			Execute()
		require.NoError(t, err)

		require.Equal(t, false, statusRsp.Status.Running)
		require.Equal(t, false, statusRsp.Status.AssetsReady)
		// will be ~ "172.17.0.1:8001"
		require.Equal(t, true, statusRsp.Status.Address != "")
	})

	t.Run("default", func(t *testing.T) {
		// getConfigRsp, _, err := userClient.ConfigSvcAPI.GetConfig(context.Background()).
		// 	Execute()
		// require.NoError(t, err)
		// require.Equal(
		// 	t,
		// 	modelservice.DefaultModelId,
		// 	dipper.Get(getConfigRsp.Config.Data, "modelSvc.currentModelId"),
		// )

		_, _, err = userClient.ModelSvcAPI.MakeDefault(context.Background(), "huggingface/TheBloke/mistral-7b-instruct-v0.2.Q2_K.gguf").
			Execute()
		// errors because it is not downloaded yet
		require.Error(t, err)

		// getConfigRsp, _, err := userClient.ConfigSvcAPI.GetConfig(context.Background()).
		// 	Execute()
		// require.NoError(t, err)
		//
		// require.Equal(
		// 	t,
		// 	modelservice.DefaultModelId,
		// 	dipper.Get(getConfigRsp.Config.Data, "modelSvc.currentModelId"),
		// )
	})
}
