/*
1Backend

Testing ModelSvcAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"testing"

	openapiclient "github.com/1backend/1backend/clients/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_ModelSvcAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ModelSvcAPIService GetDefaultModelStatus", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ModelSvcAPI.GetDefaultModelStatus(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equals(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ModelSvcAPIService GetModel", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var modelId string

		resp, httpRes, err := apiClient.ModelSvcAPI.GetModel(context.Background(), modelId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equals(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ModelSvcAPIService GetModelStatus", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var modelId string

		resp, httpRes, err := apiClient.ModelSvcAPI.GetModelStatus(context.Background(), modelId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equals(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ModelSvcAPIService ListModels", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ModelSvcAPI.ListModels(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equals(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ModelSvcAPIService MakeDefault", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var modelId string

		resp, httpRes, err := apiClient.ModelSvcAPI.MakeDefault(context.Background(), modelId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equals(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ModelSvcAPIService StartDefaultModel", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ModelSvcAPI.StartDefaultModel(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equals(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ModelSvcAPIService StartModel", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var modelId string

		resp, httpRes, err := apiClient.ModelSvcAPI.StartModel(context.Background(), modelId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equals(t, 200, httpRes.StatusCode)

	})

}
