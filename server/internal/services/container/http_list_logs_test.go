/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package containerservice_test

import (
	"context"
	"net/http/httptest"
	"testing"
	"time"

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/test"
	"github.com/1backend/1backend/server/internal/di"
	"github.com/1backend/1backend/server/internal/universe"
	"github.com/stretchr/testify/require"
)

func TestListLogs(t *testing.T) {
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

	err = universe.StarterFunc()
	require.NoError(t, err)

	adminClient, _, err := test.AdminClient(options.ClientFactory)
	require.NoError(t, err)

	ctx := context.Background()

	t.Run("run container", func(t *testing.T) {
		_, _, err := adminClient.ContainerSvcAPI.RunContainer(ctx).
			Body(openapi.ContainerSvcRunContainerRequest{
				Image: "nginx:latest",
				Ports: []openapi.ContainerSvcPortMapping{
					{
						Internal: 9080,
						Host:     9081,
					},
				},
				Names: []string{"test-container"},
				Hash:  openapi.PtrString("abc123"),
				Envs: []openapi.ContainerSvcEnvVar{
					{
						Key:   "ENV_VAR",
						Value: "value",
					},
				},
				Labels: []openapi.ContainerSvcLabel{
					{
						Key:   "app",
						Value: "test",
					},
				},
				Keeps: []openapi.ContainerSvcKeep{
					{
						Path: "/data",
					},
				},
			}).
			Execute()

		require.NoError(t, err)
	})

	time.Sleep(1000 * time.Millisecond)

	t.Run("list logs", func(t *testing.T) {
		rsp, _, err := adminClient.ContainerSvcAPI.ListContainerLogs(ctx).
			Body(openapi.ContainerSvcListLogsRequest{}).
			Execute()

		require.NoError(t, err)
		require.Equal(t, true, len(rsp.Logs) > 0)
	})
}
