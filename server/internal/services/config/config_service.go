/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package configservice

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/boot"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/lock"
	"github.com/1backend/1backend/sdk/go/middlewares"
	"github.com/1backend/1backend/sdk/go/service"
	types "github.com/1backend/1backend/server/internal/services/config/types"
)

type ConfigService struct {
	started    bool
	startupErr error

	clientFactory client.ClientFactory
	token         string

	lock lock.DistributedLock

	credentialStore datastore.DataStore
	configStore     datastore.DataStore

	datastoreFactory func(tableName string, instance any) (datastore.DataStore, error)

	configMutex sync.Mutex
	configs     map[string]map[string]any

	publicKey  string
	authorizer auth.Authorizer
	homeDir    string

	permissionChecker endpoint.PermissionChecker
}

func NewConfigService(
	lock lock.DistributedLock,
	authorizer auth.Authorizer,
	homeDir string,
	clientFactory client.ClientFactory,
) (*ConfigService, error) {
	cs := &ConfigService{
		lock:       lock,
		configs:    map[string]map[string]any{},
		authorizer: authorizer,
		homeDir:    homeDir,
	}

	return cs, nil
}

func (cs *ConfigService) SetClientFactory(clientFactory client.ClientFactory) {
	cs.clientFactory = clientFactory
	cs.permissionChecker = endpoint.NewPermissionChecker(
		clientFactory,
		cs.authorizer,
	)
}

func (cs *ConfigService) SetDataStoreFactory(
	datastoreFactory func(tableName string, instance any) (datastore.DataStore, error),
) {
	cs.datastoreFactory = datastoreFactory
}

func (cs *ConfigService) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/config-svc/config", middlewares.DefaultApplicator(service.Lazy(cs, func(w http.ResponseWriter, r *http.Request) {
		cs.Get(w, r)
	}))).
		Methods("OPTIONS", "GET")

	router.HandleFunc("/config-svc/config", middlewares.DefaultApplicator(service.Lazy(cs, func(w http.ResponseWriter, r *http.Request) {
		cs.Save(w, r)
	}))).
		Methods("OPTIONS", "PUT")
}

func (cs *ConfigService) LazyStart() error {
	if cs.started {
		return cs.startupErr
	}

	cs.startupErr = cs.start()
	if cs.startupErr != nil {
		return cs.startupErr
	}

	cs.started = true
	return nil
}

func (cs *ConfigService) start() error {
	if cs.datastoreFactory == nil {
		return errors.New("no datastore factory")
	}

	credentialStore, err := cs.datastoreFactory(
		"configSvcCredentials",
		&auth.Credential{},
	)
	if err != nil {
		return err
	}
	cs.credentialStore = credentialStore

	pk, _, err := cs.clientFactory.Client(client.WithToken(cs.token)).
		UserSvcAPI.GetPublicKey(context.Background()).
		Execute()
	if err != nil {
		return err
	}
	cs.publicKey = pk.PublicKey

	configStore, err := cs.datastoreFactory(
		"configSvcConfig",
		&types.Config{},
	)
	if err != nil {
		return err
	}
	cs.configStore = configStore

	ctx := context.Background()

	cs.lock.Acquire(ctx, "config-svc-start")
	defer cs.lock.Release(ctx, "config-svc-start")

	client := cs.clientFactory.Client()

	token, err := boot.RegisterServiceAccount(
		client.UserSvcAPI,
		"config-svc",
		"Config Svc",
		cs.credentialStore,
	)
	if err != nil {
		return err
	}
	cs.token = token.Token

	err = cs.registerPermits()
	if err != nil {
		return err
	}

	err = cs.loadConfigs()
	if err != nil {
		return err
	}
	return nil
}

func (cs *ConfigService) loadConfigs() error {
	cs.configMutex.Lock()
	defer cs.configMutex.Unlock()

	configIs, err := cs.configStore.Query().Find()
	if err != nil {
		return err
	}

	for _, configI := range configIs {
		config := configI.(*types.Config)

		v := map[string]any{}
		err := json.Unmarshal([]byte(config.DataJSON), &v)
		if err != nil {
			return errors.Wrap(err, "failed to parse config data")
		}

		cs.configs[config.Namespace] = v
	}

	return nil
}
