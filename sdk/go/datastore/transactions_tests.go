package datastore

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransactions(t *testing.T, store DataStore) {
	tx, err := store.BeginTransaction()
	require.NoError(t, err)

	obj := TestObject{Name: "test", Value: 10}
	err = tx.Create(obj)
	require.NoError(t, err)

	results, err := store.Query(Equals(Field("Name"), "test")).Find()
	require.NoError(t, err)
	require.Len(t, results, 0)

	err = tx.Commit()
	require.NoError(t, err)

	results, err = store.Query(Equals(Field("Name"), "test")).Find()
	require.NoError(t, err)
	require.Len(t, results, 1)
	readObj := results[0]
	require.Equal(t, obj, readObj)
}

func TestPointerTransactions(t *testing.T, store DataStore) {
	tx, err := store.BeginTransaction()
	require.NoError(t, err)

	obj := &TestObject{Name: "test", Value: 10}
	err = tx.Create(obj)
	require.NoError(t, err)

	results, err := store.Query(Equals(Field("Name"), "test")).Find()
	require.NoError(t, err)
	require.Len(t, results, 0)

	err = tx.Commit()
	require.NoError(t, err)

	results, err = store.Query(Equals(Field("Name"), "test")).Find()
	require.NoError(t, err)
	require.Len(t, results, 1)
	readObj := results[0]
	require.Equal(t, obj, readObj)
}
