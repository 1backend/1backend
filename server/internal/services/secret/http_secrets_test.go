/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package secretservice_test

import (
	"context"
	"testing"

	openapi "github.com/1backend/1backend/clients/go"
	sdk "github.com/1backend/1backend/sdk/go"
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

	manyClients, _, err := test.MakeClients(clientFactory, sdk.DefaultTestAppHost, 2)
	require.NoError(t, err)
	client1 := manyClients[0]
	client2 := manyClients[1]

	adminClient, _, err := test.AdminClient(clientFactory, sdk.DefaultTestAppHost)
	require.NoError(t, err)

	ctx := context.Background()

	t.Run("reading non-existent secret", func(t *testing.T) {
		readRsp, _, err := client1.SecretSvcAPI.ListSecrets(ctx).
			Body(openapi.SecretSvcListSecretsRequest{
				Id: openapi.PtrString("nonexistent"),
			}).
			Execute()

		require.NoError(t, err)
		require.Equal(t, 0, len(readRsp.Secrets))
	})

	t.Run("creating a secret without prefix as user", func(t *testing.T) {
		_, _, err := client1.SecretSvcAPI.SaveSecrets(ctx).
			Body(openapi.SecretSvcSaveSecretsRequest{
				Secrets: []openapi.SecretSvcSecretInput{
					{
						Id:    "non-prefixed",
						Value: openapi.PtrString("value"),
					}},
			}).
			Execute()

		require.Error(t, err)
	})

	t.Run("as an admin trying to claim prefixed works", func(t *testing.T) {
		_, _, err := adminClient.SecretSvcAPI.SaveSecrets(ctx).
			Body(openapi.SecretSvcSaveSecretsRequest{
				Secrets: []openapi.SecretSvcSecretInput{
					{
						Id:    "non-prefixed",
						Value: openapi.PtrString("value"),
					}},
			}).
			Execute()
		require.NoError(t, err)
	})

	t.Run("user cannot read secret owned by admin", func(t *testing.T) {
		readRsp, _, err := client1.SecretSvcAPI.ListSecrets(ctx).
			Body(openapi.SecretSvcListSecretsRequest{
				Id: openapi.PtrString("non-prefixed"),
			}).
			Execute()

		require.NoError(t, err)
		require.Equal(t, 0, len(readRsp.Secrets))
	})

	t.Run("user 1 can write and read, but user 2 will not have access", func(t *testing.T) {
		readRsp, _, err := client1.SecretSvcAPI.ListSecrets(ctx).
			Body(openapi.SecretSvcListSecretsRequest{
				Id: openapi.PtrString("test-user-slug-0/non-existent"),
			}).
			Execute()

		require.NoError(t, err)
		require.Equal(t, 0, len(readRsp.Secrets))

		_, _, err = client1.SecretSvcAPI.SaveSecrets(ctx).
			Body(openapi.SecretSvcSaveSecretsRequest{
				Secrets: []openapi.SecretSvcSecretInput{
					{
						Id:    "test-user-slug-0/non-existent",
						Value: openapi.PtrString("value"),
					}},
			}).
			Execute()

		require.NoError(t, err)

		readRsp, _, err = client1.SecretSvcAPI.ListSecrets(ctx).
			Body(openapi.SecretSvcListSecretsRequest{
				Id: openapi.PtrString("test-user-slug-0/non-existent"),
			}).
			Execute()

		require.NoError(t, err)
		require.Equal(t, 1, len(readRsp.Secrets))

		readRsp, _, err = client2.SecretSvcAPI.ListSecrets(ctx).
			Body(openapi.SecretSvcListSecretsRequest{
				Id: openapi.PtrString("test-user-slug-1/non-existent"),
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
				Secrets: []openapi.SecretSvcSecretInput{
					{
						Id:        "test-user-slug-0/enc1",
						Value:     rsp.Value,
						Encrypted: openapi.PtrBool(true),
					}},
			}).
			Execute()
		require.NoError(t, err)

		readRsp, _, err := client1.SecretSvcAPI.ListSecrets(ctx).
			Body(openapi.SecretSvcListSecretsRequest{
				Id: openapi.PtrString("test-user-slug-0/enc1"),
			}).
			Execute()
		require.NoError(t, err)
		require.Equal(t, 1, len(readRsp.Secrets), readRsp)
		require.Equal(t, "value", readRsp.Secrets[0].Value)
		//require.Equal(t, true, *readRsp.Secrets[0].Encrypted)
		//require.Equal(t, "test-user-slug-0/enc1", *readRsp.Secrets[0].Id)

	})

	t.Run("list multiple", func(t *testing.T) {
		_, _, err := adminClient.SecretSvcAPI.SaveSecrets(ctx).
			Body(openapi.SecretSvcSaveSecretsRequest{
				Secrets: []openapi.SecretSvcSecretInput{
					{
						Id:    "a",
						Value: openapi.PtrString("value"),
					}},
			}).
			Execute()
		require.NoError(t, err)

		_, _, err = adminClient.SecretSvcAPI.SaveSecrets(ctx).
			Body(openapi.SecretSvcSaveSecretsRequest{
				Secrets: []openapi.SecretSvcSecretInput{
					{
						Id:    "b",
						Value: openapi.PtrString("value"),
					}},
			}).
			Execute()
		require.NoError(t, err)

		readRsp, _, err := adminClient.SecretSvcAPI.ListSecrets(ctx).
			Body(openapi.SecretSvcListSecretsRequest{
				Ids: []string{"a", "b"},
			}).
			Execute()

		require.NoError(t, err)
		require.Equal(t, 2, len(readRsp.Secrets))
		require.Equal(t, "value", readRsp.Secrets[0].Value)
	})

}

