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

// InlineObject39 struct for InlineObject39
type InlineObject39 struct {
	// An array of port forwarding params
	Rules []NetworksNetworkIdApplianceFirewallPortForwardingRulesRules `json:"rules"`
}

// NewInlineObject39 instantiates a new InlineObject39 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject39(rules []NetworksNetworkIdApplianceFirewallPortForwardingRulesRules) *InlineObject39 {
	this := InlineObject39{}
	this.Rules = rules
	return &this
}

// NewInlineObject39WithDefaults instantiates a new InlineObject39 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject39WithDefaults() *InlineObject39 {
	this := InlineObject39{}
	return &this
}

// GetRules returns the Rules field value
func (o *InlineObject39) GetRules() []NetworksNetworkIdApplianceFirewallPortForwardingRulesRules {
	if o == nil {
		var ret []NetworksNetworkIdApplianceFirewallPortForwardingRulesRules
		return ret
	}

	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *InlineObject39) GetRulesOk() ([]NetworksNetworkIdApplianceFirewallPortForwardingRulesRules, bool) {
	if o == nil {
    return nil, false
	}
	return o.Rules, true
}

// SetRules sets field value
func (o *InlineObject39) SetRules(v []NetworksNetworkIdApplianceFirewallPortForwardingRulesRules) {
	o.Rules = v
}

func (o InlineObject39) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["rules"] = o.Rules
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject39 struct {
	value *InlineObject39
	isSet bool
}

func (v NullableInlineObject39) Get() *InlineObject39 {
	return v.value
}

func (v *NullableInlineObject39) Set(val *InlineObject39) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject39) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject39) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject39(val *InlineObject39) *NullableInlineObject39 {
	return &NullableInlineObject39{value: val, isSet: true}
}

func (v NullableInlineObject39) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject39) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


