//go:build dist
// +build dist

package fileservice_test

import (
	"context"
	"io/ioutil"

	"net/http/httptest"
	"testing"

	sdk "github.com/openorch/openorch/sdk/go"

	"github.com/openorch/openorch/sdk/go/test"
	"github.com/openorch/openorch/server/internal/di"
	"github.com/openorch/openorch/server/internal/node"

	node_types "github.com/openorch/openorch/server/internal/node/types"
	"github.com/stretchr/testify/require"
)

func TestServeUploadProxy(t *testing.T) {
	file, cleanup := createTestFile(t, "Test file things")
	defer cleanup()

	ctx := context.Background()

	hs := &di.HandlerSwitcher{}
	server := httptest.NewServer(hs)
	defer server.Close()

	dbprefix := sdk.Id("node_id")

	opt1 := &node_types.Options{
		Test:     true,
		NodeId:   "node1",
		Db:       "postgres",
		DbPrefix: dbprefix,
		Address:  server.URL,
	}

	nodeInfo1, err := node.Start(opt1)
	require.NoError(t, err)

	hs.UpdateHandler(nodeInfo1.Router)
	require.NoError(t, nodeInfo1.StarterFunc())

	require.Equal(t, true, opt1.ClientFactory != nil)
	adminClient, _, err := test.AdminClient(opt1.ClientFactory)
	require.NoError(t, err)

	t.Run("upload to node1 works", func(t *testing.T) {
		_, _, err = adminClient.FileSvcAPI.UploadFile(ctx).
			File(file).
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
		DbPrefix: dbprefix,
		Address:  server2.URL,
	}
	nodeInfo2, err := node.Start(opt2)
	hs2.UpdateHandler(nodeInfo2.Router)
	require.NoError(t, nodeInfo2.StarterFunc())

	adminClient2, _, err := test.AdminClient(opt2.ClientFactory)
	require.NoError(t, err)

	var fileId string

	t.Run("upload serve from node2 works", func(t *testing.T) {
		// listing can be done from either node as the list comes from a DB
		rsp, _, err := adminClient.FileSvcAPI.ListUploads(ctx).Execute()
		require.NoError(t, err)
		require.Equal(t, 1, len(rsp.Uploads))
		require.Equal(t, int64(16), rsp.Uploads[0].FileSize)

		fileId = *rsp.Uploads[0].FileId

		fileRsp, fileHttpRsp, err := adminClient2.FileSvcAPI.ServeUpload(ctx, fileId).Execute()
		require.NoError(t, err)
		require.Equal(t, true, fileRsp != nil)
		bs, err := ioutil.ReadAll(fileRsp)
		require.NoError(t, err)
		require.Equal(t, "Test file things", string(bs))
		require.Equal(t, "text/plain; charset=utf-8", fileHttpRsp.Header.Get("Content-Type"))
	})

	t.Run("upload serve should fail if node is down", func(t *testing.T) {
		server.Close()

		_, _, err := adminClient2.FileSvcAPI.ServeUpload(ctx, fileId).Execute()
		require.Error(t, err)
	})
}
