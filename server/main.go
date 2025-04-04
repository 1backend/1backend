/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/1backend/1backend/sdk/go/logger"
	_ "github.com/1backend/1backend/server/docs"
	"github.com/1backend/1backend/server/internal/di"
	"github.com/1backend/1backend/server/internal/router"
)

// @title           1Backend
// @version         0.3.0-rc.31
// @description     AI-native microservices platform.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://1backend.com/
// @contact.email  sales@singulatron.com

// @license.name  AGPL v3.0
// @license.url   https://www.gnu.org/licenses/agpl-3.0.html

// @host      localhost:58231
// @BasePath  /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and token acquired from the User Svc Login endpoint.

// @externalDocs.description  1Backend API
// @externalDocs.url          https://1backend.com/docs/category/1backend-api
func main() {
	logger.Info("Starting...")

	universe, err := di.BigBang(&di.Options{})
	if err != nil {
		logger.Error("Cannot start node", slog.Any("error", err))
		os.Exit(1)
	}

	srv := &http.Server{
		Handler: universe.Router,
	}

	port := router.GetPort()

	go func() {
		err := universe.StarterFunc()
		if err != nil {
			logger.Error("Cannot start universe", slog.Any("error", err))
			os.Exit(1)
		}

		time.Sleep(5 * time.Millisecond)
		logger.Info("Server started", slog.String("port", port))
	}()

	err = http.ListenAndServe(fmt.Sprintf(":%v", port), srv.Handler)
	if err != nil {
		logger.Error("HTTP listen failed", slog.String("error", err.Error()))
		os.Exit(1)
	}
}
