package lock

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	locallock "github.com/1backend/1backend/sdk/go/lock/local"
	pglock "github.com/1backend/1backend/sdk/go/lock/pg"
	"github.com/1backend/1backend/sdk/go/testutil"

	"github.com/stretchr/testify/require"
)

func TestLocks(t *testing.T) {
	pgConn := testutil.StartPostgres(t)

	lockStores := map[string]func(instance any) (DistributedLock, DistributedLock){
		"localLock": func(instance any) (DistributedLock, DistributedLock) {
			lockService := locallock.NewLocalDistributedLock()
			return lockService, lockService
		},
		"pgLock": func(instance any) (DistributedLock, DistributedLock) {
			// Use the same PostgreSQL connection string as in your existing tests
			db, err := sql.Open("postgres", pgConn)
			require.NoError(t, err)
			conn, err := db.Conn(context.Background())
			require.NoError(t, err)

			db2, err := sql.Open("postgres", pgConn)
			require.NoError(t, err)
			conn2, err := db2.Conn(context.Background())
			require.NoError(t, err)

			lockService := pglock.NewPGDistributedLock(conn)
			lockService2 := pglock.NewPGDistributedLock(conn2)

			return lockService, lockService2
		},
	}

	tests := map[string]func(t *testing.T, lock, lock2 DistributedLock){
		"AcquireCancellationDoesNotStealLock": LockAcquireCancellationDoesNotStealLock,
		"AcquireRelease":                      LockAcquireRelease,
		"ReleaseWithCanceledContext":          LockReleaseWithCanceledContext,
		"SameInstanceAcquireContention":       LockSameInstanceAcquireContention,
		"TryAcquire":                          LockTryAcquire,
		"LockContention":                      LockContention,
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

	success, err = lock.TryAcquire(ctx, key)
	require.NoError(t, err, "should try acquire the lock without error")
	require.False(t, success, "same lock instance should not reenter the lock")

	success, err = lock2.TryAcquire(ctx, key)
	require.NoError(t, err, "should try acquire the lock without error")
	require.False(t, success, "should not acquire the lock a second time")

	err = lock.Release(ctx, key)
	require.NoError(t, err, "should release the lock without error")

	success, err = lock2.TryAcquire(ctx, key)
	require.NoError(t, err, "should try acquire the lock after release")
	require.True(t, success, "should acquire the lock after release")

	err = lock2.Release(ctx, key)
	require.NoError(t, err, "should release the lock without error")
}

func LockContention(t *testing.T, lock, lock2 DistributedLock) {
	ctx := context.Background()
	key := "test_lock_contention"

	err := lock.Acquire(ctx, key)
	require.NoError(t, err, "should acquire the lock without error")

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

	err = lock2.Release(ctx, key)
	require.NoError(t, err, "should release the second acquisition")
}

func LockSameInstanceAcquireContention(t *testing.T, lock, lock2 DistributedLock) {
	ctx := context.Background()
	key := "test_lock_same_instance_contention"

	err := lock.Acquire(ctx, key)
	require.NoError(t, err, "should acquire the lock without error")

	acquireCtx, cancel := context.WithTimeout(context.Background(), 25*time.Millisecond)
	defer cancel()

	err = lock.Acquire(acquireCtx, key)
	require.Error(t, err, "same lock instance should not reenter the lock")

	err = lock.Release(ctx, key)
	require.NoError(t, err, "should release the lock without error")

	err = lock.Acquire(ctx, key)
	require.NoError(t, err, "same instance should acquire the lock after release")

	err = lock.Release(ctx, key)
	require.NoError(t, err, "should release the lock without error")
}

func LockAcquireCancellationDoesNotStealLock(t *testing.T, lock, lock2 DistributedLock) {
	ctx := context.Background()
	key := "test_lock_cancellation_does_not_steal"

	err := lock.Acquire(ctx, key)
	require.NoError(t, err, "should acquire the lock without error")

	acquireCtx, cancel := context.WithTimeout(context.Background(), 25*time.Millisecond)
	defer cancel()

	err = lock2.Acquire(acquireCtx, key)
	require.Error(t, err, "contending acquisition should stop when context is canceled")

	err = lock.Release(ctx, key)
	require.NoError(t, err, "should release the lock without error")

	time.Sleep(25 * time.Millisecond)

	retryCtx, retryCancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer retryCancel()

	err = lock2.Acquire(retryCtx, key)
	require.NoError(t, err, "canceled acquisition must not steal the lock after release")

	err = lock2.Release(ctx, key)
	require.NoError(t, err, "should release the retried acquisition")
}

func LockReleaseWithCanceledContext(t *testing.T, lock, lock2 DistributedLock) {
	ctx := context.Background()
	key := "test_lock_release_with_canceled_context"

	err := lock.Acquire(ctx, key)
	require.NoError(t, err, "should acquire the lock without error")

	releaseCtx, cancel := context.WithCancel(context.Background())
	cancel()

	err = lock.Release(releaseCtx, key)
	require.NoError(t, err, "release should not leave locks held when caller context is canceled")
	require.False(t, lock.IsHeld(key), "lock should not be held after release")

	err = lock2.Acquire(ctx, key)
	require.NoError(t, err, "another acquisition should succeed after canceled-context release")

	err = lock2.Release(ctx, key)
	require.NoError(t, err, "should release the second acquisition")
}