func TestSecretService_BatchCreate(t *testing.T) {
	server, err := test.StartService(test.Options{Test: true})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)
	manyClients, _, err := test.MakeClients(clientFactory, sdk.DefaultTestAppHost, 1)
	require.NoError(t, err)

	cl := manyClients[0]
	ctx := context.Background()
	prefix := "test-user-slug-0/"

	// Execute batch save
	_, _, err = cl.SecretSvcAPI.SaveSecrets(ctx).
		Body(openapi.SecretSvcSaveSecretsRequest{
			Secrets: []openapi.SecretSvcSecretInput{
				{Id: prefix + "b1", Value: openapi.PtrString("v1")},
				{Id: prefix + "b2", Value: openapi.PtrString("v2")},
			},
		}).
		Execute()
	require.NoError(t, err)

	// Verify both exist
	readRsp, _, err := cl.SecretSvcAPI.ListSecrets(ctx).
		Body(openapi.SecretSvcListSecretsRequest{
			Ids: []string{prefix + "b1", prefix + "b2"},
		}).
		Execute()

	require.NoError(t, err)
	require.Equal(t, 2, len(readRsp.Secrets))
}

func TestSecretService_BatchUpdate(t *testing.T) {
	server, err := test.StartService(test.Options{Test: true})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)
	manyClients, _, err := test.MakeClients(clientFactory, sdk.DefaultTestAppHost, 1)
	require.NoError(t, err)

	cl := manyClients[0]
	ctx := context.Background()
	prefix := "test-user-slug-0/"
	id1, id2 := prefix+"u1", prefix+"u2"

	// 1. Initial State
	_, _, err = cl.SecretSvcAPI.SaveSecrets(ctx).
		Body(openapi.SecretSvcSaveSecretsRequest{
			Secrets: []openapi.SecretSvcSecretInput{
				{Id: id1, Value: openapi.PtrString("old-1")},
				{Id: id2, Value: openapi.PtrString("old-2")},
			},
		}).
		Execute()
	require.NoError(t, err)

	// 2. Batch Update (Same IDs, New Values)
	_, _, err = cl.SecretSvcAPI.SaveSecrets(ctx).
		Body(openapi.SecretSvcSaveSecretsRequest{
			Secrets: []openapi.SecretSvcSecretInput{
				{Id: id1, Value: openapi.PtrString("new-1")},
				{Id: id2, Value: openapi.PtrString("new-2")},
			},
		}).
		Execute()
	require.NoError(t, err)

	// 3. Verification
	readRsp, _, err := cl.SecretSvcAPI.ListSecrets(ctx).
		Body(openapi.SecretSvcListSecretsRequest{
			Ids: []string{id1, id2},
		}).
		Execute()

	require.NoError(t, err)

	results := make(map[string]string)
	for _, s := range readRsp.Secrets {
		results[s.Id] = s.Value
	}

	require.Equal(t, "new-1", results[id1])
	require.Equal(t, "new-2", results[id2])
}

