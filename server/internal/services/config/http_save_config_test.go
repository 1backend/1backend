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
	sdk "github.com/1backend/1backend/sdk/go"
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

	manyClients, _, err := test.MakeClients(clientFactory, sdk.DefaultTestAppHost, 2)
	require.NoError(t, err)
	client1 := manyClients[0]
	client2 := manyClients[1]

	adminClient, _, err := test.AdminClient(clientFactory)
	require.NoError(t, err)

	ctx := context.Background()

	t.Run("save config client 1", func(t *testing.T) {
		_, _, err := client1.ConfigSvcAPI.SaveConfig(ctx).
			Body(openapi.ConfigSvcSaveConfigRequest{
				Data: map[string]any{
					"field1": "value1",
					"field2": "value2",
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
		require.Equal(t, "value1", rsp.Configs["testUserSlug0"].Data["field1"], rsp)
		require.Equal(t, "value2", rsp.Configs["testUserSlug0"].Data["field2"], rsp)
	})

	t.Run("save config client 2", func(t *testing.T) {
		_, _, err := client2.ConfigSvcAPI.SaveConfig(ctx).
			Body(openapi.ConfigSvcSaveConfigRequest{
				Data: map[string]any{
					"field1": "newValue1",
					"field2": "newValue2",
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
		require.Equal(t, "newValue1", rsp.Configs["testUserSlug1"].Data["field1"], rsp)
		require.Equal(t, "newValue2", rsp.Configs["testUserSlug1"].Data["field2"], rsp)
	})

	t.Run("list configs by id", func(t *testing.T) {
		rsp, _, err := client2.ConfigSvcAPI.ListConfigs(ctx).
			Body(openapi.ConfigSvcListConfigsRequest{
				Ids: []string{"testUserSlug1"},
			}).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, rsp.Configs)
		require.NotNil(t, rsp.Configs["testUserSlug1"])
		require.Equal(t, "newValue1", rsp.Configs["testUserSlug1"].Data["field1"], rsp)
		require.Equal(t, "newValue2", rsp.Configs["testUserSlug1"].Data["field2"], rsp)
	})

	t.Run("test deep merge functionality", func(t *testing.T) {
		t.Run("save nested", func(t *testing.T) {
			_, _, err := client1.ConfigSvcAPI.SaveConfig(ctx).
				Body(openapi.ConfigSvcSaveConfigRequest{
					Data: map[string]any{
						"id3": map[string]any{
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
						"id3": map[string]any{
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
					Ids: []string{"testUserSlug0"},
				}).
				Execute()

			require.NoError(t, err)

			config := rsp.Configs["testUserSlug0"]
			require.NotNil(t, config)

			require.NotNil(t, config.Data["field1"])
			require.NotNil(t, config.Data["field2"])

			id3, ok := config.Data["id3"].(map[string]any)
			require.True(t, ok, "id3 should be a nested map")

			require.Equal(t, "n1-updated", id3["nested1"], rsp)
			require.Equal(t, "n2", id3["nested2"], rsp)
			require.Equal(t, "n3", id3["nested3"], rsp)
		})
	})

	t.Run("admins can save any slug config", func(t *testing.T) {
		// Admin saves are taken at face value and the slug
		// will not be used at top level.

		_, _, err := adminClient.ConfigSvcAPI.SaveConfig(ctx).
			Body(openapi.ConfigSvcSaveConfigRequest{
				Id: openapi.PtrString("otherSvc"),
				Data: map[string]any{
					"field1": "adminValue1",
					"field2": "adminValue2",
					"field3": map[string]string{
						"secondLevel3": "adminValue3",
					},
				},
			}).
			Execute()

		require.NoError(t, err)

		rsp, _, err := adminClient.ConfigSvcAPI.ListConfigs(ctx).
			Body(openapi.ConfigSvcListConfigsRequest{
				Ids: []string{"otherSvc", "anotherSvc"},
			}).
			Execute()
		require.NoError(t, err)
		require.NotNil(t, rsp.Configs)

		require.NotNil(t, rsp.Configs["otherSvc"])
		require.Equal(t, "otherSvc", rsp.Configs["otherSvc"].Id, rsp)
		require.Equal(t, "adminValue1", rsp.Configs["otherSvc"].Data["field1"], rsp)
		require.Equal(t, "adminValue2", rsp.Configs["otherSvc"].Data["field2"], rsp)

		_, _, err = adminClient.ConfigSvcAPI.SaveConfig(ctx).
			Body(openapi.ConfigSvcSaveConfigRequest{
				AppHost: openapi.PtrString("otherApp"),
				Id:      openapi.PtrString("anotherSvc"),
				Data: map[string]any{
					"field1": "adminValue3",
					"field2": "adminValue4",
				},
			}).
			Execute()

		require.NoError(t, err)

		rsp, _, err = adminClient.ConfigSvcAPI.ListConfigs(ctx).
			Body(openapi.ConfigSvcListConfigsRequest{
				AppHost: openapi.PtrString("otherApp"),
				Ids:     []string{"anotherSvc"},
			}).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, rsp.Configs)

		require.NotNil(t, rsp.Configs["anotherSvc"])
		require.Equal(t, "adminValue3", rsp.Configs["anotherSvc"].Data["field1"], rsp)
		require.Equal(t, "adminValue4", rsp.Configs["anotherSvc"].Data["field2"], rsp)
	})

	t.Run("users cannot specify other app", func(t *testing.T) {
		_, _, err := client1.ConfigSvcAPI.SaveConfig(ctx).
			Body(openapi.ConfigSvcSaveConfigRequest{
				AppHost: openapi.PtrString("otherApp"),
				Data: map[string]any{
					"otherSvc": map[string]any{
						"field1": "userValue1",
						"field2": "userValue2",
					},
				},
			}).
			Execute()

		require.Error(t, err, "users should not be able to save configs for other apps")
	})

	t.Run("users cannot specify others id", func(t *testing.T) {
		_, _, err := client1.ConfigSvcAPI.SaveConfig(ctx).
			Body(openapi.ConfigSvcSaveConfigRequest{
				Id: openapi.PtrString("otherSvc"),
				Data: map[string]any{
					"otherSvc": map[string]any{
						"field1": "userValue1",
						"field2": "userValue2",
					},
				},
			}).
			Execute()

		require.Error(t, err, "users should not be able to save configs for other apps")
	})

	t.Run("query dotpath", func(t *testing.T) {
		rsp, _, err := client1.ConfigSvcAPI.ListConfigs(ctx).
			Body(openapi.ConfigSvcListConfigsRequest{
				Selector: &map[string][]string{
					"otherSvc": {"field1", "field3.secondLevel3"},
				},
			}).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, rsp.Configs)
		require.Len(t, rsp.Configs, 1)
		require.NotNil(t, rsp.Configs["otherSvc"])
		require.Equal(t, "adminValue1", rsp.Configs["otherSvc"].Data["field1"], rsp)
		require.Nil(t, rsp.Configs["otherSvc"].Data["field2"], rsp)
		require.NotNil(t, rsp.Configs["otherSvc"].Data["field3"].(map[string]interface{}))
		require.Equal(t, "adminValue3", rsp.Configs["otherSvc"].Data["field3"].(map[string]interface{})["secondLevel3"], rsp)
	})
}
