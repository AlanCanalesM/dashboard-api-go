/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 18 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.41.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the UpdateNetworkApplianceWarmSpareRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkApplianceWarmSpareRequest{}

// UpdateNetworkApplianceWarmSpareRequest struct for UpdateNetworkApplianceWarmSpareRequest
type UpdateNetworkApplianceWarmSpareRequest struct {
	// Enable warm spare
	Enabled bool `json:"enabled"`
	// Serial number of the warm spare appliance
	SpareSerial *string `json:"spareSerial,omitempty"`
	// Uplink mode, either virtual or public
	UplinkMode *string `json:"uplinkMode,omitempty"`
	// The WAN 1 shared IP
	VirtualIp1 *string `json:"virtualIp1,omitempty"`
	// The WAN 2 shared IP
	VirtualIp2 *string `json:"virtualIp2,omitempty"`
}

// NewUpdateNetworkApplianceWarmSpareRequest instantiates a new UpdateNetworkApplianceWarmSpareRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceWarmSpareRequest(enabled bool) *UpdateNetworkApplianceWarmSpareRequest {
	this := UpdateNetworkApplianceWarmSpareRequest{}
	this.Enabled = enabled
	return &this
}

// NewUpdateNetworkApplianceWarmSpareRequestWithDefaults instantiates a new UpdateNetworkApplianceWarmSpareRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceWarmSpareRequestWithDefaults() *UpdateNetworkApplianceWarmSpareRequest {
	this := UpdateNetworkApplianceWarmSpareRequest{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *UpdateNetworkApplianceWarmSpareRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceWarmSpareRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *UpdateNetworkApplianceWarmSpareRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetSpareSerial returns the SpareSerial field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceWarmSpareRequest) GetSpareSerial() string {
	if o == nil || IsNil(o.SpareSerial) {
		var ret string
		return ret
	}
	return *o.SpareSerial
}

// GetSpareSerialOk returns a tuple with the SpareSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceWarmSpareRequest) GetSpareSerialOk() (*string, bool) {
	if o == nil || IsNil(o.SpareSerial) {
		return nil, false
	}
	return o.SpareSerial, true
}

// HasSpareSerial returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceWarmSpareRequest) HasSpareSerial() bool {
	if o != nil && !IsNil(o.SpareSerial) {
		return true
	}

	return false
}

// SetSpareSerial gets a reference to the given string and assigns it to the SpareSerial field.
func (o *UpdateNetworkApplianceWarmSpareRequest) SetSpareSerial(v string) {
	o.SpareSerial = &v
}

// GetUplinkMode returns the UplinkMode field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceWarmSpareRequest) GetUplinkMode() string {
	if o == nil || IsNil(o.UplinkMode) {
		var ret string
		return ret
	}
	return *o.UplinkMode
}

// GetUplinkModeOk returns a tuple with the UplinkMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceWarmSpareRequest) GetUplinkModeOk() (*string, bool) {
	if o == nil || IsNil(o.UplinkMode) {
		return nil, false
	}
	return o.UplinkMode, true
}

// HasUplinkMode returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceWarmSpareRequest) HasUplinkMode() bool {
	if o != nil && !IsNil(o.UplinkMode) {
		return true
	}

	return false
}

// SetUplinkMode gets a reference to the given string and assigns it to the UplinkMode field.
func (o *UpdateNetworkApplianceWarmSpareRequest) SetUplinkMode(v string) {
	o.UplinkMode = &v
}

// GetVirtualIp1 returns the VirtualIp1 field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceWarmSpareRequest) GetVirtualIp1() string {
	if o == nil || IsNil(o.VirtualIp1) {
		var ret string
		return ret
	}
	return *o.VirtualIp1
}

// GetVirtualIp1Ok returns a tuple with the VirtualIp1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceWarmSpareRequest) GetVirtualIp1Ok() (*string, bool) {
	if o == nil || IsNil(o.VirtualIp1) {
		return nil, false
	}
	return o.VirtualIp1, true
}

// HasVirtualIp1 returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceWarmSpareRequest) HasVirtualIp1() bool {
	if o != nil && !IsNil(o.VirtualIp1) {
		return true
	}

	return false
}

// SetVirtualIp1 gets a reference to the given string and assigns it to the VirtualIp1 field.
func (o *UpdateNetworkApplianceWarmSpareRequest) SetVirtualIp1(v string) {
	o.VirtualIp1 = &v
}

// GetVirtualIp2 returns the VirtualIp2 field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceWarmSpareRequest) GetVirtualIp2() string {
	if o == nil || IsNil(o.VirtualIp2) {
		var ret string
		return ret
	}
	return *o.VirtualIp2
}

// GetVirtualIp2Ok returns a tuple with the VirtualIp2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceWarmSpareRequest) GetVirtualIp2Ok() (*string, bool) {
	if o == nil || IsNil(o.VirtualIp2) {
		return nil, false
	}
	return o.VirtualIp2, true
}

// HasVirtualIp2 returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceWarmSpareRequest) HasVirtualIp2() bool {
	if o != nil && !IsNil(o.VirtualIp2) {
		return true
	}

	return false
}

// SetVirtualIp2 gets a reference to the given string and assigns it to the VirtualIp2 field.
func (o *UpdateNetworkApplianceWarmSpareRequest) SetVirtualIp2(v string) {
	o.VirtualIp2 = &v
}

func (o UpdateNetworkApplianceWarmSpareRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkApplianceWarmSpareRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.SpareSerial) {
		toSerialize["spareSerial"] = o.SpareSerial
	}
	if !IsNil(o.UplinkMode) {
		toSerialize["uplinkMode"] = o.UplinkMode
	}
	if !IsNil(o.VirtualIp1) {
		toSerialize["virtualIp1"] = o.VirtualIp1
	}
	if !IsNil(o.VirtualIp2) {
		toSerialize["virtualIp2"] = o.VirtualIp2
	}
	return toSerialize, nil
}

type NullableUpdateNetworkApplianceWarmSpareRequest struct {
	value *UpdateNetworkApplianceWarmSpareRequest
	isSet bool
}

func (v NullableUpdateNetworkApplianceWarmSpareRequest) Get() *UpdateNetworkApplianceWarmSpareRequest {
	return v.value
}

func (v *NullableUpdateNetworkApplianceWarmSpareRequest) Set(val *UpdateNetworkApplianceWarmSpareRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceWarmSpareRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceWarmSpareRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceWarmSpareRequest(val *UpdateNetworkApplianceWarmSpareRequest) *NullableUpdateNetworkApplianceWarmSpareRequest {
	return &NullableUpdateNetworkApplianceWarmSpareRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceWarmSpareRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceWarmSpareRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


