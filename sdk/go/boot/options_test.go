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
