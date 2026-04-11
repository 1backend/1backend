package fileservice

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/1backend/1backend/sdk/go/datastore/localstore"
	filetypes "github.com/1backend/1backend/server/internal/services/file/types"
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

	storageKey := EncodeURLtoFileName(d.URL)
	_, ok := storage.data[storageKey]
	require.True(t, ok, "expected completed download to be persisted to download storage")
}

func TestServeLocalDownloadRestoresFromStorageBeforeOriginRecovery(t *testing.T) {
	tmp := t.TempDir()
	dsPath := filepath.Join(tmp, "downloads.json")
	downloadStore, err := localstore.NewLocalStore(&filetypes.InternalDownload{}, dsPath)
	require.NoError(t, err)
	defer downloadStore.Close()

	url := "https://example.com/assets/cached.txt"
	localPath := filepath.Join(tmp, "downloads", EncodeURLtoFileName(url))
	storage := newMemoryStorageProvider()
	storage.data[EncodeURLtoFileName(url)] = []byte("from-storage")

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
