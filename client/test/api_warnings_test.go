/*
Meraki Dashboard API

Testing WarningsApiService

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

func Test_client_WarningsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test WarningsApiService GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.WarningsApi.GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
