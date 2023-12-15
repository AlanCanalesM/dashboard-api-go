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

// checks if the GetNetworkVlanProfiles200ResponseInnerVlanNamesInnerAdaptivePolicyGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkVlanProfiles200ResponseInnerVlanNamesInnerAdaptivePolicyGroup{}

// GetNetworkVlanProfiles200ResponseInnerVlanNamesInnerAdaptivePolicyGroup Adaptive Policy Group assigned to Vlan ID
type GetNetworkVlanProfiles200ResponseInnerVlanNamesInnerAdaptivePolicyGroup struct {
	// Adaptive Policy Group ID
	Id *string `json:"id,omitempty"`
	// Adaptive Policy Group name
	Name *string `json:"name,omitempty"`
}

// NewGetNetworkVlanProfiles200ResponseInnerVlanNamesInnerAdaptivePolicyGroup instantiates a new GetNetworkVlanProfiles200ResponseInnerVlanNamesInnerAdaptivePolicyGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkVlanProfiles200ResponseInnerVlanNamesInnerAdaptivePolicyGroup() *GetNetworkVlanProfiles200ResponseInnerVlanNamesInnerAdaptivePolicyGroup {
	this := GetNetworkVlanProfiles200ResponseInnerVlanNamesInnerAdaptivePolicyGroup{}
	return &this
}

// NewGetNetworkVlanProfiles200ResponseInnerVlanNamesInnerAdaptivePolicyGroupWithDefaults instantiates a new GetNetworkVlanProfiles200ResponseInnerVlanNamesInnerAdaptivePolicyGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkVlanProfiles200ResponseInnerVlanNamesInnerAdaptivePolicyGroupWithDefaults() *GetNetworkVlanProfiles200ResponseInnerVlanNamesInnerAdaptivePolicyGroup {
	this := GetNetworkVlanProfiles200ResponseInnerVlanNamesInnerAdaptivePolicyGroup{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetNetworkVlanProfiles200ResponseInnerVlanNamesInnerAdaptivePolicyGroup) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkVlanProfiles200ResponseInnerVlanNamesInnerAdaptivePolicyGroup) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetNetworkVlanProfiles200ResponseInnerVlanNamesInnerAdaptivePolicyGroup) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetNetworkVlanProfiles200ResponseInnerVlanNamesInnerAdaptivePolicyGroup) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetNetworkVlanProfiles200ResponseInnerVlanNamesInnerAdaptivePolicyGroup) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkVlanProfiles200ResponseInnerVlanNamesInnerAdaptivePolicyGroup) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetNetworkVlanProfiles200ResponseInnerVlanNamesInnerAdaptivePolicyGroup) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetNetworkVlanProfiles200ResponseInnerVlanNamesInnerAdaptivePolicyGroup) SetName(v string) {
	o.Name = &v
}

func (o GetNetworkVlanProfiles200ResponseInnerVlanNamesInnerAdaptivePolicyGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkVlanProfiles200ResponseInnerVlanNamesInnerAdaptivePolicyGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableGetNetworkVlanProfiles200ResponseInnerVlanNamesInnerAdaptivePolicyGroup struct {
	value *GetNetworkVlanProfiles200ResponseInnerVlanNamesInnerAdaptivePolicyGroup
	isSet bool
}

func (v NullableGetNetworkVlanProfiles200ResponseInnerVlanNamesInnerAdaptivePolicyGroup) Get() *GetNetworkVlanProfiles200ResponseInnerVlanNamesInnerAdaptivePolicyGroup {
	return v.value
}

func (v *NullableGetNetworkVlanProfiles200ResponseInnerVlanNamesInnerAdaptivePolicyGroup) Set(val *GetNetworkVlanProfiles200ResponseInnerVlanNamesInnerAdaptivePolicyGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkVlanProfiles200ResponseInnerVlanNamesInnerAdaptivePolicyGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkVlanProfiles200ResponseInnerVlanNamesInnerAdaptivePolicyGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkVlanProfiles200ResponseInnerVlanNamesInnerAdaptivePolicyGroup(val *GetNetworkVlanProfiles200ResponseInnerVlanNamesInnerAdaptivePolicyGroup) *NullableGetNetworkVlanProfiles200ResponseInnerVlanNamesInnerAdaptivePolicyGroup {
	return &NullableGetNetworkVlanProfiles200ResponseInnerVlanNamesInnerAdaptivePolicyGroup{value: val, isSet: true}
}

func (v NullableGetNetworkVlanProfiles200ResponseInnerVlanNamesInnerAdaptivePolicyGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkVlanProfiles200ResponseInnerVlanNamesInnerAdaptivePolicyGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


