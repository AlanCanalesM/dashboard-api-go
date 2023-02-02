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

// InlineObject201 struct for InlineObject201
type InlineObject201 struct {
	// Serials of the devices that should be released
	Serials []string `json:"serials,omitempty"`
}

// NewInlineObject201 instantiates a new InlineObject201 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject201() *InlineObject201 {
	this := InlineObject201{}
	return &this
}

// NewInlineObject201WithDefaults instantiates a new InlineObject201 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject201WithDefaults() *InlineObject201 {
	this := InlineObject201{}
	return &this
}

// GetSerials returns the Serials field value if set, zero value otherwise.
func (o *InlineObject201) GetSerials() []string {
	if o == nil || isNil(o.Serials) {
		var ret []string
		return ret
	}
	return o.Serials
}

// GetSerialsOk returns a tuple with the Serials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject201) GetSerialsOk() ([]string, bool) {
	if o == nil || isNil(o.Serials) {
    return nil, false
	}
	return o.Serials, true
}

// HasSerials returns a boolean if a field has been set.
func (o *InlineObject201) HasSerials() bool {
	if o != nil && !isNil(o.Serials) {
		return true
	}

	return false
}

// SetSerials gets a reference to the given []string and assigns it to the Serials field.
func (o *InlineObject201) SetSerials(v []string) {
	o.Serials = v
}

func (o InlineObject201) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Serials) {
		toSerialize["serials"] = o.Serials
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject201 struct {
	value *InlineObject201
	isSet bool
}

func (v NullableInlineObject201) Get() *InlineObject201 {
	return v.value
}

func (v *NullableInlineObject201) Set(val *InlineObject201) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject201) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject201) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject201(val *InlineObject201) *NullableInlineObject201 {
	return &NullableInlineObject201{value: val, isSet: true}
}

func (v NullableInlineObject201) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject201) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


