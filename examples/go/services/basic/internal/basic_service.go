package basicservice

import (
	"context"
	"net/http"

	openapi "github.com/1backend/1backend/clients/go"
	basic "github.com/1backend/1backend/examples/go/services/basic/internal/types"
	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/boot"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/infra"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

const RolePetManager = "basic-svc:pet:manager"

type BasicService struct {
	Options *boot.Options

	token            string
	userSvcPublicKey string

	dataStoreFactory infra.DataStoreFactory

	petsStore       datastore.DataStore
	credentialStore datastore.DataStore

	Router *mux.Router
}

type Options struct {
	Test      bool
	ServerUrl string
	SelfUrl   string
}

func NewService(options *boot.Options) (*BasicService, error) {
	options.LoadEnvars()

	dconf := infra.DataStoreConfig{}
	if options.Test {
		dconf.Test = true
		dconf.TablePrefix = "t_" + sdk.Id("")
	}

	service := &BasicService{
		Options: options,
	}

	dsf, err := infra.NewDataStoreFactory(dconf)
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
	client := client.NewApiClientFactory(service.Options.ServerUrl).
		Client(client.WithToken(service.token))

	_, _, err := client.RegistrySvcAPI.
		RegisterInstance(context.Background()).
		Body(openapi.RegistrySvcRegisterInstanceRequest{
			Url: service.Options.SelfUrl,
		}).Execute()
	if err != nil {
		return errors.Wrap(err, "cannot register instance")
	}

	return nil
}

func (service *BasicService) registerAccount() error {
	credentialStore, err := service.dataStoreFactory.Create("basicSvcCredentials", &auth.Credential{})
	if err != nil {
		return errors.Wrap(err, "cannot create credential store")
	}
	service.credentialStore = credentialStore

	obClient := client.NewApiClientFactory(service.Options.ServerUrl).Client()
	token, err := boot.RegisterServiceAccount(
		obClient.UserSvcAPI,
		"basic-svc",
		"Basic Svc",
		service.credentialStore,
	)
	if err != nil {
		return errors.Wrap(err, "cannot register service")
	}
	service.token = token.Token

	obClient = client.NewApiClientFactory(service.Options.ServerUrl).
		Client(client.WithToken(service.token))

	_, _, err = obClient.RegistrySvcAPI.
		RegisterInstance(context.Background()).
		Body(openapi.RegistrySvcRegisterInstanceRequest{
			Url: service.Options.SelfUrl,
		}).Execute()
	if err != nil {
		return errors.Wrap(err, "cannot register instance")
	}

	pk, _, err := obClient.
		UserSvcAPI.GetPublicKey(context.Background()).
		Execute()
	if err != nil {
		return err
	}
	service.userSvcPublicKey = pk.PublicKey

	return nil
}

func (service *BasicService) registerRoutes() {
	appl := service.Options.Middlewares

	service.Router = mux.NewRouter()

	service.Router.HandleFunc("/basic-svc/pet", appl(func(w http.ResponseWriter, r *http.Request) {
		service.SavePet(w, r)
	})).
		Methods("OPTIONS", "PUT")

	service.Router.HandleFunc("/basic-svc/pets", appl(func(w http.ResponseWriter, r *http.Request) {
		service.ListPets(w, r)
	})).
		Methods("OPTIONS", "POST")

	service.Router.HandleFunc("/basic-svc/error", appl(func(w http.ResponseWriter, r *http.Request) {
		service.Error(w, r)
	})).
		Methods("OPTIONS", "POST")
}
