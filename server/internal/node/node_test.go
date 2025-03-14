//go:build dist
// +build dist

package node

import (
	"context"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	openapi "github.com/openorch/openorch/clients/go"
	sdk "github.com/openorch/openorch/sdk/go"
	"github.com/openorch/openorch/sdk/go/test"
	"github.com/openorch/openorch/server/internal/di"
	node_types "github.com/openorch/openorch/server/internal/node/types"
	"github.com/stretchr/testify/require"
)

func TestStart(t *testing.T) {
	hs1 := &di.HandlerSwitcher{}
	server1 := httptest.NewServer(hs1)
	defer server1.Close()

	dbprefix := sdk.Id("node_start")

	options1 := &node_types.Options{
		Db:       "postgres",
		DbPrefix: dbprefix,
		Address:  server1.URL,
	}

	nodeInfo1, err := Start(options1)
	require.NoError(t, err)

	hs1.UpdateHandler(nodeInfo1.Router)
	// @todo Why is this called here and also down below?
	// If I remove this there is a config service error
	// which i think points to a syncronization issue.
	err = nodeInfo1.StarterFunc()
	require.NoError(t, err)

	hs2 := &di.HandlerSwitcher{}
	server2 := httptest.NewServer(hs1)
	defer server1.Close()

	options2 := &node_types.Options{
		Db:       "postgres",
		DbPrefix: dbprefix,
		Address:  server2.URL,
	}
	nodeInfo2, err := Start(options2)
	require.NoError(t, err)

	hs2.UpdateHandler(nodeInfo2.Router)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		err := nodeInfo1.StarterFunc()
		wg.Done()
		require.NoError(t, err)

	}()

	go func() {
		err := nodeInfo2.StarterFunc()
		wg.Done()
		require.NoError(t, err)
	}()

	wg.Wait()

	// List nodes

	c := 0
	for {
		time.Sleep(100 * time.Millisecond)
		c++

		adminClient, _, err := test.AdminClient(nodeInfo1.Options.ClientFactory)
		require.NoError(t, err)

		rsp, _, err := adminClient.RegistrySvcAPI.ListNodes(context.Background()).
			Body(openapi.RegistrySvcListNodesRequest{}).
			Execute()
		require.NoError(t, err)

		if len(rsp.Nodes) == 2 {
			break
		}
		if c > 10 {
			require.Equal(t, 2, len(rsp.Nodes))
		}

	}
}
