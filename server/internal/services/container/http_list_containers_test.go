/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package containerservice_test

import (
	"context"
	"net/http/httptest"
	"testing"
	"time"

	openapi "github.com/openorch/openorch/clients/go"
	"github.com/openorch/openorch/sdk/go/test"
	"github.com/openorch/openorch/server/internal/di"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

func TestListContainers(t *testing.T) {
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

	adminClient, _, err := test.AdminClient(options.ClientFactory)
	require.NoError(t, err)

	ctx := context.Background()

	t.Run("run container", func(t *testing.T) {
		_, _, err := adminClient.ContainerSvcAPI.RunContainer(ctx).Body(openapi.ContainerSvcRunContainerRequest{
			Image: "nginx:latest",
			Ports: []openapi.ContainerSvcPortMapping{
				{
					Internal: 9080,
					Host:     9082,
				},
			},
			Names: []string{"test-container-2"},
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
		}).Execute()

		require.NoError(t, err)
	})

	time.Sleep(1100 * time.Millisecond)

	t.Run("list containers", func(t *testing.T) {
		rsp, _, err := adminClient.ContainerSvcAPI.ListContainers(ctx).Body(openapi.ContainerSvcListContainersRequest{}).Execute()
		require.NoError(t, err)

		found := false
		for _, c := range rsp.Containers {
			if lo.Contains(c.Names, "test-container-2") {
				require.Equal(t, true, len(rsp.Containers) > 0)
				require.Equal(t, "nginx:latest", *c.Image)

				portFound := false
				for _, port := range c.Ports {
					if port.Internal == 9080 && port.Host == 9082 {
						portFound = true
					}
				}
				require.Equal(t, true, portFound)
				found = true
			}
		}

		require.Equal(t, true, found)
	})
}
