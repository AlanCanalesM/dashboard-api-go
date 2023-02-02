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

// InlineObject16 struct for InlineObject16
type InlineObject16 struct {
	// List of switch ports. Example: [1, 2-5, 1_MA-MOD-8X10G_1, 1_MA-MOD-8X10G_2-1_MA-MOD-8X10G_8]
	Ports []string `json:"ports"`
}

// NewInlineObject16 instantiates a new InlineObject16 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject16(ports []string) *InlineObject16 {
	this := InlineObject16{}
	this.Ports = ports
	return &this
}

// NewInlineObject16WithDefaults instantiates a new InlineObject16 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject16WithDefaults() *InlineObject16 {
	this := InlineObject16{}
	return &this
}

// GetPorts returns the Ports field value
func (o *InlineObject16) GetPorts() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value
// and a boolean to check if the value has been set.
func (o *InlineObject16) GetPortsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Ports, true
}

// SetPorts sets field value
func (o *InlineObject16) SetPorts(v []string) {
	o.Ports = v
}

func (o InlineObject16) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ports"] = o.Ports
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject16 struct {
	value *InlineObject16
	isSet bool
}

func (v NullableInlineObject16) Get() *InlineObject16 {
	return v.value
}

func (v *NullableInlineObject16) Set(val *InlineObject16) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject16) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject16) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject16(val *InlineObject16) *NullableInlineObject16 {
	return &NullableInlineObject16{value: val, isSet: true}
}

func (v NullableInlineObject16) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject16) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


