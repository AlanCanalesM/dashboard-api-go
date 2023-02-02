/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 February, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: v1.30.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject126 struct for InlineObject126
type InlineObject126 struct {
	// The IP address of the interface to use
	InterfaceIp string `json:"interfaceIp"`
	// 'Any', or the IP address of a multicast group
	MulticastGroup string `json:"multicastGroup"`
}

// NewInlineObject126 instantiates a new InlineObject126 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject126(interfaceIp string, multicastGroup string) *InlineObject126 {
	this := InlineObject126{}
	this.InterfaceIp = interfaceIp
	this.MulticastGroup = multicastGroup
	return &this
}

// NewInlineObject126WithDefaults instantiates a new InlineObject126 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject126WithDefaults() *InlineObject126 {
	this := InlineObject126{}
	return &this
}

// GetInterfaceIp returns the InterfaceIp field value
func (o *InlineObject126) GetInterfaceIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InterfaceIp
}

// GetInterfaceIpOk returns a tuple with the InterfaceIp field value
// and a boolean to check if the value has been set.
func (o *InlineObject126) GetInterfaceIpOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.InterfaceIp, true
}

// SetInterfaceIp sets field value
func (o *InlineObject126) SetInterfaceIp(v string) {
	o.InterfaceIp = v
}

// GetMulticastGroup returns the MulticastGroup field value
func (o *InlineObject126) GetMulticastGroup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MulticastGroup
}

// GetMulticastGroupOk returns a tuple with the MulticastGroup field value
// and a boolean to check if the value has been set.
func (o *InlineObject126) GetMulticastGroupOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MulticastGroup, true
}

// SetMulticastGroup sets field value
func (o *InlineObject126) SetMulticastGroup(v string) {
	o.MulticastGroup = v
}

func (o InlineObject126) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["interfaceIp"] = o.InterfaceIp
	}
	if true {
		toSerialize["multicastGroup"] = o.MulticastGroup
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject126 struct {
	value *InlineObject126
	isSet bool
}

func (v NullableInlineObject126) Get() *InlineObject126 {
	return v.value
}

func (v *NullableInlineObject126) Set(val *InlineObject126) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject126) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject126) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject126(val *InlineObject126) *NullableInlineObject126 {
	return &NullableInlineObject126{value: val, isSet: true}
}

func (v NullableInlineObject126) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject126) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


