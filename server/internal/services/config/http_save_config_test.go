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

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/test"
	"github.com/stretchr/testify/require"
)

func TestConfigService(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		Test: true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)

	manyClients, _, err := test.MakeClients(clientFactory, 2)
	require.NoError(t, err)
	client1 := manyClients[0]
	client2 := manyClients[1]

	ctx := context.Background()

	t.Run("save config client 1", func(t *testing.T) {
		_, _, err := client1.ConfigSvcAPI.SaveConfig(ctx).
			Body(openapi.ConfigSvcSaveConfigRequest{
				Data: map[string]any{
					"key1": "value1",
					"key2": "value2",
				},
			}).
			Execute()

		require.NoError(t, err)
	})

	t.Run("read config client 1", func(t *testing.T) {
		rsp, _, err := client1.ConfigSvcAPI.ListConfigs(ctx).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, rsp.Configs)
		require.NotNil(t, rsp.Configs["testUserSlug0"])
		require.Equal(t, "value1", rsp.Configs["testUserSlug0"].Data["key1"], rsp)
		require.Equal(t, "value2", rsp.Configs["testUserSlug0"].Data["key2"], rsp)
	})

	t.Run("save config client 2", func(t *testing.T) {
		_, _, err := client2.ConfigSvcAPI.SaveConfig(ctx).
			Body(openapi.ConfigSvcSaveConfigRequest{
				Data: map[string]any{
					"key1": "newValue1",
					"key2": "newValue2",
				},
			}).
			Execute()

		require.NoError(t, err)
	})

	t.Run("read config client 2", func(t *testing.T) {
		rsp, _, err := client2.ConfigSvcAPI.ListConfigs(ctx).
			Execute()
		require.NoError(t, err)
		require.NotNil(t, rsp.Configs)
		require.NotNil(t, rsp.Configs["testUserSlug1"])
		require.Equal(t, "newValue1", rsp.Configs["testUserSlug1"].Data["key1"], rsp)
		require.Equal(t, "newValue2", rsp.Configs["testUserSlug1"].Data["key2"], rsp)
	})

	t.Run("list configs by id", func(t *testing.T) {
		rsp, _, err := client2.ConfigSvcAPI.ListConfigs(ctx).
			Body(openapi.ConfigSvcListConfigsRequest{
				Slugs: []string{"testUserSlug1"},
			}).
			Execute()
		require.NoError(t, err)
		require.NotNil(t, rsp.Configs)
		require.NotNil(t, rsp.Configs["testUserSlug1"])
		require.Equal(t, "newValue1", rsp.Configs["testUserSlug1"].Data["key1"], rsp)
		require.Equal(t, "newValue2", rsp.Configs["testUserSlug1"].Data["key2"], rsp)
	})

	t.Run("test deep merge functionality", func(t *testing.T) {
		t.Run("save nested", func(t *testing.T) {
			_, _, err := client1.ConfigSvcAPI.SaveConfig(ctx).
				Body(openapi.ConfigSvcSaveConfigRequest{
					Data: map[string]any{
						"key3": map[string]any{
							"nested1": "n1",
							"nested2": "n2",
						},
					},
				}).
				Execute()
			require.NoError(t, err)
		})

		t.Run("updated nested", func(t *testing.T) {
			_, _, err = client1.ConfigSvcAPI.SaveConfig(ctx).
				Body(openapi.ConfigSvcSaveConfigRequest{
					Data: map[string]any{
						"key3": map[string]any{
							"nested1": "n1-updated",
							"nested3": "n3",
						},
					},
				}).
				Execute()
			require.NoError(t, err)
		})

		t.Run("verify deep merge result", func(t *testing.T) {
			rsp, _, err := client1.ConfigSvcAPI.
				ListConfigs(ctx).
				Body(openapi.ConfigSvcListConfigsRequest{
					Slugs: []string{"testUserSlug0"},
				}).
				Execute()
			require.NoError(t, err)

			config := rsp.Configs["testUserSlug0"]
			require.NotNil(t, config)

			require.NotNil(t, config.Data["key1"])
			require.NotNil(t, config.Data["key2"])

			key3, ok := config.Data["key3"].(map[string]any)
			require.True(t, ok, "key3 should be a nested map")

			require.Equal(t, "n1-updated", key3["nested1"], rsp)
			require.Equal(t, "n2", key3["nested2"], rsp)
			require.Equal(t, "n3", key3["nested3"], rsp)
		})
	})

}
