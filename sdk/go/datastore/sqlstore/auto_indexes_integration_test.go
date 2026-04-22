package sqlstore

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/1backend/1backend/sdk/go/datastore"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
	tc "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

type autoIndexCapObject struct {
	Id string `json:"id"`
	F1 string `json:"f1"`
	F2 string `json:"f2"`
	F3 string `json:"f3"`
	F4 string `json:"f4"`
	F5 string `json:"f5"`
	F6 string `json:"f6"`
	F7 string `json:"f7"`
	F8 string `json:"f8"`
	F9 string `json:"f9"`
}

func (a autoIndexCapObject) GetId() string {
	return a.Id
}

func TestSQLStoreAutoIndexStatsTracksAllExecutions(t *testing.T) {
	store, db := newPostgresStore(t, datastore.TestObject{}, "auto_idx_execs")
	defer store.Close()
	defer db.Close()

	filter := datastore.Equals(datastore.Field("Name"), "missing")

	_, err := store.Query(filter).Find()
	require.NoError(t, err)

	_, _, err = store.Query(filter).FindOne()
	require.NoError(t, err)

	_, err = store.Query(filter).Count()
	require.NoError(t, err)

	err = store.Query(filter).UpdateFields(map[string]any{"Age": 7})
	require.NoError(t, err)

	err = store.Query(filter).Delete()
	require.NoError(t, err)

	stats := waitForShapeHits(t, store, 5)
	require.True(t, stats.Supported)
	require.Len(t, stats.Shapes, 1)
	require.Equal(t, 5, stats.Shapes[0].Hits)
}

func TestSQLStoreAutoIndexesManualAndObserved(t *testing.T) {
	store, db := newPostgresStore(t, datastore.TestObject{}, "auto_idx_manual")
	defer store.Close()
	defer db.Close()

	require.NoError(t, store.Create(datastore.TestObject{Name: "Alice", Value: 10}))
	require.True(t, hasIndexNamed(t, db, "auto_idx_manual", "auto_idx_manual_value_idx"))

	filter := datastore.Equals(datastore.Field("Name"), "Alice")
	_, err := store.Query(filter).Find()
	require.NoError(t, err)
	_, err = store.Query(filter).Count()
	require.NoError(t, err)

	require.Zero(t, autoIndexCount(t, db, "auto_idx_manual"))

	_, err = store.Query(filter).Find()
	require.NoError(t, err)

	stats := waitForAutoIndexes(t, store, 1)
	autoEntries := autoEntries(stats)
	require.Len(t, autoEntries, 1)
	require.Equal(t, datastore.AutoIndexStatusCreated, autoEntries[0].Status)
	require.Equal(t, 1, autoIndexCount(t, db, "auto_idx_manual"))
}

func TestSQLStoreAutoIndexesPathArrayAndTrigram(t *testing.T) {
	t.Run("path expression", func(t *testing.T) {
		store, db := newPostgresStore(t, datastore.TestObject{}, "auto_idx_path")
		defer store.Close()
		defer db.Close()

		require.NoError(t, store.Create(datastore.TestObject{
			Name:   "path-row",
			Friend: datastore.Friend{Age: 30},
		}))

		filter := datastore.LessThan(datastore.Field("Friend.Age"), 40)
		for range 3 {
			_, err := store.Query(filter).Find()
			require.NoError(t, err)
		}

		waitForAutoIndexes(t, store, 1)
		def := autoIndexDefinitions(t, db, "auto_idx_path")
		require.Len(t, def, 1)
		require.Contains(t, def[0], "friend")
		require.Contains(t, def[0], "numeric")
	})

	t.Run("array gin", func(t *testing.T) {
		store, db := newPostgresStore(t, datastore.TestObject{}, "auto_idx_array")
		defer store.Close()
		defer db.Close()

		require.NoError(t, store.Create(datastore.TestObject{
			Name:      "array-row",
			NickNames: []string{"A1", "A2"},
		}))

		filter := datastore.Intersects(datastore.Field("NickNames"), []any{"A1"})
		for range 3 {
			_, err := store.Query(filter).Find()
			require.NoError(t, err)
		}

		waitForAutoIndexes(t, store, 1)
		def := autoIndexDefinitions(t, db, "auto_idx_array")
		require.Len(t, def, 1)
		require.Contains(t, strings.ToLower(def[0]), "using gin")
	})

	t.Run("trigram", func(t *testing.T) {
		store, db := newPostgresStore(t, datastore.TestObject{}, "auto_idx_trgm")
		defer store.Close()
		defer db.Close()

		require.NoError(t, store.Create(datastore.TestObject{Name: "Alice"}))

		filter := datastore.ContainsSubstring(datastore.Field("Name"), "li")
		for range 3 {
			_, err := store.Query(filter).Find()
			require.NoError(t, err)
		}

		waitForAutoIndexes(t, store, 1)
		def := autoIndexDefinitions(t, db, "auto_idx_trgm")
		require.Len(t, def, 1)
		require.Contains(t, def[0], "gin_trgm_ops")
		require.True(t, pgTrgmEnabled(t, db))
	})
}

