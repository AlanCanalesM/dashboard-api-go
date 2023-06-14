/*
Meraki Dashboard API

Testing StaticRoutesApiService

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

func Test_client_StaticRoutesApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test StaticRoutesApiService CreateDeviceSwitchRoutingStaticRoute", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var serial string

        resp, httpRes, err := apiClient.StaticRoutesApi.CreateDeviceSwitchRoutingStaticRoute(context.Background(), serial).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test StaticRoutesApiService CreateNetworkApplianceStaticRoute", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string

        resp, httpRes, err := apiClient.StaticRoutesApi.CreateNetworkApplianceStaticRoute(context.Background(), networkId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test StaticRoutesApiService CreateNetworkSwitchStackRoutingStaticRoute", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string
        var switchStackId string

        resp, httpRes, err := apiClient.StaticRoutesApi.CreateNetworkSwitchStackRoutingStaticRoute(context.Background(), networkId, switchStackId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test StaticRoutesApiService DeleteDeviceSwitchRoutingStaticRoute", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var serial string
        var staticRouteId string

        resp, httpRes, err := apiClient.StaticRoutesApi.DeleteDeviceSwitchRoutingStaticRoute(context.Background(), serial, staticRouteId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test StaticRoutesApiService DeleteNetworkApplianceStaticRoute", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string
        var staticRouteId string

        resp, httpRes, err := apiClient.StaticRoutesApi.DeleteNetworkApplianceStaticRoute(context.Background(), networkId, staticRouteId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test StaticRoutesApiService DeleteNetworkSwitchStackRoutingStaticRoute", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string
        var switchStackId string
        var staticRouteId string

        resp, httpRes, err := apiClient.StaticRoutesApi.DeleteNetworkSwitchStackRoutingStaticRoute(context.Background(), networkId, switchStackId, staticRouteId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test StaticRoutesApiService GetDeviceSwitchRoutingStaticRoute", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var serial string
        var staticRouteId string

        resp, httpRes, err := apiClient.StaticRoutesApi.GetDeviceSwitchRoutingStaticRoute(context.Background(), serial, staticRouteId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test StaticRoutesApiService GetDeviceSwitchRoutingStaticRoutes", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var serial string

        resp, httpRes, err := apiClient.StaticRoutesApi.GetDeviceSwitchRoutingStaticRoutes(context.Background(), serial).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test StaticRoutesApiService GetNetworkApplianceStaticRoute", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string
        var staticRouteId string

        resp, httpRes, err := apiClient.StaticRoutesApi.GetNetworkApplianceStaticRoute(context.Background(), networkId, staticRouteId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test StaticRoutesApiService GetNetworkApplianceStaticRoutes", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string

        resp, httpRes, err := apiClient.StaticRoutesApi.GetNetworkApplianceStaticRoutes(context.Background(), networkId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test StaticRoutesApiService GetNetworkSwitchStackRoutingStaticRoute", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string
        var switchStackId string
        var staticRouteId string

        resp, httpRes, err := apiClient.StaticRoutesApi.GetNetworkSwitchStackRoutingStaticRoute(context.Background(), networkId, switchStackId, staticRouteId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test StaticRoutesApiService GetNetworkSwitchStackRoutingStaticRoutes", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string
        var switchStackId string

        resp, httpRes, err := apiClient.StaticRoutesApi.GetNetworkSwitchStackRoutingStaticRoutes(context.Background(), networkId, switchStackId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test StaticRoutesApiService UpdateDeviceSwitchRoutingStaticRoute", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var serial string
        var staticRouteId string

        resp, httpRes, err := apiClient.StaticRoutesApi.UpdateDeviceSwitchRoutingStaticRoute(context.Background(), serial, staticRouteId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test StaticRoutesApiService UpdateNetworkApplianceStaticRoute", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string
        var staticRouteId string

        resp, httpRes, err := apiClient.StaticRoutesApi.UpdateNetworkApplianceStaticRoute(context.Background(), networkId, staticRouteId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test StaticRoutesApiService UpdateNetworkSwitchStackRoutingStaticRoute", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string
        var switchStackId string
        var staticRouteId string

        resp, httpRes, err := apiClient.StaticRoutesApi.UpdateNetworkSwitchStackRoutingStaticRoute(context.Background(), networkId, switchStackId, staticRouteId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
