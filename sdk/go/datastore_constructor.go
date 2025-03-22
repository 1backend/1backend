package sdk

import (
	"database/sql"
	"log/slog"
	"os"
	"path"

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/datastore/localstore"
	"github.com/1backend/1backend/sdk/go/datastore/sqlstore"
	"github.com/1backend/1backend/sdk/go/logger"
	"github.com/pkg/errors"
)

type DatastoreConstructor func(tableName string, instance any) (datastore.DataStore, error)

type DatastoreConstructorOptions struct {
	Test               bool
	HomeDir            string
	Db                 string
	DbConnectionString string
	TablePrefix        string
}

func NewDatastoreConstructor(options DatastoreConstructorOptions) (DatastoreConstructor, error) {
	if options.HomeDir == "" {
		homeDir, err := HomeDir(HomeDirOptions{})
		if err != nil {
			return nil, err
		}
		options.HomeDir = homeDir
	}

	if options.Test && options.TablePrefix == "" {
		options.TablePrefix = Id("test") + "_"
	}

	if options.Db == "" {
		localStorePath := path.Join(options.HomeDir, onebackendFolder, "data")
		err := os.MkdirAll(localStorePath, 0755)
		if err != nil {
			logger.Error(
				"Creating data folder failed",
				slog.String("error", err.Error()),
			)
			os.Exit(1)
		}

		return func(tableName string, instance any) (datastore.DataStore, error) {
			return localstore.NewLocalStore(
				instance,
				path.Join(localStorePath, options.TablePrefix+tableName),
			)
		}, nil
	}

	if options.Db == "" {
		options.Db = os.Getenv("OB_DB")
	}

	if options.DbConnectionString == "" {
		options.DbConnectionString = os.Getenv("OB_DB_CONNECTION_STRING")
	}

	db, err := sql.Open(options.Db, options.DbConnectionString)
	if err != nil {
		return nil, errors.Wrap(err, "error opening sql db")
	}

	return func(tableName string, instance any) (datastore.DataStore, error) {
		return sqlstore.NewSQLStore(
			instance,
			options.Db,
			db,
			options.TablePrefix+tableName,
			false,
		)
	}, nil
}
