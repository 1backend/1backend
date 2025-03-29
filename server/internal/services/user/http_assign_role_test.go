package userservice_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/test"

	openapi "github.com/1backend/1backend/clients/go"
)

func TestAssignRoleToUser(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		Test: true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := sdk.NewApiClientFactory(server.Url)

	manyClients, tokens, err := test.MakeClients(clientFactory, 3)
	require.NoError(t, err)

	userClient := manyClients[0]

	t.Run("user creates role", func(t *testing.T) {
		_, _, err := userClient.UserSvcAPI.CreateRole(context.Background()).
			Body(openapi.UserSvcCreateRoleRequest{
				Id: "test-user-slug-0:custom-role",
			}).Execute()

		require.NoError(t, err)
	})

	// This is needed for dynamic roles.
	t.Run("can add nonexistent role to user", func(t *testing.T) {
		_, _, err := userClient.UserSvcAPI.AssignRole(
			context.Background(),
			tokens[1].UserId,
			"test-user-slug-0:custom-role-nonexistent",
		).Execute()

		require.NoError(t, err)
	})

	t.Run("assign role", func(t *testing.T) {
		_, _, err := userClient.UserSvcAPI.AssignRole(
			context.Background(),
			tokens[1].UserId,
			"test-user-slug-0:custom-role",
		).Execute()

		require.NoError(t, err)
	})

	orgId := ""

	t.Run("create organization", func(t *testing.T) {
		rsp, _, err := userClient.UserSvcAPI.CreateOrganization(
			context.Background(),
		).Body(openapi.UserSvcCreateOrganizationRequest{
			Slug: "test-org",
			Name: "Test Org",
		}).Execute()
		require.NoError(t, err)
		require.NotEmpty(t, rsp.Organization.Id)

		orgId = rsp.Organization.Id
		require.NotEmpty(t, orgId)

		require.NotEmpty(t, rsp.Token.Token)

		// Here we refresh the token to include the new org role
		userClient = clientFactory.Client(sdk.WithToken(rsp.Token.Token))
	})

	t.Run("nonexistent org role assignment", func(t *testing.T) {
		_, _, err := userClient.UserSvcAPI.AssignRole(
			context.Background(),
			tokens[1].UserId,
			"user-svc:org:{some-nonexistent-org-id}:user",
		).Execute()

		require.Error(t, err)
	})

	t.Run("org role assignment", func(t *testing.T) {
		_, _, err := userClient.UserSvcAPI.AssignRole(
			context.Background(),
			tokens[1].UserId,
			fmt.Sprintf("user-svc:org:{%v}:user", orgId),
		).Execute()

		require.NoError(t, err)
	})

	secondClient, _, err := test.LoggedInClient(
		clientFactory,
		"test-user-slug-1",
		"testUserPassword1",
	)
	require.NoError(t, err)

	t.Run("second user cannot give itself admin rights", func(t *testing.T) {
		_, _, err := secondClient.UserSvcAPI.AssignRole(
			context.Background(),
			tokens[1].UserId,
			fmt.Sprintf("user-svc:org:{%v}:user", orgId),
		).Execute()

		require.Error(t, err)
	})

	t.Run("second user cannot give a third user admin or user rights", func(t *testing.T) {
		_, _, err := secondClient.UserSvcAPI.AssignRole(
			context.Background(),
			tokens[2].UserId,
			fmt.Sprintf("user-svc:org:{%v}:user", orgId),
		).Execute()

		require.Error(t, err)

		_, _, err = secondClient.UserSvcAPI.AssignRole(
			context.Background(),
			tokens[2].UserId,
			fmt.Sprintf("user-svc:org:{%v}:admin", orgId),
		).Execute()

		require.Error(t, err)
	})

	// After making the second user admin, it can give the third user and admin rights

	_, _, err = userClient.UserSvcAPI.AssignRole(
		context.Background(),
		tokens[1].UserId,
		fmt.Sprintf("user-svc:org:{%v}:admin", orgId),
	).Execute()

	require.NoError(t, err)

	secondClient, _, err = test.LoggedInClient(
		clientFactory,
		"test-user-slug-1",
		"testUserPassword1",
	)
	require.NoError(t, err)

	t.Run("second user now can give third user user rights", func(t *testing.T) {
		_, _, err = secondClient.UserSvcAPI.AssignRole(
			context.Background(),
			tokens[2].UserId,
			fmt.Sprintf("user-svc:org:{%v}:user", orgId),
		).Execute()

		require.NoError(t, err)
	})

	t.Run("second user now can give third user user rights", func(t *testing.T) {
		_, _, err = secondClient.UserSvcAPI.AssignRole(
			context.Background(),
			tokens[2].UserId,
			fmt.Sprintf("user-svc:org:{%v}:admin", orgId),
		).Execute()

		require.NoError(t, err)
	})
}
