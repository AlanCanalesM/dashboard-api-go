/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 October, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.26.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// CreateNetworkFloorPlanRequestTopLeftCorner The longitude and latitude of the top left corner of your floor plan.
type CreateNetworkFloorPlanRequestTopLeftCorner struct {
	// Latitude
	Lat *float32 `json:"lat,omitempty"`
	// Longitude
	Lng *float32 `json:"lng,omitempty"`
}

// NewCreateNetworkFloorPlanRequestTopLeftCorner instantiates a new CreateNetworkFloorPlanRequestTopLeftCorner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkFloorPlanRequestTopLeftCorner() *CreateNetworkFloorPlanRequestTopLeftCorner {
	this := CreateNetworkFloorPlanRequestTopLeftCorner{}
	return &this
}

// NewCreateNetworkFloorPlanRequestTopLeftCornerWithDefaults instantiates a new CreateNetworkFloorPlanRequestTopLeftCorner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkFloorPlanRequestTopLeftCornerWithDefaults() *CreateNetworkFloorPlanRequestTopLeftCorner {
	this := CreateNetworkFloorPlanRequestTopLeftCorner{}
	return &this
}

// GetLat returns the Lat field value if set, zero value otherwise.
func (o *CreateNetworkFloorPlanRequestTopLeftCorner) GetLat() float32 {
	if o == nil || isNil(o.Lat) {
		var ret float32
		return ret
	}
	return *o.Lat
}

// GetLatOk returns a tuple with the Lat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkFloorPlanRequestTopLeftCorner) GetLatOk() (*float32, bool) {
	if o == nil || isNil(o.Lat) {
    return nil, false
	}
	return o.Lat, true
}

// HasLat returns a boolean if a field has been set.
func (o *CreateNetworkFloorPlanRequestTopLeftCorner) HasLat() bool {
	if o != nil && !isNil(o.Lat) {
		return true
	}

	return false
}

// SetLat gets a reference to the given float32 and assigns it to the Lat field.
func (o *CreateNetworkFloorPlanRequestTopLeftCorner) SetLat(v float32) {
	o.Lat = &v
}

// GetLng returns the Lng field value if set, zero value otherwise.
func (o *CreateNetworkFloorPlanRequestTopLeftCorner) GetLng() float32 {
	if o == nil || isNil(o.Lng) {
		var ret float32
		return ret
	}
	return *o.Lng
}

// GetLngOk returns a tuple with the Lng field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkFloorPlanRequestTopLeftCorner) GetLngOk() (*float32, bool) {
	if o == nil || isNil(o.Lng) {
    return nil, false
	}
	return o.Lng, true
}

// HasLng returns a boolean if a field has been set.
func (o *CreateNetworkFloorPlanRequestTopLeftCorner) HasLng() bool {
	if o != nil && !isNil(o.Lng) {
		return true
	}

	return false
}

// SetLng gets a reference to the given float32 and assigns it to the Lng field.
func (o *CreateNetworkFloorPlanRequestTopLeftCorner) SetLng(v float32) {
	o.Lng = &v
}

func (o CreateNetworkFloorPlanRequestTopLeftCorner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Lat) {
		toSerialize["lat"] = o.Lat
	}
	if !isNil(o.Lng) {
		toSerialize["lng"] = o.Lng
	}
	return json.Marshal(toSerialize)
}

type NullableCreateNetworkFloorPlanRequestTopLeftCorner struct {
	value *CreateNetworkFloorPlanRequestTopLeftCorner
	isSet bool
}

func (v NullableCreateNetworkFloorPlanRequestTopLeftCorner) Get() *CreateNetworkFloorPlanRequestTopLeftCorner {
	return v.value
}

func (v *NullableCreateNetworkFloorPlanRequestTopLeftCorner) Set(val *CreateNetworkFloorPlanRequestTopLeftCorner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkFloorPlanRequestTopLeftCorner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkFloorPlanRequestTopLeftCorner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkFloorPlanRequestTopLeftCorner(val *CreateNetworkFloorPlanRequestTopLeftCorner) *NullableCreateNetworkFloorPlanRequestTopLeftCorner {
	return &NullableCreateNetworkFloorPlanRequestTopLeftCorner{value: val, isSet: true}
}

func (v NullableCreateNetworkFloorPlanRequestTopLeftCorner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkFloorPlanRequestTopLeftCorner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


