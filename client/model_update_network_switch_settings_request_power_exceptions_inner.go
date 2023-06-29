/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the UpdateNetworkSwitchSettingsRequestPowerExceptionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkSwitchSettingsRequestPowerExceptionsInner{}

// UpdateNetworkSwitchSettingsRequestPowerExceptionsInner struct for UpdateNetworkSwitchSettingsRequestPowerExceptionsInner
type UpdateNetworkSwitchSettingsRequestPowerExceptionsInner struct {
	// Serial number of the switch
	Serial string `json:"serial"`
	// Per switch exception (combined, redundant, useNetworkSetting)
	PowerType string `json:"powerType"`
}

// NewUpdateNetworkSwitchSettingsRequestPowerExceptionsInner instantiates a new UpdateNetworkSwitchSettingsRequestPowerExceptionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkSwitchSettingsRequestPowerExceptionsInner(serial string, powerType string) *UpdateNetworkSwitchSettingsRequestPowerExceptionsInner {
	this := UpdateNetworkSwitchSettingsRequestPowerExceptionsInner{}
	this.Serial = serial
	this.PowerType = powerType
	return &this
}

// NewUpdateNetworkSwitchSettingsRequestPowerExceptionsInnerWithDefaults instantiates a new UpdateNetworkSwitchSettingsRequestPowerExceptionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkSwitchSettingsRequestPowerExceptionsInnerWithDefaults() *UpdateNetworkSwitchSettingsRequestPowerExceptionsInner {
	this := UpdateNetworkSwitchSettingsRequestPowerExceptionsInner{}
	return &this
}

// GetSerial returns the Serial field value
func (o *UpdateNetworkSwitchSettingsRequestPowerExceptionsInner) GetSerial() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Serial
}

// GetSerialOk returns a tuple with the Serial field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchSettingsRequestPowerExceptionsInner) GetSerialOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Serial, true
}

// SetSerial sets field value
func (o *UpdateNetworkSwitchSettingsRequestPowerExceptionsInner) SetSerial(v string) {
	o.Serial = v
}

// GetPowerType returns the PowerType field value
func (o *UpdateNetworkSwitchSettingsRequestPowerExceptionsInner) GetPowerType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PowerType
}

// GetPowerTypeOk returns a tuple with the PowerType field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchSettingsRequestPowerExceptionsInner) GetPowerTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PowerType, true
}

// SetPowerType sets field value
func (o *UpdateNetworkSwitchSettingsRequestPowerExceptionsInner) SetPowerType(v string) {
	o.PowerType = v
}

func (o UpdateNetworkSwitchSettingsRequestPowerExceptionsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkSwitchSettingsRequestPowerExceptionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["serial"] = o.Serial
	toSerialize["powerType"] = o.PowerType
	return toSerialize, nil
}

type NullableUpdateNetworkSwitchSettingsRequestPowerExceptionsInner struct {
	value *UpdateNetworkSwitchSettingsRequestPowerExceptionsInner
	isSet bool
}

func (v NullableUpdateNetworkSwitchSettingsRequestPowerExceptionsInner) Get() *UpdateNetworkSwitchSettingsRequestPowerExceptionsInner {
	return v.value
}

func (v *NullableUpdateNetworkSwitchSettingsRequestPowerExceptionsInner) Set(val *UpdateNetworkSwitchSettingsRequestPowerExceptionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkSwitchSettingsRequestPowerExceptionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkSwitchSettingsRequestPowerExceptionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkSwitchSettingsRequestPowerExceptionsInner(val *UpdateNetworkSwitchSettingsRequestPowerExceptionsInner) *NullableUpdateNetworkSwitchSettingsRequestPowerExceptionsInner {
	return &NullableUpdateNetworkSwitchSettingsRequestPowerExceptionsInner{value: val, isSet: true}
}

func (v NullableUpdateNetworkSwitchSettingsRequestPowerExceptionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkSwitchSettingsRequestPowerExceptionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


