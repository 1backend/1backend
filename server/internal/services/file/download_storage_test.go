package fileservice

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"sync/atomic"
	"testing"
	"time"

	"github.com/1backend/1backend/sdk/go/datastore/localstore"
	distlock "github.com/1backend/1backend/sdk/go/lock/local"
	filetypes "github.com/1backend/1backend/server/internal/services/file/types"
	"github.com/1backend/1backend/server/internal/universe"
	"github.com/stretchr/testify/require"
)

type memoryStorageProvider struct {
	data map[string][]byte
}

func newMemoryStorageProvider() *memoryStorageProvider {
	return &memoryStorageProvider{
		data: map[string][]byte{},
	}
}

func ptrInt(v int) *int {
	return &v
}

func (m *memoryStorageProvider) Open(_ context.Context, filePath string) (io.ReadCloser, int64, error) {
	bs, ok := m.data[filePath]
	if !ok {
		return nil, 0, os.ErrNotExist
	}
	return io.NopCloser(newBytesReader(bs)), int64(len(bs)), nil
}

func (m *memoryStorageProvider) Save(_ context.Context, u *filetypes.Upload, content io.Reader) (int64, error) {
	bs, err := io.ReadAll(content)
	if err != nil {
		return 0, err
	}
	m.data[u.FilePath] = bs
	return int64(len(bs)), nil
}

func (m *memoryStorageProvider) NewWriter(ctx context.Context, filePath string) (io.WriteCloser, error) {
	pr, pw := io.Pipe()
	go func() {
		defer pr.Close()
		_, _ = m.Save(ctx, &filetypes.Upload{FilePath: filePath}, pr)
	}()
	return pw, nil
}

func (m *memoryStorageProvider) Delete(_ context.Context, filePath string) error {
	delete(m.data, filePath)
	return nil
}

type bytesReader struct {
	bs []byte
	i  int
}

func newBytesReader(bs []byte) *bytesReader {
	return &bytesReader{bs: bs}
}

func (b *bytesReader) Read(p []byte) (int, error) {
	if b.i >= len(b.bs) {
		return 0, io.EOF
	}
	n := copy(p, b.bs[b.i:])
	b.i += n
	return n, nil
}

func TestPersistAndRestoreDownloadStorage(t *testing.T) {
	tmp := t.TempDir()
	localFile := filepath.Join(tmp, "downloads", "asset.bin")
	require.NoError(t, os.MkdirAll(filepath.Dir(localFile), 0755))
	require.NoError(t, os.WriteFile(localFile, []byte("persist-me"), 0644))

	storage := newMemoryStorageProvider()
	fs := &FileService{
		downloadStorage: storage,
	}

	require.NoError(t, fs.persistDownloadToStorage(context.Background(), "asset-key", localFile))

	// Clear local and restore from backend.
	require.NoError(t, os.Remove(localFile))
	restored, size, err := fs.restoreDownloadFromStorage(context.Background(), "asset-key", localFile)
	require.NoError(t, err)
	require.True(t, restored)
	require.Equal(t, int64(len("persist-me")), size)

	bs, err := os.ReadFile(localFile)
	require.NoError(t, err)
	require.Equal(t, "persist-me", string(bs))
}

func TestDownloadFilePersistsToDownloadStorage(t *testing.T) {
	origin := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Empty body avoids firehose/token path while still exercising completion + persist.
		w.WriteHeader(http.StatusOK)
	}))
	defer origin.Close()

	tmp := t.TempDir()
	dsPath := filepath.Join(tmp, "downloads.json")
	downloadStore, err := localstore.NewLocalStore(&filetypes.InternalDownload{}, dsPath)
	require.NoError(t, err)
	defer downloadStore.Close()

	filePath := filepath.Join(tmp, "downloads", EncodeURLtoFileName(origin.URL+"/asset.txt"))
	require.NoError(t, os.MkdirAll(filepath.Dir(filePath), 0755))
	storage := newMemoryStorageProvider()
	fs := &FileService{
		downloadStorage: storage,
		downloadStore:   downloadStore,
	}

	d := &filetypes.InternalDownload{
		Id:       "dl_test",
		URL:      origin.URL + "/asset.txt",
		NodeId:   "node-1",
		FilePath: filePath,
		Status:   filetypes.DownloadStatusInProgress,
	}

	require.NoError(t, fs.downloadFile(d))
	require.Equal(t, filetypes.DownloadStatusCompleted, d.Status)
	require.FileExists(t, filePath)

	storageKey := DownloadStorageFilePath(d.URL)
	_, ok := storage.data[storageKey]
	require.True(t, ok, "expected completed download to be persisted to download storage")
}

