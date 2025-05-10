/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
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
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/lock"
	"github.com/1backend/1backend/sdk/go/middlewares"
	"github.com/1backend/1backend/sdk/go/service"
	types "github.com/1backend/1backend/server/internal/services/file/types"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

type FileService struct {
	clientFactory    client.ClientFactory
	datastoreFactory func(tableName string, instance any) (datastore.DataStore, error)
	homeDir          string

	token string

	lock lock.DistributedLock

	uploadFolder   string
	downloadFolder string

	mutex sync.Mutex

	downloadStore datastore.DataStore
	uploadStore   datastore.DataStore

	// for testing purposes
	SyncDownloads bool

	credentialStore datastore.DataStore

	nodeId            string
	permissionChecker endpoint.PermissionChecker
}

func NewFileService(
	clientFactory client.ClientFactory,
	lock lock.DistributedLock,
	datastoreFactory func(tableName string, instance any) (datastore.DataStore, error),
	homeDir string,
) (*FileService, error) {
	ret := &FileService{
		clientFactory:     clientFactory,
		homeDir:           homeDir,
		datastoreFactory:  datastoreFactory,
		lock:              lock,
		permissionChecker: endpoint.NewPermissionChecker(clientFactory),
	}

	return ret, nil
}

func (fs *FileService) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/file-svc/download", middlewares.DefaultApplicator(service.Lazy(fs, func(w http.ResponseWriter, r *http.Request) {
		fs.Download(w, r)
	}))).
		Methods("OPTIONS", "PUT")

	router.HandleFunc("/file-svc/download/{url}/pause", middlewares.DefaultApplicator(service.Lazy(fs, func(w http.ResponseWriter, r *http.Request) {
		fs.PauseDownload(w, r)
	}))).
		Methods("OPTIONS", "PUT")

	router.HandleFunc("/file-svc/download/{url}", middlewares.DefaultApplicator(service.Lazy(fs, func(w http.ResponseWriter, r *http.Request) {
		fs.GetDownload(w, r)
	}))).
		Methods("OPTIONS", "GET")

	router.HandleFunc("/file-svc/downloads", middlewares.DefaultApplicator(service.Lazy(fs, func(w http.ResponseWriter, r *http.Request) {
		fs.ListDownloads(w, r)
	}))).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/file-svc/upload", middlewares.DefaultApplicator(service.Lazy(fs, func(w http.ResponseWriter, r *http.Request) {
		fs.UploadFile(w, r)
	}))).
		Methods("OPTIONS", "PUT")

	router.HandleFunc("/file-svc/uploads", middlewares.DefaultApplicator(service.Lazy(fs, func(w http.ResponseWriter, r *http.Request) {
		fs.ListUploads(w, r)
	}))).
		Methods("OPTIONS", "POST")

	// @todo
	// Investigate why SkipLock is needed here.
	// I placed it here because the serve proxy tests were deadlocking.
	// Not sure why though as they are not routing to the same node (themselves),
	// but to an other node.

	router.HandleFunc("/file-svc/serve/upload/{fileId}", middlewares.DefaultApplicator(service.Lazy(fs, func(w http.ResponseWriter, r *http.Request) {
		fs.ServeUpload(w, r)
	}, service.WithSkipLock()))).
		Methods("OPTIONS", "GET")

	router.HandleFunc("/file-svc/serve/download/{url}", middlewares.DefaultApplicator(service.Lazy(fs, func(w http.ResponseWriter, r *http.Request) {
		fs.ServeDownload(w, r)
	}, service.WithSkipLock()))).
		Methods("OPTIONS", "GET")
}

func (fs *FileService) Start() error {
	credentialStore, err := fs.datastoreFactory(
		"fileSvcCredentials",
		&auth.Credential{},
	)
	if err != nil {
		return err
	}
	fs.credentialStore = credentialStore

	uploadFolder := path.Join(fs.homeDir, "uploads")
	err = os.MkdirAll(uploadFolder, 0700)
	if err != nil {
		return err
	}
	fs.uploadFolder = uploadFolder

	downloadFolder := path.Join(fs.homeDir, "downloads")
	err = os.MkdirAll(downloadFolder, 0700)
	if err != nil {
		return err
	}
	fs.downloadFolder = downloadFolder

	downloadStore, err := fs.datastoreFactory(
		"fileSvcDownloads",
		&types.InternalDownload{},
	)
	if err != nil {
		return err
	}
	fs.downloadStore = downloadStore

	uploadStore, err := fs.datastoreFactory(
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
	fs.lock.Acquire(ctx, "file-svc-start")
	defer fs.lock.Release(ctx, "file-svc-start")

	token, err := boot.RegisterServiceAccount(
		fs.clientFactory.Client().UserSvcAPI,
		"file-svc",
		"File Svc",
		fs.credentialStore,
	)
	if err != nil {
		return "", err
	}
	fs.token = token.Token

	err = fs.registerPermissions()
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
