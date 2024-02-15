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

// checks if the GetOrganizationCameraBoundariesLinesByDevice200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationCameraBoundariesLinesByDevice200ResponseInner{}

// GetOrganizationCameraBoundariesLinesByDevice200ResponseInner struct for GetOrganizationCameraBoundariesLinesByDevice200ResponseInner
type GetOrganizationCameraBoundariesLinesByDevice200ResponseInner struct {
	// The network id of the camera
	NetworkId *string `json:"networkId,omitempty"`
	// The serial number of the camera
	Serial *string `json:"serial,omitempty"`
	Boundaries *GetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundaries `json:"boundaries,omitempty"`
}

// NewGetOrganizationCameraBoundariesLinesByDevice200ResponseInner instantiates a new GetOrganizationCameraBoundariesLinesByDevice200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationCameraBoundariesLinesByDevice200ResponseInner() *GetOrganizationCameraBoundariesLinesByDevice200ResponseInner {
	this := GetOrganizationCameraBoundariesLinesByDevice200ResponseInner{}
	return &this
}

// NewGetOrganizationCameraBoundariesLinesByDevice200ResponseInnerWithDefaults instantiates a new GetOrganizationCameraBoundariesLinesByDevice200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationCameraBoundariesLinesByDevice200ResponseInnerWithDefaults() *GetOrganizationCameraBoundariesLinesByDevice200ResponseInner {
	this := GetOrganizationCameraBoundariesLinesByDevice200ResponseInner{}
	return &this
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *GetOrganizationCameraBoundariesLinesByDevice200ResponseInner) GetNetworkId() string {
	if o == nil || IsNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationCameraBoundariesLinesByDevice200ResponseInner) GetNetworkIdOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkId) {
		return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *GetOrganizationCameraBoundariesLinesByDevice200ResponseInner) HasNetworkId() bool {
	if o != nil && !IsNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *GetOrganizationCameraBoundariesLinesByDevice200ResponseInner) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *GetOrganizationCameraBoundariesLinesByDevice200ResponseInner) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationCameraBoundariesLinesByDevice200ResponseInner) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *GetOrganizationCameraBoundariesLinesByDevice200ResponseInner) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *GetOrganizationCameraBoundariesLinesByDevice200ResponseInner) SetSerial(v string) {
	o.Serial = &v
}

// GetBoundaries returns the Boundaries field value if set, zero value otherwise.
func (o *GetOrganizationCameraBoundariesLinesByDevice200ResponseInner) GetBoundaries() GetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundaries {
	if o == nil || IsNil(o.Boundaries) {
		var ret GetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundaries
		return ret
	}
	return *o.Boundaries
}

// GetBoundariesOk returns a tuple with the Boundaries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationCameraBoundariesLinesByDevice200ResponseInner) GetBoundariesOk() (*GetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundaries, bool) {
	if o == nil || IsNil(o.Boundaries) {
		return nil, false
	}
	return o.Boundaries, true
}

// HasBoundaries returns a boolean if a field has been set.
func (o *GetOrganizationCameraBoundariesLinesByDevice200ResponseInner) HasBoundaries() bool {
	if o != nil && !IsNil(o.Boundaries) {
		return true
	}

	return false
}

// SetBoundaries gets a reference to the given GetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundaries and assigns it to the Boundaries field.
func (o *GetOrganizationCameraBoundariesLinesByDevice200ResponseInner) SetBoundaries(v GetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundaries) {
	o.Boundaries = &v
}

func (o GetOrganizationCameraBoundariesLinesByDevice200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationCameraBoundariesLinesByDevice200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !IsNil(o.Boundaries) {
		toSerialize["boundaries"] = o.Boundaries
	}
	return toSerialize, nil
}

type NullableGetOrganizationCameraBoundariesLinesByDevice200ResponseInner struct {
	value *GetOrganizationCameraBoundariesLinesByDevice200ResponseInner
	isSet bool
}

func (v NullableGetOrganizationCameraBoundariesLinesByDevice200ResponseInner) Get() *GetOrganizationCameraBoundariesLinesByDevice200ResponseInner {
	return v.value
}

func (v *NullableGetOrganizationCameraBoundariesLinesByDevice200ResponseInner) Set(val *GetOrganizationCameraBoundariesLinesByDevice200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationCameraBoundariesLinesByDevice200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationCameraBoundariesLinesByDevice200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationCameraBoundariesLinesByDevice200ResponseInner(val *GetOrganizationCameraBoundariesLinesByDevice200ResponseInner) *NullableGetOrganizationCameraBoundariesLinesByDevice200ResponseInner {
	return &NullableGetOrganizationCameraBoundariesLinesByDevice200ResponseInner{value: val, isSet: true}
}

func (v NullableGetOrganizationCameraBoundariesLinesByDevice200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationCameraBoundariesLinesByDevice200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


