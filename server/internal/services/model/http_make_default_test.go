package modelservice_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/test"
)

const dummyUrl = "https://raw.githubusercontent.com/github/gitignore/main/README.md"

func TestMakeDefault(t *testing.T) {
	server, err := test.StartService(test.Options{
		Test: true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)

	adminClient, _, err := test.AdminClient(clientFactory)
	require.NoError(t, err)

	t.Run("download model asset", func(t *testing.T) {
		_, _, err := adminClient.FileSvcAPI.
			DownloadFile(context.Background()).
			Body(openapi.FileSvcDownloadFileRequest{
				Url: dummyUrl,
			}).
			Execute()
		require.NoError(t, err)
	})

	t.Run("wait for file download to be ready", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			dl, _, err := adminClient.FileSvcAPI.
				GetDownload(context.Background(), dummyUrl).
				Execute()
			require.NoError(t, err)

			if dl.Exists && *dl.Download.Status == "completed" {
				return
			}

			time.Sleep(100 * time.Millisecond)
		}

		t.Fatal("Model download did not complete in time")
	})

	t.Run("make default", func(t *testing.T) {
		_, _, err := adminClient.ModelSvcAPI.
			MakeDefault(context.Background(), "dummy").
			Execute()
		require.NoError(t, err)
	})

	t.Run("default is seen in config", func(t *testing.T) {
		rsp, _, err := adminClient.ConfigSvcAPI.
			ListConfigs(context.Background()).
			Body(openapi.ConfigSvcListConfigsRequest{
				Slugs: []string{"modelSvc"},
			}).
			Execute()
		require.NoError(t, err)
		require.Equal(t, "dummy", rsp.Configs["modelSvc"].Data["currentModelId"], rsp)
	})
}
