/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package infra

import (
	"os"
	"path"

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
	DatastoreFactory DataStoreFactory
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

	dopts := DataStoreConfig{
		HomeDir:            options.HomeDir,
		Db:                 options.Db,
		DbConnectionString: options.DbConnectionString,
	}

	dataStoreFactory, err := NewDataStoreFactory(dopts)
	if err != nil {
		return nil, err
	}

	infra.DatastoreFactory = dataStoreFactory

	return infra, nil
}
