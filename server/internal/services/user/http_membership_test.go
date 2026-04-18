package userservice_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	openapi "github.com/1backend/1backend/clients/go"
	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/test"
)

func testOrgUserRole(orgId string) string {
	return "user-svc:org:{" + orgId + "}:user"
}

func testOrgAdminRole(orgId string) string {
	return "user-svc:org:{" + orgId + "}:admin"
}

func testOrgMemberRole(orgId string, userId string) string {
	return "user-svc:org:{" + orgId + "}:" + userId
}

func TestMembershipLifecycle(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{Test: true})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)
	clients, _, err := test.MakeClients(clientFactory, sdk.DefaultTestAppHost, 2)
	require.NoError(t, err)

	ownerClient := clients[0]
	memberClient := clients[1]

	publicKeyRsp, _, err := clientFactory.Client().UserSvcAPI.GetPublicKey(context.Background()).Execute()
	require.NoError(t, err)

	memberSelf, _, err := memberClient.UserSvcAPI.ReadSelf(context.Background()).Execute()
	require.NoError(t, err)
	memberUserId := memberSelf.User.Id

	org1Rsp, _, err := ownerClient.UserSvcAPI.SaveOrganization(context.Background()).
		Body(openapi.UserSvcSaveOrganizationRequest{
			Activate: openapi.PtrBool(true),
			Slug:     "membership-org-1",
			Name:     openapi.PtrString("Membership Org 1"),
		}).
		Execute()
	require.NoError(t, err)
	require.NotNil(t, org1Rsp.Token)

	org1Id := org1Rsp.Organization.Id
	ownerClient = clientFactory.Client(client.WithToken(org1Rsp.Token.Token))

	org1Claims, err := auth.AuthorizerImpl{}.ParseJWT(publicKeyRsp.PublicKey, org1Rsp.Token.Token)
	require.NoError(t, err)
	require.Equal(t, org1Id, org1Claims.ActiveOrganizationId)
	require.Contains(t, org1Claims.Roles, "user-svc:org:{"+org1Id+"}:user")
	require.Contains(t, org1Claims.Roles, "user-svc:org:{"+org1Id+"}:admin")

	creatorMemberships, httpResp := listMemberships(
		t,
		ownerClient,
		&openapi.UserSvcListMembershipsRequest{OrganizationId: openapi.PtrString(org1Id)},
	)
	require.Equal(t, http.StatusOK, httpResp.StatusCode)
	require.Len(t, creatorMemberships.Memberships, 1)
	require.Equal(t, openapi.MembershipStatusAccepted, creatorMemberships.Memberships[0].Membership.GetStatus())
	require.ElementsMatch(
		t,
		[]string{
			testOrgUserRole(org1Id),
			testOrgMemberRole(org1Id, org1Claims.UserId),
			testOrgAdminRole(org1Id),
		},
		creatorMemberships.Memberships[0].Membership.Roles,
	)

	inviteRsp, inviteHTTP := saveMembership(
		t,
		ownerClient,
		org1Id,
		memberUserId,
		&openapi.UserSvcSaveMembershipRequest{},
	)
	require.Equal(t, http.StatusOK, inviteHTTP.StatusCode)
	require.Equal(t, openapi.MembershipStatusPending, inviteRsp.Membership.GetStatus())
	require.ElementsMatch(
		t,
		[]string{
			testOrgUserRole(org1Id),
			testOrgMemberRole(org1Id, memberUserId),
		},
		inviteRsp.Membership.Roles,
	)

	memberBeforeAccept, _, err := memberClient.UserSvcAPI.ReadSelf(context.Background()).Execute()
	require.NoError(t, err)
	require.Empty(t, memberBeforeAccept.Organizations)

	ownMemberships, ownHTTP := listMemberships(
		t,
		memberClient,
		&openapi.UserSvcListMembershipsRequest{},
	)
	require.Equal(t, http.StatusOK, ownHTTP.StatusCode)
	require.Len(t, ownMemberships.Memberships, 1)
	require.Equal(t, openapi.MembershipStatusPending, ownMemberships.Memberships[0].Membership.GetStatus())

	memberLoginDevice := func(device string) (*openapi.UserSvcLoginResponse, *auth.Claims) {
		loginRsp, _, err := clientFactory.Client().UserSvcAPI.Login(context.Background()).
			Body(openapi.UserSvcLoginRequest{
				AppHost:  sdk.DefaultTestAppHost,
				Slug:     openapi.PtrString("test-user-slug-1"),
				Password: openapi.PtrString("testUserPassword1"),
				Device:   openapi.PtrString(device),
			}).
			Execute()
		require.NoError(t, err)

		claims, err := auth.AuthorizerImpl{}.ParseJWT(publicKeyRsp.PublicKey, loginRsp.Token.Token)
		require.NoError(t, err)

		return loginRsp, claims
	}

	memberLoginRsp, memberClaims := memberLoginDevice("browser-a")
	require.Empty(t, memberClaims.ActiveOrganizationId)
	require.NotContains(t, memberClaims.Roles, "user-svc:org:{"+org1Id+"}:user")
	memberClient = clientFactory.Client(client.WithToken(memberLoginRsp.Token.Token))

	memberMemberships, memberMembershipsHTTP := listMemberships(
		t,
		memberClient,
		&openapi.UserSvcListMembershipsRequest{},
	)
	require.Equal(t, http.StatusOK, memberMembershipsHTTP.StatusCode)
	require.Len(t, memberMemberships.Memberships, 1)
	require.Equal(t, openapi.MembershipStatusPending, memberMemberships.Memberships[0].Membership.GetStatus())
	require.Equal(t, org1Id, memberMemberships.Memberships[0].Organization.GetId())

	orgMemberships, orgMembershipsHTTP := listMemberships(
		t,
		ownerClient,
		&openapi.UserSvcListMembershipsRequest{OrganizationId: openapi.PtrString(org1Id)},
	)
	require.Equal(t, http.StatusOK, orgMembershipsHTTP.StatusCode)
	require.Len(t, orgMemberships.Memberships, 2)

	acceptNoActivateRsp, acceptNoActivateHTTP := acceptMembership(
		t,
		memberClient,
		org1Id,
		&openapi.UserSvcAcceptMembershipRequest{Activate: openapi.PtrBool(false)},
	)
	require.Equal(t, http.StatusOK, acceptNoActivateHTTP.StatusCode)
	require.Nil(t, acceptNoActivateRsp.Token)
	require.Equal(t, openapi.MembershipStatusAccepted, acceptNoActivateRsp.Membership.GetStatus())

	memberAfterAccept, _, err := memberClient.UserSvcAPI.ReadSelf(context.Background()).Execute()
	require.NoError(t, err)
	require.Len(t, memberAfterAccept.Organizations, 1)
	require.Equal(t, org1Id, memberAfterAccept.Organizations[0].Id)
	require.Empty(t, memberAfterAccept.GetActiveOrganizationId())

	memberLoginRsp, memberClaims = memberLoginDevice("browser-a")
	require.Empty(t, memberClaims.ActiveOrganizationId)
	require.NotContains(t, memberClaims.Roles, "user-svc:org:{"+org1Id+"}:user")
	memberClient = clientFactory.Client(client.WithToken(memberLoginRsp.Token.Token))

	org2Rsp, _, err := ownerClient.UserSvcAPI.SaveOrganization(context.Background()).
		Body(openapi.UserSvcSaveOrganizationRequest{
			Activate: openapi.PtrBool(false),
			Slug:     "membership-org-2",
			Name:     openapi.PtrString("Membership Org 2"),
		}).
		Execute()
	require.NoError(t, err)
	require.Nil(t, org2Rsp.Token)
	org2Id := org2Rsp.Organization.Id

	adminInviteRsp, adminInviteHTTP := saveMembership(
		t,
		ownerClient,
		org2Id,
		memberUserId,
		&openapi.UserSvcSaveMembershipRequest{
			Roles: []string{"user-svc:org:{" + org2Id + "}:admin"},
		},
	)
	require.Equal(t, http.StatusOK, adminInviteHTTP.StatusCode)
	require.Equal(t, openapi.MembershipStatusPending, adminInviteRsp.Membership.GetStatus())
	require.ElementsMatch(
		t,
		[]string{
			testOrgUserRole(org2Id),
			testOrgMemberRole(org2Id, memberUserId),
			testOrgAdminRole(org2Id),
		},
		adminInviteRsp.Membership.Roles,
	)

	acceptActivateRsp, acceptActivateHTTP := acceptMembership(
		t,
		memberClient,
		org2Id,
		&openapi.UserSvcAcceptMembershipRequest{Activate: openapi.PtrBool(true)},
	)
	require.Equal(t, http.StatusOK, acceptActivateHTTP.StatusCode)
	require.NotNil(t, acceptActivateRsp.Token)
	require.Equal(t, openapi.MembershipStatusAccepted, acceptActivateRsp.Membership.GetStatus())

	memberClaims, err = auth.AuthorizerImpl{}.ParseJWT(publicKeyRsp.PublicKey, acceptActivateRsp.Token.Token)
	require.NoError(t, err)
	require.Equal(t, org2Id, memberClaims.ActiveOrganizationId)
	require.Contains(t, memberClaims.Roles, "user-svc:org:{"+org2Id+"}:user")
	require.Contains(t, memberClaims.Roles, "user-svc:org:{"+org2Id+"}:admin")
	require.NotContains(t, memberClaims.Roles, "user-svc:org:{"+org1Id+"}:user")

	memberClient = clientFactory.Client(client.WithToken(acceptActivateRsp.Token.Token))

	memberLoginRsp, memberClaims = memberLoginDevice("browser-a")
	require.Equal(t, org2Id, memberClaims.ActiveOrganizationId)
	require.Contains(t, memberClaims.Roles, "user-svc:org:{"+org2Id+"}:admin")
	memberClient = clientFactory.Client(client.WithToken(memberLoginRsp.Token.Token))

	activateOrg1Rsp, _, err := memberClient.UserSvcAPI.ActivateOrganization(context.Background()).
		Body(openapi.UserSvcActivateOrganizationRequest{OrganizationId: org1Id}).
		Execute()
	require.NoError(t, err)

	memberClaims, err = auth.AuthorizerImpl{}.ParseJWT(publicKeyRsp.PublicKey, activateOrg1Rsp.Token.Token)
	require.NoError(t, err)
	require.Equal(t, org1Id, memberClaims.ActiveOrganizationId)
	require.Contains(t, memberClaims.Roles, "user-svc:org:{"+org1Id+"}:user")
	require.NotContains(t, memberClaims.Roles, "user-svc:org:{"+org1Id+"}:admin")
	require.NotContains(t, memberClaims.Roles, "user-svc:org:{"+org2Id+"}:admin")
	memberClient = clientFactory.Client(client.WithToken(activateOrg1Rsp.Token.Token))

	roleUpdateRsp, roleUpdateHTTP := saveMembership(
		t,
		ownerClient,
		org1Id,
		memberUserId,
		&openapi.UserSvcSaveMembershipRequest{
			Roles: []string{"user-svc:org:{" + org1Id + "}:admin"},
		},
	)
	require.Equal(t, http.StatusOK, roleUpdateHTTP.StatusCode)
	require.Equal(t, openapi.MembershipStatusAccepted, roleUpdateRsp.Membership.GetStatus())
	require.ElementsMatch(
		t,
		[]string{
			testOrgUserRole(org1Id),
			testOrgMemberRole(org1Id, memberUserId),
			testOrgAdminRole(org1Id),
		},
		roleUpdateRsp.Membership.Roles,
	)

	memberLoginRsp, memberClaims = memberLoginDevice("browser-a")
	require.Equal(t, org1Id, memberClaims.ActiveOrganizationId)
	require.Contains(t, memberClaims.Roles, "user-svc:org:{"+org1Id+"}:admin")
	require.Contains(t, memberClaims.Roles, "user-svc:org:{"+org1Id+"}:user")

	_, deleteHTTP, err := ownerClient.UserSvcAPI.DeleteMembership(context.Background(), org1Id, memberUserId).Execute()
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, deleteHTTP.StatusCode)

	memberLoginRsp, memberClaims = memberLoginDevice("browser-a")
	require.Empty(t, memberClaims.ActiveOrganizationId)
	require.NotContains(t, memberClaims.Roles, "user-svc:org:{"+org1Id+"}:user")
	require.NotContains(t, memberClaims.Roles, "user-svc:org:{"+org1Id+"}:admin")
	require.NotContains(t, memberClaims.Roles, "user-svc:org:{"+org2Id+"}:admin")
}

