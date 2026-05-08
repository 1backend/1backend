package di

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/1backend/1backend/sdk/go/testutil"
	"github.com/1backend/1backend/server/internal/universe"
	"github.com/stretchr/testify/require"
)

func TestReadyzRequiresStarterFunc(t *testing.T) {
	univ, err := BigBang(&universe.Options{
		Test: true,
	})
	require.NoError(t, err)

	require.Equal(t, http.StatusServiceUnavailable, readyzStatus(t, univ))
	require.Equal(t, http.StatusOK, routeStatus(t, univ, "/healthz"))
	require.Equal(t, http.StatusOK, routeStatus(t, univ, "/livez"))

	univ.started.Store(true)
	require.Equal(t, http.StatusOK, readyzStatus(t, univ))
}

func TestReadyzChecksPostgresAndDistributedLock(t *testing.T) {
	conn := testutil.StartPostgres(t)
	univ, err := BigBang(&universe.Options{
		Test:               true,
		Db:                 "postgres",
		DbConnectionString: conn,
	})
	require.NoError(t, err)

	univ.started.Store(true)
	require.Equal(t, http.StatusOK, readyzStatus(t, univ))
}

func readyzStatus(t *testing.T, univ *Universe) int {
	t.Helper()
	return routeStatus(t, univ, "/readyz")
}

func routeStatus(t *testing.T, univ *Universe, path string) int {
	t.Helper()

	req := httptest.NewRequest(http.MethodGet, path, nil)
	rec := httptest.NewRecorder()
	univ.Router.ServeHTTP(rec, req)
	return rec.Code
}
