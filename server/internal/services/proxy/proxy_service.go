/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package proxyservice

import (
	"context"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"golang.org/x/sync/singleflight"

	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/boot"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/logger"
	"github.com/1backend/1backend/sdk/go/middlewares"
	"github.com/1backend/1backend/sdk/go/secrets"
	"github.com/1backend/1backend/sdk/go/service"
	proxy "github.com/1backend/1backend/server/internal/services/proxy/types"
	"github.com/1backend/1backend/server/internal/universe"
)

type contextKey string

const (
	// unexported so other packages can't overwrite it
	targetURLKey contextKey = "proxy-target"
)

type ProxyService struct {
	options *universe.Options

	started    bool
	startupErr error

	token string

	publicKey string

	credentialStore datastore.DataStore
	certStore       datastore.DataStore
	routeStore      datastore.DataStore

	routeCache sync.Map
	sf         singleflight.Group
	backendSf  singleflight.Group
	CertStore  *CertStore

	reverseProxy  *httputil.ReverseProxy
	instanceCache sync.Map
	httpClient    *http.Client
}

func NewProxyService(
	options *universe.Options,
) (*ProxyService, error) {
	// It's best to acquire the cert store here instead of in Lazy Load.
	// Simpler to reason about instead of thinking if the frontend proxy runs LazyLoad or not.

	certStore, err := options.DataStoreFactory.Create(
		"proxySvcCerts",
		&proxy.Cert{},
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create cert store")
	}

	routeStore, err := options.DataStoreFactory.Create(
		"proxySvcRoutes",
		&proxy.Route{},
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create route store")
	}

	cs := &ProxyService{
		options:    options,
		routeStore: routeStore,
		certStore:  certStore,
		CertStore: &CertStore{
			SyncCertsToFiles: options.SyncCertsToFiles,
			CertFolder:       filepath.Join(options.ConfigPath, "certs"),
			EncryptionKey:    options.SecretEncryptionKey,
			Db:               certStore,
		},
		reverseProxy: &httputil.ReverseProxy{
			Rewrite: func(pr *httputil.ProxyRequest) {
				target, ok := pr.In.Context().Value(targetURLKey).(*url.URL)
				if !ok {
					return
				}

				pr.SetURL(target)
				pr.Out.URL.Path = target.Path
				pr.Out.URL.RawQuery = target.RawQuery

				priorFor := pr.In.Header.Get("X-Forwarded-For")
				if priorFor != "" {
					pr.Out.Header.Set("X-Forwarded-For", priorFor+", "+pr.In.RemoteAddr)
				} else {
					pr.Out.Header.Set("X-Forwarded-For", pr.In.RemoteAddr)
				}

				if proto := pr.In.Header.Get("X-Forwarded-Proto"); proto != "" {
					pr.Out.Header.Set("X-Forwarded-Proto", proto)
				} else if pr.In.TLS != nil {
					pr.Out.Header.Set("X-Forwarded-Proto", "https")
				} else {
					pr.Out.Header.Set("X-Forwarded-Proto", "http")
				}

				pr.Out.Host = pr.In.Host
			},
			Transport: &http.Transport{
				MaxIdleConns:        50000,
				MaxIdleConnsPerHost: 500,
				IdleConnTimeout:     30 * time.Second,
			},
		},
		httpClient: &http.Client{
			Transport: &http.Transport{
				MaxIdleConns:        1000, // Total across all backends
				MaxIdleConnsPerHost: 200,  // Enough for concurrent spikes to one service
				IdleConnTimeout:     90 * time.Second,
				TLSHandshakeTimeout: 10 * time.Second,
			},
		},
	}

	if cs.options.SyncCertsToFiles {
		cs.syncCertsToFiles()
	}

	return cs, nil
}

func (cs *ProxyService) syncCertsToFiles() error {
	certFolder := filepath.Join(cs.options.ConfigPath, "certs")

	err := os.MkdirAll(certFolder, 0755)
	if err != nil {
		return errors.Wrap(err, "failed to create config directory")
	}

	if cs.CertStore == nil || cs.CertStore.Db == nil {
		return errors.New("cert store is not initialized")
	}

	certs, err := cs.CertStore.Db.Query().Find()
	if err != nil {
		return errors.Wrap(err, "failed to query certs")
	}

	for _, certI := range certs {
		err := cs.syncCertToFiles(certFolder, certI)
		if err != nil {
			logger.Error(
				"Failed to sync cert to disk",
				slog.String("certId", certI.GetId()),
				slog.Any("error", err),
			)
		}
	}

	return nil
}

