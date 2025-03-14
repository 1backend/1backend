/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package proxyservice

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/exp/rand"

	openapi "github.com/openorch/openorch/clients/go"
	sdk "github.com/openorch/openorch/sdk/go"
	"github.com/openorch/openorch/sdk/go/datastore"
	"github.com/openorch/openorch/sdk/go/lock"
	"github.com/openorch/openorch/sdk/go/logger"
)

type ProxyService struct {
	clientFactory sdk.ClientFactory
	token         string

	authorizer sdk.Authorizer

	lock      lock.DistributedLock
	publicKey string

	credentialStore  datastore.DataStore
	datastoreFactory func(tableName string, instance any) (datastore.DataStore, error)
}

func NewProxyService(
	clientFactory sdk.ClientFactory,
	authorizer sdk.Authorizer,
	lock lock.DistributedLock,
	datastoreFactory func(tableName string, instance any) (datastore.DataStore, error),
) (*ProxyService, error) {
	cs := &ProxyService{
		lock:             lock,
		datastoreFactory: datastoreFactory,
		clientFactory:    clientFactory,
		authorizer:       authorizer,
	}

	credentialStore, err := cs.datastoreFactory(
		"proxySvcCredentials",
		&sdk.Credential{},
	)
	if err != nil {
		return nil, err
	}
	cs.credentialStore = credentialStore

	return cs, nil
}

func (cs *ProxyService) Route(w http.ResponseWriter, r *http.Request) {
	logger.Debug("Proxying", slog.String("path", r.URL.Path))

	// @todo cache?

	serviceSlug := getServiceSlug(r)

	rsp, _, err := cs.clientFactory.Client(sdk.WithToken(cs.token)).
		RegistrySvcAPI.ListInstances(context.Background()).
		Slug(serviceSlug).
		Execute()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errors.Wrap(err, "error listing instances").Error()))
		return
	}

	if len(rsp.Instances) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("not found"))
		return
	}

	healtyInstances := make([]openapi.RegistrySvcInstance, 0, len(rsp.Instances))
	for _, instance := range rsp.Instances {
		if instance.Status == openapi.InstanceStatusHealthy {
			healtyInstances = append(healtyInstances, instance)
		}
	}

	if len(healtyInstances) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("not healthy instance found"))
		return
	}

	randomIndex := rand.Intn(len(healtyInstances))
	instance := healtyInstances[randomIndex]

	uri := strings.TrimSuffix(instance.Url, "/") + "/" + strings.TrimLeft(r.URL.Path, "/")

	req, err := http.NewRequest(r.Method, uri, r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errors.Wrap(err, "error creating proxy request").Error()))
		return
	}

	req.Header = r.Header

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errors.Wrap(err, "error proxying request").Error()))
		return
	}

	defer resp.Body.Close()

	for k, v := range resp.Header {
		w.Header().Set(k, v[0])
	}

	logger.Debug("Proxy request returned", slog.Int("statusCode", resp.StatusCode))

	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

// gets service slug from http request path
// eg. /my-svc/my-endpoint -> my-svc
func getServiceSlug(r *http.Request) string {
	cleanedPath := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(cleanedPath, "/")
	if len(parts) > 0 && parts[0] != "" {
		return parts[0]
	}
	return ""
}

func (cs *ProxyService) Start() error {
	if cs.datastoreFactory == nil {
		return errors.New("no datastore factory")
	}

	ctx := context.Background()
	cs.lock.Acquire(ctx, "proxy-svc-start")
	defer cs.lock.Release(ctx, "proxy-svc-start")

	pk, _, err := cs.clientFactory.Client(sdk.WithToken(cs.token)).
		UserSvcAPI.GetPublicKey(context.Background()).
		Execute()
	if err != nil {
		return err
	}
	cs.publicKey = *pk.PublicKey

	client := cs.clientFactory.Client()

	token, err := sdk.RegisterServiceAccount(
		client.UserSvcAPI,
		"proxy-svc",
		"Proxy Svc",
		cs.credentialStore,
	)
	if err != nil {
		return err
	}
	cs.token = token

	return nil
}
