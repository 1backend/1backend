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

func TestModel(t *testing.T) {
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
		// 	dipper.Get(getConfigRsp.Config.Data, "model-svc.currentModelId"),
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
		// 	dipper.Get(getConfigRsp.Config.Data, "model-svc.currentModelId"),
		// )
	})
}
