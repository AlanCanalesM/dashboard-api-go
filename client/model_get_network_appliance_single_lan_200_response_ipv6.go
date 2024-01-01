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

// checks if the GetNetworkApplianceSingleLan200ResponseIpv6 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkApplianceSingleLan200ResponseIpv6{}

// GetNetworkApplianceSingleLan200ResponseIpv6 IPv6 configuration on the single LAN
type GetNetworkApplianceSingleLan200ResponseIpv6 struct {
	// Enable IPv6 on single LAN
	Enabled *bool `json:"enabled,omitempty"`
	// Prefix assignments on the single LAN
	PrefixAssignments []GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner `json:"prefixAssignments,omitempty"`
}

// NewGetNetworkApplianceSingleLan200ResponseIpv6 instantiates a new GetNetworkApplianceSingleLan200ResponseIpv6 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkApplianceSingleLan200ResponseIpv6() *GetNetworkApplianceSingleLan200ResponseIpv6 {
	this := GetNetworkApplianceSingleLan200ResponseIpv6{}
	return &this
}

// NewGetNetworkApplianceSingleLan200ResponseIpv6WithDefaults instantiates a new GetNetworkApplianceSingleLan200ResponseIpv6 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkApplianceSingleLan200ResponseIpv6WithDefaults() *GetNetworkApplianceSingleLan200ResponseIpv6 {
	this := GetNetworkApplianceSingleLan200ResponseIpv6{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GetNetworkApplianceSingleLan200ResponseIpv6) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceSingleLan200ResponseIpv6) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *GetNetworkApplianceSingleLan200ResponseIpv6) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GetNetworkApplianceSingleLan200ResponseIpv6) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetPrefixAssignments returns the PrefixAssignments field value if set, zero value otherwise.
func (o *GetNetworkApplianceSingleLan200ResponseIpv6) GetPrefixAssignments() []GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner {
	if o == nil || IsNil(o.PrefixAssignments) {
		var ret []GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner
		return ret
	}
	return o.PrefixAssignments
}

// GetPrefixAssignmentsOk returns a tuple with the PrefixAssignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceSingleLan200ResponseIpv6) GetPrefixAssignmentsOk() ([]GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner, bool) {
	if o == nil || IsNil(o.PrefixAssignments) {
		return nil, false
	}
	return o.PrefixAssignments, true
}

// HasPrefixAssignments returns a boolean if a field has been set.
func (o *GetNetworkApplianceSingleLan200ResponseIpv6) HasPrefixAssignments() bool {
	if o != nil && !IsNil(o.PrefixAssignments) {
		return true
	}

	return false
}

// SetPrefixAssignments gets a reference to the given []GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner and assigns it to the PrefixAssignments field.
func (o *GetNetworkApplianceSingleLan200ResponseIpv6) SetPrefixAssignments(v []GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner) {
	o.PrefixAssignments = v
}

func (o GetNetworkApplianceSingleLan200ResponseIpv6) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkApplianceSingleLan200ResponseIpv6) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.PrefixAssignments) {
		toSerialize["prefixAssignments"] = o.PrefixAssignments
	}
	return toSerialize, nil
}

type NullableGetNetworkApplianceSingleLan200ResponseIpv6 struct {
	value *GetNetworkApplianceSingleLan200ResponseIpv6
	isSet bool
}

func (v NullableGetNetworkApplianceSingleLan200ResponseIpv6) Get() *GetNetworkApplianceSingleLan200ResponseIpv6 {
	return v.value
}

func (v *NullableGetNetworkApplianceSingleLan200ResponseIpv6) Set(val *GetNetworkApplianceSingleLan200ResponseIpv6) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkApplianceSingleLan200ResponseIpv6) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkApplianceSingleLan200ResponseIpv6) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkApplianceSingleLan200ResponseIpv6(val *GetNetworkApplianceSingleLan200ResponseIpv6) *NullableGetNetworkApplianceSingleLan200ResponseIpv6 {
	return &NullableGetNetworkApplianceSingleLan200ResponseIpv6{value: val, isSet: true}
}

func (v NullableGetNetworkApplianceSingleLan200ResponseIpv6) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkApplianceSingleLan200ResponseIpv6) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


