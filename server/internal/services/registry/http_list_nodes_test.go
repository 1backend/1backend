//go:build dist
// +build dist

package registryservice_test

import (
	"context"
	"net/http/httptest"
	"strings"
	"testing"

	sdk "github.com/openorch/openorch/sdk/go"
	"github.com/openorch/openorch/sdk/go/test"
	"github.com/openorch/openorch/server/internal/di"
	"github.com/openorch/openorch/server/internal/node"
	node_types "github.com/openorch/openorch/server/internal/node/types"
	"github.com/stretchr/testify/require"
)

func TestNodeId(t *testing.T) {
	ctx := context.Background()

	hs := &di.HandlerSwitcher{}
	server := httptest.NewServer(hs)
	defer server.Close()

	dbprefix := sdk.Id("node_id")

	opt1 := &node_types.Options{
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

	t.Run("first node appears in node list", func(t *testing.T) {
		nodesRsp, _, err := adminClient.RegistrySvcAPI.ListNodes(ctx).Execute()
		require.NoError(t, err)
		require.Equal(t, 1, len(nodesRsp.Nodes))
		require.Equal(t, nodesRsp.Nodes[0].Url, server.URL)
		require.Equal(t, true, strings.Contains(nodesRsp.Nodes[0].Id, "node_"))
	})

	hs2 := &di.HandlerSwitcher{}
	server2 := httptest.NewServer(hs2)
	defer server2.Close()

	opt2 := &node_types.Options{
		NodeId:   "abc",
		Db:       "postgres",
		DbPrefix: dbprefix,
		Address:  server2.URL,
	}
	nodeInfo2, err := node.Start(opt2)
	hs2.UpdateHandler(nodeInfo2.Router)
	require.NoError(t, nodeInfo2.StarterFunc())

	adminClient2, _, err := test.AdminClient(opt2.ClientFactory)
	require.NoError(t, err)

	t.Run("second node appears in node list", func(t *testing.T) {
		nodesRsp2, _, err := adminClient2.RegistrySvcAPI.ListNodes(ctx).Execute()
		require.NoError(t, err)

		require.Equal(t, 2, len(nodesRsp2.Nodes), nodesRsp2.Nodes)

		found := false
		for _, node := range nodesRsp2.Nodes {
			if node.Url == server2.URL {
				found = true
				break
			}
		}
		require.True(t, found, "URL not found in node list")

		found = false
		for _, node := range nodesRsp2.Nodes {
			if node.Url == server2.URL && node.Id == "abc" {
				found = true
				break
			}
		}
		require.True(t, found, "Node with URL and ID 'abc' not found")
	})
}
