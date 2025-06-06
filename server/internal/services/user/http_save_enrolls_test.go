package userservice_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/test"

	openapi "github.com/1backend/1backend/clients/go"
)

func TestEnrollForUnregistered(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		Test: true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	manyClients, _, err := test.MakeClients(
		client.NewApiClientFactory(server.Url), 1)
	require.NoError(t, err)

	userClient := manyClients[0]

	t.Run("cannot create enroll for unowned role", func(t *testing.T) {
		_, rsp, err := userClient.UserSvcAPI.SaveEnrolls(context.Background()).
			Body(openapi.UserSvcSaveEnrollsRequest{
				Enrolls: []openapi.UserSvcEnrollInput{
					{
						ContactId: openapi.PtrString("test-user@email.com"),
						Role:      "some-other-role",
					},
				},
			}).Execute()

		require.Error(t, err, rsp)
	})

	t.Run("user creates enroll for new user", func(t *testing.T) {
		_, rsp, err := userClient.UserSvcAPI.SaveEnrolls(context.Background()).
			Body(openapi.UserSvcSaveEnrollsRequest{
				Enrolls: []openapi.UserSvcEnrollInput{
					{
						ContactId: openapi.PtrString("test-user@email.com"),
						Role:      "test-user-slug-0:custom-role",
					},
				},
			}).Execute()

		require.NoError(t, err, rsp)
	})

	t.Run("list enroll without params fail", func(t *testing.T) {
		_, _, err := userClient.UserSvcAPI.ListEnrolls(context.Background()).
			Body(openapi.UserSvcListEnrollsRequest{}).Execute()

		require.Error(t, err)
	})

	t.Run("list enroll", func(t *testing.T) {
		rsp, _, err := userClient.UserSvcAPI.ListEnrolls(context.Background()).
			Body(openapi.UserSvcListEnrollsRequest{
				Role: openapi.PtrString("test-user-slug-0:custom-role"),
			}).Execute()

		require.NoError(t, err)
		require.Len(t, rsp.Enrolls, 1)
	})

	t.Run("new user should have role", func(t *testing.T) {
		rsp, _, err := userClient.UserSvcAPI.Register(context.Background()).
			Body(openapi.UserSvcRegisterRequest{
				Slug: "test-user-slug-1",
				Contact: &openapi.UserSvcContactInput{
					Id: "test-user@email.com",
				},
				Password: openapi.PtrString("yo"),
			}).Execute()

		require.NoError(t, err)

		pkRsp, _, err := userClient.UserSvcAPI.
			GetPublicKey(context.Background()).
			Execute()

		require.NoError(t, err)

		claim, err := auth.AuthorizerImpl{}.ParseJWT(
			pkRsp.PublicKey,
			rsp.Token.Token,
		)
		require.NoError(t, err)
		require.NotNil(t, claim)
		require.Equal(t, 2, len(claim.Roles), claim.Roles)
		require.Contains(
			t,
			claim.Roles,
			"test-user-slug-0:custom-role",
			claim.Roles,
		)
	})

}

func TestEnrollForRegisteredUser(t *testing.T) {
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

	t.Run("register user", func(t *testing.T) {
		_, _, err := userClient.UserSvcAPI.Register(context.Background()).
			Body(openapi.UserSvcRegisterRequest{
				Slug: "test-user-slug-1",
				Contact: &openapi.UserSvcContactInput{
					Id: "test-user@email.com",
				},
				Password: openapi.PtrString("yo"),
			}).Execute()

		require.NoError(t, err)
	})

	t.Run("enroll already registered user - role should apply immediately", func(t *testing.T) {
		_, _, err := userClient.UserSvcAPI.SaveEnrolls(context.Background()).
			Body(openapi.UserSvcSaveEnrollsRequest{
				Enrolls: []openapi.UserSvcEnrollInput{
					{
						ContactId: openapi.PtrString("test-user@email.com"),
						Role:      "test-user-slug-0:custom-role",
					},
				},
			}).Execute()

		require.NoError(t, err)

		loginReq := openapi.UserSvcLoginRequest{
			Slug:     openapi.PtrString("test-user-slug-1"),
			Password: openapi.PtrString("yo"),
		}
		loginRsp, _, err := userClient.UserSvcAPI.Login(context.Background()).
			Body(loginReq).
			Execute()
		require.NoError(t, err)

		publicKeyRsp, _, err := clientFactory.Client().
			UserSvcAPI.GetPublicKey(context.Background()).
			Execute()
		require.NoError(t, err)

		claim, err := auth.AuthorizerImpl{}.ParseJWT(
			publicKeyRsp.PublicKey,
			loginRsp.Token.Token,
		)
		require.NoError(t, err)
		require.NotNil(t, claim)
		require.Equal(t, 2, len(claim.Roles), claim.Roles)
		require.Contains(
			t,
			claim.Roles,
			"test-user-slug-0:custom-role",
			claim.Roles,
		)
	})
}

