---
sidebar_position: 3
tags:
  - test
---

# Your first service

While 1Backend itself is written in Go, services that run on it can be written in any language.

A service only needs a few things to fully function:

- Register a user account, just like a human user. For details, see the [User Svc](/docs/built-in-services/user-svc).
- Register its instance in the registry so 1Backend knows where to route requests.

## A Go example

The following Go service demonstrates these steps:

- Registers itself as a user with the slug `basic-svc`
- Registers or updates its URL (`http://127.0.0.1:9111`) in the [Registry](/docs/built-in-services/registry-svc).

You may notice that the following code uses a "Go SDK"—this is simply a set of convenience functions built on top of the 1Backend API.
1Backend is language-agnostic and can be used with any language, even if no SDK is available in the repository.

The full code, including tests, is available in the [examples directory](https://github.com/1backend/1backend/tree/main/examples/go/services/basic).

```go
// <!-- INCLUDE: ../../../examples/go/services/basic/internal/basic_service.go -->
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
	"github.com/1backend/1backend/sdk/go/middlewares"
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
		dconf.TablePrefix = sdk.Id("t")
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
// <!-- /INCLUDE -->
```

Make sure to run it with the appropriate environment variables:

```sh
OB_SERVER_URL=http://127.0.0.1:11337 OB_SELF_URL=http://127.0.0.1:9111 go run main.go
```

Once it's running, you'll be able to call the 1Backend server proxy, which will forward the request to your basic service:

```sh
# 127.0.0.1:11337 here is the address of the 1Backend server
$ curl 127.0.0.1:11337/basic-svc/hello
{"hello": "world"}
```

This means you don't have to expose your basic service to the outside world—only the 1Backend server needs to be accessible.

Let's recap how the proxying works:

- Service registers an account, acquires the `basic-svc` slug.
- Service calls the 1Backend [Registry Svc](/docs/built-in-services/registry-svc) to tell the system an instance of the Basic service is available under the URL `http://127.0.0.1:9111`
- When you send a request to the 1Backend server with a path like `127.0.0.1:11337/basic-svc/hello`, the first section of the path is interpreted as a user account slug. The server checks what instances are owned by that slug and routes the request to one of those instances.

```sh
$ oo instance ls
ID                URL                     STATUS    OWNER SLUG       LAST HEARTBEAT
inst_eHFTNvAlk9   http://127.0.0.1:9111   Healthy   basic-svc     10s ago
```

## Things to understand

### Instance registration

Like most other things on the platform, service instances are owned by a user account slug. When the basic service calls [RegisterInstance](/docs/1backend/register-instance), the host will be associated with the `basic-svc` slug.

Updates to this host won’t be possible unless the caller is the basic service itself or an admin. In essence, the service becomes the owner of that URL.

This is the same ownership model used throughout the 1Backend system.
