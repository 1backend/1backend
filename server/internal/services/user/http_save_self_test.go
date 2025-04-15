/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package userservice_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/test"
)

func TestSaveSelf(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		Test: true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)

	manyClients, _, err := test.MakeClients(clientFactory, 1)
	require.NoError(t, err)

	client1 := manyClients[0]

	ctx := context.Background()

	t.Run("user updates their name", func(t *testing.T) {
		_, _, err := client1.UserSvcAPI.SaveSelf(ctx).Body(
			openapi.UserSvcSaveSelfRequest{
				Name: openapi.PtrString("New Name"),
			}).Execute()

		require.Error(t, err)
	})

}
