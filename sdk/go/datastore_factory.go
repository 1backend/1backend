package sdk

import (
	"database/sql"
	"log/slog"
	"os"
	"path"

	"github.com/openorch/openorch/sdk/go/datastore"
	"github.com/openorch/openorch/sdk/go/datastore/localstore"
	"github.com/openorch/openorch/sdk/go/datastore/sqlstore"
	"github.com/openorch/openorch/sdk/go/logger"
	"github.com/pkg/errors"
)

func NewDatastoreFactory(tablePrefix string) (func(tableName string, instance any) (datastore.DataStore, error), error) {
	dbType := os.Getenv("OPENORCH_DB")
	dbDriver := os.Getenv("OPENORCH_DB_DRIVER")
	dbString := os.Getenv("OPENORCH_DB_STRING")

	if dbDriver == "" {
		dbDriver = "postgres"
	}

	if dbType == "" {
		homeDir, _ := os.UserHomeDir()
		localStorePath := path.Join(homeDir, ".openorch", "data")
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
