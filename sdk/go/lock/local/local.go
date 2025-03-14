package distlock

import (
	"context"
	"errors"
	"sync"
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

	// Try to lock the mutex or wait for the context to be done
	done := make(chan struct{})
	go func() {
		mutex.Lock()
		close(done)
	}()

	select {
	case <-done:
		l.lock.Lock()
		l.heldKeys[key] = true
		l.lock.Unlock()
		return nil
	case <-ctx.Done():
		return errors.New("failed to acquire lock due to context cancellation")
	}
}

// TryAcquire tries to acquire the lock for the specified key without blocking. Returns true if successful, false otherwise.
func (l *LocalDistributedLock) TryAcquire(ctx context.Context, key string) (bool, error) {
	mutex := l.getOrCreateMutex(key)

	// Try to lock the mutex without blocking
	locked := make(chan bool)
	go func() {
		locked <- mutex.TryLock()
	}()

	select {
	case success := <-locked:
		if success {
			l.lock.Lock()
			l.heldKeys[key] = true
			l.lock.Unlock()
		}
		return success, nil
	case <-ctx.Done():
		return false, errors.New("context canceled before lock attempt")
	}
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
