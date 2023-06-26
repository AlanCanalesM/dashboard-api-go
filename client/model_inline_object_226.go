/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject226 struct for InlineObject226
type InlineObject226 struct {
	// Serial number of the source switch (must be on a network not bound to a template)
	SourceSerial string `json:"sourceSerial"`
	// Array of serial numbers of one or more target switches (must be on a network not bound to a template)
	TargetSerials []string `json:"targetSerials"`
}

// NewInlineObject226 instantiates a new InlineObject226 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject226(sourceSerial string, targetSerials []string) *InlineObject226 {
	this := InlineObject226{}
	this.SourceSerial = sourceSerial
	this.TargetSerials = targetSerials
	return &this
}

// NewInlineObject226WithDefaults instantiates a new InlineObject226 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject226WithDefaults() *InlineObject226 {
	this := InlineObject226{}
	return &this
}

// GetSourceSerial returns the SourceSerial field value
func (o *InlineObject226) GetSourceSerial() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceSerial
}

// GetSourceSerialOk returns a tuple with the SourceSerial field value
// and a boolean to check if the value has been set.
func (o *InlineObject226) GetSourceSerialOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SourceSerial, true
}

// SetSourceSerial sets field value
func (o *InlineObject226) SetSourceSerial(v string) {
	o.SourceSerial = v
}

// GetTargetSerials returns the TargetSerials field value
func (o *InlineObject226) GetTargetSerials() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TargetSerials
}

// GetTargetSerialsOk returns a tuple with the TargetSerials field value
// and a boolean to check if the value has been set.
func (o *InlineObject226) GetTargetSerialsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.TargetSerials, true
}

// SetTargetSerials sets field value
func (o *InlineObject226) SetTargetSerials(v []string) {
	o.TargetSerials = v
}

func (o InlineObject226) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sourceSerial"] = o.SourceSerial
	}
	if true {
		toSerialize["targetSerials"] = o.TargetSerials
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject226 struct {
	value *InlineObject226
	isSet bool
}

func (v NullableInlineObject226) Get() *InlineObject226 {
	return v.value
}

func (v *NullableInlineObject226) Set(val *InlineObject226) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject226) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject226) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject226(val *InlineObject226) *NullableInlineObject226 {
	return &NullableInlineObject226{value: val, isSet: true}
}

func (v NullableInlineObject226) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject226) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


