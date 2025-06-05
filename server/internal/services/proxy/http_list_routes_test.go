package proxyservice_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/test"

	openapi "github.com/1backend/1backend/clients/go"
)

func TestListRoutes(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		Test: true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)
	adminClient, _, err := test.AdminClient(clientFactory)
	require.NoError(t, err)

	t.Run("save routes", func(t *testing.T) {
		req := openapi.ProxySvcSaveRoutesRequest{
			Routes: []openapi.ProxySvcRouteInput{
				{
					Id:     openapi.PtrString("test.localhost"),
					Target: openapi.PtrString("some-backend-url"),
				},
				{
					Id:     openapi.PtrString("test.localhost2"),
					Target: openapi.PtrString("some-backend-url2"),
				},
			},
		}

		_, hrsp, err := adminClient.ProxySvcAPI.SaveRoutes(
			context.Background(),
		).Body(req).Execute()

		require.NoError(t, err, hrsp)
	})

	t.Run("list all routes", func(t *testing.T) {
		rsp, hrsp, err := adminClient.ProxySvcAPI.ListRoutes(
			context.Background(),
		).Body(openapi.ProxySvcListRoutesRequest{}).Execute()

		require.NoError(t, err, hrsp)
		require.NotNil(t, rsp)
		require.Len(t, rsp.Routes, 2)
	})

	t.Run("list route by id", func(t *testing.T) {
		rsp, hrsp, err := adminClient.ProxySvcAPI.ListRoutes(
			context.Background(),
		).Execute()

		require.NoError(t, err, hrsp)
		require.NotNil(t, rsp)
		require.Equal(t, "test.localhost", *rsp.Routes[0].Id)
		require.Equal(t, "some-backend-url", *rsp.Routes[0].Target)
	})
}
