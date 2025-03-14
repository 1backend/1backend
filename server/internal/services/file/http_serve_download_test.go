//go:build dist
// +build dist

package fileservice_test

import (
	"context"
	"io"
	"io/ioutil"

	"net/http"
	"net/http/httptest"
	"testing"

	openapi "github.com/openorch/openorch/clients/go"
	sdk "github.com/openorch/openorch/sdk/go"

	"github.com/openorch/openorch/sdk/go/test"
	"github.com/openorch/openorch/server/internal/di"
	"github.com/openorch/openorch/server/internal/node"

	node_types "github.com/openorch/openorch/server/internal/node/types"
	"github.com/stretchr/testify/require"
)

func TestServeDownloadProxy(t *testing.T) {
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

	ctx := context.Background()

	hs := &di.HandlerSwitcher{}
	server := httptest.NewServer(hs)
	defer server.Close()

	dbPrefix := sdk.Id("node_id_")

	opt1 := &node_types.Options{
		Test:     true,
		NodeId:   "node1",
		Db:       "postgres",
		DbPrefix: dbPrefix,
		Address:  server.URL,
	}

	nodeInfo1, err := node.Start(opt1)
	require.NoError(t, err)

	hs.UpdateHandler(nodeInfo1.Router)
	require.NoError(t, nodeInfo1.StarterFunc())

	require.Equal(t, true, opt1.ClientFactory != nil)
	adminClient, _, err := test.AdminClient(opt1.ClientFactory)
	require.NoError(t, err)

	downloadUrl := fileHostServer.URL + "/somefile.txt"

	t.Run("download to node1 works", func(t *testing.T) {
		_, _, err = adminClient.FileSvcAPI.DownloadFile(ctx).
			Body(openapi.FileSvcDownloadFileRequest{
				Url: downloadUrl,
			}).
			Execute()
		require.NoError(t, err)
	})

	hs2 := &di.HandlerSwitcher{}
	server2 := httptest.NewServer(hs2)
	defer server2.Close()

	opt2 := &node_types.Options{
		Test:     true,
		NodeId:   "node2",
		Db:       "postgres",
		DbPrefix: dbPrefix,
		Address:  server2.URL,
	}
	nodeInfo2, err := node.Start(opt2)
	hs2.UpdateHandler(nodeInfo2.Router)
	require.NoError(t, nodeInfo2.StarterFunc())

	adminClient2, _, err := test.AdminClient(opt2.ClientFactory)
	require.NoError(t, err)

	t.Run("download serve from node2 works", func(t *testing.T) {
		// listing can be done from either node as the list comes from a DB
		rsp, _, err := adminClient.FileSvcAPI.ListDownloads(ctx).Execute()
		require.NoError(t, err)
		require.Equal(t, 1, len(rsp.Downloads))
		require.Equal(t, int64(11), *rsp.Downloads[0].FileSize)

		require.Equal(t, downloadUrl, *rsp.Downloads[0].Url)

		fileRsp, fileHttpRsp, err := adminClient2.FileSvcAPI.ServeDownload(ctx, downloadUrl).
			Execute()
		require.NoError(t, err)
		require.Equal(t, true, fileRsp != nil)
		bs, err := ioutil.ReadAll(fileRsp)
		require.NoError(t, err)
		require.Equal(t, "Hello world", string(bs))
		require.Equal(t, "text/plain; charset=utf-8", fileHttpRsp.Header.Get("Content-Type"))
	})

	t.Run("download serve should fail if node is down", func(t *testing.T) {
		server.Close()

		_, _, err := adminClient2.FileSvcAPI.ServeDownload(ctx, downloadUrl).Execute()
		require.Error(t, err)
	})
}
