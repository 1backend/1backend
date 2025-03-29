package userservice_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	openapi "github.com/1backend/1backend/clients/go"
	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/test"
)

func TestCreateRole(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		Test: true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := sdk.NewApiClientFactory(server.Url)

	token, err := sdk.RegisterUserAccount(
		clientFactory.Client().UserSvcAPI,
		"someuser",
		"pw123",
		"Some name",
	)
	require.NoError(t, err)
	userClient := clientFactory.Client(sdk.WithToken(token.Token))

	ctx := context.Background()

	_, _, err = userClient.UserSvcAPI.GetRoles(ctx).Execute()
	require.Error(t, err)

	t.Run("created role should fail if not prefixed by user slug", func(t *testing.T) {
		_, _, err = userClient.UserSvcAPI.CreateRole(ctx).Body(openapi.UserSvcCreateRoleRequest{
			Id: "shouldfail",
		}).Execute()
		require.Error(t, err)

		_, _, err = userClient.UserSvcAPI.CreateRole(ctx).Body(openapi.UserSvcCreateRoleRequest{
			Id:   "someuser:shouldnotfail",
			Name: openapi.PtrString("Hey"),
		}).Execute()
		require.NoError(t, err)
	})
}
