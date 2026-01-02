package multiservice

import (
	"context"
	"net/http"

	openapi "github.com/1backend/1backend/clients/go"
	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/boot"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/infra"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	basicclient "github.com/1backend/1backend/examples/go/services/basic/client/go"
)

type MultiService struct {
	basicSvcClient *basicclient.APIClient

	Options *boot.Options

	token            string
	userSvcPublicKey string

	dataStoreFactory infra.DataStoreFactory

	credentialStore datastore.DataStore

	Router *mux.Router
}

type Options struct {
	Test      bool
	ServerUrl string
	SelfUrl   string
}

func NewService(options *boot.Options) (*MultiService, error) {
	options.LoadEnvars()

	dconf := infra.DataStoreConfig{}
	if options.Test {
		dconf.Test = true
		dconf.TablePrefix = "t_" + sdk.Id("")
	}

	service := &MultiService{
		Options: options,
	}

	dsf, err := infra.NewDataStoreFactory(dconf)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create datastore factory")
	}
	service.dataStoreFactory = dsf

	service.registerAccount()
	service.registerRoutes()

	return service, nil
}

func (service *MultiService) Start() error {
	client := client.NewApiClientFactory(service.Options.ServerUrl).
		Client(client.WithToken(service.token))

	_, _, err := client.RegistrySvcAPI.RegisterInstance(context.Background()).Body(openapi.RegistrySvcRegisterInstanceRequest{
		Url: service.Options.SelfUrl,
	}).Execute()
	if err != nil {
		return errors.Wrap(err, "cannot register instance")
	}

	return nil
}

func (service *MultiService) registerAccount() error {
	credentialStore, err := service.dataStoreFactory.Create("multiSvcCredentials", &auth.Credential{})
	if err != nil {
		return errors.Wrap(err, "cannot create credential store")
	}
	service.credentialStore = credentialStore

	obClient := client.NewApiClientFactory(service.Options.ServerUrl).Client()
	token, err := boot.RegisterServiceAccount(
		obClient.UserSvcAPI,
		"multi-svc",
		"Multi Svc",
		service.credentialStore,
	)
	if err != nil {
		return errors.Wrap(err, "cannot register service")
	}
	service.token = token.Token
	service.basicSvcClient = newBasicSvcClient(service.Options.ServerUrl, service.token)

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

func (service *MultiService) registerRoutes() {
	appl := service.Options.Middlewares

	service.Router = mux.NewRouter()

	service.Router.HandleFunc("/multi-svc/pets/count", appl(func(w http.ResponseWriter, r *http.Request) {
		service.CountPets(w, r)
	})).
		Methods("OPTIONS", "GET")
}

func newBasicSvcClient(url, token string) *basicclient.APIClient {
	return basicclient.NewAPIClient(&basicclient.Configuration{
		Servers: basicclient.ServerConfigurations{
			{
				URL:         url,
				Description: "Default server",
			},
		},
		DefaultHeader: map[string]string{
			"Authorization": "Bearer " + token,
		},
	})
}
