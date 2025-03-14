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

	sdk "github.com/openorch/openorch/sdk/go"
	"github.com/openorch/openorch/sdk/go/datastore"
	"github.com/openorch/openorch/sdk/go/lock"
	deploy "github.com/openorch/openorch/server/internal/services/deploy/types"
)

type DeployService struct {
	clientFactory sdk.ClientFactory
	token         string

	lock lock.DistributedLock

	credentialStore datastore.DataStore
	deploymentStore datastore.DataStore

	triggerChan chan struct{}
	triggerOnly bool
}

func NewDeployService(
	clientFactory sdk.ClientFactory,
	lock lock.DistributedLock,
	datastoreFactory func(tableName string, instance any) (datastore.DataStore, error),
	triggerOnly bool,
) (*DeployService, error) {

	credentialStore, err := datastoreFactory(
		"deploySvcCredentials",
		&sdk.Credential{},
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
	}

	return service, nil
}

func (ns *DeployService) Start() error {
	ctx := context.Background()
	ns.lock.Acquire(ctx, "deploy-svc-start")
	defer ns.lock.Release(ctx, "deploy-svc-start")

	token, err := sdk.RegisterServiceAccount(
		ns.clientFactory.Client().UserSvcAPI,
		"deploy-svc",
		"Deploy Svc",
		ns.credentialStore,
	)
	if err != nil {
		return err
	}

	ns.token = token

	go ns.loop(ns.triggerOnly)

	return ns.registerPermissions()
}
