package imageservice_test

import (
	"context"
	"image"
	"image/color"
	"image/png"
	"os"
	"testing"
	"time"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/test"

	"github.com/stretchr/testify/require"
)

func createTestImageFile(t *testing.T) (*os.File, func()) {
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			img.Set(x, y, color.RGBA{255, 0, 0, 255})
		}
	}

	tmpFile, err := os.CreateTemp("", "upload-test-*.png")
	require.NoError(t, err)

	err = png.Encode(tmpFile, img)
	require.NoError(t, err)

	_, err = tmpFile.Seek(0, 0)
	require.NoError(t, err)

	return tmpFile, func() { os.Remove(tmpFile.Name()) }
}

func TestServeUploadedImage(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		TokenExpiration: time.Second,
		Test:            true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)

	ctx := context.Background()
	_, _, err = test.MakeClients(clientFactory, sdk.DefaultTestAppHost, 1)
	require.NoError(t, err)

	adminClient, _, err := test.AdminClient(clientFactory, sdk.DefaultTestAppHost)
	require.NoError(t, err)

	imageFile, cleanup := createTestImageFile(t)
	defer cleanup()

	fileId := ""

	t.Run("upload image to node1", func(t *testing.T) {
		rsp, _, err := adminClient.FileSvcAPI.UploadFile(ctx).File(imageFile).Execute()
		require.NoError(t, err)
		fileId = rsp.Upload.FileId
	})

	t.Run("serve uploaded image", func(t *testing.T) {
		rsp, _, err := adminClient.ImageSvcAPI.ServeUploadedImage(ctx, fileId).
			Width(20).
			Height(20).
			Execute()
		require.NoError(t, err)
		require.NotNil(t, rsp)

		img, _, err := image.Decode(rsp)
		require.NoError(t, err)
		bounds := img.Bounds()
		require.Equal(t, 20, bounds.Dx())
		require.Equal(t, 20, bounds.Dy())
	})

	t.Run("serve uploaded image", func(t *testing.T) {
		rsp, _, err := adminClient.ImageSvcAPI.ServeUploadedImage(ctx, fileId).
			Width(70).
			Height(70).
			Execute()
		require.NoError(t, err)
		require.NotNil(t, rsp)

		img, _, err := image.Decode(rsp)
		require.NoError(t, err)
		bounds := img.Bounds()
		require.Equal(t, 70, bounds.Dx())
		require.Equal(t, 70, bounds.Dy())
	})

	t.Run("does not enlarge image", func(t *testing.T) {
		rsp, _, err := adminClient.ImageSvcAPI.ServeUploadedImage(ctx, fileId).
			Width(200).
			Height(200).
			Execute()
		require.NoError(t, err)
		img, _, err := image.Decode(rsp)
		require.NoError(t, err)
		bounds := img.Bounds()
		require.Equal(t, 100, bounds.Dx())
		require.Equal(t, 100, bounds.Dy())
	})
}
