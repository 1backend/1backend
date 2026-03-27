package fileservice_test

import (
	"context"
	"net/http"
	"os"
	"testing"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/test"
	"github.com/stretchr/testify/require"
)

func TestDeleteUpload(t *testing.T) {
	ctx := context.Background()

	server, err := test.StartService(test.Options{Test: true})
	require.NoError(t, err)
	defer server.Cleanup(t)

	apiFactory := client.NewApiClientFactory(server.Url)
	adminClient, adminToken, err := test.AdminClient(apiFactory, sdk.DefaultTestAppHost)
	require.NoError(t, err)

	tmpFile, err := os.CreateTemp("", "test-delete-upload-*.txt")
	require.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.WriteString("delete me")
	require.NoError(t, err)
	tmpFile.Close()

	f, err := os.Open(tmpFile.Name())
	require.NoError(t, err)
	defer f.Close()

	uplRsp, _, err := adminClient.FileSvcAPI.UploadFile(ctx).File(f).Execute()
	require.NoError(t, err)
	require.NotEmpty(t, uplRsp.Upload.FileId)

	req, err := http.NewRequest(http.MethodDelete, server.Url+"/file-svc/upload/"+uplRsp.Upload.FileId, nil)
	require.NoError(t, err)
	req.Header.Set("Authorization", "Bearer "+adminToken)

	httpRsp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer httpRsp.Body.Close()
	require.Equal(t, http.StatusOK, httpRsp.StatusCode)

	listRsp, _, err := adminClient.FileSvcAPI.ListUploads(ctx).Execute()
	require.NoError(t, err)
	require.Empty(t, listRsp.Uploads)

	_, serveHttpRsp, err := adminClient.FileSvcAPI.ServeUpload(ctx, uplRsp.Upload.FileId).Execute()
	require.Error(t, err)
	require.Equal(t, http.StatusNotFound, serveHttpRsp.StatusCode)
}
