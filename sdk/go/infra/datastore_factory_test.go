package infra

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"strconv"
	"testing"
	"time"

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/testutil"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

type autoIndexFactoryObject struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (a autoIndexFactoryObject) GetId() string {
	return a.Id
}

func TestDataStoreFactoryPostgresEnablesAutoIndexing(t *testing.T) {
	conn := testutil.StartPostgres(t)
	tablePrefix := "factory_autoidx_" + strconv.FormatInt(time.Now().UnixNano(), 10) + "_"
	tableName := tablePrefix + "objects"

	factory, err := NewDataStoreFactory(DataStoreConfig{
		Db:                 "postgres",
		DbConnectionString: conn,
		TablePrefix:        tablePrefix,
		AutoIndexes:        true,
	})
	require.NoError(t, err)

	postgresFactory := factory.(*DataStoreFactoryPostgresImpl)
	t.Cleanup(func() {
		if postgresFactory.readDB != nil {
			_ = postgresFactory.readDB.Close()
		}
		if postgresFactory.db != nil {
			_ = postgresFactory.db.Close()
		}
	})

	store, err := factory.Create("objects", autoIndexFactoryObject{})
	require.NoError(t, err)
	require.NotNil(t, postgresFactory.options.Lock)

	require.NoError(t, store.Create(autoIndexFactoryObject{
		Id:   "row-1",
		Name: "Alice",
	}))

	filter := datastore.Equals(datastore.Field("Name"), "Alice")
	for range 3 {
		_, err := store.Query(filter).Find()
		require.NoError(t, err)
	}

	handle, err := factory.Handle()
	require.NoError(t, err)
	db := handle.(*sql.DB)

	require.Eventually(t, func() bool {
		var count int
		err := db.QueryRow(
			"SELECT count(*) FROM pg_indexes WHERE schemaname = ANY(current_schemas(false)) AND tablename = $1 AND indexname LIKE $2",
			tableName,
			fmt.Sprintf("%s_autoidx_%%", tableName),
		).Scan(&count)
		require.NoError(t, err)
		return count == 1
	}, 10*time.Second, 100*time.Millisecond)
}

func TestDataStoreFactoryPostgresAutoIndexingDefaultsOff(t *testing.T) {
	conn := testutil.StartPostgres(t)
	tablePrefix := "factory_no_autoidx_" + strconv.FormatInt(time.Now().UnixNano(), 10) + "_"
	tableName := tablePrefix + "objects"

	factory, err := NewDataStoreFactory(DataStoreConfig{
		Db:                 "postgres",
		DbConnectionString: conn,
		TablePrefix:        tablePrefix,
	})
	require.NoError(t, err)

	postgresFactory := factory.(*DataStoreFactoryPostgresImpl)
	t.Cleanup(func() {
		if postgresFactory.readDB != nil {
			_ = postgresFactory.readDB.Close()
		}
		if postgresFactory.db != nil {
			_ = postgresFactory.db.Close()
		}
	})

	store, err := factory.Create("objects", autoIndexFactoryObject{})
	require.NoError(t, err)

	require.NoError(t, store.Create(autoIndexFactoryObject{
		Id:   "row-1",
		Name: "Alice",
	}))

	filter := datastore.Equals(datastore.Field("Name"), "Alice")
	for range 3 {
		_, err := store.Query(filter).Find()
		require.NoError(t, err)
	}

	handle, err := factory.Handle()
	require.NoError(t, err)
	db := handle.(*sql.DB)

	var count int
	err = db.QueryRow(
		"SELECT count(*) FROM pg_indexes WHERE schemaname = ANY(current_schemas(false)) AND tablename = $1 AND indexname LIKE $2",
		tableName,
		fmt.Sprintf("%s_autoidx_%%", tableName),
	).Scan(&count)
	require.NoError(t, err)
	require.Zero(t, count)
}

func TestNewDataStoreFactoryLoadsReadConnectionStringFromEnv(t *testing.T) {
	t.Setenv("OB_DB", "postgres")
	t.Setenv("OB_DB_CONNECTION_STRING", "postgres://writer")
	t.Setenv("OB_DB_READ_CONNECTION_STRING", "postgres://reader")

	factory, err := NewDataStoreFactory(DataStoreConfig{
		HomeDir: t.TempDir(),
	})
	require.NoError(t, err)

	postgresFactory := factory.(*DataStoreFactoryPostgresImpl)
	require.Equal(t, "postgres://writer", postgresFactory.options.DbConnectionString)
	require.Equal(t, "postgres://reader", postgresFactory.options.ReadDbConnectionString)
}

func TestNewDataStoreFactoryAppliesPostgresPoolDefaults(t *testing.T) {
	factory, err := NewDataStoreFactory(DataStoreConfig{
		HomeDir:            t.TempDir(),
		Db:                 "postgres",
		DbConnectionString: "postgres://writer",
	})
	require.NoError(t, err)

	postgresFactory := factory.(*DataStoreFactoryPostgresImpl)
	t.Cleanup(func() {
		if postgresFactory.db != nil {
			_ = postgresFactory.db.Close()
		}
	})

	handle, err := factory.Handle()
	require.NoError(t, err)

	db := handle.(*sql.DB)
	require.Equal(t, defaultPostgresMaxOpenConns, db.Stats().MaxOpenConnections)
	require.Equal(t, defaultPostgresMaxOpenConns, postgresFactory.options.DbPool.MaxOpenConns)
	require.Equal(t, defaultPostgresMaxIdleConns, postgresFactory.options.DbPool.MaxIdleConns)
	require.Equal(t, defaultPostgresConnMaxLifetime, postgresFactory.options.DbPool.ConnMaxLifetime)
	require.Equal(t, defaultPostgresConnMaxIdleTime, postgresFactory.options.DbPool.ConnMaxIdleTime)
}

