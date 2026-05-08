package infra

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

// CheckDataStoreReady verifies that the datastore's primary/write handle can
// serve work. PostgreSQL checks also reject read-only sessions, which catches
// stale primary connections after failover or restart.
func CheckDataStoreReady(ctx context.Context, factory DataStoreFactory) error {
	if factory == nil {
		return fmt.Errorf("datastore factory missing")
	}

	if checker, ok := factory.(interface {
		CheckReady(context.Context) error
	}); ok {
		return checker.CheckReady(ctx)
	}

	handle, err := factory.Handle()
	if err != nil {
		return err
	}
	if handle == nil {
		return nil
	}

	db, ok := handle.(*sql.DB)
	if !ok || db == nil {
		return fmt.Errorf("datastore handle is %T, not *sql.DB", handle)
	}
	return CheckSQLDBReady(ctx, db, "")
}

// CheckReady verifies the PostgreSQL writer pool managed by this factory.
func (df *DataStoreFactoryPostgresImpl) CheckReady(ctx context.Context) error {
	df.mutex.Lock()
	if err := df.ensureDBLocked(); err != nil {
		df.mutex.Unlock()
		return err
	}
	db := df.db
	driverName := df.options.Db
	df.mutex.Unlock()

	return CheckSQLDBReady(ctx, db, driverName)
}

// CheckReady verifies that the local datastore has no external dependency.
func (df *DataStoreFactoryLocalImpl) CheckReady(ctx context.Context) error {
	return nil
}

func CheckSQLDBReady(ctx context.Context, db *sql.DB, driverName string) error {
	if db == nil {
		return fmt.Errorf("sql db missing")
	}

	conn, err := db.Conn(ctx)
	if err != nil {
		return fmt.Errorf("get writer connection failed: %w", err)
	}

	discard := false
	defer func() {
		if discard {
			discardSQLConn(conn)
		}
		_ = conn.Close()
	}()

	if isPostgresDriver(driverName) {
		if err := checkPostgresWriterConn(ctx, conn); err != nil {
			discard = true
			return err
		}
		return nil
	}

	var one int
	if err := conn.QueryRowContext(ctx, "SELECT 1").Scan(&one); err != nil {
		discard = true
		return fmt.Errorf("writer query failed: %w", err)
	}
	return nil
}

func checkPostgresWriterConn(ctx context.Context, conn *sql.Conn) error {
	var inRecovery bool
	if err := conn.QueryRowContext(ctx, "SELECT pg_is_in_recovery()").Scan(&inRecovery); err != nil {
		return fmt.Errorf("postgres writer recovery check failed: %w", err)
	}
	if inRecovery {
		return fmt.Errorf("postgres writer is in recovery/read-only mode")
	}

	var transactionReadOnly string
	if err := conn.QueryRowContext(ctx, "SHOW transaction_read_only").Scan(&transactionReadOnly); err != nil {
		return fmt.Errorf("postgres writer read-only check failed: %w", err)
	}
	if strings.EqualFold(strings.TrimSpace(transactionReadOnly), "on") {
		return fmt.Errorf("postgres writer transaction_read_only is on/read-only")
	}

	return nil
}
