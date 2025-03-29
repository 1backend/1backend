package userservice_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	openapi "github.com/1backend/1backend/clients/go"
	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/test"
	user_svc "github.com/1backend/1backend/server/internal/services/user/types"
)

func TestGrants(t *testing.T) {
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

	adminClient, _, err := test.AdminClient(clientFactory)
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
