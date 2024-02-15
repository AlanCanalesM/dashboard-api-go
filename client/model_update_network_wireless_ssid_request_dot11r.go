/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 07 February, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.43.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the UpdateNetworkWirelessSsidRequestDot11r type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkWirelessSsidRequestDot11r{}

// UpdateNetworkWirelessSsidRequestDot11r The current setting for 802.11r
type UpdateNetworkWirelessSsidRequestDot11r struct {
	// Whether 802.11r is enabled or not.
	Enabled *bool `json:"enabled,omitempty"`
	// (Optional) Whether 802.11r is adaptive or not.
	Adaptive *bool `json:"adaptive,omitempty"`
}

// NewUpdateNetworkWirelessSsidRequestDot11r instantiates a new UpdateNetworkWirelessSsidRequestDot11r object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidRequestDot11r() *UpdateNetworkWirelessSsidRequestDot11r {
	this := UpdateNetworkWirelessSsidRequestDot11r{}
	return &this
}

// NewUpdateNetworkWirelessSsidRequestDot11rWithDefaults instantiates a new UpdateNetworkWirelessSsidRequestDot11r object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidRequestDot11rWithDefaults() *UpdateNetworkWirelessSsidRequestDot11r {
	this := UpdateNetworkWirelessSsidRequestDot11r{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidRequestDot11r) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidRequestDot11r) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidRequestDot11r) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateNetworkWirelessSsidRequestDot11r) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetAdaptive returns the Adaptive field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidRequestDot11r) GetAdaptive() bool {
	if o == nil || IsNil(o.Adaptive) {
		var ret bool
		return ret
	}
	return *o.Adaptive
}

// GetAdaptiveOk returns a tuple with the Adaptive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidRequestDot11r) GetAdaptiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Adaptive) {
		return nil, false
	}
	return o.Adaptive, true
}

// HasAdaptive returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidRequestDot11r) HasAdaptive() bool {
	if o != nil && !IsNil(o.Adaptive) {
		return true
	}

	return false
}

// SetAdaptive gets a reference to the given bool and assigns it to the Adaptive field.
func (o *UpdateNetworkWirelessSsidRequestDot11r) SetAdaptive(v bool) {
	o.Adaptive = &v
}

func (o UpdateNetworkWirelessSsidRequestDot11r) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkWirelessSsidRequestDot11r) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Adaptive) {
		toSerialize["adaptive"] = o.Adaptive
	}
	return toSerialize, nil
}

type NullableUpdateNetworkWirelessSsidRequestDot11r struct {
	value *UpdateNetworkWirelessSsidRequestDot11r
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidRequestDot11r) Get() *UpdateNetworkWirelessSsidRequestDot11r {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidRequestDot11r) Set(val *UpdateNetworkWirelessSsidRequestDot11r) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidRequestDot11r) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidRequestDot11r) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidRequestDot11r(val *UpdateNetworkWirelessSsidRequestDot11r) *NullableUpdateNetworkWirelessSsidRequestDot11r {
	return &NullableUpdateNetworkWirelessSsidRequestDot11r{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidRequestDot11r) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidRequestDot11r) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


