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

func TestSaveOrganization(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		Test: true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)
	manyClients, _, err := test.MakeClients(
		clientFactory, sdk.DefaultTestAppHost, 2)
	require.NoError(t, err)

	userClient := manyClients[0]
	otherClient := manyClients[1]
	adminClient, adminToken, err := test.AdminClient(
		clientFactory,
		sdk.DefaultTestAppHost,
	)
	require.NoError(t, err)

	orgId := ""
	otherOrgId := ""
	slug := "some-org-name"

	t.Run("save organization", func(t *testing.T) {
		orgName := "Some org name"

		req := openapi.UserSvcSaveOrganizationRequest{
			Activate: openapi.PtrBool(true),
			Name:     openapi.PtrString(orgName),
			Slug:     slug,
		}

		resp, httpResp, err := userClient.UserSvcAPI.
			SaveOrganization(
				context.Background(),
			).
			Body(req).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, httpResp.StatusCode, 200)

		orgId = resp.Organization.Id
		require.NotNil(t, resp.Token)
		require.NotEmpty(t, resp.Token.Token)
		userClient = clientFactory.
			Client(client.WithToken(resp.Token.Token))
	})

	t.Run("other user cannot update organization", func(t *testing.T) {
		orgName := "Some other org name"
		req := openapi.UserSvcSaveOrganizationRequest{
			Id:   openapi.PtrString(orgId),
			Name: openapi.PtrString(orgName),
			Slug: slug,
		}

		_, httpResp, err := otherClient.UserSvcAPI.
			SaveOrganization(

				context.Background(),
			).
			Body(req).
			Execute()

		require.Error(t, err)
		require.NotNil(t, httpResp)
		require.Equal(t, 401, httpResp.StatusCode)
	})

	t.Run("organization admin can update own organization", func(t *testing.T) {
		orgName := "Updated by organization admin"
		req := openapi.UserSvcSaveOrganizationRequest{
			Id:   openapi.PtrString(orgId),
			Name: openapi.PtrString(orgName),
			Slug: slug,
		}

		resp, httpResp, err := userClient.UserSvcAPI.
			SaveOrganization(
				context.Background(),
			).
			Body(req).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 200, httpResp.StatusCode)
		require.Equal(t, orgName, resp.Organization.Name)
	})

	t.Run("organization admin can update non-active organization", func(t *testing.T) {
		resp, _, err := userClient.UserSvcAPI.
			SaveOrganization(
				context.Background(),
			).
			Body(openapi.UserSvcSaveOrganizationRequest{
				Activate: openapi.PtrBool(true),
				Name:     openapi.PtrString("Another Org"),
				Slug:     "another-org",
			}).
			Execute()

		require.NoError(t, err)
		require.NotEmpty(t, resp.Organization.Id)
		require.NotEmpty(t, resp.Token.Token)
		otherOrgId = resp.Organization.Id
		userClient = clientFactory.
			Client(client.WithToken(resp.Token.Token))

		updateResp, httpResp, err := userClient.UserSvcAPI.
			SaveOrganization(
				context.Background(),
			).
			Body(openapi.UserSvcSaveOrganizationRequest{
				Id:   openapi.PtrString(orgId),
				Name: openapi.PtrString("Updated while another org is active"),
				Slug: slug,
			}).
			Execute()

		require.NoError(t, err)
		require.Equal(t, 200, httpResp.StatusCode)
		require.Equal(t, orgId, updateResp.Organization.Id)
		require.NotEqual(t, orgId, otherOrgId)
	})

	t.Run("create organization with activate false returns no token and keeps active org unchanged", func(t *testing.T) {
		resp, httpResp, err := userClient.UserSvcAPI.
			SaveOrganization(
				context.Background(),
			).
			Body(openapi.UserSvcSaveOrganizationRequest{
				Activate: openapi.PtrBool(false),
				Name:     openapi.PtrString("Passive Org"),
				Slug:     "passive-org",
			}).
			Execute()

		require.NoError(t, err)
		require.Equal(t, 200, httpResp.StatusCode)
		require.Nil(t, resp.Token)

		selfRsp, _, err := userClient.UserSvcAPI.
			ReadSelf(context.Background()).
			Body(openapi.UserSvcReadSelfRequest{}).
			Execute()
		require.NoError(t, err)
		require.Equal(t, otherOrgId, selfRsp.GetActiveOrganizationId())
	})

	t.Run("organization admin can update and activate organization", func(t *testing.T) {
		updateResp, httpResp, err := userClient.UserSvcAPI.
			SaveOrganization(
				context.Background(),
			).
			Body(openapi.UserSvcSaveOrganizationRequest{
				Id:       openapi.PtrString(orgId),
				Name:     openapi.PtrString("Activated via update"),
				Slug:     slug,
				Activate: openapi.PtrBool(true),
			}).
			Execute()

		require.NoError(t, err)
		require.Equal(t, 200, httpResp.StatusCode)
		require.NotNil(t, updateResp.Token)
		require.NotEmpty(t, updateResp.Token.Token)

		userClient = client.NewApiClientFactory(server.Url).
			Client(client.WithToken(updateResp.Token.Token))

		selfRsp, _, err := userClient.UserSvcAPI.
			ReadSelf(context.Background()).
			Body(openapi.UserSvcReadSelfRequest{}).
			Execute()
		require.NoError(t, err)
		require.Equal(t, orgId, selfRsp.GetActiveOrganizationId())
	})

	t.Run("organization admin update with activate false does not deactivate organization", func(t *testing.T) {
		updateResp, httpResp, err := userClient.UserSvcAPI.
			SaveOrganization(
				context.Background(),
			).
			Body(openapi.UserSvcSaveOrganizationRequest{
				Id:       openapi.PtrString(orgId),
				Name:     openapi.PtrString("Deactivated via update"),
				Slug:     slug,
				Activate: openapi.PtrBool(false),
			}).
			Execute()

		require.NoError(t, err)
		require.Equal(t, 200, httpResp.StatusCode)
		require.Nil(t, updateResp.Token)

		selfRsp, _, err := userClient.UserSvcAPI.
			ReadSelf(context.Background()).
			Body(openapi.UserSvcReadSelfRequest{}).
			Execute()
		require.NoError(t, err)
		require.Equal(t, orgId, selfRsp.GetActiveOrganizationId())
	})

	t.Run("organization admin update does not return token when activation state is unchanged", func(t *testing.T) {
		updateResp, httpResp, err := userClient.UserSvcAPI.
			SaveOrganization(
				context.Background(),
			).
			Body(openapi.UserSvcSaveOrganizationRequest{
				Id:       openapi.PtrString(orgId),
				Name:     openapi.PtrString("Still inactive"),
				Slug:     slug,
				Activate: openapi.PtrBool(false),
			}).
			Execute()

		require.NoError(t, err)
		require.Equal(t, 200, httpResp.StatusCode)
		require.Nil(t, updateResp.Token)
	})

	t.Run("admin can update organization without acquiring org membership", func(t *testing.T) {
		adminSelf, _, err := adminClient.UserSvcAPI.
			ReadSelf(context.Background()).
			Body(openapi.UserSvcReadSelfRequest{}).
			Execute()
		require.NoError(t, err)

		orgName := "Updated by admin"
		req := openapi.UserSvcSaveOrganizationRequest{
			Id:   openapi.PtrString(orgId),
			Name: openapi.PtrString(orgName),
			Slug: slug,
		}

		resp, httpResp, err := adminClient.UserSvcAPI.
			SaveOrganization(
				context.Background(),
			).
			Body(req).
			Execute()

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 200, httpResp.StatusCode)
		require.Equal(t, orgName, resp.Organization.Name)
		require.Nil(t, resp.Token)

		membershipRsp, httpResp := listMemberships(
			t,
			clientFactory.Client(client.WithToken(adminToken)),
			&openapi.UserSvcListMembershipsRequest{OrganizationId: openapi.PtrString(orgId)},
		)
		require.Equal(t, 200, httpResp.StatusCode)
		for _, membership := range membershipRsp.Memberships {
			require.NotEqual(t, adminSelf.User.Id, membership.User.Id)
		}
	})
}