func TestNewDataStoreFactoryLoadsPostgresPoolFromEnv(t *testing.T) {
	t.Setenv("OB_DB", "postgres")
	t.Setenv("OB_DB_CONNECTION_STRING", "postgres://writer")
	t.Setenv("OB_DB_MAX_OPEN_CONNS", "7")
	t.Setenv("OB_DB_MAX_IDLE_CONNS", "3")
	t.Setenv("OB_DB_CONN_MAX_LIFETIME", "11m")
	t.Setenv("OB_DB_CONN_MAX_IDLE_TIME", "90s")
	t.Setenv("OB_DB_APPLICATION_NAME", "onebackend-prod")

	factory, err := NewDataStoreFactory(DataStoreConfig{
		HomeDir: t.TempDir(),
	})
	require.NoError(t, err)

	postgresFactory := factory.(*DataStoreFactoryPostgresImpl)
	t.Cleanup(func() {
		if postgresFactory.db != nil {
			_ = postgresFactory.db.Close()
		}
	})

	handle, err := factory.Handle()
	require.NoError(t, err)

	db := handle.(*sql.DB)
	require.Equal(t, 7, db.Stats().MaxOpenConnections)
	require.Equal(t, 7, postgresFactory.options.DbPool.MaxOpenConns)
	require.Equal(t, 3, postgresFactory.options.DbPool.MaxIdleConns)
	require.Equal(t, 11*time.Minute, postgresFactory.options.DbPool.ConnMaxLifetime)
	require.Equal(t, 90*time.Second, postgresFactory.options.DbPool.ConnMaxIdleTime)
	require.Equal(t, "onebackend-prod", postgresFactory.options.DbApplicationName)
}

func TestNewDataStoreFactoryRejectsTooSmallPostgresPool(t *testing.T) {
	_, err := NewDataStoreFactory(DataStoreConfig{
		HomeDir:            t.TempDir(),
		Db:                 "postgres",
		DbConnectionString: "postgres://writer",
		DbPool: DbPoolConfig{
			MaxOpenConns: 1,
		},
	})

	require.Error(t, err)
	require.Contains(t, err.Error(), "at least 2")
}

func TestConnectionStringForRoleAddsPostgresApplicationName(t *testing.T) {
	cfg := DataStoreConfig{
		Db:                "postgres",
		DbApplicationName: "onebackend-prod",
	}

	got := cfg.connectionStringForRole(
		"write",
		"postgres://user:password@example.com:5432/app?sslmode=disable",
	)

	parsed, err := url.Parse(got)
	require.NoError(t, err)
	require.Equal(t, "disable", parsed.Query().Get("sslmode"))
	require.Equal(t, "onebackend-prod:write", parsed.Query().Get("application_name"))
}

func TestConnectionStringForRoleLabelsKeywordValuePostgresDSN(t *testing.T) {
	cfg := DataStoreConfig{
		Db:                "postgres",
		DbApplicationName: "onebackend prod",
	}

	got := cfg.connectionStringForRole(
		"read",
		"host=example.com user=onebackend dbname=app sslmode=disable",
	)

	require.Contains(t, got, "application_name='onebackend-prod:read'")
}

func TestCheckDataStoreReadyAcceptsWritablePostgres(t *testing.T) {
	conn := testutil.StartPostgres(t)

	factory, err := NewDataStoreFactory(DataStoreConfig{
		HomeDir:            t.TempDir(),
		Db:                 "postgres",
		DbConnectionString: conn,
	})
	require.NoError(t, err)

	postgresFactory := factory.(*DataStoreFactoryPostgresImpl)
	t.Cleanup(func() {
		if postgresFactory.db != nil {
			_ = postgresFactory.db.Close()
		}
	})

	require.NoError(t, CheckDataStoreReady(context.Background(), factory))
}

func TestCheckDataStoreReadyRejectsReadOnlyPostgres(t *testing.T) {
	conn := testutil.StartPostgres(t)

	factory, err := NewDataStoreFactory(DataStoreConfig{
		HomeDir:            t.TempDir(),
		Db:                 "postgres",
		DbConnectionString: connWithQueryValue(t, conn, "default_transaction_read_only", "on"),
	})
	require.NoError(t, err)

	postgresFactory := factory.(*DataStoreFactoryPostgresImpl)
	t.Cleanup(func() {
		if postgresFactory.db != nil {
			_ = postgresFactory.db.Close()
		}
	})

	err = CheckDataStoreReady(context.Background(), factory)
	require.Error(t, err)
	require.Contains(t, err.Error(), "read-only")
}

func connWithQueryValue(t *testing.T, conn string, key string, value string) string {
	t.Helper()

	parsed, err := url.Parse(conn)
	require.NoError(t, err)

	query := parsed.Query()
	query.Set(key, value)
	parsed.RawQuery = query.Encode()
	return parsed.String()
}
