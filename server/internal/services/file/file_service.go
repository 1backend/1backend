/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package fileservice

import (
	"context"
	"net/http"
	"os"
	"path"
	"sync"

	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/boot"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/service"
	types "github.com/1backend/1backend/server/internal/services/file/types"
	"github.com/1backend/1backend/server/internal/universe"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

type FileService struct {
	options *universe.Options
	token   string

	uploadFolder   string
	downloadFolder string

	mutex sync.Mutex

	downloadStore datastore.DataStore
	uploadStore   datastore.DataStore

	// for testing purposes
	SyncDownloads bool

	credentialStore datastore.DataStore

	nodeId string
}

func NewFileService(
	options *universe.Options,
) (*FileService, error) {
	ret := &FileService{
		options: options,
	}

	return ret, nil
}

func (fs *FileService) RegisterRoutes(router *mux.Router) {
	appl := fs.options.Middlewares

	router.HandleFunc("/file-svc/download", appl(service.Lazy(fs, func(w http.ResponseWriter, r *http.Request) {
		fs.Download(w, r)
	}))).
		Methods("OPTIONS", "PUT")

	router.HandleFunc("/file-svc/download/{url}/pause", appl(service.Lazy(fs, func(w http.ResponseWriter, r *http.Request) {
		fs.PauseDownload(w, r)
	}))).
		Methods("OPTIONS", "PUT")

	router.HandleFunc("/file-svc/download/{url}", appl(service.Lazy(fs, func(w http.ResponseWriter, r *http.Request) {
		fs.GetDownload(w, r)
	}))).
		Methods("OPTIONS", "GET")

	router.HandleFunc("/file-svc/downloads", appl(service.Lazy(fs, func(w http.ResponseWriter, r *http.Request) {
		fs.ListDownloads(w, r)
	}))).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/file-svc/upload", appl(service.Lazy(fs, func(w http.ResponseWriter, r *http.Request) {
		fs.UploadFile(w, r)
	}))).
		Methods("OPTIONS", "PUT")

	router.HandleFunc("/file-svc/uploads", appl(service.Lazy(fs, func(w http.ResponseWriter, r *http.Request) {
		fs.ListUploads(w, r)
	}))).
		Methods("OPTIONS", "POST")

	// @todo
	// Investigate why SkipLock is needed here.
	// I placed it here because the serve proxy tests were deadlocking.
	// Not sure why though as they are not routing to the same node (themselves),
	// but to an other node.

	router.HandleFunc("/file-svc/serve/upload/{fileId}", appl(service.Lazy(fs, func(w http.ResponseWriter, r *http.Request) {
		fs.ServeUpload(w, r)
	}, service.WithSkipLock()))).
		Methods("OPTIONS", "GET")

	router.HandleFunc("/file-svc/serve/download/{url}", appl(service.Lazy(fs, func(w http.ResponseWriter, r *http.Request) {
		fs.ServeDownload(w, r)
	}, service.WithSkipLock()))).
		Methods("OPTIONS", "GET")
}

func (fs *FileService) Start() error {
	credentialStore, err := fs.options.DataStoreFactory.Create(
		"fileSvcCredentials",
		&auth.Credential{},
	)
	if err != nil {
		return err
	}
	fs.credentialStore = credentialStore

	uploadFolder := path.Join(fs.options.HomeDir, "uploads")
	err = os.MkdirAll(uploadFolder, 0700)
	if err != nil {
		return err
	}
	fs.uploadFolder = uploadFolder

	downloadFolder := path.Join(fs.options.HomeDir, "downloads")
	err = os.MkdirAll(downloadFolder, 0700)
	if err != nil {
		return err
	}
	fs.downloadFolder = downloadFolder

	downloadStore, err := fs.options.DataStoreFactory.Create(
		"fileSvcDownloads",
		&types.InternalDownload{},
	)
	if err != nil {
		return err
	}
	fs.downloadStore = downloadStore

	uploadStore, err := fs.options.DataStoreFactory.Create(
		"fileSvcUploads",
		&types.Upload{},
	)
	if err != nil {
		return err
	}
	fs.uploadStore = uploadStore

	downloads, err := fs.downloadStore.Query(
		datastore.Equals([]string{"status"},
			types.DownloadStatusInProgress,
		)).Find()
	if err != nil {
		return nil
	}

	for _, downloadI := range downloads {
		download := downloadI.(*types.InternalDownload)

		if download.Status == types.DownloadStatusInProgress {
			err = fs.download(context.Background(), download.URL, path.Dir(download.FilePath))
			if err != nil {
				return err
			}
		}
	}

	return err
}

func (fs *FileService) LazyStart() error {
	_, err := fs.getToken()
	if err != nil {
		return errors.Wrap(err, "failed to get token")
	}

	return nil
}

func (fs *FileService) getToken() (string, error) {
	if fs.token != "" {
		return fs.token, nil
	}

	ctx := context.Background()
	fs.options.Lock.Acquire(ctx, "file-svc-start")
	defer fs.options.Lock.Release(ctx, "file-svc-start")

	token, err := boot.RegisterServiceAccount(
		fs.options.ClientFactory.Client().UserSvcAPI,
		"file-svc",
		"File Svc",
		fs.credentialStore,
	)
	if err != nil {
		return "", err
	}
	fs.token = token.Token

	err = fs.registerPermits()
	if err != nil {
		return "", errors.Wrap(err, "failed to register permissions")
	}

	return fs.token, nil
}

func (fs *FileService) getDownload(url string) (*types.InternalDownload, bool) {
	fs.mutex.Lock()
	defer fs.mutex.Unlock()

	downloadIs, err := fs.downloadStore.Query(
		datastore.Equals([]string{"url"},
			url,
		)).Find()
	if err != nil {
		return nil, false
	}

	if len(downloadIs) == 0 {
		return nil, false
	}

	return downloadIs[0].(*types.InternalDownload), true
}
