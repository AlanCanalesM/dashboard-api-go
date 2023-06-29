/*
Meraki Dashboard API

Testing RequestsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package client

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/meraki/dashboard-api-go/client"
)

func Test_client_RequestsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RequestsApiService CreateNetworkPiiRequest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.RequestsApi.CreateNetworkPiiRequest(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RequestsApiService DeleteNetworkPiiRequest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var requestId string

		httpRes, err := apiClient.RequestsApi.DeleteNetworkPiiRequest(context.Background(), networkId, requestId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RequestsApiService GetNetworkPiiRequest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var requestId string

		resp, httpRes, err := apiClient.RequestsApi.GetNetworkPiiRequest(context.Background(), networkId, requestId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RequestsApiService GetNetworkPiiRequests", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.RequestsApi.GetNetworkPiiRequests(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
