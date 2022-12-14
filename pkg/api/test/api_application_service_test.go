/*
Consolidate Services

Testing ApplicationServiceApiService

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

func Test_api_ApplicationServiceApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ApplicationServiceApiService ApplicationServiceCreate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ApplicationServiceApi.ApplicationServiceCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationServiceApiService ApplicationServiceDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var name string

		resp, httpRes, err := apiClient.ApplicationServiceApi.ApplicationServiceDelete(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationServiceApiService ApplicationServiceDeleteResource", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var name string

		resp, httpRes, err := apiClient.ApplicationServiceApi.ApplicationServiceDeleteResource(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationServiceApiService ApplicationServiceGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var name string

		resp, httpRes, err := apiClient.ApplicationServiceApi.ApplicationServiceGet(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationServiceApiService ApplicationServiceGetApplicationSyncWindows", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var name string

		resp, httpRes, err := apiClient.ApplicationServiceApi.ApplicationServiceGetApplicationSyncWindows(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationServiceApiService ApplicationServiceGetManifests", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var name string

		resp, httpRes, err := apiClient.ApplicationServiceApi.ApplicationServiceGetManifests(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationServiceApiService ApplicationServiceGetManifestsWithFiles", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ApplicationServiceApi.ApplicationServiceGetManifestsWithFiles(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationServiceApiService ApplicationServiceGetResource", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var name string

		resp, httpRes, err := apiClient.ApplicationServiceApi.ApplicationServiceGetResource(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationServiceApiService ApplicationServiceList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ApplicationServiceApi.ApplicationServiceList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationServiceApiService ApplicationServiceListResourceActions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var name string

		resp, httpRes, err := apiClient.ApplicationServiceApi.ApplicationServiceListResourceActions(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationServiceApiService ApplicationServiceListResourceEvents", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var name string

		resp, httpRes, err := apiClient.ApplicationServiceApi.ApplicationServiceListResourceEvents(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationServiceApiService ApplicationServiceManagedResources", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var applicationName string

		resp, httpRes, err := apiClient.ApplicationServiceApi.ApplicationServiceManagedResources(context.Background(), applicationName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationServiceApiService ApplicationServicePatch", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var name string

		resp, httpRes, err := apiClient.ApplicationServiceApi.ApplicationServicePatch(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationServiceApiService ApplicationServicePatchResource", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var name string

		resp, httpRes, err := apiClient.ApplicationServiceApi.ApplicationServicePatchResource(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationServiceApiService ApplicationServicePodLogs", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var name string
		var podName string

		resp, httpRes, err := apiClient.ApplicationServiceApi.ApplicationServicePodLogs(context.Background(), name, podName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationServiceApiService ApplicationServicePodLogs2", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var name string

		resp, httpRes, err := apiClient.ApplicationServiceApi.ApplicationServicePodLogs2(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationServiceApiService ApplicationServiceResourceTree", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var applicationName string

		resp, httpRes, err := apiClient.ApplicationServiceApi.ApplicationServiceResourceTree(context.Background(), applicationName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationServiceApiService ApplicationServiceRevisionMetadata", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var name string
		var revision string

		resp, httpRes, err := apiClient.ApplicationServiceApi.ApplicationServiceRevisionMetadata(context.Background(), name, revision).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationServiceApiService ApplicationServiceRollback", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var name string

		resp, httpRes, err := apiClient.ApplicationServiceApi.ApplicationServiceRollback(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationServiceApiService ApplicationServiceRunResourceAction", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var name string

		resp, httpRes, err := apiClient.ApplicationServiceApi.ApplicationServiceRunResourceAction(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationServiceApiService ApplicationServiceSync", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var name string

		resp, httpRes, err := apiClient.ApplicationServiceApi.ApplicationServiceSync(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationServiceApiService ApplicationServiceTerminateOperation", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var name string

		resp, httpRes, err := apiClient.ApplicationServiceApi.ApplicationServiceTerminateOperation(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationServiceApiService ApplicationServiceUpdate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var applicationMetadataName string

		resp, httpRes, err := apiClient.ApplicationServiceApi.ApplicationServiceUpdate(context.Background(), applicationMetadataName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationServiceApiService ApplicationServiceUpdateSpec", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var name string

		resp, httpRes, err := apiClient.ApplicationServiceApi.ApplicationServiceUpdateSpec(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationServiceApiService ApplicationServiceWatch", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ApplicationServiceApi.ApplicationServiceWatch(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationServiceApiService ApplicationServiceWatchResourceTree", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var applicationName string

		resp, httpRes, err := apiClient.ApplicationServiceApi.ApplicationServiceWatchResourceTree(context.Background(), applicationName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
