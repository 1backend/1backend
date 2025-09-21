package registryservice_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	openapi "github.com/1backend/1backend/clients/go"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/test"
)

func TestListInstances(t *testing.T) {
	t.Parallel()

	mockBackend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer mockBackend.Close()

	server, err := test.StartService(test.Options{
		TokenExpiration: 1 * time.Second,
		Test:            true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	client := client.NewApiClientFactory(server.Url)

	manyClients, _, err := test.MakeClients(client, sdk.DefaultTestAppHost, 1)
	require.NoError(t, err)

	client1 := manyClients[0]

	t.Run("register instance", func(t *testing.T) {
		_, _, err = client1.RegistrySvcAPI.RegisterInstance(
			context.Background(),
		).Body(
			openapi.RegistrySvcRegisterInstanceRequest{
				Url: mockBackend.URL,
			}).Execute()
		require.NoError(t, err, err)
	})

	t.Run("list instances", func(t *testing.T) {
		resp, _, err := client1.RegistrySvcAPI.ListInstances(
			context.Background(),
		).Execute()
		require.NoError(t, err)

		require.Len(t, resp.Instances, 1)
		require.Equal(t, mockBackend.URL, resp.Instances[0].Url)
	})

	t.Run("list instances works after the token expires", func(t *testing.T) {
		time.Sleep(1 * time.Second)

		resp, _, err := client1.RegistrySvcAPI.ListInstances(
			context.Background(),
		).Execute()
		require.NoError(t, err)

		require.Len(t, resp.Instances, 1)
		require.Equal(t, mockBackend.URL, resp.Instances[0].Url)
	})
}