func TestServeLocalDownloadRecoversFromStorageWhenLocalMissing(t *testing.T) {
	origin := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Empty body avoids firehose/token path while still exercising fallback behavior.
		w.WriteHeader(http.StatusOK)
	}))
	defer origin.Close()

	tmp := t.TempDir()
	dsPath := filepath.Join(tmp, "downloads.json")
	downloadStore, err := localstore.NewLocalStore(&filetypes.InternalDownload{}, dsPath)
	require.NoError(t, err)
	defer downloadStore.Close()

	url := origin.URL + "/assets/cached.txt"
	localPath := filepath.Join(tmp, "downloads", EncodeURLtoFileName(url))
	require.NoError(t, os.MkdirAll(filepath.Dir(localPath), 0755))
	storage := newMemoryStorageProvider()
	storage.data[DownloadStorageFilePath(url)] = []byte("from-storage")

	download := &filetypes.InternalDownload{
		Id:       "dl_cached",
		URL:      url,
		NodeId:   "node-1",
		FilePath: localPath,
		Status:   filetypes.DownloadStatusInProgress,
	}
	require.NoError(t, downloadStore.Upsert(download))

	fs := &FileService{
		nodeId:          "node-1",
		downloadStore:   downloadStore,
		downloadStorage: storage,
		downloadFolder:  filepath.Join(tmp, "downloads"),
	}

	req := httptest.NewRequest(http.MethodGet, "/file-svc/serve/download/ignored", nil)
	w := httptest.NewRecorder()
	fs.serveLocalDownload([]*filetypes.InternalDownload{download}, w, req)

	resp := w.Result()
	defer resp.Body.Close()
	require.Equal(t, http.StatusOK, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.Equal(t, "from-storage", string(body))
	require.FileExists(t, localPath)

	updatedI, exists, err := downloadStore.Query().FindOne()
	require.NoError(t, err)
	require.True(t, exists)
	updated := updatedI.(*filetypes.InternalDownload)
	require.Equal(t, filetypes.DownloadStatusCompleted, updated.Status)
}

func TestServeLocalDownloadRedownloadsWhenLocalAndStorageMissing(t *testing.T) {
	originHits := 0
	origin := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		originHits++
		w.WriteHeader(http.StatusOK)
	}))
	defer origin.Close()

	tmp := t.TempDir()
	dsPath := filepath.Join(tmp, "downloads.json")
	downloadStore, err := localstore.NewLocalStore(&filetypes.InternalDownload{}, dsPath)
	require.NoError(t, err)
	defer downloadStore.Close()

	url := origin.URL + "/assets/missing.txt"
	localPath := filepath.Join(tmp, "downloads", EncodeURLtoFileName(url))
	require.NoError(t, os.MkdirAll(filepath.Dir(localPath), 0755))
	storage := newMemoryStorageProvider()

	download := &filetypes.InternalDownload{
		Id:       "dl_missing",
		URL:      url,
		NodeId:   "node-1",
		FilePath: localPath,
		Status:   filetypes.DownloadStatusInProgress,
	}
	require.NoError(t, downloadStore.Upsert(download))

	fs := &FileService{
		nodeId:          "node-1",
		downloadStore:   downloadStore,
		downloadStorage: storage,
		downloadFolder:  filepath.Join(tmp, "downloads"),
		SyncDownloads:   true,
	}

	req := httptest.NewRequest(http.MethodGet, "/file-svc/serve/download/ignored", nil)
	w := httptest.NewRecorder()
	fs.serveLocalDownload([]*filetypes.InternalDownload{download}, w, req)

	resp := w.Result()
	defer resp.Body.Close()
	require.Equal(t, http.StatusOK, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.Empty(t, string(body))
	require.Equal(t, 1, originHits, "expected one redownload from origin")
	require.FileExists(t, localPath)

	storageKey := DownloadStorageFilePath(url)
	require.Equal(t, []byte(""), storage.data[storageKey], "expected redownloaded object persisted to storage")
}

