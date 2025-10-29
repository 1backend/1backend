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

	adminClient, _, err := test.AdminClient(clientFactory, sdk.DefaultTestAppHost)
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
			Body(openapi.ConfigSvcListConfigsRequest{
				AppHost: sdk.DefaultTestAppHost,
			}).
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
			Body(openapi.ConfigSvcListConfigsRequest{
				AppHost: sdk.DefaultTestAppHost,
			}).
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
				AppHost: sdk.DefaultTestAppHost,
				Ids:     []string{"testUserSlug1"},
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
					AppHost: sdk.DefaultTestAppHost,
					Ids:     []string{"testUserSlug0"},
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
				AppHost: sdk.DefaultTestAppHost,
				Ids:     []string{"otherSvc", "anotherSvc"},
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
				AppHost: "otherApp",
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
				AppHost: sdk.DefaultTestAppHost,
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

	t.Run("branch-based save and read", func(t *testing.T) {
		// Client 1 saves different data into main and preview branches
		_, _, err := client1.ConfigSvcAPI.SaveConfig(ctx).
			Body(openapi.ConfigSvcSaveConfigRequest{
				Data: map[string]any{
					"env": "production",
				},
				Branch: openapi.PtrString("main"),
			}).Execute()
		require.NoError(t, err)

		_, _, err = client1.ConfigSvcAPI.SaveConfig(ctx).
			Body(openapi.ConfigSvcSaveConfigRequest{
				Data: map[string]any{
					"env": "staging",
				},
				Branch: openapi.PtrString("preview"),
			}).Execute()
		require.NoError(t, err)

		// Read from main branch
		mainRsp, _, err := client1.ConfigSvcAPI.ListConfigs(ctx).
			Body(openapi.ConfigSvcListConfigsRequest{
				AppHost: sdk.DefaultTestAppHost,
				Branch:  openapi.PtrString("main"),
				Ids:     []string{"testUserSlug0"},
			}).Execute()
		require.NoError(t, err)
		require.NotNil(t, mainRsp.Configs["testUserSlug0"])
		require.Equal(t, "production", mainRsp.Configs["testUserSlug0"].Data["env"], mainRsp)

		// Read from preview branch
		previewRsp, _, err := client1.ConfigSvcAPI.ListConfigs(ctx).
			Body(openapi.ConfigSvcListConfigsRequest{
				AppHost: sdk.DefaultTestAppHost,
				Branch:  openapi.PtrString("preview"),
				Ids:     []string{"testUserSlug0"},
			}).Execute()
		require.NoError(t, err)
		require.NotNil(t, previewRsp.Configs["testUserSlug0"])
		require.Equal(t, "staging", previewRsp.Configs["testUserSlug0"].Data["env"], previewRsp)

		// Ensure branches are isolated
		require.NotEqual(t,
			mainRsp.Configs["testUserSlug0"].Data["env"],
			previewRsp.Configs["testUserSlug0"].Data["env"],
			"main and preview branches must hold distinct data",
		)
	})

	t.Run("branch isolation after multiple saves", func(t *testing.T) {
		// Save main and preview branches multiple times
		for i := 0; i < 2; i++ {
			_, _, err := client1.ConfigSvcAPI.SaveConfig(ctx).
				Body(openapi.ConfigSvcSaveConfigRequest{
					Data:   map[string]any{"counter": i},
					Branch: openapi.PtrString("main"),
				}).Execute()
			require.NoError(t, err)
		}

		_, _, err := client1.ConfigSvcAPI.SaveConfig(ctx).
			Body(openapi.ConfigSvcSaveConfigRequest{
				Data:   map[string]any{"counter": "preview-only"},
				Branch: openapi.PtrString("preview"),
			}).Execute()
		require.NoError(t, err)

		// Confirm branch data divergence
		mainRsp, _, _ := client1.ConfigSvcAPI.ListConfigs(ctx).
			Body(openapi.ConfigSvcListConfigsRequest{
				AppHost: sdk.DefaultTestAppHost,
				Branch:  openapi.PtrString("main"),
				Ids:     []string{"testUserSlug0"},
			}).Execute()

		previewRsp, _, _ := client1.ConfigSvcAPI.ListConfigs(ctx).
			Body(openapi.ConfigSvcListConfigsRequest{
				AppHost: sdk.DefaultTestAppHost,
				Branch:  openapi.PtrString("preview"),
				Ids:     []string{"testUserSlug0"},
			}).Execute()

		require.NotEqual(t,
			mainRsp.Configs["testUserSlug0"].Data["counter"],
			previewRsp.Configs["testUserSlug0"].Data["counter"],
		)
	})

	t.Run("default branch is main", func(t *testing.T) {
		_, _, err := client1.ConfigSvcAPI.SaveConfig(ctx).
			Body(openapi.ConfigSvcSaveConfigRequest{
				Data: map[string]any{"auto": "defaultBranch"},
			}).Execute()
		require.NoError(t, err)

		rsp, _, err := client1.ConfigSvcAPI.ListConfigs(ctx).
			Body(openapi.ConfigSvcListConfigsRequest{
				AppHost: sdk.DefaultTestAppHost,
				Branch:  openapi.PtrString("main"),
				Ids:     []string{"testUserSlug0"},
			}).Execute()
		require.NoError(t, err)
		require.Equal(t, "defaultBranch", rsp.Configs["testUserSlug0"].Data["auto"])
	})

	t.Run("pagination returns empty when no more data", func(t *testing.T) {
		rsp, _, _ := client1.ConfigSvcAPI.ListConfigVersions(ctx).
			Body(openapi.ConfigSvcListVersionsRequest{
				AppHost: sdk.DefaultTestAppHost,
				Limit:   openapi.PtrInt32(1000),
			}).Execute()
		require.NoError(t, err)
		cursor := []any{rsp.Versions[len(rsp.Versions)-1].CreatedAt, rsp.Versions[len(rsp.Versions)-1].Id}
		afterJSON := sdk.Marshal(cursor)

		empty, _, _ := client1.ConfigSvcAPI.ListConfigVersions(ctx).
			Body(openapi.ConfigSvcListVersionsRequest{
				AppHost:   sdk.DefaultTestAppHost,
				AfterJson: afterJSON,
			}).Execute()
		require.Empty(t, empty.Versions)
	})

	t.Run("version history preserves previous values", func(t *testing.T) {
		// Save first version
		_, _, err := client1.ConfigSvcAPI.SaveConfig(ctx).
			Body(openapi.ConfigSvcSaveConfigRequest{
				Data: map[string]any{"featureFlag": "off"},
			}).Execute()
		require.NoError(t, err)

		time.Sleep(5 * time.Millisecond)

		// Save updated version
		_, _, err = client1.ConfigSvcAPI.SaveConfig(ctx).
			Body(openapi.ConfigSvcSaveConfigRequest{
				Data: map[string]any{"featureFlag": "on"},
			}).Execute()
		require.NoError(t, err)

		// List versions, newest first
		rsp, _, err := client1.ConfigSvcAPI.ListConfigVersions(ctx).
			Body(openapi.ConfigSvcListVersionsRequest{
				AppHost: sdk.DefaultTestAppHost,
				Ids:     []string{"testUserSlug0"},
				Limit:   openapi.PtrInt32(2),
			}).Execute()
		require.NoError(t, err)
		require.GreaterOrEqual(t, len(rsp.Versions), 2)

		latest := rsp.Versions[0]
		previous := rsp.Versions[1]

		require.Equal(t, "on", latest.Data["featureFlag"], "latest version must contain the new value")
		require.Equal(t, "off", previous.Data["featureFlag"], "previous version must preserve the old value")

		t1, err := parseStringTime(latest.CreatedAt)
		require.NoError(t, err)
		t2, err := parseStringTime(previous.CreatedAt)
		require.NoError(t, err)
		require.True(t, t1.After(t2), "latest version must have newer timestamp")
	})

	t.Run("scope all returns separate entries per branch", func(t *testing.T) {
		// Save different values under main and preview
		_, _, err := client1.ConfigSvcAPI.SaveConfig(ctx).
			Body(openapi.ConfigSvcSaveConfigRequest{
				Data:   map[string]any{"scopeFlag": "mainValue"},
				Branch: openapi.PtrString("main"),
			}).Execute()
		require.NoError(t, err)

		_, _, err = client1.ConfigSvcAPI.SaveConfig(ctx).
			Body(openapi.ConfigSvcSaveConfigRequest{
				Data:   map[string]any{"scopeFlag": "previewValue"},
				Branch: openapi.PtrString("preview"),
			}).Execute()
		require.NoError(t, err)

		scope := openapi.ListConfigsScopeAll

		// Query all branches
		rsp, _, err := client1.ConfigSvcAPI.ListConfigs(ctx).
			Body(openapi.ConfigSvcListConfigsRequest{
				AppHost: sdk.DefaultTestAppHost,
				Ids:     []string{"testUserSlug0"},
				Scope:   &scope,
			}).Execute()
		require.NoError(t, err)

		// Expect two distinct entries: <id>::main and <id>::preview
		mainKey := "testUserSlug0:main"
		previewKey := "testUserSlug0:preview"

		require.NotNil(t, rsp.Configs[mainKey])
		require.NotNil(t, rsp.Configs[previewKey])

		require.Equal(t, "mainValue", rsp.Configs[mainKey].Data["scopeFlag"])
		require.Equal(t, "previewValue", rsp.Configs[previewKey].Data["scopeFlag"])
	})
}

func parseStringTime(s string) (time.Time, error) {
	if s == "" {
		return time.Time{}, nil
	}
	// RFC3339 is Go’s default JSON timestamp format.
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		// attempt fallback for non-RFC3339 timestamps
		t, err = time.Parse("2006-01-02T15:04:05Z07:00", s)
	}
	return t, err
}
