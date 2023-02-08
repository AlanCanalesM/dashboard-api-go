/*
Meraki Dashboard API

Testing BluetoothClientsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package client

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_client_BluetoothClientsApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test BluetoothClientsApiService GetNetworkBluetoothClient", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string
        var bluetoothClientId string

        resp, httpRes, err := apiClient.BluetoothClientsApi.GetNetworkBluetoothClient(context.Background(), networkId, bluetoothClientId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test BluetoothClientsApiService GetNetworkBluetoothClients", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string

        resp, httpRes, err := apiClient.BluetoothClientsApi.GetNetworkBluetoothClients(context.Background(), networkId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
