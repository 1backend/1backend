/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package infra

import (
	"context"
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

	options  DataStoreConfig
	db       *sql.DB
	lockConn *sql.Conn
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
	Test               bool
	HomeDir            string
	Db                 string
	DbConnectionString string
	TablePrefix        string
	Lock               lock.DistributedLock
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

	if df.db == nil {
		db, err := sql.Open(df.options.Db, df.options.DbConnectionString)
		if err != nil {
			return nil, errors.Wrap(err, "error opening sql db")
		}
		df.db = db
	}
	if err := df.initAutoIndexLockLocked(); err != nil {
		return nil, err
	}

	var opts []sqlstore.SQLStoreOption
	if df.options.Lock != nil {
		opts = append(opts, sqlstore.WithAutoIndexLock(df.options.Lock))
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

	conn, err := df.db.Conn(context.Background())
	if err != nil {
		return errors.Wrap(err, "error creating automatic index lock connection")
	}
	df.lockConn = conn
	df.options.Lock = pglock.NewPGDistributedLock(conn)
	return nil
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

	if df.db == nil {
		db, err := sql.Open(df.options.Db, df.options.DbConnectionString)
		if err != nil {
			return nil, errors.Wrap(err, "error opening sql db")
		}
		df.db = db
	}

	return df.db, nil
}

func (df *DataStoreFactoryPostgresImpl) SetLock(distributedLock lock.DistributedLock) {
	df.mutex.Lock()
	defer df.mutex.Unlock()

	if df.lockConn != nil {
		_ = df.lockConn.Close()
		df.lockConn = nil
	}
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
