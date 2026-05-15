package boot

import (
	"testing"

	"github.com/1backend/1backend/sdk/go/client"
	"github.com/stretchr/testify/require"
)

func TestOptionsLoadEnvarsLoadsReadDbConnectionString(t *testing.T) {
	t.Setenv("OB_DB_READ_CONNECTION_STRING", "postgres://reader")

	options := NewOptions("read-db-test", WithTelemetryDisabled())
	require.NoError(t, options.LoadEnvars())

	require.Equal(t, "postgres://reader", options.ReadDbConnectionString)
	require.Equal(t, "postgres://reader", options.DataStoreConfig().ReadDbConnectionString)
}

func TestDataStoreConfigUsesServiceNameAsDbApplicationName(t *testing.T) {
	options := NewOptions("campaign-svc", WithTelemetryDisabled())

	require.Equal(t, "campaign-svc", options.DataStoreConfig().DbApplicationName)
}

func TestOptionsLoadEnvarsLoadsDbApplicationName(t *testing.T) {
	t.Setenv("OB_DB_APPLICATION_NAME", "onebackend-prod")

	options := NewOptions("campaign-svc", WithTelemetryDisabled())
	require.NoError(t, options.LoadEnvars())

	require.Equal(t, "onebackend-prod", options.DbApplicationName)
	require.Equal(t, "onebackend-prod", options.DataStoreConfig().DbApplicationName)
}

func TestOptionsLoadEnvarsSeparatesServerAndSelfURL(t *testing.T) {
	t.Setenv("OB_INTERNAL_SERVER_URL", "http://127.0.0.1:11337")
	t.Setenv("OB_SERVER_URL", "https://deprecated-server.example.test")
	t.Setenv("OB_PUBLIC_URL", "https://basic-svc.example.test")
	t.Setenv("OB_SELF_URL", "https://deprecated-basic-svc.example.test")

	options := NewOptions("basic-svc", WithTelemetryDisabled())
	require.NoError(t, options.LoadEnvars())

	require.Equal(t, "http://127.0.0.1:11337", options.ServerUrl)
	require.Equal(t, "https://basic-svc.example.test", options.SelfUrl)
	require.Equal(t, "http://127.0.0.1:11337", bootClientFactoryURL(t, options.ClientFactory))
}

func TestOptionsLoadEnvarsFallsBackToDeprecatedURLVariables(t *testing.T) {
	t.Setenv("OB_INTERNAL_SERVER_URL", "")
	t.Setenv("OB_SERVER_URL", "http://127.0.0.1:11337")
	t.Setenv("OB_PUBLIC_URL", "")
	t.Setenv("OB_SELF_URL", "https://basic-svc.example.test")

	options := NewOptions("basic-svc", WithTelemetryDisabled())
	require.NoError(t, options.LoadEnvars())

	require.Equal(t, "http://127.0.0.1:11337", options.ServerUrl)
	require.Equal(t, "https://basic-svc.example.test", options.SelfUrl)
	require.Equal(t, "http://127.0.0.1:11337", bootClientFactoryURL(t, options.ClientFactory))
}

func bootClientFactoryURL(t *testing.T, clientFactory client.ClientFactory) string {
	t.Helper()

	factory, ok := clientFactory.(*client.APIClientFactory)
	require.True(t, ok)

	servers := factory.Client().GetConfig().Servers
	require.Len(t, servers, 1)

	return servers[0].URL
}
