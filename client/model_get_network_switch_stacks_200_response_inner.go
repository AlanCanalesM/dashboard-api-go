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

// checks if the GetNetworkSwitchStacks200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSwitchStacks200ResponseInner{}

// GetNetworkSwitchStacks200ResponseInner struct for GetNetworkSwitchStacks200ResponseInner
type GetNetworkSwitchStacks200ResponseInner struct {
	// Switch stacks id
	Id *string `json:"id,omitempty"`
	// Switch stacks name
	Name *string `json:"name,omitempty"`
	// Serials of the switches in the switch stack
	Serials []string `json:"serials,omitempty"`
}

// NewGetNetworkSwitchStacks200ResponseInner instantiates a new GetNetworkSwitchStacks200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSwitchStacks200ResponseInner() *GetNetworkSwitchStacks200ResponseInner {
	this := GetNetworkSwitchStacks200ResponseInner{}
	return &this
}

// NewGetNetworkSwitchStacks200ResponseInnerWithDefaults instantiates a new GetNetworkSwitchStacks200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSwitchStacks200ResponseInnerWithDefaults() *GetNetworkSwitchStacks200ResponseInner {
	this := GetNetworkSwitchStacks200ResponseInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetNetworkSwitchStacks200ResponseInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchStacks200ResponseInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetNetworkSwitchStacks200ResponseInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetNetworkSwitchStacks200ResponseInner) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetNetworkSwitchStacks200ResponseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchStacks200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetNetworkSwitchStacks200ResponseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetNetworkSwitchStacks200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetSerials returns the Serials field value if set, zero value otherwise.
func (o *GetNetworkSwitchStacks200ResponseInner) GetSerials() []string {
	if o == nil || IsNil(o.Serials) {
		var ret []string
		return ret
	}
	return o.Serials
}

// GetSerialsOk returns a tuple with the Serials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchStacks200ResponseInner) GetSerialsOk() ([]string, bool) {
	if o == nil || IsNil(o.Serials) {
		return nil, false
	}
	return o.Serials, true
}

// HasSerials returns a boolean if a field has been set.
func (o *GetNetworkSwitchStacks200ResponseInner) HasSerials() bool {
	if o != nil && !IsNil(o.Serials) {
		return true
	}

	return false
}

// SetSerials gets a reference to the given []string and assigns it to the Serials field.
func (o *GetNetworkSwitchStacks200ResponseInner) SetSerials(v []string) {
	o.Serials = v
}

func (o GetNetworkSwitchStacks200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSwitchStacks200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Serials) {
		toSerialize["serials"] = o.Serials
	}
	return toSerialize, nil
}

type NullableGetNetworkSwitchStacks200ResponseInner struct {
	value *GetNetworkSwitchStacks200ResponseInner
	isSet bool
}

func (v NullableGetNetworkSwitchStacks200ResponseInner) Get() *GetNetworkSwitchStacks200ResponseInner {
	return v.value
}

func (v *NullableGetNetworkSwitchStacks200ResponseInner) Set(val *GetNetworkSwitchStacks200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSwitchStacks200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSwitchStacks200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSwitchStacks200ResponseInner(val *GetNetworkSwitchStacks200ResponseInner) *NullableGetNetworkSwitchStacks200ResponseInner {
	return &NullableGetNetworkSwitchStacks200ResponseInner{value: val, isSet: true}
}

func (v NullableGetNetworkSwitchStacks200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSwitchStacks200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


