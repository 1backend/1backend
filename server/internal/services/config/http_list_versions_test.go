/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package configservice_test

import (
	"context"
	"testing"
	"time"

	openapi "github.com/1backend/1backend/clients/go"
	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/test"
	"github.com/stretchr/testify/require"
)

func TestConfigVersioning(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{Test: true})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)
	clients, _, err := test.MakeClients(clientFactory, sdk.DefaultTestAppHost, 1)
	require.NoError(t, err)
	c := clients[0]

	ctx := context.Background()

	t.Run("create multiple versions", func(t *testing.T) {
		for i := 0; i < 3; i++ {
			_, _, err := c.ConfigSvcAPI.SaveConfig(ctx).
				Body(openapi.ConfigSvcSaveConfigRequest{
					Data: map[string]any{
						"field": i,
					},
				}).Execute()
			require.NoError(t, err)
			time.Sleep(10 * time.Millisecond) // ensure distinct timestamps
		}
	})

	t.Run("list all versions", func(t *testing.T) {
		rsp, _, err := c.ConfigSvcAPI.ListConfigVersions(ctx).
			Body(openapi.ConfigSvcListVersionsRequest{
				AppHost: sdk.DefaultTestAppHost,
				Ids:     []string{"testUserSlug0"},
				Limit:   openapi.PtrInt32(10),
			}).Execute()
		require.NoError(t, err)
		require.NotNil(t, rsp.Versions)
		require.GreaterOrEqual(t, len(rsp.Versions), 3)
	})

	t.Run("filter by branch", func(t *testing.T) {
		rsp, _, err := c.ConfigSvcAPI.ListConfigVersions(ctx).
			Body(openapi.ConfigSvcListVersionsRequest{
				AppHost: sdk.DefaultTestAppHost,
				Branch:  openapi.PtrString("main"),
			}).Execute()
		require.NoError(t, err)
		require.NotNil(t, rsp.Versions)
		for _, v := range rsp.Versions {
			require.Equal(t, "main", v.Branch)
		}
	})

	t.Run("after cursor pagination", func(t *testing.T) {
		firstPage, _, err := c.ConfigSvcAPI.ListConfigVersions(ctx).
			Body(openapi.ConfigSvcListVersionsRequest{
				AppHost: sdk.DefaultTestAppHost,
				Limit:   openapi.PtrInt32(1),
			}).Execute()
		require.NoError(t, err)
		require.Len(t, firstPage.Versions, 1)

		cursor := []any{
			firstPage.Versions[0].CreatedAt,
			firstPage.Versions[0].Id,
		}
		afterJSON := sdk.Marshal(cursor)

		secondPage, _, err := c.ConfigSvcAPI.ListConfigVersions(ctx).
			Body(openapi.ConfigSvcListVersionsRequest{
				AppHost:   sdk.DefaultTestAppHost,
				AfterJson: afterJSON,
				Limit:     openapi.PtrInt32(5),
			}).Execute()

		require.NoError(t, err)
		require.NotNil(t, secondPage.Versions)
		require.NotEqual(t, firstPage.Versions[0].VersionId, secondPage.Versions[0].VersionId)
	})

	t.Run("selector only returns requested fields", func(t *testing.T) {
		rsp, _, err := c.ConfigSvcAPI.ListConfigVersions(ctx).
			Body(openapi.ConfigSvcListVersionsRequest{
				AppHost:  sdk.DefaultTestAppHost,
				Selector: &map[string][]string{"testUserSlug0": {"field"}},
			}).Execute()

		require.NoError(t, err)
		require.NotNil(t, rsp.Versions)
		for _, v := range rsp.Versions {
			require.NotNil(t, v.Data["field"])
			require.Len(t, v.Data, 1, "selector should filter out other fields")
		}
	})

	t.Run("read config by versionId", func(t *testing.T) {
		// Save a new version and capture its versionId
		_, _, err := c.ConfigSvcAPI.SaveConfig(ctx).
			Body(openapi.ConfigSvcSaveConfigRequest{
				Data: map[string]any{"field": "v1"},
			}).Execute()
		require.NoError(t, err)

		// List latest version to get versionId
		listRsp, _, err := c.ConfigSvcAPI.ListConfigVersions(ctx).
			Body(openapi.ConfigSvcListVersionsRequest{
				AppHost: sdk.DefaultTestAppHost,
				Limit:   openapi.PtrInt32(1),
			}).Execute()
		require.NoError(t, err)
		require.NotEmpty(t, listRsp.Versions)

		versionId := listRsp.Versions[0].VersionId
		require.NotEmpty(t, versionId)

		// Read config by versionId
		readRsp, _, err := c.ConfigSvcAPI.ListConfigVersions(ctx).
			Body(openapi.ConfigSvcListVersionsRequest{
				AppHost:    sdk.DefaultTestAppHost,
				VersionIds: []string{versionId},
			}).Execute()

		require.NoError(t, err)
		require.NotNil(t, readRsp)
		require.Len(t, readRsp.Versions, 1)
		require.Equal(t, versionId, readRsp.Versions[0].VersionId)
		require.Equal(t, "v1", readRsp.Versions[0].Data["field"])
	})

	t.Run("version entries separated by branch", func(t *testing.T) {
		// Create versions in two branches
		_, _, err := c.ConfigSvcAPI.SaveConfig(ctx).
			Body(openapi.ConfigSvcSaveConfigRequest{
				Data:   map[string]any{"b": "mainVer"},
				Branch: openapi.PtrString("main"),
			}).Execute()
		require.NoError(t, err)

		_, _, err = c.ConfigSvcAPI.SaveConfig(ctx).
			Body(openapi.ConfigSvcSaveConfigRequest{
				Data:   map[string]any{"b": "previewVer"},
				Branch: openapi.PtrString("preview"),
			}).Execute()
		require.NoError(t, err)

		// Default: no branch specified → should return only main
		rsp, _, err := c.ConfigSvcAPI.ListConfigVersions(ctx).
			Body(openapi.ConfigSvcListVersionsRequest{
				AppHost: sdk.DefaultTestAppHost,
			}).Execute()
		require.NoError(t, err)

		branches := map[string]bool{}
		for _, v := range rsp.Versions {
			branches[v.Branch] = true
		}

		require.Contains(t, branches, "main", "default listing must include main branch")
		require.Len(t, branches, 1, "only main branch should be returned when no branch filter is provided")

		// Explicit preview request → should return only preview versions
		previewRsp, _, err := c.ConfigSvcAPI.ListConfigVersions(ctx).
			Body(openapi.ConfigSvcListVersionsRequest{
				AppHost: sdk.DefaultTestAppHost,
				Branch:  openapi.PtrString("preview"),
			}).Execute()
		require.NoError(t, err)

		for _, v := range previewRsp.Versions {
			require.Equal(t, "preview", v.Branch)
		}
	})
	t.Run("read multiple versions by versionIds", func(t *testing.T) {
		// Save several versions
		for i := 0; i < 3; i++ {
			_, _, err := c.ConfigSvcAPI.SaveConfig(ctx).
				Body(openapi.ConfigSvcSaveConfigRequest{
					Data: map[string]any{"multiVerField": i},
				}).Execute()
			require.NoError(t, err)
			time.Sleep(5 * time.Millisecond)
		}

		// List to get the latest 3 versions
		listRsp, _, err := c.ConfigSvcAPI.ListConfigVersions(ctx).
			Body(openapi.ConfigSvcListVersionsRequest{
				AppHost: sdk.DefaultTestAppHost,
				Ids:     []string{"testUserSlug0"},
				Limit:   openapi.PtrInt32(3),
			}).Execute()
		require.NoError(t, err)
		require.GreaterOrEqual(t, len(listRsp.Versions), 3)

		// Collect versionIds for batch read
		versionIds := []string{
			listRsp.Versions[0].VersionId,
			listRsp.Versions[1].VersionId,
		}

		readRsp, _, err := c.ConfigSvcAPI.ListConfigVersions(ctx).
			Body(openapi.ConfigSvcListVersionsRequest{
				AppHost:    sdk.DefaultTestAppHost,
				VersionIds: versionIds, // new field in request struct
			}).Execute()

		require.NoError(t, err)
		require.NotNil(t, readRsp)
		require.Len(t, readRsp.Versions, 2)

		// Check both requested versionIds are returned and in data
		found := map[string]bool{}
		for _, v := range readRsp.Versions {
			found[v.VersionId] = true
			require.Contains(t, versionIds, v.VersionId)
			require.Contains(t, v.Data, "multiVerField")
		}
		require.True(t, found[versionIds[0]])
		require.True(t, found[versionIds[1]])
	})

}
