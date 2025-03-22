package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	openapi "github.com/1backend/1backend/clients/go"
	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/pkg/errors"
)

func main() {
	skeletonService, err := NewService()
	if err != nil {
		log.Fatalf("Failed to initialize skeleton service: %v", err)
	}

	router := http.NewServeMux()

	router.HandleFunc("/skeleton-svc/hello", func(w http.ResponseWriter, r *http.Request) {
		skeletonService.Hello(w, r)
	})

	log.Println("Server started on :9311")
	log.Fatal(http.ListenAndServe(":9311", router))

}

type SkeletonService struct {
	token string
}

func NewService() (*SkeletonService, error) {
	spUrl := os.Getenv("OB_URL")
	if spUrl == "" {
		return nil, errors.New("OB_URL cannot be found")
	}

	selfUrl := os.Getenv("SELF_URL")

	dsf, err := sdk.NewDatastoreFactory("")
	if err != nil {
		return nil, errors.Wrap(err, "cannot create datastore factory")
	}

	credentialStore, err := dsf("skeletonSvcCredentials", &sdk.Credential{})
	if err != nil {
		return nil, errors.Wrap(err, "cannot create credential store")
	}

	client := sdk.NewApiClientFactory(spUrl).Client()
	token, err := sdk.RegisterServiceAccount(
		client.UserSvcAPI,
		"skeleton-svc",
		"Skeleton Svc",
		credentialStore,
	)
	if err != nil {
		return nil, errors.Wrap(err, "cannot register service")
	}

	client = sdk.NewApiClientFactory(spUrl).Client(sdk.WithToken(token))
	_, _, err = client.RegistrySvcAPI.
		RegisterInstance(context.Background()).
		Body(openapi.RegistrySvcRegisterInstanceRequest{
			Url: selfUrl,
		}).Execute()
	if err != nil {
		return nil, errors.Wrap(err, "cannot register instance")
	}

	repo := &SkeletonService{
		token: token,
	}

	return repo, nil

}

func (skeleton *SkeletonService) Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"hello": "world"}`)
}