func (cs *ProxyService) syncCertToFiles(
	certFolder string,
	certI datastore.Row,
) error {
	cert, ok := certI.(*proxy.Cert)
	if !ok {
		return errors.Errorf("expected cert type, got %T", certI)
	}

	decrypted, err := secrets.Decrypt(cert.Cert, cs.options.SecretEncryptionKey)
	if err != nil {
		return errors.Wrapf(err, "failed to decrypt cert '%s'", cert.Id)
	}

	if err := WriteCertKeyChainToFilesWithHost(
		certFolder,
		cert.Id,
		decrypted,
	); err != nil {
		return errors.Wrapf(err, "failed to write cert '%s' to disk", cert.Id)
	}

	return nil
}

func (cs *ProxyService) RegisterRoutes(router *mux.Router) {
	appl := cs.options.Middlewares

	router.HandleFunc("/proxy-svc/routes", appl(service.Lazy(cs, func(w http.ResponseWriter, r *http.Request) {
		cs.SaveRoutes(w, r)
	}))).
		Methods("OPTIONS", "PUT")

	router.HandleFunc("/proxy-svc/routes", appl(service.Lazy(cs, func(w http.ResponseWriter, r *http.Request) {
		cs.ListRoutes(w, r)
	}))).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/proxy-svc/routes", appl(service.Lazy(cs, func(w http.ResponseWriter, r *http.Request) {
		cs.DeleteRoutes(w, r)
	}))).
		Methods("OPTIONS", "DELETE")

	router.HandleFunc("/proxy-svc/certs", appl(service.Lazy(cs, func(w http.ResponseWriter, r *http.Request) {
		cs.ListCerts(w, r)
	}))).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/proxy-svc/certs", appl(service.Lazy(cs, func(w http.ResponseWriter, r *http.Request) {
		cs.SaveCerts(w, r)
	}))).
		Methods("OPTIONS", "PUT")

	tokenRefresherMiddleware := middlewares.TokenRefreshMiddleware(
		cs.options.TokenRefresher,
		cs.options.TokenAutoRefreshOff,
	)

	router.PathPrefix("/").HandlerFunc(tokenRefresherMiddleware(func(w http.ResponseWriter, r *http.Request) {
		cs.RouteBackend(w, r)
	}))
}

// RegisterFrontendRoutes is a special method for the proxy service. Unlike typical `RegisterRoutes`
// implementations that register internal service-specific routes, this method dynamically loads
// a list of frontend routes from the datastore and configures them here.
//
// It is only used if `OB_EDGE_PROXY` is set to `true`.
//
// A "frontend route" refers to traffic that will be forwarded to another port on the same machine
// or to another host altogether. This enables the proxy to handle external domain-based routing.
//
// The `RegisterRoutes` method is intended for the internal HTTP server (typically on port 11337, or
// as defined by `OB_SERVER_URL`). In contrast, `RegisterFrontendRoutes` is meant for the external
// HTTP server that listens on ports 80 (to handle ACME/Let's Encrypt challenges) and 443 (to handle
// HTTPS requests and act as the front-facing smart proxy).
func (cs *ProxyService) RegisterFrontendRoutes(router *mux.Router) {
	handler := WithCompression(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cs.RouteFrontend(w, r)
	}))

	router.PathPrefix("/").Handler(handler)
}

func (cs *ProxyService) LazyStart() error {
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

func (cs *ProxyService) start() error {
	if cs.options.DataStoreFactory == nil {
		return errors.New("no datastore factory")
	}

	credentialStore, err := cs.options.DataStoreFactory.Create(
		"proxySvcCredentials",
		&auth.Credential{},
	)
	if err != nil {
		return errors.Wrap(err, "failed to create credential store")
	}
	cs.credentialStore = credentialStore

	ctx := context.Background()
	cs.options.Lock.Acquire(ctx, "proxy-svc-start")
	defer cs.options.Lock.Release(ctx, "proxy-svc-start")

	pk, _, err := cs.options.ClientFactory.Client(client.WithToken(cs.token)).
		UserSvcAPI.GetPublicKey(context.Background()).
		Execute()
	if err != nil {
		return errors.Wrap(err, "failed to get public key")
	}
	cs.publicKey = pk.PublicKey

	client := cs.options.ClientFactory.Client()

	token, err := boot.RegisterServiceAccount(
		client.UserSvcAPI,
		"proxy-svc",
		"Proxy Svc",
		cs.credentialStore,
	)
	if err != nil {
		return errors.Wrap(err, "failed to register service account")
	}
	cs.token = token.Token

	if err := cs.registerPermits(); err != nil {
		return errors.Wrap(err, "failed to register permits")
	}

	return nil
}
