package sqlstore

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type Schema1 struct {
	Id        string `json:"id"`
	Namespace string `json:"namespace"`
}

func TestSchema(t *testing.T) {
	s := &SQLStore{}
	db := &DebugDB{
		skipExec: true,
		debug:    true,
	}
	_, err := s.createTable(Schema1{}, db, "schema")
	require.NoError(t, err)
	require.Equal(t, "CREATE TABLE IF NOT EXISTS schema (); ALTER TABLE schema ADD COLUMN IF NOT EXISTS id TEXT; ALTER TABLE schema ADD COLUMN IF NOT EXISTS namespace TEXT;", db.queries[0])
}
