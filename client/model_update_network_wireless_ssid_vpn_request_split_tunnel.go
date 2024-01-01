/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 18 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.41.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the UpdateNetworkWirelessSsidVpnRequestSplitTunnel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkWirelessSsidVpnRequestSplitTunnel{}

// UpdateNetworkWirelessSsidVpnRequestSplitTunnel The VPN split tunnel settings for this SSID.
type UpdateNetworkWirelessSsidVpnRequestSplitTunnel struct {
	// If true, VPN split tunnel is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// List of VPN split tunnel rules.
	Rules []UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner `json:"rules,omitempty"`
}

// NewUpdateNetworkWirelessSsidVpnRequestSplitTunnel instantiates a new UpdateNetworkWirelessSsidVpnRequestSplitTunnel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidVpnRequestSplitTunnel() *UpdateNetworkWirelessSsidVpnRequestSplitTunnel {
	this := UpdateNetworkWirelessSsidVpnRequestSplitTunnel{}
	return &this
}

// NewUpdateNetworkWirelessSsidVpnRequestSplitTunnelWithDefaults instantiates a new UpdateNetworkWirelessSsidVpnRequestSplitTunnel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidVpnRequestSplitTunnelWithDefaults() *UpdateNetworkWirelessSsidVpnRequestSplitTunnel {
	this := UpdateNetworkWirelessSsidVpnRequestSplitTunnel{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidVpnRequestSplitTunnel) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidVpnRequestSplitTunnel) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidVpnRequestSplitTunnel) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateNetworkWirelessSsidVpnRequestSplitTunnel) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidVpnRequestSplitTunnel) GetRules() []UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner {
	if o == nil || IsNil(o.Rules) {
		var ret []UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidVpnRequestSplitTunnel) GetRulesOk() ([]UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidVpnRequestSplitTunnel) HasRules() bool {
	if o != nil && !IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner and assigns it to the Rules field.
func (o *UpdateNetworkWirelessSsidVpnRequestSplitTunnel) SetRules(v []UpdateNetworkWirelessSsidVpnRequestSplitTunnelRulesInner) {
	o.Rules = v
}

func (o UpdateNetworkWirelessSsidVpnRequestSplitTunnel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkWirelessSsidVpnRequestSplitTunnel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	return toSerialize, nil
}

type NullableUpdateNetworkWirelessSsidVpnRequestSplitTunnel struct {
	value *UpdateNetworkWirelessSsidVpnRequestSplitTunnel
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidVpnRequestSplitTunnel) Get() *UpdateNetworkWirelessSsidVpnRequestSplitTunnel {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidVpnRequestSplitTunnel) Set(val *UpdateNetworkWirelessSsidVpnRequestSplitTunnel) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidVpnRequestSplitTunnel) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidVpnRequestSplitTunnel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidVpnRequestSplitTunnel(val *UpdateNetworkWirelessSsidVpnRequestSplitTunnel) *NullableUpdateNetworkWirelessSsidVpnRequestSplitTunnel {
	return &NullableUpdateNetworkWirelessSsidVpnRequestSplitTunnel{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidVpnRequestSplitTunnel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidVpnRequestSplitTunnel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


