/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package deployservice

import (
	"context"
	"net/http"

	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/boot"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/lock"
	"github.com/1backend/1backend/sdk/go/middlewares"
	deploy "github.com/1backend/1backend/server/internal/services/deploy/types"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

type DeployService struct {
	clientFactory client.ClientFactory
	token         string

	lock lock.DistributedLock

	credentialStore datastore.DataStore
	deploymentStore datastore.DataStore

	triggerChan       chan struct{}
	triggerOnly       bool
	permissionChecker endpoint.PermissionChecker
}

func NewDeployService(
	clientFactory client.ClientFactory,
	lock lock.DistributedLock,
	datastoreFactory func(tableName string, instance any) (datastore.DataStore, error),
	triggerOnly bool,
	authorizer auth.Authorizer,
) (*DeployService, error) {

	credentialStore, err := datastoreFactory(
		"deploySvcCredentials",
		&auth.Credential{},
	)
	if err != nil {
		return nil, err
	}
	deploymentStore, err := datastoreFactory(
		"deploySvcDeployments",
		&deploy.Deployment{},
	)
	if err != nil {
		return nil, err
	}

	service := &DeployService{
		clientFactory: clientFactory,

		lock:            lock,
		credentialStore: credentialStore,
		deploymentStore: deploymentStore,

		triggerChan: make(chan struct{}),

		triggerOnly: triggerOnly,
		permissionChecker: endpoint.NewPermissionChecker(
			clientFactory,
			authorizer,
		),
	}

	return service, nil
}

func (ds *DeployService) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/deploy-svc/deployment", middlewares.DefaultApplicator(func(w http.ResponseWriter, r *http.Request) {
		ds.SaveDeployment(w, r)
	})).
		Methods("OPTIONS", "PUT")

	router.HandleFunc("/deploy-svc/deployments", middlewares.DefaultApplicator(func(w http.ResponseWriter, r *http.Request) {
		ds.ListDeployments(w, r)
	})).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/deploy-svc/deployment", middlewares.DefaultApplicator(func(w http.ResponseWriter, r *http.Request) {
		ds.DeleteDeployment(w, r)
	})).
		Methods("OPTIONS", "DELETE")
}

func (ns *DeployService) Start() error {
	go ns.loop(ns.triggerOnly)

	return nil
}

func (cs *DeployService) getToken() (string, error) {
	if cs.token != "" {
		return cs.token, nil
	}

	ctx := context.Background()
	cs.lock.Acquire(ctx, "deploy-svc-start")
	defer cs.lock.Release(ctx, "deploy-svc-start")

	token, err := boot.RegisterServiceAccount(
		cs.clientFactory.Client().UserSvcAPI,
		"deploy-svc",
		"Deploy Svc",
		cs.credentialStore,
	)
	if err != nil {
		return "", err
	}
	cs.token = token.Token

	err = cs.registerPermits()
	if err != nil {
		return "", errors.Wrap(err, "failed to register permissions")
	}

	return cs.token, nil
}
