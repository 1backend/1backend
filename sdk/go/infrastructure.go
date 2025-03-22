package sdk

import (
	"os"
	"path"

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/pkg/errors"
)

type InfrastructureOptions struct {
	// Is this a test environment?
	Test bool

	// eg. postgres
	Db string

	DbConnectionString string

	HomeDir string
}

type Infrastructure struct {
	DatastoreConstructor func(tableName string, instance any) (datastore.DataStore, error)
}

func NewInfrastructure(options InfrastructureOptions) (*Infrastructure, error) {
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

	infra := &Infrastructure{}

	if options.Db == "" {
		options.Db = os.Getenv("OB_DB")
	}
	if options.DbConnectionString == "" {
		options.DbConnectionString = os.Getenv("OB_DB_CONNECTION_STRING")
	}

	dopts := DatastoreConstructorOptions{
		HomeDir:            options.HomeDir,
		Db:                 options.Db,
		DbConnectionString: options.DbConnectionString,
	}

	datastoreConstructor, err := NewDatastoreConstructor(dopts)
	if err != nil {
		return nil, err
	}

	infra.DatastoreConstructor = datastoreConstructor

	return infra, nil
}
