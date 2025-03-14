/*
OpenOrch

Testing ChatSvcAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"testing"

	openapiclient "github.com/openorch/openorch/clients/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_ChatSvcAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ChatSvcAPIService AddMessage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var threadId string

		resp, httpRes, err := apiClient.ChatSvcAPI.AddMessage(context.Background(), threadId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equals(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChatSvcAPIService AddThread", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ChatSvcAPI.AddThread(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equals(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChatSvcAPIService DeleteMessage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var messageId string

		resp, httpRes, err := apiClient.ChatSvcAPI.DeleteMessage(context.Background(), messageId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equals(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChatSvcAPIService DeleteThread", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var threadId string

		resp, httpRes, err := apiClient.ChatSvcAPI.DeleteThread(context.Background(), threadId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equals(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChatSvcAPIService GetMessages", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var threadId string

		resp, httpRes, err := apiClient.ChatSvcAPI.GetMessages(context.Background(), threadId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equals(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChatSvcAPIService GetThread", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var threadId string

		resp, httpRes, err := apiClient.ChatSvcAPI.GetThread(context.Background(), threadId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equals(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChatSvcAPIService GetThreads", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ChatSvcAPI.GetThreads(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equals(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChatSvcAPIService UpdateThread", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var threadId string

		resp, httpRes, err := apiClient.ChatSvcAPI.UpdateThread(context.Background(), threadId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equals(t, 200, httpRes.StatusCode)

	})

}