func TestMembershipInviteAuthorization(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{Test: true})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)
	clients, _, err := test.MakeClients(clientFactory, sdk.DefaultTestAppHost, 3)
	require.NoError(t, err)

	ownerClient := clients[0]
	memberClient := clients[1]
	thirdClient := clients[2]

	publicKeyRsp, _, err := clientFactory.Client().UserSvcAPI.GetPublicKey(context.Background()).Execute()
	require.NoError(t, err)

	memberSelf, _, err := memberClient.UserSvcAPI.ReadSelf(context.Background()).Execute()
	require.NoError(t, err)
	memberUserId := memberSelf.User.Id

	orgRsp, _, err := ownerClient.UserSvcAPI.SaveOrganization(context.Background()).
		Body(openapi.UserSvcSaveOrganizationRequest{
			Activate: openapi.PtrBool(true),
			Slug:     "membership-invite-auth-org",
			Name:     openapi.PtrString("Membership Invite Auth Org"),
		}).
		Execute()
	require.NoError(t, err)
	require.NotNil(t, orgRsp.Token)

	orgId := orgRsp.Organization.Id
	ownerClient = clientFactory.Client(client.WithToken(orgRsp.Token.Token))

	inviteRsp, inviteHTTP := saveMembership(
		t,
		ownerClient,
		orgId,
		memberUserId,
		&openapi.UserSvcSaveMembershipRequest{},
	)
	require.Equal(t, http.StatusOK, inviteHTTP.StatusCode)
	require.Equal(t, openapi.MembershipStatusPending, inviteRsp.Membership.GetStatus())

	loginRsp, _, err := clientFactory.Client().UserSvcAPI.Login(context.Background()).
		Body(openapi.UserSvcLoginRequest{
			AppHost:  sdk.DefaultTestAppHost,
			Slug:     openapi.PtrString("test-user-slug-1"),
			Password: openapi.PtrString("testUserPassword1"),
		}).
		Execute()
	require.NoError(t, err)

	memberClaims, err := auth.AuthorizerImpl{}.ParseJWT(publicKeyRsp.PublicKey, loginRsp.Token.Token)
	require.NoError(t, err)
	require.Empty(t, memberClaims.ActiveOrganizationId)
	require.NotContains(t, memberClaims.Roles, "user-svc:org:{"+orgId+"}:user")

	_, thirdDeleteHTTP, err := thirdClient.UserSvcAPI.DeleteMembership(
		context.Background(),
		orgId,
		memberUserId,
	).Execute()
	require.Error(t, err)
	require.NotNil(t, thirdDeleteHTTP)

	_, ownerDeleteHTTP, err := ownerClient.UserSvcAPI.DeleteMembership(
		context.Background(),
		orgId,
		memberUserId,
	).Execute()
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, ownerDeleteHTTP.StatusCode)
}

