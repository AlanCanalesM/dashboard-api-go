/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 07 February, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.43.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner{}

// GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner struct for GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner
type GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner struct {
	// Description of the rule (optional)
	Comment *string `json:"comment,omitempty"`
	// 'allow' or 'deny' traffic specified by this rule
	Policy *string `json:"policy,omitempty"`
	// The type of protocol (must be 'tcp', 'udp', 'icmp', 'icmp6' or 'any')
	Protocol *string `json:"protocol,omitempty"`
	// Comma-separated list of source port(s) (integer in the range 1-65535), or 'any'
	SrcPort *string `json:"srcPort,omitempty"`
	// Comma-separated list of source IP address(es) (in IP or CIDR notation), or 'any' (note: FQDN not supported for source addresses)
	SrcCidr *string `json:"srcCidr,omitempty"`
	// Comma-separated list of destination port(s) (integer in the range 1-65535), or 'any'
	DestPort *string `json:"destPort,omitempty"`
	// Comma-separated list of destination IP address(es) (in IP or CIDR notation), fully-qualified domain names (FQDN) or 'any'
	DestCidr *string `json:"destCidr,omitempty"`
	// Log this rule to syslog (true or false, boolean value) - only applicable if a syslog has been configured (optional)
	SyslogEnabled *bool `json:"syslogEnabled,omitempty"`
}

// NewGetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner instantiates a new GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner() *GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner {
	this := GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner{}
	return &this
}

// NewGetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInnerWithDefaults instantiates a new GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInnerWithDefaults() *GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner {
	this := GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner) SetComment(v string) {
	o.Comment = &v
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner) GetPolicy() string {
	if o == nil || IsNil(o.Policy) {
		var ret string
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner) GetPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.Policy) {
		return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner) HasPolicy() bool {
	if o != nil && !IsNil(o.Policy) {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given string and assigns it to the Policy field.
func (o *GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner) SetPolicy(v string) {
	o.Policy = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner) SetProtocol(v string) {
	o.Protocol = &v
}

// GetSrcPort returns the SrcPort field value if set, zero value otherwise.
func (o *GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner) GetSrcPort() string {
	if o == nil || IsNil(o.SrcPort) {
		var ret string
		return ret
	}
	return *o.SrcPort
}

// GetSrcPortOk returns a tuple with the SrcPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner) GetSrcPortOk() (*string, bool) {
	if o == nil || IsNil(o.SrcPort) {
		return nil, false
	}
	return o.SrcPort, true
}

// HasSrcPort returns a boolean if a field has been set.
func (o *GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner) HasSrcPort() bool {
	if o != nil && !IsNil(o.SrcPort) {
		return true
	}

	return false
}

// SetSrcPort gets a reference to the given string and assigns it to the SrcPort field.
func (o *GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner) SetSrcPort(v string) {
	o.SrcPort = &v
}

// GetSrcCidr returns the SrcCidr field value if set, zero value otherwise.
func (o *GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner) GetSrcCidr() string {
	if o == nil || IsNil(o.SrcCidr) {
		var ret string
		return ret
	}
	return *o.SrcCidr
}

// GetSrcCidrOk returns a tuple with the SrcCidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner) GetSrcCidrOk() (*string, bool) {
	if o == nil || IsNil(o.SrcCidr) {
		return nil, false
	}
	return o.SrcCidr, true
}

// HasSrcCidr returns a boolean if a field has been set.
func (o *GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner) HasSrcCidr() bool {
	if o != nil && !IsNil(o.SrcCidr) {
		return true
	}

	return false
}

// SetSrcCidr gets a reference to the given string and assigns it to the SrcCidr field.
func (o *GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner) SetSrcCidr(v string) {
	o.SrcCidr = &v
}

// GetDestPort returns the DestPort field value if set, zero value otherwise.
func (o *GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner) GetDestPort() string {
	if o == nil || IsNil(o.DestPort) {
		var ret string
		return ret
	}
	return *o.DestPort
}

// GetDestPortOk returns a tuple with the DestPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner) GetDestPortOk() (*string, bool) {
	if o == nil || IsNil(o.DestPort) {
		return nil, false
	}
	return o.DestPort, true
}

// HasDestPort returns a boolean if a field has been set.
func (o *GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner) HasDestPort() bool {
	if o != nil && !IsNil(o.DestPort) {
		return true
	}

	return false
}

// SetDestPort gets a reference to the given string and assigns it to the DestPort field.
func (o *GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner) SetDestPort(v string) {
	o.DestPort = &v
}

// GetDestCidr returns the DestCidr field value if set, zero value otherwise.
func (o *GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner) GetDestCidr() string {
	if o == nil || IsNil(o.DestCidr) {
		var ret string
		return ret
	}
	return *o.DestCidr
}

// GetDestCidrOk returns a tuple with the DestCidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner) GetDestCidrOk() (*string, bool) {
	if o == nil || IsNil(o.DestCidr) {
		return nil, false
	}
	return o.DestCidr, true
}

// HasDestCidr returns a boolean if a field has been set.
func (o *GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner) HasDestCidr() bool {
	if o != nil && !IsNil(o.DestCidr) {
		return true
	}

	return false
}

// SetDestCidr gets a reference to the given string and assigns it to the DestCidr field.
func (o *GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner) SetDestCidr(v string) {
	o.DestCidr = &v
}

// GetSyslogEnabled returns the SyslogEnabled field value if set, zero value otherwise.
func (o *GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner) GetSyslogEnabled() bool {
	if o == nil || IsNil(o.SyslogEnabled) {
		var ret bool
		return ret
	}
	return *o.SyslogEnabled
}

// GetSyslogEnabledOk returns a tuple with the SyslogEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner) GetSyslogEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.SyslogEnabled) {
		return nil, false
	}
	return o.SyslogEnabled, true
}

// HasSyslogEnabled returns a boolean if a field has been set.
func (o *GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner) HasSyslogEnabled() bool {
	if o != nil && !IsNil(o.SyslogEnabled) {
		return true
	}

	return false
}

// SetSyslogEnabled gets a reference to the given bool and assigns it to the SyslogEnabled field.
func (o *GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner) SetSyslogEnabled(v bool) {
	o.SyslogEnabled = &v
}

func (o GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.Policy) {
		toSerialize["policy"] = o.Policy
	}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.SrcPort) {
		toSerialize["srcPort"] = o.SrcPort
	}
	if !IsNil(o.SrcCidr) {
		toSerialize["srcCidr"] = o.SrcCidr
	}
	if !IsNil(o.DestPort) {
		toSerialize["destPort"] = o.DestPort
	}
	if !IsNil(o.DestCidr) {
		toSerialize["destCidr"] = o.DestCidr
	}
	if !IsNil(o.SyslogEnabled) {
		toSerialize["syslogEnabled"] = o.SyslogEnabled
	}
	return toSerialize, nil
}

type NullableGetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner struct {
	value *GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner
	isSet bool
}

func (v NullableGetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner) Get() *GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner {
	return v.value
}

func (v *NullableGetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner) Set(val *GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner(val *GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner) *NullableGetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner {
	return &NullableGetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner{value: val, isSet: true}
}

func (v NullableGetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


