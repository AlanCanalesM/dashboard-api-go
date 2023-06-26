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

// NetworksNetworkIdWirelessSsidsNumberFirewallL7FirewallRulesRules struct for NetworksNetworkIdWirelessSsidsNumberFirewallL7FirewallRulesRules
type NetworksNetworkIdWirelessSsidsNumberFirewallL7FirewallRulesRules struct {
	// 'Deny' traffic specified by this rule
	Policy *string `json:"policy,omitempty"`
	// Type of the L7 firewall rule. One of: 'application', 'applicationCategory', 'host', 'port', 'ipRange'
	Type *string `json:"type,omitempty"`
	// The value of what needs to get blocked. Format of the value varies depending on type of the firewall rule selected.
	Value *string `json:"value,omitempty"`
}

// NewNetworksNetworkIdWirelessSsidsNumberFirewallL7FirewallRulesRules instantiates a new NetworksNetworkIdWirelessSsidsNumberFirewallL7FirewallRulesRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessSsidsNumberFirewallL7FirewallRulesRules() *NetworksNetworkIdWirelessSsidsNumberFirewallL7FirewallRulesRules {
	this := NetworksNetworkIdWirelessSsidsNumberFirewallL7FirewallRulesRules{}
	return &this
}

// NewNetworksNetworkIdWirelessSsidsNumberFirewallL7FirewallRulesRulesWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberFirewallL7FirewallRulesRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessSsidsNumberFirewallL7FirewallRulesRulesWithDefaults() *NetworksNetworkIdWirelessSsidsNumberFirewallL7FirewallRulesRules {
	this := NetworksNetworkIdWirelessSsidsNumberFirewallL7FirewallRulesRules{}
	return &this
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberFirewallL7FirewallRulesRules) GetPolicy() string {
	if o == nil || isNil(o.Policy) {
		var ret string
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberFirewallL7FirewallRulesRules) GetPolicyOk() (*string, bool) {
	if o == nil || isNil(o.Policy) {
    return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberFirewallL7FirewallRulesRules) HasPolicy() bool {
	if o != nil && !isNil(o.Policy) {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given string and assigns it to the Policy field.
func (o *NetworksNetworkIdWirelessSsidsNumberFirewallL7FirewallRulesRules) SetPolicy(v string) {
	o.Policy = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberFirewallL7FirewallRulesRules) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberFirewallL7FirewallRulesRules) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberFirewallL7FirewallRulesRules) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NetworksNetworkIdWirelessSsidsNumberFirewallL7FirewallRulesRules) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberFirewallL7FirewallRulesRules) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberFirewallL7FirewallRulesRules) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
    return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberFirewallL7FirewallRulesRules) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *NetworksNetworkIdWirelessSsidsNumberFirewallL7FirewallRulesRules) SetValue(v string) {
	o.Value = &v
}

func (o NetworksNetworkIdWirelessSsidsNumberFirewallL7FirewallRulesRules) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Policy) {
		toSerialize["policy"] = o.Policy
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessSsidsNumberFirewallL7FirewallRulesRules struct {
	value *NetworksNetworkIdWirelessSsidsNumberFirewallL7FirewallRulesRules
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberFirewallL7FirewallRulesRules) Get() *NetworksNetworkIdWirelessSsidsNumberFirewallL7FirewallRulesRules {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberFirewallL7FirewallRulesRules) Set(val *NetworksNetworkIdWirelessSsidsNumberFirewallL7FirewallRulesRules) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberFirewallL7FirewallRulesRules) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberFirewallL7FirewallRulesRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessSsidsNumberFirewallL7FirewallRulesRules(val *NetworksNetworkIdWirelessSsidsNumberFirewallL7FirewallRulesRules) *NullableNetworksNetworkIdWirelessSsidsNumberFirewallL7FirewallRulesRules {
	return &NullableNetworksNetworkIdWirelessSsidsNumberFirewallL7FirewallRulesRules{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberFirewallL7FirewallRulesRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberFirewallL7FirewallRulesRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