func TestMembershipRoleAssignmentOwnership(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{Test: true})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)
	clients, _, err := test.MakeClients(clientFactory, sdk.DefaultTestAppHost, 2)
	require.NoError(t, err)

	ownerClient := clients[0]
	memberClient := clients[1]

	memberSelf, _, err := memberClient.UserSvcAPI.ReadSelf(context.Background()).Execute()
	require.NoError(t, err)
	memberUserId := memberSelf.User.Id

	orgRsp, _, err := ownerClient.UserSvcAPI.SaveOrganization(context.Background()).
		Body(openapi.UserSvcSaveOrganizationRequest{
			Activate: openapi.PtrBool(true),
			Slug:     "membership-role-ownership-org",
			Name:     openapi.PtrString("Membership Role Ownership Org"),
		}).
		Execute()
	require.NoError(t, err)
	require.NotNil(t, orgRsp.Token)

	orgId := orgRsp.Organization.Id
	ownerClient = clientFactory.Client(client.WithToken(orgRsp.Token.Token))

	t.Run("rejects roles outside organization scope", func(t *testing.T) {
		invalidRole := "some-random-role"

		rsp, httpResp, err := ownerClient.UserSvcAPI.SaveMembership(
			context.Background(),
			orgId,
			memberUserId,
		).
			Body(openapi.UserSvcSaveMembershipRequest{
				Roles: []string{invalidRole},
			}).
			Execute()
		require.Error(t, err)
		require.Nil(t, rsp)
		require.NotNil(t, httpResp)
		t.Cleanup(func() {
			closeResponseBody(t, httpResp)
		})
		require.Equal(t, http.StatusBadRequest, httpResp.StatusCode)

		membershipsRsp, membershipsHTTP := listMemberships(
			t,
			ownerClient,
			&openapi.UserSvcListMembershipsRequest{OrganizationId: openapi.PtrString(orgId)},
		)
		require.Equal(t, http.StatusOK, membershipsHTTP.StatusCode)
		require.Len(t, membershipsRsp.Memberships, 1)
	})

	t.Run("rejects org-scoped roles the caller does not own", func(t *testing.T) {
		unownedRole := "user-svc:org:{" + orgId + "}:team:admin"

		rsp, httpResp, err := ownerClient.UserSvcAPI.SaveMembership(
			context.Background(),
			orgId,
			memberUserId,
		).
			Body(openapi.UserSvcSaveMembershipRequest{
				Roles: []string{unownedRole},
			}).
			Execute()
		require.Error(t, err)
		require.Nil(t, rsp)
		require.NotNil(t, httpResp)
		t.Cleanup(func() {
			closeResponseBody(t, httpResp)
		})
		require.Equal(t, http.StatusForbidden, httpResp.StatusCode)

		membershipsRsp, membershipsHTTP := listMemberships(
			t,
			ownerClient,
			&openapi.UserSvcListMembershipsRequest{OrganizationId: openapi.PtrString(orgId)},
		)
		require.Equal(t, http.StatusOK, membershipsHTTP.StatusCode)
		require.Len(t, membershipsRsp.Memberships, 1)
	})

	t.Run("allows org-scoped roles the caller owns", func(t *testing.T) {
		inviteRsp, inviteHTTP := saveMembership(
			t,
			ownerClient,
			orgId,
			memberUserId,
			&openapi.UserSvcSaveMembershipRequest{
				Roles: []string{"user-svc:org:{" + orgId + "}:admin"},
			},
		)
		require.Equal(t, http.StatusOK, inviteHTTP.StatusCode)
		require.Equal(t, openapi.MembershipStatusPending, inviteRsp.Membership.GetStatus())
		require.ElementsMatch(
			t,
			[]string{
				testOrgUserRole(orgId),
				testOrgMemberRole(orgId, memberUserId),
				testOrgAdminRole(orgId),
			},
			inviteRsp.Membership.Roles,
		)
	})
}

