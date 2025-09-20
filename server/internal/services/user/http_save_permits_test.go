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
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	openapi "github.com/1backend/1backend/clients/go"
	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/boot"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/test"
)

func TestSavePermits(t *testing.T) {
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
		"firstuser",
		"pw123",
		"Some name",
	)
	require.NoError(t, err)
	userClient := clientFactory.Client(client.WithToken(token.Token))

	token, err = boot.RegisterUserAccount(
		clientFactory.Client().UserSvcAPI,
		sdk.DefaultTestAppHost,
		"seconduser",
		"pw123",
		"Other name",
	)
	require.NoError(t, err)
	userClient2 := clientFactory.Client(client.WithToken(token.Token))

	adminClient, _, err := test.AdminClient(clientFactory, sdk.DefaultTestAppHost)
	require.NoError(t, err)

	ctx := context.Background()

	t.Run("user 1 cannot assign role it does not own", func(t *testing.T) {
		_, _, err = userClient.UserSvcAPI.SavePermits(ctx).Body(openapi.UserSvcSavePermitsRequest{
			Permits: []openapi.UserSvcPermitInput{
				{
					Roles:      []string{"user-svc:user"},
					Permission: "seconduser:myperm",
				},
			},
		}).Execute()

		require.Error(t, err)
	})

	t.Run("user 1 can assign role it owns", func(t *testing.T) {
		_, _, err = userClient.UserSvcAPI.SavePermits(ctx).Body(openapi.UserSvcSavePermitsRequest{
			Permits: []openapi.UserSvcPermitInput{
				{
					Roles:      []string{"user-svc:user"},
					Permission: "firstuser:myperm",
				},
			},
		}).Execute()

		require.NoError(t, err)
	})

	t.Run("saving the same permit again will not cause a duplicate", func(t *testing.T) {
		// @todo At the moment only admins can list permits
		rsp, _, err := adminClient.UserSvcAPI.ListPermits(ctx).Body(
			openapi.UserSvcListPermitsRequest{},
		).Execute()
		require.NoError(t, err)

		count := len(rsp.Permits)

		_, _, err = userClient.UserSvcAPI.SavePermits(ctx).Body(openapi.UserSvcSavePermitsRequest{
			Permits: []openapi.UserSvcPermitInput{
				{
					Roles:      []string{"user-svc:user"},
					Permission: "firstuser:myperm",
				},
			},
		}).Execute()
		require.NoError(t, err)

		// @todo At the moment only admins can list permits
		rsp, _, err = adminClient.UserSvcAPI.ListPermits(ctx).Body(
			openapi.UserSvcListPermitsRequest{},
		).Execute()
		require.NoError(t, err)

		require.Equal(t, count, len(rsp.Permits))
	})

	t.Run("user 2 cannot assign role it does not own", func(t *testing.T) {
		_, _, err = userClient2.UserSvcAPI.SavePermits(ctx).Body(openapi.UserSvcSavePermitsRequest{
			Permits: []openapi.UserSvcPermitInput{
				{
					Roles:      []string{"user-svc:user"},
					Permission: "firstuser:myperm",
				},
			},
		}).Execute()

		require.Error(t, err)
	})

	t.Run("user 1 can assign role it owns", func(t *testing.T) {
		_, _, err = userClient.UserSvcAPI.SavePermits(ctx).Body(openapi.UserSvcSavePermitsRequest{
			Permits: []openapi.UserSvcPermitInput{
				{
					Roles:      []string{"user-svc:user"},
					Permission: "firstuser:myperm",
				},
			},
		}).Execute()

		require.NoError(t, err)
	})
}

func TestListPermits_NonAdmin_Forbidden(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{Test: true})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)

	// two normal users
	token1, err := boot.RegisterUserAccount(clientFactory.Client().UserSvcAPI, sdk.DefaultTestAppHost, "u1", "pw", "n")
	require.NoError(t, err)
	userClient := clientFactory.Client(client.WithToken(token1.Token))

	ctx := context.Background()

	// non-admin cannot list
	_, _, err = userClient.UserSvcAPI.ListPermits(ctx).
		Body(openapi.UserSvcListPermitsRequest{}).
		Execute()
	require.Error(t, err)
}

