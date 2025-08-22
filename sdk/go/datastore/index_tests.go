package datastore

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// Doesn't really test indexes but helps us to play around,
// printline some values and compare.
func TestIndex(t *testing.T, store DataStore) {
	store.SetDebug(false)

	c := 0
	for range 100 {
		objs := []Row{}
		for range 100 {
			c++
			objs = append(objs, TestObject{
				Name:  fmt.Sprintf("%v", c),
				Value: c,
				Age:   c,
			})
		}

		err := store.UpsertMany(objs)
		require.NoError(t, err)
	}

	store.SetDebug(true)

	now := time.Now()
	records, err := store.Query().OrderBy(
		OrderByField("value", true),
	).After(5000).Limit(10).Find()
	dur1 := time.Since(now)
	require.NoError(t, err)
	require.Len(t, records, 10)
	require.Equal(t, "4999", records[0].GetId())
	require.Equal(t, "4998", records[1].GetId())

	now = time.Now()
	records, err = store.Query().OrderBy(
		OrderByField("age", true),
	).After(5000).Limit(10).Find()
	dur2 := time.Since(now)
	require.NoError(t, err)
	require.Len(t, records, 10)
	require.Equal(t, "4999", records[0].GetId())
	require.Equal(t, "4998", records[1].GetId())

	fmt.Println("Duration with index", dur1, "without index", dur2)
}

// Doesn't really test indexes but helps us to play around,
// printline some values and compare.
func TestPointerIndex(t *testing.T, store DataStore) {
	store.SetDebug(false)

	c := 0
	for range 100 {
		objs := []Row{}
		for range 100 {
			c++
			objs = append(objs, &TestObject{
				Name:  fmt.Sprintf("%v", c),
				Value: c,
				Age:   c,
			})
		}

		err := store.UpsertMany(objs)
		require.NoError(t, err)
	}

	store.SetDebug(true)

	now := time.Now()
	records, err := store.Query().OrderBy(
		OrderByField("value", true),
	).After(5000).Limit(10).Find()
	dur1 := time.Since(now)
	require.NoError(t, err)
	require.Len(t, records, 10)
	require.Equal(t, "4999", records[0].GetId())
	require.Equal(t, "4998", records[1].GetId())

	now = time.Now()
	records, err = store.Query().OrderBy(
		OrderByField("age", true),
	).After(5000).Limit(10).Find()
	dur2 := time.Since(now)
	require.NoError(t, err)
	require.Len(t, records, 10)
	require.Equal(t, "4999", records[0].GetId())
	require.Equal(t, "4998", records[1].GetId())

	fmt.Println("Duration with index", dur1, "without index", dur2)
}
