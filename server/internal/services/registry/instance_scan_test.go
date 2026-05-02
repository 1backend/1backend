package registryservice

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/datastore/localstore"
	registry "github.com/1backend/1backend/server/internal/services/registry/types"
	"github.com/stretchr/testify/require"
)

func TestScanInstanceSkipsRecentHealthyHeartbeatWrite(t *testing.T) {
	t.Parallel()

	backend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer backend.Close()

	store, err := localstore.NewLocalStore(&registry.Instance{}, "")
	require.NoError(t, err)

	previousHeartbeat := time.Now()
	instance := &registry.Instance{
		Id:            "inst-test",
		URL:           backend.URL,
		Status:        registry.InstanceStatusHealthy,
		LastHeartbeat: previousHeartbeat,
	}
	require.NoError(t, store.Create(instance))

	rs := &RegistryService{instanceStore: store}
	require.NoError(t, rs.scanInstance(instance))

	row, found, err := store.Query(datastore.Id(instance.Id)).FindOne()
	require.NoError(t, err)
	require.True(t, found)

	stored := row.(*registry.Instance)
	require.Equal(t, registry.InstanceStatusHealthy, stored.Status)
	require.True(t, stored.LastHeartbeat.Equal(previousHeartbeat))
}

func TestScanInstancePersistsStaleHealthyHeartbeat(t *testing.T) {
	t.Parallel()

	backend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer backend.Close()

	store, err := localstore.NewLocalStore(&registry.Instance{}, "")
	require.NoError(t, err)

	previousHeartbeat := time.Now().Add(-instanceHeartbeatWriteInterval - time.Second)
	instance := &registry.Instance{
		Id:            "inst-test",
		URL:           backend.URL,
		Status:        registry.InstanceStatusHealthy,
		LastHeartbeat: previousHeartbeat,
	}
	require.NoError(t, store.Create(instance))

	rs := &RegistryService{instanceStore: store}
	require.NoError(t, rs.scanInstance(instance))

	row, found, err := store.Query(datastore.Id(instance.Id)).FindOne()
	require.NoError(t, err)
	require.True(t, found)

	stored := row.(*registry.Instance)
	require.Equal(t, registry.InstanceStatusHealthy, stored.Status)
	require.True(t, stored.LastHeartbeat.After(previousHeartbeat))
}