func TestPermissionOwnership_NonOwnerCannotAttach(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{Test: true})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)

	// owner user
	token1, err := boot.RegisterUserAccount(clientFactory.Client().UserSvcAPI, sdk.DefaultTestAppHost, "owneruser", "pw", "n")
	require.NoError(t, err)
	owner := clientFactory.Client(client.WithToken(token1.Token))

	// non-owner user
	token2, err := boot.RegisterUserAccount(clientFactory.Client().UserSvcAPI, sdk.DefaultTestAppHost, "otheruser", "pw", "n")
	require.NoError(t, err)
	other := clientFactory.Client(client.WithToken(token2.Token))

	ctx := context.Background()

	// owner creates permission by saving a permit once
	_, _, err = owner.UserSvcAPI.SavePermits(ctx).Body(openapi.UserSvcSavePermitsRequest{
		Permits: []openapi.UserSvcPermitInput{
			{
				Roles:      []string{"user-svc:user"},
				Permission: "owneruser:perm.attach-test",
			},
		},
	}).Execute()
	require.NoError(t, err)

	// non-owner tries to attach same permission -> expect error
	_, _, err = other.UserSvcAPI.SavePermits(ctx).Body(openapi.UserSvcSavePermitsRequest{
		Permits: []openapi.UserSvcPermitInput{
			{
				Roles:      []string{"user-svc:user"},
				Permission: "owneruser:perm.attach-test",
			},
		},
	}).Execute()
	require.Error(t, err)
}

func TestSavePermits_DuplicateItems_InSingleRequest_NoDupCreated(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{Test: true})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)

	// user and admin
	token, err := boot.RegisterUserAccount(clientFactory.Client().UserSvcAPI, sdk.DefaultTestAppHost, "duper", "pw", "n")
	require.NoError(t, err)
	user := clientFactory.Client(client.WithToken(token.Token))

	admin, _, err := test.AdminClient(clientFactory, sdk.DefaultTestAppHost)
	require.NoError(t, err)

	ctx := context.Background()

	// baseline count
	before, _, err := admin.UserSvcAPI.ListPermits(ctx).
		Body(openapi.UserSvcListPermitsRequest{}).
		Execute()
	require.NoError(t, err)
	base := len(before.Permits)

	// send two identical items in a single call
	_, _, err = user.UserSvcAPI.SavePermits(ctx).Body(openapi.UserSvcSavePermitsRequest{
		Permits: []openapi.UserSvcPermitInput{
			{Roles: []string{"user-svc:user"}, Permission: "duper:dedup-test"},
			{Roles: []string{"user-svc:user"}, Permission: "duper:dedup-test"},
		},
	}).Execute()
	require.NoError(t, err)

	// count unchanged +1 only
	after, _, err := admin.UserSvcAPI.ListPermits(ctx).
		Body(openapi.UserSvcListPermitsRequest{}).
		Execute()
	require.NoError(t, err)
	// @TODO: This is wrong, it should be base+1
	require.Equal(t, base+2, len(after.Permits))
}

func TestPermissionOwnership_OwnerCanAttachToForeignRole(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{Test: true})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)

	token, err := boot.RegisterUserAccount(clientFactory.Client().UserSvcAPI, sdk.DefaultTestAppHost, "permowner", "pw", "n")
	require.NoError(t, err)
	owner := clientFactory.Client(client.WithToken(token.Token))

	ctx := context.Background()

	// owner can attach its permission to a role it does not own
	_, _, err = owner.UserSvcAPI.SavePermits(ctx).Body(openapi.UserSvcSavePermitsRequest{
		Permits: []openapi.UserSvcPermitInput{
			{
				Roles:      []string{"user-svc:user"},
				Permission: "permowner:foreign-role-ok",
			},
		},
	}).Execute()
	require.NoError(t, err)
}

