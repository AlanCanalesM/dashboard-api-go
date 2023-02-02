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

// InlineObject131 struct for InlineObject131
type InlineObject131 struct {
	// The serial of the switch to be removed
	Serial string `json:"serial"`
}

// NewInlineObject131 instantiates a new InlineObject131 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject131(serial string) *InlineObject131 {
	this := InlineObject131{}
	this.Serial = serial
	return &this
}

// NewInlineObject131WithDefaults instantiates a new InlineObject131 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject131WithDefaults() *InlineObject131 {
	this := InlineObject131{}
	return &this
}

// GetSerial returns the Serial field value
func (o *InlineObject131) GetSerial() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Serial
}

// GetSerialOk returns a tuple with the Serial field value
// and a boolean to check if the value has been set.
func (o *InlineObject131) GetSerialOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Serial, true
}

// SetSerial sets field value
func (o *InlineObject131) SetSerial(v string) {
	o.Serial = v
}

func (o InlineObject131) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["serial"] = o.Serial
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject131 struct {
	value *InlineObject131
	isSet bool
}

func (v NullableInlineObject131) Get() *InlineObject131 {
	return v.value
}

func (v *NullableInlineObject131) Set(val *InlineObject131) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject131) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject131) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject131(val *InlineObject131) *NullableInlineObject131 {
	return &NullableInlineObject131{value: val, isSet: true}
}

func (v NullableInlineObject131) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject131) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


