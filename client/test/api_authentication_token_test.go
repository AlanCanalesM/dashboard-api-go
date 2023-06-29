/*
Meraki Dashboard API

Testing AuthenticationTokenApiService

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

func Test_client_AuthenticationTokenApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AuthenticationTokenApiService CreateDeviceApplianceVmxAuthenticationToken", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serial string

		resp, httpRes, err := apiClient.AuthenticationTokenApi.CreateDeviceApplianceVmxAuthenticationToken(context.Background(), serial).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
