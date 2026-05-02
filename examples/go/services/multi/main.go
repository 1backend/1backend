package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	multi "github.com/1backend/1backend/examples/go/services/multi/internal"
	"github.com/1backend/1backend/sdk/go/boot"
)

// @title           Multi Svc
// @version         0.3.0-rc.8
// @description     An example service for bootstrapping.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    https://1backend.com/
// @contact.email  sales@singulatron.com

// @license.name  Proprietary

// @host      localhost:11337
// @BasePath  /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and token acquired from the User Svc Login endpoint.
func main() {
	selfUrl := os.Getenv("OB_SELF_URL")
	if selfUrl == "" {
		selfUrl = "http://127.0.0.1:9211"
	}

	options := &boot.Options{
		SelfUrl: selfUrl,
	}

	telemetryShutdown, metricsPath, err := options.SetupTelemetry(context.Background(), "multi-svc")
	if err != nil {
		log.Fatalf("Failed to initialize telemetry: %v", err)
	}
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := telemetryShutdown(ctx); err != nil {
			log.Printf("Failed to shut down telemetry: %v", err)
		}
	}()

	multiService, err := multi.NewService(options)
	if err != nil {
		log.Fatalf("Failed to initialize multi service: %v", err)
	}
	metricsRoute := options.InstrumentRouter(multiService.Router, "multi-svc", metricsPath)

	err = multiService.Start()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	log.Println("Server started on " + selfUrl)
	if metricsRoute != "" {
		log.Println("Telemetry metrics endpoint enabled at " + metricsRoute)
	}
	log.Fatal(http.ListenAndServe(boot.ListenAddress(selfUrl), multiService.Router))
}
