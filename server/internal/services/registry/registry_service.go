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
	"net/http"
	"os"
	"strings"

	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/boot"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/lock"
	"github.com/1backend/1backend/sdk/go/middlewares"
	registry "github.com/1backend/1backend/server/internal/services/registry/types"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

type RegistryService struct {
	publicKey  string
	token      string
	authorizer auth.Authorizer

	clientFactory client.ClientFactory

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

	permissionChecker endpoint.PermissionChecker
}

func NewRegistryService(
	address string,
	az string,
	region string,
	clientFactory client.ClientFactory,
	lock lock.DistributedLock,
	datastoreFactory func(tableName string, instance any) (datastore.DataStore, error),
	nodeId string,
	authorizer auth.Authorizer,
) (*RegistryService, error) {

	nodeUrl := address

	var err error

	if nodeUrl == "" {
		nodeUrl, err = os.Hostname()
		if err != nil {
			return nil, err
		}
		nodeUrl = fmt.Sprintf("%v:%v", nodeUrl, "11337")
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
		authorizer:       authorizer,

		triggerChan: make(chan struct{}),
		permissionChecker: endpoint.NewPermissionChecker(
			clientFactory,
			authorizer,
		),
	}

	return service, nil
}

func (rs *RegistryService) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/registry-svc/node/self", middlewares.DefaultApplicator(func(w http.ResponseWriter, r *http.Request) {
		rs.NodeSelf(w, r)
	})).
		Methods("OPTIONS", "GET")

	router.HandleFunc("/registry-svc/nodes", middlewares.DefaultApplicator(func(w http.ResponseWriter, r *http.Request) {
		rs.ListNodes(w, r)
	})).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/registry-svc/instances", middlewares.DefaultApplicator(func(w http.ResponseWriter, r *http.Request) {
		rs.ListInstances(w, r)
	})).
		Methods("OPTIONS", "GET")

	router.HandleFunc("/registry-svc/definitions", middlewares.DefaultApplicator(func(w http.ResponseWriter, r *http.Request) {
		rs.ListDefinitions(w, r)
	})).
		Methods("OPTIONS", "GET")

	router.HandleFunc("/registry-svc/instance", middlewares.DefaultApplicator(func(w http.ResponseWriter, r *http.Request) {
		rs.RegisterInstance(w, r)
	})).
		Methods("OPTIONS", "PUT")

	router.HandleFunc("/registry-svc/definition", middlewares.DefaultApplicator(func(w http.ResponseWriter, r *http.Request) {
		rs.SaveDefinition(w, r)
	})).
		Methods("OPTIONS", "PUT")

	router.HandleFunc("/registry-svc/instance/{id}", middlewares.DefaultApplicator(func(w http.ResponseWriter, r *http.Request) {
		rs.RemoveInstance(w, r)
	})).
		Methods("OPTIONS", "DELETE")

	router.HandleFunc("/registry-svc/definition/{id}", middlewares.DefaultApplicator(func(w http.ResponseWriter, r *http.Request) {
		rs.DeleteDefinition(w, r)
	})).
		Methods("OPTIONS", "DELETE")

	router.HandleFunc("/registry-svc/node/{url}", middlewares.DefaultApplicator(func(w http.ResponseWriter, r *http.Request) {
		rs.DeleteNode(w, r)
	})).
		Methods("OPTIONS", "DELETE")
}

func (ns *RegistryService) Start() error {
	ns.heartbeatCycle()

	go ns.nodeHeartbeat()
	go ns.instanceScan()

	// Only to trigger registration of permissions
	_, err := ns.getToken()
	if err != nil {
		return errors.Wrap(err, "failed to get token")
	}

	return nil
}

func (cs *RegistryService) getToken() (string, error) {
	if cs.token != "" {
		return cs.token, nil
	}

	ctx := context.Background()
	cs.lock.Acquire(ctx, "registry-svc-start")
	defer cs.lock.Release(ctx, "registry-svc-start")

	token, err := boot.RegisterServiceAccount(
		cs.clientFactory.Client().UserSvcAPI,
		"registry-svc",
		"Registry Svc",
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

func (ns *RegistryService) getPublicKey() (string, error) {
	if ns.publicKey != "" {
		return ns.publicKey, nil
	}

	pk, _, err := ns.clientFactory.Client(client.WithToken(ns.token)).
		UserSvcAPI.GetPublicKey(context.Background()).
		Execute()
	if err != nil {
		return "", errors.Wrap(err, "failed to get public key")
	}
	ns.publicKey = pk.PublicKey

	return ns.publicKey, nil
}
