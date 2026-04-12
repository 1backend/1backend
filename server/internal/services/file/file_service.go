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
	"fmt"
	"hash/fnv"
	"log/slog"
	"net/http"
	"os"
	"path"
	"sync"
	"time"

	"cloud.google.com/go/storage"
	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/boot"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/logger"
	"github.com/1backend/1backend/sdk/go/service"
	file "github.com/1backend/1backend/server/internal/services/file/types"
	types "github.com/1backend/1backend/server/internal/services/file/types"
	"github.com/1backend/1backend/server/internal/universe"
	"github.com/gorilla/mux"
	lru "github.com/hashicorp/golang-lru/v2"
	"github.com/pkg/errors"
	"google.golang.org/api/option"
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

	storage StorageProvider
	// downloadStorage is an optional backing storage used for downloaded
	// internet assets (e.g. GCS). Local disk remains the serving cache.
	downloadStorage StorageProvider

	nodeId string
	cache  *lru.Cache[string, *file.Upload]

	accessFlushInterval time.Duration
	lastAccessMu        sync.Mutex
	pendingLastAccessAt map[string]pendingAccess
}

type pendingAccess struct {
	FileId         string
	LastAccessedAt time.Time
}

const restartDownloadBaseBackoff = 3 * time.Second

func NewFileService(
	options *universe.Options,
) (*FileService, error) {
	fs := &FileService{
		options: options,
	}

	fs.uploadFolder = path.Join(fs.options.HomeDir, "uploads")
	fs.downloadFolder = path.Join(fs.options.HomeDir, "downloads")

	localProvider := &LocalProvider{
		uploadFolder: fs.uploadFolder,
	}
	localDownloadProvider := &LocalProvider{
		uploadFolder: fs.downloadFolder,
	}

	fs.cache, _ = lru.New[string, *file.Upload](100000)
	fs.pendingLastAccessAt = map[string]pendingAccess{}
	fs.accessFlushInterval = 30 * time.Second
	if options.Test {
		fs.accessFlushInterval = 250 * time.Millisecond
	}

	// Determine Strategy
	if options.FileGcs {
		ctx := context.Background()

		if options.GcpSaKey == "" || options.GcsBucket == "" {
			return nil, fmt.Errorf("GCS enabled but OB_GCP_SA_KEY or OB_GCS_BUCKET is missing")
		}

		// Initialize GCS Client
		gcsClient, err := storage.NewClient(ctx, option.WithCredentialsFile(options.GcpSaKey))
		if err != nil {
			return nil, fmt.Errorf("failed to init GCS client: %w", err)
		}

		gcsProvider := &GCSProvider{
			client: gcsClient,
			bucket: options.GcsBucket,
		}

		// Wrap uploads in Cloud+Local cache provider.
		fs.storage = &CloudCacheProvider{
			cloud: gcsProvider,
			local: localProvider,
		}
		// Downloads are backed by cloud and cached under ~/.1backend/downloads.
		fs.downloadStorage = &CloudCacheProvider{
			cloud: &PrefixedStorageProvider{
				base:   gcsProvider,
				prefix: "downloads",
			},
			local: localDownloadProvider,
		}
		logger.Info("File service initialized with GCS Cloud Cache")
	} else {
		// Fallback to standard distributed behavior
		fs.storage = localProvider
		fs.downloadStorage = nil
		logger.Info("File service initialized with Local Storage")
	}

	return fs, nil
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

	router.HandleFunc("/file-svc/upload/{fileId}", appl(service.Lazy(fs, func(w http.ResponseWriter, r *http.Request) {
		fs.DeleteUpload(w, r)
	}))).
		Methods("OPTIONS", "DELETE")

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

	err = os.MkdirAll(fs.uploadFolder, 0700)
	if err != nil {
		return err
	}

	err = os.MkdirAll(fs.downloadFolder, 0700)
	if err != nil {
		return err
	}

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

	if fs.nodeId == "" {
		fs.nodeId = fs.options.NodeId
	}

	downloads, err := fs.downloadStore.Query(
		datastore.Equals([]string{"status"}, types.DownloadStatusInProgress),
	).Find()
	if err != nil {
		return nil
	}

	for _, downloadI := range downloads {
		download := downloadI.(*types.InternalDownload)

		if download.Status == types.DownloadStatusInProgress {
			go fs.restartDownloadWithBackoffLoop(context.Background(), download)
		}
	}

	go fs.flushAccessLoop()

	return err
}

func (fs *FileService) restartDownloadWithBackoffLoop(
	ctx context.Context,
	download *types.InternalDownload,
) {
	for {
		shouldRetry := fs.restartDownloadWithBackoffOnce(ctx, download)
		if !shouldRetry {
			return
		}
	}
}

