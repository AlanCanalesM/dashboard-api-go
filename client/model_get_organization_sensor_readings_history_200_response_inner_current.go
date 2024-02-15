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

// checks if the GetOrganizationSensorReadingsHistory200ResponseInnerCurrent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationSensorReadingsHistory200ResponseInnerCurrent{}

// GetOrganizationSensorReadingsHistory200ResponseInnerCurrent Reading for the 'current' metric. This will only be present if the 'metric' property equals 'current'.
type GetOrganizationSensorReadingsHistory200ResponseInnerCurrent struct {
	// Electrical current reading in amperes.
	Draw *float32 `json:"draw,omitempty"`
}

// NewGetOrganizationSensorReadingsHistory200ResponseInnerCurrent instantiates a new GetOrganizationSensorReadingsHistory200ResponseInnerCurrent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationSensorReadingsHistory200ResponseInnerCurrent() *GetOrganizationSensorReadingsHistory200ResponseInnerCurrent {
	this := GetOrganizationSensorReadingsHistory200ResponseInnerCurrent{}
	return &this
}

// NewGetOrganizationSensorReadingsHistory200ResponseInnerCurrentWithDefaults instantiates a new GetOrganizationSensorReadingsHistory200ResponseInnerCurrent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationSensorReadingsHistory200ResponseInnerCurrentWithDefaults() *GetOrganizationSensorReadingsHistory200ResponseInnerCurrent {
	this := GetOrganizationSensorReadingsHistory200ResponseInnerCurrent{}
	return &this
}

// GetDraw returns the Draw field value if set, zero value otherwise.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerCurrent) GetDraw() float32 {
	if o == nil || IsNil(o.Draw) {
		var ret float32
		return ret
	}
	return *o.Draw
}

// GetDrawOk returns a tuple with the Draw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerCurrent) GetDrawOk() (*float32, bool) {
	if o == nil || IsNil(o.Draw) {
		return nil, false
	}
	return o.Draw, true
}

// HasDraw returns a boolean if a field has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerCurrent) HasDraw() bool {
	if o != nil && !IsNil(o.Draw) {
		return true
	}

	return false
}

// SetDraw gets a reference to the given float32 and assigns it to the Draw field.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerCurrent) SetDraw(v float32) {
	o.Draw = &v
}

func (o GetOrganizationSensorReadingsHistory200ResponseInnerCurrent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationSensorReadingsHistory200ResponseInnerCurrent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Draw) {
		toSerialize["draw"] = o.Draw
	}
	return toSerialize, nil
}

type NullableGetOrganizationSensorReadingsHistory200ResponseInnerCurrent struct {
	value *GetOrganizationSensorReadingsHistory200ResponseInnerCurrent
	isSet bool
}

func (v NullableGetOrganizationSensorReadingsHistory200ResponseInnerCurrent) Get() *GetOrganizationSensorReadingsHistory200ResponseInnerCurrent {
	return v.value
}

func (v *NullableGetOrganizationSensorReadingsHistory200ResponseInnerCurrent) Set(val *GetOrganizationSensorReadingsHistory200ResponseInnerCurrent) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationSensorReadingsHistory200ResponseInnerCurrent) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationSensorReadingsHistory200ResponseInnerCurrent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationSensorReadingsHistory200ResponseInnerCurrent(val *GetOrganizationSensorReadingsHistory200ResponseInnerCurrent) *NullableGetOrganizationSensorReadingsHistory200ResponseInnerCurrent {
	return &NullableGetOrganizationSensorReadingsHistory200ResponseInnerCurrent{value: val, isSet: true}
}

func (v NullableGetOrganizationSensorReadingsHistory200ResponseInnerCurrent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationSensorReadingsHistory200ResponseInnerCurrent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


