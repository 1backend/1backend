/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package sqlstore

import (
	"database/sql"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/openorch/openorch/sdk/go/logger"
)

type DebugDB struct {
	*sql.DB
	debug bool

	// skipExec is only here to avoid having to use a mock for the tests
	skipExec bool

	tableName string

	// when debug is enabled all queries are stored here
	queries []string
	values  [][]interface{}
}

type DebugTx struct {
	*sql.Tx
	debug     bool
	tableName string

	// skipExec is only here to avoid having to use a mock for the tests
	skipExec bool

	// when debug is enabled all queries are stored here
	queries []string
	values  [][]interface{}
}

func NewDebugDB(db *sql.DB, tableName string) *DebugDB {
	return &DebugDB{
		DB:        db,
		tableName: tableName,
	}
}

func (db *DebugDB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	res, err := db.DB.Query(query, args...)
	db.logQuery(query, err, args...)
	return res, err
}

func (db *DebugDB) SetDebug(debug bool) {
	db.debug = debug
	if !debug {
		db.queries = nil
		db.values = nil
	}
}

func (db *DebugDB) SkipExec(skip bool) {
	db.skipExec = skip
}

func (db *DebugDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	var (
		res sql.Result
		err error
	)
	if !db.skipExec {
		res, err = db.DB.Exec(query, args...)
	}
	db.logQuery(query, err, args...)
	return res, err
}

func (db *DebugDB) Prepare(query string) (*sql.Stmt, error) {
	var (
		res *sql.Stmt
		err error
	)
	if !db.skipExec {
		res, err = db.DB.Prepare(query)
	}
	db.logQuery(query, err, nil)
	return res, err
}

func (db *DebugDB) Tablename() string {
	return db.tableName
}

func (db *DebugDB) logQuery(query string, err error, args ...interface{}) {
	if db.debug {
		logger.Info(fmt.Sprintf("[%v] [ERROR: %v] Executing query: %s, args: %v", db.tableName, err, query, spew.Sdump(args)))
		db.queries = append(db.queries, query)
		db.values = append(db.values, args)
	}
}

func (db *DebugDB) Begin() (Tx, error) {
	tx, err := db.DB.Begin()
	if err != nil {
		return nil, err
	}
	return &DebugTx{
		Tx:        tx,
		debug:     db.debug,
		tableName: db.tableName,
	}, nil
}

func (db *DebugTx) Query(query string, args ...interface{}) (*sql.Rows, error) {
	var (
		res *sql.Rows
		err error
	)
	if !db.skipExec {
		res, err = db.Tx.Query(query, args...)
	}

	db.logQuery(query, err, args...)
	return res, err
}

func (db *DebugTx) Exec(query string, args ...interface{}) (sql.Result, error) {
	var (
		res sql.Result
		err error
	)
	if !db.skipExec {
		res, err = db.Tx.Exec(query, args...)
	}
	db.logQuery(query, err, args...)
	return res, err
}

func (db *DebugTx) Prepare(query string) (*sql.Stmt, error) {
	db.logQuery(query, nil)
	if db.skipExec {
		return db.Tx.Prepare(query)
	}
	return nil, nil
}

func (db *DebugTx) logQuery(query string, err error, args ...interface{}) {
	if db.debug {
		logger.Info(fmt.Sprintf("[%v] [ERROR: %v] [TRANSACTION] Executing query: %s, args: %v", db.tableName, err, query, spew.Sdump(args)))
		db.queries = append(db.queries, query)
		db.values = append(db.values, args)
	}
}

func (db *DebugTx) SetDebug(debug bool) {
	db.debug = debug
}

func (db *DebugTx) SkipExec(skip bool) {
	db.skipExec = skip
}
