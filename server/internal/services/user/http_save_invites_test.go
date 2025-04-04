package userservice_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/test"

	openapi "github.com/1backend/1backend/clients/go"
)

func TestInviteForUnregistered(t *testing.T) {
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

	t.Run("cannot create invite for unowned role", func(t *testing.T) {
		_, rsp, err := userClient.UserSvcAPI.SaveInvites(context.Background()).
			Body(openapi.UserSvcSaveInvitesRequest{
				Invites: []openapi.UserSvcNewInvite{
					{
						ContactId: "test-user@email.com",
						RoleId:    "some-other-role",
					},
				},
			}).Execute()

		require.Error(t, err, rsp)
	})

	t.Run("user creates invite for new user", func(t *testing.T) {
		_, rsp, err := userClient.UserSvcAPI.SaveInvites(context.Background()).
			Body(openapi.UserSvcSaveInvitesRequest{
				Invites: []openapi.UserSvcNewInvite{
					{
						ContactId: "test-user@email.com",
						RoleId:    "test-user-slug-0:custom-role",
					},
				},
			}).Execute()

		require.NoError(t, err, rsp)
	})

	t.Run("list invite", func(t *testing.T) {
		rsp, _, err := userClient.UserSvcAPI.ListInvites(context.Background()).
			Body(openapi.UserSvcListInvitesRequest{}).Execute()

		require.NoError(t, err)
		require.Len(t, rsp.Invites, 1)
	})

	t.Run("new user should have role", func(t *testing.T) {
		rsp, _, err := userClient.UserSvcAPI.Register(context.Background()).
			Body(openapi.UserSvcRegisterRequest{
				Slug: "test-user-slug-1",
				Contact: &openapi.UserSvcContact{
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
		require.Equal(t, 2, len(claim.RoleIds), claim.RoleIds)
		require.Contains(
			t,
			claim.RoleIds,
			"test-user-slug-0:custom-role",
			claim.RoleIds,
		)
	})

}

func TestInviteForRegisteredUser(t *testing.T) {
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
				Contact: &openapi.UserSvcContact{
					Id: "test-user@email.com",
				},
				Password: openapi.PtrString("yo"),
			}).Execute()

		require.NoError(t, err)
	})

	t.Run("invite already registered user - role should apply immediately", func(t *testing.T) {
		_, _, err := userClient.UserSvcAPI.SaveInvites(context.Background()).
			Body(openapi.UserSvcSaveInvitesRequest{
				Invites: []openapi.UserSvcNewInvite{
					{
						ContactId: "test-user@email.com",
						RoleId:    "test-user-slug-0:custom-role",
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
		require.Equal(t, 2, len(claim.RoleIds), claim.RoleIds)
		require.Contains(
			t,
			claim.RoleIds,
			"test-user-slug-0:custom-role",
			claim.RoleIds,
		)
	})
}

func TestListInviteAuthorization(t *testing.T) {
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

	t.Run("user adds role to second user so both can invite", func(t *testing.T) {
		_, _, err := userClient.UserSvcAPI.AssignRole(
			context.Background(),
			tokens[1].UserId,
			"test-user-slug-0:custom-role",
		).Execute()

		require.NoError(t, err)
	})

	t.Run("first user invites", func(t *testing.T) {
		_, _, err := userClient.UserSvcAPI.SaveInvites(context.Background()).
			Body(openapi.UserSvcSaveInvitesRequest{
				Invites: []openapi.UserSvcNewInvite{
					{
						ContactId: "test-user@email.com",
						RoleId:    "test-user-slug-0:custom-role",
					},
				},
			}).Execute()

		require.NoError(t, err)

		rsp, _, err := userClient.UserSvcAPI.ListInvites(context.Background()).
			Body(openapi.UserSvcListInvitesRequest{}).Execute()

		require.NoError(t, err)
		require.Len(t, rsp.Invites, 1)
		require.Len(t, rsp.Invites[0].OwnerIds, 0)
	})

	secondUserClient, _, err = test.LoggedInClient(clientFactory, "test-user-slug-1", "testUserPassword1")
	require.NoError(t, err)

	t.Run("second user cannot invite as it has the role but does not own it", func(t *testing.T) {
		_, _, err := secondUserClient.UserSvcAPI.SaveInvites(context.Background()).
			Body(openapi.UserSvcSaveInvitesRequest{
				Invites: []openapi.UserSvcNewInvite{
					{
						ContactId: "test-user@email.com",
						RoleId:    "test-user-slug-0:custom-role",
					},
				},
			}).Execute()

		require.Error(t, err)

		rsp, _, err := secondUserClient.UserSvcAPI.ListInvites(context.Background()).
			Body(openapi.UserSvcListInvitesRequest{}).Execute()

		require.NoError(t, err)
		require.Len(t, rsp.Invites, 0)
	})
}
