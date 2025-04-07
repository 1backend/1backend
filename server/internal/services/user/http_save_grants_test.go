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
	"github.com/1backend/1backend/sdk/go/boot"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/test"
)

func TestSaveGrants(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		Test: true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)

	token, err := boot.RegisterUserAccount(
		clientFactory.Client().UserSvcAPI,
		"firstuser",
		"pw123",
		"Some name",
	)
	require.NoError(t, err)
	userClient := clientFactory.Client(client.WithToken(token.Token))

	token, err = boot.RegisterUserAccount(
		clientFactory.Client().UserSvcAPI,
		"seconduser",
		"pw123",
		"Other name",
	)
	require.NoError(t, err)
	userClient2 := clientFactory.Client(client.WithToken(token.Token))

	ctx := context.Background()

	t.Run("user 1 cannot assign role it does not own", func(t *testing.T) {
		_, _, err = userClient.UserSvcAPI.SaveGrants(ctx).Body(openapi.UserSvcSaveGrantsRequest{
			Grants: []openapi.UserSvcGrant{
				{
					Roles:      []string{"user-svc:user"},
					Permission: "seconduser:myperm",
				},
			},
		}).Execute()

		require.Error(t, err)
	})

	t.Run("user 1 can assign role it owns", func(t *testing.T) {
		_, _, err = userClient.UserSvcAPI.SaveGrants(ctx).Body(openapi.UserSvcSaveGrantsRequest{
			Grants: []openapi.UserSvcGrant{
				{
					Roles:      []string{"user-svc:user"},
					Permission: "firstuser:myperm",
				},
			},
		}).Execute()

		require.NoError(t, err)
	})

	t.Run("saving the same grant again will not cause a duplicate", func(t *testing.T) {
		rsp, _, err := userClient.UserSvcAPI.ListGrants(ctx).Body(
			openapi.UserSvcListGrantsRequest{},
		).Execute()
		require.NoError(t, err)

		count := len(rsp.Grants)

		_, _, err = userClient.UserSvcAPI.SaveGrants(ctx).Body(openapi.UserSvcSaveGrantsRequest{
			Grants: []openapi.UserSvcGrant{
				{
					Roles:      []string{"user-svc:user"},
					Permission: "firstuser:myperm",
				},
			},
		}).Execute()
		require.NoError(t, err)

		rsp, _, err = userClient.UserSvcAPI.ListGrants(ctx).Body(
			openapi.UserSvcListGrantsRequest{},
		).Execute()
		require.NoError(t, err)

		require.Equal(t, count, len(rsp.Grants))
	})

	t.Run("user 2 cannot assign role it does not own", func(t *testing.T) {
		_, _, err = userClient2.UserSvcAPI.SaveGrants(ctx).Body(openapi.UserSvcSaveGrantsRequest{
			Grants: []openapi.UserSvcGrant{
				{
					Roles:      []string{"user-svc:user"},
					Permission: "firstuser:myperm",
				},
			},
		}).Execute()

		require.Error(t, err)
	})

	t.Run("user 1 can assign role it owns", func(t *testing.T) {
		_, _, err = userClient.UserSvcAPI.SaveGrants(ctx).Body(openapi.UserSvcSaveGrantsRequest{
			Grants: []openapi.UserSvcGrant{
				{
					Roles:      []string{"user-svc:user"},
					Permission: "firstuser:myperm",
				},
			},
		}).Execute()

		require.NoError(t, err)
	})
}