func TestListEnrollAuthorization(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		Test: true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)
	manyClients, tokens, err := test.MakeClients(clientFactory, 2)
	require.NoError(t, err)

	userClient := manyClients[0]
	secondUserClient := manyClients[1]

	t.Run("user adds role to second user so both can enroll", func(t *testing.T) {
		_, _, err := userClient.UserSvcAPI.SaveEnrolls(context.Background()).
			Body(openapi.UserSvcSaveEnrollsRequest{
				Enrolls: []openapi.UserSvcEnrollInput{
					{
						UserId: openapi.PtrString(tokens[1].UserId),
						Role:   "test-user-slug-0:custom-role",
					},
				},
			}).Execute()

		require.NoError(t, err)
	})

	t.Run("first user enrolls someone", func(t *testing.T) {
		_, _, err := userClient.UserSvcAPI.SaveEnrolls(context.Background()).
			Body(openapi.UserSvcSaveEnrollsRequest{
				Enrolls: []openapi.UserSvcEnrollInput{
					{
						ContactId: openapi.PtrString("test-user@email.com"),
						Role:      "test-user-slug-0:custom-role",
					},
				},
			}).Execute()

		require.NoError(t, err)

		rsp, _, err := userClient.UserSvcAPI.ListEnrolls(context.Background()).
			Body(openapi.UserSvcListEnrollsRequest{
				Role: openapi.PtrString("test-user-slug-0:custom-role"),
			}).Execute()

		require.NoError(t, err)
		require.Len(t, rsp.Enrolls, 2, rsp)
	})

	secondUserClient, _, err = test.LoggedInClient(clientFactory, "test-user-slug-1", "testUserPassword1")
	require.NoError(t, err)

	t.Run("second user cannot enroll someone as it has the role but does not own it", func(t *testing.T) {
		_, _, err := secondUserClient.UserSvcAPI.SaveEnrolls(context.Background()).
			Body(openapi.UserSvcSaveEnrollsRequest{
				Enrolls: []openapi.UserSvcEnrollInput{
					{
						ContactId: openapi.PtrString("test-user@email.com"),
						Role:      "test-user-slug-0:custom-role",
					},
				},
			}).Execute()

		require.Error(t, err)

		rsp, _, err := secondUserClient.UserSvcAPI.ListEnrolls(context.Background()).
			Body(openapi.UserSvcListEnrollsRequest{
				Role: openapi.PtrString("test-user-slug-0:custom-role"),
			}).Execute()

		require.Error(t, err, rsp)
	})
}

