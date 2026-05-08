package distlock

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/binary"
	"fmt"
	"strings"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

// PGDistributedLock implements DistributedLock using PostgreSQL advisory locks
type PGDistributedLock struct {
	conn       *sql.Conn
	connFunc   func(context.Context) (*sql.Conn, error)
	heldKeys   map[string]bool
	mutexes    map[string]*sync.Mutex
	stateMutex sync.Mutex
	connMutex  sync.Mutex
}

// NewPGDistributedLock creates a new instance of PGDistributedLock
func NewPGDistributedLock(conn *sql.Conn) *PGDistributedLock {
	return &PGDistributedLock{
		conn:     conn,
		heldKeys: map[string]bool{},
		mutexes:  map[string]*sync.Mutex{},
	}
}

// NewPGDistributedLockFromDB creates a PostgreSQL advisory lock that can
// reacquire its dedicated session if PostgreSQL restarts or the writer endpoint
// moves. Advisory locks are bound to a server session, so the implementation
// still uses one checked-out connection while healthy.
func NewPGDistributedLockFromDB(db *sql.DB) *PGDistributedLock {
	return &PGDistributedLock{
		connFunc: db.Conn,
		heldKeys: map[string]bool{},
		mutexes:  map[string]*sync.Mutex{},
	}
}

// hashKey hashes the key string into a 64-bit integer for advisory locking
func hashKey(key string) int64 {
	hash := sha256.Sum256([]byte(key))
	return int64(binary.BigEndian.Uint64(hash[:8]))
}

func (l *PGDistributedLock) getOrCreateMutex(key string) *sync.Mutex {
	l.stateMutex.Lock()
	defer l.stateMutex.Unlock()

	if m, exists := l.mutexes[key]; exists {
		return m
	}

	m := &sync.Mutex{}
	l.mutexes[key] = m
	return m
}

func (l *PGDistributedLock) acquireLocalMutex(ctx context.Context, key string) (*sync.Mutex, error) {
	mutex := l.getOrCreateMutex(key)
	ticker := time.NewTicker(time.Millisecond)
	defer ticker.Stop()

	for {
		if mutex.TryLock() {
			return mutex, nil
		}

		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("failed to acquire local lock for key '%s': %w", key, ctx.Err())
		case <-ticker.C:
		}
	}
}

func (l *PGDistributedLock) ensureConnLocked(ctx context.Context) error {
	if l.conn != nil {
		return nil
	}
	if l.connFunc == nil {
		return fmt.Errorf("postgres lock connection is not available")
	}

	conn, err := l.connFunc(ctx)
	if err != nil {
		return err
	}
	l.conn = conn
	return nil
}

func (l *PGDistributedLock) reconnectLocked(ctx context.Context) error {
	if l.connFunc == nil {
		return fmt.Errorf("postgres lock connection cannot be refreshed")
	}
	if l.conn != nil {
		_ = l.conn.Close()
		l.conn = nil
	}
	return l.ensureConnLocked(ctx)
}

func recoverableConnErr(err error) bool {
	if err == nil {
		return false
	}
	msg := strings.ToLower(err.Error())
	for _, fragment := range []string{
		"bad connection",
		"connection is already closed",
		"connection refused",
		"connection reset",
		"broken pipe",
		"server closed the connection",
		"eof",
	} {
		if strings.Contains(msg, fragment) {
			return true
		}
	}
	return false
}

func (l *PGDistributedLock) queryLockLocked(ctx context.Context, query string, lockKey int64, out *bool) error {
	if err := l.ensureConnLocked(ctx); err != nil {
		return err
	}

	queryCtx := ctx
	if l.connFunc == nil {
		queryCtx = context.Background()
	}

	err := l.conn.QueryRowContext(queryCtx, query, lockKey).Scan(out)
	if err == nil || !recoverableConnErr(err) {
		return err
	}

	if reconnectErr := l.reconnectLocked(ctx); reconnectErr != nil {
		return fmt.Errorf("%w; reconnect failed: %v", err, reconnectErr)
	}
	return l.conn.QueryRowContext(queryCtx, query, lockKey).Scan(out)
}