func TestServeRemoteDownloadPromotesStoredDownloadToCurrentNode(t *testing.T) {
	tmp := t.TempDir()
	dsPath := filepath.Join(tmp, "downloads.json")
	downloadStore, err := localstore.NewLocalStore(&filetypes.InternalDownload{}, dsPath)
	require.NoError(t, err)
	defer downloadStore.Close()

	url := "https://example.com/assets/stale.txt"
	stalePath := filepath.Join("/old-node/downloads", EncodeURLtoFileName(url))
	storage := newMemoryStorageProvider()
	storage.data[DownloadStorageFilePath(url)] = []byte("from-storage")

	download := &filetypes.InternalDownload{
		Id:             "dl_stale",
		URL:            url,
		NodeId:         "stale-node",
		FilePath:       stalePath,
		Status:         filetypes.DownloadStatusCompleted,
		TotalSize:      0,
		DownloadedSize: int64(len("from-storage")),
	}
	require.NoError(t, downloadStore.Upsert(download))

	fs := &FileService{
		nodeId:          "current-node",
		downloadStore:   downloadStore,
		downloadStorage: storage,
		downloadFolder:  filepath.Join(tmp, "downloads"),
	}

	req := httptest.NewRequest(http.MethodGet, "/file-svc/serve/download/ignored", nil)
	w := httptest.NewRecorder()
	fs.serveRemoteDownload([]*filetypes.InternalDownload{download}, w, req)

	resp := w.Result()
	defer resp.Body.Close()
	require.Equal(t, http.StatusOK, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.Equal(t, "from-storage", string(body))

	localPath := filepath.Join(fs.downloadFolder, EncodeURLtoFileName(url))
	require.FileExists(t, localPath)

	updated, exists := fs.getDownload(url)
	require.True(t, exists)
	require.Equal(t, "current-node", updated.NodeId)
	require.Equal(t, localPath, updated.FilePath)
	require.Equal(t, filetypes.DownloadStatusCompleted, updated.Status)
	require.Equal(t, int64(len("from-storage")), updated.TotalSize)
	require.Equal(t, int64(len("from-storage")), updated.DownloadedSize)
}

func TestRestoreDownloadFromStorageReturnsFalseWhenObjectMissing(t *testing.T) {
	tmp := t.TempDir()
	localFile := filepath.Join(tmp, "downloads", "missing.bin")

	url := "https://example.com/missing.bin"
	storage := newMemoryStorageProvider()

	fs := &FileService{
		downloadStorage: storage,
	}

	restored, size, err := fs.restoreDownloadFromStorage(
		context.Background(),
		DownloadStorageFilePath(url),
		localFile,
	)
	require.NoError(t, err)
	require.False(t, restored)
	require.Equal(t, int64(0), size)
	require.NoFileExists(t, localFile)
}

func TestPrefixedStorageProviderAppliesPrefix(t *testing.T) {
	base := newMemoryStorageProvider()
	p := &PrefixedStorageProvider{
		base:   base,
		prefix: "downloads",
	}

	_, err := p.Save(context.Background(), &filetypes.Upload{
		FilePath: "key.txt",
	}, newBytesReader([]byte("abc")))
	require.NoError(t, err)

	_, ok := base.data["downloads/key.txt"]
	require.True(t, ok)

	rc, _, err := p.Open(context.Background(), "key.txt")
	require.NoError(t, err)
	defer rc.Close()

	bs, err := io.ReadAll(rc)
	require.NoError(t, err)
	require.Equal(t, "abc", string(bs))
}

func TestDownloadStorageUsesPrefixAndShardedPath(t *testing.T) {
	base := newMemoryStorageProvider()
	p := &PrefixedStorageProvider{
		base:   base,
		prefix: "downloads",
	}

	url := "https://example.com/a/b/c.png"
	shardedPath := DownloadStorageFilePath(url)

	_, err := p.Save(context.Background(), &filetypes.Upload{
		FilePath: shardedPath,
	}, newBytesReader([]byte("abc")))
	require.NoError(t, err)

	_, ok := base.data["downloads/"+filepath.ToSlash(shardedPath)]
	require.True(t, ok, "expected objects saved as downloads/<2>/<2>/<full-hash>")
}

func TestShardStoragePathAlwaysUsesTwoCharBuckets(t *testing.T) {
	require.Equal(t, filepath.Join("ab", "cd", "abcd"), shardStoragePath("abcd"))
	require.Equal(t, filepath.Join("ab", "__", "ab"), shardStoragePath("ab"))
	require.Equal(t, filepath.Join("a_", "__", "a"), shardStoragePath("a"))
}

func TestShardStoragePathWithBasisExhaustive(t *testing.T) {
	testCases := []struct {
		name     string
		basis    string
		fileName string
		expected string
	}{
		{
			name:     "empty basis and file",
			basis:    "",
			fileName: "",
			expected: filepath.Join("__", "__", ""),
		},
		{
			name:     "one char basis",
			basis:    "a",
			fileName: "a",
			expected: filepath.Join("a_", "__", "a"),
		},
		{
			name:     "two char basis",
			basis:    "ab",
			fileName: "ab",
			expected: filepath.Join("ab", "__", "ab"),
		},
		{
			name:     "three char basis",
			basis:    "abc",
			fileName: "abc",
			expected: filepath.Join("ab", "c_", "abc"),
		},
		{
			name:     "four char basis",
			basis:    "abcd",
			fileName: "abcd",
			expected: filepath.Join("ab", "cd", "abcd"),
		},
		{
			name:     "five char basis",
			basis:    "abcde",
			fileName: "abcde",
			expected: filepath.Join("ab", "cd", "abcde"),
		},
		{
			name:     "long basis",
			basis:    "this-is-a-long-basis",
			fileName: "this-is-a-long-basis",
			expected: filepath.Join("th", "is", "this-is-a-long-basis"),
		},
		{
			name:     "upload style basis and file name differ",
			basis:    "81d259fc",
			fileName: "file_81d259fc",
			expected: filepath.Join("81", "d2", "file_81d259fc"),
		},
		{
			name:     "short upload style basis and full file id",
			basis:    "x",
			fileName: "file_x",
			expected: filepath.Join("x_", "__", "file_x"),
		},
		{
			name:     "basis with symbols",
			basis:    "-_",
			fileName: "file-special",
			expected: filepath.Join("-_", "__", "file-special"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, shardStoragePathWithBasis(tc.basis, tc.fileName))
		})
	}
}

