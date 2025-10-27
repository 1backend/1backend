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
					Tags: []string{"release"},
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

	t.Run("filter by tag", func(t *testing.T) {
		rsp, _, err := c.ConfigSvcAPI.ListConfigVersions(ctx).
			Body(openapi.ConfigSvcListVersionsRequest{
				AppHost: sdk.DefaultTestAppHost,
				Tags:    []string{"release"},
			}).Execute()
		require.NoError(t, err)
		require.NotNil(t, rsp.Versions)
		for _, v := range rsp.Versions {
			require.Contains(t, v.Tags, "release")
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
		require.NotEqual(t, firstPage.Versions[0].Version, secondPage.Versions[0].Version)
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
}