func TestMembershipActivationPerDevice(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{Test: true})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)
	clients, _, err := test.MakeClients(clientFactory, sdk.DefaultTestAppHost, 2)
	require.NoError(t, err)

	ownerClient := clients[0]
	memberClient := clients[1]

	publicKeyRsp, _, err := clientFactory.Client().UserSvcAPI.GetPublicKey(context.Background()).Execute()
	require.NoError(t, err)

	memberSelf, _, err := memberClient.UserSvcAPI.ReadSelf(context.Background()).Execute()
	require.NoError(t, err)
	memberUserId := memberSelf.User.Id

	memberRole := func(orgId string) string {
		return "user-svc:org:{" + orgId + "}:user"
	}

	loginOnDevice := func(device string) (string, *auth.Claims, *openapi.UserSvcReadSelfResponse) {
		loginRsp, _, err := clientFactory.Client().UserSvcAPI.Login(context.Background()).
			Body(openapi.UserSvcLoginRequest{
				AppHost:  sdk.DefaultTestAppHost,
				Slug:     openapi.PtrString("test-user-slug-1"),
				Password: openapi.PtrString("testUserPassword1"),
				Device:   openapi.PtrString(device),
			}).
			Execute()
		require.NoError(t, err)

		claims, err := auth.AuthorizerImpl{}.ParseJWT(publicKeyRsp.PublicKey, loginRsp.Token.Token)
		require.NoError(t, err)

		selfRsp, _, err := clientFactory.Client(client.WithToken(loginRsp.Token.Token)).
			UserSvcAPI.ReadSelf(context.Background()).
			Execute()
		require.NoError(t, err)

		return loginRsp.Token.Token, claims, selfRsp
	}

	createPendingMembership := func(slug string, name string) string {
		orgRsp, _, err := ownerClient.UserSvcAPI.SaveOrganization(context.Background()).
			Body(openapi.UserSvcSaveOrganizationRequest{
				Activate: openapi.PtrBool(true),
				Slug:     slug,
				Name:     openapi.PtrString(name),
			}).
			Execute()
		require.NoError(t, err)
		require.NotNil(t, orgRsp.Token)

		ownerClient = clientFactory.Client(client.WithToken(orgRsp.Token.Token))

		orgId := orgRsp.Organization.Id
		inviteRsp, inviteHTTP := saveMembership(
			t,
			ownerClient,
			orgId,
			memberUserId,
			&openapi.UserSvcSaveMembershipRequest{},
		)
		require.Equal(t, http.StatusOK, inviteHTTP.StatusCode)
		require.Equal(t, openapi.MembershipStatusPending, inviteRsp.Membership.GetStatus())

		return orgId
	}

	t.Run("fresh membership can be accepted as active=true", func(t *testing.T) {
		orgId := createPendingMembership(
			"membership-accept-active-org",
			"Membership Accept Active Org",
		)
		device := "membership-accept-active-device"

		memberToken, _, _ := loginOnDevice(device)

		acceptRsp, acceptHTTP := acceptMembership(
			t,
			clientFactory.Client(client.WithToken(memberToken)),
			orgId,
			&openapi.UserSvcAcceptMembershipRequest{Activate: openapi.PtrBool(true)},
		)
		require.Equal(t, http.StatusOK, acceptHTTP.StatusCode)
		require.NotNil(t, acceptRsp.Token)
		require.Equal(t, openapi.MembershipStatusAccepted, acceptRsp.Membership.GetStatus())

		claims, err := auth.AuthorizerImpl{}.ParseJWT(publicKeyRsp.PublicKey, acceptRsp.Token.Token)
		require.NoError(t, err)
		require.Equal(t, orgId, claims.ActiveOrganizationId)
		require.Contains(t, claims.Roles, memberRole(orgId))

		selfRsp, _, err := clientFactory.Client(client.WithToken(acceptRsp.Token.Token)).
			UserSvcAPI.ReadSelf(context.Background()).
			Execute()
		require.NoError(t, err)
		require.Equal(t, orgId, selfRsp.GetActiveOrganizationId())

		_, otherClaims, otherSelfRsp := loginOnDevice("membership-accept-active-other-device")
		require.Empty(t, otherClaims.ActiveOrganizationId)
		require.NotContains(t, otherClaims.Roles, memberRole(orgId))
		require.Empty(t, otherSelfRsp.GetActiveOrganizationId())
	})

	t.Run("fresh membership can be accepted as active=false", func(t *testing.T) {
		orgId := createPendingMembership(
			"membership-accept-inactive-org",
			"Membership Accept Inactive Org",
		)
		device := "membership-accept-inactive-device"

		memberToken, _, _ := loginOnDevice(device)

		acceptRsp, acceptHTTP := acceptMembership(
			t,
			clientFactory.Client(client.WithToken(memberToken)),
			orgId,
			&openapi.UserSvcAcceptMembershipRequest{Activate: openapi.PtrBool(false)},
		)
		require.Equal(t, http.StatusOK, acceptHTTP.StatusCode)
		require.Nil(t, acceptRsp.Token)
		require.Equal(t, openapi.MembershipStatusAccepted, acceptRsp.Membership.GetStatus())

		_, claims, selfRsp := loginOnDevice(device)
		require.Empty(t, claims.ActiveOrganizationId)
		require.NotContains(t, claims.Roles, memberRole(orgId))
		require.Empty(t, selfRsp.GetActiveOrganizationId())
	})

	t.Run("accepted membership can transition from inactive to active", func(t *testing.T) {
		orgId := createPendingMembership(
			"membership-transition-active-org",
			"Membership Transition Active Org",
		)
		device := "membership-transition-active-device"

		memberToken, _, _ := loginOnDevice(device)

		acceptRsp, acceptHTTP := acceptMembership(
			t,
			clientFactory.Client(client.WithToken(memberToken)),
			orgId,
			&openapi.UserSvcAcceptMembershipRequest{Activate: openapi.PtrBool(false)},
		)
		require.Equal(t, http.StatusOK, acceptHTTP.StatusCode)
		require.Nil(t, acceptRsp.Token)

		memberToken, claims, selfRsp := loginOnDevice(device)
		require.Empty(t, claims.ActiveOrganizationId)
		require.NotContains(t, claims.Roles, memberRole(orgId))
		require.Empty(t, selfRsp.GetActiveOrganizationId())

		activateRsp, activateHTTP, err := clientFactory.Client(client.WithToken(memberToken)).
			UserSvcAPI.ActivateOrganization(context.Background()).
			Body(openapi.UserSvcActivateOrganizationRequest{OrganizationId: orgId}).
			Execute()
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, activateHTTP.StatusCode)

		claims, err = auth.AuthorizerImpl{}.ParseJWT(publicKeyRsp.PublicKey, activateRsp.Token.Token)
		require.NoError(t, err)
		require.Equal(t, orgId, claims.ActiveOrganizationId)
		require.Contains(t, claims.Roles, memberRole(orgId))

		selfRsp, _, err = clientFactory.Client(client.WithToken(activateRsp.Token.Token)).
			UserSvcAPI.ReadSelf(context.Background()).
			Execute()
		require.NoError(t, err)
		require.Equal(t, orgId, selfRsp.GetActiveOrganizationId())
	})

	t.Run("accepted membership can transition from active to inactive", func(t *testing.T) {
		orgId := createPendingMembership(
			"membership-transition-inactive-org",
			"Membership Transition Inactive Org",
		)
		device := "membership-transition-inactive-device"

		memberToken, _, _ := loginOnDevice(device)

		acceptRsp, acceptHTTP := acceptMembership(
			t,
			clientFactory.Client(client.WithToken(memberToken)),
			orgId,
			&openapi.UserSvcAcceptMembershipRequest{Activate: openapi.PtrBool(true)},
		)
		require.Equal(t, http.StatusOK, acceptHTTP.StatusCode)
		require.NotNil(t, acceptRsp.Token)

		deactivateRsp, deactivateHTTP, err := clientFactory.Client(client.WithToken(acceptRsp.Token.Token)).
			UserSvcAPI.DeactivateOrganization(context.Background()).
			Execute()
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, deactivateHTTP.StatusCode)

		claims, err := auth.AuthorizerImpl{}.ParseJWT(publicKeyRsp.PublicKey, deactivateRsp.Token.Token)
		require.NoError(t, err)
		require.Empty(t, claims.ActiveOrganizationId)
		require.NotContains(t, claims.Roles, memberRole(orgId))

		selfRsp, _, err := clientFactory.Client(client.WithToken(deactivateRsp.Token.Token)).
			UserSvcAPI.ReadSelf(context.Background()).
			Execute()
		require.NoError(t, err)
		require.Empty(t, selfRsp.GetActiveOrganizationId())
	})
}

