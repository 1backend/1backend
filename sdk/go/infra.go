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
	Test               bool
	DbName             string
	DbConnectionString string
	HomeDir            string
}

type Infra struct {
	DatastoreFactory func(tableName string, instance any) (datastore.DataStore, error)
}

func InfraFactory(options InfraOptions) (*Infra, error) {
	if options.HomeDir == "" {
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

		options.HomeDir = homeDir
	}

	infra := &Infra{}

	if options.DbName == "" {
		options.DbName = os.Getenv("OB_DB")
	}
	if options.DbConnectionString == "" {
		options.DbConnectionString = os.Getenv("OB_DB_CONNECTION_STRING")
	}

	if options.DbName == "" {
		localStorePath := path.Join(options.HomeDir, "data")
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
	} else if options.DbName == "postgres" {
		db, err := sql.Open("postgres", options.DbConnectionString)
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
