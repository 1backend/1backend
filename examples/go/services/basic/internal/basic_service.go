package basicservice

import (
	"context"
	"net/http"
	"os"

	openapi "github.com/1backend/1backend/clients/go"
	basic "github.com/1backend/1backend/examples/go/services/basic/internal/types"
	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/middlewares"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

const RolePetManager = "basic-svc:pet:manager"

type BasicService struct {
	Options *Options

	token            string
	userSvcPublicKey string

	dataStoreFactory sdk.DataStoreFactory

	petsStore       datastore.DataStore
	credentialStore datastore.DataStore

	Router *mux.Router
}

type Options struct {
	Test      bool
	ServerUrl string
	SelfUrl   string
}

func NewService(options *Options) (*BasicService, error) {
	if options.ServerUrl == "" {
		options.ServerUrl = os.Getenv("OB_SERVER_URL")
	}
	if options.ServerUrl == "" {
		options.ServerUrl = "http://127.0.0.1:58231"
	}
	if options.SelfUrl == "" {
		options.SelfUrl = os.Getenv("OB_SELF_URL")
	}

	dconf := sdk.DataStoreConfig{}
	if options.Test {
		dconf.TablePrefix = sdk.Id("t")
	}

	service := &BasicService{
		Options: options,
	}

	dsf, err := sdk.NewDataStoreFactory(dconf)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create datastore factory")
	}
	service.dataStoreFactory = dsf

	petStore, err := dsf.Create("basicSvcPets", &basic.Pet{})
	if err != nil {
		return nil, err
	}
	service.petsStore = petStore

	service.registerAccount()
	service.registerRoutes()

	return service, nil
}

func (service *BasicService) Start() error {
	client := sdk.NewApiClientFactory(service.Options.ServerUrl).Client(sdk.WithToken(service.token))

	_, _, err := client.RegistrySvcAPI.RegisterInstance(context.Background()).Body(openapi.RegistrySvcRegisterInstanceRequest{
		Url: service.Options.SelfUrl,
	}).Execute()
	if err != nil {
		return errors.Wrap(err, "cannot register instance")
	}

	return nil
}

func (service *BasicService) registerAccount() error {
	credentialStore, err := service.dataStoreFactory.Create("petSvcCredentials", &sdk.Credential{})
	if err != nil {
		return errors.Wrap(err, "cannot create credential store")
	}
	service.credentialStore = credentialStore

	client := sdk.NewApiClientFactory(service.Options.ServerUrl).Client()
	token, err := sdk.RegisterServiceAccount(
		client.UserSvcAPI,
		"basic-svc",
		"Basic Svc",
		service.credentialStore,
	)
	if err != nil {
		return errors.Wrap(err, "cannot register service")
	}
	service.token = token.Token

	client = sdk.NewApiClientFactory(service.Options.ServerUrl).Client(sdk.WithToken(service.token))

	_, _, err = client.RegistrySvcAPI.RegisterInstance(context.Background()).Body(openapi.RegistrySvcRegisterInstanceRequest{
		Url: service.Options.SelfUrl,
	}).Execute()
	if err != nil {
		return errors.Wrap(err, "cannot register instance")
	}

	pk, _, err := client.
		UserSvcAPI.GetPublicKey(context.Background()).
		Execute()
	if err != nil {
		return err
	}
	service.userSvcPublicKey = pk.PublicKey

	return nil
}

func (service *BasicService) registerRoutes() {
	mws := []middlewares.Middleware{
		middlewares.ThrottledLogger,
		middlewares.Recover,
		middlewares.CORS,
		middlewares.GzipDecodeMiddleware,
	}
	appl := applicator(mws)

	service.Router = mux.NewRouter()

	service.Router.HandleFunc("/basic-svc/pet", appl(func(w http.ResponseWriter, r *http.Request) {
		service.SavePet(w, r)
	})).
		Methods("OPTIONS", "PUT")

	service.Router.HandleFunc("/basic-svc/pets", appl(func(w http.ResponseWriter, r *http.Request) {
		service.ListPets(w, r)
	})).
		Methods("OPTIONS", "POST")
}

func applicator(
	mws []middlewares.Middleware,
) func(http.HandlerFunc) http.HandlerFunc {
	return func(h http.HandlerFunc) http.HandlerFunc {
		for _, middleware := range mws {
			h = middleware(h)
		}

		return h
	}
}
