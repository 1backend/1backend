package sqlstore

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type Schema1 struct {
	Id        string `json:"id"`
	Namespace string `json:"namespace"`
}

type Schema2 struct {
	InternalId int64  `json:"internalId,omitempty"`
	Id         string `json:"id"`
	Name       string `json:"name"`
}

func TestSchema1PK(t *testing.T) {
	s := &SQLStore{}
	db := &DebugDB{
		skipExec: true,
		debug:    true,
	}
	_, err := s.createTable(Schema1{}, db, "schema1")
	require.NoError(t, err)

	require.Equal(t,
		"CREATE TABLE IF NOT EXISTS schema1 (); ALTER TABLE schema1 ADD COLUMN IF NOT EXISTS id TEXT PRIMARY KEY; ALTER TABLE schema1 ADD COLUMN IF NOT EXISTS namespace TEXT;",
		db.queries[0],
	)
}

func TestSchema2PK(t *testing.T) {
	s := &SQLStore{}
	db := &DebugDB{
		skipExec: true,
		debug:    true,
	}
	_, err := s.createTable(Schema2{}, db, "schema2")
	require.NoError(t, err)

	require.Equal(t,
		"CREATE TABLE IF NOT EXISTS schema2 (); ALTER TABLE schema2 ADD COLUMN IF NOT EXISTS internalId TEXT PRIMARY KEY; ALTER TABLE schema2 ADD COLUMN IF NOT EXISTS id TEXT; ALTER TABLE schema2 ADD COLUMN IF NOT EXISTS name TEXT;",
		db.queries[0],
	)
}