func TestRestartDownloadWithBackoffPersistsRetryMetadata(t *testing.T) {
	origin := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
	}))
	defer origin.Close()

	tmp := t.TempDir()
	dsPath := filepath.Join(tmp, "downloads.json")
	downloadStore, err := localstore.NewLocalStore(&filetypes.InternalDownload{}, dsPath)
	require.NoError(t, err)
	defer downloadStore.Close()

	fs := &FileService{
		nodeId:         "node-1",
		downloadFolder: filepath.Join(tmp, "downloads"),
		downloadStore:  downloadStore,
		SyncDownloads:  true,
	}
	require.NoError(t, os.MkdirAll(fs.downloadFolder, 0755))

	d := &filetypes.InternalDownload{
		Id:       "dl_retry",
		URL:      origin.URL + "/asset.txt",
		NodeId:   "node-1",
		FilePath: filepath.Join(fs.downloadFolder, EncodeURLtoFileName(origin.URL+"/asset.txt")),
		Status:   filetypes.DownloadStatusInProgress,
	}
	require.NoError(t, downloadStore.Upsert(d))

	fs.restartDownloadWithBackoffOnce(context.Background(), d)

	dl, exists := fs.getDownload(d.URL)
	require.True(t, exists)
	require.NotNil(t, dl)
	require.NotNil(t, dl.RetryCount)
	require.Equal(t, 1, *dl.RetryCount)
	require.NotNil(t, dl.NextRetryAt)
	require.NotNil(t, dl.LastError)
	require.NotEmpty(t, *dl.LastError)
	require.Equal(t, filetypes.DownloadStatusInProgress, dl.Status)
}

func TestRestartDownloadWithBackoffHonorsNextRetryAt(t *testing.T) {
	var requestCount atomic.Int32
	origin := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount.Add(1)
		w.WriteHeader(http.StatusOK)
	}))
	defer origin.Close()

	tmp := t.TempDir()
	dsPath := filepath.Join(tmp, "downloads.json")
	downloadStore, err := localstore.NewLocalStore(&filetypes.InternalDownload{}, dsPath)
	require.NoError(t, err)
	defer downloadStore.Close()

	nextRetry := time.Now().UTC().Add(250 * time.Millisecond)
	fs := &FileService{
		nodeId:         "node-1",
		downloadFolder: filepath.Join(tmp, "downloads"),
		downloadStore:  downloadStore,
		SyncDownloads:  true,
	}
	require.NoError(t, os.MkdirAll(fs.downloadFolder, 0755))

	d := &filetypes.InternalDownload{
		Id:          "dl_delayed",
		URL:         origin.URL + "/delayed.txt",
		NodeId:      "node-1",
		FilePath:    filepath.Join(fs.downloadFolder, EncodeURLtoFileName(origin.URL+"/delayed.txt")),
		Status:      filetypes.DownloadStatusInProgress,
		RetryCount:  ptrInt(2),
		NextRetryAt: &nextRetry,
	}
	require.NoError(t, downloadStore.Upsert(d))

	done := make(chan struct{})
	go func() {
		defer close(done)
		fs.restartDownloadWithBackoffOnce(context.Background(), d)
	}()

	time.Sleep(120 * time.Millisecond)
	require.Equal(t, int32(0), requestCount.Load())

	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatal("timeout waiting for delayed restart")
	}
	require.GreaterOrEqual(t, requestCount.Load(), int32(1))
}

