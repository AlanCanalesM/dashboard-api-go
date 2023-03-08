/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 22 February, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: v1.31.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject221 struct for InlineObject221
type InlineObject221 struct {
	// Serial number of the source switch (must be on a network not bound to a template)
	SourceSerial string `json:"sourceSerial"`
	// Array of serial numbers of one or more target switches (must be on a network not bound to a template)
	TargetSerials []string `json:"targetSerials"`
}

// NewInlineObject221 instantiates a new InlineObject221 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject221(sourceSerial string, targetSerials []string) *InlineObject221 {
	this := InlineObject221{}
	this.SourceSerial = sourceSerial
	this.TargetSerials = targetSerials
	return &this
}

// NewInlineObject221WithDefaults instantiates a new InlineObject221 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject221WithDefaults() *InlineObject221 {
	this := InlineObject221{}
	return &this
}

// GetSourceSerial returns the SourceSerial field value
func (o *InlineObject221) GetSourceSerial() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceSerial
}

// GetSourceSerialOk returns a tuple with the SourceSerial field value
// and a boolean to check if the value has been set.
func (o *InlineObject221) GetSourceSerialOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SourceSerial, true
}

// SetSourceSerial sets field value
func (o *InlineObject221) SetSourceSerial(v string) {
	o.SourceSerial = v
}

// GetTargetSerials returns the TargetSerials field value
func (o *InlineObject221) GetTargetSerials() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TargetSerials
}

// GetTargetSerialsOk returns a tuple with the TargetSerials field value
// and a boolean to check if the value has been set.
func (o *InlineObject221) GetTargetSerialsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.TargetSerials, true
}

// SetTargetSerials sets field value
func (o *InlineObject221) SetTargetSerials(v []string) {
	o.TargetSerials = v
}

func (o InlineObject221) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sourceSerial"] = o.SourceSerial
	}
	if true {
		toSerialize["targetSerials"] = o.TargetSerials
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject221 struct {
	value *InlineObject221
	isSet bool
}

func (v NullableInlineObject221) Get() *InlineObject221 {
	return v.value
}

func (v *NullableInlineObject221) Set(val *InlineObject221) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject221) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject221) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject221(val *InlineObject221) *NullableInlineObject221 {
	return &NullableInlineObject221{value: val, isSet: true}
}

func (v NullableInlineObject221) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject221) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


