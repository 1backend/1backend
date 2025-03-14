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
	"sync"

	sdk "github.com/openorch/openorch/sdk/go"
	"github.com/openorch/openorch/sdk/go/datastore"
	"github.com/openorch/openorch/sdk/go/lock"

	policytypes "github.com/openorch/openorch/server/internal/services/policy/types"
)

type PolicyService struct {
	clientFactory sdk.ClientFactory
	token         string

	lock lock.DistributedLock

	instancesStore  datastore.DataStore
	credentialStore datastore.DataStore

	instances []*policytypes.Instance

	rateLimiters sync.Map // Map to store rate limiters
	mutex        sync.Mutex
}

func NewPolicyService(
	clientFactory sdk.ClientFactory,
	lock lock.DistributedLock,
	datastoreFactory func(tableName string, instance any) (datastore.DataStore, error),
) (*PolicyService, error) {

	instancesStore, err := datastoreFactory(
		"policySvcInstances",
		&policytypes.Instance{},
	)
	if err != nil {
		return nil, err
	}

	credentialStore, err := datastoreFactory(
		"policySvcCredentials",
		&sdk.Credential{},
	)
	if err != nil {
		return nil, err
	}

	service := &PolicyService{
		clientFactory: clientFactory,

		lock: lock,

		instancesStore:  instancesStore,
		credentialStore: credentialStore,
	}

	return service, nil
}

func (cs *PolicyService) Start() error {
	ctx := context.Background()
	cs.lock.Acquire(ctx, "policy-svc-start")
	defer cs.lock.Release(ctx, "policy-svc-start")

	instances, err := cs.instancesStore.Query().Find()
	if err != nil {
		return err
	}

	for _, instanceI := range instances {
		instance := instanceI.(*policytypes.Instance)
		cs.instances = append(cs.instances, instance)
	}

	token, err := sdk.RegisterServiceAccount(
		cs.clientFactory.Client().UserSvcAPI,
		"policy-svc",
		"Policy Svc",
		cs.credentialStore,
	)
	if err != nil {
		return err
	}
	cs.token = token

	return cs.registerPermissions()
}
