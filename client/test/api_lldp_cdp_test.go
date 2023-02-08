/*
Meraki Dashboard API

Testing LldpCdpApiService

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

func Test_client_LldpCdpApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test LldpCdpApiService GetDeviceLldpCdp", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var serial string

        resp, httpRes, err := apiClient.LldpCdpApi.GetDeviceLldpCdp(context.Background(), serial).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
