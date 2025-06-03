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
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/test"
	"github.com/1backend/1backend/server/internal/di"
	"github.com/1backend/1backend/server/internal/universe"
)

func TestChangePassword(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		Test: true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)

	manyClients, _, err := test.MakeClients(clientFactory, 2)
	require.NoError(t, err)

	client1 := manyClients[0]

	adminClient, _, err := test.AdminClient(clientFactory)
	require.NoError(t, err)

	ctx := context.Background()

	t.Run("without current password, change password should fail", func(t *testing.T) {
		newPw := "a test that should not work"

		_, hrsp, err := client1.UserSvcAPI.ChangePassword(ctx).Body(
			openapi.UserSvcChangePasswordRequest{
				NewPassword: newPw,
			}).Execute()

		require.Error(t, err, hrsp)

		_, _, err = client1.UserSvcAPI.Login(ctx).Body(
			openapi.UserSvcLoginRequest{
				Slug:     openapi.PtrString("test-user-slug-0"),
				Password: openapi.PtrString(newPw),
			}).Execute()

		require.Error(t, err)
	})

	t.Run("user 1 cannot 'use' password of user 2", func(t *testing.T) {
		_, hrsp, err := client1.UserSvcAPI.ChangePassword(ctx).Body(
			openapi.UserSvcChangePasswordRequest{
				CurrentPassword: "testUserPassword1",
				NewPassword:     "a test",
			}).Execute()

		require.Error(t, err, hrsp)
	})

	t.Run("user 1 can change their own password", func(t *testing.T) {
		_, _, err := client1.UserSvcAPI.ChangePassword(ctx).Body(
			openapi.UserSvcChangePasswordRequest{
				CurrentPassword: "testUserPassword0",
				NewPassword:     "a test",
			}).Execute()

		require.NoError(t, err)
	})

	t.Run("admin can change their own password", func(t *testing.T) {
		_, _, err := adminClient.UserSvcAPI.ChangePassword(ctx).Body(
			openapi.UserSvcChangePasswordRequest{
				CurrentPassword: "changeme",
				NewPassword:     "a test",
			}).Execute()

		require.NoError(t, err)
	})

	t.Run("user 1 can login with new password", func(t *testing.T) {
		rsp, _, err := client1.UserSvcAPI.Login(ctx).Body(
			openapi.UserSvcLoginRequest{
				Slug:     openapi.PtrString("test-user-slug-0"),
				Password: openapi.PtrString("a test"),
			}).Execute()

		require.NoError(t, err)
		require.NotEmpty(t, rsp.Token.Token)
	})
}

func TestPasswordChange(t *testing.T) {
	hs := &di.HandlerSwitcher{}
	server := httptest.NewServer(hs)
	defer server.Close()

	options := &universe.Options{
		Test: true,
		Url:  server.URL,
	}
	universe, err := di.BigBang(options)
	require.NoError(t, err)

	hs.UpdateHandler(universe.Router)

	err = universe.StarterFunc()
	require.NoError(t, err)

	userSvc := options.ClientFactory.Client().UserSvcAPI

	manyClients, _, err := test.MakeClients(options.ClientFactory, 1)

	require.NoError(t, err)

	userClient := manyClients[0]
	userToken := auth.TokenFromClient(userClient)

	require.Equal(t, true, len(userToken) > 0)

	publicKeyRsp, _, err := userSvc.GetPublicKey(context.Background()).Execute()
	require.NoError(t, err)

	t.Run("user password change", func(t *testing.T) {
		claim, err := options.Authorizer.ParseJWT(
			publicKeyRsp.PublicKey,
			userToken,
		)
		require.NoError(t, err)

		selfRsp, _, err := userClient.UserSvcAPI.ReadSelf(context.Background()).
			Execute()
		require.NoError(t, err)

		require.Equal(t, "test-user-slug-0", selfRsp.User.Slug)

		require.Equal(t, claim.UserId, selfRsp.User.Id)

		changePassReq := openapi.UserSvcChangePasswordRequest{
			CurrentPassword: "testUserPassword0",
			NewPassword:     "yo",
		}
		_, _, err = userSvc.ChangePassword(context.Background()).
			Body(changePassReq).
			Execute()
		require.Error(t, err)

		_, _, err = userClient.UserSvcAPI.ChangePassword(context.Background()).
			Body(changePassReq).
			Execute()
		require.NoError(t, err)

		// changing with wrong password should error
		changePassReq.CurrentPassword = "yoWRONG"
		changePassReq.NewPassword = "yo1"

		_, _, err = userClient.UserSvcAPI.ChangePassword(context.Background()).
			Body(changePassReq).
			Execute()
		require.Error(t, err)
	})
}
