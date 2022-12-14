/*
Consolidate Services

Testing ProjectServiceApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package api

import (
	"context"
	openapiclient "github.com/eschercloudai/argocd-client-go/pkg/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_api_ProjectServiceApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ProjectServiceApiService ProjectServiceCreate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ProjectServiceApi.ProjectServiceCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectServiceApiService ProjectServiceCreateToken", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var project string
		var role string

		resp, httpRes, err := apiClient.ProjectServiceApi.ProjectServiceCreateToken(context.Background(), project, role).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectServiceApiService ProjectServiceDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var name string

		resp, httpRes, err := apiClient.ProjectServiceApi.ProjectServiceDelete(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectServiceApiService ProjectServiceDeleteToken", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var project string
		var role string
		var iat string

		resp, httpRes, err := apiClient.ProjectServiceApi.ProjectServiceDeleteToken(context.Background(), project, role, iat).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectServiceApiService ProjectServiceGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var name string

		resp, httpRes, err := apiClient.ProjectServiceApi.ProjectServiceGet(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectServiceApiService ProjectServiceGetDetailedProject", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var name string

		resp, httpRes, err := apiClient.ProjectServiceApi.ProjectServiceGetDetailedProject(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectServiceApiService ProjectServiceGetGlobalProjects", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var name string

		resp, httpRes, err := apiClient.ProjectServiceApi.ProjectServiceGetGlobalProjects(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectServiceApiService ProjectServiceGetSyncWindowsState", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var name string

		resp, httpRes, err := apiClient.ProjectServiceApi.ProjectServiceGetSyncWindowsState(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectServiceApiService ProjectServiceList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ProjectServiceApi.ProjectServiceList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectServiceApiService ProjectServiceListEvents", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var name string

		resp, httpRes, err := apiClient.ProjectServiceApi.ProjectServiceListEvents(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectServiceApiService ProjectServiceUpdate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectMetadataName string

		resp, httpRes, err := apiClient.ProjectServiceApi.ProjectServiceUpdate(context.Background(), projectMetadataName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
