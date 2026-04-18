/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package userservice_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	openapi "github.com/1backend/1backend/clients/go"
	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/boot"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/test"
	user_svc "github.com/1backend/1backend/server/internal/services/user/types"
)

func TestUnauthorizedShouldNotReturnError(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		Test: true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)

	token, err := boot.RegisterUserAccount(
		clientFactory.Client().UserSvcAPI,
		sdk.DefaultTestAppHost,
		"someuser",
		"pw123",
		"Some name",
	)
	require.NoError(t, err)
	userClient := clientFactory.Client(client.WithToken(token.Token))

	ctx := context.Background()

	rsp, _, err := userClient.UserSvcAPI.
		HasPermission(ctx, "user.view").
		Execute()
	require.NoError(t, err)
	require.False(t, rsp.Authorized)
	require.NotNil(t, rsp.AppId)
	require.Equal(t, sdk.DefaultTestAppHost, rsp.App.Host)
	require.NotEmpty(t, rsp.User)

	t.Run("not logged in user should not return error", func(t *testing.T) {
		_, hrsp, err := clientFactory.Client().UserSvcAPI.
			HasPermission(context.Background(), "user.view").
			Execute()

		require.Error(t, err)
		require.Equal(t, 422, hrsp.StatusCode)
	})
}

func TestPermitsBySlug(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		Test: true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)

	token, err := boot.RegisterUserAccount(
		clientFactory.Client().UserSvcAPI,
		sdk.DefaultTestAppHost,
		"someuser",
		"pw123",
		"Some name",
	)
	require.NoError(t, err)
	userClient := clientFactory.Client(client.WithToken(token.Token))

	ctx := context.Background()

	_, _, err = userClient.UserSvcAPI.ListUsers(ctx).Execute()
	require.Error(t, err)

	adminClient, _, err := test.AdminClient(clientFactory, sdk.DefaultTestAppHost)
	require.NoError(t, err)

	_, _, err = adminClient.UserSvcAPI.SavePermits(ctx).Body(openapi.UserSvcSavePermitsRequest{
		Permits: []openapi.UserSvcPermitInput{
			{
				Slugs:      []string{"someuser"},
				Permission: user_svc.PermissionUserView,
			},
		},
	}).Execute()
	require.NoError(t, err)

	rsp, _, err := userClient.UserSvcAPI.ListUsers(ctx).Execute()
	require.NoError(t, err)
	require.NotEmpty(t, len(rsp.Users))
}

func TestPermitsByRoleId(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		Test: true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)

	token, err := boot.RegisterUserAccount(
		clientFactory.Client().UserSvcAPI,
		sdk.DefaultTestAppHost,
		"someuser",
		"pw123",
		"Some name",
	)
	require.NoError(t, err)
	userClient := clientFactory.Client(client.WithToken(token.Token))

	ctx := context.Background()

	_, _, err = userClient.UserSvcAPI.ListUsers(ctx).Execute()
	require.Error(t, err)

	adminClient, _, err := test.AdminClient(clientFactory, sdk.DefaultTestAppHost)
	require.NoError(t, err)

	_, _, err = adminClient.UserSvcAPI.SavePermits(ctx).Body(openapi.UserSvcSavePermitsRequest{
		Permits: []openapi.UserSvcPermitInput{
			{
				Roles:      []string{"user-svc:user"},
				Permission: user_svc.PermissionUserView,
			},
		},
	}).Execute()
	require.NoError(t, err)

	rsp, _, err := userClient.UserSvcAPI.ListUsers(ctx).Execute()
	require.NoError(t, err)
	require.NotEmpty(t, len(rsp.Users))
}

