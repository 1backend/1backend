package proxyservice_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/test"

	openapi "github.com/1backend/1backend/clients/go"
)

func TestListCerts(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		TokenExpiration: 1 * time.Second,
		Test:            true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)
	adminClient, _, err := test.AdminClient(clientFactory)
	require.NoError(t, err)

	clients, _, err := test.MakeClients(
		clientFactory, 1)
	require.NoError(t, err)

	client := clients[0]

	t.Run("user cannot save cert", func(t *testing.T) {
		req := openapi.ProxySvcSaveCertsRequest{
			Certs: []openapi.ProxySvcCertInput{
				{
					Id:   "test.localhost",
					Cert: certString1,
				},
			},
		}

		_, hrsp, err := client.ProxySvcAPI.SaveCerts(
			context.Background(),
		).Body(req).Execute()

		require.Error(t, err, hrsp)

		_, hrsp, err = adminClient.ProxySvcAPI.SaveCerts(
			context.Background(),
		).Body(req).Execute()

		require.NoError(t, err, hrsp)
	})

	t.Run("unauthorized cannot list routes", func(t *testing.T) {
		_, _, err := client.ProxySvcAPI.ListCerts(
			context.Background(),
		).Execute()

		require.Error(t, err)
	})

	t.Run("list all certs", func(t *testing.T) {
		rsp, hrsp, err := adminClient.ProxySvcAPI.ListCerts(
			context.Background(),
		).Execute()

		require.NoError(t, err, hrsp)
		require.NotNil(t, rsp)
		require.Len(t, rsp.Certs, 1)
	})

	t.Run("list cert by id", func(t *testing.T) {
		rsp, hrsp, err := adminClient.ProxySvcAPI.ListCerts(
			context.Background(),
		).Execute()

		require.NoError(t, err, hrsp)
		require.NotNil(t, rsp)

		var certIDs []string
		for _, cert := range rsp.Certs {
			if cert.Id != "" {
				certIDs = append(certIDs, cert.Id)
			}
		}

		require.Contains(t, certIDs, "test.localhost")
	})
}
