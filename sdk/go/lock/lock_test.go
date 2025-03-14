//go:build dist
// +build dist

package lock

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	pglock "github.com/openorch/openorch/sdk/go/lock/pg"

	"github.com/stretchr/testify/require"
)

func TestLocks(t *testing.T) {
	lockStores := map[string]func(instance any) (DistributedLock, DistributedLock){
		"pgLock": func(instance any) (DistributedLock, DistributedLock) {
			// Use the same PostgreSQL connection string as in your existing tests
			db, err := sql.Open("postgres", "postgres://postgres:mysecretpassword@localhost:5432/mydatabase?sslmode=disable")
			require.NoError(t, err)
			conn, err := db.Conn(context.Background())
			require.NoError(t, err)

			db2, err := sql.Open("postgres", "postgres://postgres:mysecretpassword@localhost:5432/mydatabase?sslmode=disable")
			require.NoError(t, err)
			conn2, err := db2.Conn(context.Background())
			require.NoError(t, err)

			lockService := pglock.NewPGDistributedLock(conn)
			lockService2 := pglock.NewPGDistributedLock(conn2)

			return lockService, lockService2
		},
	}

	tests := map[string]func(t *testing.T, lock, lock2 DistributedLock){
		"AcquireRelease": LockAcquireRelease,
		"TryAcquire":     LockTryAcquire,
		"LockContention": LockContention,
	}

	for testName, test := range tests {
		for storeName, storeFunc := range lockStores {
			t.Run(fmt.Sprintf("%v %v", storeName, testName), func(t *testing.T) {
				lock, lock2 := storeFunc(nil)
				test(t, lock, lock2)
			})
		}
	}
}

func LockAcquireRelease(t *testing.T, lock, lock2 DistributedLock) {
	ctx := context.Background()
	key := "test_lock_acquire_release"

	err := lock.Acquire(ctx, key)
	require.NoError(t, err, "should acquire the lock without error")

	held := lock.IsHeld(key)
	require.True(t, held, "lock should be held after acquire")

	err = lock.Release(ctx, key)
	require.NoError(t, err, "should release the lock without error")

	held = lock.IsHeld(key)
	require.False(t, held, "lock should not be held after release")
}

func LockTryAcquire(t *testing.T, lock, lock2 DistributedLock) {
	ctx := context.Background()
	key := "test_lock_try_acquire"

	success, err := lock.TryAcquire(ctx, key)
	require.NoError(t, err, "should try acquire the lock without error")
	require.True(t, success, "should acquire the lock successfully")

	// second try
	success, err = lock.TryAcquire(ctx, key)
	require.NoError(t, err, "should try acquire the lock without error")
	require.True(t, success, "should acquire the lock successfully")

	success, err = lock2.TryAcquire(ctx, key)
	require.NoError(t, err, "should try acquire the lock without error")
	require.False(t, success, "should not acquire the lock a second time")

	err = lock.Release(ctx, key)
	require.NoError(t, err, "should release the lock without error")
}

func LockContention(t *testing.T, lock, lock2 DistributedLock) {
	ctx := context.Background()
	key := "test_lock_contention"

	err := lock.Acquire(ctx, key)
	require.NoError(t, err, "should acquire the lock without error")

	// this blocks here but not in the goroutine below?
	//err = lock.Acquire(ctx, key)
	//require.NoError(t, err, "should acquire the lock without error second time too")

	acquireResult := make(chan error, 1)
	go func() {
		err := lock2.Acquire(ctx, key)
		acquireResult <- err
	}()

	select {
	case <-acquireResult:
		t.Fatal("lock acquisition should be blocked by the first acquisition")
	case <-time.After(100 * time.Millisecond):
		// No response, as expected
	}

	err = lock.Release(ctx, key)
	require.NoError(t, err, "should release the lock without error")

	err = <-acquireResult
	require.NoError(t, err, "second acquisition should succeed after release")
}
