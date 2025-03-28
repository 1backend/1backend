/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package sdk

import (
	"database/sql"
	"os"
	"path"
	"sync"

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/datastore/localstore"
	"github.com/1backend/1backend/sdk/go/datastore/sqlstore"
	"github.com/pkg/errors"
)

type DataStoreFactoryPostgresImpl struct {
	mutex sync.Mutex

	options DataStoreConfig
	db      *sql.DB
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
}

func NewDataStoreFactory(options DataStoreConfig) (DataStoreFactory, error) {
	if options.HomeDir == "" {
		homeDir, err := HomeDir(HomeDirOptions{})
		if err != nil {
			return nil, err
		}
		options.HomeDir = homeDir
	}

	if options.TablePrefix == "" {
		options.TablePrefix = os.Getenv("OB_DB_PREFIX")
	}
	if options.Test && options.TablePrefix == "" {
		options.TablePrefix = Id("test") + "_"
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

	d, err := sqlstore.NewSQLStore(
		instance,
		df.options.Db,
		df.db,
		df.options.TablePrefix+tableName,
		false,
	)
	if err != nil {
		return nil, err
	}

	// d.SetDebug(df.options.Test)

	return d, nil

}

func (df *DataStoreFactoryLocalImpl) Create(tableName string, instance any) (datastore.DataStore, error) {
	df.mutex.Lock()
	defer df.mutex.Unlock()

	if df.localStorePath == "" {
		localStorePath := path.Join(df.options.HomeDir, onebackendFolder, "data")
		err := os.MkdirAll(localStorePath, 0755)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create data folder")
		}

		df.localStorePath = localStorePath
	}

	d, err := localstore.NewLocalStore(
		instance,
		path.Join(df.localStorePath, df.options.TablePrefix+tableName),
	)
	if err != nil {
		return nil, err
	}

	//d.SetDebug(df.options.Test)

	return d, nil
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
