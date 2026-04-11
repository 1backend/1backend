package fileservice_test

import (
	"context"
	"io"
	"net/http/httptest"
	"testing"
	"time"

	openapi "github.com/1backend/1backend/clients/go"
	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/test"
	"github.com/1backend/1backend/server/internal/di"
	"github.com/1backend/1backend/server/internal/universe"
	"github.com/stretchr/testify/require"
)

func TestServeUploadRecordsLastAccessedAt(t *testing.T) {
	hs := &di.HandlerSwitcher{}
	server := httptest.NewServer(hs)
	defer server.Close()

	options := &universe.Options{
		Test: true,
		Url:  server.URL,
	}
	universe, err := di.BigBang(options)
	require.NoError(t, err)

	hs.UpdateHandler(universe.Router)
	require.NoError(t, universe.StarterFunc())

	adminClient, _, err := test.AdminClient(options.ClientFactory, sdk.DefaultAppHost)
	require.NoError(t, err)

	ctx := context.Background()
	file, cleanup := createTestFile(t, "access-tracking-content")
	defer cleanup()

	uplRsp, _, err := adminClient.FileSvcAPI.UploadFile(ctx).File(file).Execute()
	require.NoError(t, err)
	require.NotEmpty(t, uplRsp.Upload.FileId)

	initialUpload := waitForUploadByFileID(t, adminClient, ctx, uplRsp.Upload.FileId, 3*time.Second)
	require.Nil(t, initialUpload.LastAccessedAt, "lastAccessedAt should be empty before serving")
	initialUpdatedAt := initialUpload.UpdatedAt

	// First access: should be flushed by interval in test mode (250ms).
	serveUploadAndReadBody(t, adminClient, ctx, uplRsp.Upload.FileId)
	firstAccessUpload := waitForLastAccessedAt(t, adminClient, ctx, uplRsp.Upload.FileId, 4*time.Second)
	require.NotNil(t, firstAccessUpload.LastAccessedAt)
	require.Equal(t, initialUpdatedAt, firstAccessUpload.UpdatedAt, "serving should not mutate updatedAt")

	firstAccessTs, err := time.Parse(time.RFC3339Nano, *firstAccessUpload.LastAccessedAt)
	require.NoError(t, err)

	// Repeated access: verify lastAccessedAt advances over time via interval-based flushing.
	time.Sleep(300 * time.Millisecond)
	serveUploadAndReadBody(t, adminClient, ctx, uplRsp.Upload.FileId)

	require.Eventually(t, func() bool {
		upload := waitForUploadByFileID(t, adminClient, ctx, uplRsp.Upload.FileId, 500*time.Millisecond)
		if upload.LastAccessedAt == nil {
			return false
		}
		nextTs, err := time.Parse(time.RFC3339Nano, *upload.LastAccessedAt)
		if err != nil {
			return false
		}
		return nextTs.After(firstAccessTs)
	}, 3*time.Second, 100*time.Millisecond, "expected lastAccessedAt to advance after repeated serves")

	finalUpload := waitForUploadByFileID(t, adminClient, ctx, uplRsp.Upload.FileId, 2*time.Second)
	require.Equal(t, initialUpdatedAt, finalUpload.UpdatedAt, "updatedAt must remain stable across access tracking updates")
}

func serveUploadAndReadBody(
	t *testing.T,
	adminClient *openapi.APIClient,
	ctx context.Context,
	fileID string,
) {
	t.Helper()

	stream, _, err := adminClient.FileSvcAPI.ServeUpload(ctx, fileID).Execute()
	require.NoError(t, err)
	defer stream.Close()

	_, err = io.ReadAll(stream)
	require.NoError(t, err)
}

func waitForLastAccessedAt(
	t *testing.T,
	adminClient *openapi.APIClient,
	ctx context.Context,
	fileID string,
	timeout time.Duration,
) *openapi.FileSvcUpload {
	t.Helper()

	var upload *openapi.FileSvcUpload
	require.Eventually(t, func() bool {
		upload = waitForUploadByFileID(t, adminClient, ctx, fileID, 500*time.Millisecond)
		return upload.LastAccessedAt != nil
	}, timeout, 100*time.Millisecond, "expected lastAccessedAt to be recorded")

	return upload
}

func waitForUploadByFileID(
	t *testing.T,
	adminClient *openapi.APIClient,
	ctx context.Context,
	fileID string,
	timeout time.Duration,
) *openapi.FileSvcUpload {
	t.Helper()

	deadline := time.Now().Add(timeout)
	for {
		rsp, _, err := adminClient.FileSvcAPI.ListUploads(ctx).Execute()
		require.NoError(t, err)
		for _, upload := range rsp.Uploads {
			if upload.FileId == fileID {
				u := upload
				return &u
			}
		}

		if time.Now().After(deadline) {
			require.FailNow(t, "timed out waiting for upload row", "fileId=%s", fileID)
		}
		time.Sleep(50 * time.Millisecond)
	}
}
