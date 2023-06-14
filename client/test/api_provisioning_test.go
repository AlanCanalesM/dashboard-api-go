/*
Meraki Dashboard API

Testing ProvisioningApiService

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

func Test_client_ProvisioningApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test ProvisioningApiService GetOrganizationDevicesProvisioningStatuses", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var organizationId string

        resp, httpRes, err := apiClient.ProvisioningApi.GetOrganizationDevicesProvisioningStatuses(context.Background(), organizationId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
