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

// NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsageC An object containing current disk usage details.
type NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsageC struct {
	// The used disk space.
	Used *int32 `json:"used,omitempty"`
	// The available disk space.
	Space *int32 `json:"space,omitempty"`
}

// NewNetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsageC instantiates a new NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsageC object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsageC() *NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsageC {
	this := NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsageC{}
	return &this
}

// NewNetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsageCWithDefaults instantiates a new NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsageC object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsageCWithDefaults() *NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsageC {
	this := NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsageC{}
	return &this
}

// GetUsed returns the Used field value if set, zero value otherwise.
func (o *NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsageC) GetUsed() int32 {
	if o == nil || isNil(o.Used) {
		var ret int32
		return ret
	}
	return *o.Used
}

// GetUsedOk returns a tuple with the Used field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsageC) GetUsedOk() (*int32, bool) {
	if o == nil || isNil(o.Used) {
    return nil, false
	}
	return o.Used, true
}

// HasUsed returns a boolean if a field has been set.
func (o *NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsageC) HasUsed() bool {
	if o != nil && !isNil(o.Used) {
		return true
	}

	return false
}

// SetUsed gets a reference to the given int32 and assigns it to the Used field.
func (o *NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsageC) SetUsed(v int32) {
	o.Used = &v
}

// GetSpace returns the Space field value if set, zero value otherwise.
func (o *NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsageC) GetSpace() int32 {
	if o == nil || isNil(o.Space) {
		var ret int32
		return ret
	}
	return *o.Space
}

// GetSpaceOk returns a tuple with the Space field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsageC) GetSpaceOk() (*int32, bool) {
	if o == nil || isNil(o.Space) {
    return nil, false
	}
	return o.Space, true
}

// HasSpace returns a boolean if a field has been set.
func (o *NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsageC) HasSpace() bool {
	if o != nil && !isNil(o.Space) {
		return true
	}

	return false
}

// SetSpace gets a reference to the given int32 and assigns it to the Space field.
func (o *NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsageC) SetSpace(v int32) {
	o.Space = &v
}

func (o NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsageC) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Used) {
		toSerialize["used"] = o.Used
	}
	if !isNil(o.Space) {
		toSerialize["space"] = o.Space
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsageC struct {
	value *NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsageC
	isSet bool
}

func (v NullableNetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsageC) Get() *NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsageC {
	return v.value
}

func (v *NullableNetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsageC) Set(val *NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsageC) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsageC) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsageC) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsageC(val *NetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsageC) *NullableNetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsageC {
	return &NullableNetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsageC{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsageC) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSmDevicesDeviceIdPerformanceHistoryDiskUsageC) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


