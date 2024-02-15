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

// checks if the UpdateDeviceWirelessBluetoothSettingsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDeviceWirelessBluetoothSettingsRequest{}

// UpdateDeviceWirelessBluetoothSettingsRequest struct for UpdateDeviceWirelessBluetoothSettingsRequest
type UpdateDeviceWirelessBluetoothSettingsRequest struct {
	// Desired UUID of the beacon. If the value is set to null it will reset to Dashboard's automatically generated value.
	Uuid *string `json:"uuid,omitempty"`
	// Desired major value of the beacon. If the value is set to null it will reset to Dashboard's automatically generated value.
	Major *int32 `json:"major,omitempty"`
	// Desired minor value of the beacon. If the value is set to null it will reset to Dashboard's automatically generated value.
	Minor *int32 `json:"minor,omitempty"`
}

// NewUpdateDeviceWirelessBluetoothSettingsRequest instantiates a new UpdateDeviceWirelessBluetoothSettingsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDeviceWirelessBluetoothSettingsRequest() *UpdateDeviceWirelessBluetoothSettingsRequest {
	this := UpdateDeviceWirelessBluetoothSettingsRequest{}
	return &this
}

// NewUpdateDeviceWirelessBluetoothSettingsRequestWithDefaults instantiates a new UpdateDeviceWirelessBluetoothSettingsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDeviceWirelessBluetoothSettingsRequestWithDefaults() *UpdateDeviceWirelessBluetoothSettingsRequest {
	this := UpdateDeviceWirelessBluetoothSettingsRequest{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *UpdateDeviceWirelessBluetoothSettingsRequest) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceWirelessBluetoothSettingsRequest) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *UpdateDeviceWirelessBluetoothSettingsRequest) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *UpdateDeviceWirelessBluetoothSettingsRequest) SetUuid(v string) {
	o.Uuid = &v
}

// GetMajor returns the Major field value if set, zero value otherwise.
func (o *UpdateDeviceWirelessBluetoothSettingsRequest) GetMajor() int32 {
	if o == nil || IsNil(o.Major) {
		var ret int32
		return ret
	}
	return *o.Major
}

// GetMajorOk returns a tuple with the Major field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceWirelessBluetoothSettingsRequest) GetMajorOk() (*int32, bool) {
	if o == nil || IsNil(o.Major) {
		return nil, false
	}
	return o.Major, true
}

// HasMajor returns a boolean if a field has been set.
func (o *UpdateDeviceWirelessBluetoothSettingsRequest) HasMajor() bool {
	if o != nil && !IsNil(o.Major) {
		return true
	}

	return false
}

// SetMajor gets a reference to the given int32 and assigns it to the Major field.
func (o *UpdateDeviceWirelessBluetoothSettingsRequest) SetMajor(v int32) {
	o.Major = &v
}

// GetMinor returns the Minor field value if set, zero value otherwise.
func (o *UpdateDeviceWirelessBluetoothSettingsRequest) GetMinor() int32 {
	if o == nil || IsNil(o.Minor) {
		var ret int32
		return ret
	}
	return *o.Minor
}

// GetMinorOk returns a tuple with the Minor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceWirelessBluetoothSettingsRequest) GetMinorOk() (*int32, bool) {
	if o == nil || IsNil(o.Minor) {
		return nil, false
	}
	return o.Minor, true
}

// HasMinor returns a boolean if a field has been set.
func (o *UpdateDeviceWirelessBluetoothSettingsRequest) HasMinor() bool {
	if o != nil && !IsNil(o.Minor) {
		return true
	}

	return false
}

// SetMinor gets a reference to the given int32 and assigns it to the Minor field.
func (o *UpdateDeviceWirelessBluetoothSettingsRequest) SetMinor(v int32) {
	o.Minor = &v
}

func (o UpdateDeviceWirelessBluetoothSettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDeviceWirelessBluetoothSettingsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !IsNil(o.Major) {
		toSerialize["major"] = o.Major
	}
	if !IsNil(o.Minor) {
		toSerialize["minor"] = o.Minor
	}
	return toSerialize, nil
}

type NullableUpdateDeviceWirelessBluetoothSettingsRequest struct {
	value *UpdateDeviceWirelessBluetoothSettingsRequest
	isSet bool
}

func (v NullableUpdateDeviceWirelessBluetoothSettingsRequest) Get() *UpdateDeviceWirelessBluetoothSettingsRequest {
	return v.value
}

func (v *NullableUpdateDeviceWirelessBluetoothSettingsRequest) Set(val *UpdateDeviceWirelessBluetoothSettingsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDeviceWirelessBluetoothSettingsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDeviceWirelessBluetoothSettingsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDeviceWirelessBluetoothSettingsRequest(val *UpdateDeviceWirelessBluetoothSettingsRequest) *NullableUpdateDeviceWirelessBluetoothSettingsRequest {
	return &NullableUpdateDeviceWirelessBluetoothSettingsRequest{value: val, isSet: true}
}

func (v NullableUpdateDeviceWirelessBluetoothSettingsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDeviceWirelessBluetoothSettingsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


