/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package fileservice_test

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	openapi "github.com/openorch/openorch/clients/go"
	sdk "github.com/openorch/openorch/sdk/go"
	"github.com/openorch/openorch/sdk/go/test"
	"github.com/openorch/openorch/server/internal/di"
	fileservice "github.com/openorch/openorch/server/internal/services/file"
	types "github.com/openorch/openorch/server/internal/services/file/types"
	"github.com/stretchr/testify/require"
)

func TestDownloadFile(t *testing.T) {
	fileHostServer := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rangeHeader := r.Header.Get("Range")
			if rangeHeader != "" {
				w.Header().Set("Content-Range", "bytes 0-10/11")
				w.WriteHeader(http.StatusPartialContent)
				io.WriteString(w, "Hello world")
			} else {
				w.WriteHeader(http.StatusOK)
				io.WriteString(w, "Hello world")
			}
		}),
	)
	defer fileHostServer.Close()

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

	err = starterFunc()
	require.NoError(t, err)

	token, err := sdk.RegisterUserAccount(
		options.ClientFactory.Client().UserSvcAPI,
		"someuser",
		"pw123",
		"Some name",
	)
	require.NoError(t, err)
	userClient := options.ClientFactory.Client(sdk.WithToken(token))

	downloadUrl := fileHostServer.URL + "/somefile.txt"

	_, _, err = userClient.FileSvcAPI.DownloadFile(context.Background()).
		Body(openapi.FileSvcDownloadFileRequest{
			Url: downloadUrl,
		}).
		Execute()
	require.Error(t, err)

	adminClient, _, err := test.AdminClient(options.ClientFactory)
	require.NoError(t, err)

	_, _, err = adminClient.FileSvcAPI.DownloadFile(context.Background()).
		Body(openapi.FileSvcDownloadFileRequest{
			Url: downloadUrl,
		}).
		Execute()
	require.NoError(t, err)

	timeout := time.After(5 * time.Second)
	ticker := time.NewTicker(5 * time.Millisecond)
	defer ticker.Stop()

outer:
	for {
		select {
		case <-timeout:
			t.Fatal("Timeout reached while waiting for download to complete")
		case <-ticker.C:
			_, _, err := userClient.FileSvcAPI.GetDownload(context.Background(), downloadUrl).
				Execute()
			require.Error(t, err)

			rsp, _, err := adminClient.FileSvcAPI.GetDownload(context.Background(), downloadUrl).
				Execute()
			require.NoError(t, err)

			if rsp.Exists &&
				*rsp.Download.Status == string(types.DownloadStatusCompleted) {
				break outer
			}
		}
	}

	expectedFilePath := filepath.Join(
		options.HomeDir,
		"downloads",
		fileservice.EncodeURLtoFileName(downloadUrl),
	)
	data, err := os.ReadFile(expectedFilePath)
	require.NoError(t, err)
	require.Equal(t, "Hello world", string(data))

	fileRsp, fileHttpRsp, err := userClient.FileSvcAPI.ServeDownload(context.Background(), downloadUrl).
		Execute()
	require.NoError(t, err)
	bs, err := ioutil.ReadAll(fileRsp)
	require.NoError(t, err)
	require.Equal(t, "Hello world", string(bs))
	require.Equal(t, "text/plain; charset=utf-8", fileHttpRsp.Header.Get("Content-Type"))
}

func TestDownloadFileWithPartFile(t *testing.T) {
	fileHostServer := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rangeHeader := r.Header.Get("Range")
			if rangeHeader != "bytes=5-" {
				t.Errorf("Expected 'bytes=5-' got '%s'", rangeHeader)
			}

			w.Header().Set("Content-Range", "bytes 5-10/11")
			w.WriteHeader(http.StatusPartialContent)
			io.WriteString(w, " world")
		}),
	)
	defer fileHostServer.Close()

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

	err = starterFunc()
	require.NoError(t, err)

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

	downloadURL := fileHostServer.URL + "/file"

	partFilePath := filepath.Join(
		options.HomeDir,
		"downloads",
		fileservice.EncodeURLtoFileName(downloadURL)+".part",
	)
	if err := os.WriteFile(partFilePath, []byte("Hello"), 0644); err != nil {
		t.Fatalf("Failed to create part file: %s", err)
	}

	_, _, err = userClient.FileSvcAPI.DownloadFile(context.Background()).
		Body(openapi.FileSvcDownloadFileRequest{
			Url: downloadURL,
		}).
		Execute()
	require.Error(t, err)

	_, _, err = adminClient.FileSvcAPI.DownloadFile(context.Background()).
		Body(openapi.FileSvcDownloadFileRequest{
			Url: downloadURL,
		}).
		Execute()
	require.NoError(t, err)

	timeout := time.After(5 * time.Second)
	ticker := time.NewTicker(5 * time.Millisecond)
	defer ticker.Stop()

outer:
	for {
		select {
		case <-timeout:
			t.Fatal("Timeout reached while waiting for download to complete")
		case <-ticker.C:
			_, _, err := userClient.FileSvcAPI.GetDownload(context.Background(), downloadURL).
				Execute()
			require.Error(t, err)

			rsp, _, err := adminClient.FileSvcAPI.GetDownload(context.Background(), downloadURL).
				Execute()
			require.NoError(t, err)

			if rsp.Exists &&
				*rsp.Download.Status == string(types.DownloadStatusCompleted) {
				break outer
			}
		}
	}

	expectedFilePath := filepath.Join(
		options.HomeDir,
		"downloads",
		fileservice.EncodeURLtoFileName(downloadURL),
	)
	data, err := os.ReadFile(expectedFilePath)
	require.NoError(t, err)
	require.Equal(t, "Hello world", string(data))
}

func TestDownloadFileWithFullFile(t *testing.T) {
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

	err = starterFunc()
	require.NoError(t, err)

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

	downloadURL := "full-file"
	fullFilePath := filepath.Join(
		options.HomeDir,
		"downloads",
		fileservice.EncodeURLtoFileName(downloadURL),
	)
	require.NoError(t, os.WriteFile(fullFilePath, []byte("Hello world"), 0644))

	_, _, err = userClient.FileSvcAPI.DownloadFile(context.Background()).
		Body(openapi.FileSvcDownloadFileRequest{
			Url: downloadURL,
		}).
		Execute()
	require.Error(t, err)

	_, _, err = adminClient.FileSvcAPI.DownloadFile(context.Background()).
		Body(openapi.FileSvcDownloadFileRequest{
			Url: downloadURL,
		}).
		Execute()
	require.NoError(t, err)

	timeout := time.After(5 * time.Second)
	ticker := time.NewTicker(5 * time.Millisecond)
	defer ticker.Stop()

	var d *openapi.FileSvcDownload

outer:
	for {
		select {
		case <-timeout:
			t.Fatal("Timeout reached while waiting for download to complete")
		case <-ticker.C:
			_, _, err := userClient.FileSvcAPI.GetDownload(context.Background(), downloadURL).
				Execute()
			require.Error(t, err)

			rsp, _, err := adminClient.FileSvcAPI.GetDownload(context.Background(), downloadURL).
				Execute()
			require.NoError(t, err)

			if rsp.Exists &&
				*rsp.Download.Status == string(types.DownloadStatusCompleted) {
				d = rsp.Download
				break outer
			}
		}
	}

	require.Equal(t, int64(11), *d.DownloadedBytes)
	require.Equal(t, int64(11), *d.FileSize)
}
