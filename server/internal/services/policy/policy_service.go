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
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/lock"
	"github.com/1backend/1backend/sdk/go/middlewares"
	"github.com/1backend/1backend/sdk/go/service"
	"github.com/gorilla/mux"

	policytypes "github.com/1backend/1backend/server/internal/services/policy/types"
)

type PolicyService struct {
	started    bool
	startupErr error

	clientFactory    client.ClientFactory
	datastoreFactory func(tableName string, instance any) (datastore.DataStore, error)
	token            string

	lock lock.DistributedLock

	instancesStore  datastore.DataStore
	credentialStore datastore.DataStore

	instances []*policytypes.Instance

	rateLimiters sync.Map // Map to store rate limiters
	mutex        sync.Mutex
}

func NewPolicyService(
	clientFactory client.ClientFactory,
	lock lock.DistributedLock,
	datastoreFactory func(tableName string, instance any) (datastore.DataStore, error),
) (*PolicyService, error) {

	service := &PolicyService{
		clientFactory:    clientFactory,
		datastoreFactory: datastoreFactory,
		lock:             lock,
	}

	return service, nil
}

func (ps *PolicyService) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/policy-svc/check", middlewares.DefaultApplicator(service.Lazy(ps, func(w http.ResponseWriter, r *http.Request) {
		ps.Check(w, r)
	}))).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/policy-svc/instance/{instanceId}", middlewares.DefaultApplicator(service.Lazy(ps, func(w http.ResponseWriter, r *http.Request) {
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
	instancesStore, err := ps.datastoreFactory(
		"policySvcInstances",
		&policytypes.Instance{},
	)
	if err != nil {
		return err
	}
	ps.instancesStore = instancesStore

	credentialStore, err := ps.datastoreFactory(
		"policySvcCredentials",
		&auth.Credential{},
	)
	if err != nil {
		return err
	}
	ps.credentialStore = credentialStore

	ctx := context.Background()
	ps.lock.Acquire(ctx, "policy-svc-start")
	defer ps.lock.Release(ctx, "policy-svc-start")

	instances, err := ps.instancesStore.Query().Find()
	if err != nil {
		return err
	}

	for _, instanceI := range instances {
		instance := instanceI.(*policytypes.Instance)
		ps.instances = append(ps.instances, instance)
	}

	token, err := boot.RegisterServiceAccount(
		ps.clientFactory.Client().UserSvcAPI,
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
