/*
Meraki Dashboard API

Testing AdaptivePolicyApiService

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

func Test_client_AdaptivePolicyApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test AdaptivePolicyApiService CreateOrganizationAdaptivePolicyAcl", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var organizationId string

        resp, httpRes, err := apiClient.AdaptivePolicyApi.CreateOrganizationAdaptivePolicyAcl(context.Background(), organizationId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AdaptivePolicyApiService CreateOrganizationAdaptivePolicyGroup", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var organizationId string

        resp, httpRes, err := apiClient.AdaptivePolicyApi.CreateOrganizationAdaptivePolicyGroup(context.Background(), organizationId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AdaptivePolicyApiService CreateOrganizationAdaptivePolicyPolicy", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var organizationId string

        resp, httpRes, err := apiClient.AdaptivePolicyApi.CreateOrganizationAdaptivePolicyPolicy(context.Background(), organizationId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AdaptivePolicyApiService DeleteOrganizationAdaptivePolicyAcl", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var organizationId string
        var aclId string

        resp, httpRes, err := apiClient.AdaptivePolicyApi.DeleteOrganizationAdaptivePolicyAcl(context.Background(), organizationId, aclId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AdaptivePolicyApiService DeleteOrganizationAdaptivePolicyGroup", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var organizationId string
        var id string

        resp, httpRes, err := apiClient.AdaptivePolicyApi.DeleteOrganizationAdaptivePolicyGroup(context.Background(), organizationId, id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AdaptivePolicyApiService DeleteOrganizationAdaptivePolicyPolicy", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var organizationId string
        var id string

        resp, httpRes, err := apiClient.AdaptivePolicyApi.DeleteOrganizationAdaptivePolicyPolicy(context.Background(), organizationId, id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AdaptivePolicyApiService GetOrganizationAdaptivePolicyAcl", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var organizationId string
        var aclId string

        resp, httpRes, err := apiClient.AdaptivePolicyApi.GetOrganizationAdaptivePolicyAcl(context.Background(), organizationId, aclId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AdaptivePolicyApiService GetOrganizationAdaptivePolicyAcls", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var organizationId string

        resp, httpRes, err := apiClient.AdaptivePolicyApi.GetOrganizationAdaptivePolicyAcls(context.Background(), organizationId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AdaptivePolicyApiService GetOrganizationAdaptivePolicyGroup", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var organizationId string
        var id string

        resp, httpRes, err := apiClient.AdaptivePolicyApi.GetOrganizationAdaptivePolicyGroup(context.Background(), organizationId, id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AdaptivePolicyApiService GetOrganizationAdaptivePolicyGroups", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var organizationId string

        resp, httpRes, err := apiClient.AdaptivePolicyApi.GetOrganizationAdaptivePolicyGroups(context.Background(), organizationId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AdaptivePolicyApiService GetOrganizationAdaptivePolicyOverview", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var organizationId string

        resp, httpRes, err := apiClient.AdaptivePolicyApi.GetOrganizationAdaptivePolicyOverview(context.Background(), organizationId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AdaptivePolicyApiService GetOrganizationAdaptivePolicyPolicies", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var organizationId string

        resp, httpRes, err := apiClient.AdaptivePolicyApi.GetOrganizationAdaptivePolicyPolicies(context.Background(), organizationId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AdaptivePolicyApiService GetOrganizationAdaptivePolicyPolicy", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var organizationId string
        var id string

        resp, httpRes, err := apiClient.AdaptivePolicyApi.GetOrganizationAdaptivePolicyPolicy(context.Background(), organizationId, id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AdaptivePolicyApiService GetOrganizationAdaptivePolicySettings", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var organizationId string

        resp, httpRes, err := apiClient.AdaptivePolicyApi.GetOrganizationAdaptivePolicySettings(context.Background(), organizationId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AdaptivePolicyApiService UpdateOrganizationAdaptivePolicyAcl", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var organizationId string
        var aclId string

        resp, httpRes, err := apiClient.AdaptivePolicyApi.UpdateOrganizationAdaptivePolicyAcl(context.Background(), organizationId, aclId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AdaptivePolicyApiService UpdateOrganizationAdaptivePolicyGroup", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var organizationId string
        var id string

        resp, httpRes, err := apiClient.AdaptivePolicyApi.UpdateOrganizationAdaptivePolicyGroup(context.Background(), organizationId, id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AdaptivePolicyApiService UpdateOrganizationAdaptivePolicyPolicy", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var organizationId string
        var id string

        resp, httpRes, err := apiClient.AdaptivePolicyApi.UpdateOrganizationAdaptivePolicyPolicy(context.Background(), organizationId, id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AdaptivePolicyApiService UpdateOrganizationAdaptivePolicySettings", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var organizationId string

        resp, httpRes, err := apiClient.AdaptivePolicyApi.UpdateOrganizationAdaptivePolicySettings(context.Background(), organizationId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
