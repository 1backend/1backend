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
	"fmt"
	"log/slog"
	"os"
	"runtime/debug"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/logger"
	"github.com/1backend/1backend/server/internal/di"
	node_types "github.com/1backend/1backend/server/internal/node/types"
	"github.com/gorilla/mux"

	_ "github.com/1backend/1backend/server/docs"
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
		options.GpuPlatform = os.Getenv("OB_GPU_PLATFORM")
	}
	if options.Address == "" {
		options.Address = os.Getenv("OB_SERVER_URL")
	}
	if options.NodeId == "" {
		options.NodeId = os.Getenv("OB_NODE_ID")
	}
	if options.Az == "" {
		options.Az = os.Getenv("OB_AZ")
	}
	if options.Region == "" {
		options.Region = os.Getenv("OB_AZ")
	}
	if options.LLMHost == "" {
		options.LLMHost = os.Getenv("OB_LLM_HOST")
	}
	if options.VolumeName == "" {
		options.VolumeName = os.Getenv("OB_VOLUME_NAME")
	}
	if options.DbPrefix == "" {
		options.VolumeName = os.Getenv("OB_VOLUME_NAME")
	}
	if options.DbPrefix == "" {
		options.DbPrefix = os.Getenv("OB_DB_PREFIX")
	}
	if options.Db == "" {
		options.Db = os.Getenv("OB_DB")
	}
	if options.DbConnectionString == "" {
		options.DbConnectionString = os.Getenv("OB_DB_CONNECTION_STRING")
	}
	if options.SecretEncryptionKey == "" {
		options.SecretEncryptionKey = os.Getenv("OB_ENCRYPTION_KEY")
		if options.SecretEncryptionKey == "" {
			options.SecretEncryptionKey = "changeMeToSomethingSecureForReal"
		}
	}

	diopt := &di.Options{
		NodeOptions: options,
		Test:        options.Test,
		Url:         options.Address,
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
