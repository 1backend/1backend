package localstore

import (
	"testing"

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/stretchr/testify/require"
)

func TestLocalStoreAutoIndexStatsNoop(t *testing.T) {
	store, err := NewLocalStore(datastore.TestObject{}, "")
	require.NoError(t, err)
	defer store.Close()

	filter := datastore.Equals(datastore.Field("Name"), "missing")

	_, err = store.Query(filter).Find()
	require.NoError(t, err)

	_, _, err = store.Query(filter).FindOne()
	require.NoError(t, err)

	_, err = store.Query(filter).Count()
	require.NoError(t, err)

	err = store.Query(filter).UpdateFields(map[string]any{"Age": 42})
	require.NoError(t, err)

	err = store.Query(filter).Delete()
	require.NoError(t, err)

	stats := store.AutoIndexStats()
	require.False(t, stats.Supported)
	require.Equal(t, "localstore", stats.Backend)
	require.Empty(t, stats.Indexes)
	require.Len(t, stats.Shapes, 1)
	require.Equal(t, 5, stats.Shapes[0].Hits)
}
