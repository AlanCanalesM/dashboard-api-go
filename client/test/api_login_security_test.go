/*
Meraki Dashboard API

Testing LoginSecurityApiService

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

func Test_client_LoginSecurityApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test LoginSecurityApiService GetOrganizationLoginSecurity", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var organizationId string

        resp, httpRes, err := apiClient.LoginSecurityApi.GetOrganizationLoginSecurity(context.Background(), organizationId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test LoginSecurityApiService UpdateOrganizationLoginSecurity", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var organizationId string

        resp, httpRes, err := apiClient.LoginSecurityApi.UpdateOrganizationLoginSecurity(context.Background(), organizationId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