func TestDeclineMembership(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{Test: true})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)
	clients, _, err := test.MakeClients(clientFactory, sdk.DefaultTestAppHost, 2)
	require.NoError(t, err)

	ownerClient := clients[0]
	memberClient := clients[1]

	publicKeyRsp, _, err := clientFactory.Client().UserSvcAPI.GetPublicKey(context.Background()).Execute()
	require.NoError(t, err)

	memberSelf, _, err := memberClient.UserSvcAPI.ReadSelf(context.Background()).Execute()
	require.NoError(t, err)
	memberUserId := memberSelf.User.Id

	loginMember := func(device string) (*openapi.UserSvcLoginResponse, *auth.Claims) {
		loginRsp, _, err := clientFactory.Client().UserSvcAPI.Login(context.Background()).
			Body(openapi.UserSvcLoginRequest{
				AppHost:  sdk.DefaultTestAppHost,
				Slug:     openapi.PtrString("test-user-slug-1"),
				Password: openapi.PtrString("testUserPassword1"),
				Device:   openapi.PtrString(device),
			}).
			Execute()
		require.NoError(t, err)

		claims, err := auth.AuthorizerImpl{}.ParseJWT(publicKeyRsp.PublicKey, loginRsp.Token.Token)
		require.NoError(t, err)

		return loginRsp, claims
	}

	orgRsp, _, err := ownerClient.UserSvcAPI.SaveOrganization(context.Background()).
		Body(openapi.UserSvcSaveOrganizationRequest{
			Activate: openapi.PtrBool(true),
			Slug:     "membership-decline-org",
			Name:     openapi.PtrString("Membership Decline Org"),
		}).
		Execute()
	require.NoError(t, err)
	require.NotNil(t, orgRsp.Token)

	orgId := orgRsp.Organization.Id
	ownerClient = clientFactory.Client(client.WithToken(orgRsp.Token.Token))

	inviteRsp, inviteHTTP := saveMembership(
		t,
		ownerClient,
		orgId,
		memberUserId,
		&openapi.UserSvcSaveMembershipRequest{
			Roles: []string{"user-svc:org:{" + orgId + "}:admin"},
		},
	)
	require.Equal(t, http.StatusOK, inviteHTTP.StatusCode)
	require.Equal(t, openapi.MembershipStatusPending, inviteRsp.Membership.GetStatus())
	require.ElementsMatch(
		t,
		[]string{
			testOrgUserRole(orgId),
			testOrgMemberRole(orgId, memberUserId),
			testOrgAdminRole(orgId),
		},
		inviteRsp.Membership.Roles,
	)

	memberLoginRsp, memberClaims := loginMember("membership-decline-device")
	memberClient = clientFactory.Client(client.WithToken(memberLoginRsp.Token.Token))
	require.Empty(t, memberClaims.ActiveOrganizationId)
	require.NotContains(t, memberClaims.Roles, "user-svc:org:{"+orgId+"}:user")
	require.NotContains(t, memberClaims.Roles, "user-svc:org:{"+orgId+"}:admin")

	declineRsp, declineHTTP := declineMembership(
		t,
		memberClient,
		orgId,
	)
	require.Equal(t, http.StatusOK, declineHTTP.StatusCode)
	require.Equal(t, openapi.MembershipStatusDeclined, declineRsp.Membership.GetStatus())
	require.False(t, declineRsp.Membership.HasAcceptedAt())
	require.ElementsMatch(t, inviteRsp.Membership.Roles, declineRsp.Membership.Roles)

	ownDeclinedMemberships, ownDeclinedHTTP := listMemberships(
		t,
		memberClient,
		&openapi.UserSvcListMembershipsRequest{Status: openapi.MembershipStatusDeclined.Ptr()},
	)
	require.Equal(t, http.StatusOK, ownDeclinedHTTP.StatusCode)
	require.Len(t, ownDeclinedMemberships.Memberships, 1)
	require.Equal(t, openapi.MembershipStatusDeclined, ownDeclinedMemberships.Memberships[0].Membership.GetStatus())
	require.Equal(t, orgId, ownDeclinedMemberships.Memberships[0].Organization.GetId())

	memberLoginRsp, memberClaims = loginMember("membership-decline-device")
	memberClient = clientFactory.Client(client.WithToken(memberLoginRsp.Token.Token))
	require.Empty(t, memberClaims.ActiveOrganizationId)
	require.NotContains(t, memberClaims.Roles, "user-svc:org:{"+orgId+"}:user")
	require.NotContains(t, memberClaims.Roles, "user-svc:org:{"+orgId+"}:admin")

	memberSelfRsp, _, err := clientFactory.Client(client.WithToken(memberLoginRsp.Token.Token)).
		UserSvcAPI.ReadSelf(context.Background()).
		Execute()
	require.NoError(t, err)
	require.Empty(t, memberSelfRsp.GetActiveOrganizationId())
	require.Empty(t, memberSelfRsp.Organizations)

	reinviteRsp, reinviteHTTP := saveMembership(
		t,
		ownerClient,
		orgId,
		memberUserId,
		&openapi.UserSvcSaveMembershipRequest{},
	)
	require.Equal(t, http.StatusOK, reinviteHTTP.StatusCode)
	require.Equal(t, openapi.MembershipStatusPending, reinviteRsp.Membership.GetStatus())

	memberLoginRsp, memberClaims = loginMember("membership-decline-device")
	memberClient = clientFactory.Client(client.WithToken(memberLoginRsp.Token.Token))
	require.Empty(t, memberClaims.ActiveOrganizationId)
	require.NotContains(t, memberClaims.Roles, "user-svc:org:{"+orgId+"}:user")

	acceptRsp, acceptHTTP := acceptMembership(
		t,
		memberClient,
		orgId,
		&openapi.UserSvcAcceptMembershipRequest{Activate: openapi.PtrBool(true)},
	)
	require.Equal(t, http.StatusOK, acceptHTTP.StatusCode)
	require.NotNil(t, acceptRsp.Token)
	require.Equal(t, openapi.MembershipStatusAccepted, acceptRsp.Membership.GetStatus())

	memberClaims, err = auth.AuthorizerImpl{}.ParseJWT(publicKeyRsp.PublicKey, acceptRsp.Token.Token)
	require.NoError(t, err)
	require.Equal(t, orgId, memberClaims.ActiveOrganizationId)
	require.Contains(t, memberClaims.Roles, "user-svc:org:{"+orgId+"}:user")
}

