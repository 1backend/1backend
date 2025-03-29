package multiservice

import (
	"context"
	"net/http"
	"os"

	openapi "github.com/1backend/1backend/clients/go"
	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/middlewares"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	basicclient "github.com/1backend/1backend/examples/go/services/basic/client/go"
)

type MultiService struct {
	basicSvcClient *basicclient.APIClient

	Options *Options

	token            string
	userSvcPublicKey string

	dataStoreFactory sdk.DataStoreFactory

	credentialStore datastore.DataStore

	Router *mux.Router
}

type Options struct {
	Test      bool
	ServerUrl string
	SelfUrl   string
}

func NewService(options *Options) (*MultiService, error) {
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

	service := &MultiService{
		Options: options,
	}

	dsf, err := sdk.NewDataStoreFactory(dconf)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create datastore factory")
	}
	service.dataStoreFactory = dsf

	service.registerAccount()
	service.registerRoutes()

	return service, nil
}

func (service *MultiService) Start() error {
	client := sdk.NewApiClientFactory(service.Options.ServerUrl).Client(sdk.WithToken(service.token))

	_, _, err := client.RegistrySvcAPI.RegisterInstance(context.Background()).Body(openapi.RegistrySvcRegisterInstanceRequest{
		Url: service.Options.SelfUrl,
	}).Execute()
	if err != nil {
		return errors.Wrap(err, "cannot register instance")
	}

	return nil
}

func (service *MultiService) registerAccount() error {
	credentialStore, err := service.dataStoreFactory.Create("petSvcCredentials", &sdk.Credential{})
	if err != nil {
		return errors.Wrap(err, "cannot create credential store")
	}
	service.credentialStore = credentialStore

	client := sdk.NewApiClientFactory(service.Options.ServerUrl).Client()
	token, err := sdk.RegisterServiceAccount(
		client.UserSvcAPI,
		"multi-svc",
		"Multi Svc",
		service.credentialStore,
	)
	if err != nil {
		return errors.Wrap(err, "cannot register service")
	}
	service.token = token.Token
	service.basicSvcClient = newBasicSvcClient(service.Options.ServerUrl, service.token)

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

func (service *MultiService) registerRoutes() {
	mws := []middlewares.Middleware{
		middlewares.ThrottledLogger,
		middlewares.Recover,
		middlewares.CORS,
		middlewares.GzipDecodeMiddleware,
	}
	appl := applicator(mws)

	service.Router = mux.NewRouter()

	service.Router.HandleFunc("/multi-svc/pets/count", appl(func(w http.ResponseWriter, r *http.Request) {
		service.CountPets(w, r)
	})).
		Methods("OPTIONS", "GET")
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
