package multiservice_test

import (
	"context"
	"net/http/httptest"
	"testing"

	basicclient "github.com/1backend/1backend/examples/go/services/basic/client/go"
	multiservice "github.com/1backend/1backend/examples/go/services/multi/internal"

	openapi "github.com/1backend/1backend/clients/go"
	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	multiclient "github.com/1backend/1backend/examples/go/services/multi/client/go"
)

func TestCountPets(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		Test: true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	basicSvcServer, err := test.StartService(test.Options{
		Name:      "basic",
		Test:      true,
		ServerUrl: server.Url,
	})
	require.NoError(t, err)
	defer basicSvcServer.Cleanup(t)

	selfUrl := test.NewSelfUrl(t)

	service, err := multiservice.NewService(&multiservice.Options{
		Test:      true,
		ServerUrl: server.Url,
		SelfUrl:   selfUrl,
	})
	require.NoError(t, err)

	multiServer := httptest.NewServer(service.Router)
	defer multiServer.Close()

	service.Options.SelfUrl = multiServer.URL
	require.NoError(t, service.Start())

	clientFactory := sdk.NewApiClientFactory(server.Url)

	serverAdminClient, _, err := test.AdminClient(clientFactory)
	assert.NoError(t, err)

	serverClients, tokens, err := test.MakeClients(clientFactory, 1)
	require.NoError(t, err)

	t.Run("admin assigns role to user so user can save a pet first", func(t *testing.T) {
		_, rsp, err := serverAdminClient.UserSvcAPI.AssignRole(
			context.Background(),
			tokens[0].UserId,
			"basic-svc:pet:manager",
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

	basicSvcClient := newBasicSvcClient(server.Url, loginRsp.Token.Token)

	t.Run("save pet", func(t *testing.T) {
		_, rsp, err := basicSvcClient.BasicSvcAPI.SavePet(context.Background()).Body(
			basicclient.BasicSvcSavePetRequest{
				Name: "test",
			},
		).Execute()

		require.NoError(t, err, rsp)
	})

	multiSvcClient := newMultiSvcClient(server.Url, loginRsp.Token.Token)

	t.Run("count pets", func(t *testing.T) {
		rsp, _, err := multiSvcClient.MultiSvcAPI.
			CountPets(context.Background()).Execute()

		require.NoError(t, err, rsp)
		require.Equal(t, 1, int(rsp.PetCount))
	})
}

func newBasicSvcClient(url, token string) *basicclient.APIClient {
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

func newMultiSvcClient(url, token string) *multiclient.APIClient {
	return multiclient.NewAPIClient(&multiclient.Configuration{
		Servers: multiclient.ServerConfigurations{
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
