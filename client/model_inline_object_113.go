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

// InlineObject113 struct for InlineObject113
type InlineObject113 struct {
	// The mac address of the trusted server being added
	Mac string `json:"mac"`
	// The VLAN of the trusted server being added. It must be between 1 and 4094
	Vlan int32 `json:"vlan"`
	Ipv4 NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41 `json:"ipv4"`
}

// NewInlineObject113 instantiates a new InlineObject113 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject113(mac string, vlan int32, ipv4 NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41) *InlineObject113 {
	this := InlineObject113{}
	this.Mac = mac
	this.Vlan = vlan
	this.Ipv4 = ipv4
	return &this
}

// NewInlineObject113WithDefaults instantiates a new InlineObject113 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject113WithDefaults() *InlineObject113 {
	this := InlineObject113{}
	return &this
}

// GetMac returns the Mac field value
func (o *InlineObject113) GetMac() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mac
}

// GetMacOk returns a tuple with the Mac field value
// and a boolean to check if the value has been set.
func (o *InlineObject113) GetMacOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Mac, true
}

// SetMac sets field value
func (o *InlineObject113) SetMac(v string) {
	o.Mac = v
}

// GetVlan returns the Vlan field value
func (o *InlineObject113) GetVlan() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value
// and a boolean to check if the value has been set.
func (o *InlineObject113) GetVlanOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Vlan, true
}

// SetVlan sets field value
func (o *InlineObject113) SetVlan(v int32) {
	o.Vlan = v
}

// GetIpv4 returns the Ipv4 field value
func (o *InlineObject113) GetIpv4() NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41 {
	if o == nil {
		var ret NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41
		return ret
	}

	return o.Ipv4
}

// GetIpv4Ok returns a tuple with the Ipv4 field value
// and a boolean to check if the value has been set.
func (o *InlineObject113) GetIpv4Ok() (*NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Ipv4, true
}

// SetIpv4 sets field value
func (o *InlineObject113) SetIpv4(v NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41) {
	o.Ipv4 = v
}

func (o InlineObject113) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mac"] = o.Mac
	}
	if true {
		toSerialize["vlan"] = o.Vlan
	}
	if true {
		toSerialize["ipv4"] = o.Ipv4
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject113 struct {
	value *InlineObject113
	isSet bool
}

func (v NullableInlineObject113) Get() *InlineObject113 {
	return v.value
}

func (v *NullableInlineObject113) Set(val *InlineObject113) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject113) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject113) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject113(val *InlineObject113) *NullableInlineObject113 {
	return &NullableInlineObject113{value: val, isSet: true}
}

func (v NullableInlineObject113) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject113) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


