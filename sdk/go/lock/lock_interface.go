package lock

import "context"

// DistributedLock represents the interface for acquiring and releasing distributed locks
type DistributedLock interface {
	// Acquire tries to acquire the lock for the specified key. It blocks until the lock is acquired or the context is done.
	// Returns an error if the lock couldn't be acquired (due to context cancellation or another error).
	Acquire(ctx context.Context, key string) error

	// TryAcquire tries to acquire the lock for the specified key without blocking.
	// It immediately returns true if successful, false otherwise.
	TryAcquire(ctx context.Context, key string) (bool, error)

	// Release releases the lock for the specified key. It should return an error if the lock wasn't held or couldn't be released.
	Release(ctx context.Context, key string) error

	// IsHeld returns whether the lock is currently held for the specified key by this instance.
	IsHeld(key string) bool
}
