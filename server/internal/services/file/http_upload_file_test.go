package fileservice_test

import (
	"context"
	"io/ioutil"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	sdk "github.com/openorch/openorch/sdk/go"
	"github.com/openorch/openorch/sdk/go/test"
	"github.com/openorch/openorch/server/internal/di"
	"github.com/stretchr/testify/require"
)

func createTestFile(t *testing.T, content string) (*os.File, func()) {
	file, err := os.CreateTemp("", "upload-test-*.txt")
	require.NoError(t, err)
	_, err = file.WriteString(content)
	require.NoError(t, err)
	file.Seek(0, 0)
	return file, func() { os.Remove(file.Name()) }
}

func TestUploadFile(t *testing.T) {
	file, cleanup := createTestFile(t, "Test file content")
	defer cleanup()

	hs := &di.HandlerSwitcher{}
	server := httptest.NewServer(hs)
	defer server.Close()

	options := &di.Options{
		Test: true,
		Url:  server.URL,
	}
	universe, starterFunc, err := di.BigBang(options)
	require.NoError(t, err)

	hs.UpdateHandler(universe)
	require.NoError(t, starterFunc())

	token, err := sdk.RegisterUserAccount(
		options.ClientFactory.Client().UserSvcAPI,
		"someuser",
		"pw123",
		"Some name",
	)
	require.NoError(t, err)
	userClient := options.ClientFactory.Client(sdk.WithToken(token))

	adminClient, _, err := test.AdminClient(options.ClientFactory)
	require.NoError(t, err)

	ctx := context.Background()

	_, _, err = userClient.FileSvcAPI.UploadFile(ctx).
		File(file).
		Execute()
	require.Error(t, err)

	// Need to create new file as the previous upload request closes it
	// for some reason...
	file2, cleanup2 := createTestFile(t, "Test file content")
	defer cleanup2()

	uplRsp, _, err := adminClient.FileSvcAPI.UploadFile(ctx).
		File(file2).
		Execute()
	require.NoError(t, err)
	require.Equal(t, int64(17), uplRsp.Upload.FileSize)

	timeout := time.After(5 * time.Second)
	ticker := time.NewTicker(50 * time.Millisecond)
	defer ticker.Stop()

	var uploaded bool
outer:
	for {
		select {
		case <-timeout:
			t.Fatal("Timeout reached while waiting for upload to complete")
		case <-ticker.C:
			files, err := os.ReadDir(filepath.Join(options.HomeDir, "uploads"))
			require.NoError(t, err)
			for _, f := range files {
				if f.Name() == filepath.Base(file2.Name()) {
					uploaded = true
					break outer
				}
			}
		}
	}

	require.True(t, uploaded, "File was not uploaded successfully")

	rsp, _, err := adminClient.FileSvcAPI.ListUploads(ctx).Execute()
	require.NoError(t, err)
	require.Equal(t, 1, len(rsp.Uploads))
	require.Equal(t, int64(17), rsp.Uploads[0].FileSize)

	fileRsp, fileHttpRsp, err := userClient.FileSvcAPI.ServeUpload(ctx, *rsp.Uploads[0].FileId).
		Execute()
	require.NoError(t, err)
	require.Equal(t, true, fileRsp != nil)
	bs, err := ioutil.ReadAll(fileRsp)
	require.NoError(t, err)
	require.Equal(t, "Test file content", string(bs))
	require.Equal(t, "text/plain; charset=utf-8", fileHttpRsp.Header.Get("Content-Type"))
}
