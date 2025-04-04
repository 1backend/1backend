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
	"github.com/1backend/1backend/sdk/go/boot"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/test"
)

func TestSavePermissions(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		Test: true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)

	token, err := boot.RegisterUserAccount(
		clientFactory.Client().UserSvcAPI,
		"someuser",
		"pw123",
		"Some name",
	)
	require.NoError(t, err)
	userClient := clientFactory.Client(client.WithToken(token.Token))

	ctx := context.Background()

	_, _, err = userClient.UserSvcAPI.SavePermissions(ctx).Body(openapi.UserSvcSavePermissionsRequest{
		Permissions: []openapi.UserSvcPermission{
			{
				Id: openapi.PtrString("someuser:myperm"),
			},
		},
	}).Execute()
	require.NoError(t, err)

	_, _, err = userClient.UserSvcAPI.SavePermissions(ctx).Body(openapi.UserSvcSavePermissionsRequest{
		Permissions: []openapi.UserSvcPermission{
			{
				Id: openapi.PtrString("other:myperm"),
			},
		},
	}).Execute()
	require.Error(t, err)
}
