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

// InlineResponse20017TrafficFilters struct for InlineResponse20017TrafficFilters
type InlineResponse20017TrafficFilters struct {
	// Traffic filter type. Must be \"custom\"
	Type string `json:"type"`
	Value InlineResponse20017Value `json:"value"`
}

// NewInlineResponse20017TrafficFilters instantiates a new InlineResponse20017TrafficFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20017TrafficFilters(type_ string, value InlineResponse20017Value) *InlineResponse20017TrafficFilters {
	this := InlineResponse20017TrafficFilters{}
	this.Type = type_
	this.Value = value
	return &this
}

// NewInlineResponse20017TrafficFiltersWithDefaults instantiates a new InlineResponse20017TrafficFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20017TrafficFiltersWithDefaults() *InlineResponse20017TrafficFilters {
	this := InlineResponse20017TrafficFilters{}
	return &this
}

// GetType returns the Type field value
func (o *InlineResponse20017TrafficFilters) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20017TrafficFilters) GetTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InlineResponse20017TrafficFilters) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *InlineResponse20017TrafficFilters) GetValue() InlineResponse20017Value {
	if o == nil {
		var ret InlineResponse20017Value
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20017TrafficFilters) GetValueOk() (*InlineResponse20017Value, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *InlineResponse20017TrafficFilters) SetValue(v InlineResponse20017Value) {
	o.Value = v
}

func (o InlineResponse20017TrafficFilters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20017TrafficFilters struct {
	value *InlineResponse20017TrafficFilters
	isSet bool
}

func (v NullableInlineResponse20017TrafficFilters) Get() *InlineResponse20017TrafficFilters {
	return v.value
}

func (v *NullableInlineResponse20017TrafficFilters) Set(val *InlineResponse20017TrafficFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20017TrafficFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20017TrafficFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20017TrafficFilters(val *InlineResponse20017TrafficFilters) *NullableInlineResponse20017TrafficFilters {
	return &NullableInlineResponse20017TrafficFilters{value: val, isSet: true}
}

func (v NullableInlineResponse20017TrafficFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20017TrafficFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


