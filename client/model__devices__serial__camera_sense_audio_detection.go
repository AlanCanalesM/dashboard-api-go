/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// DevicesSerialCameraSenseAudioDetection The details of the audio detection config.
type DevicesSerialCameraSenseAudioDetection struct {
	// Boolean indicating if audio detection is enabled(true) or disabled(false) on the camera
	Enabled *bool `json:"enabled,omitempty"`
}

// NewDevicesSerialCameraSenseAudioDetection instantiates a new DevicesSerialCameraSenseAudioDetection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesSerialCameraSenseAudioDetection() *DevicesSerialCameraSenseAudioDetection {
	this := DevicesSerialCameraSenseAudioDetection{}
	return &this
}

// NewDevicesSerialCameraSenseAudioDetectionWithDefaults instantiates a new DevicesSerialCameraSenseAudioDetection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesSerialCameraSenseAudioDetectionWithDefaults() *DevicesSerialCameraSenseAudioDetection {
	this := DevicesSerialCameraSenseAudioDetection{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *DevicesSerialCameraSenseAudioDetection) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialCameraSenseAudioDetection) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *DevicesSerialCameraSenseAudioDetection) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *DevicesSerialCameraSenseAudioDetection) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o DevicesSerialCameraSenseAudioDetection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableDevicesSerialCameraSenseAudioDetection struct {
	value *DevicesSerialCameraSenseAudioDetection
	isSet bool
}

func (v NullableDevicesSerialCameraSenseAudioDetection) Get() *DevicesSerialCameraSenseAudioDetection {
	return v.value
}

func (v *NullableDevicesSerialCameraSenseAudioDetection) Set(val *DevicesSerialCameraSenseAudioDetection) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesSerialCameraSenseAudioDetection) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesSerialCameraSenseAudioDetection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesSerialCameraSenseAudioDetection(val *DevicesSerialCameraSenseAudioDetection) *NullableDevicesSerialCameraSenseAudioDetection {
	return &NullableDevicesSerialCameraSenseAudioDetection{value: val, isSet: true}
}

func (v NullableDevicesSerialCameraSenseAudioDetection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesSerialCameraSenseAudioDetection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


