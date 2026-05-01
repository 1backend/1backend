package infra

import (
	"database/sql"
	"fmt"
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
	})
	require.NoError(t, err)

	postgresFactory := factory.(*DataStoreFactoryPostgresImpl)
	t.Cleanup(func() {
		if postgresFactory.lockConn != nil {
			_ = postgresFactory.lockConn.Close()
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
