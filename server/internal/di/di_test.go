//go:build dist
// +build dist

package di_test

import (
	"context"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	openapi "github.com/1backend/1backend/clients/go"
	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/test"
	"github.com/1backend/1backend/server/internal/di"
	"github.com/1backend/1backend/server/internal/universe"
	"github.com/stretchr/testify/require"
)

func TestStart(t *testing.T) {
	hs1 := &di.HandlerSwitcher{}
	server1 := httptest.NewServer(hs1)
	defer server1.Close()

	dbprefix := sdk.Id("node_start")

	options1 := &universe.Options{
		Test:     true,
		Db:       "postgres",
		DbPrefix: dbprefix,
		Url:      server1.URL,
	}

	universe1, err := di.BigBang(options1)
	require.NoError(t, err)

	hs1.UpdateHandler(universe1.Router)
	// @todo Why is this called here and also down below?
	// If I remove this there is a config service error
	// which i think points to a syncronization issue.
	err = universe1.StarterFunc()
	require.NoError(t, err)

	hs2 := &di.HandlerSwitcher{}
	server2 := httptest.NewServer(hs1)
	defer server1.Close()

	options2 := &universe.Options{
		Test:     true,
		Db:       "postgres",
		DbPrefix: dbprefix,
		Url:      server2.URL,
	}
	universe2, err := di.BigBang(options2)
	require.NoError(t, err)

	hs2.UpdateHandler(universe2.Router)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		err := universe1.StarterFunc()
		wg.Done()
		require.NoError(t, err)

	}()

	go func() {
		err := universe2.StarterFunc()
		wg.Done()
		require.NoError(t, err)
	}()

	wg.Wait()

	// List nodes

	c := 0
	for {
		time.Sleep(100 * time.Millisecond)
		c++

		adminClient, _, err := test.AdminClient(universe1.Options.ClientFactory, sdk.DefaultTestAppHost)
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