func TestSQLStoreAutoIndexesRediscoverAfterRestart(t *testing.T) {
	conn := postgresConnString(t)
	if conn == "" {
		return
	}
	db, err := sql.Open("postgres", conn)
	require.NoError(t, err)

	store, err := NewSQLStore(datastore.TestObject{}, DriverPostGRES, db, "auto_idx_restart", false)
	require.NoError(t, err)

	require.NoError(t, store.Create(datastore.TestObject{Name: "Alice"}))
	filter := datastore.Equals(datastore.Field("Name"), "Alice")
	for range 3 {
		_, err := store.Query(filter).Find()
		require.NoError(t, err)
	}

	waitForAutoIndexes(t, store, 1)
	require.NoError(t, store.Close())
	require.NoError(t, db.Close())

	db, err = sql.Open("postgres", conn)
	require.NoError(t, err)
	defer db.Close()

	store, err = NewSQLStore(datastore.TestObject{}, DriverPostGRES, db, "auto_idx_restart", false)
	require.NoError(t, err)
	defer store.Close()

	stats := store.AutoIndexStats()
	foundDiscovered := false
	for _, entry := range autoEntries(stats) {
		if entry.Status == datastore.AutoIndexStatusDiscovered {
			foundDiscovered = true
			break
		}
	}
	require.True(t, foundDiscovered)
}

func TestSQLStoreAutoIndexesCapStopsFurtherCreation(t *testing.T) {
	store, db := newPostgresStore(t, autoIndexCapObject{}, "auto_idx_cap")
	defer store.Close()
	defer db.Close()

	require.NoError(t, store.Create(autoIndexCapObject{
		Id: "row-1",
		F1: "v", F2: "v", F3: "v", F4: "v", F5: "v", F6: "v", F7: "v", F8: "v", F9: "v",
	}))

	fields := []string{"f1", "f2", "f3", "f4", "f5", "f6", "f7", "f8", "f9"}
	for _, field := range fields {
		filter := datastore.Equals(datastore.Field(field), "v")
		for range 3 {
			_, err := store.Query(filter).Find()
			require.NoError(t, err)
		}
	}

	stats := waitForAutoIndexes(t, store, 8)
	require.Len(t, autoEntries(stats), 8)
	require.Equal(t, 8, autoIndexCount(t, db, "auto_idx_cap"))

	capReached := false
	for _, shape := range stats.Shapes {
		if shape.Reason == "auto index cap reached" {
			capReached = true
			break
		}
	}
	require.True(t, capReached)
}

func newPostgresStore(t *testing.T, instance any, tableName string) (*SQLStore, *sql.DB) {
	t.Helper()

	conn := postgresConnString(t)
	if conn == "" {
		return nil, nil
	}
	db, err := sql.Open("postgres", conn)
	require.NoError(t, err)

	store, err := NewSQLStore(instance, DriverPostGRES, db, tableName, false)
	require.NoError(t, err)
	return store, db
}

