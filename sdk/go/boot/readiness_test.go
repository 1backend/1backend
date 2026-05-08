package boot

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

type failingHealthDataStoreFactory struct{}

func (f failingHealthDataStoreFactory) Create(tableName string, instance any) (datastore.DataStore, error) {
	return nil, errors.New("not implemented")
}

func (f failingHealthDataStoreFactory) Handle() (any, error) {
	return nil, errors.New("writer unavailable")
}

func TestRegisterHealthRoutesReadyzAllowsServicesWithoutDatastore(t *testing.T) {
	router := mux.NewRouter()
	RegisterHealthRoutes(router, HealthRoutesOptions{})

	require.Equal(t, http.StatusOK, routeStatus(t, router, "/healthz"))
	require.Equal(t, http.StatusOK, routeStatus(t, router, "/livez"))
	require.Equal(t, http.StatusOK, routeStatus(t, router, "/readyz"))
}

func TestRegisterHealthRoutesReadyzFailsWhenDatastoreFails(t *testing.T) {
	router := mux.NewRouter()
	RegisterHealthRoutes(router, HealthRoutesOptions{
		DataStoreFactory: failingHealthDataStoreFactory{},
	})

	require.Equal(t, http.StatusServiceUnavailable, routeStatus(t, router, "/readyz"))
}

func routeStatus(t *testing.T, router *mux.Router, path string) int {
	t.Helper()

	req := httptest.NewRequest(http.MethodGet, path, nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	return rec.Code
}
