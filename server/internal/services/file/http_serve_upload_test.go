package fileservice_test

import (
	"context"
	"io"
	"net/http"
	"os"
	"testing"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/test"
	"github.com/stretchr/testify/require"
)

func TestServeUpload(t *testing.T) {
	// 1. Setup environment
	ctx := context.Background()

	// Start the full service (Local Storage by default in tests)
	server, err := test.StartService(test.Options{
		Test: true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	// Create an admin client
	apiFactory := client.NewApiClientFactory(server.Url)
	adminClient, _, err := test.AdminClient(apiFactory, sdk.DefaultTestAppHost)
	require.NoError(t, err)

	// Prepare a test file
	tmpFile, err := os.CreateTemp("", "test-upload-*.txt")
	require.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	content := "Hello 1Backend Storage"
	_, err = tmpFile.WriteString(content)
	require.NoError(t, err)
	tmpFile.Close()

	var fileId string

	// 2. Test Upload
	t.Run("upload file works", func(t *testing.T) {
		f, err := os.Open(tmpFile.Name())
		require.NoError(t, err)
		defer f.Close()

		_, _, err = adminClient.FileSvcAPI.UploadFile(ctx).
			File(f).
			Execute()
		require.NoError(t, err)
	})

	// 3. Test Metadata Retrieval (List)
	t.Run("list uploads shows the file", func(t *testing.T) {
		rsp, _, err := adminClient.FileSvcAPI.ListUploads(ctx).Execute()
		require.NoError(t, err)
		require.Equal(t, 1, len(rsp.Uploads))

		fileId = rsp.Uploads[0].FileId
		require.NotEmpty(t, fileId)
		require.Equal(t, int64(len(content)), rsp.Uploads[0].FileSize)
	})

	// 4. Test Serve (The core functionality)
	t.Run("serve upload returns correct content and type", func(t *testing.T) {
		fileRsp, httpRsp, err := adminClient.FileSvcAPI.ServeUpload(ctx, fileId).Execute()
		require.NoError(t, err)
		require.NotNil(t, fileRsp)
		defer fileRsp.Close()

		// Verify Content-Type (Detection works because ID includes extension)
		require.Equal(t, http.StatusOK, httpRsp.StatusCode)
		require.Equal(t, "text/plain; charset=utf-8", httpRsp.Header.Get("Content-Type"))

		// Verify Body
		bs, err := io.ReadAll(fileRsp)
		require.NoError(t, err)
		require.Equal(t, content, string(bs))
	})

	// 5. Test 404
	t.Run("serving non-existent file returns 404", func(t *testing.T) {
		_, httpRsp, err := adminClient.FileSvcAPI.ServeUpload(ctx, "non-existent.txt").Execute()
		require.Error(t, err)
		require.Equal(t, http.StatusNotFound, httpRsp.StatusCode)
	})
}
