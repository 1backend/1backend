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

	"github.com/1backend/1backend/sdk/go/pubsub"
	"github.com/1backend/1backend/sdk/go/pubsub/localpubsub"
	"github.com/1backend/1backend/sdk/go/pubsub/pgpubsub"
)

type PubSubFactory interface {
	Create(name string) (pubsub.PubSub, error)
}

type PubSubFactoryLocalImpl struct {
	mu      sync.Mutex
	options DataStoreConfig
	homeDir string
}

type PubSubFactoryPostgresImpl struct {
	mu      sync.Mutex
	options DataStoreConfig
	db      *sql.DB
}

func NewPubSubFactory(options DataStoreConfig) (PubSubFactory, error) {
	if options.HomeDir == "" {
		homeDir, err := HomeDir(HomeDirOptions{Test: options.Test})
		if err != nil {
			return nil, err
		}
		options.HomeDir = homeDir
	}

	if options.Db == "" {
		options.Db = os.Getenv("OB_DB")
	}
	if options.DbConnectionString == "" {
		options.DbConnectionString = os.Getenv("OB_DB_CONNECTION_STRING")
	}
	if err := options.loadDbConnectionRuntimeOptionsFromEnv(); err != nil {
		return nil, err
	}
	if options.DbConnectionString == "" {
		options.DbConnectionString = "postgres://postgres:mysecretpassword@localhost:5432/mydatabase?sslmode=disable"
	}

	if options.Db == "" {
		return &PubSubFactoryLocalImpl{options: options}, nil
	}

	return &PubSubFactoryPostgresImpl{options: options}, nil
}

func (pf *PubSubFactoryLocalImpl) Create(name string) (pubsub.PubSub, error) {
	pf.mu.Lock()
	defer pf.mu.Unlock()

	if pf.homeDir == "" {
		pf.homeDir = path.Join(pf.options.HomeDir, "pubsub")
		if err := os.MkdirAll(pf.homeDir, 0755); err != nil {
			return nil, err
		}
	}

	return localpubsub.NewLocalPubSub(path.Join(pf.homeDir, name+".log"))
}

func (pf *PubSubFactoryPostgresImpl) Create(name string) (pubsub.PubSub, error) {
	pf.mu.Lock()
	defer pf.mu.Unlock()

	if pf.db == nil {
		db, err := pf.options.openSQLDB("pubsub", pf.options.DbConnectionString)
		if err != nil {
			return nil, err
		}
		pf.db = db
	}

	listenerConnectionString := pf.options.connectionStringForRole("pubsub-listener", pf.options.DbConnectionString)
	return pgpubsub.NewPGPubSub(listenerConnectionString, pf.db, "")
}
