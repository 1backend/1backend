package di

import (
	"testing"

	sdkclient "github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/server/internal/router"
	"github.com/1backend/1backend/server/internal/universe"
	"github.com/stretchr/testify/require"
)

func TestBigBangUsesOBInternalServerURLForInternalClientFactory(t *testing.T) {
	t.Cleanup(func() {
		router.SetPort(11337)
	})
	t.Setenv("OB_PUBLIC_URL", "https://api.example.test")
	t.Setenv("OB_SELF_URL", "https://deprecated-api.example.test")
	t.Setenv("OB_INTERNAL_SERVER_URL", "http://127.0.0.1:11337")
	t.Setenv("OB_SERVER_URL", "https://deprecated-internal.example.test")

	univ, err := BigBang(&universe.Options{Test: true})
	require.NoError(t, err)

	require.Equal(t, "https://api.example.test", univ.Options.Url)
	require.Equal(t, "http://127.0.0.1:11337", univ.Options.ServerUrl)
	require.Equal(t, "http://127.0.0.1:11337", apiClientFactoryURL(t, univ.Options.ClientFactory))
}

func TestBigBangFallsBackToDeprecatedURLVariables(t *testing.T) {
	t.Cleanup(func() {
		router.SetPort(11337)
	})
	t.Setenv("OB_PUBLIC_URL", "")
	t.Setenv("OB_SELF_URL", "https://api.example.test")
	t.Setenv("OB_INTERNAL_SERVER_URL", "")
	t.Setenv("OB_SERVER_URL", "http://127.0.0.1:11337")

	univ, err := BigBang(&universe.Options{Test: true})
	require.NoError(t, err)

	require.Equal(t, "https://api.example.test", univ.Options.Url)
	require.Equal(t, "http://127.0.0.1:11337", univ.Options.ServerUrl)
	require.Equal(t, "http://127.0.0.1:11337", apiClientFactoryURL(t, univ.Options.ClientFactory))
}

func TestBigBangFallsBackToSelfURLForInternalClientFactory(t *testing.T) {
	t.Cleanup(func() {
		router.SetPort(11337)
	})
	t.Setenv("OB_PUBLIC_URL", "")
	t.Setenv("OB_SELF_URL", "")
	t.Setenv("OB_INTERNAL_SERVER_URL", "")
	t.Setenv("OB_SERVER_URL", "")

	univ, err := BigBang(&universe.Options{
		Test: true,
		Url:  "http://127.0.0.1:17777",
	})
	require.NoError(t, err)

	require.Equal(t, "http://127.0.0.1:17777", univ.Options.Url)
	require.Equal(t, "http://127.0.0.1:17777", univ.Options.ServerUrl)
	require.Equal(t, "http://127.0.0.1:17777", apiClientFactoryURL(t, univ.Options.ClientFactory))
}

func apiClientFactoryURL(t *testing.T, clientFactory sdkclient.ClientFactory) string {
	t.Helper()

	factory, ok := clientFactory.(*sdkclient.APIClientFactory)
	require.True(t, ok)

	servers := factory.Client().GetConfig().Servers
	require.Len(t, servers, 1)

	return servers[0].URL
}
