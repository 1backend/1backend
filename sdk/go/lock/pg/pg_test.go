package distlock

import (
	"context"
	"database/sql"
	"testing"

	"github.com/1backend/1backend/sdk/go/testutil"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func TestPGDistributedLockFromDBReconnectsClosedConnection(t *testing.T) {
	connString := testutil.StartPostgres(t)
	db, err := sql.Open("postgres", connString)
	require.NoError(t, err)
	t.Cleanup(func() {
		_ = db.Close()
	})

	lock := NewPGDistributedLockFromDB(db)
	closedConn, err := db.Conn(context.Background())
	require.NoError(t, err)
	require.NoError(t, closedConn.Close())

	lock.conn = closedConn

	require.NoError(t, lock.Acquire(context.Background(), "reconnects-closed-connection"))
	require.True(t, lock.IsHeld("reconnects-closed-connection"))
	require.NoError(t, lock.Release(context.Background(), "reconnects-closed-connection"))
	require.False(t, lock.IsHeld("reconnects-closed-connection"))
}

func TestPGDistributedLockReleaseClearsLocalStateAfterLostConnection(t *testing.T) {
	connString := testutil.StartPostgres(t)
	db, err := sql.Open("postgres", connString)
	require.NoError(t, err)
	t.Cleanup(func() {
		_ = db.Close()
	})

	lock := NewPGDistributedLockFromDB(db)
	key := "release-clears-lost-connection"

	require.NoError(t, lock.Acquire(context.Background(), key))

	lock.connMutex.Lock()
	require.NoError(t, lock.conn.Close())
	lock.connMutex.Unlock()

	require.NoError(t, lock.Release(context.Background(), key))
	require.False(t, lock.IsHeld(key))

	require.NoError(t, lock.Acquire(context.Background(), key))
	require.NoError(t, lock.Release(context.Background(), key))
}