func TestSaveEnrollsOldAssignTests(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		Test: true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)

	manyClients, tokens, err := test.MakeClients(clientFactory, 3)
	require.NoError(t, err)

	userClient := manyClients[0]

	// This is needed for dynamic roles.
	t.Run("can add nonexistent role to user", func(t *testing.T) {
		_, _, err := userClient.UserSvcAPI.SaveEnrolls(context.Background()).
			Body(openapi.UserSvcSaveEnrollsRequest{
				Enrolls: []openapi.UserSvcEnrollInput{
					{
						UserId: openapi.PtrString(tokens[1].UserId),
						Role:   "test-user-slug-0:custom-role-nonexistent",
					},
				},
			}).Execute()

		require.NoError(t, err)
	})

	t.Run("assign role", func(t *testing.T) {
		_, _, err := userClient.UserSvcAPI.SaveEnrolls(context.Background()).
			Body(openapi.UserSvcSaveEnrollsRequest{
				Enrolls: []openapi.UserSvcEnrollInput{
					{
						UserId: openapi.PtrString(tokens[1].UserId),
						Role:   "test-user-slug-0:custom-role",
					},
				},
			}).Execute()

		require.NoError(t, err)
	})

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

	t.Run("nonexistent org role assignment", func(t *testing.T) {
		_, _, err := userClient.UserSvcAPI.SaveEnrolls(context.Background()).
			Body(openapi.UserSvcSaveEnrollsRequest{
				Enrolls: []openapi.UserSvcEnrollInput{
					{
						UserId: openapi.PtrString(tokens[1].UserId),
						Role:   "user-svc:org:{some-nonexistent-org-id}:user",
					},
				},
			}).Execute()

		require.Error(t, err)
	})

	t.Run("org role assignment", func(t *testing.T) {
		_, _, err := userClient.UserSvcAPI.SaveEnrolls(context.Background()).
			Body(openapi.UserSvcSaveEnrollsRequest{
				Enrolls: []openapi.UserSvcEnrollInput{
					{
						UserId: openapi.PtrString(tokens[1].UserId),
						Role:   fmt.Sprintf("user-svc:org:{%v}:user", orgId),
					},
				},
			}).Execute()

		require.NoError(t, err)
	})

	secondClient, _, err := test.LoggedInClient(
		clientFactory,
		"test-user-slug-1",
		"testUserPassword1",
	)
	require.NoError(t, err)

	t.Run("second user cannot give itself admin rights", func(t *testing.T) {
		_, _, err := secondClient.UserSvcAPI.SaveEnrolls(context.Background()).
			Body(openapi.UserSvcSaveEnrollsRequest{
				Enrolls: []openapi.UserSvcEnrollInput{
					{
						UserId: openapi.PtrString(tokens[1].UserId),
						Role:   fmt.Sprintf("user-svc:org:{%v}:user", orgId),
					},
				},
			}).Execute()

		require.Error(t, err)
	})

	t.Run("second user cannot give a third user admin or user rights", func(t *testing.T) {
		_, _, err := secondClient.UserSvcAPI.SaveEnrolls(context.Background()).
			Body(openapi.UserSvcSaveEnrollsRequest{
				Enrolls: []openapi.UserSvcEnrollInput{
					{
						UserId: openapi.PtrString(tokens[1].UserId),
						Role:   fmt.Sprintf("user-svc:org:{%v}:user", orgId),
					},
				},
			}).Execute()

		require.Error(t, err)

		_, _, err = secondClient.UserSvcAPI.SaveEnrolls(context.Background()).
			Body(openapi.UserSvcSaveEnrollsRequest{
				Enrolls: []openapi.UserSvcEnrollInput{
					{
						UserId: openapi.PtrString(tokens[1].UserId),
						Role:   fmt.Sprintf("user-svc:org:{%v}:admin", orgId),
					},
				},
			}).Execute()

		require.Error(t, err)
	})

	// After making the second user admin, it can give the third user user and admin rights

	_, _, err = userClient.UserSvcAPI.SaveEnrolls(context.Background()).
		Body(openapi.UserSvcSaveEnrollsRequest{
			Enrolls: []openapi.UserSvcEnrollInput{
				{
					UserId: openapi.PtrString(tokens[1].UserId),
					Role:   fmt.Sprintf("user-svc:org:{%v}:admin", orgId),
				},
			},
		}).Execute()

	require.NoError(t, err)

	secondClient, _, err = test.LoggedInClient(
		clientFactory,
		"test-user-slug-1",
		"testUserPassword1",
	)
	require.NoError(t, err)

	t.Run("second user now can give third user user rights", func(t *testing.T) {
		_, _, err = secondClient.UserSvcAPI.SaveEnrolls(context.Background()).
			Body(openapi.UserSvcSaveEnrollsRequest{
				Enrolls: []openapi.UserSvcEnrollInput{
					{
						UserId: openapi.PtrString(tokens[2].UserId),
						Role:   fmt.Sprintf("user-svc:org:{%v}:user", orgId),
					},
				},
			}).Execute()

		require.NoError(t, err, orgId)
	})

	t.Run("second user now can give third user admin rights", func(t *testing.T) {
		_, _, err = secondClient.UserSvcAPI.SaveEnrolls(context.Background()).
			Body(openapi.UserSvcSaveEnrollsRequest{
				Enrolls: []openapi.UserSvcEnrollInput{
					{
						UserId: openapi.PtrString(tokens[2].UserId),
						Role:   fmt.Sprintf("user-svc:org:{%v}:admin", orgId),
					},
				},
			}).Execute()

		require.NoError(t, err)
	})
}
