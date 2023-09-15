/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 06 September, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.37.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the UpdateDeviceSwitchWarmSpareRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDeviceSwitchWarmSpareRequest{}

// UpdateDeviceSwitchWarmSpareRequest struct for UpdateDeviceSwitchWarmSpareRequest
type UpdateDeviceSwitchWarmSpareRequest struct {
	// Enable or disable warm spare for a switch
	Enabled bool `json:"enabled"`
	// Serial number of the warm spare switch
	SpareSerial *string `json:"spareSerial,omitempty"`
}

// NewUpdateDeviceSwitchWarmSpareRequest instantiates a new UpdateDeviceSwitchWarmSpareRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDeviceSwitchWarmSpareRequest(enabled bool) *UpdateDeviceSwitchWarmSpareRequest {
	this := UpdateDeviceSwitchWarmSpareRequest{}
	this.Enabled = enabled
	return &this
}

// NewUpdateDeviceSwitchWarmSpareRequestWithDefaults instantiates a new UpdateDeviceSwitchWarmSpareRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDeviceSwitchWarmSpareRequestWithDefaults() *UpdateDeviceSwitchWarmSpareRequest {
	this := UpdateDeviceSwitchWarmSpareRequest{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *UpdateDeviceSwitchWarmSpareRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *UpdateDeviceSwitchWarmSpareRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *UpdateDeviceSwitchWarmSpareRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetSpareSerial returns the SpareSerial field value if set, zero value otherwise.
func (o *UpdateDeviceSwitchWarmSpareRequest) GetSpareSerial() string {
	if o == nil || IsNil(o.SpareSerial) {
		var ret string
		return ret
	}
	return *o.SpareSerial
}

// GetSpareSerialOk returns a tuple with the SpareSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceSwitchWarmSpareRequest) GetSpareSerialOk() (*string, bool) {
	if o == nil || IsNil(o.SpareSerial) {
		return nil, false
	}
	return o.SpareSerial, true
}

// HasSpareSerial returns a boolean if a field has been set.
func (o *UpdateDeviceSwitchWarmSpareRequest) HasSpareSerial() bool {
	if o != nil && !IsNil(o.SpareSerial) {
		return true
	}

	return false
}

// SetSpareSerial gets a reference to the given string and assigns it to the SpareSerial field.
func (o *UpdateDeviceSwitchWarmSpareRequest) SetSpareSerial(v string) {
	o.SpareSerial = &v
}

func (o UpdateDeviceSwitchWarmSpareRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDeviceSwitchWarmSpareRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.SpareSerial) {
		toSerialize["spareSerial"] = o.SpareSerial
	}
	return toSerialize, nil
}

type NullableUpdateDeviceSwitchWarmSpareRequest struct {
	value *UpdateDeviceSwitchWarmSpareRequest
	isSet bool
}

func (v NullableUpdateDeviceSwitchWarmSpareRequest) Get() *UpdateDeviceSwitchWarmSpareRequest {
	return v.value
}

func (v *NullableUpdateDeviceSwitchWarmSpareRequest) Set(val *UpdateDeviceSwitchWarmSpareRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDeviceSwitchWarmSpareRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDeviceSwitchWarmSpareRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDeviceSwitchWarmSpareRequest(val *UpdateDeviceSwitchWarmSpareRequest) *NullableUpdateDeviceSwitchWarmSpareRequest {
	return &NullableUpdateDeviceSwitchWarmSpareRequest{value: val, isSet: true}
}

func (v NullableUpdateDeviceSwitchWarmSpareRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDeviceSwitchWarmSpareRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


