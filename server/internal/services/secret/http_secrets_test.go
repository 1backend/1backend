/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package secretservice_test

import (
	"context"
	"testing"

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/test"
	"github.com/stretchr/testify/require"
)

func TestSecretService(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		Test: true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)

	manyClients, _, err := test.MakeClients(clientFactory, 2)
	require.NoError(t, err)
	client1 := manyClients[0]
	client2 := manyClients[1]

	adminClient, _, err := test.AdminClient(clientFactory)
	require.NoError(t, err)

	ctx := context.Background()

	t.Run("reading non-existent secret", func(t *testing.T) {
		readRsp, _, err := client1.SecretSvcAPI.ListSecrets(ctx).
			Body(openapi.SecretSvcListSecretsRequest{
				Key: openapi.PtrString("nonexistent"),
			}).
			Execute()

		require.NoError(t, err)
		require.Equal(t, 0, len(readRsp.Secrets))
	})

	t.Run("creating a secret without prefix as user", func(t *testing.T) {
		_, _, err := client1.SecretSvcAPI.SaveSecrets(ctx).
			Body(openapi.SecretSvcSaveSecretsRequest{
				Secrets: []openapi.SecretSvcSecret{
					{
						Key:   openapi.PtrString("non-prefixed"),
						Value: openapi.PtrString("value"),
					}},
			}).
			Execute()

		require.Error(t, err)
	})

	t.Run("as an admin trying to claim prefixed works", func(t *testing.T) {
		_, _, err := adminClient.SecretSvcAPI.SaveSecrets(ctx).
			Body(openapi.SecretSvcSaveSecretsRequest{
				Secrets: []openapi.SecretSvcSecret{
					{
						Key:   openapi.PtrString("non-prefixed"),
						Value: openapi.PtrString("value"),
					}},
			}).
			Execute()
		require.NoError(t, err)
	})

	t.Run("user cannot read secret owned by admin", func(t *testing.T) {
		readRsp, _, err := client1.SecretSvcAPI.ListSecrets(ctx).
			Body(openapi.SecretSvcListSecretsRequest{
				Key: openapi.PtrString("non-prefixed"),
			}).
			Execute()

		require.NoError(t, err)
		require.Equal(t, 0, len(readRsp.Secrets))
	})

	t.Run("user 1 can write and read, but user 2 will not have access", func(t *testing.T) {
		readRsp, _, err := client1.SecretSvcAPI.ListSecrets(ctx).
			Body(openapi.SecretSvcListSecretsRequest{
				Key: openapi.PtrString("test-user-slug-0/non-existent"),
			}).
			Execute()

		require.NoError(t, err)
		require.Equal(t, 0, len(readRsp.Secrets))

		_, _, err = client1.SecretSvcAPI.SaveSecrets(ctx).
			Body(openapi.SecretSvcSaveSecretsRequest{
				Secrets: []openapi.SecretSvcSecret{
					{
						Key:   openapi.PtrString("test-user-slug-0/non-existent"),
						Value: openapi.PtrString("value"),
					}},
			}).
			Execute()

		require.NoError(t, err)

		readRsp, _, err = client1.SecretSvcAPI.ListSecrets(ctx).
			Body(openapi.SecretSvcListSecretsRequest{
				Key: openapi.PtrString("test-user-slug-0/non-existent"),
			}).
			Execute()

		require.NoError(t, err)
		require.Equal(t, 1, len(readRsp.Secrets))

		readRsp, _, err = client2.SecretSvcAPI.ListSecrets(ctx).
			Body(openapi.SecretSvcListSecretsRequest{
				Key: openapi.PtrString("test-user-slug-1/non-existent"),
			}).
			Execute()

		require.NoError(t, err)
		require.Equal(t, 0, len(readRsp.Secrets))
	})

	t.Run("as a user writing already encrypted values", func(t *testing.T) {
		rsp, _, err := client1.SecretSvcAPI.EncryptValue(ctx).
			Body(openapi.SecretSvcEncryptValueRequest{
				Value: openapi.PtrString("value"),
			}).
			Execute()
		require.NoError(t, err)

		_, _, err = client1.SecretSvcAPI.SaveSecrets(ctx).
			Body(openapi.SecretSvcSaveSecretsRequest{
				Secrets: []openapi.SecretSvcSecret{
					{
						Key:       openapi.PtrString("test-user-slug-0/enc1"),
						Value:     rsp.Value,
						Encrypted: openapi.PtrBool(true),
					}},
			}).
			Execute()
		require.NoError(t, err)

		readRsp, _, err := client1.SecretSvcAPI.ListSecrets(ctx).
			Body(openapi.SecretSvcListSecretsRequest{
				Key: openapi.PtrString("test-user-slug-0/enc1"),
			}).
			Execute()
		require.NoError(t, err)
		require.Equal(t, 1, len(readRsp.Secrets))
		require.Equal(t, "value", *readRsp.Secrets[0].Value)
		//require.Equal(t, true, *readRsp.Secrets[0].Encrypted)
		//require.Equal(t, "test-user-slug-0/enc1", *readRsp.Secrets[0].Key)

	})

	t.Run("list multiple", func(t *testing.T) {
		_, _, err := adminClient.SecretSvcAPI.SaveSecrets(ctx).
			Body(openapi.SecretSvcSaveSecretsRequest{
				Secrets: []openapi.SecretSvcSecret{
					{
						Key:   openapi.PtrString("a"),
						Value: openapi.PtrString("value"),
					}},
			}).
			Execute()
		require.NoError(t, err)

		_, _, err = adminClient.SecretSvcAPI.SaveSecrets(ctx).
			Body(openapi.SecretSvcSaveSecretsRequest{
				Secrets: []openapi.SecretSvcSecret{
					{
						Key:   openapi.PtrString("b"),
						Value: openapi.PtrString("value"),
					}},
			}).
			Execute()
		require.NoError(t, err)

		readRsp, _, err := adminClient.SecretSvcAPI.ListSecrets(ctx).
			Body(openapi.SecretSvcListSecretsRequest{
				Keys: []string{"a", "b"},
			}).
			Execute()

		require.NoError(t, err)
		require.Equal(t, 2, len(readRsp.Secrets))
		require.Equal(t, "value", *readRsp.Secrets[0].Value)
	})

}
