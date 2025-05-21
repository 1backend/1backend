/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package policyservice

import (
	"context"
	"net/http"
	"sync"

	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/boot"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/service"
	"github.com/gorilla/mux"

	policytypes "github.com/1backend/1backend/server/internal/services/policy/types"
	"github.com/1backend/1backend/server/internal/universe"
)

type PolicyService struct {
	options *universe.Options

	started    bool
	startupErr error

	token string

	instancesStore  datastore.DataStore
	credentialStore datastore.DataStore

	instances []*policytypes.Instance

	rateLimiters sync.Map // Map to store rate limiters
	mutex        sync.Mutex
}

func NewPolicyService(
	options *universe.Options,
) (*PolicyService, error) {

	service := &PolicyService{
		options: options,
	}

	return service, nil
}

func (ps *PolicyService) RegisterRoutes(router *mux.Router) {
	appl := ps.options.Middlewares

	router.HandleFunc("/policy-svc/check", appl(service.Lazy(ps, func(w http.ResponseWriter, r *http.Request) {
		ps.Check(w, r)
	}))).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/policy-svc/instance/{instanceId}", appl(service.Lazy(ps, func(w http.ResponseWriter, r *http.Request) {
		ps.UpsertInstance(w, r)
	}))).
		Methods("OPTIONS", "PUT")
}

func (cs *PolicyService) LazyStart() error {
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

func (ps *PolicyService) start() error {
	instancesStore, err := ps.options.DataStoreFactory.Create(
		"policySvcInstances",
		&policytypes.Instance{},
	)
	if err != nil {
		return err
	}
	ps.instancesStore = instancesStore

	credentialStore, err := ps.options.DataStoreFactory.Create(
		"policySvcCredentials",
		&auth.Credential{},
	)
	if err != nil {
		return err
	}
	ps.credentialStore = credentialStore

	ctx := context.Background()
	ps.options.Lock.Acquire(ctx, "policy-svc-start")
	defer ps.options.Lock.Release(ctx, "policy-svc-start")

	instances, err := ps.instancesStore.Query().Find()
	if err != nil {
		return err
	}

	for _, instanceI := range instances {
		instance := instanceI.(*policytypes.Instance)
		ps.instances = append(ps.instances, instance)
	}

	token, err := boot.RegisterServiceAccount(
		ps.options.ClientFactory.Client().UserSvcAPI,
		"policy-svc",
		"Policy Svc",
		ps.credentialStore,
	)
	if err != nil {
		return err
	}
	ps.token = token.Token

	return ps.registerPermits()
}
