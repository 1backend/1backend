package userservice_test

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	openapi "github.com/openorch/openorch/clients/go"
	sdk "github.com/openorch/openorch/sdk/go"
	"github.com/openorch/openorch/sdk/go/test"
	"github.com/openorch/openorch/server/internal/di"
	user_svc "github.com/openorch/openorch/server/internal/services/user/types"
)

func TestGrants(t *testing.T) {
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

	adminClient, _, err := test.AdminClient(options.ClientFactory)
	require.NoError(t, err)

	_, _, err = adminClient.UserSvcAPI.SaveGrants(ctx).Body(openapi.UserSvcSaveGrantsRequest{
		Grants: []openapi.UserSvcGrant{
			{
				Slugs:        []string{"someuser"},
				PermissionId: openapi.PtrString(user_svc.PermissionRoleView.Id),
			},
		},
	}).Execute()
	require.NoError(t, err)

	rsp, _, err := userClient.UserSvcAPI.GetRoles(ctx).Execute()
	require.NoError(t, err)
	require.Equal(t, 2, len(rsp.Roles))
}
