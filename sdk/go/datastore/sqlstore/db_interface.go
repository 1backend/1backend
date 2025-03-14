package sqlstore

import "database/sql"

type Tx interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Rollback() error
	Commit() error
}

type DB interface {
	Close() error

	// Enable or disable debug logs
	SetDebug(debug bool)

	// Enable or disable query execution
	SkipExec(skip bool)

	Tablename() string
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Begin() (Tx, error)
}
