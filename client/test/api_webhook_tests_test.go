/*
Meraki Dashboard API

Testing WebhookTestsApiService

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

func Test_client_WebhookTestsApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test WebhookTestsApiService CreateNetworkHttpServersWebhookTest", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string

        resp, httpRes, err := apiClient.WebhookTestsApi.CreateNetworkHttpServersWebhookTest(context.Background(), networkId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test WebhookTestsApiService CreateNetworkWebhooksWebhookTest", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string

        resp, httpRes, err := apiClient.WebhookTestsApi.CreateNetworkWebhooksWebhookTest(context.Background(), networkId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test WebhookTestsApiService GetNetworkHttpServersWebhookTest", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string
        var id string

        resp, httpRes, err := apiClient.WebhookTestsApi.GetNetworkHttpServersWebhookTest(context.Background(), networkId, id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test WebhookTestsApiService GetNetworkWebhooksWebhookTest", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string
        var webhookTestId string

        resp, httpRes, err := apiClient.WebhookTestsApi.GetNetworkWebhooksWebhookTest(context.Background(), networkId, webhookTestId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