func (fs *FileService) restartDownloadWithBackoffOnce(
	ctx context.Context,
	download *types.InternalDownload,
) bool {
	latest, exists := fs.getDownloadById(download.Id)
	if exists {
		download = latest
	}
	if download.Status != types.DownloadStatusInProgress {
		return false
	}

	now := time.Now().UTC()
	if download.NextRetryAt != nil && download.NextRetryAt.After(now) {
		time.Sleep(time.Until(*download.NextRetryAt))
	}

	lockKey := fmt.Sprintf("file-svc-download-restart:%s", download.Id)
	if lockKey == "file-svc-download-restart:" {
		lockKey = fmt.Sprintf("file-svc-download-restart:%s", EncodeURLtoFileName(download.URL))
	}
	if fs.options != nil && fs.options.Lock != nil {
		acquired, lockErr := fs.options.Lock.TryAcquire(ctx, lockKey)
		if lockErr != nil || !acquired {
			return false
		}
		defer func() {
			if err := fs.options.Lock.Release(ctx, lockKey); err != nil {
				logger.Warn("Failed to release restart lock",
					slog.String("lockKey", lockKey),
					slog.String("error", err.Error()),
				)
			}
		}()
	}

	latest, exists = fs.getDownloadById(download.Id)
	if exists {
		download = latest
	}
	if download.Status != types.DownloadStatusInProgress {
		return false
	}

	err := fs.download(ctx, download.URL, path.Dir(download.FilePath), fs.SyncDownloads)
	if err == nil {
		return false
	}

	retryCount := 1
	if download.RetryCount != nil && *download.RetryCount > 0 {
		retryCount = *download.RetryCount + 1
	}
	download.RetryCount = &retryCount
	errMessage := err.Error()
	download.LastError = &errMessage

	waitFor := restartDownloadBackoff(restartDownloadBaseBackoff, retryCount, download.URL)
	if waitFor > 0 {
		nextRetry := time.Now().UTC().Add(waitFor)
		download.NextRetryAt = &nextRetry
	}

	if upsertErr := fs.downloadStore.Upsert(download); upsertErr != nil {
		logger.Error("Failed to persist restart download retry state",
			slog.String("url", download.URL),
			slog.String("error", upsertErr.Error()),
		)
		return false
	}

	logger.Warn("Restart download attempt failed; persisted retry metadata",
		slog.String("url", download.URL),
		slog.Int("retryCount", retryCount),
		slog.Duration("nextBackoff", waitFor),
		slog.String("error", err.Error()),
	)
	return true
}

func restartDownloadBackoff(base time.Duration, retryCount int, url string) time.Duration {
	if base <= 0 || retryCount <= 0 {
		return 0
	}

	backoff := base * time.Duration(1<<uint(retryCount-1))

	h := fnv.New32a()
	_, _ = h.Write([]byte(url))
	jitter := time.Duration(h.Sum32()%1000) * (base / 1000)
	if jitter <= 0 {
		jitter = time.Duration(h.Sum32()%10) * time.Millisecond
	}

	return backoff + jitter
}

func (fs *FileService) markUploadAccess(upload *types.Upload) {
	if upload == nil || upload.Id == "" {
		return
	}

	now := time.Now().UTC()

	fs.lastAccessMu.Lock()
	prev, exists := fs.pendingLastAccessAt[upload.Id]
	if !exists || now.After(prev.LastAccessedAt) {
		fs.pendingLastAccessAt[upload.Id] = pendingAccess{
			FileId:         upload.FileId,
			LastAccessedAt: now,
		}
	}
	fs.lastAccessMu.Unlock()
}

func (fs *FileService) flushAccessLoop() {
	ticker := time.NewTicker(fs.accessFlushInterval)
	defer ticker.Stop()

	for range ticker.C {
		fs.flushAccesses()
	}
}

func (fs *FileService) flushAccesses() {
	pending := fs.swapPendingAccesses()
	if len(pending) == 0 {
		return
	}

	patches := make([]datastore.Patch, 0, len(pending))
	cacheUpdates := make(map[string]time.Time, len(pending))
	for uploadId, update := range pending {
		patches = append(patches, datastore.Patch{
			ID: uploadId,
			Fields: map[string]any{
				"lastAccessedAt": update.LastAccessedAt,
			},
		})

		prev, exists := cacheUpdates[update.FileId]
		if !exists || update.LastAccessedAt.After(prev) {
			cacheUpdates[update.FileId] = update.LastAccessedAt
		}
	}

	if len(patches) > 0 {
		if err := fs.uploadStore.PatchMany(patches); err != nil {
			// non-critical: never fail file serving due to analytics-like metadata updates
			logger.Error("Failed to flush lastAccessedAt batch",
				slog.Int("count", len(patches)),
				slog.Any("error", err))
			return
		}
	}

	for fileId, lastAccessedAt := range cacheUpdates {
		if cachedUpload, ok := fs.cache.Get(fileId); ok {
			accessedAt := lastAccessedAt
			cachedUpload.LastAccessedAt = &accessedAt
			fs.cache.Add(fileId, cachedUpload)
		}
	}
}

func (fs *FileService) swapPendingAccesses() map[string]pendingAccess {
	fs.lastAccessMu.Lock()
	defer fs.lastAccessMu.Unlock()

	if len(fs.pendingLastAccessAt) == 0 {
		return nil
	}

	pending := fs.pendingLastAccessAt
	fs.pendingLastAccessAt = map[string]pendingAccess{}
	return pending
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

func (fs *FileService) getDownloadById(id string) (*types.InternalDownload, bool) {
	if id == "" {
		return nil, false
	}

	fs.mutex.Lock()
	defer fs.mutex.Unlock()

	downloadIs, err := fs.downloadStore.Query(datastore.Id(id)).Find()
	if err != nil || len(downloadIs) == 0 {
		return nil, false
	}

	return downloadIs[0].(*types.InternalDownload), true
}
