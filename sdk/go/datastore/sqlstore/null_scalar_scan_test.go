package sqlstore

import (
	"database/sql"
	"fmt"
	"strings"
	"testing"

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

type nullScalarScanTestObject struct {
	Id           string                `json:"id"`
	StringValue  string                `json:"stringValue,omitempty"`
	NamedValue   datastore.NamedString `json:"namedValue,omitempty"`
	BoolValue    bool                  `json:"boolValue,omitempty"`
	IntValue     int                   `json:"intValue,omitempty"`
	Int8Value    int8                  `json:"int8Value,omitempty"`
	UintValue    uint                  `json:"uintValue,omitempty"`
	Uint8Value   uint8                 `json:"uint8Value,omitempty"`
	FloatValue   float64               `json:"floatValue,omitempty"`
	Float32Value float32               `json:"float32Value,omitempty"`
}

func (t nullScalarScanTestObject) GetId() string {
	return t.Id
}

func TestFindTreatsNullScalarsAsZeroValues(t *testing.T) {
	defer func() {
		if recovered := recover(); recovered != nil {
			t.Skipf("docker-backed sqlstore unavailable: %v", recovered)
		}
	}()

	pgConn := postgresConnString(t)

	db, err := sql.Open("postgres", pgConn)
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, db.Close())
	})

	table := "table_" + strings.ReplaceAll(uuid.New().String(), "-", "")[:10]

	store, err := NewSQLStore(
		nullScalarScanTestObject{},
		DriverPostGRES,
		db,
		table,
		true,
	)
	require.NoError(t, err)

	obj := nullScalarScanTestObject{
		Id:           "null-scalars",
		StringValue:  "tracked",
		NamedValue:   datastore.NamedStringTwo,
		BoolValue:    true,
		IntValue:     42,
		Int8Value:    7,
		UintValue:    99,
		Uint8Value:   5,
		FloatValue:   3.5,
		Float32Value: 1.25,
	}

	require.NoError(t, store.Upsert(obj))

	_, err = db.Exec(fmt.Sprintf(
		"UPDATE %s SET stringValue = NULL, namedValue = NULL, boolValue = NULL, intValue = NULL, int8Value = NULL, uintValue = NULL, uint8Value = NULL, floatValue = NULL, float32Value = NULL WHERE id = $1",
		table,
	), obj.Id)
	require.NoError(t, err)

	rows, err := store.Query(datastore.Equals(datastore.Field("Id"), obj.Id)).Find()
	require.NoError(t, err)
	require.Len(t, rows, 1)

	stored := rows[0].(nullScalarScanTestObject)
	require.Equal(t, obj.Id, stored.Id)
	require.Zero(t, stored.StringValue)
	require.Zero(t, stored.NamedValue)
	require.Zero(t, stored.BoolValue)
	require.Zero(t, stored.IntValue)
	require.Zero(t, stored.Int8Value)
	require.Zero(t, stored.UintValue)
	require.Zero(t, stored.Uint8Value)
	require.Zero(t, stored.FloatValue)
	require.Zero(t, stored.Float32Value)
}
