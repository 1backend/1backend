package userservice_test

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/test"
)

func TestAppRoles(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		Test: true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)
	ctx := context.Background()

	adminClient, _, err := test.AdminClient(clientFactory)
	require.NoError(t, err)

	var userId string

	t.Run("anyone can register", func(t *testing.T) {
		rsp, _, err := clientFactory.Client().UserSvcAPI.Register(ctx).Body(
			openapi.UserSvcRegisterRequest{
				Slug:     "test-slug-1",
				Name:     openapi.PtrString("Test Name"),
				Password: openapi.PtrString("testPass123"),
			},
		).Execute()
		require.NoError(t, err)
		require.NotNil(t, rsp.Token)
		require.NotNil(t, rsp.Token.AppId)
		require.Equal(t,
			strings.Replace(rsp.Token.App.Host, "http://", "", -1),
			rsp.Token.App.Host,
		)
		userId = rsp.Token.UserId
	})

	t.Run("cannot re-register in a different app", func(t *testing.T) {
		_, _, err := clientFactory.Client().UserSvcAPI.Register(ctx).Body(
			openapi.UserSvcRegisterRequest{
				AppHost:  openapi.PtrString("helloapp"),
				Slug:     "test-slug-1",
				Name:     openapi.PtrString("Test Name"),
				Password: openapi.PtrString("testPass123"),
			},
		).Execute()
		require.Error(t, err)
	})

	t.Run("admin gives user a new role in the default app", func(t *testing.T) {
		_, _, err := adminClient.UserSvcAPI.SaveEnrolls(ctx).
			Body(openapi.UserSvcSaveEnrollsRequest{
				Enrolls: []openapi.UserSvcEnrollInput{
					{
						UserId: openapi.PtrString(userId),
						Role:   "some-role",
					},
				},
			}).Execute()
		require.NoError(t, err)
	})

	t.Run("log in to default app and see the role", func(t *testing.T) {
		rsp, _, err := clientFactory.Client().UserSvcAPI.Login(ctx).Body(
			openapi.UserSvcLoginRequest{
				Slug:     openapi.PtrString("test-slug-1"),
				Password: openapi.PtrString("testPass123"),
			},
		).Execute()
		require.NoError(t, err)

		selfRsp, _, err := clientFactory.Client(client.WithToken(rsp.Token.Token)).
			UserSvcAPI.ReadSelf(ctx).Execute()
		require.NoError(t, err)
		require.Equal(t, 2, len(selfRsp.Roles), selfRsp.Roles)
		require.Contains(t, selfRsp.Roles, "some-role", selfRsp.Roles)
	})

	t.Run("log in to helloapp and NOT see the role", func(t *testing.T) {
		rsp, _, err := clientFactory.Client().UserSvcAPI.Login(ctx).Body(
			openapi.UserSvcLoginRequest{
				AppHost:  openapi.PtrString("helloapp"),
				Slug:     openapi.PtrString("test-slug-1"),
				Password: openapi.PtrString("testPass123"),
			},
		).Execute()
		require.NoError(t, err)

		selfRsp, _, err := clientFactory.Client(client.WithToken(rsp.Token.Token)).
			UserSvcAPI.ReadSelf(ctx).Execute()
		require.NoError(t, err)
		require.Equal(t, 1, len(selfRsp.Roles), selfRsp.Roles)
	})

	t.Run("admin logs in to an other app", func(t *testing.T) {
		rsp, _, err := adminClient.UserSvcAPI.Login(ctx).Body(
			openapi.UserSvcLoginRequest{
				AppHost:  openapi.PtrString("helloapp"),
				Slug:     openapi.PtrString("1backend"),
				Password: openapi.PtrString("changeme"),
			},
		).Execute()
		require.NoError(t, err)
		require.Equal(t, "helloapp", rsp.Token.App.Host)

		clientFactory.Client(client.WithToken(rsp.Token.Token)).
			UserSvcAPI.SaveEnrolls(ctx).
			Body(openapi.UserSvcSaveEnrollsRequest{
				Enrolls: []openapi.UserSvcEnrollInput{
					{
						UserId: openapi.PtrString(userId),
						Role:   "some-role",
					},
				},
			}).Execute()
		require.NoError(t, err)

		t.Run("log in to helloapp and see the role", func(t *testing.T) {
			rsp, _, err := clientFactory.Client().UserSvcAPI.Login(ctx).Body(
				openapi.UserSvcLoginRequest{
					AppHost:  openapi.PtrString("helloapp"),
					Slug:     openapi.PtrString("test-slug-1"),
					Password: openapi.PtrString("testPass123"),
				},
			).Execute()
			require.NoError(t, err)
			require.Equal(t, "helloapp", rsp.Token.App.Host)

			selfRsp, _, err := clientFactory.Client(client.WithToken(rsp.Token.Token)).
				UserSvcAPI.ReadSelf(ctx).Execute()
			require.NoError(t, err)
			require.Equal(t, 2, len(selfRsp.Roles), selfRsp.Roles)
			require.Contains(t, selfRsp.Roles, "some-role", selfRsp.Roles)
		})
	})
}
