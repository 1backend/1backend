package distlock

import (
	"context"
	"errors"
	"sync"
	"time"
)

// LocalDistributedLock implements the DistributedLock interface using sync.Mutex for local locking.
type LocalDistributedLock struct {
	mutexes  map[string]*sync.Mutex
	heldKeys map[string]bool
	lock     sync.Mutex // Protects access to the map of mutexes and held keys
}

// NewLocalDistributedLock creates a new LocalDistributedLock.
func NewLocalDistributedLock() *LocalDistributedLock {
	return &LocalDistributedLock{
		mutexes:  make(map[string]*sync.Mutex),
		heldKeys: make(map[string]bool),
	}
}

// getOrCreateMutex retrieves the mutex for the given key, or creates a new one if it doesn't exist.
func (l *LocalDistributedLock) getOrCreateMutex(key string) *sync.Mutex {
	l.lock.Lock()
	defer l.lock.Unlock()

	if m, exists := l.mutexes[key]; exists {
		return m
	}

	m := &sync.Mutex{}
	l.mutexes[key] = m
	return m
}

// Acquire tries to acquire the lock for the specified key. It blocks until the lock is acquired or the context is done.
func (l *LocalDistributedLock) Acquire(ctx context.Context, key string) error {
	mutex := l.getOrCreateMutex(key)
	ticker := time.NewTicker(time.Millisecond)
	defer ticker.Stop()

	for {
		if mutex.TryLock() {
			l.lock.Lock()
			l.heldKeys[key] = true
			l.lock.Unlock()
			return nil
		}

		select {
		case <-ctx.Done():
			return errors.New("failed to acquire lock due to context cancellation")
		case <-ticker.C:
		}
	}
}

// TryAcquire tries to acquire the lock for the specified key without blocking. Returns true if successful, false otherwise.
func (l *LocalDistributedLock) TryAcquire(ctx context.Context, key string) (bool, error) {
	select {
	case <-ctx.Done():
		return false, errors.New("context canceled before lock attempt")
	default:
	}

	mutex := l.getOrCreateMutex(key)
	if mutex.TryLock() {
		l.lock.Lock()
		l.heldKeys[key] = true
		l.lock.Unlock()
		return true, nil
	}

	return false, nil
}

// Release releases the lock for the specified key.
func (l *LocalDistributedLock) Release(ctx context.Context, key string) error {
	l.lock.Lock()
	defer l.lock.Unlock()

	if !l.heldKeys[key] {
		return errors.New("lock not held for key")
	}

	// Unlock the mutex
	mutex := l.mutexes[key]
	mutex.Unlock()

	// Clean up the held keys and mutexes
	delete(l.heldKeys, key)
	return nil
}

// IsHeld returns whether the lock is currently held for the specified key by this instance.
func (l *LocalDistributedLock) IsHeld(key string) bool {
	l.lock.Lock()
	defer l.lock.Unlock()

	return l.heldKeys[key]
}