func TestRestartDownloadWithBackoffLoopRetriesUntilSuccess(t *testing.T) {
	var requestCount atomic.Int32
	origin := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		n := requestCount.Add(1)
		if n <= 2 {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer origin.Close()

	tmp := t.TempDir()
	dsPath := filepath.Join(tmp, "downloads.json")
	downloadStore, err := localstore.NewLocalStore(&filetypes.InternalDownload{}, dsPath)
	require.NoError(t, err)
	defer downloadStore.Close()

	fs := &FileService{
		nodeId:         "node-1",
		downloadFolder: filepath.Join(tmp, "downloads"),
		downloadStore:  downloadStore,
		SyncDownloads:  true,
	}
	require.NoError(t, os.MkdirAll(fs.downloadFolder, 0755))

	d := &filetypes.InternalDownload{
		Id:       "dl_loop",
		URL:      origin.URL + "/loop.txt",
		NodeId:   "node-1",
		FilePath: filepath.Join(fs.downloadFolder, EncodeURLtoFileName(origin.URL+"/loop.txt")),
		Status:   filetypes.DownloadStatusInProgress,
	}
	require.NoError(t, downloadStore.Upsert(d))

	done := make(chan struct{})
	go func() {
		defer close(done)
		fs.restartDownloadWithBackoffLoop(context.Background(), d)
	}()

	select {
	case <-done:
	case <-time.After(15 * time.Second):
		t.Fatal("timeout waiting for retry loop success")
	}

	updated, exists := fs.getDownload(d.URL)
	require.True(t, exists)
	require.Equal(t, filetypes.DownloadStatusCompleted, updated.Status)
	require.GreaterOrEqual(t, requestCount.Load(), int32(3))
}

func TestRestartDownloadWithBackoffUsesDistributedLock(t *testing.T) {
	var requestCount atomic.Int32
	origin := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount.Add(1)
		w.WriteHeader(http.StatusOK)
	}))
	defer origin.Close()

	tmp := t.TempDir()
	dsPath := filepath.Join(tmp, "downloads.json")
	downloadStore, err := localstore.NewLocalStore(&filetypes.InternalDownload{}, dsPath)
	require.NoError(t, err)
	defer downloadStore.Close()

	lock := distlock.NewLocalDistributedLock()
	fs1 := &FileService{
		nodeId:         "node-1",
		downloadFolder: filepath.Join(tmp, "downloads"),
		downloadStore:  downloadStore,
		SyncDownloads:  true,
		options:        &universe.Options{Lock: lock},
	}
	fs2 := &FileService{
		nodeId:         "node-2",
		downloadFolder: filepath.Join(tmp, "downloads"),
		downloadStore:  downloadStore,
		SyncDownloads:  true,
		options:        &universe.Options{Lock: lock},
	}
	require.NoError(t, os.MkdirAll(fs1.downloadFolder, 0755))

	d := &filetypes.InternalDownload{
		Id:       "dl_lock",
		URL:      origin.URL + "/lock.txt",
		NodeId:   "node-1",
		FilePath: filepath.Join(fs1.downloadFolder, EncodeURLtoFileName(origin.URL+"/lock.txt")),
		Status:   filetypes.DownloadStatusInProgress,
	}
	require.NoError(t, downloadStore.Upsert(d))

	done := make(chan struct{})
	go func() {
		defer close(done)
		_ = fs1.restartDownloadWithBackoffOnce(context.Background(), d)
	}()
	_ = fs2.restartDownloadWithBackoffOnce(context.Background(), d)
	<-done

	require.Equal(t, int32(1), requestCount.Load())
}
