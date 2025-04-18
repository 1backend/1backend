/*
1Backend

Testing DynamicSvcAPIService

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

func Test_openapi_DynamicSvcAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DynamicSvcAPIService CreateObject", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DynamicSvcAPI.CreateObject(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equals(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DynamicSvcAPIService DeleteObjects", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var objectId string

		resp, httpRes, err := apiClient.DynamicSvcAPI.DeleteObjects(context.Background(), objectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equals(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DynamicSvcAPIService Query", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DynamicSvcAPI.Query(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equals(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DynamicSvcAPIService UpdateObjects", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DynamicSvcAPI.UpdateObjects(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equals(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DynamicSvcAPIService UpsertObject", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var objectId string

		resp, httpRes, err := apiClient.DynamicSvcAPI.UpsertObject(context.Background(), objectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equals(t, 200, httpRes.StatusCode)

	})

}