func postgresConnString(t *testing.T) (conn string) {
	t.Helper()

	defer func() {
		if recovered := recover(); recovered != nil {
			t.Skipf("docker-backed sqlstore unavailable: %v", recovered)
		}
	}()

	ctx := context.Background()

	container, err := postgres.Run(ctx,
		"postgres:16-alpine",
		postgres.WithDatabase("mydatabase"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("mysecretpassword"),
		tc.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(90*time.Second),
		),
	)
	if err != nil {
		t.Skipf("docker-backed sqlstore unavailable: %v", err)
	}

	t.Cleanup(func() {
		_ = container.Terminate(ctx)
	})

	host, err := container.Host(ctx)
	require.NoError(t, err)

	port, err := container.MappedPort(ctx, "5432/tcp")
	require.NoError(t, err)

	ready := false
	for i := 0; i < 30; i++ {
		exitCode, _, err := container.Exec(ctx, []string{"pg_isready", "-U", "postgres", "-d", "mydatabase"})
		if err == nil && exitCode == 0 {
			ready = true
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
	require.True(t, ready, "postgres container did not become ready")

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", "postgres", "mysecretpassword", host, port.Port(), "mydatabase")
}

func waitForShapeHits(t *testing.T, store *SQLStore, hits int) datastore.AutoIndexStats {
	t.Helper()

	var stats datastore.AutoIndexStats
	require.Eventually(t, func() bool {
		stats = store.AutoIndexStats()
		return len(stats.Shapes) == 1 && stats.Shapes[0].Hits >= hits
	}, 5*time.Second, 50*time.Millisecond)

	return stats
}

func waitForAutoIndexes(t *testing.T, store *SQLStore, count int) datastore.AutoIndexStats {
	t.Helper()

	var stats datastore.AutoIndexStats
	require.Eventually(t, func() bool {
		stats = store.AutoIndexStats()
		return len(autoEntries(stats)) >= count
	}, 10*time.Second, 100*time.Millisecond)

	return stats
}

func autoEntries(stats datastore.AutoIndexStats) []datastore.AutoIndexEntry {
	out := []datastore.AutoIndexEntry{}
	for _, entry := range stats.Indexes {
		if entry.Kind == datastore.AutoIndexKindAuto && (entry.Status == datastore.AutoIndexStatusCreated || entry.Status == datastore.AutoIndexStatusDiscovered || entry.Status == datastore.AutoIndexStatusPending) {
			out = append(out, entry)
		}
	}
	return out
}

func autoIndexCount(t *testing.T, db *sql.DB, tableName string) int {
	t.Helper()

	var count int
	err := db.QueryRow(
		"SELECT count(*) FROM pg_indexes WHERE schemaname = ANY(current_schemas(false)) AND tablename = $1 AND indexname LIKE $2",
		tableName,
		fmt.Sprintf("%s_autoidx_%%", tableName),
	).Scan(&count)
	require.NoError(t, err)
	return count
}

func autoIndexDefinitions(t *testing.T, db *sql.DB, tableName string) []string {
	t.Helper()

	rows, err := db.Query(
		"SELECT indexdef FROM pg_indexes WHERE schemaname = ANY(current_schemas(false)) AND tablename = $1 AND indexname LIKE $2 ORDER BY indexname",
		tableName,
		fmt.Sprintf("%s_autoidx_%%", tableName),
	)
	require.NoError(t, err)
	defer rows.Close()

	defs := []string{}
	for rows.Next() {
		var def string
		require.NoError(t, rows.Scan(&def))
		defs = append(defs, def)
	}
	require.NoError(t, rows.Err())
	return defs
}

func hasIndexNamed(t *testing.T, db *sql.DB, tableName, indexName string) bool {
	t.Helper()

	var count int
	err := db.QueryRow(
		"SELECT count(*) FROM pg_indexes WHERE schemaname = ANY(current_schemas(false)) AND tablename = $1 AND indexname = $2",
		tableName,
		indexName,
	).Scan(&count)
	require.NoError(t, err)
	return count == 1
}

func pgTrgmEnabled(t *testing.T, db *sql.DB) bool {
	t.Helper()

	var count int
	err := db.QueryRow("SELECT count(*) FROM pg_extension WHERE extname = 'pg_trgm'").Scan(&count)
	require.NoError(t, err)
	return count == 1
}
