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
	"os"
	"path"
	"sync"

	sdk "github.com/openorch/openorch/sdk/go"
	"github.com/openorch/openorch/sdk/go/datastore"
	"github.com/openorch/openorch/sdk/go/lock"
	types "github.com/openorch/openorch/server/internal/services/file/types"
)

type FileService struct {
	clientFactory sdk.ClientFactory
	token         string

	dlock lock.DistributedLock

	uploadFolder   string
	downloadFolder string

	lock sync.Mutex

	downloadStore datastore.DataStore
	uploadStore   datastore.DataStore

	// for testing purposes
	SyncDownloads bool

	credentialStore datastore.DataStore

	nodeId string
}

func NewFileService(
	clientFactory sdk.ClientFactory,
	lock lock.DistributedLock,
	datastoreFactory func(tableName string, instance any) (datastore.DataStore, error),
	homeDir string,
) (*FileService, error) {
	credentialStore, err := datastoreFactory(
		"fileSvcCredentials",
		&sdk.Credential{},
	)
	if err != nil {
		return nil, err
	}

	uploadFolder := path.Join(homeDir, "uploads")
	err = os.MkdirAll(uploadFolder, 0700)
	if err != nil {
		return nil, err
	}

	downloadFolder := path.Join(homeDir, "downloads")
	err = os.MkdirAll(downloadFolder, 0700)
	if err != nil {
		return nil, err
	}

	downloadStore, err := datastoreFactory(
		"fileSvcDownloads",
		&types.InternalDownload{},
	)
	if err != nil {
		return nil, err
	}

	uploadStore, err := datastoreFactory(
		"fileSvcUploads",
		&types.Upload{},
	)
	if err != nil {
		return nil, err
	}

	ret := &FileService{
		clientFactory: clientFactory,

		credentialStore: credentialStore,
		dlock:           lock,

		uploadFolder:   uploadFolder,
		downloadFolder: downloadFolder,

		downloadStore: downloadStore,
		uploadStore:   uploadStore,
	}

	return ret, nil
}

func (dm *FileService) Start() error {
	ctx := context.Background()
	dm.dlock.Acquire(ctx, "file-svc-start")
	defer dm.dlock.Release(ctx, "file-svc-start")

	token, err := sdk.RegisterServiceAccount(
		dm.clientFactory.Client().UserSvcAPI,
		"file-svc",
		"File Svc",
		dm.credentialStore,
	)
	if err != nil {
		return err
	}
	dm.token = token

	err = dm.registerPermissions()
	if err != nil {
		return err
	}

	downloads, err := dm.downloadStore.Query(
		datastore.Equals([]string{"status"},
			types.DownloadStatusInProgress,
		)).Find()
	if err != nil {
		return nil
	}

	for _, downloadI := range downloads {
		download := downloadI.(*types.InternalDownload)

		if download.Status == types.DownloadStatusInProgress {
			err = dm.download(context.Background(), download.URL, path.Dir(download.FilePath))
			if err != nil {
				return err
			}
		}
	}

	return err
}

func (dm *FileService) getDownload(url string) (*types.InternalDownload, bool) {
	dm.lock.Lock()
	defer dm.lock.Unlock()

	downloadIs, err := dm.downloadStore.Query(
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
