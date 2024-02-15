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

// checks if the GetDeviceApplianceRadioSettings200ResponseFiveGhzSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDeviceApplianceRadioSettings200ResponseFiveGhzSettings{}

// GetDeviceApplianceRadioSettings200ResponseFiveGhzSettings Manual radio settings for 5 GHz
type GetDeviceApplianceRadioSettings200ResponseFiveGhzSettings struct {
	// Manual channel for 5 GHz
	Channel *int32 `json:"channel,omitempty"`
	// Manual channel width for 5 GHz
	ChannelWidth *int32 `json:"channelWidth,omitempty"`
	// Manual target power for 5 GHz
	TargetPower *int32 `json:"targetPower,omitempty"`
}

// NewGetDeviceApplianceRadioSettings200ResponseFiveGhzSettings instantiates a new GetDeviceApplianceRadioSettings200ResponseFiveGhzSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDeviceApplianceRadioSettings200ResponseFiveGhzSettings() *GetDeviceApplianceRadioSettings200ResponseFiveGhzSettings {
	this := GetDeviceApplianceRadioSettings200ResponseFiveGhzSettings{}
	return &this
}

// NewGetDeviceApplianceRadioSettings200ResponseFiveGhzSettingsWithDefaults instantiates a new GetDeviceApplianceRadioSettings200ResponseFiveGhzSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDeviceApplianceRadioSettings200ResponseFiveGhzSettingsWithDefaults() *GetDeviceApplianceRadioSettings200ResponseFiveGhzSettings {
	this := GetDeviceApplianceRadioSettings200ResponseFiveGhzSettings{}
	return &this
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *GetDeviceApplianceRadioSettings200ResponseFiveGhzSettings) GetChannel() int32 {
	if o == nil || IsNil(o.Channel) {
		var ret int32
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceApplianceRadioSettings200ResponseFiveGhzSettings) GetChannelOk() (*int32, bool) {
	if o == nil || IsNil(o.Channel) {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *GetDeviceApplianceRadioSettings200ResponseFiveGhzSettings) HasChannel() bool {
	if o != nil && !IsNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given int32 and assigns it to the Channel field.
func (o *GetDeviceApplianceRadioSettings200ResponseFiveGhzSettings) SetChannel(v int32) {
	o.Channel = &v
}

// GetChannelWidth returns the ChannelWidth field value if set, zero value otherwise.
func (o *GetDeviceApplianceRadioSettings200ResponseFiveGhzSettings) GetChannelWidth() int32 {
	if o == nil || IsNil(o.ChannelWidth) {
		var ret int32
		return ret
	}
	return *o.ChannelWidth
}

// GetChannelWidthOk returns a tuple with the ChannelWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceApplianceRadioSettings200ResponseFiveGhzSettings) GetChannelWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.ChannelWidth) {
		return nil, false
	}
	return o.ChannelWidth, true
}

// HasChannelWidth returns a boolean if a field has been set.
func (o *GetDeviceApplianceRadioSettings200ResponseFiveGhzSettings) HasChannelWidth() bool {
	if o != nil && !IsNil(o.ChannelWidth) {
		return true
	}

	return false
}

// SetChannelWidth gets a reference to the given int32 and assigns it to the ChannelWidth field.
func (o *GetDeviceApplianceRadioSettings200ResponseFiveGhzSettings) SetChannelWidth(v int32) {
	o.ChannelWidth = &v
}

// GetTargetPower returns the TargetPower field value if set, zero value otherwise.
func (o *GetDeviceApplianceRadioSettings200ResponseFiveGhzSettings) GetTargetPower() int32 {
	if o == nil || IsNil(o.TargetPower) {
		var ret int32
		return ret
	}
	return *o.TargetPower
}

// GetTargetPowerOk returns a tuple with the TargetPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceApplianceRadioSettings200ResponseFiveGhzSettings) GetTargetPowerOk() (*int32, bool) {
	if o == nil || IsNil(o.TargetPower) {
		return nil, false
	}
	return o.TargetPower, true
}

// HasTargetPower returns a boolean if a field has been set.
func (o *GetDeviceApplianceRadioSettings200ResponseFiveGhzSettings) HasTargetPower() bool {
	if o != nil && !IsNil(o.TargetPower) {
		return true
	}

	return false
}

// SetTargetPower gets a reference to the given int32 and assigns it to the TargetPower field.
func (o *GetDeviceApplianceRadioSettings200ResponseFiveGhzSettings) SetTargetPower(v int32) {
	o.TargetPower = &v
}

func (o GetDeviceApplianceRadioSettings200ResponseFiveGhzSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDeviceApplianceRadioSettings200ResponseFiveGhzSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if !IsNil(o.ChannelWidth) {
		toSerialize["channelWidth"] = o.ChannelWidth
	}
	if !IsNil(o.TargetPower) {
		toSerialize["targetPower"] = o.TargetPower
	}
	return toSerialize, nil
}

type NullableGetDeviceApplianceRadioSettings200ResponseFiveGhzSettings struct {
	value *GetDeviceApplianceRadioSettings200ResponseFiveGhzSettings
	isSet bool
}

func (v NullableGetDeviceApplianceRadioSettings200ResponseFiveGhzSettings) Get() *GetDeviceApplianceRadioSettings200ResponseFiveGhzSettings {
	return v.value
}

func (v *NullableGetDeviceApplianceRadioSettings200ResponseFiveGhzSettings) Set(val *GetDeviceApplianceRadioSettings200ResponseFiveGhzSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDeviceApplianceRadioSettings200ResponseFiveGhzSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDeviceApplianceRadioSettings200ResponseFiveGhzSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDeviceApplianceRadioSettings200ResponseFiveGhzSettings(val *GetDeviceApplianceRadioSettings200ResponseFiveGhzSettings) *NullableGetDeviceApplianceRadioSettings200ResponseFiveGhzSettings {
	return &NullableGetDeviceApplianceRadioSettings200ResponseFiveGhzSettings{value: val, isSet: true}
}

func (v NullableGetDeviceApplianceRadioSettings200ResponseFiveGhzSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDeviceApplianceRadioSettings200ResponseFiveGhzSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


