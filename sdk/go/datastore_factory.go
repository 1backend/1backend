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

func NewDatastoreFactory(tablePrefix string) (func(tableName string, instance any) (datastore.DataStore, error), error) {
	dbType := os.Getenv("OB_DB")
	dbDriver := os.Getenv("OB_DB_DRIVER")
	dbString := os.Getenv("OB_DB_STRING")

	if dbDriver == "" {
		dbDriver = "postgres"
	}

	if dbType == "" {
		homeDir, _ := os.UserHomeDir()
		localStorePath := path.Join(homeDir, ".1backend", "data")
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
				path.Join(localStorePath, tableName),
			)
		}, nil
	}

	db, err := sql.Open(dbDriver, dbString)
	if err != nil {
		return nil, errors.Wrap(err, "error opening sql db")
	}

	return func(tableName string, instance any) (datastore.DataStore, error) {
		return sqlstore.NewSQLStore(
			instance,
			dbDriver,
			db,
			tablePrefix+tableName,
			false,
		)
	}, nil
}