func TestSecretService_GranularWriters_Wildcard(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{Test: true})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)
	ctx := t.Context()

	adminClient, _, err := test.AdminClient(clientFactory, sdk.DefaultTestAppHost)
	require.NoError(t, err)

	manyClients, _, err := test.MakeClients(clientFactory, sdk.DefaultTestAppHost, 2)
	require.NoError(t, err)
	client1 := manyClients[0]
	client2 := manyClients[1]

	t.Run("Admin grants wildcard permit to client1", func(t *testing.T) {
		_, _, err = adminClient.UserSvcAPI.SavePermits(ctx).
			Body(openapi.UserSvcSavePermitsRequest{
				Permits: []openapi.UserSvcPermitInput{
					{
						AppHost:    openapi.PtrString("*"),
						Slugs:      []string{"test-user-slug-0"},
						Permission: "secret-svc:secret:save:random-secret-*",
					},
				},
			}).
			Execute()
		require.NoError(t, err)
	})

	t.Run("Client1 can save matching secret", func(t *testing.T) {
		secretId := "random-secret-1" // Matches "random-secret-*"

		_, _, err = client1.SecretSvcAPI.SaveSecrets(ctx).
			Body(openapi.SecretSvcSaveSecretsRequest{
				Secrets: []openapi.SecretSvcSecretInput{
					{
						Id:    secretId,
						Value: openapi.PtrString("wildcard-value"),
					},
				},
			}).
			Execute()

		require.NoError(t, err, "Should be authorized by hierarchical wildcard logic")

		readRsp, _, err := adminClient.SecretSvcAPI.ListSecrets(ctx).
			Body(openapi.SecretSvcListSecretsRequest{
				Id: openapi.PtrString(secretId),
			}).
			Execute()
		require.NoError(t, err)
		require.Equal(t, 1, len(readRsp.Secrets))
		require.Equal(t, "wildcard-value", readRsp.Secrets[0].Value)
	})

	t.Run("Client1 cannot save non-matching secret", func(t *testing.T) {
		secretId := "random-sec1111ret-1"

		_, _, err = client1.SecretSvcAPI.SaveSecrets(ctx).
			Body(openapi.SecretSvcSaveSecretsRequest{
				Secrets: []openapi.SecretSvcSecretInput{
					{
						Id:    secretId,
						Value: openapi.PtrString("wildcard-value"),
					},
				},
			}).
			Execute()

		require.Error(t, err)
	})

	t.Run("Client2 cannot save non-matching secret", func(t *testing.T) {
		_, _, err = client2.SecretSvcAPI.SaveSecrets(ctx).
			Body(openapi.SecretSvcSaveSecretsRequest{
				Secrets: []openapi.SecretSvcSecretInput{
					{
						Id:    "other-prefix-123",
						Value: openapi.PtrString("forbidden"),
					},
				},
			}).
			Execute()

		require.Error(t, err, "Should fail because it doesn't match the wildcard and has no user prefix")
	})
}