// Acquire tries to acquire the PostgreSQL advisory lock for the given key.
func (l *PGDistributedLock) Acquire(ctx context.Context, key string) error {
	localMutex, err := l.acquireLocalMutex(ctx, key)
	if err != nil {
		return err
	}

	lockKey := hashKey(key)
	query := "SELECT pg_try_advisory_lock($1)"
	ticker := time.NewTicker(time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			localMutex.Unlock()
			return fmt.Errorf("failed to acquire lock for key '%s': %w", key, ctx.Err())
		default:
		}

		var success bool
		l.connMutex.Lock()
		err = l.queryLockLocked(ctx, query, lockKey, &success)
		l.connMutex.Unlock()
		if err != nil {
			localMutex.Unlock()
			return fmt.Errorf("failed to acquire lock for key '%s': %w", key, err)
		}
		if success {
			l.stateMutex.Lock()
			l.heldKeys[key] = true
			l.stateMutex.Unlock()

			return nil
		}

		select {
		case <-ctx.Done():
			localMutex.Unlock()
			return fmt.Errorf("failed to acquire lock for key '%s': %w", key, ctx.Err())
		case <-ticker.C:
		}
	}
}

// TryAcquire tries to acquire the lock for the given key without blocking.
func (l *PGDistributedLock) TryAcquire(ctx context.Context, key string) (bool, error) {
	select {
	case <-ctx.Done():
		return false, fmt.Errorf("context canceled before lock attempt for key '%s': %w", key, ctx.Err())
	default:
	}

	localMutex := l.getOrCreateMutex(key)
	if !localMutex.TryLock() {
		return false, nil
	}

	lockKey := hashKey(key)

	query := "SELECT pg_try_advisory_lock($1)"

	var success bool
	l.connMutex.Lock()
	err := l.queryLockLocked(ctx, query, lockKey, &success)
	l.connMutex.Unlock()
	if err != nil {
		localMutex.Unlock()
		return false, fmt.Errorf("failed to try acquire lock for key '%s': %w", key, err)
	}

	if success {
		l.stateMutex.Lock()
		l.heldKeys[key] = true
		l.stateMutex.Unlock()
		return true, nil
	}

	localMutex.Unlock()
	return false, nil
}

// Release releases the PostgreSQL advisory lock for the given key.
func (l *PGDistributedLock) Release(ctx context.Context, key string) error {
	localMutex := l.getOrCreateMutex(key)

	l.stateMutex.Lock()
	if !l.heldKeys[key] {
		l.stateMutex.Unlock()
		return fmt.Errorf("lock not held for key '%s'", key)
	}
	l.stateMutex.Unlock()

	lockKey := hashKey(key)

	query := "SELECT pg_advisory_unlock($1)"

	var unlocked bool
	l.connMutex.Lock()
	err := l.queryLockLocked(context.Background(), query, lockKey, &unlocked)
	l.connMutex.Unlock()
	if err != nil {
		if recoverableConnErr(err) {
			l.stateMutex.Lock()
			delete(l.heldKeys, key)
			l.stateMutex.Unlock()
			localMutex.Unlock()
		}
		return fmt.Errorf("failed to release lock for key '%s': %w", key, err)
	}
	if !unlocked {
		l.stateMutex.Lock()
		delete(l.heldKeys, key)
		l.stateMutex.Unlock()
		localMutex.Unlock()
		return nil
	}

	l.stateMutex.Lock()
	delete(l.heldKeys, key)
	l.stateMutex.Unlock()
	localMutex.Unlock()

	return nil
}

// IsHeld returns whether the lock for the given key is held by this instance.
// Note: PostgreSQL advisory locks do not provide a direct way to check if the lock is held,
// so this method may have limited utility without additional state management.
func (l *PGDistributedLock) IsHeld(key string) bool {
	l.stateMutex.Lock()
	defer l.stateMutex.Unlock()

	return l.heldKeys[key]
}
