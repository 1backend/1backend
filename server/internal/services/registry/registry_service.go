/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package registryservice

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/boot"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/lock"
	registry "github.com/1backend/1backend/server/internal/services/registry/types"
)

type RegistryService struct {
	publicKey     string
	clientFactory client.ClientFactory
	token         string

	URL              string
	AvailabilityZone string
	Region           string

	lock lock.DistributedLock

	credentialStore datastore.DataStore
	definitionStore datastore.DataStore
	instanceStore   datastore.DataStore
	nodeStore       datastore.DataStore

	triggerChan chan struct{}
	nodeId      string
}

func NewRegistryService(
	address string,
	az string,
	region string,
	clientFactory client.ClientFactory,
	lock lock.DistributedLock,
	datastoreFactory func(tableName string, instance any) (datastore.DataStore, error),
	nodeId string,
) (*RegistryService, error) {

	nodeUrl := address

	var err error

	if nodeUrl == "" {
		nodeUrl, err = os.Hostname()
		if err != nil {
			return nil, err
		}
		nodeUrl = fmt.Sprintf("%v:%v", nodeUrl, "58231")
	}

	if !strings.HasPrefix(nodeUrl, "http") {
		nodeUrl = "http://" + nodeUrl
	}

	credentialStore, err := datastoreFactory(
		"registrySvcCredentials",
		&auth.Credential{},
	)
	if err != nil {
		return nil, err
	}
	instanceStore, err := datastoreFactory(
		"registrySvcInstances",
		&registry.Instance{},
	)
	if err != nil {
		return nil, err
	}
	definitionStore, err := datastoreFactory(
		"registrySvcDefinitions",
		&registry.Definition{},
	)
	if err != nil {
		return nil, err
	}
	nodeStore, err := datastoreFactory("registrySvcNodes", &registry.Node{})
	if err != nil {
		return nil, err
	}

	service := &RegistryService{
		URL:              nodeUrl,
		clientFactory:    clientFactory,
		lock:             lock,
		credentialStore:  credentialStore,
		definitionStore:  definitionStore,
		instanceStore:    instanceStore,
		nodeStore:        nodeStore,
		AvailabilityZone: az,
		Region:           region,
		nodeId:           nodeId,

		triggerChan: make(chan struct{}),
	}

	return service, nil
}

func (ns *RegistryService) Start() error {
	ns.heartbeatCycle()

	go ns.nodeHeartbeat()
	go ns.instanceScan()

	ctx := context.Background()
	ns.lock.Acquire(ctx, "registry-svc-start")
	defer ns.lock.Release(ctx, "registry-svc-start")

	token, err := boot.RegisterServiceAccount(
		ns.clientFactory.Client().UserSvcAPI,
		"registry-svc",
		"Registry Svc",
		ns.credentialStore,
	)
	if err != nil {
		return err
	}
	ns.token = token.Token

	pk, _, err := ns.clientFactory.Client(client.WithToken(ns.token)).
		UserSvcAPI.GetPublicKey(context.Background()).
		Execute()
	if err != nil {
		return err
	}
	ns.publicKey = pk.PublicKey

	return ns.registerPermissions()
}
