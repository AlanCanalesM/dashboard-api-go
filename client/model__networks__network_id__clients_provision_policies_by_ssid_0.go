/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 February, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: v1.30.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdClientsProvisionPoliciesBySsid0 The number for the SSID
type NetworksNetworkIdClientsProvisionPoliciesBySsid0 struct {
	// The policy to apply to the specified client. Can be 'Allowed', 'Blocked', 'Normal' or 'Group policy'. Required.
	DevicePolicy string `json:"devicePolicy"`
	// The ID of the desired group policy to apply to the client. Required if 'devicePolicy' is set to \"Group policy\". Otherwise this is ignored.
	GroupPolicyId *string `json:"groupPolicyId,omitempty"`
}

// NewNetworksNetworkIdClientsProvisionPoliciesBySsid0 instantiates a new NetworksNetworkIdClientsProvisionPoliciesBySsid0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdClientsProvisionPoliciesBySsid0(devicePolicy string) *NetworksNetworkIdClientsProvisionPoliciesBySsid0 {
	this := NetworksNetworkIdClientsProvisionPoliciesBySsid0{}
	this.DevicePolicy = devicePolicy
	return &this
}

// NewNetworksNetworkIdClientsProvisionPoliciesBySsid0WithDefaults instantiates a new NetworksNetworkIdClientsProvisionPoliciesBySsid0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdClientsProvisionPoliciesBySsid0WithDefaults() *NetworksNetworkIdClientsProvisionPoliciesBySsid0 {
	this := NetworksNetworkIdClientsProvisionPoliciesBySsid0{}
	return &this
}

// GetDevicePolicy returns the DevicePolicy field value
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid0) GetDevicePolicy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DevicePolicy
}

// GetDevicePolicyOk returns a tuple with the DevicePolicy field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid0) GetDevicePolicyOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DevicePolicy, true
}

// SetDevicePolicy sets field value
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid0) SetDevicePolicy(v string) {
	o.DevicePolicy = v
}

// GetGroupPolicyId returns the GroupPolicyId field value if set, zero value otherwise.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid0) GetGroupPolicyId() string {
	if o == nil || isNil(o.GroupPolicyId) {
		var ret string
		return ret
	}
	return *o.GroupPolicyId
}

// GetGroupPolicyIdOk returns a tuple with the GroupPolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid0) GetGroupPolicyIdOk() (*string, bool) {
	if o == nil || isNil(o.GroupPolicyId) {
    return nil, false
	}
	return o.GroupPolicyId, true
}

// HasGroupPolicyId returns a boolean if a field has been set.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid0) HasGroupPolicyId() bool {
	if o != nil && !isNil(o.GroupPolicyId) {
		return true
	}

	return false
}

// SetGroupPolicyId gets a reference to the given string and assigns it to the GroupPolicyId field.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid0) SetGroupPolicyId(v string) {
	o.GroupPolicyId = &v
}

func (o NetworksNetworkIdClientsProvisionPoliciesBySsid0) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["devicePolicy"] = o.DevicePolicy
	}
	if !isNil(o.GroupPolicyId) {
		toSerialize["groupPolicyId"] = o.GroupPolicyId
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdClientsProvisionPoliciesBySsid0 struct {
	value *NetworksNetworkIdClientsProvisionPoliciesBySsid0
	isSet bool
}

func (v NullableNetworksNetworkIdClientsProvisionPoliciesBySsid0) Get() *NetworksNetworkIdClientsProvisionPoliciesBySsid0 {
	return v.value
}

func (v *NullableNetworksNetworkIdClientsProvisionPoliciesBySsid0) Set(val *NetworksNetworkIdClientsProvisionPoliciesBySsid0) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdClientsProvisionPoliciesBySsid0) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdClientsProvisionPoliciesBySsid0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdClientsProvisionPoliciesBySsid0(val *NetworksNetworkIdClientsProvisionPoliciesBySsid0) *NullableNetworksNetworkIdClientsProvisionPoliciesBySsid0 {
	return &NullableNetworksNetworkIdClientsProvisionPoliciesBySsid0{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdClientsProvisionPoliciesBySsid0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdClientsProvisionPoliciesBySsid0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


