package save

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExpandEndpoint(t *testing.T) {
	endpoint, err := expandEndpoint("/policy-svc/instance/{id}", map[string]interface{}{
		"id": "login rate limit",
	})
	require.NoError(t, err)
	require.Equal(t, "/policy-svc/instance/login%20rate%20limit", endpoint)
}

func TestExpandEndpointNestedField(t *testing.T) {
	endpoint, err := expandEndpoint("/apps/{app.host}/configs/{id}", map[string]interface{}{
		"app": map[string]interface{}{
			"host": "example.com",
		},
		"id": "theme",
	})
	require.NoError(t, err)
	require.Equal(t, "/apps/example.com/configs/theme", endpoint)
}

func TestExpandEndpointMissingField(t *testing.T) {
	_, err := expandEndpoint("/policy-svc/instance/{id}", map[string]interface{}{})
	require.Error(t, err)
}
