/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package imageservice

import (
	"context"
	"net/http"
	"os"
	"path"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/boot"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/lock"
	"github.com/1backend/1backend/sdk/go/middlewares"
	"github.com/1backend/1backend/sdk/go/service"
)

type ImageService struct {
	homeDir       string
	clientFactory client.ClientFactory

	started    bool
	startupErr error

	lock lock.DistributedLock

	token string

	imageCacheFolder string

	publicKey string

	credentialStore   datastore.DataStore
	datastoreFactory  func(tableName string, instance any) (datastore.DataStore, error)
	permissionChecker endpoint.PermissionChecker
}

func NewImageService(
	clientFactory client.ClientFactory,
	homeDir string,
	datastoreFactory func(tableName string, instance any) (datastore.DataStore, error),
	lock lock.DistributedLock,
) (*ImageService, error) {
	cs := &ImageService{
		homeDir:          homeDir,
		datastoreFactory: datastoreFactory,
		lock:             lock,
		clientFactory:    clientFactory,
	}

	return cs, nil
}

func (cs *ImageService) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/image-svc/serve/upload/{fileId}", middlewares.DefaultApplicator(service.Lazy(cs, func(w http.ResponseWriter, r *http.Request) {
		cs.ServeUploadedImage(w, r)
	}))).
		Methods("OPTIONS", "GET")
}

func (cs *ImageService) LazyStart() error {
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

func (cs *ImageService) start() error {
	if cs.datastoreFactory == nil {
		return errors.New("no datastore factory")
	}
	if cs.homeDir == "" {
		return errors.New("no home dir")
	}

	credentialStore, err := cs.datastoreFactory(
		"imageSvcCredentials",
		&auth.Credential{},
	)
	if err != nil {
		return err
	}
	cs.credentialStore = credentialStore

	imageCacheFolder := path.Join(cs.homeDir, "image-cache")
	err = os.MkdirAll(imageCacheFolder, 0700)
	if err != nil {
		return err
	}
	cs.imageCacheFolder = imageCacheFolder

	ctx := context.Background()
	cs.lock.Acquire(ctx, "image-svc-start")
	defer cs.lock.Release(ctx, "image-svc-start")

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
		"image-svc",
		"Image Svc",
		cs.credentialStore,
	)
	if err != nil {
		return err
	}
	cs.token = token.Token

	return nil
}