func TestSaveEnrollsRejectsOrgScopedRoles(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{Test: true})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)
	clients, _, err := test.MakeClients(clientFactory, sdk.DefaultTestAppHost, 2)
	require.NoError(t, err)

	ownerClient := clients[0]
	receiverClient := clients[1]

	receiverSelf, _, err := receiverClient.UserSvcAPI.ReadSelf(context.Background()).Execute()
	require.NoError(t, err)

	orgRsp, _, err := ownerClient.UserSvcAPI.SaveOrganization(context.Background()).
		Body(openapi.UserSvcSaveOrganizationRequest{
			Activate: openapi.PtrBool(true),
			Slug:     "enroll-reject-org",
			Name:     openapi.PtrString("Enroll Reject Org"),
		}).
		Execute()
	require.NoError(t, err)

	_, httpResp, err := ownerClient.UserSvcAPI.SaveEnrolls(context.Background()).
		Body(openapi.UserSvcSaveEnrollsRequest{
			Enrolls: []openapi.UserSvcEnrollInput{
				{
					UserId: openapi.PtrString(receiverSelf.User.Id),
					Role:   "user-svc:org:{" + orgRsp.Organization.Id + "}:user",
				},
			},
		}).
		Execute()
	require.Error(t, err)
	require.Equal(t, http.StatusBadRequest, httpResp.StatusCode)
}
