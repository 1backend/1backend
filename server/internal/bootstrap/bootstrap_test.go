package bootstrap

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	user "github.com/1backend/1backend/server/internal/services/user/types"
)

func TestApplyLoadsSupportedManifestsAndSkipsSecrets(t *testing.T) {
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
data:
  publicKey: "pk_test"
`)
	writeFile(t, filepath.Join(dir, "routes", "_meta.yaml"), `entity: "proxy-svc:route"`)
	writeFile(t, filepath.Join(dir, "routes", "api.yaml"), `id: "api.example.com"
target: "http://1backend:11337"
`)
	writeFile(t, filepath.Join(dir, "apps", "example.com", "secrets", "_meta.yaml"), `entity: "secret-svc:secret"`)
	writeFile(t, filepath.Join(dir, "apps", "example.com", "secrets", "api-key.yaml"), `id: "api-key"
value: "secret"
`)
	writeFile(t, filepath.Join(dir, "apps", "example.com", "unclassified.yaml"), `id: "looks-secretish"
value: "but-has-no-meta"
`)

	var permits []user.PermitInput
	var enrolls []user.EnrollInput

	summary, err := Apply(context.Background(), dir, Services{
		SavePermits: func(_ context.Context, items []user.PermitInput) error {
			permits = append(permits, items...)
			return nil
		},
		SaveEnrolls: func(_ context.Context, items []user.EnrollInput) error {
			enrolls = append(enrolls, items...)
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

	require.Equal(t, 4, summary.AppliedPermits)
	require.Equal(t, 1, summary.AppliedEnrolls)
	require.Equal(t, 1, summary.SkippedSecrets)
	require.Equal(t, 3, summary.SkippedUnsupported)
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
