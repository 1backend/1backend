package sqlstore

import (
	"database/sql"
	"net/url"
	"strings"
	"testing"

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

type readDBRouteTestObject struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (o readDBRouteTestObject) GetId() string {
	return o.Id
}

func TestReadDBRoutesNonTransactionalReads(t *testing.T) {
	conn := postgresConnString(t)

	adminDB, err := sql.Open("postgres", conn)
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, adminDB.Close())
	})

	suffix := strings.ReplaceAll(uuid.New().String(), "-", "")[:10]
	writerSchema := "writer_" + suffix
	readerSchema := "reader_" + suffix
	table := "read_route_" + suffix

	_, err = adminDB.Exec("CREATE SCHEMA " + writerSchema)
	require.NoError(t, err)
	_, err = adminDB.Exec("CREATE SCHEMA " + readerSchema)
	require.NoError(t, err)

	writerDB, err := sql.Open("postgres", connWithSearchPath(t, conn, writerSchema))
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, writerDB.Close())
	})

	readerDB, err := sql.Open("postgres", connWithSearchPath(t, conn, readerSchema))
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, readerDB.Close())
	})

	readerSetupStore, err := NewSQLStore(readDBRouteTestObject{}, DriverPostGRES, readerDB, table, false)
	require.NoError(t, err)
	require.NoError(t, readerSetupStore.Create(readDBRouteTestObject{Id: "reader-1", Name: "Reader One"}))
	require.NoError(t, readerSetupStore.Create(readDBRouteTestObject{Id: "reader-2", Name: "Reader Two"}))

	store, err := NewSQLStore(
		readDBRouteTestObject{},
		DriverPostGRES,
		writerDB,
		table,
		false,
		WithReadDB(readerDB),
	)
	require.NoError(t, err)
	require.NoError(t, store.Create(readDBRouteTestObject{Id: "writer-1", Name: "Writer One"}))

	rows, err := store.Query().Find()
	require.NoError(t, err)
	require.Len(t, rows, 2)
	require.ElementsMatch(t, []string{"reader-1", "reader-2"}, rowIDs(rows))

	count, err := store.Query().Count()
	require.NoError(t, err)
	require.Equal(t, int64(2), count)

	_, found, err := store.Query(datastore.Id("reader-1")).FindOne()
	require.NoError(t, err)
	require.True(t, found)

	_, found, err = store.Query(datastore.Id("writer-1")).FindOne()
	require.NoError(t, err)
	require.False(t, found)

	tx, err := store.BeginTransaction()
	require.NoError(t, err)
	require.NoError(t, tx.Create(readDBRouteTestObject{Id: "tx-1", Name: "Transaction"}))

	row, found, err := tx.Query(datastore.Id("tx-1")).FindOne()
	require.NoError(t, err)
	require.True(t, found)
	require.Equal(t, "tx-1", row.(readDBRouteTestObject).Id)
	require.NoError(t, tx.Rollback())
}

func connWithSearchPath(t *testing.T, conn, schema string) string {
	t.Helper()

	u, err := url.Parse(conn)
	require.NoError(t, err)

	q := u.Query()
	q.Set("options", "-c search_path="+schema)
	u.RawQuery = q.Encode()
	return u.String()
}

func rowIDs(rows []datastore.Row) []string {
	ids := make([]string, 0, len(rows))
	for _, row := range rows {
		ids = append(ids, row.(readDBRouteTestObject).Id)
	}
	return ids
}
