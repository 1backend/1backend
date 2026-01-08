package userservice_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/test"

	openapi "github.com/1backend/1backend/clients/go"
)

func TestListApps(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		Test: true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)

	manyClients, _, err := test.MakeClients(clientFactory, sdk.DefaultTestAppHost, 1)
	require.NoError(t, err)

	userClient := manyClients[0]
	adminClient, _, err := test.AdminClient(clientFactory, sdk.DefaultTestAppHost)
	require.NoError(t, err)

	// unnamed + this tests app
	t.Run("only two apps initially", func(t *testing.T) {
		rsp, _, err := adminClient.UserSvcAPI.ListApps(
			context.Background(),
		).Body(openapi.UserSvcListAppsRequest{}).Execute()

		require.NoError(t, err)
		require.Len(t, rsp.Apps, 2)

		var foundUnnamed, foundServerApp bool
		for _, app := range rsp.Apps {
			if app.Host == "unnamed" {
				foundUnnamed = true
			}
			if app.Host == sdk.DefaultTestAppHost {
				foundServerApp = true
			}
		}
		require.True(t, foundUnnamed)
		require.True(t, foundServerApp)
	})

	t.Run("create an app by making a failed login request", func(t *testing.T) {
		_, _, err := adminClient.UserSvcAPI.Login(
			context.Background(),
		).Body(openapi.UserSvcLoginRequest{
			Slug:    openapi.PtrString("non-existent"),
			AppHost: "example.com",
		}).Execute()

		require.Error(t, err)
	})

	t.Run("three apps now", func(t *testing.T) {
		rsp, _, err := adminClient.UserSvcAPI.ListApps(
			context.Background(),
		).Body(openapi.UserSvcListAppsRequest{}).Execute()

		require.NoError(t, err)
		require.Len(t, rsp.Apps, 3)
	})

	t.Run("users cannot list all apps", func(t *testing.T) {
		_, _, err := userClient.UserSvcAPI.ListApps(
			context.Background(),
		).Body(openapi.UserSvcListAppsRequest{}).Execute()

		require.Error(t, err)
	})

	t.Run("but users can list a single app by id or host", func(t *testing.T) {
		rsp, _, err := userClient.UserSvcAPI.ListApps(
			context.Background(),
		).Body(openapi.UserSvcListAppsRequest{
			Host: []string{"example.com"},
		}).Execute()

		require.NoError(t, err)
		require.Len(t, rsp.Apps, 1)
		require.Equal(t, "example.com", rsp.Apps[0].Host)
	})
}

func TestUpdateAppHost(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		Test: true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)

	// Setup: Create a regular user and an admin
	manyClients, _, err := test.MakeClients(clientFactory, sdk.DefaultTestAppHost, 1)
	require.NoError(t, err)
	userClient := manyClients[0]

	adminClient, _, err := test.AdminClient(clientFactory, sdk.DefaultTestAppHost)
	require.NoError(t, err)

	// 1. Identify an app to update (e.g., the 'unnamed' app created at startup)
	var targetApp openapi.UserSvcApp
	t.Run("find target app", func(t *testing.T) {
		rsp, _, err := adminClient.UserSvcAPI.ListApps(context.Background()).
			Body(openapi.UserSvcListAppsRequest{
				Host: []string{"unnamed"},
			}).Execute()
		require.NoError(t, err)
		require.Equal(t, 1, len(rsp.Apps))
		require.NotEmpty(t, rsp.Apps)
		targetApp = rsp.Apps[0]
	})

	const newHostName = "new-awesome-host.io"

	// 2. Test Authorization: Regular users should not be able to change hosts
	t.Run("non-admin cannot update host", func(t *testing.T) {
		_, _, err := userClient.UserSvcAPI.UpdateApp(context.Background()).
			Body(openapi.UserSvcUpdateAppRequest{
				Id:      targetApp.Id,
				NewHost: newHostName,
			}).Execute()

		require.Error(t, err)
		// Verify the host hasn't changed
		checkRsp, _, _ := adminClient.UserSvcAPI.ListApps(context.Background()).
			Body(openapi.UserSvcListAppsRequest{Ids: []string{targetApp.Id}}).Execute()
		require.Equal(t, "unnamed", checkRsp.Apps[0].Host)
	})

	// 3. Test Success: Admin updates the host
	t.Run("admin can update host", func(t *testing.T) {
		_, _, err := adminClient.UserSvcAPI.UpdateApp(context.Background()).
			Body(openapi.UserSvcUpdateAppRequest{
				Id:      targetApp.Id,
				NewHost: newHostName,
			}).Execute()

		require.NoError(t, err)
	})

	// 4. Persistence Check: Verify the change is reflected in subsequent lists
	t.Run("verify update persistence", func(t *testing.T) {
		rsp, _, err := adminClient.UserSvcAPI.ListApps(context.Background()).
			Body(openapi.UserSvcListAppsRequest{
				Host: []string{newHostName},
			}).Execute()

		require.NoError(t, err)
		require.Len(t, rsp.Apps, 1)
		require.Equal(t, targetApp.Id, rsp.Apps[0].Id)
		require.Equal(t, newHostName, rsp.Apps[0].Host)
	})
}
