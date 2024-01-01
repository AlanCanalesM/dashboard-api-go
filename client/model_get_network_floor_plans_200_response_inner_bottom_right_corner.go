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

// checks if the GetNetworkFloorPlans200ResponseInnerBottomRightCorner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkFloorPlans200ResponseInnerBottomRightCorner{}

// GetNetworkFloorPlans200ResponseInnerBottomRightCorner The longitude and latitude of the bottom right corner of your floor plan.
type GetNetworkFloorPlans200ResponseInnerBottomRightCorner struct {
	// Latitude
	Lat *float32 `json:"lat,omitempty"`
	// Longitude
	Lng *float32 `json:"lng,omitempty"`
}

// NewGetNetworkFloorPlans200ResponseInnerBottomRightCorner instantiates a new GetNetworkFloorPlans200ResponseInnerBottomRightCorner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkFloorPlans200ResponseInnerBottomRightCorner() *GetNetworkFloorPlans200ResponseInnerBottomRightCorner {
	this := GetNetworkFloorPlans200ResponseInnerBottomRightCorner{}
	return &this
}

// NewGetNetworkFloorPlans200ResponseInnerBottomRightCornerWithDefaults instantiates a new GetNetworkFloorPlans200ResponseInnerBottomRightCorner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkFloorPlans200ResponseInnerBottomRightCornerWithDefaults() *GetNetworkFloorPlans200ResponseInnerBottomRightCorner {
	this := GetNetworkFloorPlans200ResponseInnerBottomRightCorner{}
	return &this
}

// GetLat returns the Lat field value if set, zero value otherwise.
func (o *GetNetworkFloorPlans200ResponseInnerBottomRightCorner) GetLat() float32 {
	if o == nil || IsNil(o.Lat) {
		var ret float32
		return ret
	}
	return *o.Lat
}

// GetLatOk returns a tuple with the Lat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFloorPlans200ResponseInnerBottomRightCorner) GetLatOk() (*float32, bool) {
	if o == nil || IsNil(o.Lat) {
		return nil, false
	}
	return o.Lat, true
}

// HasLat returns a boolean if a field has been set.
func (o *GetNetworkFloorPlans200ResponseInnerBottomRightCorner) HasLat() bool {
	if o != nil && !IsNil(o.Lat) {
		return true
	}

	return false
}

// SetLat gets a reference to the given float32 and assigns it to the Lat field.
func (o *GetNetworkFloorPlans200ResponseInnerBottomRightCorner) SetLat(v float32) {
	o.Lat = &v
}

// GetLng returns the Lng field value if set, zero value otherwise.
func (o *GetNetworkFloorPlans200ResponseInnerBottomRightCorner) GetLng() float32 {
	if o == nil || IsNil(o.Lng) {
		var ret float32
		return ret
	}
	return *o.Lng
}

// GetLngOk returns a tuple with the Lng field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFloorPlans200ResponseInnerBottomRightCorner) GetLngOk() (*float32, bool) {
	if o == nil || IsNil(o.Lng) {
		return nil, false
	}
	return o.Lng, true
}

// HasLng returns a boolean if a field has been set.
func (o *GetNetworkFloorPlans200ResponseInnerBottomRightCorner) HasLng() bool {
	if o != nil && !IsNil(o.Lng) {
		return true
	}

	return false
}

// SetLng gets a reference to the given float32 and assigns it to the Lng field.
func (o *GetNetworkFloorPlans200ResponseInnerBottomRightCorner) SetLng(v float32) {
	o.Lng = &v
}

func (o GetNetworkFloorPlans200ResponseInnerBottomRightCorner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkFloorPlans200ResponseInnerBottomRightCorner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Lat) {
		toSerialize["lat"] = o.Lat
	}
	if !IsNil(o.Lng) {
		toSerialize["lng"] = o.Lng
	}
	return toSerialize, nil
}

type NullableGetNetworkFloorPlans200ResponseInnerBottomRightCorner struct {
	value *GetNetworkFloorPlans200ResponseInnerBottomRightCorner
	isSet bool
}

func (v NullableGetNetworkFloorPlans200ResponseInnerBottomRightCorner) Get() *GetNetworkFloorPlans200ResponseInnerBottomRightCorner {
	return v.value
}

func (v *NullableGetNetworkFloorPlans200ResponseInnerBottomRightCorner) Set(val *GetNetworkFloorPlans200ResponseInnerBottomRightCorner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkFloorPlans200ResponseInnerBottomRightCorner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkFloorPlans200ResponseInnerBottomRightCorner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkFloorPlans200ResponseInnerBottomRightCorner(val *GetNetworkFloorPlans200ResponseInnerBottomRightCorner) *NullableGetNetworkFloorPlans200ResponseInnerBottomRightCorner {
	return &NullableGetNetworkFloorPlans200ResponseInnerBottomRightCorner{value: val, isSet: true}
}

func (v NullableGetNetworkFloorPlans200ResponseInnerBottomRightCorner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkFloorPlans200ResponseInnerBottomRightCorner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


