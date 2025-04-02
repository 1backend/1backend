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

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/boot"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/lock"
	"github.com/1backend/1backend/sdk/go/logger"
)

type ProxyService struct {
	clientFactory client.ClientFactory
	token         string

	authorizer auth.Authorizer

	lock      lock.DistributedLock
	publicKey string

	credentialStore  datastore.DataStore
	datastoreFactory func(tableName string, instance any) (datastore.DataStore, error)
}

func NewProxyService(
	clientFactory client.ClientFactory,
	authorizer auth.Authorizer,
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
		&auth.Credential{},
	)
	if err != nil {
		return nil, err
	}
	cs.credentialStore = credentialStore

	return cs, nil
}

func (cs *ProxyService) Route(w http.ResponseWriter, r *http.Request) {
	logger.Debug("Proxying",
		slog.String("path", r.URL.Path),
		slog.String("method", r.Method),
	)

	// @todo cache?

	serviceSlug := getServiceSlug(r)

	rsp, _, err := cs.clientFactory.Client(client.WithToken(cs.token)).
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

	// Prioritize healthy instances
	selectedInstances := make([]openapi.RegistrySvcInstance, 0, len(rsp.Instances))
	for _, instance := range rsp.Instances {
		if instance.Status == openapi.InstanceStatusHealthy {
			selectedInstances = append(selectedInstances, instance)
		}
	}

	// But fall back to any instances if none found
	if len(selectedInstances) == 0 {
		selectedInstances = rsp.Instances
	}

	randomIndex := rand.Intn(len(selectedInstances))
	instance := selectedInstances[randomIndex]

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
		// Skip "Content-Length" and "Transfer-Encoding" to prevent HTTP response errors.
		// - Go automatically sets "Content-Length" based on the actual body size.
		// - "Transfer-Encoding: chunked" conflicts with "Content-Length", so we avoid copying it.
		if k == "Content-Length" || k == "Transfer-Encoding" {
			continue
		}
		for _, vv := range v {
			w.Header().Add(k, vv)
		}
	}

	logger.Debug("Proxy request returned",
		slog.Int("statusCode", resp.StatusCode),
	)

	w.WriteHeader(resp.StatusCode)

	// Unfortunately io.Copy does not work here and causes "invalid write result"

	//_, err = io.Copy(w, resp.Body)
	//if err != nil {
	//	logger.Error("Error proxying body",
	//		slog.String("error", err.Error()),
	//	)
	//}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errors.Wrap(err, "failed to read body while proxying").Error()))
		return
	}

	_, err = w.Write(body)
	if err != nil {
		logger.Error("Error writing response body", slog.String("error", err.Error()))
	}

	if flusher, ok := w.(http.Flusher); ok {
		flusher.Flush()
	}
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

	pk, _, err := cs.clientFactory.Client(client.WithToken(cs.token)).
		UserSvcAPI.GetPublicKey(context.Background()).
		Execute()
	if err != nil {
		return err
	}
	cs.publicKey = pk.PublicKey

	client := cs.clientFactory.Client()

	token, err := boot.RegisterServiceAccount(
		client.UserSvcAPI,
		"proxy-svc",
		"Proxy Svc",
		cs.credentialStore,
	)
	if err != nil {
		return err
	}
	cs.token = token.Token

	return nil
}
