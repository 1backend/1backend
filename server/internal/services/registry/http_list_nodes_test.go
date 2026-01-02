//go:build dist
// +build dist

package registryservice_test

import (
	"context"
	"net/http/httptest"
	"strings"
	"testing"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/test"
	"github.com/1backend/1backend/server/internal/di"
	"github.com/1backend/1backend/server/internal/universe"
	"github.com/stretchr/testify/require"
)

func TestNodeId(t *testing.T) {
	ctx := context.Background()

	hs := &di.HandlerSwitcher{}
	server := httptest.NewServer(hs)
	defer server.Close()

	dbprefix := "node_id_" + sdk.Id("")

	opt1 := &universe.Options{
		Test:     true,
		Db:       "postgres",
		DbPrefix: dbprefix,
		Url:      server.URL,
	}

	nodeInfo1, err := di.BigBang(opt1)
	require.NoError(t, err)

	hs.UpdateHandler(nodeInfo1.Router)
	require.NoError(t, nodeInfo1.StarterFunc())

	require.Equal(t, true, opt1.ClientFactory != nil)
	adminClient, _, err := test.AdminClient(opt1.ClientFactory, sdk.DefaultTestAppHost)
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

	opt2 := &universe.Options{
		Test:     true,
		NodeId:   "abc",
		Db:       "postgres",
		DbPrefix: dbprefix,
		Url:      server2.URL,
	}
	nodeInfo2, err := di.BigBang(opt2)
	hs2.UpdateHandler(nodeInfo2.Router)
	require.NoError(t, nodeInfo2.StarterFunc())

	adminClient2, _, err := test.AdminClient(opt2.ClientFactory, sdk.DefaultTestAppHost)
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
		require.True(t, found, "Node with URL and ID 'abc' not found", nodesRsp2.Nodes)
	})
}