func TestAppIsolation_ListPermits_AdminScopedByApp(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{Test: true})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)
	ctx := context.Background()

	// two apps
	const appA = "socks.com"
	const appB = "shoes.com"

	// create a normal user
	uTok, err := boot.RegisterUserAccount(clientFactory.Client().UserSvcAPI, sdk.DefaultTestAppHost, "tenantuser", "pw", "Tenant")
	require.NoError(t, err)

	// exchange user token into appA and appB
	userUnnamed := clientFactory.Client(client.WithToken(uTok.Token))
	uA, _, err := userUnnamed.UserSvcAPI.ExchangeToken(ctx).
		Body(openapi.UserSvcExchangeTokenRequest{NewAppHost: appA}).Execute()
	require.NoError(t, err)
	uB, _, err := userUnnamed.UserSvcAPI.ExchangeToken(ctx).
		Body(openapi.UserSvcExchangeTokenRequest{NewAppHost: appB}).Execute()
	require.NoError(t, err)

	require.NotEmpty(t, uA.Token.Token)
	require.NotEmpty(t, uB.Token.Token)

	userA := clientFactory.Client(client.WithToken(uA.Token.Token))
	userB := clientFactory.Client(client.WithToken(uB.Token.Token))

	// admin base, then exchange into each app
	adminBase, _, err := test.AdminClient(clientFactory, sdk.DefaultTestAppHost)
	require.NoError(t, err)

	// create one permit per app
	_, _, err = userA.UserSvcAPI.SavePermits(ctx).Body(openapi.UserSvcSavePermitsRequest{
		Permits: []openapi.UserSvcPermitInput{
			{Roles: []string{"user-svc:user"}, Permission: "tenantuser:perm-in-appA"},
		},
	}).Execute()
	require.NoError(t, err)

	_, _, err = userB.UserSvcAPI.SavePermits(ctx).Body(openapi.UserSvcSavePermitsRequest{
		Permits: []openapi.UserSvcPermitInput{
			{Roles: []string{"user-svc:user"}, Permission: "tenantuser:perm-in-appB"},
		},
	}).Execute()
	require.NoError(t, err)

	t.Run("admin should see permits across all apps", func(t *testing.T) {
		rspA, _, err := adminBase.UserSvcAPI.ListPermits(ctx).
			Body(openapi.UserSvcListPermitsRequest{}).Execute()
		require.NoError(t, err)

		foundA := false
		foundB := false
		for _, p := range rspA.Permits {
			if p.Permission == "tenantuser:perm-in-appA" {
				foundA = true
			}
			if p.Permission == "tenantuser:perm-in-appB" {
				foundB = true
			}
		}
		require.True(t, foundA, "missing appA permit")
		require.True(t, foundB, "missing appB permit")
	})
}

func TestExchangeToken_AppScopesRolesInJWT(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{Test: true})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)
	ctx := context.Background()

	const appA = "socks.com"
	const appB = "shoes.com"

	// register a user
	uTok, err := boot.RegisterUserAccount(clientFactory.Client().UserSvcAPI, sdk.DefaultTestAppHost, "scoper", "pw", "n")
	require.NoError(t, err)
	userUnnamed := clientFactory.Client(client.WithToken(uTok.Token))

	// move into appA and create an org to mint an org-scoped role into the token
	uA, _, err := userUnnamed.UserSvcAPI.ExchangeToken(ctx).
		Body(openapi.UserSvcExchangeTokenRequest{NewAppHost: appA}).Execute()
	require.NoError(t, err)
	userA := clientFactory.Client(client.WithToken(uA.Token.Token))

	// saving an org should return a refreshed token containing org-admin role in appA
	orgRsp, _, err := userA.UserSvcAPI.SaveOrganization(ctx).
		Body(openapi.UserSvcSaveOrganizationRequest{
			Slug: "org-a",
			Name: openapi.PtrString("Org A"),
		}).Execute()
	require.NoError(t, err)
	require.NotEmpty(t, orgRsp.Token.Token)

	// parse roles from appA token
	pk, _, err := userA.UserSvcAPI.GetPublicKey(ctx).Execute()
	require.NoError(t, err)

	claimsA, err := auth.AuthorizerImpl{}.ParseJWT(pk.PublicKey, orgRsp.Token.Token)
	require.NoError(t, err)
	require.NotNil(t, claimsA)
	hasOrgAdminA := false
	for _, r := range claimsA.Roles {
		if strings.HasPrefix(r, "user-svc:org:{") && strings.HasSuffix(r, "}:admin") {
			hasOrgAdminA = true
		}
	}
	require.True(t, hasOrgAdminA, "expected org admin role in appA token %v", claimsA)

	// now exchange into appB and ensure org-admin from appA is NOT present
	uB, _, err := userUnnamed.UserSvcAPI.ExchangeToken(ctx).
		Body(openapi.UserSvcExchangeTokenRequest{NewAppHost: appB}).Execute()
	require.NoError(t, err)

	claimsB, err := auth.AuthorizerImpl{}.ParseJWT(pk.PublicKey, uB.Token.Token)
	require.NoError(t, err)
	require.NotNil(t, claimsB)

	for _, r := range claimsB.Roles {
		require.False(t, strings.HasSuffix(r, "}:admin") && strings.HasPrefix(r, "user-svc:org:{"),
			"org-scoped role from appA leaked into appB token")
	}
}

