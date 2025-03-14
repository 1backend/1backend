package userservice_test

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	openapi "github.com/openorch/openorch/clients/go"
	sdk "github.com/openorch/openorch/sdk/go"
	"github.com/openorch/openorch/server/internal/di"
)

func TestCreateRole(t *testing.T) {
	hs := &di.HandlerSwitcher{}
	server := httptest.NewServer(hs)
	defer server.Close()

	options := &di.Options{
		Test: true,
		Url:  server.URL,
	}
	universe, starterFunc, err := di.BigBang(options)
	require.NoError(t, err)

	hs.UpdateHandler(universe)

	err = starterFunc()
	require.NoError(t, err)

	token, err := sdk.RegisterUserAccount(
		options.ClientFactory.Client().UserSvcAPI,
		"someuser",
		"pw123",
		"Some name",
	)
	require.NoError(t, err)
	userClient := options.ClientFactory.Client(sdk.WithToken(token))

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
