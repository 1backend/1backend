/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package infra

import (
	"database/sql"
	"os"
	"path"
	"sync"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/datastore/localstore"
	"github.com/1backend/1backend/sdk/go/datastore/sqlstore"
	lock "github.com/1backend/1backend/sdk/go/lock"
	pglock "github.com/1backend/1backend/sdk/go/lock/pg"
	"github.com/1backend/1backend/sdk/go/telemetry"
	"github.com/pkg/errors"
)

type DataStoreFactoryPostgresImpl struct {
	mutex sync.Mutex

	options DataStoreConfig
	db      *sql.DB
	readDB  *sql.DB
}

type DataStoreFactoryLocalImpl struct {
	mutex sync.Mutex

	options        DataStoreConfig
	localStorePath string
}

type DataStoreFactory interface {
	Create(tableName string, instance any) (datastore.DataStore, error)

	// eg. *sql.Db
	// Don't use this, it's only for system level hacks
	Handle() (any, error)
}

type DataStoreConfig struct {
	Test                   bool
	HomeDir                string
	Db                     string
	DbConnectionString     string
	ReadDbConnectionString string
	DbApplicationName      string
	DbPool                 DbPoolConfig
	TablePrefix            string
	Lock                   lock.DistributedLock
	AutoIndexes            bool
}

func NewDataStoreFactory(options DataStoreConfig) (DataStoreFactory, error) {
	if options.HomeDir == "" {
		homeDir, err := HomeDir(HomeDirOptions{
			Test: options.Test,
		})
		if err != nil {
			return nil, err
		}
		options.HomeDir = homeDir
	}

	if options.TablePrefix == "" {
		options.TablePrefix = os.Getenv("OB_DB_PREFIX")
	}
	if options.Test && options.TablePrefix == "" {
		options.TablePrefix = "test_" + sdk.Id("")
	}

	if options.Db == "" {
		options.Db = os.Getenv("OB_DB")
	}

	if options.DbConnectionString == "" {
		options.DbConnectionString = os.Getenv("OB_DB_CONNECTION_STRING")
	}

	if options.ReadDbConnectionString == "" {
		options.ReadDbConnectionString = os.Getenv("OB_DB_READ_CONNECTION_STRING")
	}
	if !options.AutoIndexes && os.Getenv("OB_AUTO_INDEXES") == "true" {
		options.AutoIndexes = true
	}

	if err := options.loadDbConnectionRuntimeOptionsFromEnv(); err != nil {
		return nil, err
	}
	if err := options.validateDatastoreDbPoolConfig(); err != nil {
		return nil, err
	}

	// Default used for tests
	if options.DbConnectionString == "" {
		options.DbConnectionString = "postgres://postgres:mysecretpassword@localhost:5432/mydatabase?sslmode=disable"
	}

	if options.Db == "" {
		return &DataStoreFactoryLocalImpl{
			options: options,
		}, nil
	}

	return &DataStoreFactoryPostgresImpl{
		options: options,
	}, nil
}

func (df *DataStoreFactoryPostgresImpl) Create(tableName string, instance any) (datastore.DataStore, error) {
	df.mutex.Lock()
	defer df.mutex.Unlock()

	if err := df.ensureDBLocked(); err != nil {
		return nil, err
	}

	readDB, err := df.readDBLocked()
	if err != nil {
		return nil, err
	}
	if err := df.initAutoIndexLockLocked(); err != nil {
		return nil, err
	}

	var opts []sqlstore.SQLStoreOption
	if df.options.Lock != nil {
		opts = append(opts, sqlstore.WithAutoIndexLock(df.options.Lock))
	}
	opts = append(opts, sqlstore.WithAutoIndexes(df.options.AutoIndexes))
	if readDB != df.db {
		opts = append(opts, sqlstore.WithReadDB(readDB))
	}

	fullTableName := df.options.TablePrefix + tableName
	d, err := sqlstore.NewSQLStore(
		instance,
		df.options.Db,
		df.db,
		fullTableName,
		false,
		opts...,
	)
	if err != nil {
		return nil, err
	}

	// d.SetDebug(df.options.Test)

	return telemetry.InstrumentDataStore(df.options.Db, fullTableName, instance, d), nil

}

func (df *DataStoreFactoryPostgresImpl) initAutoIndexLockLocked() error {
	if df.options.Lock != nil || df.options.Db != string(sqlstore.DriverPostGRES) {
		return nil
	}

	df.options.Lock = pglock.NewPGDistributedLockFromDB(df.db)
	return nil
}

func (df *DataStoreFactoryPostgresImpl) ensureDBLocked() error {
	if df.db != nil {
		return nil
	}

	db, err := df.options.openSQLDB("write", df.options.DbConnectionString)
	if err != nil {
		return errors.Wrap(err, "error opening sql db")
	}
	df.db = db
	return nil
}

func (df *DataStoreFactoryPostgresImpl) readDBLocked() (*sql.DB, error) {
	if df.options.ReadDbConnectionString == "" {
		return df.db, nil
	}
	if df.readDB != nil {
		return df.readDB, nil
	}

	db, err := df.options.openSQLDB("read", df.options.ReadDbConnectionString)
	if err != nil {
		return nil, errors.Wrap(err, "error opening read sql db")
	}
	df.readDB = db
	return df.readDB, nil
}

func (df *DataStoreFactoryLocalImpl) Create(tableName string, instance any) (datastore.DataStore, error) {
	df.mutex.Lock()
	defer df.mutex.Unlock()

	if df.localStorePath == "" {
		localStorePath := path.Join(df.options.HomeDir, "data")
		err := os.MkdirAll(localStorePath, 0755)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create data folder")
		}

		df.localStorePath = localStorePath
	}

	fullTableName := df.options.TablePrefix + tableName
	d, err := localstore.NewLocalStore(
		instance,
		path.Join(df.localStorePath, fullTableName),
	)
	if err != nil {
		return nil, err
	}

	//d.SetDebug(df.options.Test)

	return telemetry.InstrumentDataStore("localstore", fullTableName, instance, d), nil
}

func (df *DataStoreFactoryLocalImpl) Handle() (any, error) {
	return nil, nil
}

func (df *DataStoreFactoryPostgresImpl) Handle() (any, error) {
	df.mutex.Lock()
	defer df.mutex.Unlock()

	if err := df.ensureDBLocked(); err != nil {
		return nil, err
	}

	return df.db, nil
}

func (df *DataStoreFactoryPostgresImpl) SetLock(distributedLock lock.DistributedLock) {
	df.mutex.Lock()
	defer df.mutex.Unlock()

	df.options.Lock = distributedLock
}

func SetDataStoreFactoryLock(factory DataStoreFactory, distributedLock lock.DistributedLock) {
	lockAwareFactory, ok := factory.(interface {
		SetLock(lock.DistributedLock)
	})
	if !ok {
		return
	}
	lockAwareFactory.SetLock(distributedLock)
}

func (df *DataStoreFactoryLocalImpl) SetLock(distributedLock lock.DistributedLock) {
	df.mutex.Lock()
	defer df.mutex.Unlock()

	df.options.Lock = distributedLock
}
