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
	"github.com/1backend/1backend/sdk/go/service"
	types "github.com/1backend/1backend/server/internal/services/config/types"
	"github.com/1backend/1backend/server/internal/universe"
)

type ConfigService struct {
	options    *universe.Options
	started    bool
	startupErr error

	token string

	credentialStore datastore.DataStore
	configStore     datastore.DataStore

	configMutex sync.Mutex
	configs     map[string]map[string]any

	publicKey string
}

func NewConfigService(
	options *universe.Options,
) (*ConfigService, error) {

	cs := &ConfigService{
		options: options,
		configs: map[string]map[string]any{},
	}

	return cs, nil
}

func (cs *ConfigService) RegisterRoutes(router *mux.Router) {
	appl := cs.options.Middlewares

	router.HandleFunc("/config-svc/config", appl(service.Lazy(cs, func(w http.ResponseWriter, r *http.Request) {
		cs.Get(w, r)
	}))).
		Methods("OPTIONS", "GET")

	router.HandleFunc("/config-svc/config", appl(service.Lazy(cs, func(w http.ResponseWriter, r *http.Request) {
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
	if cs.options.DataStoreFactory == nil {
		return errors.New("no datastore factory")
	}

	credentialStore, err := cs.options.DataStoreFactory.Create(
		"configSvcCredentials",
		&auth.Credential{},
	)
	if err != nil {
		return err
	}
	cs.credentialStore = credentialStore

	pk, _, err := cs.options.ClientFactory.Client(client.WithToken(cs.token)).
		UserSvcAPI.GetPublicKey(context.Background()).
		Execute()
	if err != nil {
		return err
	}
	cs.publicKey = pk.PublicKey

	configStore, err := cs.options.DataStoreFactory.Create(
		"configSvcConfig",
		&types.Config{},
	)
	if err != nil {
		return err
	}
	cs.configStore = configStore

	ctx := context.Background()

	cs.options.Lock.Acquire(ctx, "config-svc-start")
	defer cs.options.Lock.Release(ctx, "config-svc-start")

	client := cs.options.ClientFactory.Client()

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
