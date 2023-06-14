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

// UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner struct for UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner
type UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner struct {
	// A description of the rule
	Name *string `json:"name,omitempty"`
	// 'tcp' or 'udp'
	Protocol *string `json:"protocol,omitempty"`
	// Destination port of the traffic that is arriving on the WAN
	PublicPort *string `json:"publicPort,omitempty"`
	// Local IP address to which traffic will be forwarded
	LocalIp *string `json:"localIp,omitempty"`
	// Destination port of the forwarded traffic that will be sent from the MX to the specified host on the LAN. If you simply wish to forward the traffic without translating the port, this should be the same as the Public port
	LocalPort *string `json:"localPort,omitempty"`
	// Remote IP addresses or ranges that are permitted to access the internal resource via this port forwarding rule, or 'any'
	AllowedIps []string `json:"allowedIps,omitempty"`
}

// NewUpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner instantiates a new UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner() *UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner {
	this := UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner{}
	return &this
}

// NewUpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInnerWithDefaults instantiates a new UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInnerWithDefaults() *UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner {
	this := UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner) SetName(v string) {
	o.Name = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner) GetProtocol() string {
	if o == nil || isNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner) GetProtocolOk() (*string, bool) {
	if o == nil || isNil(o.Protocol) {
    return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner) HasProtocol() bool {
	if o != nil && !isNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner) SetProtocol(v string) {
	o.Protocol = &v
}

// GetPublicPort returns the PublicPort field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner) GetPublicPort() string {
	if o == nil || isNil(o.PublicPort) {
		var ret string
		return ret
	}
	return *o.PublicPort
}

// GetPublicPortOk returns a tuple with the PublicPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner) GetPublicPortOk() (*string, bool) {
	if o == nil || isNil(o.PublicPort) {
    return nil, false
	}
	return o.PublicPort, true
}

// HasPublicPort returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner) HasPublicPort() bool {
	if o != nil && !isNil(o.PublicPort) {
		return true
	}

	return false
}

// SetPublicPort gets a reference to the given string and assigns it to the PublicPort field.
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner) SetPublicPort(v string) {
	o.PublicPort = &v
}

// GetLocalIp returns the LocalIp field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner) GetLocalIp() string {
	if o == nil || isNil(o.LocalIp) {
		var ret string
		return ret
	}
	return *o.LocalIp
}

// GetLocalIpOk returns a tuple with the LocalIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner) GetLocalIpOk() (*string, bool) {
	if o == nil || isNil(o.LocalIp) {
    return nil, false
	}
	return o.LocalIp, true
}

// HasLocalIp returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner) HasLocalIp() bool {
	if o != nil && !isNil(o.LocalIp) {
		return true
	}

	return false
}

// SetLocalIp gets a reference to the given string and assigns it to the LocalIp field.
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner) SetLocalIp(v string) {
	o.LocalIp = &v
}

// GetLocalPort returns the LocalPort field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner) GetLocalPort() string {
	if o == nil || isNil(o.LocalPort) {
		var ret string
		return ret
	}
	return *o.LocalPort
}

// GetLocalPortOk returns a tuple with the LocalPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner) GetLocalPortOk() (*string, bool) {
	if o == nil || isNil(o.LocalPort) {
    return nil, false
	}
	return o.LocalPort, true
}

// HasLocalPort returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner) HasLocalPort() bool {
	if o != nil && !isNil(o.LocalPort) {
		return true
	}

	return false
}

// SetLocalPort gets a reference to the given string and assigns it to the LocalPort field.
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner) SetLocalPort(v string) {
	o.LocalPort = &v
}

// GetAllowedIps returns the AllowedIps field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner) GetAllowedIps() []string {
	if o == nil || isNil(o.AllowedIps) {
		var ret []string
		return ret
	}
	return o.AllowedIps
}

// GetAllowedIpsOk returns a tuple with the AllowedIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner) GetAllowedIpsOk() ([]string, bool) {
	if o == nil || isNil(o.AllowedIps) {
    return nil, false
	}
	return o.AllowedIps, true
}

// HasAllowedIps returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner) HasAllowedIps() bool {
	if o != nil && !isNil(o.AllowedIps) {
		return true
	}

	return false
}

// SetAllowedIps gets a reference to the given []string and assigns it to the AllowedIps field.
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner) SetAllowedIps(v []string) {
	o.AllowedIps = v
}

func (o UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !isNil(o.PublicPort) {
		toSerialize["publicPort"] = o.PublicPort
	}
	if !isNil(o.LocalIp) {
		toSerialize["localIp"] = o.LocalIp
	}
	if !isNil(o.LocalPort) {
		toSerialize["localPort"] = o.LocalPort
	}
	if !isNil(o.AllowedIps) {
		toSerialize["allowedIps"] = o.AllowedIps
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner struct {
	value *UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner
	isSet bool
}

func (v NullableUpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner) Get() *UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner {
	return v.value
}

func (v *NullableUpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner) Set(val *UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner(val *UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner) *NullableUpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner {
	return &NullableUpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


