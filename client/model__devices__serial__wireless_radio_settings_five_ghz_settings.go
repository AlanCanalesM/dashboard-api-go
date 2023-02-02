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

// DevicesSerialWirelessRadioSettingsFiveGhzSettings Manual radio settings for 5 GHz.
type DevicesSerialWirelessRadioSettingsFiveGhzSettings struct {
	// Sets a manual channel for 5 GHz. Can be '36', '40', '44', '48', '52', '56', '60', '64', '100', '104', '108', '112', '116', '120', '124', '128', '132', '136', '140', '144', '149', '153', '157', '161', '165', '169', '173' or '177' or null for using auto channel.
	Channel *int32 `json:"channel,omitempty"`
	// Sets a manual channel for 5 GHz. Can be '0', '20', '40', '80' or '160' or null for using auto channel width.
	ChannelWidth *int32 `json:"channelWidth,omitempty"`
	// Set a manual target power for 5 GHz. Can be between '8' or '30' or null for using auto power range.
	TargetPower *int32 `json:"targetPower,omitempty"`
}

// NewDevicesSerialWirelessRadioSettingsFiveGhzSettings instantiates a new DevicesSerialWirelessRadioSettingsFiveGhzSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesSerialWirelessRadioSettingsFiveGhzSettings() *DevicesSerialWirelessRadioSettingsFiveGhzSettings {
	this := DevicesSerialWirelessRadioSettingsFiveGhzSettings{}
	return &this
}

// NewDevicesSerialWirelessRadioSettingsFiveGhzSettingsWithDefaults instantiates a new DevicesSerialWirelessRadioSettingsFiveGhzSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesSerialWirelessRadioSettingsFiveGhzSettingsWithDefaults() *DevicesSerialWirelessRadioSettingsFiveGhzSettings {
	this := DevicesSerialWirelessRadioSettingsFiveGhzSettings{}
	return &this
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *DevicesSerialWirelessRadioSettingsFiveGhzSettings) GetChannel() int32 {
	if o == nil || isNil(o.Channel) {
		var ret int32
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialWirelessRadioSettingsFiveGhzSettings) GetChannelOk() (*int32, bool) {
	if o == nil || isNil(o.Channel) {
    return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *DevicesSerialWirelessRadioSettingsFiveGhzSettings) HasChannel() bool {
	if o != nil && !isNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given int32 and assigns it to the Channel field.
func (o *DevicesSerialWirelessRadioSettingsFiveGhzSettings) SetChannel(v int32) {
	o.Channel = &v
}

// GetChannelWidth returns the ChannelWidth field value if set, zero value otherwise.
func (o *DevicesSerialWirelessRadioSettingsFiveGhzSettings) GetChannelWidth() int32 {
	if o == nil || isNil(o.ChannelWidth) {
		var ret int32
		return ret
	}
	return *o.ChannelWidth
}

// GetChannelWidthOk returns a tuple with the ChannelWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialWirelessRadioSettingsFiveGhzSettings) GetChannelWidthOk() (*int32, bool) {
	if o == nil || isNil(o.ChannelWidth) {
    return nil, false
	}
	return o.ChannelWidth, true
}

// HasChannelWidth returns a boolean if a field has been set.
func (o *DevicesSerialWirelessRadioSettingsFiveGhzSettings) HasChannelWidth() bool {
	if o != nil && !isNil(o.ChannelWidth) {
		return true
	}

	return false
}

// SetChannelWidth gets a reference to the given int32 and assigns it to the ChannelWidth field.
func (o *DevicesSerialWirelessRadioSettingsFiveGhzSettings) SetChannelWidth(v int32) {
	o.ChannelWidth = &v
}

// GetTargetPower returns the TargetPower field value if set, zero value otherwise.
func (o *DevicesSerialWirelessRadioSettingsFiveGhzSettings) GetTargetPower() int32 {
	if o == nil || isNil(o.TargetPower) {
		var ret int32
		return ret
	}
	return *o.TargetPower
}

// GetTargetPowerOk returns a tuple with the TargetPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialWirelessRadioSettingsFiveGhzSettings) GetTargetPowerOk() (*int32, bool) {
	if o == nil || isNil(o.TargetPower) {
    return nil, false
	}
	return o.TargetPower, true
}

// HasTargetPower returns a boolean if a field has been set.
func (o *DevicesSerialWirelessRadioSettingsFiveGhzSettings) HasTargetPower() bool {
	if o != nil && !isNil(o.TargetPower) {
		return true
	}

	return false
}

// SetTargetPower gets a reference to the given int32 and assigns it to the TargetPower field.
func (o *DevicesSerialWirelessRadioSettingsFiveGhzSettings) SetTargetPower(v int32) {
	o.TargetPower = &v
}

func (o DevicesSerialWirelessRadioSettingsFiveGhzSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if !isNil(o.ChannelWidth) {
		toSerialize["channelWidth"] = o.ChannelWidth
	}
	if !isNil(o.TargetPower) {
		toSerialize["targetPower"] = o.TargetPower
	}
	return json.Marshal(toSerialize)
}

type NullableDevicesSerialWirelessRadioSettingsFiveGhzSettings struct {
	value *DevicesSerialWirelessRadioSettingsFiveGhzSettings
	isSet bool
}

func (v NullableDevicesSerialWirelessRadioSettingsFiveGhzSettings) Get() *DevicesSerialWirelessRadioSettingsFiveGhzSettings {
	return v.value
}

func (v *NullableDevicesSerialWirelessRadioSettingsFiveGhzSettings) Set(val *DevicesSerialWirelessRadioSettingsFiveGhzSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesSerialWirelessRadioSettingsFiveGhzSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesSerialWirelessRadioSettingsFiveGhzSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesSerialWirelessRadioSettingsFiveGhzSettings(val *DevicesSerialWirelessRadioSettingsFiveGhzSettings) *NullableDevicesSerialWirelessRadioSettingsFiveGhzSettings {
	return &NullableDevicesSerialWirelessRadioSettingsFiveGhzSettings{value: val, isSet: true}
}

func (v NullableDevicesSerialWirelessRadioSettingsFiveGhzSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesSerialWirelessRadioSettingsFiveGhzSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


