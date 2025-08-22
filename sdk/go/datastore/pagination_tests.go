package datastore

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPagination(t *testing.T, store DataStore) {
	for i := 1; i <= 10; i++ {
		now := time.Now()
		obj := TestObject{
			Name:      fmt.Sprintf("PaginationTest%d", i),
			Value:     i,
			CreatedAt: now,
		}
		err := store.Create(obj)
		time.Sleep(1 * time.Millisecond)
		require.NoError(t, err)
	}

	t.Run("test int", func(t *testing.T) {
		results, err := store.Query().OrderBy(
			OrderByField("Value", true),
		).
			Limit(5).
			Find()
		require.NoError(t, err)
		require.Len(t, results, 5)
		require.Equal(t, "PaginationTest10", results[0].(TestObject).Name)
		require.Equal(t, 10, results[0].(TestObject).Value)

		lastValue := results[len(results)-1].(TestObject).Value
		results, err = store.Query().OrderBy(
			OrderByField("Value", true),
		).
			Limit(5).
			After(lastValue).
			Find()
		require.NoError(t, err)
		require.Len(t, results, 5)
		require.Equal(t, "PaginationTest5", results[0].(TestObject).Name)
		require.Equal(t, 5, results[0].(TestObject).Value)
	})

	t.Run("test date", func(t *testing.T) {
		results, err := store.Query().OrderBy(
			OrderByField("createdAt", true),
		).
			Limit(5).
			Find()
		require.NoError(t, err)
		require.Len(t, results, 5)
		require.Equal(t, "PaginationTest10", results[0].(TestObject).Name)
		require.Equal(t, 10, results[0].(TestObject).Value)

		lastValue := results[len(results)-1].(TestObject).CreatedAt
		results, err = store.Query().OrderBy(
			OrderByField("createdAt", true),
		).
			Limit(5).
			After(lastValue).
			Find()
		require.NoError(t, err)
		require.Len(t, results, 5)
		require.Equal(t, "PaginationTest5", results[0].(TestObject).Name)
		require.Equal(t, 5, results[0].(TestObject).Value)
	})
}

func TestPaginationMultipleKeys(t *testing.T, store DataStore) {
	base := time.Now()
	data := []TestObject{
		// Insert intentionally out of desired CreatedAt order
		{Name: "C3", Value: 30, CreatedAt: base.Add(1 * time.Second)}, // oldest
		{Name: "C1", Value: 30, CreatedAt: base.Add(3 * time.Second)}, // newest
		{Name: "C2", Value: 30, CreatedAt: base.Add(2 * time.Second)}, // middle

		{Name: "B2", Value: 20, CreatedAt: base.Add(2 * time.Second)},
		{Name: "B1", Value: 20, CreatedAt: base.Add(3 * time.Second)},
		{Name: "B3", Value: 20, CreatedAt: base.Add(1 * time.Second)},

		{Name: "A1", Value: 0, CreatedAt: base.Add(3 * time.Second)},
		{Name: "A2", Value: 0, CreatedAt: base.Add(2 * time.Second)},
		{Name: "A3", Value: 0, CreatedAt: base.Add(1 * time.Second)},

		{Name: "A4", Value: 0, CreatedAt: base.Add(-1 * time.Second)},
		{Name: "A5", Value: 0, CreatedAt: base.Add(-2 * time.Second)},
		{Name: "A6", Value: 0, CreatedAt: base.Add(-3 * time.Second)},
		{Name: "A7", Value: 0, CreatedAt: base.Add(-4 * time.Second)},
		{Name: "A8", Value: 0, CreatedAt: base.Add(-5 * time.Second)},
	}

	for _, obj := range data {
		require.NoError(t, store.Create(obj))
	}

	t.Run("composite key pagination", func(t *testing.T) {
		results, err := store.Query().OrderBy(
			OrderByField("Value", true),
			OrderByField("CreatedAt", true),
		).Limit(4).Find()
		require.NoError(t, err)
		require.Len(t, results, 4)

		require.Equal(t, "C1", results[0].(TestObject).Name)
		require.Equal(t, "C2", results[1].(TestObject).Name)
		require.Equal(t, "C3", results[2].(TestObject).Name)
		require.Equal(t, "B1", results[3].(TestObject).Name)

		last := results[len(results)-1].(TestObject)

		t.Run("after with time", func(t *testing.T) {
			results, err = store.Query().OrderBy(
				OrderByField("Value", true),
				OrderByField("CreatedAt", true),
			).After(last.Value, last.CreatedAt).Limit(4).Find()
			require.NoError(t, err)
			require.Len(t, results, 4)

			require.Equal(t, "B2", results[0].(TestObject).Name, "Pagination did not respect second sort key")
			require.Equal(t, "A2", results[3].(TestObject).Name)
		})

		t.Run("after with string time", func(t *testing.T) {
			results, err = store.Query().OrderBy(
				OrderByField("Value", true),
				OrderByField("CreatedAt", true),
			).After(last.Value, last.CreatedAt.Format(time.RFC3339Nano)).Limit(4).Find()
			require.NoError(t, err)
			require.Len(t, results, 4)

			require.Equal(t, "B2", results[0].(TestObject).Name, "Pagination did not respect second sort key")

			last = results[len(results)-1].(TestObject) // Update last to the new last result

		})

		t.Run("test null values", func(t *testing.T) {
			results, err = store.Query().OrderBy(
				OrderByField("Value", true),
				OrderByField("CreatedAt", true),
			).After(last.Value, last.CreatedAt.Format(time.RFC3339Nano)).Limit(4).Find()
			require.NoError(t, err)
			require.Len(t, results, 4)

			require.Equal(t, "A3", results[0].(TestObject).Name, "Pagination did not respect second sort key")
			last = results[len(results)-1].(TestObject) // Update last to the new last result
			require.Equal(t, "A6", last.Name, "Last result should be A6")
		})

		t.Run("test null second page", func(t *testing.T) {
			results, err = store.Query().OrderBy(
				OrderByField("Value", true),
				OrderByField("CreatedAt", true),
			).After(last.Value, last.CreatedAt.Format(time.RFC3339Nano)).Limit(4).Find()
			require.NoError(t, err)
			require.Len(t, results, 2)

			require.Equal(t, "A7", results[0].(TestObject).Name, "Pagination did not respect second sort key")
		})
	})
}

