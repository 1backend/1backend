package userservice_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

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

	manyClients, _, err := test.MakeClients(
		client.NewApiClientFactory(server.Url), 2)
	require.NoError(t, err)

	userClient := manyClients[0]
	otherClient := manyClients[1]

	orgId := ""
	slug := "some-org-name"

	t.Run("save organization", func(t *testing.T) {
		orgName := "Some org name"

		req := openapi.UserSvcSaveOrganizationRequest{
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
		require.Equal(t, httpResp.StatusCode, 200)

		orgId = resp.Organization.Id
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
}