// Mirrors the documented org-specific permission pattern in user-svc.md:
// 1. org admin creates membership
// 2. membership gets canonical role user-svc:org:{orgId}:{userId}
// 3. org admin saves an org-owned permission to that role
// 4. HasPermission succeeds only while that org is active
func TestOrgScopedPermissionPattern_DocsFlow(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{Test: true})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)
	ctx := context.Background()

	clients, _, err := test.MakeClients(clientFactory, sdk.DefaultTestAppHost, 2)
	require.NoError(t, err)

	ownerClient := clients[0]
	memberClient := clients[1]

	pk, _, err := clientFactory.Client().UserSvcAPI.GetPublicKey(ctx).Execute()
	require.NoError(t, err)

	memberSelf, _, err := memberClient.UserSvcAPI.ReadSelf(ctx).Execute()
	require.NoError(t, err)
	memberUserId := memberSelf.User.Id

	org1Rsp, _, err := ownerClient.UserSvcAPI.SaveOrganization(ctx).
		Body(openapi.UserSvcSaveOrganizationRequest{
			Activate: openapi.PtrBool(true),
			Slug:     "drift-member-role-org-1",
			Name:     openapi.PtrString("Notes Member Role Org 1"),
		}).
		Execute()
	require.NoError(t, err)
	require.NotNil(t, org1Rsp.Token)

	org1Id := org1Rsp.Organization.Id
	ownerClient = clientFactory.Client(client.WithToken(org1Rsp.Token.Token))

	org2Rsp, _, err := ownerClient.UserSvcAPI.SaveOrganization(ctx).
		Body(openapi.UserSvcSaveOrganizationRequest{
			Activate: openapi.PtrBool(false),
			Slug:     "drift-member-role-org-2",
			Name:     openapi.PtrString("Notes Member Role Org 2"),
		}).
		Execute()
	require.NoError(t, err)

	org2Id := org2Rsp.Organization.Id

	inviteOrg1Rsp, inviteOrg1HTTP := saveMembership(
		t,
		ownerClient,
		org1Id,
		memberUserId,
		&openapi.UserSvcSaveMembershipRequest{},
	)
	require.Equal(t, 200, inviteOrg1HTTP.StatusCode)
	require.ElementsMatch(
		t,
		[]string{
			"user-svc:org:{" + org1Id + "}:user",
			"user-svc:org:{" + org1Id + "}:" + memberUserId,
		},
		inviteOrg1Rsp.Membership.Roles,
	)

	_, inviteOrg2HTTP := saveMembership(
		t,
		ownerClient,
		org2Id,
		memberUserId,
		&openapi.UserSvcSaveMembershipRequest{},
	)
	require.Equal(t, 200, inviteOrg2HTTP.StatusCode)

	loginRsp, _, err := clientFactory.Client().UserSvcAPI.Login(ctx).
		Body(openapi.UserSvcLoginRequest{
			AppHost:  sdk.DefaultTestAppHost,
			Slug:     openapi.PtrString("test-user-slug-1"),
			Password: openapi.PtrString("testUserPassword1"),
			Device:   openapi.PtrString("drift-browser"),
		}).
		Execute()
	require.NoError(t, err)

	memberClient = clientFactory.Client(client.WithToken(loginRsp.Token.Token))

	acceptOrg1Rsp, acceptOrg1HTTP := acceptMembership(
		t,
		memberClient,
		org1Id,
		&openapi.UserSvcAcceptMembershipRequest{Activate: openapi.PtrBool(true)},
	)
	require.Equal(t, 200, acceptOrg1HTTP.StatusCode)
	require.NotNil(t, acceptOrg1Rsp.Token)

	memberClient = clientFactory.Client(client.WithToken(acceptOrg1Rsp.Token.Token))

	acceptOrg2Rsp, acceptOrg2HTTP := acceptMembership(
		t,
		memberClient,
		org2Id,
		&openapi.UserSvcAcceptMembershipRequest{Activate: openapi.PtrBool(false)},
	)
	require.Equal(t, 200, acceptOrg2HTTP.StatusCode)
	require.Nil(t, acceptOrg2Rsp.Token)

	memberRole := "user-svc:org:{" + org1Id + "}:" + memberUserId
	claims, err := auth.AuthorizerImpl{}.ParseJWT(pk.PublicKey, acceptOrg1Rsp.Token.Token)
	require.NoError(t, err)
	require.Equal(t, org1Id, claims.ActiveOrganizationId)
	require.Contains(t, claims.Roles, memberRole)

	orgScopedPermission := "user-svc:org:{" + org1Id + "}:notes-svc:note:edit"
	_, _, err = ownerClient.UserSvcAPI.SavePermits(ctx).Body(openapi.UserSvcSavePermitsRequest{
		Permits: []openapi.UserSvcPermitInput{
			{
				Permission: orgScopedPermission,
				Roles:      []string{memberRole},
			},
		},
	}).Execute()
	require.NoError(t, err)

	hasInOrg1, _, err := memberClient.UserSvcAPI.HasPermission(ctx, orgScopedPermission).Execute()
	require.NoError(t, err)
	require.True(t, hasInOrg1.Authorized)

	activateOrg2Rsp, _, err := memberClient.UserSvcAPI.ActivateOrganization(ctx).
		Body(openapi.UserSvcActivateOrganizationRequest{OrganizationId: org2Id}).
		Execute()
	require.NoError(t, err)

	memberClient = clientFactory.Client(client.WithToken(activateOrg2Rsp.Token.Token))

	org2Claims, err := auth.AuthorizerImpl{}.ParseJWT(pk.PublicKey, activateOrg2Rsp.Token.Token)
	require.NoError(t, err)
	require.Equal(t, org2Id, org2Claims.ActiveOrganizationId)
	require.NotContains(t, org2Claims.Roles, memberRole)

	hasInOrg2, _, err := memberClient.UserSvcAPI.HasPermission(ctx, orgScopedPermission).Execute()
	require.NoError(t, err)
	require.False(t, hasInOrg2.Authorized)
}

func TestAutoRefresh(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		// JWT expiration is only second granular
		TokenExpiration: time.Second,
		Test:            true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)

	token, err := boot.RegisterUserAccount(
		clientFactory.Client().UserSvcAPI,
		sdk.DefaultTestAppHost,
		"someuser",
		"pw123",
		"Some name",
	)
	require.NoError(t, err)
	userClient := clientFactory.Client(client.WithToken(token.Token))

	ctx := context.Background()

	rsp, _, err := userClient.UserSvcAPI.ReadSelf(ctx).Execute()
	require.NoError(t, err)
	require.Equal(t, "someuser", rsp.User.Slug)

	time.Sleep(2 * time.Second)

	rsp, _, err = userClient.UserSvcAPI.ReadSelf(ctx).Execute()
	require.NoError(t, err)
	require.Equal(t, "someuser", rsp.User.Slug)
}

func TestAutoRefreshOff(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		// JWT expiration is only second granular
		TokenExpiration:     time.Second,
		TokenAutoRefreshOff: true,
		Test:                true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)

	token, err := boot.RegisterUserAccount(
		clientFactory.Client().UserSvcAPI,
		sdk.DefaultTestAppHost,
		"someuser",
		"pw123",
		"Some name",
	)
	require.NoError(t, err)
	userClient := clientFactory.Client(client.WithToken(token.Token))

	ctx := context.Background()

	rsp, _, err := userClient.UserSvcAPI.ReadSelf(ctx).Execute()
	require.NoError(t, err)
	require.Equal(t, "someuser", rsp.User.Slug)

	time.Sleep(2 * time.Second)

	_, _, err = userClient.UserSvcAPI.ReadSelf(ctx).Execute()
	require.Error(t, err)
}
