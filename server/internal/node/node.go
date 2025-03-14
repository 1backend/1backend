/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package node

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"runtime/debug"

	"github.com/gorilla/mux"
	sdk "github.com/openorch/openorch/sdk/go"
	"github.com/openorch/openorch/sdk/go/datastore"
	"github.com/openorch/openorch/sdk/go/datastore/sqlstore"
	pglock "github.com/openorch/openorch/sdk/go/lock/pg"
	"github.com/openorch/openorch/sdk/go/logger"
	"github.com/openorch/openorch/server/internal/di"
	node_types "github.com/openorch/openorch/server/internal/node/types"
	"github.com/pkg/errors"

	_ "github.com/openorch/openorch/server/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

type NodeInfo struct {
	Options       *di.Options
	Router        *mux.Router
	StarterFunc   func() error
	ClientFactory sdk.ClientFactory
}

// Start wraps the dependency injection universe creation
// so getting envars happens outside of that. The two could probably be merged.
// Node options are a set of node specific configuration options and secrets required for bootstrapping.
func Start(options *node_types.Options) (*NodeInfo, error) {

	defer func() {
		if r := recover(); r != nil {
			logger.Error("Panic in node.Start()",
				slog.String("error", fmt.Sprintf("%v", r)),
				slog.String("trace", string(debug.Stack())),
			)
			os.Exit(1)
		}
	}()

	// @todo GPU platform maybe this could be autodetected
	if options.GpuPlatform == "" {
		options.GpuPlatform = os.Getenv("OPENORCH_GPU_PLATFORM")
	}
	if options.Address == "" {
		options.Address = os.Getenv("OPENORCH_URL")
	}
	if options.NodeId == "" {
		options.NodeId = os.Getenv("OPENORCH_NODE_ID")
	}
	if options.Az == "" {
		options.Az = os.Getenv("OPENORCH_AZ")
	}
	if options.Region == "" {
		options.Region = os.Getenv("OPENORCH_AZ")
	}
	if options.LLMHost == "" {
		options.LLMHost = os.Getenv("OPENORCH_LLM_HOST")
	}
	if options.VolumeName == "" {
		options.VolumeName = os.Getenv("OPENORCH_VOLUME_NAME")
	}
	if options.ConfigPath == "" {
		options.ConfigPath = os.Getenv("OPENORCH_CONFIG_PATH")
	}
	if options.Db == "" {
		options.Db = os.Getenv("OPENORCH_DB")
	}
	if options.DbDriver == "" {
		options.DbDriver = os.Getenv("OPENORCH_DB_DRIVER")
	}
	if options.DbString == "" {
		options.DbString = os.Getenv("OPENORCH_DB_STRING")
	}
	if options.SecretEncryptionKey == "" {
		options.SecretEncryptionKey = os.Getenv("OPENORCH_ENCRYPTION_KEY")
		if options.SecretEncryptionKey == "" {
			options.SecretEncryptionKey = "changeMeToSomethingSecureForReal"
		}
	}

	diopt := &di.Options{
		NodeOptions: options,
		Test:        options.Test,
		Url:         options.Address,
	}

	var tablePrefix string
	if options.DbPrefix != "" {
		tablePrefix = options.DbPrefix
	}

	ctx := context.Background()

	if options.Db != "" {
		if options.DbDriver == "" {
			options.DbDriver = "postgres"
		}
		if options.DbString == "" {
			options.DbString = "postgres://postgres:mysecretpassword@localhost:5432/mydatabase?sslmode=disable"
		}

		db, err := sql.Open(options.DbDriver, options.DbString)
		if err != nil {
			return nil, errors.Wrap(err, "error opening sql db")
		}

		conn, err := db.Conn(ctx)
		if err != nil {
			return nil, err
		}
		diopt.Lock = pglock.NewPGDistributedLock(conn)

		diopt.DatastoreFactory = func(tableName string, instance any) (datastore.DataStore, error) {
			return sqlstore.NewSQLStore(
				instance,
				options.DbDriver,
				db,
				tablePrefix+tableName,
				false,
			)
		}
	}

	router, starter, err := di.BigBang(diopt)
	if err != nil {
		logger.Error("Cannot make universe", slog.Any("error", err))
		os.Exit(1)
	}

	router.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	return &NodeInfo{
		Router:      router,
		StarterFunc: starter,
		Options:     diopt,
	}, nil
}
