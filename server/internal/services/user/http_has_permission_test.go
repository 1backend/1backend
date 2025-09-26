/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package userservice_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	openapi "github.com/1backend/1backend/clients/go"
	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/boot"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/test"
	user_svc "github.com/1backend/1backend/server/internal/services/user/types"
)

func TestUnauthorizedShouldNotReturnError(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		Test: true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)

	token, err := boot.RegisterUserAccount(
		clientFactory.Client().UserSvcAPI,
		sdk.DefaultTestAppHost,
		"someuser",
		"pw123",
		"Some name",
	)
	require.NoError(t, err)
	userClient := clientFactory.Client(client.WithToken(token.Token))

	ctx := context.Background()

	rsp, _, err := userClient.UserSvcAPI.
		HasPermission(ctx, "user.view").
		Execute()
	require.NoError(t, err)
	require.False(t, rsp.Authorized)
	require.NotNil(t, rsp.AppId)
	require.Equal(t, sdk.DefaultTestAppHost, rsp.App.Host)
	require.NotEmpty(t, rsp.User)

	t.Run("not logged in user should not return error", func(t *testing.T) {
		_, hrsp, err := clientFactory.Client().UserSvcAPI.
			HasPermission(context.Background(), "user.view").
			Execute()

		require.Error(t, err)
		require.Equal(t, 422, hrsp.StatusCode)
	})
}

func TestPermitsBySlug(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		Test: true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)

	token, err := boot.RegisterUserAccount(
		clientFactory.Client().UserSvcAPI,
		sdk.DefaultTestAppHost,
		"someuser",
		"pw123",
		"Some name",
	)
	require.NoError(t, err)
	userClient := clientFactory.Client(client.WithToken(token.Token))

	ctx := context.Background()

	_, _, err = userClient.UserSvcAPI.ListUsers(ctx).Execute()
	require.Error(t, err)

	adminClient, _, err := test.AdminClient(clientFactory, sdk.DefaultTestAppHost)
	require.NoError(t, err)

	_, _, err = adminClient.UserSvcAPI.SavePermits(ctx).Body(openapi.UserSvcSavePermitsRequest{
		Permits: []openapi.UserSvcPermitInput{
			{
				Slugs:      []string{"someuser"},
				Permission: user_svc.PermissionUserView,
			},
		},
	}).Execute()
	require.NoError(t, err)

	rsp, _, err := userClient.UserSvcAPI.ListUsers(ctx).Execute()
	require.NoError(t, err)
	require.NotEmpty(t, len(rsp.Users))
}

func TestPermitsByRoleId(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		Test: true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)

	token, err := boot.RegisterUserAccount(
		clientFactory.Client().UserSvcAPI,
		sdk.DefaultTestAppHost,
		"someuser",
		"pw123",
		"Some name",
	)
	require.NoError(t, err)
	userClient := clientFactory.Client(client.WithToken(token.Token))

	ctx := context.Background()

	_, _, err = userClient.UserSvcAPI.ListUsers(ctx).Execute()
	require.Error(t, err)

	adminClient, _, err := test.AdminClient(clientFactory, sdk.DefaultTestAppHost)
	require.NoError(t, err)

	_, _, err = adminClient.UserSvcAPI.SavePermits(ctx).Body(openapi.UserSvcSavePermitsRequest{
		Permits: []openapi.UserSvcPermitInput{
			{
				Roles:      []string{"user-svc:user"},
				Permission: user_svc.PermissionUserView,
			},
		},
	}).Execute()
	require.NoError(t, err)

	rsp, _, err := userClient.UserSvcAPI.ListUsers(ctx).Execute()
	require.NoError(t, err)
	require.NotEmpty(t, len(rsp.Users))
}

func TestAutoRefresh(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		// JWT expiration is only second granular
		TokenExpiration: time.Second,
		Test:            true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)

	token, err := boot.RegisterUserAccount(
		clientFactory.Client().UserSvcAPI,
		sdk.DefaultTestAppHost,
		"someuser",
		"pw123",
		"Some name",
	)
	require.NoError(t, err)
	userClient := clientFactory.Client(client.WithToken(token.Token))

	ctx := context.Background()

	rsp, _, err := userClient.UserSvcAPI.ReadSelf(ctx).Execute()
	require.NoError(t, err)
	require.Equal(t, "someuser", rsp.User.Slug)

	time.Sleep(2 * time.Second)

	rsp, _, err = userClient.UserSvcAPI.ReadSelf(ctx).Execute()
	require.NoError(t, err)
	require.Equal(t, "someuser", rsp.User.Slug)
}

func TestAutoRefreshOff(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		// JWT expiration is only second granular
		TokenExpiration:     time.Second,
		TokenAutoRefreshOff: true,
		Test:                true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)

	token, err := boot.RegisterUserAccount(
		clientFactory.Client().UserSvcAPI,
		sdk.DefaultTestAppHost,
		"someuser",
		"pw123",
		"Some name",
	)
	require.NoError(t, err)
	userClient := clientFactory.Client(client.WithToken(token.Token))

	ctx := context.Background()

	rsp, _, err := userClient.UserSvcAPI.ReadSelf(ctx).Execute()
	require.NoError(t, err)
	require.Equal(t, "someuser", rsp.User.Slug)

	time.Sleep(2 * time.Second)

	_, _, err = userClient.UserSvcAPI.ReadSelf(ctx).Execute()
	require.Error(t, err)
}
