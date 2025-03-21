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

const onebackendFolder = ".1backend"

type InfraOptions struct {
	Test bool
}

type Infra struct {
	DatastoreFactory func(tableName string, instance any) (datastore.DataStore, error)
}

func InfraFactory(options InfraOptions) (*Infra, error) {
	var homeDir string
	var err error
	if options.Test {
		homeDir, err = os.MkdirTemp("", "1backend-")
		if err != nil {
			return nil, errors.Wrap(err,
				"homedir creation failed",
			)
		}
	} else {
		homeDir, err = os.UserHomeDir()
		if err != nil {
			return nil, errors.Wrap(err, "homedir creation failed")
		}
		homeDir = path.Join(homeDir, onebackendFolder)
	}

	infra := &Infra{}

	dbName := os.Getenv("OB_DB")
	if dbName == "" {
		localStorePath := path.Join(homeDir, "data")
		err := os.MkdirAll(localStorePath, 0755)
		if err != nil {
			logger.Error(
				"Creating data folder failed",
				slog.String("error", err.Error()),
			)
			os.Exit(1)
		}

		infra.DatastoreFactory = func(tableName string, instance any) (datastore.DataStore, error) {
			return localstore.NewLocalStore(
				instance,
				path.Join(localStorePath, tableName),
			)
		}
	} else if dbName == "postgres" {
		db, err := sql.Open("postgres", os.Getenv("OB_DB_CONNECTION_STRING"))
		if err != nil {
			return nil, errors.Wrap(err, "error opening sql db")
		}

		infra.DatastoreFactory = func(tableName string, instance any) (datastore.DataStore, error) {
			return sqlstore.NewSQLStore(
				instance,
				"postgres",
				db,
				tableName,
				false,
			)
		}
	}

	return infra, nil
}
