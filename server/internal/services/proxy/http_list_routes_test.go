package proxyservice_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/test"

	openapi "github.com/1backend/1backend/clients/go"
)

func TestListRoutes(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		TokenExpiration: 1 * time.Second,
		Test:            true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)
	adminClient, _, err := test.AdminClient(clientFactory, sdk.DefaultTestAppHost)
	require.NoError(t, err)

	clients, _, err := test.MakeClients(
		clientFactory,
		sdk.DefaultTestAppHost,
		1,
	)
	require.NoError(t, err)

	client := clients[0]

	t.Run("save routes", func(t *testing.T) {
		req := openapi.ProxySvcSaveRoutesRequest{
			Routes: []openapi.ProxySvcRouteInput{
				{
					Id:     "test.localhost",
					Target: openapi.PtrString("some-backend-url"),
				},
				{
					Id:     "test.localhost2",
					Target: openapi.PtrString("some-backend-url2"),
				},
			},
		}

		_, hrsp, err := adminClient.ProxySvcAPI.SaveRoutes(
			context.Background(),
		).Body(req).Execute()

		require.NoError(t, err, hrsp)
	})

	t.Run("unauthorized cannot list routes", func(t *testing.T) {
		_, _, err := client.ProxySvcAPI.ListRoutes(
			context.Background(),
		).Execute()

		require.Error(t, err)
	})

	t.Run("list all routes", func(t *testing.T) {
		rsp, hrsp, err := adminClient.ProxySvcAPI.ListRoutes(
			context.Background(),
		).Execute()

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

		var routeIDs []string
		var routeTargets []string
		for _, route := range rsp.Routes {
			if route.Id != "" {
				routeIDs = append(routeIDs, route.Id)
			}
			if route.Target != "" {
				routeTargets = append(routeTargets, route.Target)
			}
		}

		require.Contains(t, routeIDs, "test.localhost")
		require.Contains(t, routeIDs, "test.localhost2")
		require.Contains(t, routeTargets, "some-backend-url")
		require.Contains(t, routeTargets, "some-backend-url2")
	})

	t.Run("list routes with expired token", func(t *testing.T) {
		time.Sleep(1200 * time.Millisecond)

		rsp, hrsp, err := adminClient.ProxySvcAPI.ListRoutes(
			context.Background(),
		).Execute()

		require.NoError(t, err, hrsp)
		require.NotNil(t, rsp)
		require.Len(t, rsp.Routes, 2)
	})
}

func TestDeleteRoutes(t *testing.T) {
	t.Parallel()

	// 1. Setup Service
	server, err := test.StartService(test.Options{
		Test: true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)
	adminClient, _, err := test.AdminClient(clientFactory, sdk.DefaultTestAppHost)
	require.NoError(t, err)

	userClient, _, err := test.MakeClients(clientFactory, sdk.DefaultTestAppHost, 1)
	require.NoError(t, err)

	// 2. Seed Data: Save two routes to delete later
	initialRoutes := openapi.ProxySvcSaveRoutesRequest{
		Routes: []openapi.ProxySvcRouteInput{
			{Id: "route.one", Target: openapi.PtrString("target.one")},
			{Id: "route.two", Target: openapi.PtrString("target.two")},
		},
	}
	_, hrsp, err := adminClient.ProxySvcAPI.SaveRoutes(context.Background()).Body(initialRoutes).Execute()
	require.NoError(t, err, hrsp)

	// 3. Test: Unauthorized user cannot delete
	t.Run("unauthorized user is blocked", func(t *testing.T) {
		req := openapi.ProxySvcDeleteRoutesRequest{
			Ids: []string{"route.one"},
		}
		_, hrsp, err := userClient[0].ProxySvcAPI.DeleteRoutes(context.Background()).Body(req).Execute()

		require.Error(t, err)
		require.Equal(t, 401, hrsp.StatusCode)
	})

	// 4. Test: Admin deletes one route
	t.Run("admin deletes a single route", func(t *testing.T) {
		req := openapi.ProxySvcDeleteRoutesRequest{
			Ids: []string{"route.one"},
		}
		_, hrsp, err := adminClient.ProxySvcAPI.DeleteRoutes(context.Background()).Body(req).Execute()
		require.NoError(t, err, hrsp)

		// Verify count is now 1
		list, _, _ := adminClient.ProxySvcAPI.ListRoutes(context.Background()).Execute()
		require.Len(t, list.Routes, 1)
		require.Equal(t, "route.two", list.Routes[0].Id)
	})

	// 5. Test: Admin deletes multiple/remaining routes
	t.Run("admin deletes remaining routes", func(t *testing.T) {
		req := openapi.ProxySvcDeleteRoutesRequest{
			Ids: []string{"route.two"},
		}
		_, hrsp, err := adminClient.ProxySvcAPI.DeleteRoutes(context.Background()).Body(req).Execute()
		require.NoError(t, err, hrsp)

		// Verify store is empty
		list, _, _ := adminClient.ProxySvcAPI.ListRoutes(context.Background()).Execute()
		require.Empty(t, list.Routes)
	})

	// 6. Test: Delete non-existent route (should typically be idempotent/no-op)
	t.Run("deleting non-existent route is idempotent", func(t *testing.T) {
		req := openapi.ProxySvcDeleteRoutesRequest{
			Ids: []string{"ghost.route"},
		}
		_, hrsp, err := adminClient.ProxySvcAPI.DeleteRoutes(context.Background()).Body(req).Execute()

		// Usually, deleting something already gone returns 200 OK
		require.NoError(t, err, hrsp)
		require.Equal(t, 200, hrsp.StatusCode)
	})
}
