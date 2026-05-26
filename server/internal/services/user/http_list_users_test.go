package userservice_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	openapi "github.com/1backend/1backend/clients/go"
	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/test"
)

func TestListUsers(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		Test: true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)

	manyClients, tokens, err := test.MakeClients(clientFactory, sdk.DefaultTestAppHost, 5)
	require.NoError(t, err)

	userClient := manyClients[0]

	adminClient, _, err := test.AdminClient(clientFactory, sdk.DefaultTestAppHost)
	require.NoError(t, err)

	contactId := "list-users-contact@example.com"
	contactUser, _, err := clientFactory.Client().UserSvcAPI.Register(
		context.Background(),
	).Body(openapi.UserSvcRegisterRequest{
		AppHost:  sdk.DefaultTestAppHost,
		Slug:     "list-users-contact-target",
		Name:     openapi.PtrString("List Users Contact Target"),
		Password: openapi.PtrString("testUserPasswordContact"),
		Contact: &openapi.UserSvcContactInput{
			Id:       contactId,
			Platform: "email",
		},
	}).Execute()
	require.NoError(t, err)
	require.NotNil(t, contactUser.Token)

	secondContactId := "list-users-contact-2@example.com"
	secondContactUser, _, err := clientFactory.Client().UserSvcAPI.Register(
		context.Background(),
	).Body(openapi.UserSvcRegisterRequest{
		AppHost:  sdk.DefaultTestAppHost,
		Slug:     "list-users-contact-target-2",
		Name:     openapi.PtrString("List Users Contact Target 2"),
		Password: openapi.PtrString("testUserPasswordContact2"),
		Contact: &openapi.UserSvcContactInput{
			Id:       secondContactId,
			Platform: "email",
		},
	}).Execute()
	require.NoError(t, err)
	require.NotNil(t, secondContactUser.Token)

	t.Run("users can not list users", func(t *testing.T) {
		_, _, err := userClient.UserSvcAPI.ListUsers(
			context.Background(),
		).Execute()
		require.Error(t, err)
	})

	t.Run("users can read a single user by contact id", func(t *testing.T) {
		rsp, _, err := userClient.UserSvcAPI.ListUsers(
			context.Background(),
		).Body(openapi.UserSvcListUsersRequest{
			ContactId: openapi.PtrString(contactId),
		}).Execute()
		require.NoError(t, err)

		require.Len(t, rsp.Users, 1)
		require.Equal(t, contactUser.Token.UserId, rsp.Users[0].Id)
		require.Contains(t, rsp.Users[0].ContactIds, contactId)
	})

	t.Run("users can read multiple users by contact ids", func(t *testing.T) {
		body, err := json.Marshal(map[string][]string{
			"contactIds": {contactId, secondContactId},
		})
		require.NoError(t, err)

		req, err := http.NewRequestWithContext(
			context.Background(),
			http.MethodPost,
			server.Url+"/user-svc/users",
			bytes.NewReader(body),
		)
		require.NoError(t, err)
		req.Header.Set("Authorization", "Bearer "+tokens[0].Token)
		req.Header.Set("Content-Type", "application/json")

		httpRsp, err := http.DefaultClient.Do(req)
		require.NoError(t, err)
		defer httpRsp.Body.Close()

		responseBody, err := io.ReadAll(httpRsp.Body)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, httpRsp.StatusCode, string(responseBody))

		rsp := openapi.UserSvcListUsersResponse{}
		err = json.Unmarshal(responseBody, &rsp)
		require.NoError(t, err)

		userIds := []string{}
		for _, user := range rsp.Users {
			userIds = append(userIds, user.Id)
		}
		require.ElementsMatch(t, []string{
			contactUser.Token.UserId,
			secondContactUser.Token.UserId,
		}, userIds)
	})

	t.Run("admins can list users", func(t *testing.T) {
		rsp, _, err := adminClient.UserSvcAPI.ListUsers(
			context.Background(),
		).Execute()
		require.NoError(t, err)

		require.NotEmpty(t, rsp.Users)
		require.True(t, len(rsp.Users) > 6, rsp)
	})

	t.Run("limit", func(t *testing.T) {
		rsp, _, err := adminClient.UserSvcAPI.ListUsers(
			context.Background(),
		).Body(openapi.UserSvcListUsersRequest{
			Limit: openapi.PtrInt32(3),
		}).Execute()
		require.NoError(t, err)

		require.NotEmpty(t, rsp.Users)
		require.True(t, len(rsp.Users) == 3, rsp)
	})

	t.Run("by ids", func(t *testing.T) {
		rsp, _, err := adminClient.UserSvcAPI.ListUsers(
			context.Background(),
		).Body(openapi.UserSvcListUsersRequest{
			Ids: []string{
				tokens[0].UserId,
				tokens[1].UserId,
				tokens[2].UserId,
			},
		}).Execute()
		require.NoError(t, err)

		require.NotEmpty(t, rsp.Users)
		require.True(t, len(rsp.Users) == 3, rsp)

		for _, user := range rsp.Users {
			require.Contains(t, []string{
				tokens[0].UserId,
				tokens[1].UserId,
				tokens[2].UserId,
			}, user.Id)
		}
	})
}
