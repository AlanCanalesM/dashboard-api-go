/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 May, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.33.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject109 struct for InlineObject109
type InlineObject109 struct {
	// An ordered array of the access control list rules (not including the default rule). An empty array will clear the rules.
	Rules []NetworksNetworkIdSwitchAccessControlListsRules `json:"rules"`
}

// NewInlineObject109 instantiates a new InlineObject109 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject109(rules []NetworksNetworkIdSwitchAccessControlListsRules) *InlineObject109 {
	this := InlineObject109{}
	this.Rules = rules
	return &this
}

// NewInlineObject109WithDefaults instantiates a new InlineObject109 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject109WithDefaults() *InlineObject109 {
	this := InlineObject109{}
	return &this
}

// GetRules returns the Rules field value
func (o *InlineObject109) GetRules() []NetworksNetworkIdSwitchAccessControlListsRules {
	if o == nil {
		var ret []NetworksNetworkIdSwitchAccessControlListsRules
		return ret
	}

	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *InlineObject109) GetRulesOk() ([]NetworksNetworkIdSwitchAccessControlListsRules, bool) {
	if o == nil {
    return nil, false
	}
	return o.Rules, true
}

// SetRules sets field value
func (o *InlineObject109) SetRules(v []NetworksNetworkIdSwitchAccessControlListsRules) {
	o.Rules = v
}

func (o InlineObject109) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["rules"] = o.Rules
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject109 struct {
	value *InlineObject109
	isSet bool
}

func (v NullableInlineObject109) Get() *InlineObject109 {
	return v.value
}

func (v *NullableInlineObject109) Set(val *InlineObject109) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject109) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject109) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject109(val *InlineObject109) *NullableInlineObject109 {
	return &NullableInlineObject109{value: val, isSet: true}
}

func (v NullableInlineObject109) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject109) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


