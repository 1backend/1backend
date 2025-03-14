/*
OpenOrch

Testing UserSvcAPIService

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

func Test_openapi_UserSvcAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test UserSvcAPIService AddPermissionToRole", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var roleId string
		var permissionId string

		resp, httpRes, err := apiClient.UserSvcAPI.AddPermissionToRole(context.Background(), roleId, permissionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equals(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserSvcAPIService AddUserToOrganization", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.UserSvcAPI.AddUserToOrganization(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equals(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserSvcAPIService ChangePassword", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.UserSvcAPI.ChangePassword(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equals(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserSvcAPIService ChangePasswordAdmin", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.UserSvcAPI.ChangePasswordAdmin(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equals(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserSvcAPIService CreateOrganization", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.UserSvcAPI.CreateOrganization(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equals(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserSvcAPIService CreateRole", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.UserSvcAPI.CreateRole(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equals(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserSvcAPIService CreateUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.UserSvcAPI.CreateUser(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equals(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserSvcAPIService DeleteRole", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var roleId string

		resp, httpRes, err := apiClient.UserSvcAPI.DeleteRole(context.Background(), roleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equals(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserSvcAPIService DeleteUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string

		resp, httpRes, err := apiClient.UserSvcAPI.DeleteUser(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equals(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserSvcAPIService GetPermissionsByRole", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var roleId string

		resp, httpRes, err := apiClient.UserSvcAPI.GetPermissionsByRole(context.Background(), roleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equals(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserSvcAPIService GetPublicKey", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.UserSvcAPI.GetPublicKey(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equals(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserSvcAPIService GetRoles", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.UserSvcAPI.GetRoles(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equals(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserSvcAPIService GetUsers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.UserSvcAPI.GetUsers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equals(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserSvcAPIService IsAuthorized", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var permissionId string

		resp, httpRes, err := apiClient.UserSvcAPI.IsAuthorized(context.Background(), permissionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equals(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserSvcAPIService Login", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.UserSvcAPI.Login(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equals(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserSvcAPIService ReadUserByToken", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.UserSvcAPI.ReadUserByToken(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equals(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserSvcAPIService Register", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.UserSvcAPI.Register(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equals(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserSvcAPIService RemoveUserFromOrganization", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var userId string

		resp, httpRes, err := apiClient.UserSvcAPI.RemoveUserFromOrganization(context.Background(), organizationId, userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equals(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserSvcAPIService SaveUserProfile", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string

		resp, httpRes, err := apiClient.UserSvcAPI.SaveUserProfile(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equals(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserSvcAPIService SetRolePermission", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var roleId string

		resp, httpRes, err := apiClient.UserSvcAPI.SetRolePermission(context.Background(), roleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equals(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserSvcAPIService UpsertPermission", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var permissionId string

		resp, httpRes, err := apiClient.UserSvcAPI.UpsertPermission(context.Background(), permissionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equals(t, 200, httpRes.StatusCode)

	})

}
