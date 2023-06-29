/*
Meraki Dashboard API

Testing LatencyStatsApiService

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

func Test_client_LatencyStatsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test LatencyStatsApiService GetDeviceWirelessLatencyStats", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serial string

		resp, httpRes, err := apiClient.LatencyStatsApi.GetDeviceWirelessLatencyStats(context.Background(), serial).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LatencyStatsApiService GetNetworkWirelessClientLatencyStats", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var clientId string

		resp, httpRes, err := apiClient.LatencyStatsApi.GetNetworkWirelessClientLatencyStats(context.Background(), networkId, clientId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LatencyStatsApiService GetNetworkWirelessClientsLatencyStats", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.LatencyStatsApi.GetNetworkWirelessClientsLatencyStats(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LatencyStatsApiService GetNetworkWirelessDevicesLatencyStats", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.LatencyStatsApi.GetNetworkWirelessDevicesLatencyStats(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LatencyStatsApiService GetNetworkWirelessLatencyStats", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.LatencyStatsApi.GetNetworkWirelessLatencyStats(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
