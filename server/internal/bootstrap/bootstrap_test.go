package bootstrap

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	config "github.com/1backend/1backend/server/internal/services/config/types"
	policy "github.com/1backend/1backend/server/internal/services/policy/types"
	proxy "github.com/1backend/1backend/server/internal/services/proxy/types"
	secret "github.com/1backend/1backend/server/internal/services/secret/types"
	user "github.com/1backend/1backend/server/internal/services/user/types"
)

func TestApplyLoadsSupportedManifests(t *testing.T) {
	dir := t.TempDir()

	writeFile(t, filepath.Join(dir, "permits", "_meta.yaml"), `entity: "user-svc:permit"`)
	writeFile(t, filepath.Join(dir, "permits", "report.yaml"), `id: "report-svc-list-users"
appHost: "*"
permission: "user-svc:user:view"
slugs:
  - "report-svc"
`)
	writeFile(t, filepath.Join(dir, "apps", "site.example.com", "permits", "_meta.yaml"), `entity: "user-svc:permit"`)
	writeFile(t, filepath.Join(dir, "apps", "site.example.com", "permits", "site.yaml"), `id: "site-example-config"
appHost: "site.example.com"
permission: "config-svc:config:edit-on-behalf"
slugs:
  - "site-svc"
`)
	writeFile(t, filepath.Join(dir, "apps", "_all", "permits", "_meta.yaml"), `entity: "user-svc:permit"`)
	writeFile(t, filepath.Join(dir, "apps", "_all", "permits", "global.yaml"), `id: "global-list-users"
appHost: "*"
permission: "user-svc:user:view"
slugs:
  - "report-svc"
`)
	writeFile(t, filepath.Join(dir, "apps", "example.com", "permits", "explicitness.yaml"), `id: "path-does-not-default-app"
permission: "user-svc:user:view"
slugs:
  - "report-svc"
`)
	writeFile(t, filepath.Join(dir, "apps", "tenant.example.com", "enrolls", "_meta.yaml"), `entity: "user-svc:enroll"`)
	writeFile(t, filepath.Join(dir, "apps", "tenant.example.com", "enrolls", "owner.yaml"), `id: "site-owner"
appHost: "tenant.example.com"
contactId: "owner@example.com"
role: "site-svc:admin"
`)
	writeFile(t, filepath.Join(dir, "apps", "tenant.example.com", "configs", "_meta.yaml"), `entity: "config-svc:config"`)
	writeFile(t, filepath.Join(dir, "apps", "tenant.example.com", "configs", "payment.yaml"), `id: "paymentSvc"
appHost: "tenant.example.com"
data:
  publicKey: "pk_test"
`)
	writeFile(t, filepath.Join(dir, "routes", "_meta.yaml"), `entity: "proxy-svc:route"`)
	writeFile(t, filepath.Join(dir, "routes", "api.yaml"), `id: "api.example.com"
target: "http://1backend:11337"
`)
	writeFile(t, filepath.Join(dir, "redirects", "_meta.yaml"), `entity: "proxy-svc:redirect"`)
	writeFile(t, filepath.Join(dir, "redirects", "old-api.yaml"), `id: "old-api.example.com"
target: "https://api.example.com"
statusCode: 301
`)
	writeFile(t, filepath.Join(dir, "apps", "example.com", "secrets", "_meta.yaml"), `entity: "secret-svc:secret"`)
	writeFile(t, filepath.Join(dir, "apps", "example.com", "secrets", "api-key.yaml"), `id: "api-key"
appHost: "example.com"
value: "encrypted-secret"
encrypted: true
checksum: "12345678"
checksumAlgorithm: "CRC32"
`)
	writeFile(t, filepath.Join(dir, "policies", "login-rate-limit.yaml"), `id: "login-rate-limit"
endpoint: "/user-svc/login"
templateId: "rate-limit"
parameters:
  rateLimit:
    maxRequests: 20
    timeWindow: "1m"
    entity: "ip"
    scope: "endpoint"
`)
	writeFile(t, filepath.Join(dir, "apps", "example.com", "unclassified.yaml"), `id: "looks-secretish"
value: "but-has-no-meta"
`)

	var permits []user.PermitInput
	var enrolls []user.EnrollInput
	var routes []proxy.RouteInput
	var redirects []proxy.RedirectInput
	var configs []config.SaveConfigRequest
	var secrets []*secret.SecretInput
	var policyInstances []*policy.Instance

	summary, err := Apply(context.Background(), dir, Services{
		SavePermits: func(_ context.Context, items []user.PermitInput) error {
			permits = append(permits, items...)
			return nil
		},
		SaveEnrolls: func(_ context.Context, items []user.EnrollInput) error {
			enrolls = append(enrolls, items...)
			return nil
		},
		SaveRoutes: func(_ context.Context, items []proxy.RouteInput) error {
			routes = append(routes, items...)
			return nil
		},
		SaveRedirects: func(_ context.Context, items []proxy.RedirectInput) error {
			redirects = append(redirects, items...)
			return nil
		},
		SaveConfigs: func(_ context.Context, items []config.SaveConfigRequest) error {
			configs = append(configs, items...)
			return nil
		},
		SaveSecrets: func(_ context.Context, items []*secret.SecretInput) error {
			secrets = append(secrets, items...)
			return nil
		},
		SavePolicyInstances: func(_ context.Context, items []*policy.Instance) error {
			policyInstances = append(policyInstances, items...)
			return nil
		},
	})
	require.NoError(t, err)

	require.Len(t, permits, 4)
	requirePermit(t, permits, "report-svc-list-users", "*")
	requirePermit(t, permits, "site-example-config", "site.example.com")
	requirePermit(t, permits, "global-list-users", "*")
	requirePermit(t, permits, "path-does-not-default-app", "")

	require.Len(t, enrolls, 1)
	require.Equal(t, "site-owner", enrolls[0].Id)
	require.Equal(t, "tenant.example.com", enrolls[0].AppHost)
	require.Equal(t, "owner@example.com", enrolls[0].ContactId)
	require.Equal(t, "site-svc:admin", enrolls[0].Role)

	require.Len(t, routes, 1)
	require.Equal(t, "api.example.com", routes[0].Id)
	require.Equal(t, "http://1backend:11337", routes[0].Target)

	require.Len(t, redirects, 1)
	require.Equal(t, "old-api.example.com", redirects[0].Id)
	require.Equal(t, "https://api.example.com", redirects[0].Target)
	require.Equal(t, 301, redirects[0].StatusCode)

	require.Len(t, configs, 1)
	require.Equal(t, "paymentSvc", configs[0].Id)
	require.Equal(t, "tenant.example.com", configs[0].AppHost)
	require.Equal(t, "pk_test", configs[0].Data["publicKey"])

	require.Len(t, secrets, 1)
	require.Equal(t, "api-key", secrets[0].Id)
	require.Equal(t, "example.com", secrets[0].AppHost)
	require.Equal(t, "encrypted-secret", secrets[0].Value)
	require.True(t, secrets[0].Encrypted)
	require.Equal(t, "12345678", secrets[0].Checksum)

	require.Len(t, policyInstances, 1)
	require.Equal(t, "login-rate-limit", policyInstances[0].Id)
	require.Equal(t, "/user-svc/login", policyInstances[0].Endpoint)
	require.Equal(t, policy.TemplateIdRateLimit, policyInstances[0].TemplateId)
	require.Equal(t, 20, policyInstances[0].Parameters.RateLimit.MaxRequests)
	require.Equal(t, policy.EntityIP, policyInstances[0].Parameters.RateLimit.Entity)

	require.Equal(t, 4, summary.AppliedPermits)
	require.Equal(t, 1, summary.AppliedEnrolls)
	require.Equal(t, 1, summary.AppliedRoutes)
	require.Equal(t, 1, summary.AppliedRedirects)
	require.Equal(t, 1, summary.AppliedConfigs)
	require.Equal(t, 1, summary.AppliedSecrets)
	require.Equal(t, 1, summary.AppliedPolicyInstances)
	require.Equal(t, 1, summary.SkippedUnsupported)
}

func requirePermit(t *testing.T, permits []user.PermitInput, id, appHost string) {
	t.Helper()
	for _, permit := range permits {
		if permit.Id == id {
			require.Equal(t, appHost, permit.AppHost)
			return
		}
	}
	require.Failf(t, "missing permit", "permit %s not found", id)
}

func writeFile(t *testing.T, path, content string) {
	t.Helper()
	require.NoError(t, os.MkdirAll(filepath.Dir(path), 0o755))
	require.NoError(t, os.WriteFile(path, []byte(content), 0o644))
}