func TestPointerPaginationMultipleKeys(t *testing.T, store DataStore) {
	base := time.Now()
	data := []*TestObject{
		// Insert intentionally out of desired CreatedAt order
		{Name: "C3", Value: 30, CreatedAt: base.Add(1 * time.Second)}, // oldest
		{Name: "C1", Value: 30, CreatedAt: base.Add(3 * time.Second)}, // newest
		{Name: "C2", Value: 30, CreatedAt: base.Add(2 * time.Second)}, // middle

		{Name: "B2", Value: 20, CreatedAt: base.Add(2 * time.Second)},
		{Name: "B1", Value: 20, CreatedAt: base.Add(3 * time.Second)},
		{Name: "B3", Value: 20, CreatedAt: base.Add(1 * time.Second)},

		{Name: "A1", Value: 0, CreatedAt: base.Add(3 * time.Second)},
		{Name: "A2", Value: 0, CreatedAt: base.Add(2 * time.Second)},
		{Name: "A3", Value: 0, CreatedAt: base.Add(1 * time.Second)},

		{Name: "A4", Value: 0, CreatedAt: base.Add(-1 * time.Second)},
		{Name: "A5", Value: 0, CreatedAt: base.Add(-2 * time.Second)},
		{Name: "A6", Value: 0, CreatedAt: base.Add(-3 * time.Second)},
		{Name: "A7", Value: 0, CreatedAt: base.Add(-4 * time.Second)},
		{Name: "A8", Value: 0, CreatedAt: base.Add(-5 * time.Second)},
	}

	for _, obj := range data {
		require.NoError(t, store.Create(obj))
	}

	t.Run("composite key pagination", func(t *testing.T) {
		results, err := store.Query().OrderBy(
			OrderByField("Value", true),
			OrderByField("CreatedAt", true),
		).Limit(4).Find()
		require.NoError(t, err)
		require.Len(t, results, 4)

		require.Equal(t, "C1", results[0].(*TestObject).Name)
		require.Equal(t, "C2", results[1].(*TestObject).Name)
		require.Equal(t, "C3", results[2].(*TestObject).Name)
		require.Equal(t, "B1", results[3].(*TestObject).Name)

		last := results[len(results)-1].(*TestObject)

		t.Run("after with time", func(t *testing.T) {
			results, err = store.Query().OrderBy(
				OrderByField("Value", true),
				OrderByField("CreatedAt", true),
			).After(last.Value, last.CreatedAt).Limit(4).Find()
			require.NoError(t, err)
			require.Len(t, results, 4)

			require.Equal(t, "B2", results[0].(*TestObject).Name, "Pagination did not respect second sort key")
			require.Equal(t, "A2", results[3].(*TestObject).Name)
		})

		t.Run("after with string time", func(t *testing.T) {
			results, err = store.Query().OrderBy(
				OrderByField("Value", true),
				OrderByField("CreatedAt", true),
			).After(last.Value, last.CreatedAt.Format(time.RFC3339Nano)).Limit(4).Find()
			require.NoError(t, err)
			require.Len(t, results, 4)

			require.Equal(t, "B2", results[0].(*TestObject).Name, "Pagination did not respect second sort key")

			last = results[len(results)-1].(*TestObject) // Update last to the new last result

		})

		t.Run("test null values", func(t *testing.T) {
			results, err = store.Query().OrderBy(
				OrderByField("Value", true),
				OrderByField("CreatedAt", true),
			).After(last.Value, last.CreatedAt.Format(time.RFC3339Nano)).Limit(4).Find()
			require.NoError(t, err)
			require.Len(t, results, 4)

			require.Equal(t, "A3", results[0].(*TestObject).Name, "Pagination did not respect second sort key")
			last = results[len(results)-1].(*TestObject) // Update last to the new last result
			require.Equal(t, "A6", last.Name, "Last result should be A6")
		})

		t.Run("test null second page", func(t *testing.T) {
			results, err = store.Query().OrderBy(
				OrderByField("Value", true),
				OrderByField("CreatedAt", true),
			).After(last.Value, last.CreatedAt.Format(time.RFC3339Nano)).Limit(4).Find()
			require.NoError(t, err)
			require.Len(t, results, 2)

			require.Equal(t, "A7", results[0].(*TestObject).Name, "Pagination did not respect second sort key")
		})
	})
}

