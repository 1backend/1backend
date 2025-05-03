package userservice_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/test"

	openapi "github.com/1backend/1backend/clients/go"
)

func TestReadSelf(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		Test: true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)
	var userClient *openapi.APIClient

	t.Run("register", func(t *testing.T) {
		rsp, _, err := clientFactory.Client().UserSvcAPI.Register(context.Background()).
			Body(openapi.UserSvcRegisterRequest{
				Slug:     "some-slug",
				Password: openapi.PtrString("some password"),
				Contact: &openapi.UserSvcContact{
					Id:       "xyz",
					Platform: "email",
				},
			}).Execute()

		require.NoError(t, err)
		userClient = clientFactory.Client(client.WithToken(rsp.Token.Token))
	})

	t.Run("read self", func(t *testing.T) {
		rsp, _, err := userClient.UserSvcAPI.ReadSelf(context.Background()).Execute()
		require.NoError(t, err)

		require.Equal(t, rsp.User.Slug, "some-slug")
		require.NotEmpty(t, rsp.Contacts)
		require.Equal(t, rsp.Contacts[0].Id, "xyz")
		require.Equal(t, rsp.Contacts[0].Platform, "email")
	})
}
