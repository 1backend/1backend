package userservice_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/test"

	openapi "github.com/1backend/1backend/clients/go"
)

func TestListOrganizations(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		Test: true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)

	manyClients, _, err := test.MakeClients(clientFactory, 1)
	require.NoError(t, err)

	userClient := manyClients[0]

	orgId := ""

	t.Run("create organization", func(t *testing.T) {
		rsp, _, err := userClient.UserSvcAPI.SaveOrganization(
			context.Background(),
		).Body(openapi.UserSvcSaveOrganizationRequest{
			Slug: "test-org",
			Name: openapi.PtrString("Test Org"),
		}).Execute()
		require.NoError(t, err)
		require.NotEmpty(t, rsp.Organization.Id)

		orgId = rsp.Organization.Id
		require.NotEmpty(t, orgId)

		require.NotEmpty(t, rsp.Token.Token)

		// Here we refresh the token to include the new org role
		userClient = clientFactory.Client(client.WithToken(rsp.Token.Token))
	})

	t.Run("list all organizations", func(t *testing.T) {
		rsp, _, err := userClient.UserSvcAPI.ListOrganizations(
			context.Background(),
		).Body(openapi.UserSvcListOrganizationsRequest{}).Execute()

		require.NoError(t, err)
		require.NotEmpty(t, rsp.Organizations)
	})

	t.Run("list existing by id", func(t *testing.T) {
		rsp, _, err := userClient.UserSvcAPI.ListOrganizations(
			context.Background(),
		).Body(openapi.UserSvcListOrganizationsRequest{
			Ids: []string{orgId},
		}).Execute()

		require.NoError(t, err)
		require.NotEmpty(t, rsp.Organizations)
	})

	t.Run("list non-existent by id", func(t *testing.T) {
		rsp, _, err := userClient.UserSvcAPI.ListOrganizations(
			context.Background(),
		).Body(openapi.UserSvcListOrganizationsRequest{
			Ids: []string{"some random id"},
		}).Execute()

		require.NoError(t, err)
		require.Empty(t, rsp.Organizations)
	})
}
