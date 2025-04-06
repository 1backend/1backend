package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	basic "github.com/1backend/1backend/examples/go/services/basic/internal"
	"github.com/1backend/1backend/sdk/go/boot"
)

// @title           Basic Svc
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
		selfUrl = "http://127.0.0.1:9111"
	}

	basicService, err := basic.NewService(&basic.Options{
		SelfUrl: selfUrl,
	})
	if err != nil {
		log.Fatalf("Failed to initialize basic service: %v", err)
	}

	err = basicService.Start()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	log.Println("Server started on " + selfUrl)
	log.Fatal(http.ListenAndServe(boot.ListenAddress(selfUrl), basicService.Router))
}
