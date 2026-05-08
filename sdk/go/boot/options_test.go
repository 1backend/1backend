package boot

import (
	"testing"

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
