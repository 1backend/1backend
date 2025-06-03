/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
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
	"github.com/1backend/1backend/sdk/go/service"
	"github.com/1backend/1backend/server/internal/universe"
)

type ImageService struct {
	options *universe.Options

	started    bool
	startupErr error

	token string

	imageCacheFolder string

	publicKey string

	credentialStore datastore.DataStore
}

func NewImageService(
	options *universe.Options,
) (*ImageService, error) {
	cs := &ImageService{
		options: options,
	}

	return cs, nil
}

func (cs *ImageService) RegisterRoutes(router *mux.Router) {
	appl := cs.options.Middlewares

	router.HandleFunc("/image-svc/serve/upload/{fileId}", appl(service.Lazy(cs, func(w http.ResponseWriter, r *http.Request) {
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
	if cs.options.HomeDir == "" {
		return errors.New("no home dir")
	}

	credentialStore, err := cs.options.DataStoreFactory.Create(
		"imageSvcCredentials",
		&auth.Credential{},
	)
	if err != nil {
		return err
	}
	cs.credentialStore = credentialStore

	imageCacheFolder := path.Join(cs.options.HomeDir, "image-cache")
	err = os.MkdirAll(imageCacheFolder, 0700)
	if err != nil {
		return err
	}
	cs.imageCacheFolder = imageCacheFolder

	ctx := context.Background()
	cs.options.Lock.Acquire(ctx, "image-svc-start")
	defer cs.options.Lock.Release(ctx, "image-svc-start")

	pk, _, err := cs.options.ClientFactory.Client(client.WithToken(cs.token)).
		UserSvcAPI.GetPublicKey(context.Background()).
		Execute()
	if err != nil {
		return err
	}
	cs.publicKey = pk.PublicKey

	client := cs.options.ClientFactory.Client()

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
