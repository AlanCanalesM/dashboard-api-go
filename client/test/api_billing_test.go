/*
Meraki Dashboard API

Testing BillingApiService

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

func Test_client_BillingApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test BillingApiService GetNetworkWirelessBilling", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string

        resp, httpRes, err := apiClient.BillingApi.GetNetworkWirelessBilling(context.Background(), networkId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test BillingApiService UpdateNetworkWirelessBilling", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string

        resp, httpRes, err := apiClient.BillingApi.UpdateNetworkWirelessBilling(context.Background(), networkId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
