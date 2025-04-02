package basicservice_test

import (
	"context"
	"net/http/httptest"
	"testing"

	basicservice "github.com/1backend/1backend/examples/go/services/basic/internal"

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	basicclient "github.com/1backend/1backend/examples/go/services/basic/client/go"
)

func TestListPets(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		Test: true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	selfUrl := test.NewSelfUrl(t)

	service, err := basicservice.NewService(&basicservice.Options{
		Test:      true,
		ServerUrl: server.Url,
		SelfUrl:   selfUrl,
	})
	require.NoError(t, err)

	basicServer := httptest.NewServer(service.Router)
	defer basicServer.Close()

	service.Options.SelfUrl = basicServer.URL
	require.NoError(t, service.Start())

	clientFactory := client.NewApiClientFactory(server.Url)

	serverAdminClient, _, err := test.AdminClient(clientFactory)
	assert.NoError(t, err)

	serverClients, tokens, err := test.MakeClients(clientFactory, 1)
	require.NoError(t, err)

	//basicminClient := newClient(server.Url, basicminToken)
	client1 := newClient(server.Url, tokens[0].Token)
	//client2 := newClient(server.Url, tokens[1].Token)

	t.Run("users cannot create pets", func(t *testing.T) {
		_, rsp, err := client1.BasicSvcAPI.SavePet(context.Background()).Body(
			basicclient.BasicSvcSavePetRequest{
				Name: "test",
			},
		).Execute()

		require.Error(t, err, rsp)
	})

	t.Run("admin assigns role to user", func(t *testing.T) {
		_, rsp, err := serverAdminClient.UserSvcAPI.AssignRole(
			context.Background(),
			tokens[0].UserId,
			basicservice.RolePetManager,
		).Execute()

		require.NoError(t, err, rsp)
	})

	loginRsp, _, err := serverClients[0].UserSvcAPI.Login(context.Background()).Body(
		openapi.UserSvcLoginRequest{
			Slug:     openapi.PtrString("test-user-slug-0"),
			Password: openapi.PtrString("testUserPassword0"),
		},
	).Execute()
	require.NoError(t, err)

	client1 = newClient(server.Url, loginRsp.Token.Token)

	t.Run("user now can create pets", func(t *testing.T) {
		_, rsp, err := client1.BasicSvcAPI.SavePet(context.Background()).Body(
			basicclient.BasicSvcSavePetRequest{
				Name: "test",
			},
		).Execute()

		require.NoError(t, err, rsp)
	})
}

func newClient(url, token string) *basicclient.APIClient {
	return basicclient.NewAPIClient(&basicclient.Configuration{
		Servers: basicclient.ServerConfigurations{
			{
				URL:         url,
				Description: "Default server",
			},
		},
		DefaultHeader: map[string]string{
			"Authorization": "Bearer " + token,
		},
	})
}
