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

// checks if the UpdateDeviceCellularSimsRequestSimFailover type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDeviceCellularSimsRequestSimFailover{}

// UpdateDeviceCellularSimsRequestSimFailover SIM Failover settings.
type UpdateDeviceCellularSimsRequestSimFailover struct {
	// Failover to secondary SIM (optional)
	Enabled *bool `json:"enabled,omitempty"`
	// Failover timeout in seconds (optional)
	Timeout *int32 `json:"timeout,omitempty"`
}

// NewUpdateDeviceCellularSimsRequestSimFailover instantiates a new UpdateDeviceCellularSimsRequestSimFailover object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDeviceCellularSimsRequestSimFailover() *UpdateDeviceCellularSimsRequestSimFailover {
	this := UpdateDeviceCellularSimsRequestSimFailover{}
	return &this
}

// NewUpdateDeviceCellularSimsRequestSimFailoverWithDefaults instantiates a new UpdateDeviceCellularSimsRequestSimFailover object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDeviceCellularSimsRequestSimFailoverWithDefaults() *UpdateDeviceCellularSimsRequestSimFailover {
	this := UpdateDeviceCellularSimsRequestSimFailover{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateDeviceCellularSimsRequestSimFailover) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceCellularSimsRequestSimFailover) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateDeviceCellularSimsRequestSimFailover) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateDeviceCellularSimsRequestSimFailover) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *UpdateDeviceCellularSimsRequestSimFailover) GetTimeout() int32 {
	if o == nil || IsNil(o.Timeout) {
		var ret int32
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceCellularSimsRequestSimFailover) GetTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.Timeout) {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *UpdateDeviceCellularSimsRequestSimFailover) HasTimeout() bool {
	if o != nil && !IsNil(o.Timeout) {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int32 and assigns it to the Timeout field.
func (o *UpdateDeviceCellularSimsRequestSimFailover) SetTimeout(v int32) {
	o.Timeout = &v
}

func (o UpdateDeviceCellularSimsRequestSimFailover) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDeviceCellularSimsRequestSimFailover) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Timeout) {
		toSerialize["timeout"] = o.Timeout
	}
	return toSerialize, nil
}

type NullableUpdateDeviceCellularSimsRequestSimFailover struct {
	value *UpdateDeviceCellularSimsRequestSimFailover
	isSet bool
}

func (v NullableUpdateDeviceCellularSimsRequestSimFailover) Get() *UpdateDeviceCellularSimsRequestSimFailover {
	return v.value
}

func (v *NullableUpdateDeviceCellularSimsRequestSimFailover) Set(val *UpdateDeviceCellularSimsRequestSimFailover) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDeviceCellularSimsRequestSimFailover) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDeviceCellularSimsRequestSimFailover) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDeviceCellularSimsRequestSimFailover(val *UpdateDeviceCellularSimsRequestSimFailover) *NullableUpdateDeviceCellularSimsRequestSimFailover {
	return &NullableUpdateDeviceCellularSimsRequestSimFailover{value: val, isSet: true}
}

func (v NullableUpdateDeviceCellularSimsRequestSimFailover) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDeviceCellularSimsRequestSimFailover) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