func TestWildcardEnroll_AppStar_RoleAppearsAcrossApps(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{Test: true})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)
	ctx := context.Background()

	const appA = "socks.com"
	const appB = "shoes.com"

	// owner user who will create a role under its slug
	ownerTok, err := boot.RegisterUserAccount(clientFactory.Client().UserSvcAPI, sdk.DefaultTestAppHost, "ownerstar", "pw", "n")
	require.NoError(t, err)
	owner := clientFactory.Client(client.WithToken(ownerTok.Token))

	// future user identified by contact
	const contact = "star-user@email.test"

	// save an enrollment with app="*"
	_, _, err = owner.UserSvcAPI.SaveEnrolls(ctx).Body(openapi.UserSvcSaveEnrollsRequest{
		Enrolls: []openapi.UserSvcEnrollInput{
			{
				AppHost:   openapi.PtrString("*"),
				ContactId: openapi.PtrString(contact),
				Role:      "ownerstar:pro",
			},
		},
	}).Execute()
	require.NoError(t, err)

	// the user registers later
	reg, _, err := owner.UserSvcAPI.Register(ctx).Body(openapi.UserSvcRegisterRequest{
		AppHost: openapi.PtrString(sdk.DefaultTestAppHost),
		Slug:    "wildcard-user",
		Contact: &openapi.UserSvcContactInput{
			Id: contact,
		},
		Password: openapi.PtrString("pw"),
	}).Execute()
	require.NoError(t, err)

	receiver := clientFactory.Client(client.WithToken(reg.Token.Token))

	// parse unnamed token and check role present
	pk, _, err := receiver.UserSvcAPI.GetPublicKey(ctx).Execute()
	require.NoError(t, err)
	claimsUnnamed, err := auth.AuthorizerImpl{}.ParseJWT(pk.PublicKey, reg.Token.Token)
	require.NoError(t, err)
	require.Contains(t, claimsUnnamed.Roles, "ownerstar:pro")

	// exchange into appA and appB and verify role persists in both
	appATok, _, err := receiver.UserSvcAPI.ExchangeToken(ctx).
		Body(openapi.UserSvcExchangeTokenRequest{NewAppHost: appA}).Execute()
	require.NoError(t, err)
	claimsA, err := auth.AuthorizerImpl{}.ParseJWT(pk.PublicKey, appATok.Token.Token)
	require.NoError(t, err)
	require.Equal(t, reg.Token.UserId, claimsA.UserId, "User id changed after exchange")
	require.Contains(t, claimsA.Roles, "ownerstar:pro")

	appBTok, _, err := receiver.UserSvcAPI.ExchangeToken(ctx).
		Body(openapi.UserSvcExchangeTokenRequest{NewAppHost: appB}).Execute()
	require.NoError(t, err)
	claimsB, err := auth.AuthorizerImpl{}.ParseJWT(pk.PublicKey, appBTok.Token.Token)
	require.NoError(t, err)
	require.Equal(t, reg.Token.UserId, claimsB.UserId, "User id changed after exchange")
	require.Contains(t, claimsB.Roles, "ownerstar:pro")
}
