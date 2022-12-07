/*
Consolidate Services

Testing GPGKeyServiceApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package client

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/eschercloudai/argocd-client-go/client"
)

func Test_client_GPGKeyServiceApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test GPGKeyServiceApiService GPGKeyServiceCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GPGKeyServiceApi.GPGKeyServiceCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GPGKeyServiceApiService GPGKeyServiceDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GPGKeyServiceApi.GPGKeyServiceDelete(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GPGKeyServiceApiService GPGKeyServiceGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var keyID string

		resp, httpRes, err := apiClient.GPGKeyServiceApi.GPGKeyServiceGet(context.Background(), keyID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GPGKeyServiceApiService GPGKeyServiceList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GPGKeyServiceApi.GPGKeyServiceList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
