/*
Meraki Dashboard API

Testing SmApiService

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

func Test_client_SmApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SmApiService CheckinNetworkSmDevices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SmApi.CheckinNetworkSmDevices(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmApiService CreateNetworkSmBypassActivationLockAttempt", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SmApi.CreateNetworkSmBypassActivationLockAttempt(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmApiService CreateNetworkSmTargetGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SmApi.CreateNetworkSmTargetGroup(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmApiService CreateOrganizationSmAdminsRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.SmApi.CreateOrganizationSmAdminsRole(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmApiService DeleteNetworkSmTargetGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var targetGroupId string

		httpRes, err := apiClient.SmApi.DeleteNetworkSmTargetGroup(context.Background(), networkId, targetGroupId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmApiService DeleteNetworkSmUserAccessDevice", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var userAccessDeviceId string

		httpRes, err := apiClient.SmApi.DeleteNetworkSmUserAccessDevice(context.Background(), networkId, userAccessDeviceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmApiService DeleteOrganizationSmAdminsRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string
		var roleId string

		httpRes, err := apiClient.SmApi.DeleteOrganizationSmAdminsRole(context.Background(), organizationId, roleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmApiService GetNetworkSmBypassActivationLockAttempt", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var attemptId string

		resp, httpRes, err := apiClient.SmApi.GetNetworkSmBypassActivationLockAttempt(context.Background(), networkId, attemptId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmApiService GetNetworkSmDeviceCellularUsageHistory", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var deviceId string

		resp, httpRes, err := apiClient.SmApi.GetNetworkSmDeviceCellularUsageHistory(context.Background(), networkId, deviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmApiService GetNetworkSmDeviceCerts", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var deviceId string

		resp, httpRes, err := apiClient.SmApi.GetNetworkSmDeviceCerts(context.Background(), networkId, deviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmApiService GetNetworkSmDeviceConnectivity", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var deviceId string

		resp, httpRes, err := apiClient.SmApi.GetNetworkSmDeviceConnectivity(context.Background(), networkId, deviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmApiService GetNetworkSmDeviceDesktopLogs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var deviceId string

		resp, httpRes, err := apiClient.SmApi.GetNetworkSmDeviceDesktopLogs(context.Background(), networkId, deviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmApiService GetNetworkSmDeviceDeviceCommandLogs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var deviceId string

		resp, httpRes, err := apiClient.SmApi.GetNetworkSmDeviceDeviceCommandLogs(context.Background(), networkId, deviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmApiService GetNetworkSmDeviceDeviceProfiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var deviceId string

		resp, httpRes, err := apiClient.SmApi.GetNetworkSmDeviceDeviceProfiles(context.Background(), networkId, deviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmApiService GetNetworkSmDeviceNetworkAdapters", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var deviceId string

		resp, httpRes, err := apiClient.SmApi.GetNetworkSmDeviceNetworkAdapters(context.Background(), networkId, deviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmApiService GetNetworkSmDevicePerformanceHistory", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var deviceId string

		resp, httpRes, err := apiClient.SmApi.GetNetworkSmDevicePerformanceHistory(context.Background(), networkId, deviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmApiService GetNetworkSmDeviceRestrictions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var deviceId string

		resp, httpRes, err := apiClient.SmApi.GetNetworkSmDeviceRestrictions(context.Background(), networkId, deviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmApiService GetNetworkSmDeviceSecurityCenters", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var deviceId string

		resp, httpRes, err := apiClient.SmApi.GetNetworkSmDeviceSecurityCenters(context.Background(), networkId, deviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmApiService GetNetworkSmDeviceSoftwares", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var deviceId string

		resp, httpRes, err := apiClient.SmApi.GetNetworkSmDeviceSoftwares(context.Background(), networkId, deviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmApiService GetNetworkSmDeviceWlanLists", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var deviceId string

		resp, httpRes, err := apiClient.SmApi.GetNetworkSmDeviceWlanLists(context.Background(), networkId, deviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmApiService GetNetworkSmDevices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SmApi.GetNetworkSmDevices(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmApiService GetNetworkSmProfiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SmApi.GetNetworkSmProfiles(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmApiService GetNetworkSmTargetGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var targetGroupId string

		resp, httpRes, err := apiClient.SmApi.GetNetworkSmTargetGroup(context.Background(), networkId, targetGroupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmApiService GetNetworkSmTargetGroups", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SmApi.GetNetworkSmTargetGroups(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmApiService GetNetworkSmTrustedAccessConfigs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SmApi.GetNetworkSmTrustedAccessConfigs(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmApiService GetNetworkSmUserAccessDevices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SmApi.GetNetworkSmUserAccessDevices(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmApiService GetNetworkSmUserDeviceProfiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var userId string

		resp, httpRes, err := apiClient.SmApi.GetNetworkSmUserDeviceProfiles(context.Background(), networkId, userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmApiService GetNetworkSmUserSoftwares", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var userId string

		resp, httpRes, err := apiClient.SmApi.GetNetworkSmUserSoftwares(context.Background(), networkId, userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmApiService GetNetworkSmUsers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SmApi.GetNetworkSmUsers(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmApiService GetOrganizationSmAdminsRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string
		var roleId string

		resp, httpRes, err := apiClient.SmApi.GetOrganizationSmAdminsRole(context.Background(), organizationId, roleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmApiService GetOrganizationSmAdminsRoles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.SmApi.GetOrganizationSmAdminsRoles(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmApiService GetOrganizationSmApnsCert", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.SmApi.GetOrganizationSmApnsCert(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmApiService GetOrganizationSmSentryPoliciesAssignmentsByNetwork", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.SmApi.GetOrganizationSmSentryPoliciesAssignmentsByNetwork(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmApiService GetOrganizationSmVppAccount", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string
		var vppAccountId string

		resp, httpRes, err := apiClient.SmApi.GetOrganizationSmVppAccount(context.Background(), organizationId, vppAccountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmApiService GetOrganizationSmVppAccounts", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.SmApi.GetOrganizationSmVppAccounts(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmApiService InstallNetworkSmDeviceApps", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var deviceId string

		httpRes, err := apiClient.SmApi.InstallNetworkSmDeviceApps(context.Background(), networkId, deviceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmApiService LockNetworkSmDevices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SmApi.LockNetworkSmDevices(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmApiService ModifyNetworkSmDevicesTags", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SmApi.ModifyNetworkSmDevicesTags(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmApiService MoveNetworkSmDevices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SmApi.MoveNetworkSmDevices(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmApiService RebootNetworkSmDevices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SmApi.RebootNetworkSmDevices(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmApiService RefreshNetworkSmDeviceDetails", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var deviceId string

		httpRes, err := apiClient.SmApi.RefreshNetworkSmDeviceDetails(context.Background(), networkId, deviceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmApiService ShutdownNetworkSmDevices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SmApi.ShutdownNetworkSmDevices(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmApiService UnenrollNetworkSmDevice", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var deviceId string

		resp, httpRes, err := apiClient.SmApi.UnenrollNetworkSmDevice(context.Background(), networkId, deviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmApiService UninstallNetworkSmDeviceApps", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var deviceId string

		httpRes, err := apiClient.SmApi.UninstallNetworkSmDeviceApps(context.Background(), networkId, deviceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmApiService UpdateNetworkSmDevicesFields", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SmApi.UpdateNetworkSmDevicesFields(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmApiService UpdateNetworkSmTargetGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var targetGroupId string

		resp, httpRes, err := apiClient.SmApi.UpdateNetworkSmTargetGroup(context.Background(), networkId, targetGroupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmApiService UpdateOrganizationSmAdminsRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string
		var roleId string

		resp, httpRes, err := apiClient.SmApi.UpdateOrganizationSmAdminsRole(context.Background(), organizationId, roleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmApiService UpdateOrganizationSmSentryPoliciesAssignments", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.SmApi.UpdateOrganizationSmSentryPoliciesAssignments(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmApiService WipeNetworkSmDevices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SmApi.WipeNetworkSmDevices(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
