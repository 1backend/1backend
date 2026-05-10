package di

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/1backend/1backend/server/internal/universe"
)

func TestLoadContactAuthProvidersFromMapEnv(t *testing.T) {
	clearContactAuthProviderEnv(t)
	t.Setenv("OB_AUTH_PROVIDERS", `{
		"google": {
			"clientId": "google-client-id",
			"clientSecret": "google-client-secret"
		},
		"okta": {
			"name": "Okta",
			"kind": "oidc",
			"issuerUrl": "https://example.okta.com",
			"clientId": "okta-client-id",
			"clientSecret": "okta-client-secret"
		}
	}`)

	options := &universe.Options{}
	require.NoError(t, loadContactAuthProviders(options))

	require.Equal(t, "google-client-id", options.ContactAuthProviders["google"].ClientID)
	require.Equal(t, "google-client-secret", options.ContactAuthProviders["google"].ClientSecret)
	require.Equal(t, "Okta", options.ContactAuthProviders["okta"].Name)
	require.Equal(t, "oidc", options.ContactAuthProviders["okta"].Kind)
	require.Equal(t, "https://example.okta.com", options.ContactAuthProviders["okta"].IssuerURL)
}

func TestLoadContactAuthProvidersFromGoogleEnv(t *testing.T) {
	clearContactAuthProviderEnv(t)
	t.Setenv("OB_GOOGLE_OIDC_CLIENT_ID", "google-client-id")
	t.Setenv("OB_GOOGLE_OIDC_CLIENT_SECRET", "google-client-secret")

	options := &universe.Options{}
	require.NoError(t, loadContactAuthProviders(options))

	require.Equal(t, "google", options.ContactAuthProviders["google"].Id)
	require.Equal(t, "google-client-id", options.ContactAuthProviders["google"].ClientID)
	require.Equal(t, "google-client-secret", options.ContactAuthProviders["google"].ClientSecret)
}

func clearContactAuthProviderEnv(t *testing.T) {
	t.Helper()
	for _, envName := range []string{
		"OB_AUTH_PROVIDERS",
		"OB_CONTACT_AUTH_PROVIDERS",
		"OB_OIDC_PROVIDERS",
		"OB_GOOGLE_OIDC_CLIENT_ID",
		"OB_GOOGLE_OIDC_CLIENT_SECRET",
		"OB_FACEBOOK_APP_ID",
		"OB_FACEBOOK_APP_SECRET",
		"OB_FACEBOOK_GRAPH_VERSION",
		"OB_SLACK_OIDC_CLIENT_ID",
		"OB_SLACK_OIDC_CLIENT_SECRET",
		"OB_LINKEDIN_OIDC_CLIENT_ID",
		"OB_LINKEDIN_OIDC_CLIENT_SECRET",
		"OB_GITLAB_OIDC_CLIENT_ID",
		"OB_GITLAB_OIDC_CLIENT_SECRET",
		"OB_GITLAB_OIDC_ISSUER_URL",
		"OB_GITHUB_CLIENT_ID",
		"OB_GITHUB_CLIENT_SECRET",
		"OB_DISCORD_CLIENT_ID",
		"OB_DISCORD_CLIENT_SECRET",
	} {
		t.Setenv(envName, "")
	}
}
