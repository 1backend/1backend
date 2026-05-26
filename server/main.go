/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/1backend/1backend/sdk/go/logger"
	"github.com/1backend/1backend/sdk/go/telemetry"
	_ "github.com/1backend/1backend/server/docs"
	"github.com/1backend/1backend/server/internal/di"
	"github.com/1backend/1backend/server/internal/router"
	"github.com/1backend/1backend/server/internal/universe"
)

// @title           1Backend
// @version         0.9.15
// @description     AI-native microservices platform.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://1backend.com/
// @contact.email  sales@singulatron.com

// @license.name  AGPL v3.0
// @license.url   https://www.gnu.org/licenses/agpl-3.0.html

// @host      localhost:11337
// @BasePath  /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and token acquired from the User Svc Login endpoint.

// @externalDocs.description  1Backend API
// @externalDocs.url          https://1backend.com/docs/category/1backend-api
func main() {
	autoIndexes := flag.Bool("auto-indexes", false, "enable query-observed automatic datastore index creation")
	flag.Parse()

	logger.Info("Starting...")

	telemetryShutdown, metricsPath, err := telemetry.Setup(context.Background(), telemetry.Config{
		ServiceVersion: "0.9.15",
	})
	if err != nil {
		logger.Error("Cannot initialize telemetry", slog.Any("error", err))
		os.Exit(1)
	}
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := telemetryShutdown(ctx); err != nil {
			logger.Error("Cannot shutdown telemetry", slog.Any("error", err))
		}
	}()
	if metricsPath != "" {
		logger.Info("OpenTelemetry metrics endpoint enabled", slog.String("path", metricsPath))
	}

	universe, err := di.BigBang(&universe.Options{
		AutoIndexes: *autoIndexes,
	})
	if err != nil {
		logger.Error("Cannot start node", slog.Any("error", err))
		os.Exit(1)
	}

	srv := &http.Server{
		Handler: universe.Router,
	}

	port := router.GetPort()
	srv.Addr = fmt.Sprintf(":%v", port)

	shutdownCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	go func() {
		<-shutdownCtx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			logger.Error("HTTP shutdown failed", slog.String("error", err.Error()))
			_ = srv.Close()
		}
	}()

	go func() {
		err := universe.StarterFunc()
		if err != nil {
			logger.Error("Cannot start universe", slog.Any("error", err))
			os.Exit(1)
		}

		time.Sleep(5 * time.Millisecond)
		logger.Info("Server started", slog.String("port", port))
	}()

	err = srv.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Error("HTTP listen failed", slog.String("error", err.Error()))
		os.Exit(1)
	}
}
