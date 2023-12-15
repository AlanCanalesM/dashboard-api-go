/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 06 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.40.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the UpdateNetworkApplianceSsidRequestDot11w type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkApplianceSsidRequestDot11w{}

// UpdateNetworkApplianceSsidRequestDot11w The current setting for Protected Management Frames (802.11w).
type UpdateNetworkApplianceSsidRequestDot11w struct {
	// Whether 802.11w is enabled or not.
	Enabled *bool `json:"enabled,omitempty"`
	// (Optional) Whether 802.11w is required or not.
	Required *bool `json:"required,omitempty"`
}

// NewUpdateNetworkApplianceSsidRequestDot11w instantiates a new UpdateNetworkApplianceSsidRequestDot11w object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceSsidRequestDot11w() *UpdateNetworkApplianceSsidRequestDot11w {
	this := UpdateNetworkApplianceSsidRequestDot11w{}
	return &this
}

// NewUpdateNetworkApplianceSsidRequestDot11wWithDefaults instantiates a new UpdateNetworkApplianceSsidRequestDot11w object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceSsidRequestDot11wWithDefaults() *UpdateNetworkApplianceSsidRequestDot11w {
	this := UpdateNetworkApplianceSsidRequestDot11w{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceSsidRequestDot11w) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceSsidRequestDot11w) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceSsidRequestDot11w) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateNetworkApplianceSsidRequestDot11w) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceSsidRequestDot11w) GetRequired() bool {
	if o == nil || IsNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceSsidRequestDot11w) GetRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceSsidRequestDot11w) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *UpdateNetworkApplianceSsidRequestDot11w) SetRequired(v bool) {
	o.Required = &v
}

func (o UpdateNetworkApplianceSsidRequestDot11w) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkApplianceSsidRequestDot11w) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	return toSerialize, nil
}

type NullableUpdateNetworkApplianceSsidRequestDot11w struct {
	value *UpdateNetworkApplianceSsidRequestDot11w
	isSet bool
}

func (v NullableUpdateNetworkApplianceSsidRequestDot11w) Get() *UpdateNetworkApplianceSsidRequestDot11w {
	return v.value
}

func (v *NullableUpdateNetworkApplianceSsidRequestDot11w) Set(val *UpdateNetworkApplianceSsidRequestDot11w) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceSsidRequestDot11w) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceSsidRequestDot11w) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceSsidRequestDot11w(val *UpdateNetworkApplianceSsidRequestDot11w) *NullableUpdateNetworkApplianceSsidRequestDot11w {
	return &NullableUpdateNetworkApplianceSsidRequestDot11w{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceSsidRequestDot11w) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceSsidRequestDot11w) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