func TestPointerPagination(t *testing.T, store DataStore) {
	for i := 1; i <= 10; i++ {
		now := time.Now()
		obj := &TestObject{
			Name:      fmt.Sprintf("PaginationTest%d", i),
			Value:     i,
			CreatedAt: now,
		}
		err := store.Create(obj)
		time.Sleep(1 * time.Millisecond)
		require.NoError(t, err)
	}

	t.Run("test int", func(t *testing.T) {
		results, err := store.Query().OrderBy(
			OrderByField("Value", true),
		).
			Limit(5).
			Find()
		require.NoError(t, err)
		require.Len(t, results, 5)
		require.Equal(t, "PaginationTest10", results[0].(*TestObject).Name)
		require.Equal(t, 10, results[0].(*TestObject).Value)

		lastValue := results[len(results)-1].(*TestObject).Value
		results, err = store.Query().OrderBy(
			OrderByField("Value", true),
		).
			Limit(5).
			After(lastValue).
			Find()
		require.NoError(t, err)
		require.Len(t, results, 5)
		require.Equal(t, "PaginationTest5", results[0].(*TestObject).Name)
		require.Equal(t, 5, results[0].(*TestObject).Value)
	})

	t.Run("test date", func(t *testing.T) {
		results, err := store.Query().OrderBy(
			OrderByField("createdAt", true),
		).
			Limit(5).
			Find()
		require.NoError(t, err)
		require.Len(t, results, 5)
		require.Equal(t, "PaginationTest10", results[0].(*TestObject).Name)
		require.Equal(t, 10, results[0].(*TestObject).Value)

		lastValue := results[len(results)-1].(*TestObject).CreatedAt
		results, err = store.Query().OrderBy(
			OrderByField("createdAt", true),
		).
			Limit(5).
			After(lastValue).
			Find()
		require.NoError(t, err)
		require.Len(t, results, 5)
		require.Equal(t, "PaginationTest5", results[0].(*TestObject).Name)
		require.Equal(t, 5, results[0].(*TestObject).Value)
	})
}
