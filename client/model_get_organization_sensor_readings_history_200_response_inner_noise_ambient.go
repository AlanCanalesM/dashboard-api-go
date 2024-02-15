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

// checks if the GetOrganizationSensorReadingsHistory200ResponseInnerNoiseAmbient type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationSensorReadingsHistory200ResponseInnerNoiseAmbient{}

// GetOrganizationSensorReadingsHistory200ResponseInnerNoiseAmbient Ambient noise reading.
type GetOrganizationSensorReadingsHistory200ResponseInnerNoiseAmbient struct {
	// Ambient noise reading in adjusted decibels.
	Level *int32 `json:"level,omitempty"`
}

// NewGetOrganizationSensorReadingsHistory200ResponseInnerNoiseAmbient instantiates a new GetOrganizationSensorReadingsHistory200ResponseInnerNoiseAmbient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationSensorReadingsHistory200ResponseInnerNoiseAmbient() *GetOrganizationSensorReadingsHistory200ResponseInnerNoiseAmbient {
	this := GetOrganizationSensorReadingsHistory200ResponseInnerNoiseAmbient{}
	return &this
}

// NewGetOrganizationSensorReadingsHistory200ResponseInnerNoiseAmbientWithDefaults instantiates a new GetOrganizationSensorReadingsHistory200ResponseInnerNoiseAmbient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationSensorReadingsHistory200ResponseInnerNoiseAmbientWithDefaults() *GetOrganizationSensorReadingsHistory200ResponseInnerNoiseAmbient {
	this := GetOrganizationSensorReadingsHistory200ResponseInnerNoiseAmbient{}
	return &this
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerNoiseAmbient) GetLevel() int32 {
	if o == nil || IsNil(o.Level) {
		var ret int32
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerNoiseAmbient) GetLevelOk() (*int32, bool) {
	if o == nil || IsNil(o.Level) {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerNoiseAmbient) HasLevel() bool {
	if o != nil && !IsNil(o.Level) {
		return true
	}

	return false
}

// SetLevel gets a reference to the given int32 and assigns it to the Level field.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerNoiseAmbient) SetLevel(v int32) {
	o.Level = &v
}

func (o GetOrganizationSensorReadingsHistory200ResponseInnerNoiseAmbient) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationSensorReadingsHistory200ResponseInnerNoiseAmbient) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Level) {
		toSerialize["level"] = o.Level
	}
	return toSerialize, nil
}

type NullableGetOrganizationSensorReadingsHistory200ResponseInnerNoiseAmbient struct {
	value *GetOrganizationSensorReadingsHistory200ResponseInnerNoiseAmbient
	isSet bool
}

func (v NullableGetOrganizationSensorReadingsHistory200ResponseInnerNoiseAmbient) Get() *GetOrganizationSensorReadingsHistory200ResponseInnerNoiseAmbient {
	return v.value
}

func (v *NullableGetOrganizationSensorReadingsHistory200ResponseInnerNoiseAmbient) Set(val *GetOrganizationSensorReadingsHistory200ResponseInnerNoiseAmbient) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationSensorReadingsHistory200ResponseInnerNoiseAmbient) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationSensorReadingsHistory200ResponseInnerNoiseAmbient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationSensorReadingsHistory200ResponseInnerNoiseAmbient(val *GetOrganizationSensorReadingsHistory200ResponseInnerNoiseAmbient) *NullableGetOrganizationSensorReadingsHistory200ResponseInnerNoiseAmbient {
	return &NullableGetOrganizationSensorReadingsHistory200ResponseInnerNoiseAmbient{value: val, isSet: true}
}

func (v NullableGetOrganizationSensorReadingsHistory200ResponseInnerNoiseAmbient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationSensorReadingsHistory200ResponseInnerNoiseAmbient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


