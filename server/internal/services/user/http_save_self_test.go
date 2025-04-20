/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package userservice_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/test"
)

func TestSaveSelf(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		Test: true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)

	manyClients, _, err := test.MakeClients(clientFactory, 1)
	require.NoError(t, err)

	client1 := manyClients[0]

	ctx := context.Background()

	t.Run("user updates their name", func(t *testing.T) {
		_, _, err := client1.UserSvcAPI.SaveSelf(ctx).Body(
			openapi.UserSvcSaveSelfRequest{
				Name: openapi.PtrString("New Name"),
			}).Execute()

		require.NoError(t, err)
	})

	t.Run("name has changed", func(t *testing.T) {
		selfRsp, _, err := client1.UserSvcAPI.ReadSelf(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, selfRsp.User)
		require.NotNil(t, selfRsp.User.Name)
		require.Equal(t, "New Name", *selfRsp.User.Name)
	})

	t.Run("user updates their meta", func(t *testing.T) {
		_, _, err := client1.UserSvcAPI.SaveSelf(ctx).Body(
			openapi.UserSvcSaveSelfRequest{
				Labels: &map[string]string{
					"key": "value",
				},
			}).Execute()

		require.NoError(t, err)
	})

	t.Run("name has not changed, meta is saved", func(t *testing.T) {
		selfRsp, _, err := client1.UserSvcAPI.ReadSelf(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, selfRsp.User)
		require.NotNil(t, selfRsp.User.Name)
		require.Equal(t, "New Name", *selfRsp.User.Name)

		require.NotNil(t, selfRsp.User.Labels)
		require.Equal(t, "value", (*selfRsp.User.Labels)["key"])
	})

	t.Run("update meta with a new field", func(t *testing.T) {
		_, _, err := client1.UserSvcAPI.SaveSelf(ctx).Body(
			openapi.UserSvcSaveSelfRequest{
				Labels: &map[string]string{
					"key2": "value2",
				},
			}).Execute()

		require.NoError(t, err)
	})

	t.Run("check if new meta fields are saved", func(t *testing.T) {
		selfRsp, _, err := client1.UserSvcAPI.ReadSelf(ctx).Execute()

		require.NoError(t, err)
		require.NotNil(t, selfRsp.User)
		require.NotNil(t, selfRsp.User.Name)
		require.Equal(t, "New Name", *selfRsp.User.Name)

		require.NotNil(t, selfRsp.User.Labels)
		require.Equal(t, "value", (*selfRsp.User.Labels)["key"])
		require.Equal(t, "value2", (*selfRsp.User.Labels)["key2"])
	})

	t.Run("make sure login still works after update", func(t *testing.T) {
		_, _, err := client1.UserSvcAPI.Login(ctx).Body(
			openapi.UserSvcLoginRequest{
				Slug:     openapi.PtrString("test-user-slug-0"),
				Password: openapi.PtrString("testUserPassword0"),
			}).Execute()

		require.NoError(t, err)
	})
}
