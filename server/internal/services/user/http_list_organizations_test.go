package userservice_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/1backend/1backend/sdk/go"
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

	manyClients, _, err := test.MakeClients(clientFactory, sdk.DefaultTestAppHost, 2)
	require.NoError(t, err)

	userClient := manyClients[0]
	otherUserClient := manyClients[1]
	adminClient, _, err := test.AdminClient(clientFactory, sdk.DefaultTestAppHost)
	require.NoError(t, err)

	orgId := ""
	otherOrgId := ""

	t.Run("create organization", func(t *testing.T) {
		rsp, _, err := userClient.UserSvcAPI.SaveOrganization(
			context.Background(),
		).Body(openapi.UserSvcSaveOrganizationRequest{
			Activate: openapi.PtrBool(true),
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

	t.Run("create other user organization", func(t *testing.T) {
		rsp, _, err := otherUserClient.UserSvcAPI.SaveOrganization(
			context.Background(),
		).Body(openapi.UserSvcSaveOrganizationRequest{
			Activate: openapi.PtrBool(true),
			Slug:     "other-test-org",
			Name:     openapi.PtrString("Other Test Org"),
		}).Execute()

		require.NoError(t, err)
		require.NotEmpty(t, rsp.Organization.Id)
		otherOrgId = rsp.Organization.Id
	})

	t.Run("user lists only own organizations", func(t *testing.T) {
		rsp, _, err := userClient.UserSvcAPI.ListOrganizations(
			context.Background(),
		).Body(openapi.UserSvcListOrganizationsRequest{}).Execute()

		require.NoError(t, err)
		require.Len(t, rsp.Organizations, 1)
		require.Equal(t, orgId, rsp.Organizations[0].Id)
	})

	t.Run("admin without all flag lists only own organizations", func(t *testing.T) {
		rsp, _, err := adminClient.UserSvcAPI.ListOrganizations(
			context.Background(),
		).Body(openapi.UserSvcListOrganizationsRequest{}).Execute()

		require.NoError(t, err)
		require.Empty(t, rsp.Organizations)
	})

	t.Run("admin with all flag lists all organizations", func(t *testing.T) {
		rsp, _, err := adminClient.UserSvcAPI.ListOrganizations(
			context.Background(),
		).Body(openapi.UserSvcListOrganizationsRequest{
			All: openapi.PtrBool(true),
		}).Execute()

		require.NoError(t, err)
		require.NotEmpty(t, rsp.Organizations)

		organizationIds := []string{}
		for _, organization := range rsp.Organizations {
			organizationIds = append(organizationIds, organization.Id)
		}

		require.Contains(t, organizationIds, orgId)
		require.Contains(t, organizationIds, otherOrgId)
	})

	t.Run("user cannot use all flag to list all organizations", func(t *testing.T) {
		_, httpResp, err := userClient.UserSvcAPI.ListOrganizations(
			context.Background(),
		).Body(openapi.UserSvcListOrganizationsRequest{
			All: openapi.PtrBool(true),
		}).Execute()

		require.Error(t, err)
		require.NotNil(t, httpResp)
		require.Equal(t, 401, httpResp.StatusCode)
	})

	t.Run("list existing by id", func(t *testing.T) {
		rsp, _, err := adminClient.UserSvcAPI.ListOrganizations(
			context.Background(),
		).Body(openapi.UserSvcListOrganizationsRequest{
			All: openapi.PtrBool(true),
			Ids: []string{orgId},
		}).Execute()

		require.NoError(t, err)
		require.NotEmpty(t, rsp.Organizations)
	})

	t.Run("list non-existent by id", func(t *testing.T) {
		rsp, _, err := adminClient.UserSvcAPI.ListOrganizations(
			context.Background(),
		).Body(openapi.UserSvcListOrganizationsRequest{
			All: openapi.PtrBool(true),
			Ids: []string{"some random id"},
		}).Execute()

		require.NoError(t, err)
		require.Empty(t, rsp.Organizations)
	})
}
