/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package sqlstore

import (
	"context"
	"database/sql"
	"time"

	"github.com/1backend/1backend/sdk/go/telemetry"
)

func instrumentDB(driverName, tableName string, db DB) DB {
	if db == nil {
		return nil
	}
	return &otelDB{
		db:         db,
		driverName: driverName,
		tableName:  tableName,
	}
}

type otelDB struct {
	db         DB
	driverName string
	tableName  string
}

func (db *otelDB) Close() error {
	return db.db.Close()
}

func (db *otelDB) SetDebug(debug bool) {
	db.db.SetDebug(debug)
}

func (db *otelDB) SkipExec(skip bool) {
	db.db.SkipExec(skip)
}

func (db *otelDB) Tablename() string {
	return db.db.Tablename()
}

func (db *otelDB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	started := time.Now()
	res, err := db.db.Query(query, args...)
	telemetry.RecordSQLStatement(context.Background(), db.driverName, db.tableName, query, started, err)
	return res, err
}

func (db *otelDB) QueryRow(query string, args ...interface{}) *sql.Row {
	started := time.Now()
	row := db.db.QueryRow(query, args...)
	telemetry.RecordSQLStatement(context.Background(), db.driverName, db.tableName, query, started, nil)
	return row
}

func (db *otelDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	started := time.Now()
	res, err := db.db.Exec(query, args...)
	telemetry.RecordSQLStatement(context.Background(), db.driverName, db.tableName, query, started, err)
	return res, err
}

func (db *otelDB) Prepare(query string) (*sql.Stmt, error) {
	started := time.Now()
	stmt, err := db.db.Prepare(query)
	telemetry.RecordSQLStatement(context.Background(), db.driverName, db.tableName, query, started, err)
	return stmt, err
}

func (db *otelDB) Begin() (Tx, error) {
	tx, err := db.db.Begin()
	if err != nil {
		return nil, err
	}
	return &otelTx{
		tx:         tx,
		driverName: db.driverName,
		tableName:  db.tableName,
	}, nil
}

type otelTx struct {
	tx         Tx
	driverName string
	tableName  string
}

func (tx *otelTx) Query(query string, args ...interface{}) (*sql.Rows, error) {
	started := time.Now()
	res, err := tx.tx.Query(query, args...)
	telemetry.RecordSQLStatement(context.Background(), tx.driverName, tx.tableName, query, started, err)
	return res, err
}

func (tx *otelTx) QueryRow(query string, args ...interface{}) *sql.Row {
	started := time.Now()
	row := tx.tx.QueryRow(query, args...)
	telemetry.RecordSQLStatement(context.Background(), tx.driverName, tx.tableName, query, started, nil)
	return row
}

func (tx *otelTx) Exec(query string, args ...interface{}) (sql.Result, error) {
	started := time.Now()
	res, err := tx.tx.Exec(query, args...)
	telemetry.RecordSQLStatement(context.Background(), tx.driverName, tx.tableName, query, started, err)
	return res, err
}

func (tx *otelTx) Prepare(query string) (*sql.Stmt, error) {
	started := time.Now()
	stmt, err := tx.tx.Prepare(query)
	telemetry.RecordSQLStatement(context.Background(), tx.driverName, tx.tableName, query, started, err)
	return stmt, err
}

func (tx *otelTx) Rollback() error {
	return tx.tx.Rollback()
}

func (tx *otelTx) Commit() error {
	return tx.tx.Commit()
}
