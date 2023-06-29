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

// checks if the GetNetworkSensorRelationships200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSensorRelationships200ResponseInner{}

// GetNetworkSensorRelationships200ResponseInner struct for GetNetworkSensorRelationships200ResponseInner
type GetNetworkSensorRelationships200ResponseInner struct {
	Device *GetNetworkSensorRelationships200ResponseInnerDevice `json:"device,omitempty"`
	Relationships *GetNetworkSensorRelationships200ResponseInnerRelationships `json:"relationships,omitempty"`
}

// NewGetNetworkSensorRelationships200ResponseInner instantiates a new GetNetworkSensorRelationships200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSensorRelationships200ResponseInner() *GetNetworkSensorRelationships200ResponseInner {
	this := GetNetworkSensorRelationships200ResponseInner{}
	return &this
}

// NewGetNetworkSensorRelationships200ResponseInnerWithDefaults instantiates a new GetNetworkSensorRelationships200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSensorRelationships200ResponseInnerWithDefaults() *GetNetworkSensorRelationships200ResponseInner {
	this := GetNetworkSensorRelationships200ResponseInner{}
	return &this
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *GetNetworkSensorRelationships200ResponseInner) GetDevice() GetNetworkSensorRelationships200ResponseInnerDevice {
	if o == nil || IsNil(o.Device) {
		var ret GetNetworkSensorRelationships200ResponseInnerDevice
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorRelationships200ResponseInner) GetDeviceOk() (*GetNetworkSensorRelationships200ResponseInnerDevice, bool) {
	if o == nil || IsNil(o.Device) {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *GetNetworkSensorRelationships200ResponseInner) HasDevice() bool {
	if o != nil && !IsNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given GetNetworkSensorRelationships200ResponseInnerDevice and assigns it to the Device field.
func (o *GetNetworkSensorRelationships200ResponseInner) SetDevice(v GetNetworkSensorRelationships200ResponseInnerDevice) {
	o.Device = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *GetNetworkSensorRelationships200ResponseInner) GetRelationships() GetNetworkSensorRelationships200ResponseInnerRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret GetNetworkSensorRelationships200ResponseInnerRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorRelationships200ResponseInner) GetRelationshipsOk() (*GetNetworkSensorRelationships200ResponseInnerRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *GetNetworkSensorRelationships200ResponseInner) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given GetNetworkSensorRelationships200ResponseInnerRelationships and assigns it to the Relationships field.
func (o *GetNetworkSensorRelationships200ResponseInner) SetRelationships(v GetNetworkSensorRelationships200ResponseInnerRelationships) {
	o.Relationships = &v
}

func (o GetNetworkSensorRelationships200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSensorRelationships200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullableGetNetworkSensorRelationships200ResponseInner struct {
	value *GetNetworkSensorRelationships200ResponseInner
	isSet bool
}

func (v NullableGetNetworkSensorRelationships200ResponseInner) Get() *GetNetworkSensorRelationships200ResponseInner {
	return v.value
}

func (v *NullableGetNetworkSensorRelationships200ResponseInner) Set(val *GetNetworkSensorRelationships200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSensorRelationships200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSensorRelationships200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSensorRelationships200ResponseInner(val *GetNetworkSensorRelationships200ResponseInner) *NullableGetNetworkSensorRelationships200ResponseInner {
	return &NullableGetNetworkSensorRelationships200ResponseInner{value: val, isSet: true}
}

func (v NullableGetNetworkSensorRelationships200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSensorRelationships200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


