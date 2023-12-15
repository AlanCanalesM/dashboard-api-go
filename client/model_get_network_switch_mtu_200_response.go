/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 06 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.40.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the GetNetworkSwitchMtu200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSwitchMtu200Response{}

// GetNetworkSwitchMtu200Response struct for GetNetworkSwitchMtu200Response
type GetNetworkSwitchMtu200Response struct {
	// MTU size for the entire network. Default value is 9578.
	DefaultMtuSize *int32 `json:"defaultMtuSize,omitempty"`
	// Override MTU size for individual switches or switch templates.       An empty array will clear overrides.
	Overrides []GetNetworkSwitchMtu200ResponseOverridesInner `json:"overrides,omitempty"`
}

// NewGetNetworkSwitchMtu200Response instantiates a new GetNetworkSwitchMtu200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSwitchMtu200Response() *GetNetworkSwitchMtu200Response {
	this := GetNetworkSwitchMtu200Response{}
	return &this
}

// NewGetNetworkSwitchMtu200ResponseWithDefaults instantiates a new GetNetworkSwitchMtu200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSwitchMtu200ResponseWithDefaults() *GetNetworkSwitchMtu200Response {
	this := GetNetworkSwitchMtu200Response{}
	return &this
}

// GetDefaultMtuSize returns the DefaultMtuSize field value if set, zero value otherwise.
func (o *GetNetworkSwitchMtu200Response) GetDefaultMtuSize() int32 {
	if o == nil || IsNil(o.DefaultMtuSize) {
		var ret int32
		return ret
	}
	return *o.DefaultMtuSize
}

// GetDefaultMtuSizeOk returns a tuple with the DefaultMtuSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchMtu200Response) GetDefaultMtuSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.DefaultMtuSize) {
		return nil, false
	}
	return o.DefaultMtuSize, true
}

// HasDefaultMtuSize returns a boolean if a field has been set.
func (o *GetNetworkSwitchMtu200Response) HasDefaultMtuSize() bool {
	if o != nil && !IsNil(o.DefaultMtuSize) {
		return true
	}

	return false
}

// SetDefaultMtuSize gets a reference to the given int32 and assigns it to the DefaultMtuSize field.
func (o *GetNetworkSwitchMtu200Response) SetDefaultMtuSize(v int32) {
	o.DefaultMtuSize = &v
}

// GetOverrides returns the Overrides field value if set, zero value otherwise.
func (o *GetNetworkSwitchMtu200Response) GetOverrides() []GetNetworkSwitchMtu200ResponseOverridesInner {
	if o == nil || IsNil(o.Overrides) {
		var ret []GetNetworkSwitchMtu200ResponseOverridesInner
		return ret
	}
	return o.Overrides
}

// GetOverridesOk returns a tuple with the Overrides field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchMtu200Response) GetOverridesOk() ([]GetNetworkSwitchMtu200ResponseOverridesInner, bool) {
	if o == nil || IsNil(o.Overrides) {
		return nil, false
	}
	return o.Overrides, true
}

// HasOverrides returns a boolean if a field has been set.
func (o *GetNetworkSwitchMtu200Response) HasOverrides() bool {
	if o != nil && !IsNil(o.Overrides) {
		return true
	}

	return false
}

// SetOverrides gets a reference to the given []GetNetworkSwitchMtu200ResponseOverridesInner and assigns it to the Overrides field.
func (o *GetNetworkSwitchMtu200Response) SetOverrides(v []GetNetworkSwitchMtu200ResponseOverridesInner) {
	o.Overrides = v
}

func (o GetNetworkSwitchMtu200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSwitchMtu200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DefaultMtuSize) {
		toSerialize["defaultMtuSize"] = o.DefaultMtuSize
	}
	if !IsNil(o.Overrides) {
		toSerialize["overrides"] = o.Overrides
	}
	return toSerialize, nil
}

type NullableGetNetworkSwitchMtu200Response struct {
	value *GetNetworkSwitchMtu200Response
	isSet bool
}

func (v NullableGetNetworkSwitchMtu200Response) Get() *GetNetworkSwitchMtu200Response {
	return v.value
}

func (v *NullableGetNetworkSwitchMtu200Response) Set(val *GetNetworkSwitchMtu200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSwitchMtu200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSwitchMtu200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSwitchMtu200Response(val *GetNetworkSwitchMtu200Response) *NullableGetNetworkSwitchMtu200Response {
	return &NullableGetNetworkSwitchMtu200Response{value: val, isSet: true}
}

func (v NullableGetNetworkSwitchMtu200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSwitchMtu200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


