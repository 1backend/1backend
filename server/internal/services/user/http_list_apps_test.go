package userservice_test

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

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

	manyClients, _, err := test.MakeClients(clientFactory, 1)
	require.NoError(t, err)

	userClient := manyClients[0]
	adminClient, _, err := test.AdminClient(clientFactory)
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
			if app.Host == strings.Replace(server.Url, "http://", "", 1) {
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
			AppHost: openapi.PtrString("example.com"),
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
