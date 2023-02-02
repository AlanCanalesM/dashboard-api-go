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

// InlineObject11 struct for InlineObject11
type InlineObject11 struct {
	// An array of port forwarding params
	Rules []DevicesSerialCellularGatewayPortForwardingRulesRules `json:"rules,omitempty"`
}

// NewInlineObject11 instantiates a new InlineObject11 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject11() *InlineObject11 {
	this := InlineObject11{}
	return &this
}

// NewInlineObject11WithDefaults instantiates a new InlineObject11 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject11WithDefaults() *InlineObject11 {
	this := InlineObject11{}
	return &this
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *InlineObject11) GetRules() []DevicesSerialCellularGatewayPortForwardingRulesRules {
	if o == nil || isNil(o.Rules) {
		var ret []DevicesSerialCellularGatewayPortForwardingRulesRules
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject11) GetRulesOk() ([]DevicesSerialCellularGatewayPortForwardingRulesRules, bool) {
	if o == nil || isNil(o.Rules) {
    return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *InlineObject11) HasRules() bool {
	if o != nil && !isNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []DevicesSerialCellularGatewayPortForwardingRulesRules and assigns it to the Rules field.
func (o *InlineObject11) SetRules(v []DevicesSerialCellularGatewayPortForwardingRulesRules) {
	o.Rules = v
}

func (o InlineObject11) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject11 struct {
	value *InlineObject11
	isSet bool
}

func (v NullableInlineObject11) Get() *InlineObject11 {
	return v.value
}

func (v *NullableInlineObject11) Set(val *InlineObject11) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject11) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject11) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject11(val *InlineObject11) *NullableInlineObject11 {
	return &NullableInlineObject11{value: val, isSet: true}
}

func (v NullableInlineObject11) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject11) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


