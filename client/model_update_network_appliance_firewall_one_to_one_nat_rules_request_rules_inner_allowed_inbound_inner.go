/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 06 September, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.37.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerAllowedInboundInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerAllowedInboundInner{}

// UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerAllowedInboundInner struct for UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerAllowedInboundInner
type UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerAllowedInboundInner struct {
	// Either of the following: 'tcp', 'udp', 'icmp-ping' or 'any'
	Protocol *string `json:"protocol,omitempty"`
	// An array of ports or port ranges that will be forwarded to the host on the LAN
	DestinationPorts []string `json:"destinationPorts,omitempty"`
	// An array of ranges of WAN IP addresses that are allowed to make inbound connections on the specified ports or port ranges, or 'any'
	AllowedIps []string `json:"allowedIps,omitempty"`
}

// NewUpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerAllowedInboundInner instantiates a new UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerAllowedInboundInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerAllowedInboundInner() *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerAllowedInboundInner {
	this := UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerAllowedInboundInner{}
	return &this
}

// NewUpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerAllowedInboundInnerWithDefaults instantiates a new UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerAllowedInboundInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerAllowedInboundInnerWithDefaults() *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerAllowedInboundInner {
	this := UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerAllowedInboundInner{}
	return &this
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerAllowedInboundInner) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerAllowedInboundInner) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerAllowedInboundInner) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerAllowedInboundInner) SetProtocol(v string) {
	o.Protocol = &v
}

// GetDestinationPorts returns the DestinationPorts field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerAllowedInboundInner) GetDestinationPorts() []string {
	if o == nil || IsNil(o.DestinationPorts) {
		var ret []string
		return ret
	}
	return o.DestinationPorts
}

// GetDestinationPortsOk returns a tuple with the DestinationPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerAllowedInboundInner) GetDestinationPortsOk() ([]string, bool) {
	if o == nil || IsNil(o.DestinationPorts) {
		return nil, false
	}
	return o.DestinationPorts, true
}

// HasDestinationPorts returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerAllowedInboundInner) HasDestinationPorts() bool {
	if o != nil && !IsNil(o.DestinationPorts) {
		return true
	}

	return false
}

// SetDestinationPorts gets a reference to the given []string and assigns it to the DestinationPorts field.
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerAllowedInboundInner) SetDestinationPorts(v []string) {
	o.DestinationPorts = v
}

// GetAllowedIps returns the AllowedIps field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerAllowedInboundInner) GetAllowedIps() []string {
	if o == nil || IsNil(o.AllowedIps) {
		var ret []string
		return ret
	}
	return o.AllowedIps
}

// GetAllowedIpsOk returns a tuple with the AllowedIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerAllowedInboundInner) GetAllowedIpsOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowedIps) {
		return nil, false
	}
	return o.AllowedIps, true
}

// HasAllowedIps returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerAllowedInboundInner) HasAllowedIps() bool {
	if o != nil && !IsNil(o.AllowedIps) {
		return true
	}

	return false
}

// SetAllowedIps gets a reference to the given []string and assigns it to the AllowedIps field.
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerAllowedInboundInner) SetAllowedIps(v []string) {
	o.AllowedIps = v
}

func (o UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerAllowedInboundInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerAllowedInboundInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.DestinationPorts) {
		toSerialize["destinationPorts"] = o.DestinationPorts
	}
	if !IsNil(o.AllowedIps) {
		toSerialize["allowedIps"] = o.AllowedIps
	}
	return toSerialize, nil
}

type NullableUpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerAllowedInboundInner struct {
	value *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerAllowedInboundInner
	isSet bool
}

func (v NullableUpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerAllowedInboundInner) Get() *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerAllowedInboundInner {
	return v.value
}

func (v *NullableUpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerAllowedInboundInner) Set(val *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerAllowedInboundInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerAllowedInboundInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerAllowedInboundInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerAllowedInboundInner(val *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerAllowedInboundInner) *NullableUpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerAllowedInboundInner {
	return &NullableUpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerAllowedInboundInner{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerAllowedInboundInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerAllowedInboundInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


